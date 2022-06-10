package cache

import (
	"github.com/fengziren/designPatterns/prototype/protyp"
	"github.com/fengziren/designPatterns/prototype/protyp/realization"
)

type ShapeCache struct {
	cache map[uint64]protyp.Clone
}

func (s *ShapeCache) GetShape(id uint64) protyp.Clone {
	return s.cache[id]
}

func (s *ShapeCache) LoadCache() {
	if s.cache == nil {
		s.cache = make(map[uint64]protyp.Clone)
	}

	r := realization.GetRectangle()
	r.SetID(1)
	s.cache[r.GetID()] = r

}
