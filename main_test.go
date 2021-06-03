package cfiscale

import (
	"testing"
)

func TestCreationCall(t *testing.T) {
	p := NewPerson("silvio", "berlusconi", "milano", "29/09/1936", "M")
	res, err := p.DoRequest()

	if err != nil {
		t.Error("http call not completed")
	}

	expected := "BRLSLV36P29F205W"

	if res != expected {
		t.Error("expected code different from real")
	}
}

func TestVerificationCallOK(t *testing.T) {
	p := NewPerson("loredana", "panico", "san pietro vernotico", "12/12/1987", "F")

	fc := "PNCLDN87T52I119C"
	ok, err := p.Verify(fc)

	if !ok {
		t.Errorf("Fiscal code %v should be valid", fc)
	}

	if err != nil {
		t.Error(err)
	}
}

func TestVerificationCallFail(t *testing.T) {
	p := NewPerson("loredana", "panico", "san pietro vernotico", "12/12/1978", "M")

	// forced to be invalid
	fc := "BRLSLV36P29F205X"

	ok, err := p.Verify(fc)

	if ok {
		t.Errorf("Fiscal code %v should be invalid", fc)
	}

	if err != nil {
		t.Error(err)
	}
}
