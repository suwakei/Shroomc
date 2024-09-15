package code

import (
	"fmt"
)


type Instructions []byte

type OptimizedCode byte


const (
	OptimizedConstant = iota
)

type Definition struct {
	Name string // OptimizedCodeを読みやすくする
	OperandWidths []int // 
}

var definitions = map[OptimizedCode]*Definition{
	OptimizedConstant: {"OptimizedConstant", []int{2}},
}

func lookup(op byte) (*Definition, error) {
	def, ok := definitions[OptimizedCode(op)]
	if !ok {
		return nil, fmt.Errorf("OptimizedCode %d undifined", op)
		return def, nil
	}
}