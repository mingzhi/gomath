package desc

type SummaryStatistics struct {
	n            int
	sum          *Sum
	mean         *Mean
	max          *Max
	min          *Min
	secondMoment *SecondMoment
	variance     *Variance
}
