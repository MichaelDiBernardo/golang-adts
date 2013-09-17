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
type ArrayRational [2]int

const n_i = 0
const d_i = 1

const (
	TypeHintArray = iota
	TypeHintStruct
)

func New(num, denom, hint int) Rational {
	reduced_denom := gcd(num, denom)
	numerator := num / reduced_denom
	denominator := denom / reduced_denom

	if hint == TypeHintArray {
		return &ArrayRational{numerator, denominator}
	} else {
		return &StructRational{numerator, denominator}
	}
}

func (self *ArrayRational) Numerator() int {
	return self[n_i]
}

func (self *ArrayRational) Denominator() int {
	return self[d_i]
}

func (self *ArrayRational) String() string {
	return fmt.Sprintf("%d/%d", self.Numerator(), self.Denominator())
}

func (self *ArrayRational) Equal(other Rational) bool {
	return self.Numerator() == other.Numerator() && self.Denominator() == other.Denominator()
}

func (self *ArrayRational) Add(other Rational) Rational {
	new_numerator := self.Numerator()*other.Denominator() + self.Denominator()*other.Numerator()
	new_denominator := self.Denominator() * other.Denominator()
	return New(new_numerator, new_denominator, TypeHintArray)
}

// StructRational implementation
type StructRational struct {
	numerator   int
	denominator int
}

func (self *StructRational) Numerator() int {
	return self.numerator
}

func (self *StructRational) Denominator() int {
	return self.denominator
}

func (self *StructRational) String() string {
	return fmt.Sprintf("%d/%d", self.Numerator(), self.Denominator())
}

func (self *StructRational) Equal(other Rational) bool {
	return self.Numerator() == other.Numerator() && self.Denominator() == other.Denominator()
}

func (self *StructRational) Add(other Rational) Rational {
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
