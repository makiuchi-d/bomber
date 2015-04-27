package main

import (
	"bytes"
	"testing"

	"bom"
)

func Test_addBom8(t *testing.T) {
	buf1 := make([]byte, 0, 16)
	in := bytes.NewReader([]byte{0x42})
	out := bytes.NewBuffer(buf1)

	addBom(in, out, bom.UTF8)
	if !bytes.Equal([]byte{0xef, 0xbb, 0xbf, 0x42}, out.Bytes()) {
		t.Errorf("invalid UTF8 bom data %v", out.Bytes())
	}
}

func Test_addBom8Again(t *testing.T) {
	buf1 := make([]byte, 0, 16)
	in := bytes.NewReader([]byte{0xef, 0xbb, 0xbf, 0x42})
	out := bytes.NewBuffer(buf1)

	addBom(in, out, bom.UTF8)
	if !bytes.Equal([]byte{0xef, 0xbb, 0xbf, 0x42}, out.Bytes()) {
		t.Errorf("extra BOM %v", out.Bytes())
	}
}

func Test_addBom32BE(t *testing.T) {
	buf1 := make([]byte, 0, 16)
	in := bytes.NewReader([]byte{0x42})
	out := bytes.NewBuffer(buf1)

	addBom(in, out, bom.UTF32BE)

	if !bytes.Equal([]byte{0x00, 0x00, 0xfe, 0xff, 0x42}, out.Bytes()) {
		t.Errorf("invalid UTF-32 BE bom data %v", out.Bytes())
	}
}

func Test_removeUTF16BE(t *testing.T) {
	in := bytes.NewReader([]byte{0xfe, 0xff, 0x42})
	out := bytes.NewBuffer(make([]byte, 0, 16))

	removeBom(in, out, bom.UTF16BE)
	if !bytes.Equal([]byte{0x42}, out.Bytes()) {
		t.Errorf("fail to remove UTF16 BE bom: ", out.Bytes())
	}
}

func Test_removeAnotherTargetBom(t *testing.T) {
	in := bytes.NewReader([]byte{0xfe, 0xff, 0x42, 0x43})
	out := bytes.NewBuffer(make([]byte, 0, 16))

	removeBom(in, out, bom.UTF32LE)
	if !bytes.Equal([]byte{0xfe, 0xff, 0x42, 0x43}, out.Bytes()) {
		t.Errorf("fail to remove UTF16 BE bom: ", out.Bytes())
	}
}

func Test_removeFromNoBomString(t *testing.T) {
	buf1 := make([]byte, 0, 16)
	in := bytes.NewReader([]byte{0x42, 0x43, 0x44})

	out := bytes.NewBuffer(buf1)

	removeBom(in, out, bom.UTF8)
	if !bytes.Equal([]byte{0x42, 0x43, 0x44}, out.Bytes()) {
		t.Errorf("fail to remove UTF16 BE bom: ", out.Bytes())
	}
}
