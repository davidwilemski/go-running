package unit

import (
	"testing"
)

func TestMilesTypeMilesFunc(t *testing.T) {
	m := Miles(1.0)
	if m.Miles() != 1.0 {
		t.Errorf("Wrong output from m.Miles()")
	}
}

func TestMilesTypeKilometersFunc(t *testing.T) {
	m := Miles(1.0)

	if m.Kilometers() != kiloInMiles {
		t.Errorf("Wrong output from m.Kilometers()")
	}
}

func Test_Kilometers_MilesFunc(t *testing.T) {
	k := Kilometers(1.0)

	if k.Miles() != 1.0*kiloInMiles {
		t.Errorf("Wrong output from k.Miles()")
	}
}

func Test_Kilometers_KilometersFunc(t *testing.T) {
	k := Kilometers(1.0)

	if k.Kilometers() != 1.0 {
		t.Errorf("Wrong output from k.Kilometers()")
	}
}

func Test_Miles_ToKilometers(t *testing.T) {
	m := Miles(1.0)

	if m.ToKilometers() != Kilometers(m.Kilometers()) {
		t.Errorf("Wrong output from m.ToKilometers()")
	}
}

func Test_Miles_ToMiles(t *testing.T) {
	m := Miles(1.0)

	if m.ToMiles() != m {
		t.Errorf("Wrong output from m.ToMiles()")
	}
}

func Test_Kilometers_ToKilometers(t *testing.T) {
	k := Kilometers(1.0)

	if k.ToKilometers() != Kilometers(k.Kilometers()) {
		t.Errorf("Wrong output from m.ToKilometers()")
	}
}

func Test_Kilometers_ToMiles(t *testing.T) {
	k := Kilometers(1.0)

	if k.ToMiles() != Miles(k.Miles()) {
		t.Errorf("Wrong output from m.ToMiles()")
	}
}

func Test_Miles_Addition(t *testing.T) {
	m := Miles(1.0)

	if m+m != Miles(2.0) {
		t.Errorf("Miles addition failed or wrong result")
	}
}

func Test_Kilometers_Addition(t *testing.T) {
	k := Kilometers(1.0)

	if k+k != Kilometers(2.0) {
		t.Errorf("Kilometers addition failed or incorrect")
	}
}
