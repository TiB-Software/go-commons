package utils

import "testing"

func TestNormalizeDocument(t *testing.T) {
	got := NormalizeDocument("04.252.011/0001-10")
	want := "04252011000110"

	if got != want {
		t.Fatalf("Expected normalized document %q, got %q", want, got)
	}
}

func TestIsValidCPF(t *testing.T) {
	tests := []struct {
		name     string
		document string
		expected bool
	}{
		{name: "valid masked cpf", document: "529.982.247-25", expected: true},
		{name: "valid unmasked cpf", document: "52998224725", expected: true},
		{name: "invalid check digit", document: "529.982.247-26", expected: false},
		{name: "repeated digits", document: "111.111.111-11", expected: false},
		{name: "too short", document: "5299822472", expected: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValidCPF(tt.document)
			if got != tt.expected {
				t.Fatalf("Expected IsValidCPF(%q) to be %v, got %v", tt.document, tt.expected, got)
			}
		})
	}
}

func TestIsValidCNPJ(t *testing.T) {
	tests := []struct {
		name     string
		document string
		expected bool
	}{
		{name: "valid masked cnpj", document: "04.252.011/0001-10", expected: true},
		{name: "valid unmasked cnpj", document: "04252011000110", expected: true},
		{name: "invalid check digit", document: "04.252.011/0001-11", expected: false},
		{name: "repeated digits", document: "11.111.111/1111-11", expected: false},
		{name: "too short", document: "0425201100011", expected: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValidCNPJ(tt.document)
			if got != tt.expected {
				t.Fatalf("Expected IsValidCNPJ(%q) to be %v, got %v", tt.document, tt.expected, got)
			}
		})
	}
}
