package main

import "fmt"

func main() {
	type Weapon int

	const (
		Arrow Weapon = iota
		Shuriken
		SniperRifle
		Rifle
		Blower
	)

	// 输出枚举值
	fmt.Println(Arrow, Shuriken, SniperRifle, Rifle, Blower)
}
