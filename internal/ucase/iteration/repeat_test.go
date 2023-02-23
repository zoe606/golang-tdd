package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("Expected %q but got %q", expected, repeated)
	}
}

func TestClone(t *testing.T) {
	rep := Repeat("a", 5)
	exp := strings.Clone(rep)

	if rep != exp {
		t.Errorf("Expected %q but got %q", exp, rep)

	}
}

func TestCompare(t *testing.T) {
	rep := Repeat("a", 5)
	comStr := strings.Compare(rep, "aaaaa")

	if comStr != 0 {
		t.Errorf("Expected %q but got %q", rep, comStr)
	}
}

func TestContains(t *testing.T) {
	word := "some string"
	rep := Repeat("a", 5)

	if strings.Contains(word, rep); false {
		t.Errorf("Expected %q but got %q", rep, word)
	}
}

func TestContainsAny(t *testing.T) {
	word := "some string"
	str := "i"

	fmt.Println(strings.ContainsAny(word, str))
	if strings.ContainsAny(word, str); false {
		t.Errorf("Expected %q but got %q", word, str)
	}
}

func TestContainsRune(t *testing.T) {
	t.Run("it should true", func(t *testing.T) {
		if got := strings.ContainsRune("asome", 97); got != true {
			t.Errorf("Expected %t but got %t", true, got)
		}
	})

	t.Run("it should false", func(t *testing.T) {
		if got := strings.ContainsRune("words", 97); got != false {
			t.Errorf("Expected %t but got %t", false, got)
		}
	})
}

func TestCount(t *testing.T) {
	word := Repeat("e", 3)

	t.Run("it should true", func(t *testing.T) {
		if got := strings.Count(word, "e"); got != 3 {
			t.Errorf("Expected %b but got %b", 1, got)
		}
	})

	t.Run("it should false", func(t *testing.T) {
		//var got return binary code/reference (100)
		if got := strings.Count(word, ""); got != 4 {
			t.Errorf("Expected %b but got %b", 1, got)
		}
	})
}

func TestCut(t *testing.T) {
	word := Repeat("aku", 3)

	before, after, found := strings.Cut(word, "uak")
	t.Run("it should true", func(t *testing.T) {
		if before != "ak" {
			t.Errorf("Expected %s but got %s", "akuakuaku", before)
		}
		if found != true {
			t.Errorf("Expected %t but got %t", true, found)
		}
		if after != "uaku" {
			t.Errorf("Expected %s but got %s", "akuakuaku", after)
		}
	})

	t.Run("it should false", func(t *testing.T) {
		if before == "akuakuaku" {
			t.Errorf("Expected %s but got %s", "ak", before)
		}
		if after == "aku" {
			t.Errorf("Expected %s but got %s", "uaku", after)
		}
	})
}

func TestCutPrefix(t *testing.T) {
	word := Repeat("xyuhuu", 3)

	after, found := strings.CutPrefix(word, "xyu")
	t.Run("it should true", func(t *testing.T) {
		if found != true {
			t.Errorf("Expected %t but got %t", true, found)
		}
		if after != "huuxyuhuuxyuhuu" {
			t.Errorf("Expected %s but got %s", "akuakuaku", after)
		}
	})
	t.Run("it should false", func(t *testing.T) {
		if after == "huuhuuhuu" {
			t.Errorf("Expected %s but got %s", "akuakuaku", after)
		}
	})
}

func TestCutSuffix(t *testing.T) {
	word := Repeat("xyuhuu", 3)

	before, found := strings.CutSuffix(word, "huu")
	t.Run("it should true", func(t *testing.T) {
		if found != true {
			t.Errorf("Expected %t but got %t", true, found)
		}
		if before != "xyuhuuxyuhuuxyu" {
			t.Errorf("Expected %s but got %s", "akuakuaku", before)
		}
	})
	t.Run("it should false", func(t *testing.T) {
		if before == "xyuxyuxyu" {
			t.Errorf("Expected %s but got %s", "akuakuaku", before)
		}
	})
}

func TestFields(t *testing.T) {
	word := Repeat("Say Hi! ", 5)
	fmt.Println(strings.Fields(word))
}

func TestHasPrefix(t *testing.T) {
	word := Repeat("xypher", 5)
	if got := strings.HasPrefix(word, "xyp"); got == false {
		t.Errorf("Expected %t but got %t", true, got)
	}
}

func TestHasSuffix(t *testing.T) {
	word := Repeat("pherxyp", 5)
	if got := strings.HasSuffix(word, "xyp"); got == false {
		t.Errorf("Expected %t but got %t", true, got)
	}
}

func TestIndex(t *testing.T) {
	word := Repeat("pherxyp", 3)
	idx := strings.Index(word, "xyp")
	lastIdx := strings.LastIndex(word, "xyp")

	if idx != 4 {
		t.Errorf("Expected %b but got %b", 4, idx)
	}

	if lastIdx != 18 {
		t.Errorf("Expected %b but got %b", 18, lastIdx)

	}
}

func TestIndexBytes(t *testing.T) {
	word := Repeat("uye", 3)
	if got := strings.IndexByte(word, 'x'); got != -1 {
		t.Errorf("Expected %b but got %b", -1, got)
	}
	if got := strings.IndexByte(word, 'u'); got != 0 {
		t.Errorf("Expected %b but got %b", 0, got)
	}
	if got := strings.IndexByte(word, 'e'); got != 2 {
		t.Errorf("Expected %b but got %b", 2, got)
	}

	if got := strings.LastIndexByte(word, 'e'); got != 8 {
		t.Errorf("Expected %b but got %b", 8, got)
	}
	if got := strings.LastIndexByte(word, 'y'); got != 7 {
		t.Errorf("Expected %b but got %b", 7, got)
	}
	if got := strings.LastIndexByte(word, 'u'); got != 6 {
		t.Errorf("Expected %b but got %b", 6, got)
	}
}

func TestJoin(t *testing.T) {
	word := []string{"abc", "def", "ghi"}
	if got := strings.Join(word, ","); got != "abc,def,ghi" {
		t.Errorf("Expected %s but got %s", "abc,def,ghi", got)
	}
}

func TestReplace(t *testing.T) {
	word := Repeat("huha", 5)
	if got := strings.Replace(word, "u", "a", 5); got != Repeat("haha", 5) {
		t.Errorf("Expected %s but got %s", Repeat("haha ", 5), got)
	}

	if got := strings.ReplaceAll(word, "huha", "hahu"); got != Repeat("hahu", 5) {
		t.Errorf("Expected %s but got %s", Repeat("hahu ", 5), got)
	}
}

func TestSplit(t *testing.T) {
	word := Repeat("this,string,should,be,splited,", 3)
	s := strings.Split(word, ",")
	sAfter := strings.SplitAfter(word, ",")
	sBefore := strings.SplitAfterN(word, ",", 4)
	sToN := strings.SplitN(word, ",", 1)
	fmt.Println(s)
	fmt.Println(sAfter)
	fmt.Println(sBefore)
	fmt.Println(sToN)
}

func TestStringFormatter(t *testing.T) {
	word := "this word shold be formatted"
	toLower := strings.ToLower(word)
	toTitle := strings.ToTitle(word)
	toUpper := strings.ToUpper(word)
	if word != toLower {
		t.Errorf("Expected %s but got %s", word, toLower)
	}

	if word == toTitle {
		t.Errorf("Expected %s but got %s", toTitle, word)
	}

	if word == toUpper {
		t.Errorf("Expected %s but got %s", toTitle, word)
	}
}

func TestTrimmer(t *testing.T) {
	word := Repeat(" some string with many many word ", 3)
	if got := strings.Trim(word, " "); got == word {
		t.Errorf("Expected %s but got %s", got, word)
	}
	if got := strings.TrimLeft(word, " "); got == word {
		t.Errorf("Expected %s but got %s", got, word)
	}
	if got := strings.TrimRight(word, " "); got == word {
		t.Errorf("Expected %s but got %s", got, word)
	}
	if got := strings.TrimPrefix(word, " some"); got == word {
		t.Errorf("Expected %s but got %s", got, word)
	}
	if got := strings.TrimSuffix(word, "word "); got == word {
		t.Errorf("Expected %s but got %s", got, word)
	}
	if got := strings.TrimSpace(word); got == word {
		t.Errorf("Expected %s but got %s", got, word)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
