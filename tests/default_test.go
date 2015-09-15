package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/mnhkahn/asciiimg"
)

func TestMain(t *testing.T) {
	file_name := "eimyymF.png"
	file, err := os.Open("./" + file_name)
	ai, err := asciiimg.NewAsciiImg(file)
	if err == nil {
		ascii_txt, _ := os.Create(fmt.Sprintf("%s_ascii.txt", file_name))
		ascii_txt.Write([]byte(ai.Do(4, 8)))
		ascii_txt.Close()
	}
	fmt.Println("Test Complete.", err)
}
