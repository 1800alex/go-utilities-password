package password

import (
	"log"
	"strings"
	"testing"
)

const (
	N = 10000
)

func testHasDuplicates(tb testing.TB, s string) bool {
	found := make(map[rune]struct{}, len(s))
	for _, ch := range s {
		if _, ok := found[ch]; ok {
			return true
		}
		found[ch] = struct{}{}
	}
	return false
}

func TestGenerator_Generate(t *testing.T) {
	t.Parallel()

	gen, err := NewGenerator(nil)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("exceeds_available", func(t *testing.T) {
		t.Parallel()

		if _, err := gen.Generate(53, false, false, false, false); err != ErrLettersExceedsAvailable {
			t.Errorf("expected %q to be %q", err, ErrLettersExceedsAvailable)
		}

		if _, err := gen.Generate(83, false, true, false, false); err != ErrLettersExceedsAvailable {
			t.Errorf("expected %q to be %q", err, ErrLettersExceedsAvailable)
		}

		if _, err := gen.Generate(63, false, false, true, false); err != ErrLettersExceedsAvailable {
			t.Errorf("expected %q to be %q", err, ErrLettersExceedsAvailable)
		}

		if _, err := gen.Generate(93, false, true, true, false); err != ErrLettersExceedsAvailable {
			t.Errorf("expected %q to be %q", err, ErrLettersExceedsAvailable)
		}
	})

	t.Run("gen_lowercase", func(t *testing.T) {
		t.Parallel()

		for i := 0; i < N; i++ {
			res, err := gen.Generate(i%len(LowerLetters), false, false, true, true)
			if err != nil {
				t.Error(err)
			}

			if res != strings.ToLower(res) {
				t.Errorf("%q is not lowercase", res)
			}
		}
	})

	t.Run("gen_uppercase", func(t *testing.T) {
		t.Parallel()

		res, err := gen.Generate(1000, false, false, false, true)
		if err != nil {
			t.Error(err)
		}

		if res == strings.ToLower(res) {
			t.Errorf("%q does not include uppercase", res)
		}
	})

	t.Run("gen_no_repeats", func(t *testing.T) {
		t.Parallel()

		for i := 0; i < N; i++ {
			res, err := gen.Generate(52, true, true, false, false)
			if err != nil {
				t.Error(err)
			}

			if testHasDuplicates(t, res) {
				t.Errorf("%q should not have duplicates", res)
			}
		}
	})
}

func TestGenerator_Generate_Custom(t *testing.T) {
	t.Parallel()

	gen, err := NewGenerator(&GeneratorInput{
		LowerLetters: "abcde",
		UpperLetters: "ABCDE",
		Symbols:      "!@#$%",
		Digits:       "01234",
	})
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < N; i++ {
		res, err := gen.Generate(52, true, true, false, true)
		if err != nil {
			t.Error(err)
		}

		if strings.Contains(res, "f") {
			t.Errorf("%q should only contain lower letters abcde", res)
		}

		if strings.Contains(res, "F") {
			t.Errorf("%q should only contain upper letters ABCDE", res)
		}

		if strings.Contains(res, "&") {
			t.Errorf("%q should only include symbols !@#$%%", res)
		}

		if strings.Contains(res, "5") {
			t.Errorf("%q should only contain digits 01234", res)
		}
	}
}

func TestGenerator_Generate_Standalone(t *testing.T) {
	t.Parallel()

	t.Run("gen_no_repeats", func(t *testing.T) {
		t.Parallel()

		for i := 0; i < N; i++ {
			res, err := Generate(52, true, true, false, false)
			if err != nil {
				t.Error(err)
			}

			if testHasDuplicates(t, res) {
				t.Errorf("%q should not have duplicates", res)
			}
		}
	})
}

func ExampleGenerate() {
	res, err := Generate(64, true, true, false, false)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(res)
}

func ExampleGenerator_Generate() {
	gen, err := NewGenerator(nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := gen.Generate(64, true, true, false, false)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(res)
}

func ExampleNewGenerator_nil() {
	// This is exactly the same as calling "Generate" directly. It will use all
	// the default values.
	gen, err := NewGenerator(nil)
	if err != nil {
		log.Fatal(err)
	}

	_ = gen // gen.Generate(...)
}

func ExampleNewGenerator_custom() {
	// Customize the list of symbols.
	gen, err := NewGenerator(&GeneratorInput{
		Symbols: "!@#$%^()",
	})
	if err != nil {
		log.Fatal(err)
	}

	_ = gen // gen.Generate(...)
}
