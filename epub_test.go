package epub

import (
	"bytes"
	"testing"
)

func TestReader(t *testing.T) {
	fn := "./data/test.epub"

	bk, err := Open(fn)
	if err != nil {
		t.Fatal(err)
	}

	if len(bk.Files()) != 211 {
		t.Fatal("invalid files counter")
	}

	bk.Close()

	i := 0

	Reader(fn, func(n string, data []byte) bool {
		i++
		if data == nil {
			t.Fatal("reader failed")
		}
		return true
	})

	if i != 189 {
		t.Fatalf("Invalid chapter numbers: %d != 182", i)
	}

	buf := bytes.NewBuffer(nil)

	ToTxt(fn, buf)

	if buf.Len() == 0 {
		t.Fatal("ToTxt failed")
	}
}

func TestGutenbergEBook(t *testing.T) {
	fn := "./data/kim.epub"

	bk, err := Open(fn)
	if err != nil {
		t.Fatal(err)
	}

	if len(bk.Files()) != 25 {
		t.Fatalf("invalid files counter: %d vs 22", len(bk.Files()))
	}

	bk.Close()

	i := 0

	Reader(fn, func(n string, data []byte) bool {
		i++
		t.Logf("Got chapter: %s", n)
		if data == nil {
			t.Fatal("reader failed: empty data")
		}
		return true
	})

	if i != 18 {
		t.Fatalf("Invalid chapter numbers: %d", i)
	}

	buf := bytes.NewBuffer(nil)

	ToTxt(fn, buf)

	if buf.Len() == 0 {
		t.Fatal("ToTxt failed")
	}
}
