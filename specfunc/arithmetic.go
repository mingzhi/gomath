/*
 *   Copyright (C) 2012 Mingzhi Lin
 *
 * Permission is hereby granted, free of charge, to any person obtaining 
 * a copy of this software and associated documentation files (the "Software"), 
 * to deal in the Software without restriction, including without limitation 
 * the rights to use, copy, modify, merge, publish, distribute, sublicense, 
 * and/or sell copies of the Software, and to permit persons to whom the 
 * Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included 
 * in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS 
 * OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, 
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE 
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER 
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, 
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE 
 * SOFTWARE.
 */

package specfunc

import (
	"math"
)

// LogFactorial returns log(k!).
// Tries to avoid overflows.
// For k<30 simply looks up a table in O(1).
// For k>=30 uses stirlings approximation.
// This is a port from Arithmetic.java from Colt Java library.
func LogFactorial(k int) (f float64) {
	logFactorials := []float64{
		0.00000000000000000, 0.00000000000000000, 0.69314718055994531,
		1.79175946922805500, 3.17805383034794562, 4.78749174278204599,
		6.57925121201010100, 8.52516136106541430, 10.60460290274525023,
		12.80182748008146961, 15.10441257307551530, 17.50230784587388584,
		19.98721449566188615, 22.55216385312342289, 25.19122118273868150,
		27.89927138384089157, 30.67186010608067280, 33.50507345013688888,
		36.39544520803305358, 39.33988418719949404, 42.33561646075348503,
		45.38013889847690803, 48.47118135183522388, 51.60667556776437357,
		54.78472939811231919, 58.00360522298051994, 61.26170176100200198,
		64.55753862700633106, 67.88974313718153498, 71.25703896716800901,
	}
	if k >= 30 {
		C0 := 9.18938533204672742e-01
		C1 := 8.33333333333333333e-02
		C3 := -2.77777777777777778e-03
		C5 := 7.93650793650793651e-04
		C7 := -5.95238095238095238e-04

		r := 1.0 / float64(k)
		rr := r * r
		f = (float64(k)+0.5)*math.Log(float64(k)) - float64(k) + C0 + r*(C1+rr*(C3+rr*(C5+rr*C7)))
	} else {
		f = logFactorials[k]
	}
	return
}

// StirlingCorrection returns the StirlingCorrection
// Correction term of the Stirling approximation for log(k!)
// (series in 1/k, or table values for small k)                         
// with int parameter k. 
// log k! = (k + 1/2)log(k + 1) - (k + 1) + (1/2)log(2Pi) +
//          stirlingCorrection(k + 1)                                       
// log k! = (k + 1/2)log(k)     -  k      + (1/2)log(2Pi) +              
//          stirlingCorrection(k)
// This is a port of Arithmetic.java from Colt java library.
func StirlingCorrection(k int) (sc float64) {
	stirlingCorrection := []float64{
		0.0,
		8.106146679532726e-02, 4.134069595540929e-02,
		2.767792568499834e-02, 2.079067210376509e-02,
		1.664469118982119e-02, 1.387612882307075e-02,
		1.189670994589177e-02, 1.041126526197209e-02,
		9.255462182712733e-03, 8.330563433362871e-03,
		7.573675487951841e-03, 6.942840107209530e-03,
		6.408994188004207e-03, 5.951370112758848e-03,
		5.554733551962801e-03, 5.207655919609640e-03,
		4.901395948434738e-03, 4.629153749334029e-03,
		4.385560249232324e-03, 4.166319691996922e-03,
		3.967954218640860e-03, 3.787618068444430e-03,
		3.622960224683090e-03, 3.472021382978770e-03,
		3.333155636728090e-03, 3.204970228055040e-03,
		3.086278682608780e-03, 2.976063983550410e-03,
		2.873449362352470e-03, 2.777674929752690e-03,
	}

	C1 := 8.33333333333333333e-02  //  +1/12 
	C3 := -2.77777777777777778e-03 //  -1/360
	C5 := 7.93650793650793651e-04  //  +1/1260
	C7 := -5.95238095238095238e-04 //  -1/1680

	var r, rr float64

	if k > 30 {
		r = 1.0 / float64(k)
		rr = r * r
		sc = r * (C1 + rr*(C3+rr*(C5+rr*C7)))
	} else {
		sc = stirlingCorrection[k]
	}

	return
}
