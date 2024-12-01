package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/mnhkahn/asciiimg"
)

func TestAscii(t *testing.T) {
	file_name := "eimyymF.png"
	file, err := os.Open("./" + fileName)
	ai, err := asciiimg.NewAsciiImg(file)
	if err == nil {
		asciiTxt, _ := os.Create(fmt.Sprintf("%s_ascii.txt", fileName))
		asciiTxt.Write([]byte(ai.Do(4, 12)))
		asciiTxt.Close()
	}
	fmt.Println("Test Complete.", err)
}
