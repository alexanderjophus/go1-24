package pkg1

import (
	"golang.org/x/exp/constraints"
)

type Constraint constraints.Ordered
type G[P Constraint] struct {
	Field P
}
