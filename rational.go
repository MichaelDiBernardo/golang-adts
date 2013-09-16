package rational

import "fmt"

// Completely change the representation - tests don't break!
type Rational [2]int

const n_i = 0;
const d_i = 1;

func New(num int, denom int) *Rational {
	reduced_denom := gcd(num, denom)
	return &Rational{num / reduced_denom, denom / reduced_denom}
}

func (self *Rational) Numerator() int {
    return self[n_i]
}

func (self *Rational) Denominator() int {
    return self[d_i]
}

func (self *Rational) String() string {
	return fmt.Sprintf("%d/%d", self[n_i], self[d_i])
}

func (self *Rational) Equal(other *Rational) bool {
	return self[n_i] == other[n_i] && self[d_i] == other[d_i]
}

func (self *Rational) Add(other *Rational) *Rational {
	new_numerator := self[n_i]*other[d_i] + self[d_i]*other[n_i]
	new_denominator := self[d_i] * other[d_i]
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
