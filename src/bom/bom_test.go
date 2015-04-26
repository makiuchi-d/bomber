package bom

import (
	"testing"
)

func Test_match(t *testing.T) {
	s := []byte{0xef, 0xbb, 0xbf}
	if !match(s, UTF8) {
		t.Error()
	}
}

func Test_UTF8(t *testing.T) {
	b := GetBom([]byte{0xef, 0xbb, 0xbf, 0x41})
	if b != UTF8 {
		t.Error("returns " + b.ToString())
	}
}

func Test_UTF16LE(t *testing.T) {
	b := GetBom([]byte{0xff, 0xfe, 0x41, 0x00})
	if b != UTF16LE {
		t.Error("returns " + b.ToString())
	}
}

func Test_UTF16BE(t *testing.T) {
	b := GetBom([]byte{0xfe, 0xff, 0x00, 0x41})
	if b != UTF16BE {
		t.Error("returns " + b.ToString())
	}
}

func Test_UTF32LE(t *testing.T) {
	b := GetBom([]byte{0xff, 0xfe, 0x00, 0x00, 0x41, 0x00, 0x00, 0x00})
	if b != UTF32LE {
		t.Error("returns " + b.ToString())
	}
}

func Test_UTF32BE(t *testing.T) {
	b := GetBom([]byte{0x00, 0x00, 0xfe, 0xff, 0x00, 0x00, 0x00, 0x41})
	if b != UTF32BE {
		t.Error("returns " + b.ToString())
	}
}

func Test_NOTBOM(t *testing.T) {
	b := GetBom([]byte{0x41})
	if b != NOTBOM {
		t.Errorf("returns %s BOM (%s)", b.ToString(), string(b))
	}
}

func Test_Len(t *testing.T) {
	b := UTF8
	l := 3
	if len(b) != l {
		t.Errorf("length of %s is not %d (%d)", b.ToString(), l, len(b))
	}
	b = UTF16LE
	l = 2
	if len(b) != l {
		t.Errorf("length of %s is not %d (%d)", b.ToString(), l, len(b))
	}
	b = UTF32BE
	l = 4
	if len(b) != l {
		t.Errorf("length of %s is not %d (%d)", b.ToString(), l, len(b))
	}
}
