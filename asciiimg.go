package asciiimg

import (
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	"image/png"
	"io"
	"os"

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

func (this *AsciiImg) Do() string {
	ascii := ""
	fmt.Println(this.img.Bounds(), this.img.ColorModel())

	img_test := image.NewNRGBA(this.img.Bounds())
	for y := 0; y < this.img.Bounds().Dy(); y++ {
		for x := 0; x < this.img.Bounds().Dx(); x++ {
			Red, Green, Blue, Alpha := this.img.At(x, y).RGBA()
			Grey := Red*3/10 + Green*59/100 + Blue*11/100
			img_test.Set(x, y, color.NRGBA{uint8(Grey), uint8(Grey), uint8(Grey), uint8(Alpha)})
			// fmt.Println(Red, Green, Blue, Alpha, Grey)
		}
	}

	imgfile, _ := os.Create(fmt.Sprintf("%s_grey.png", this.name))
	defer imgfile.Close()

	err := png.Encode(imgfile, img_test)
	if err != nil {
		panic(err)
	}

	w, h := 4, 8
	rows := this.img.Bounds().Dy() / h
	cols := this.img.Bounds().Dx() / w

	fmt.Println(rows, cols)
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

func (this *AsciiImg) getBlockInfo(x, y, w, h int) uint32 {
	var sumGray uint32
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			Red, Green, Blue, _ := this.img.At(i+x, j+y).RGBA()
			Gray := Red*3/10 + Green*59/100 + Blue*11/100
			sumGray += Gray
		}
	}
	return sumGray / uint32(w*h)
}
