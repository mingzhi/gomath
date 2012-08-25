//   Copyright (C) 2012 Mingzhi Lin
//
// Permission is hereby granted, free of charge, to any person obtaining 
// a copy of this software and associated documentation files (the "Software"), 
// to deal in the Software without restriction, including without limitation 
// the rights to use, copy, modify, merge, publish, distribute, sublicense, 
// and/or sell copies of the Software, and to permit persons to whom the 
// Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included 
// in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS 
// OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, 
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE 
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER 
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, 
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE 
// SOFTWARE.

package random

import (
	"github.com/mingzhi/gomath/specfunc"
	"math"
	"sync"
)

// A Binomial distribution.
// p(x) = k * p^k * (1-p)^(n-k) with k = n! / (k! * (n-k)!)
type Binomial struct {
	N int
	P float64

	randomGenerator RandomEngine
	locker          sync.Mutex

	// cache vars for method generateBinomial(...)
	n_last, n_prev                                        int
	par, np, p0, q, p_last, p_prev                        float64
	b, m, nm                                              int
	pq, rc, ss, xm, xl, xr, ll, lr, c, p1, p2, p3, p4, ch float64

	// cache vars for method Pdf(...)
	log_p, log_q, log_n float64
}

// NewBinomial returns a Binomial distribution.
func NewBinomial(n int, p float64, randomGenerator RandomEngine) *Binomial {
	if float64(n)*math.Min(p, 1.0-p) <= 0 {
		panic("Illegal argument for Binomial distribution: n*p <= 0")
	}
	binomial := &Binomial{N: n, P: p, randomGenerator: randomGenerator}

	binomial.n_last = -1
	binomial.n_prev = -1
	binomial.p_last = -1.0
	binomial.p_prev = -1.0

	binomial.log_p = math.Log(binomial.P)
	binomial.log_q = math.Log(1.0 - binomial.P)
	binomial.log_n = specfunc.LogFactorial(binomial.N)

	return binomial
}

// Cdf returns the cumulative distribution function.
func (binomial Binomial) Cdf(k int) (p float64) {
	panic("have not implemented")
	return
}

// Int returns a random number from the Binomial distribution.
func (binomial Binomial) Int() (k int) {
	binomial.locker.Lock()
	k = binomial.random()
	binomial.locker.Unlock()
	return
}

// Pdf returns the probability distribution function.
func (binomial Binomial) Pdf(k int) (p float64) {
	r := binomial.N - k
	p = math.Exp(binomial.log_n - specfunc.LogFactorial(k) - specfunc.LogFactorial(r) + binomial.log_p*float64(k) + binomial.log_q*float64(r))
	return
}

// Generate a random number from the Binomial distribution.
// This is a port of Biomial.java from Colt java library, which is based on
// V. Kachitvichyanukul, B.W. Schmeiser (1988): Binomial random variate generation, Communications of the ACM 31, 216-222.
func (binomial *Binomial) random() int {
	/******************************************************************
	 *                                                                *
	 *     Binomial-Distribution - Acceptance Rejection/Inversion     *
	 *                                                                *
	 ******************************************************************
	 *                                                                *
	 * Acceptance Rejection method combined with Inversion for        *
	 * generating Binomial random numbers with parameters             *
	 * n (number of trials) and p (probability of success).           *
	 * For  min(n*p,n*(1-p)) < 10  the Inversion method is applied:   *
	 * The random numbers are generated via sequential search,        *
	 * starting at the lowest index k=0. The cumulative probabilities *
	 * are avoided by using the technique of chop-down.               *
	 * For  min(n*p,n*(1-p)) >= 10  Acceptance Rejection is used:     *
	 * The algorithm is based on a hat-function which is uniform in   *
	 * the centre region and exponential in the tails.                *
	 * A triangular immediate acceptance region in the centre speeds  *
	 * up the generation of binomial variates.                        *
	 * If candidate k is near the mode, f(k) is computed recursively  *
	 * starting at the mode m.                                        *
	 * The acceptance test by Stirling's formula is modified          *
	 * according to W. Hoermann (1992): The generation of binomial    *
	 * random variates, to appear in J. Statist. Comput. Simul.       *
	 * If  p < .5  the algorithm is applied to parameters n, p.       *
	 * Otherwise p is replaced by 1-p, and k is replaced by n - k.    *
	 *                                                                *
	 ******************************************************************
	 *                                                                *
	 * FUNCTION:    - samples a random number from the binomial       *
	 *                distribution with parameters n and p  and is    *
	 *                valid for  n*min(p,1-p)  >  0.                  *
	 * REFERENCE:   - V. Kachitvichyanukul, B.W. Schmeiser (1988):    *
	 *                Binomial random variate generation,             *
	 *                Communications of the ACM 31, 216-222.          *
	 * SUBPROGRAMS: - StirlingCorrection()                            *
	 *                            ... Correction term of the Stirling *
	 *                                approximation for log(k!)       *
	 *                                (series in 1/k or table values  *
	 *                                for small k) with long int k    *
	 *              - randomGenerator    ... (0,1)-Uniform engine     * 
	 *                                                                *
	 ******************************************************************/

	C1_3 := 0.33333333333333333
	C5_8 := 0.62500000000000000
	C1_6 := 0.16666666666666667
	DMAX_KM := 20

	var (
		bh, i, K, Km, nK     int
		f, rm, U, V, X, T, E float64
	)

	if binomial.N != binomial.n_last || binomial.P != binomial.p_last { // set-up
		binomial.n_last = binomial.N
		binomial.p_last = binomial.P
		binomial.par = math.Min(binomial.P, 1.0-binomial.P)
		binomial.q = 1.0 - binomial.par
		binomial.np = float64(binomial.N) * binomial.par

		// check for invaid input values
		if binomial.np <= 0.0 {
			return -1
		}

		rm = float64(binomial.np) + binomial.par
		binomial.m = int(rm)
		if binomial.np < 10 {
			binomial.p0 = math.Exp(float64(binomial.N) * math.Log(binomial.q))
			bh = int(binomial.np + 10.0*math.Sqrt(binomial.np*binomial.q))
			if binomial.N < bh {
				binomial.b = binomial.N
			} else {
				binomial.b = bh
			}
		} else {
			binomial.pq = binomial.par / binomial.q
			binomial.rc = float64(binomial.N+1) * binomial.pq
			binomial.ss = binomial.np + binomial.q
			i = int(2.195*math.Sqrt(binomial.ss) - 4.6*binomial.q)
			binomial.xm = float64(binomial.m) + 0.5
			binomial.xl = float64(binomial.m - i)
			binomial.xr = float64(binomial.m + i + 1)
			f = (rm - binomial.xl) / (rm - binomial.xl*binomial.par)
			binomial.ll = f * (1.0 + 0.5*f)
			f = (binomial.xr - rm) / (binomial.xr * binomial.q)
			binomial.lr = f * (1.0 + 0.5*f)
			binomial.c = 0.134 + 20.5/(15.3+float64(binomial.m))

			binomial.p1 = float64(i) + 0.5
			binomial.p2 = binomial.p1 * (1.0 + binomial.c + binomial.c) // probabilities
			binomial.p3 = binomial.p2 + binomial.c/binomial.ll          // of regions 1-4
			binomial.p4 = binomial.p3 + binomial.c/binomial.lr
		}
	}

	if binomial.np < 10 {
		var pk float64
		K = 0
		pk = binomial.p0
		U = binomial.randomGenerator.Float64()
		for U > pk {
			K++
			if K > binomial.b {
				U = binomial.randomGenerator.Float64()
				K = 0
				pk = binomial.p0
			} else {
				U -= pk
				pk = float64((float64(binomial.N-K+1) * binomial.par * pk) / (float64(K) * binomial.q))
			}
		}

		if binomial.P > 0.5 {
			return binomial.N - K
		} else {
			return K
		}
	}

	for {
		V = binomial.randomGenerator.Float64()
		U = binomial.randomGenerator.Float64() * binomial.p4
		if U <= binomial.p1 {
			K = int(binomial.xm - U + binomial.p1*V)
			if binomial.P > 0.5 {
				return binomial.N - K
			} else {
				return K
			}
		}
		if U <= binomial.p2 {
			X = binomial.xl + (U-binomial.p1)/binomial.c
			V = V*binomial.c + 1.0 - math.Abs(binomial.xm-X)/binomial.p1
			if V >= 1.0 {
				continue
			}
			K = int(X)
		} else if U <= binomial.p3 {
			X = binomial.xl + math.Log(V)/binomial.ll
			if X < 0.0 {
				continue
			}
			K = int(X)
			V *= (U - binomial.p2) * binomial.ll
		} else {
			K = int(binomial.xr - math.Log(V)/binomial.lr)
			if K > binomial.N {
				continue
			}
			V *= (U - binomial.p3) * binomial.lr
		}

		// acceptance test :  two cases, depending on |K - m|
		if K > binomial.m {
			Km = K - binomial.m
		} else {
			Km = binomial.m - K
		}
		if Km <= DMAX_KM || float64(Km+Km+2) >= binomial.ss {
			f = 1.0
			if binomial.m < K {
				for i := binomial.m; i < K; {
					i++
					f *= (binomial.rc/float64(i) - binomial.pq)
					if f < V {
						break
					}
				}
			} else {
				for i := K; i < binomial.m; {
					i++
					V *= (binomial.rc/float64(i) - binomial.pq)
					if V > f {
						break
					}
				}
			}
			if V <= f {
				break
			}
		} else {
			V = math.Log(V)
			T = float64(-Km*Km) / (binomial.ss + binomial.ss)
			E = (float64(Km) / binomial.ss) * ((float64(Km)*(float64(Km)*C1_3+C5_8)+C1_6)/binomial.ss + 0.5)
			if V <= T-E {
				break
			}
			if V <= T+E {
				if binomial.N != binomial.n_prev || binomial.par != binomial.p_prev {
					binomial.n_prev = binomial.N
					binomial.p_prev = binomial.par

					binomial.nm = binomial.N - binomial.m + 1
					binomial.ch = binomial.xm*math.Log(float64(binomial.m+1.0)/(binomial.pq*float64(binomial.nm))) + specfunc.StirlingCorrection(binomial.m+1) - specfunc.StirlingCorrection(binomial.nm)
				}
				nK = binomial.N - K + 1

				if V <= binomial.ch+(float64(binomial.N)+1.0)*math.Log(float64(binomial.nm)/float64(nK))+(float64(K)+0.5)*math.Log(float64(nK)*binomial.pq/(float64(K)+1.0))-specfunc.StirlingCorrection(K+1)-specfunc.StirlingCorrection(nK) {
					break
				}
			}
		}
	}
	var result int
	if binomial.P > 0.5 {
		result = binomial.N - K
	} else {
		result = K
	}
	return result
}
