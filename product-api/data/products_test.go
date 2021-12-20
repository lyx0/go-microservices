package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Testea",
		Price: 1.0,
		SKU:   "abc-adsad-dasd",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
