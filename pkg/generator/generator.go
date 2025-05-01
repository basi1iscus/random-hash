package generator

import (
	"math/rand"
	"time"
)

type RandomValueType string
const (
	Integer RandomValueType = "int"
	Float RandomValueType = "float"
	Mixed RandomValueType = "mixed"
	String RandomValueType = "string"
)

const (
	firstChar = 32
	lastChar  = 127
)

type RandomGenerator struct {
	rnd *rand.Rand
}

func NewGenerator() *RandomGenerator {
	source := rand.NewSource(time.Now().UnixNano())
	return &RandomGenerator{
		rnd: rand.New(source),
	}
}

func (g *RandomGenerator) Generate(count int, valueType RandomValueType, min int, max int) []any {
	randoms := make([]any, count)

	for i := 0; i < count; i++ {
		switch valueType {
		case Integer:
			randoms[i] = g.generateInt(min, max)
		case Float:
			randoms[i] = g.generateFloat(min, max)
		case Mixed:
			randoms[i] = func() any {
				if g.rnd.Intn(2) == 0 {
					return g.generateInt(min, max)
				} else {
					return g.generateFloat(min, max)
				}
			}()
		case String:
			randoms[i] = g.generateString(min, max)
		}
	}

	return randoms
}

func (g *RandomGenerator) generateInt(min int, max int) int {
	return g.rnd.Intn(max - min + 1) + min
}

func (g *RandomGenerator) generateFloat(min int, max int) float64 {
	return g.rnd.Float64()*float64(max-min) + float64(min)
}

func (g *RandomGenerator) generateString(min int, max int) string {
	length := g.generateInt(min, max)
	str := make([]byte, length)
	for j := 0; j < length; j++ {
		str[j] = byte(g.generateInt(firstChar, lastChar))
	}
	return string(str)
}
