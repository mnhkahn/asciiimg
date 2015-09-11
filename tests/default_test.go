package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/mnhkahn/asciiimg"
)

func TestMain(t *testing.T) {
	file, err := os.Open("./eimyymF.png")
	ai, err := asciiimg.NewAsciiImg(file)
	fmt.Println(ai, err, ai.Do())
}
