package rational

type Rational struct {
	Numerator int
	Denominator int
}

func New(num int, denom int) *Rational {
	return &Rational{num, denom}
}
