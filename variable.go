package main

import "fmt"

var (
	hp                 = 100
	attack             = 40
	defence            = 20
	damageRate float32 = 0.17
	damage             = float32(attack-defence) * damageRate
)

func main() {
	leftHp := float32(hp) - damage
	fmt.Println("leftHp is %f\n", leftHp)
}
