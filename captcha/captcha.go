package captcha

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"rezerwacje-duw-go/log"
	"strconv"
	"strings"
)

var zero = []byte("\x89PNG\r\n\x1a\n\x00\x00\x00\rIHDR\x00\x00\x00\f\x00\x00\x00\x10\x02\x03\x00\x00\x00_\x0fv\x94\x00\x00\x00\fPLTE\xff\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf8\x8c\x02M\x00\x00\x00>IDATx\xdacd\xd0*cah\xf8\xc7\xc4\xc4\xc0\xc4\xc2\xf1\xad\x8bI\xe1\xc1\x03\x06-\x06.\xa6\u007f\f\fL\r@\xcc\x00\xc6\nP\xba\x8e\xa9\t\xa8\x90a\xc5\x0f\xa6\a\f\x0f\x98~\xfck\x00\x00\x96S\x0eV\x19X\xb8\x14\x00\x00\x00\x00IEND\xaeB`\x82")
var one = []byte("\x89PNG\r\n\x1a\n\x00\x00\x00\rIHDR\x00\x00\x00\f\x00\x00\x00\x10\x02\x03\x00\x00\x00_\x0fv\x94\x00\x00\x00\fPLTE\xff\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf8\x8c\x02M\x00\x00\x00 IDATx\xdac`\x00\x83\x06&\x06\x10d\x80b\x0e8\x9b\x81A\vI\xbc\x0e\xc1f\xfa\x01\xa6\x015*\x02G\xac\xe9\x8b:\x00\x00\x00\x00IEND\xaeB`\x82")
var two = []byte("\x89PNG\r\n\x1a\n\x00\x00\x00\rIHDR\x00\x00\x00\f\x00\x00\x00\x10\x02\x03\x00\x00\x00_\x0fv\x94\x00\x00\x00\fPLTE\xff\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf8\x8c\x02M\x00\x00\x00EIDATx\xdacd\xe0*caX\xf0\x8d\x89\x8b\x81\x81\x85\xe1\xda\n\x06-\x06\r\xa6{\f\fL?\x18\x1a\x98\x18\x18~01\x80 \a\x10+<``X\xc5\xc0\xc4\xf5\x8d\x81Q\xab\xe1\x1bS\x03\x03\x03\xf3U\x87\xdf\x00\x9d\xa9\x0eB\v\xb3\xa8?\x00\x00\x00\x00IEND\xaeB`\x82")
var three = []byte("\x89PNG\r\n\x1a\n\x00\x00\x00\rIHDR\x00\x00\x00\f\x00\x00\x00\x10\x02\x03\x00\x00\x00_\x0fv\x94\x00\x00\x00\fPLTE\xff\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf8\x8c\x02M\x00\x00\x00>IDATx\xdacd\xd0\ncaj\xf8\xc7\U0010f049\x81\x81A\vD\x00q\x03\x13\x83\xd6? }\x03\x88\x15\x98\x18\xae\x01E\x194\xc0r\nL\x15L\\+\x18\x98\x1a\x18\x1e0p\xadb\x00\x00H-\nD֤P\t\x00\x00\x00\x00IEND\xaeB`\x82")
var four = []byte("\x89PNG\r\n\x1a\n\x00\x00\x00\rIHDR\x00\x00\x00\f\x00\x00\x00\x10\x02\x03\x00\x00\x00_\x0fv\x94\x00\x00\x00\fPLTE\xff\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf8\x8c\x02M\x00\x00\x00:IDATx\xda\x1d\x89\xb1\r\xc0 \f\xc0\x8c'\x94+2\xf7\xca\\\xd2\x03z\x15g0\xa1\x02\x96,\x0f\x86K\b\xb9-\t$\a|\xa1+\xed\x13s\xac\xf6\x14\x16\xfbA\xc3\xf7T:?\xff\x18\bq\x80\xe9\xfa\b\x00\x00\x00\x00IEND\xaeB`\x82")
var five = []byte("\x89PNG\r\n\x1a\n\x00\x00\x00\rIHDR\x00\x00\x00\f\x00\x00\x00\x10\x02\x03\x00\x00\x00_\x0fv\x94\x00\x00\x00\fPLTE\xff\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf8\x8c\x02M\x00\x00\x009IDATx\xdacdX\xc5\xc0\xc4\xc0\xf0\x0f\x88\x19\x18\x98\x16\x80\xd8`\xcc\xcc\x19\xaa\r\xa45\x98@\\\x05\x06-\xa6\a`\xf1\x06&\x06\xa6\u007f\x8c\xab\x18\xbe\x01\xd9\tL\f\xd7\x18\x00\xe3N\b\x99\xf9/:5\x00\x00\x00\x00IEND\xaeB`\x82")
var six = []byte("\x89PNG\r\n\x1a\n\x00\x00\x00\rIHDR\x00\x00\x00\f\x00\x00\x00\x10\x02\x03\x00\x00\x00_\x0fv\x94\x00\x00\x00\fPLTE\xff\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf8\x8c\x02M\x00\x00\x00AIDATx\xda\x15\x88\xb1\t\x80@\x10\xc0bl\x1c\xc3F\x90s\n\x97p\x14Kq\x9cǩn\n\xb1\x12OB \x04\xa0#v\b\xe4@\x1f\x1c\xb2z\xa6\x9f\xd6E\x18=7|\xa9\xf7[ؒ\xb80o>\xfd\xda\n(\xf5\xaao\b\x00\x00\x00\x00IEND\xaeB`\x82")
var seven = []byte("\x89PNG\r\n\x1a\n\x00\x00\x00\rIHDR\x00\x00\x00\f\x00\x00\x00\x10\x02\x03\x00\x00\x00_\x0fv\x94\x00\x00\x00\fPLTE\xff\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf8\x8c\x02M\x00\x00\x001IDATx\xdac\xd4j\xf8\xc7\xc4\xc0\xc0\x00\u008c\fL\xf3\x80\xf4\x03&\x06\x0e O\x8b\x01\"^\a\xa1\x99~@\xf9\x0f\x18\x988 \xea\x19\xb8\x18\x18\x00\xe8\x16\x05\xdf\x04\x9d7h\x00\x00\x00\x00IEND\xaeB`\x82")
var eight = []byte("\x89PNG\r\n\x1a\n\x00\x00\x00\rIHDR\x00\x00\x00\f\x00\x00\x00\x10\x02\x03\x00\x00\x00_\x0fv\x94\x00\x00\x00\fPLTE\xff\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf8\x8c\x02M\x00\x00\x00@IDATx\xdacd\xe0\x9a\xc6İ\x80\x81\x89\x89\x81\x81\x89\xe1[\a\x13\xc7\x03 \xcdP\xc1\xf4C\x8b\x81\x89\x03$\x06\x04Z@\x05+\x18\x14\x98~\x80\xf9\r\x8cZ\tbL\f\n\x0f\x98\x1e\xfck\x00\x00R\x11\fm{\fd\xb4\x00\x00\x00\x00IEND\xaeB`\x82")
var nine = []byte("\x89PNG\r\n\x1a\n\x00\x00\x00\rIHDR\x00\x00\x00\f\x00\x00\x00\x10\x02\x03\x00\x00\x00_\x0fv\x94\x00\x00\x00\fPLTE\xff\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00\x00\xf8\x8c\x02M\x00\x00\x00DIDATx\xdac`\xe0b`aX\xf0\x8d\x89\xeb\x1f\a\vÍ\x15\fZ\f\x1aL\f\f\f@\xdc\xc0İ\x8a\x81\xe9\x01\xc3\x0f\xa6\x1f\f\f\xcc\xff\xb7j31(\x00e\xb4\x80ru\fL@A&\x8e\a\f\x00\xaa\x11\r%\x87\xa9\x00e\x00\x00\x00\x00IEND\xaeB`\x82")

type mutablePalettedImage interface {
	Set(x, y int, c color.Color)
	image.PalettedImage
}

func decode(imgData []byte) mutablePalettedImage {
	img, err := png.Decode(bytes.NewReader(imgData))
	if err != nil {
		log.Errorf("Unable to decode reader")
	}
	return img.(mutablePalettedImage)
}

func preloadReferenceDigits() []mutablePalettedImage {
	data := [][]byte{zero, one, two, three, four, five, six, seven, eight, nine}
	digits := []mutablePalettedImage{}
	for _, undecodedDigit := range data {
		digits = append(digits, decode(undecodedDigit))
	}
	return digits
}

var digits = preloadReferenceDigits()

func deleteNoise(captcha *mutablePalettedImage) *mutablePalettedImage {
	bounds := (*captcha).Bounds()

	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			pixel := (*captcha).ColorIndexAt(x, y)
			if pixel != 2 {
				(*captcha).Set(x, y, color.RGBA{255, 255, 255, 255})
			}
		}
	}
	return captcha
}

func compareDigits(digitLeft *mutablePalettedImage, digitRight *mutablePalettedImage, originX int) int {
	bounds := (*digitLeft).Bounds()
	numberOfDifferences := 0

	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			pixelLeft := (*digitLeft).ColorIndexAt(x, y)
			pixelRight := (*digitRight).ColorIndexAt(originX+x, 2+y)
			if pixelLeft != pixelRight {
				numberOfDifferences = +1
			}
		}
	}
	return numberOfDifferences
}

func convertToString(captcha *mutablePalettedImage, originX int) string {
	minDiff := 1000
	index := 0
	for i, digit := range digits {
		result := compareDigits(&digit, captcha, originX)
		if result < minDiff {
			minDiff = result
			index = i
		}
	}
	return strconv.Itoa(index)
}

func RecognizeCaptcha(captcha *[]byte) string {
	decodedCaptcha := decode(*captcha)
	cleanCaptcha := deleteNoise(&decodedCaptcha)
	stringDigits := [6]string{}
	for i := 0; i <= 5; i++ {
		originX := 10 + i*12
		stringDigits[i] = convertToString(cleanCaptcha, originX)
	}
	return strings.Join(stringDigits[:], "")
}