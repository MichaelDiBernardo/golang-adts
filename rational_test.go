package rational_test

import (
	"rational"
	"testing"
)

func TestCreateRational(t *testing.T) {
	rat := rational.New(1, 2)
	if rat.Numerator() != 1 {
		t.Errorf("Numerator not 1! %d", rat.Numerator())
	}
	if rat.Denominator() != 2 {
		t.Errorf("Denominator not 2! %d", rat.Denominator())
	}
}

// Can't do this!
//func TestDodgeConstructor(t *testing.T) {
//	rat := rational.Rational{1, 2}
//	if rat.Numerator != 1 {
//		t.Errorf("Numerator not 1! %d", rat.Numerator)
//	}
//	if rat.Denominator != 2 {
//		t.Errorf("Denominator not 2! %d", rat.Denominator)
//	}
//}

func TestString(t *testing.T) {
	rat := rational.New(1, 2)
	actual := rat.String()
	if actual != "1/2" {
		t.Errorf("String of 1/2 was %s", actual)
	}
}

func TestEqualityOfIrreducible(t *testing.T) {
	same_1, same_2 := rational.New(1, 2), rational.New(1, 2)
	diff := rational.New(1, 3)

	if !same_1.Equal(same_2) {
		t.Errorf("Two same values were not equal.")
	}

	if same_1.Equal(diff) {
		t.Errorf("Two different values were equal.")
	}
}

func TestAddWithoutReduce(t *testing.T) {
	first, second := rational.New(1, 2), rational.New(1, 3)
	expected := rational.New(5, 6)

	sum := first.Add(second)
	if !sum.Equal(expected) {
		t.Errorf("Sum to irreducible fraction: E - %s, A - %s", expected.String(), sum.String())
	}
}

func CreateReducibleFraction(t *testing.T) {
	rat := rational.New(10, 18)
	if rat.Numerator() != 5 {
		t.Errorf("Numerator did not reduce! E - 5, A - %d", rat.Numerator)
	}
	if rat.Denominator() != 9 {
		t.Errorf("Denominator did not reduce! E - 9, A - %d", rat.Denominator)
	}
}

func TestAddWithReduce(t *testing.T) {
	first, second := rational.New(1, 2), rational.New(1, 3)
	expected := rational.New(5, 6)

	sum := first.Add(second)
	if !sum.Equal(expected) {
		t.Errorf("Sum to irreducible fraction: E - %s, A - %s", expected.String(), sum.String())
	}
}

//func TestTryToCreateConcreteRational(t *testing.T) {
//	r := rational.arrayRational{1, 2}
//}
//
//func TestTryToViolateEncapsulation(t *testing.T) {
//	r := rational.New(1, 2)
//	r[0] = 1
//}
