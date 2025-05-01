package generator

import (
	"testing"
)

func TestNewGenerator(t *testing.T) {
	gen := NewGenerator()
	if gen == nil {
		t.Fatal("Expected generator to be non-nil")
	}
	if gen.rnd == nil {
		t.Fatal("Expected generator's random source to be non-nil")
	}
}

func TestGenerateIntegers(t *testing.T) {
	gen := NewGenerator()
	count := 5
	min, max := 1, 10
	values := gen.Generate(count, Integer, min, max)

	if len(values) != count {
		t.Fatalf("Expected %d values, got %d", count, len(values))
	}

	for _, v := range values {
		val, ok := v.(int)
		if !ok {
			t.Fatalf("Expected value to be of type int, got %T", v)
		}
		if val < min || val > max {
			t.Fatalf("Value %d out of range [%d, %d]", val, min, max)
		}
	}
}

func TestGenerateFloats(t *testing.T) {
	gen := NewGenerator()
	count := 5
	min, max := 1, 10
	values := gen.Generate(count, Float, min, max)

	if len(values) != count {
		t.Fatalf("Expected %d values, got %d", count, len(values))
	}

	for _, v := range values {
		val, ok := v.(float64)
		if !ok {
			t.Fatalf("Expected value to be of type float64, got %T", v)
		}
		if val < 0 || val > float64(max) {
			t.Fatalf("Value %f out of range [0, %d]", val, max)
		}
	}
}

func TestGenerateMixed(t *testing.T) {
	gen := NewGenerator()
	count := 5
	min, max := 1, 10
	values := gen.Generate(count, Mixed, min, max)

	if len(values) != count {
		t.Fatalf("Expected %d values, got %d", count, len(values))
	}

	for _, v := range values {
		switch val := v.(type) {
		case int:
			if val < min || val > max {
				t.Fatalf("Integer value %d out of range [%d, %d]", val, min, max)
			}
		case float64:
			if val < 0 || val > float64(max) {
				t.Fatalf("Float value %f out of range [0, %d]", val, max)
			}
		default:
			t.Fatalf("Unexpected value type %T", v)
		}
	}
}

func TestGenerateStrings(t *testing.T) {
	gen := NewGenerator()
	count := 5
	min, max := 3, 8
	values := gen.Generate(count, String, min, max)

	if len(values) != count {
		t.Fatalf("Expected %d values, got %d", count, len(values))
	}

	for _, v := range values {
		str, ok := v.(string)
		if !ok {
			t.Fatalf("Expected value to be of type []byte, got %T", v)
		}
		length := len(str)
		if length < min || length > max {
			t.Fatalf("String length %d out of range [%d, %d]", length, min, max)
		}
		for _, char := range str {
			if char < firstChar || char >= lastChar {
				t.Fatalf("Character %d out of range [%d, %d)", char, firstChar, lastChar)
			}
		}
	}
}