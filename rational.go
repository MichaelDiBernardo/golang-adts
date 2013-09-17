package rational

import "fmt"

type Rational interface {
	Equal(other Rational) bool
	String() string
	Numerator() int
	Denominator() int
	Add(other Rational) Rational
}

const (
	TypeHintArray = iota
	TypeHintStruct
)

// Factory function for rationals.
func New(num, denom, hint int) Rational {
	reduced_denom := gcd(num, denom)
	numerator := num / reduced_denom
	denominator := denom / reduced_denom

	if hint == TypeHintArray {
		return &arrayRational{numerator, denominator}
	} else {
		return &structRational{numerator, denominator}
	}
}

// v v v FROM HEREON IN, FOR OUR EYES ONLY v v v 
// arrayRational implementation
type arrayRational [2]int

const n_i = 0
const d_i = 1

func (self *arrayRational) Numerator() int {
	return self[n_i]
}

func (self *arrayRational) Denominator() int {
	return self[d_i]
}

func (self *arrayRational) String() string {
	return fmt.Sprintf("%d/%d", self.Numerator(), self.Denominator())
}

func (self *arrayRational) Equal(other Rational) bool {
	return self.Numerator() == other.Numerator() && self.Denominator() == other.Denominator()
}

func (self *arrayRational) Add(other Rational) Rational {
	new_numerator := self.Numerator()*other.Denominator() + self.Denominator()*other.Numerator()
	new_denominator := self.Denominator() * other.Denominator()
	return New(new_numerator, new_denominator, TypeHintArray)
}

// structRational implementation
type structRational struct {
	numerator   int
	denominator int
}

func (self *structRational) Numerator() int {
	return self.numerator
}

func (self *structRational) Denominator() int {
	return self.denominator
}

func (self *structRational) String() string {
	return fmt.Sprintf("%d/%d", self.Numerator(), self.Denominator())
}

func (self *structRational) Equal(other Rational) bool {
	return self.Numerator() == other.Numerator() && self.Denominator() == other.Denominator()
}

func (self *structRational) Add(other Rational) Rational {
	new_numerator := self.Numerator()*other.Denominator() + self.Denominator()*other.Numerator()
	new_denominator := self.Denominator() * other.Denominator()
	return New(new_numerator, new_denominator, TypeHintStruct)
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
