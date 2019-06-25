package nextGenerator

import (
	"math/rand"
	"time"
)

type Generator struct {
	Vals   []uint32
	r      *rand.Rand
	length uint32
}

func (g *Generator) GetNext() uint32 {
	next := g.getRandomValue()
	for g.hasValue(next) || next == 0 {
		next = g.getRandomValue()
		//log.Println(g.Vals, next)
	}

	return next
}

func (g *Generator) hasValue(val uint32) bool {
	for _, v := range g.Vals {
		if val == v {
			return true
		}
	}

	return false
}

func (g *Generator) getRandomValue() uint32 {
	return uint32(g.r.Int63n(int64(g.length)))
}

func MakeGenerator(length uint32, initialVals []uint32) Generator {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return Generator{initialVals, r, length}
}
