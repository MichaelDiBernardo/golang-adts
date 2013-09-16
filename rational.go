package rational

import "fmt"

type Rational struct {
	Numerator   int
	Denominator int
}

func New(num int, denom int) *Rational {
	reduced_denom := gcd(num, denom)
	return &Rational{num / reduced_denom, denom / reduced_denom}
}

func (self *Rational) String() string {
	return fmt.Sprintf("%d/%d", self.Numerator, self.Denominator)
}

func (self *Rational) Equal(other *Rational) bool {
	return self.Numerator == other.Numerator && self.Denominator == other.Denominator
}

func (self *Rational) Add(other *Rational) *Rational {
	new_numerator := self.Numerator*other.Denominator + self.Denominator*other.Numerator
	new_denominator := self.Denominator * other.Denominator
	return New(new_numerator, new_denominator)
}

func gcd(m int, n int) int {
	if m == n {
		return m
	} else if m > n {
		return gcd(m-n, n)
	} else {
		return gcd(m, n-m)
	}
}
