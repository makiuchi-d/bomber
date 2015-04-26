package bom

type BOM string

const UTF8 BOM = "\xef\xbb\xbf"
const UTF16LE BOM = "\xff\xfe"
const UTF16BE BOM = "\xfe\xff"
const UTF32LE BOM = "\xff\xfe\x00\x00"
const UTF32BE BOM = "\x00\x00\xfe\xff"
const NOTBOM BOM = ""

func match(data []byte, bom BOM) bool {
	for i, b := range []byte(bom) {
		if data[i] != b {
			return false
		}
	}
	return true
}

func GetBom(data []byte) BOM {
	switch {
	case match(data, UTF8):
		return UTF8
	case match(data, UTF32LE):
		return UTF32LE
	case match(data, UTF16LE):
		return UTF16LE
	case match(data, UTF16BE):
		return UTF16BE
	case match(data, UTF32BE):
		return UTF32BE
	}
	return NOTBOM
}

func (b BOM) ToString() string {
	switch b {
	case UTF8:
		return "UTF-8"
	case UTF16LE:
		return "UTF-16 LE"
	case UTF16BE:
		return "UTF-16 BE"
	case UTF32LE:
		return "UTF-32 LE"
	case UTF32BE:
		return "UTF-32 BE"
	default:
		return "Unknown"
	}
}
