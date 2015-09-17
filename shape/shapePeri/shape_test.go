package shapePeri

import "testing"

func TestShapePeri(t *testing.T) {
	result1 := 37.69911184307752
	result2 := 20.0
	c := Circle{}
	r := Rectangle{}
	c.SetRadius(6)
	r.SetRectValues(3, 7)

	peri := Perimeter(&c)
	if peri != result1 {
		t.Error("Expected", result1, " got ", peri)
	}

	peri = Perimeter(&r)
	if peri != result2 {
		t.Error("Expected", result2, " got ", peri)
	}
}
