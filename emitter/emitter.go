package emitter

import (
	"github.com/olebedev/emitter"
)

func NewEmitter() *emitter.Emitter {
	return emitter.New(10000)
}
