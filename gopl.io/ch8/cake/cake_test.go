package cake_test

import (
	"gopl.io/ch8/cake"
	"testing"
	"time"
)

var defaults = cake.Shop{
	Verbose:      testing.Verbose(),
	Cakes:        20,
	BakeTime:     10 * time.Millisecond,
	NumIcers:     1,
	IceTime:      10 * time.Millisecond,
	InscribeTime: 10 * time.Millisecond,
}

func BenchmarkWork(b *testing.B) {
	cakeShop := defaults
	cakeShop.Work(b.N)
}

func BenchmarkBuffers(b *testing.B) {
	cakeShop := defaults
	cakeShop.BakeBuf = 10
	cakeShop.IceBuf = 10
	cakeShop.Work(b.N)
}

func BenchmarkVariable(b *testing.B) {
	cakeShop := defaults
	cakeShop.BakeStdDev = cakeShop.BakeTime / 4
	cakeShop.IceStdDev = cakeShop.IceTime / 4
	cakeShop.InscribeStdDev = cakeShop.InscribeTime / 4
	cakeShop.Work(b.N)
}

func BenchmarkVariableBuffers(b *testing.B) {
	cakeShop := defaults
	cakeShop.BakeStdDev = cakeShop.BakeTime / 4
	cakeShop.IceStdDev = cakeShop.IceTime / 4
	cakeShop.InscribeStdDev = cakeShop.InscribeTime / 4
	cakeShop.BakeBuf = 10
	cakeShop.IceBuf = 10
	cakeShop.Work(b.N)
}

func BenchmarkSlowIcing(b *testing.B) {
	cakeShop := defaults
	cakeShop.IceTime = 50 * time.Millisecond
	cakeShop.Work(b.N)
}

func BenchmarkSlowIcingManyIcers(b *testing.B) {
	cakeShop := defaults
	cakeShop.IceTime = 50 * time.Millisecond
	cakeShop.NumIcers = 5
	cakeShop.Work(b.N)
}
