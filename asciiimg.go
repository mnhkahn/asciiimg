package asciiimg

import (
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"math"

	"github.com/mnhkahn/asciiimg/gray"
)

type AsciiImg struct {
	name string
	img  image.Image
}

func NewAsciiImg(r io.Reader) (*AsciiImg, error) {
	ai := new(AsciiImg)
	var err error
	ai.img, ai.name, err = image.Decode(r)
	return ai, err
}

func (this *AsciiImg) DoByCol(cols int) string {
	ascii := ""

	if this.img == nil {
		return ascii
	}

	w := this.img.Bounds().Dx() / cols
	h := w * 2
	rows := this.img.Bounds().Dy() / h

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			x, y := c*w, r*h

			avg := this.getBlockInfo(x, y, w, h)
			ascii += string(gray.Get("default", avg))
		}
		ascii += "\r\n"
	}

	return ascii
}

func (this *AsciiImg) Do(w, h int) string {
	ascii := ""

	if this.img == nil {
		return ascii
	}

	rows := int(math.Ceil(float64(this.img.Bounds().Dy()) / float64(h)))
	cols := int(math.Ceil(float64(this.img.Bounds().Dx()) / float64(w)))

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			x, y := c*w, r*h

			avg := this.getBlockInfo(x, y, w, h)
			ascii += string(gray.Get("default", avg))
		}
		ascii += "\r\n"
	}

	return ascii
}

func (this *AsciiImg) getSize(x, y, w, h int) (int, int) {
	if x+w > this.img.Bounds().Dx() {
		w = this.img.Bounds().Dx() - x
	}
	if y+h > this.img.Bounds().Dy() {
		h = this.img.Bounds().Dy() - y
	}
	return w, h
}

func (this *AsciiImg) getBlockInfo(x, y, w, h int) uint32 {
	w, h = this.getSize(x, y, w, h)
	var sumGray uint32
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			Red, Green, Blue, _ := this.img.At(x+j, y+i).RGBA()
			Gray := Red*3/10 + Green*59/100 + Blue*11/100
			sumGray += Gray
		}
	}
	return sumGray / uint32(w*h)
}
