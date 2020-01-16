package main

import "backend/objects"

func main() {
	println("test")
	triangle := objects.Triangle{
		X: objects.Vector2{
			X: 10,
			Y: 0,
		},
		Y: objects.Vector2{},
		Z: objects.Vector2{},
	}

	moved := triangle.GetMovedTriangle(objects.Vector2{
		X: 5,
		Y: 10,
	})

	println(moved.X.X)
}
