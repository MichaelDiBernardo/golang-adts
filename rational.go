package rational

import "fmt"

type Rational struct {
	n int
	d int
}

func New(num int, denom int) *Rational {
	reduced_denom := gcd(num, denom)
	return &Rational{num / reduced_denom, denom / reduced_denom}
}

func (self *Rational) Numerator() int {
    return self.n
}

func (self *Rational) Denominator() int {
    return self.d
}

func (self *Rational) String() string {
	return fmt.Sprintf("%d/%d", self.n, self.d)
}

func (self *Rational) Equal(other *Rational) bool {
	return self.n == other.n && self.d == other.d
}

func (self *Rational) Add(other *Rational) *Rational {
	new_numerator := self.n*other.d + self.d*other.n
	new_denominator := self.d * other.d
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
