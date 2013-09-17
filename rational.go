package rational

import "fmt"

type Rational interface {
    Equal(other Rational) bool
    String() string
	Numerator() int
	Denominator() int
	Add(other Rational) Rational
}

// ArrayRational implementation
type arrayRational [2]int

const n_i = 0
const d_i = 1

func New(num int, denom int) Rational {
	reduced_denom := gcd(num, denom)
	return &arrayRational{num / reduced_denom, denom / reduced_denom}
}

func (self *arrayRational) Numerator() int {
	return self[n_i]
}

func (self *arrayRational) Denominator() int {
	return self[d_i]
}

func (self *arrayRational) String() string {
	return fmt.Sprintf("%d/%d", self[n_i], self[d_i])
}

func (self *arrayRational) Equal(other Rational) bool {
	return self.Numerator() == other.Numerator() && self.Denominator() == other.Denominator()
}

func (self *arrayRational) Add(other Rational) Rational {
	new_numerator := self.Numerator()*other.Denominator() + self.Denominator()*other.Numerator()
	new_denominator := self.Denominator() * other.Denominator()
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
