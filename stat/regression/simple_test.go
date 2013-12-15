package regression

import (
	"math"
	"testing"
)

var (
	norrisData [][]float64
	corrData   [][]float64
	infData    [][]float64
)

const TOLERANCE float64 = 10E-12
const TOLERANCE9 float64 = 10E-9

func init() {
	/*
	 * NIST "Norris" refernce data set from
	 * http://www.itl.nist.gov/div898/strd/lls/data/LINKS/DATA/Norris.dat
	 * Strangely, order is {y,x}
	 */
	norrisData = [][]float64{{0.1, 0.2}, {338.8, 337.4}, {118.1, 118.2},
		{888.0, 884.6}, {9.2, 10.1}, {228.1, 226.5}, {668.5, 666.3}, {998.5, 996.3},
		{449.1, 448.6}, {778.9, 777.0}, {559.2, 558.2}, {0.3, 0.4}, {0.1, 0.6}, {778.1, 775.5},
		{668.8, 666.9}, {339.3, 338.0}, {448.9, 447.5}, {10.8, 11.6}, {557.7, 556.0},
		{228.3, 228.1}, {998.0, 995.8}, {888.8, 887.6}, {119.6, 120.2}, {0.3, 0.3},
		{0.6, 0.3}, {557.6, 556.8}, {339.3, 339.1}, {888.0, 887.2}, {998.5, 999.0},
		{778.9, 779.0}, {10.2, 11.1}, {117.6, 118.3}, {228.9, 229.2}, {668.4, 669.1},
		{449.2, 448.9}, {0.2, 0.5},
	}

	corrData = [][]float64{{101.0, 99.2}, {100.1, 99.0}, {100.0, 100.0},
		{90.6, 111.6}, {86.5, 122.2}, {89.7, 117.6}, {90.6, 121.1}, {82.8, 136.0},
		{70.1, 154.2}, {65.4, 153.6}, {61.3, 158.5}, {62.5, 140.6}, {63.6, 136.2},
		{52.6, 168.0}, {59.7, 154.3}, {59.5, 149.0}, {61.3, 165.5},
	}

	infData = [][]float64{{15.6, 5.2}, {26.8, 6.1}, {37.8, 8.7}, {36.4, 8.5},
		{35.5, 8.8}, {18.6, 4.9}, {15.3, 4.5}, {7.9, 2.5}, {0.0, 1.1},
	}
}

func TestNorris(t *testing.T) {
	simple := NewSimple()
	for _, d := range norrisData {
		simple.Add(d[1], d[0])
	}
	// Tests against certified values from
	// http://www.itl.nist.gov/div898/strd/lls/data/LINKS/DATA/Norris.dat
	slope := simple.Slope()
	expectedSlope := 1.00211681802045
	slopeDistance := math.Abs(slope - expectedSlope)
	if slopeDistance > TOLERANCE {
		t.Errorf("slope %g does not match the expect value %g, the distance is %g, but the tolerance is %g\n", slope, expectedSlope, slopeDistance, TOLERANCE)
	}

	slopeStdErr := simple.SlopeStdErr()
	expectedSlopeStdErr := 0.429796848199937E-03
	slopeStdErrDistance := math.Abs(slopeStdErr - expectedSlopeStdErr)
	if slopeStdErrDistance > TOLERANCE {
		t.Errorf("slope std err %g does not match the expected value %g, the distance is %g, but the tolerance is %g\n", slopeStdErr, expectedSlopeStdErr, TOLERANCE)
	}

	intercept := simple.Intercept()
	expectedIntercept := -0.262323073774029
	interceptDistance := math.Abs(intercept - expectedIntercept)
	if interceptDistance > TOLERANCE {
		t.Errorf("intercept %g does not match the expect value %g, the distance is %g, but the tolerance is %g\n", intercept, expectedIntercept, interceptDistance, TOLERANCE)
	}

	interceptStdErr := simple.InterceptStdErr()
	expectedInterceptStdErr := 0.232818234301152
	interceptStdErrDistance := math.Abs(interceptStdErr - expectedInterceptStdErr)
	if interceptStdErrDistance > TOLERANCE {
		t.Errorf("intercept std err %g does not match the expected value %g, the distance is %g, but the tolerance is %g\n", interceptStdErr, expectedInterceptStdErr, interceptStdErrDistance, TOLERANCE)
	}

	n := simple.N()
	expectedN := 36
	if n != expectedN {
		t.Errorf("number of observations %d does not match the expect value %d\n", n, expectedN)
	}

	rsquare := simple.RSquare()
	expectedRSquare := 0.999993745883712
	rsquareDistance := math.Abs(rsquare - expectedRSquare)
	if rsquareDistance > TOLERANCE {
		t.Errorf("r-square %g does not match the expected value %g, the distance is %g, but the tolerance is %g\n", rsquare, expectedRSquare, rsquareDistance, TOLERANCE)
	}

	mse := simple.MeanSquareError()
	expectedMSE := 0.782864662630069
	mseDistance := math.Abs(mse - expectedMSE)
	if mseDistance > TOLERANCE9 {
		t.Errorf("MSE %g does not match the expected value %g, the distance is %g, but the tolerance is %g", mse, expectedMSE, mseDistance, TOLERANCE9)
	}

	sse := simple.SumSquaredErrors()
	expectedSSE := 26.6173985294224
	sseDistance := math.Abs(sse - expectedSSE)
	if sseDistance > TOLERANCE9 {
		t.Errorf("SSE %g does not match the expected value %g, the distance is %g, but the tolerance is %g", sse, expectedSSE, sseDistance, TOLERANCE9)
	}

	ssr := simple.RegressionSumSquares()
	expectedSSR := 4255954.13232369
	ssrDistance := math.Abs(ssr - expectedSSR)
	if ssrDistance > TOLERANCE9 {
		t.Errorf("SSR %g does not match the expected value %g, the distance is %g, but the tolerance is %g", ssr, expectedSSR, ssrDistance, TOLERANCE9)
	}
}

func TestNaNs(t *testing.T) {
	simple := NewSimple()
	if !math.IsNaN(simple.Intercept()) {
		t.Error("intercept not NaN")
	}
	if !math.IsNaN(simple.Slope()) {
		t.Error("slope not NaN")
	}
	if !math.IsNaN(simple.SlopeStdErr()) {
		t.Error("slope std err not NaN")
	}
	if !math.IsNaN(simple.InterceptStdErr()) {
		t.Error("intercept std err not NaN")
	}
	if !math.IsNaN(simple.MeanSquareError()) {
		t.Error("MSE not NaN")
	}
	if !math.IsNaN(simple.RSquare()) {
		t.Error("r-square not NaN")
	}
	if !math.IsNaN(simple.SumSquaredErrors()) {
		t.Error("SSE not NaN")
	}
	if !math.IsNaN(simple.TotalSumSquares()) {
		t.Error("SSTO not NaN")
	}
	if !math.IsNaN(simple.R()) {
		t.Error("R not NaN")
	}

	simple.Add(1, 2)
	simple.Add(1, 3)
	if !math.IsNaN(simple.Intercept()) {
		t.Error("intercept not NaN")
	}
	if !math.IsNaN(simple.Slope()) {
		t.Error("slope not NaN")
	}
	if !math.IsNaN(simple.SlopeStdErr()) {
		t.Error("slope std err not NaN")
	}
	if !math.IsNaN(simple.InterceptStdErr()) {
		t.Error("intercept std err not NaN")
	}
	if !math.IsNaN(simple.MeanSquareError()) {
		t.Error("MSE not NaN")
	}
	if !math.IsNaN(simple.RSquare()) {
		t.Error("r-square not NaN")
	}
	if !math.IsNaN(simple.SumSquaredErrors()) {
		t.Error("SSE not NaN")
	}
	if !math.IsNaN(simple.R()) {
		t.Error("R not NaN")
	}

	simple = NewSimple()
	simple.Add(1, 2)
	simple.Add(3, 3)
	if math.IsNaN(simple.Intercept()) {
		t.Error("intercept NaN")
	}
	if math.IsNaN(simple.Slope()) {
		t.Error("slope NaN")
	}
	if !math.IsNaN(simple.SlopeStdErr()) {
		t.Error("slope std err not NaN")
	}
	if !math.IsNaN(simple.InterceptStdErr()) {
		t.Error("intercept std err not NaN")
	}
	if !math.IsNaN(simple.MeanSquareError()) {
		t.Error("MSE not NaN")
	}
	if math.IsNaN(simple.RSquare()) {
		t.Error("r-square NaN")
	}
	if math.IsNaN(simple.SumSquaredErrors()) {
		t.Error("SSE NaN")
	}
	if math.IsNaN(simple.TotalSumSquares()) {
		t.Error("SSTO NaN")
	}
	if math.IsNaN(simple.R()) {
		t.Error("R NaN")
	}
	simple.Add(1, 4)
	if math.IsNaN(simple.SlopeStdErr()) {
		t.Error("slope std err NaN")
	}
	if math.IsNaN(simple.InterceptStdErr()) {
		t.Error("intercept std err NaN")
	}
	if math.IsNaN(simple.MeanSquareError()) {
		t.Error("MSE NaN")
	}
}

func TestCorr(t *testing.T) {
	simple := NewSimple()
	for _, d := range corrData {
		simple.Add(d[0], d[1])
	}

	n := simple.N()
	expectedN := 17
	if n != expectedN {
		t.Errorf("number of observations %d does not match the expect value %d\n", n, expectedN)
	}

	rsquare := simple.RSquare()
	expectedRSquare := 0.896123
	rsquareDistance := math.Abs(rsquare - expectedRSquare)
	if rsquareDistance > 10E-6 {
		t.Errorf("r-square %g does not match the expected value %g, the distance is %g, but the tolerance is %g\n", rsquare, expectedRSquare, rsquareDistance, 10E-6)
	}

	r := simple.R()
	expectedR := -0.94663767742
	rDistance := math.Abs(r - expectedR)
	if rDistance > 10E-10 {
		t.Errorf("r %g does not match the expected value %g, the distance is %g, but the tolerance is %g\n", r, expectedR, rDistance, 10E-10)
	}
}

func TestClear(t *testing.T) {
	simple := NewSimple()
	for _, d := range corrData {
		simple.Add(d[0], d[1])
	}
	n := simple.N()
	expectedN := 17
	if n != expectedN {
		t.Errorf("number of observations %d does not match the expect value %d\n", n, expectedN)
	}
	simple.Clear()
	n = simple.N()
	if n != 0 {
		t.Errorf("after clearance the number of observations should be zero, but we got %d\n", n)
	}

	for _, d := range corrData {
		simple.Add(d[0], d[1])
	}
	rsquare := simple.RSquare()
	expectedRSquare := 0.896123
	rsquareDistance := math.Abs(rsquare - expectedRSquare)
	if rsquareDistance > 10E-6 {
		t.Errorf("r-square %g does not match the expected value %g, the distance is %g, but the tolerance is %g\n", rsquare, expectedRSquare, rsquareDistance, 10E-6)
	}
}

func TestReomve(t *testing.T) {
	simple := NewSimple()
	for _, d := range infData {
		simple.Add(d[0], d[1])
	}
	removeX := infData[0][0]
	removeY := infData[0][1]
	simple.Remove(removeX, removeY)
	simple.Add(removeX, removeY)
	slopeStdErr := simple.SlopeStdErr()
	expectedSlopeStdErr := 0.011448491
	slopeStdErrDistance := math.Abs(slopeStdErr - expectedSlopeStdErr)
	if slopeStdErrDistance > 1E-10 {
		t.Errorf("slope std err %g does not match the expected value %g, the distance is %g, but the tolerance is %g\n", slopeStdErr, expectedSlopeStdErr, 1E-10)
	}

	interceptStdErr := simple.InterceptStdErr()
	expectedInterceptStdErr := 0.286036932
	interceptStdErrDistance := math.Abs(interceptStdErr - expectedInterceptStdErr)
	if interceptStdErrDistance > 1E-8 {
		t.Errorf("intercept std err %g does not match the expected value %g, the distance is %g, but the tolerance is %g\n", interceptStdErr, expectedInterceptStdErr, interceptStdErrDistance, 1E-8)
	}
}
