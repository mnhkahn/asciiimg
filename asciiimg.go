package asciiimg

import (
	"fmt"
	"image"
	"image/color"
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

	fmt.Println(this.img.Bounds(), this.img.ColorModel(), string(gray.Get("default", 0)))

	w, h := 2, 2
	rows := this.img.Bounds().Dy() / h
	cols := this.img.Bounds().Dx() / w

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			x, y := c*w, c*h

			avg := this.getBlockInfo(w, h)
			ascii += string(gray.Gray("default", avg))
		}
		ascii += "\r\n"
	}

	return ascii
}

func (this *AsciiImg) getBlockInfo(w, h int) uint8 {
	var sumGray uint8
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			Red, Green, Blue, _ := this.img.At(i, j).RGBA()
			Grey := Red*3/10 + Green*59/100 + Blue*11/100
			sumGray += Grey
		}
	}
	return sumGray / w * h
}
