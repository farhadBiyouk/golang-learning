package main

import "fmt"

type Runner interface {
	run()
}

type Shooter interface {
	shoot()
}

type Player struct {
	Name     string
	Age      int
	Height   float64
	Weight   float64
	Position string
}

func main() {
	player := &Player{Name: "messi", Age: 34, Height: 1.76, Weight: 70.23, Position: "Front"}
	var runner Runner = player
	runner.run()

	var shooter Shooter = player
	shooter.shoot()
}

func (player *Player) run() {
	fmt.Printf("%s in position %s, player is running\n", player.Name, player.Position)
}
func (player *Player) shoot() {
	fmt.Printf("%s in position %s, player is shooting\n", player.Name, player.Position)
}
