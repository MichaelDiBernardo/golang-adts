package rational_test

import (
    "rational"
	"testing"
)

func TestCreateRational(t *testing.T) {
	rat := rational.New(1, 2)
	if rat.Numerator != 1 {
		t.Errorf("Numerator not 1! %d", rat.Numerator)
	}
	if rat.Denominator != 2 {
		t.Errorf("Denominator not 2! %d", rat.Denominator)
	}
}

func TestDodgeConstructor(t *testing.T) {
	rat := rational.Rational{1, 2}
	if rat.Numerator != 1 {
		t.Errorf("Numerator not 1! %d", rat.Numerator)
	}
	if rat.Denominator != 2 {
		t.Errorf("Denominator not 2! %d", rat.Denominator)
	}
}
