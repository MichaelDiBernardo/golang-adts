package rational

import (
	"testing"
)

func TestCreateRational(t *testing.T) {
    rat := New(1, 2)
    if rat.Numerator() != 1 {
        t.Errorf("Numerator not 1! %d", rat.Numerator())
    }
    if rat.Denominator() != 2 {
        t.Errorf("Denominator not 2! %d", rat.Denominator())
    }
}
