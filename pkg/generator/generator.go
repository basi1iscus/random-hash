package generator

import (
	"math/rand"
	"time"
)

type RandomValueType int

const (
	Integer RandomValueType = iota
	Float
	Mixed
	String
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

func (g *RandomGenerator) Generate(count int, valueType RandomValueType, min int, max int) []interface{} {
	randoms := make([]interface{}, count)

	for i := 0; i < count; i++ {
		switch valueType {
		case Integer:
			randoms[i] = g.generateInt(min, max)
		case Float:
			randoms[i] = g.rnd.Float64() * float64(max)
		case Mixed:
			randoms[i] = func() interface{} {
				if g.rnd.Intn(2) == 0 {
					return g.generateInt(min, max)
				} else {
					return g.rnd.Float64() * float64(max)
				}
			}
		case String:
			randoms[i] = func() []byte {
				str := make([]byte, max)
				length := g.generateInt(min, max)
				for j := 0; j < length; j++ {
					str[j] = byte(g.generateInt(firstChar, lastChar))
				}
				return str
			}
		}
	}

	return randoms
}

func (g *RandomGenerator) generateInt(min int, max int) int {
	return g.rnd.Intn(max-min+1) + min
}
