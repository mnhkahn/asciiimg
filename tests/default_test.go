package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/mnhkahn/asciiimg"
	"github.com/mnhkahn/asciiimg/gray"
)

func TestMain(t *testing.T) {
	fmt.Println(len(gray.GetGray("default")))
	file_name := "eimyymF.png"
	file, err := os.Open("./" + file_name)
	ai, err := asciiimg.NewAsciiImg(file)
	if err == nil {
		ascii_txt, _ := os.Create(fmt.Sprintf("%s_ascii.txt", file_name))
		_, err := ascii_txt.Write([]byte(ai.Do()))
		fmt.Println(err)
		ascii_txt.Close()
	}
	fmt.Println("Test Complete.", err)
}
