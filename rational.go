package rational

type Rational struct {
	numerator   int
	denominator int
}

func New(num int, denom int) *Rational {
	return &Rational{num, denom}
}

func (self *Rational) Numerator() int {
	return self.numerator
}

func (self *Rational) Denominator() int {
	return self.denominator
}
