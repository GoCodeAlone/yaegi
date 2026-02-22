package p1

import "math/rand"

var rng = rand.New(rand.NewSource(1))

var Uint32 = rng.Uint32
