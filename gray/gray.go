package gray

import (
	"math"
)

var DEFAULT_GRAY *Gray

type Gray struct {
	data map[string][]rune
}

func NewGray() *Gray {
	g := new(Gray)
	g.data = map[string][]rune{}
	return g
}

func (this *Gray) Add(name string, chars []rune) {
	DEFAULT_GRAY.data[name] = []rune{}
	step := int(math.Floor(float64(65535/len(chars)))) + 1
	for i := 0; i <= 65535; i++ {
		index := int(math.Floor(float64(i / step)))
		DEFAULT_GRAY.data[name] = append(DEFAULT_GRAY.data[name], chars[index])
	}
}

func GetGray(name string) []rune {
	return DEFAULT_GRAY.data[name]
}

func Get(name string, gray uint32) rune {
	return DEFAULT_GRAY.data[name][gray]
}

func init() {
	DEFAULT_GRAY = NewGray()
	DEFAULT_GRAY.Add("default", []rune{'@', 'w', '#', '$', 'k', 'd', 't', 'j', 'i', '.', ' '})
}
