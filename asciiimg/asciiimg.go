package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mnhkahn/asciiimg"
)

var srcImg = flag.String("s", "", "Source image directory.")
var destAscii = flag.String("d", "", "Destination ascii direcotry.")

func main() {
	file, err := os.Open(*srcImg)
	if err != nil {
		panic(fmt.Sprintln("File path error", err, *srcImg))
	}
	ai, _ := asciiimg.NewAsciiImg(file)

	if len(*destAscii) == 0 {
		*destAscii = fmt.Sprintf("%s_ascii.txt", file.Name())
	}

	ascii_txt, _ := os.Create(*destAscii)
	ascii_txt.Write([]byte(ai.Do(4, 8)))
	defer ascii_txt.Close()

	fmt.Printf("File save in %s", ascii_txt.Name())
}

func init() {
	flag.Parse()
}
