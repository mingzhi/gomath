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
package desc

import (
	"bitbucket.org/mingzhi/goutils/assert"
	"math"
	"testing"
)

var (
	mean         float64
	variance     float64
	secondMoment float64
	std          float64
	testArray    []float64
)

func init() {
	mean = 12.404545454545455
	variance = 10.00235930735931
	secondMoment = 210.04954545454547
	std = math.Sqrt(variance)
	testArray = []float64{12.5, 12.0, 11.8, 14.2, 14.9, 14.5, 21.0, 8.2, 10.3, 11.3,
		14.1, 9.9, 12.2, 12.0, 12.1, 11.0, 19.8, 11.0, 10.0, 8.8,
		9.0, 12.3}
}

func TestIncrement(t *testing.T) {
	var (
		result    float64
		expected  float64
		tolerance float64
		statistic StorelessUnivariateStatistic
	)
	tolerance = 1e-6
	// FIRST MOMENT
	statistic = NewFirstMoment()
	expected = mean
	for i := 0; i < len(testArray); i++ {
		statistic.Increment(testArray[i])
	}
	result = statistic.GetResult()
	if !assert.EqualFloat64(result, expected, tolerance, 1) {
		t.Errorf("FirstMoment Increment, result: %.10f, expected: %.10f\n", result, expected)
	}
	// SECOND MOMENT
	statistic = NewSecondMoment()
	expected = secondMoment
	for i := 0; i < len(testArray); i++ {
		statistic.Increment(testArray[i])
	}
	result = statistic.GetResult()
	if !assert.EqualFloat64(result, expected, tolerance, 1) {
		t.Errorf("SecondMoment Increment, result: %.10f, expected: %.10f\n", result, expected)
	}
	// Variance
	statistic = NewVarianceWithBiasCorrection()
	expected = variance
	for i := 0; i < len(testArray); i++ {
		statistic.Increment(testArray[i])
	}
	result = statistic.GetResult()
	if !assert.EqualFloat64(result, expected, tolerance, 1) {
		t.Errorf("Variance Increment, result:%.10f, expected: %.10f\n", result, expected)
	}
	// StandardDeviation
	statistic = NewStandardDeviationWithBiasCorrection()
	expected = std
	for i := 0; i < len(testArray); i++ {
		statistic.Increment(testArray[i])
	}
	result = statistic.GetResult()
	if !assert.EqualFloat64(result, expected, tolerance, 1) {
		t.Errorf("StandardDeviation Increment, result:%.10f, expected: %.10f\n", result, expected)
	}
}
