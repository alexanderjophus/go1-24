package pkg2

import "github.com/alexanderjophus/go124/generic-type-aliases/pkg1"

type Constraint = pkg1.Constraint // pkg1.Constraint could also be used directly in G

// type G = pkg1.G
type G[P Constraint] = pkg1.G[P]
