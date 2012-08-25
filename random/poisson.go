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
	"log"
	"math"
	"sync"
)

const (
	SWITCH_MEAN = 10.0
)

// A Poisson distribution.
type Poisson struct {
	Mean float64

	randGenerator RandomEngine
	locker        sync.Mutex

	// precomputed and cached values (for performance only)
	// cached for < SWICH_MEAN
	my_old, p, q, p0 float64
	pp               []float64
	llll             int

	// cached for >= SWICH_MEAN
	my_last, ll                            float64
	k2, k4, k1, k5                         int
	dl, dr, r1, r2, r4, r5, lr, l_my, c_pm float64
	f1, f2, f4, f5, p1, p2, p3, p4, p5, p6 float64

	// cached for both
	m int
}

// NewPoisson returens a new Poisson with provided mean and random generator.
func NewPoisson(mean float64, randGenerator RandomEngine) (poisson *Poisson) {
	if mean < 0 {
		log.Fatalf("Poisson Mean should >= 0, but we got: %f\n", mean)
	}
	poisson = &Poisson{Mean: mean, randGenerator: randGenerator}
	poisson.my_old = -1.0
	poisson.my_last = -1.0
	poisson.pp = make([]float64, 36)

	return
}

// Cdf returns the cumulative distribution function.
func (poisson Poisson) Cdf(k int) (p float64) {
	return
}

// Pdf returns the probability distribution function.
func (poisson Poisson) Pdf(k int) (p float64) {
	p = math.Exp(float64(k)*math.Log(poisson.Mean) - specfunc.LogFactorial(k) - poisson.Mean)
	return
}

// Seed uses the provided seed value to initialize the generator to a deterministic state.
func (poisson *Poisson) Seed(seed int64) {
	poisson.locker.Lock()
	poisson.randGenerator.Seed(seed)
	poisson.locker.Unlock()
}

// Int returns a non-negative pseudo-random int from a poisson distribution.
func (poisson *Poisson) Int() (k int) {
	poisson.locker.Lock()
	k = poisson.random()
	poisson.locker.Unlock()
	return
}

// Generate poisson random number
// Implementation: High performance implementation.
// Patchwork Rejection/Inversion method.
// This is a port of Poisson.java from the colt java library, which is based upon
// H. Zechner (1994): Efficient sampling from continuous and discrete unimodal distributions,
// Doctoral Dissertation, 156 pp., Technical University Graz, Austria.
// Also see
// Stadlober E., H. Zechner (1999), The patchwork rejection method for sampling from unimodal distributions,
// to appear in ACM Transactions on Modelling and Simulation.
func (poisson *Poisson) random() (k int) {
	/******************************************************************
	 *                                                                *
	 * Poisson Distribution - Patchwork Rejection/Inversion           *
	 *                                                                *
	 ******************************************************************
	 *                                                                *
	 * For parameter  my < 10  Tabulated Inversion is applied.        *
	 * For my >= 10  Patchwork Rejection is employed:                 *
	 * The area below the histogram function f(x) is rearranged in    *
	 * its body by certain point reflections. Within a large center   *
	 * interval variates are sampled efficiently by rejection from    *
	 * uniform hats. Rectangular immediate acceptance regions speed   *
	 * up the generation. The remaining tails are covered by          *
	 * exponential functions.                                         *
	 *                                                                *
	 *****************************************************************/
	my := poisson.Mean

	if my < SWITCH_MEAN { // CASE B: Inversion- start new table and calculate p0
		if my != poisson.my_old {
			poisson.my_old = my
			poisson.llll = 0
			poisson.p = math.Exp(-my)
			poisson.q = poisson.p
			poisson.p0 = poisson.p
		}

		if my > 1.0 {
			poisson.m = int(my)
		} else {
			poisson.m = 1
		}

		for {
			u := poisson.randGenerator.Float64()
			if u <= poisson.p0 {
				return k
			}
			if poisson.llll != 0 {
				if u > 0.458 {
					if poisson.llll < poisson.m {
						k = poisson.llll
					} else {
						k = poisson.m
					}
				} else {
					k = 1
				}
				for ; k <= poisson.llll; k++ {
					if u <= poisson.pp[k] {
						return k
					}
				}
				if poisson.llll == 35 {
					continue
				}
			}
			for k = poisson.llll + 1; k <= 35; k++ {
				poisson.p *= my / float64(k)
				poisson.q += poisson.p
				poisson.pp[k] = poisson.q
				if u <= poisson.q {
					poisson.llll = k
					return k
				}
			}
			poisson.llll = 35
		}
	} else {
		var (
			Dk, X, Y    int
			Ds, U, V, W float64
		)

		poisson.m = int(my)
		if my != poisson.my_last {
			poisson.my_last = my

			// approximate deviation of reflection points k2, k4 from my - 1/2 
			Ds = math.Sqrt(my + 0.25)

			// mode m, reflection points k2 and k4, and points k1 and k5, which    
			// delimit the centre region of h(x)                                    
			poisson.k2 = int(math.Ceil(my - 0.5 - Ds))
			poisson.k4 = int(my - 0.5 + Ds)
			poisson.k1 = int(poisson.k2 + poisson.k2 - poisson.m + 1.0)
			poisson.k5 = int(poisson.k4 + poisson.k4 - poisson.m)

			// range width of the critical left and right centre region 
			poisson.dl = float64(poisson.k2 - poisson.k1)
			poisson.dr = float64(poisson.k5 - poisson.k4)

			// recurrence constants r(k) = p(k)/p(k-1) at k = k1, k2, k4+1, k5+1  
			poisson.r1 = my / float64(poisson.k1)
			poisson.r2 = my / float64(poisson.k2)
			poisson.r4 = my / float64(poisson.k4+1)
			poisson.r5 = my / float64(poisson.k5+1)

			// reciprocal values of the scale parameters of expon. tail envelopes   
			poisson.ll = math.Log(poisson.r1)  // expon. tail left 
			poisson.lr = -math.Log(poisson.r5) // expon. tail right

			// Poisson constants, necessary for computing function values f(k)      
			poisson.l_my = math.Log(my)
			poisson.c_pm = float64(poisson.m)*poisson.l_my - specfunc.LogFactorial(poisson.m)

			// function values f(k) = p(k)/p(m) at k = k2, k4, k1, k5               
			poisson.f2 = poisson.f(poisson.k2, poisson.l_my, poisson.c_pm)
			poisson.f4 = poisson.f(poisson.k4, poisson.l_my, poisson.c_pm)
			poisson.f1 = poisson.f(poisson.k1, poisson.l_my, poisson.c_pm)
			poisson.f5 = poisson.f(poisson.k5, poisson.l_my, poisson.c_pm)

			// area of the two centre and the two exponential tail regions          
			// area of the two immediate acceptance regions between k2, k4         
			poisson.p1 = poisson.f2 * (poisson.dl + 1.0)          // immed. left    
			poisson.p2 = poisson.f2*poisson.dl + poisson.p1       // centre left    
			poisson.p3 = poisson.f4*(poisson.dr+1.0) + poisson.p2 // immed. right     
			poisson.p4 = poisson.f4*poisson.dr + poisson.p3       // centre right     
			poisson.p5 = poisson.f1/poisson.ll + poisson.p4       // expon. tail left 
			poisson.p6 = poisson.f5/poisson.lr + poisson.p5       // expon. tail right
		} // end set-up

		for {
			// generate uniform number U -- U(0, p6)                                
			// case distinction corresponding to U 
			U = poisson.randGenerator.Float64() * poisson.p6
			if U < poisson.p2 {
				// immediate acceptance region R2 = [k2, m) *[0, f2),  X = k2, ... m -1 
				V = U - poisson.p1
				if V < 0.0 {
					k = poisson.k2 + int(U/poisson.f2)
					return
				}
				// immediate acceptance region R1 = [k1, k2)*[0, f1),  X = k1, ... k2-1 
				W = V / poisson.dl
				if W < poisson.f1 {
					k = poisson.k1 + int(V/poisson.f1)
					return
				}

				// computation of candidate X < k2, and its counterpart Y > k2          
				// either squeeze-acceptance of X or acceptance-rejection of Y         
				Dk = int(poisson.dl*poisson.randGenerator.Float64()) + 1
				if W <= poisson.f2-float64(Dk)*(poisson.f2-poisson.f2/poisson.r2) {
					k = poisson.k2 - Dk
					return
				}
				V = poisson.f2 + poisson.f2 - W
				if V < 1.0 {
					Y = poisson.k2 + Dk
					if V < poisson.f2+float64(Dk)*(1.0-poisson.f2)/(poisson.dl+1.0) {
						k = Y
						return
					}
					if V <= poisson.f(Y, poisson.l_my, poisson.c_pm) {
						k = Y
						return
					}
				}
				X = poisson.k2 - Dk
			} else if U < poisson.p4 { // centre right
				// immediate acceptance region R3 = [m, k4+1)*[0, f4), X = m, ... k4
				V = U - poisson.p3
				if V < 0.0 {
					k = poisson.k4 - int((U-poisson.p2)/poisson.f4)
					return
				}
				// immediate acceptance region R4 = [k4+1, k5+1)*[0, f5) 
				W = V / poisson.dr
				if W < poisson.f5 {
					k = poisson.k5 - int(V/poisson.f5)
					return
				}

				// computation of candidate X > k4, and its counterpart Y < k4          
				// either squeeze-acceptance of X or acceptance-rejection of Y 
				Dk = int(poisson.dr*poisson.randGenerator.Float64()) + 1
				if W <= poisson.f4-float64(Dk)*(poisson.f4-poisson.f4*poisson.f4) {
					k = poisson.k4 + Dk
					return
				}
				V = poisson.f4 + poisson.f4 - W
				if V < 1.0 {
					Y = poisson.k4 - Dk
					if V <= poisson.f4+float64(Dk)*(1.0-poisson.f4)/poisson.dr {
						k = Y
						return
					}
					if V <= poisson.f(Y, poisson.l_my, poisson.c_pm) {
						k = Y
						return
					}
				}
				X = poisson.k4 + Dk
			} else {
				W = poisson.randGenerator.Float64()
				if U < poisson.p5 {
					Dk = int(1.0 - math.Log(W)/poisson.ll)
					X = poisson.k1 - Dk
					if X < 0 {
						continue
					}
					W *= (U - poisson.p4) * poisson.ll
					if W <= poisson.f1-float64(Dk)*(poisson.f1-poisson.f1/poisson.r1) {
						k = X
						return
					}
				} else {
					Dk = int(1.0 - math.Log(W)/poisson.lr)
					X = poisson.k5 + Dk
					W *= (U - poisson.p5) * poisson.lr
					if W <= poisson.f5-float64(Dk)*(poisson.f5-poisson.f5*poisson.f5) {
						k = X
						return
					}
				}
			}

			// acceptance-rejection test of candidate X from the original area   
			// test, whether  W <= f(k),    with  W = U*h(x)  and  U -- U(0, 1)  
			// log f(X) = (X - m)*log(my) - log X! + log m! 
			if math.Log(W) <= float64(X)*poisson.l_my-specfunc.LogFactorial(X)-poisson.c_pm {
				k = X
				return
			}
		}
	}
	return
}

// Helper function
func (poisson Poisson) f(k int, l_nu, c_pm float64) float64 {
	return math.Exp(float64(k)*l_nu - specfunc.LogFactorial(k) - c_pm)
}
