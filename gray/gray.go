package gray

var DEFAULT_GRAY *Gray

type Gray struct {
	data map[string]map[uint8]rune
}

func NewGray() *Gray {
	g := new(Gray)
	g.data = map[string]map[uint8]rune{}
	return g
}

func (this *Gray) Add(name string, chars []rune) {
	DEFAULT_GRAY.data[name] = map[uint8]rune{}
	step := 256 / len(chars)
	for i := 0; i <= 255; i++ {
		index := i / step
		DEFAULT_GRAY.data[name][uint8(i)] = chars[index]
		println(i, index, string(chars[index]))
	}
}

func Get(name string, gray uint8) rune {
	return DEFAULT_GRAY.data[name][gray]
}

func init() {
	DEFAULT_GRAY = NewGray()
	DEFAULT_GRAY.Add("default", []rune{'@', 'w'})
}
