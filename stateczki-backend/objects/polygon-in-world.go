package objects

type GameObject struct {
	Position Vector2
	Rotation float32
	Velocity float32
}


func (object *GameObject) MoveForwardBy(distance float32) {
	moveForwardWithoutRotationVector := Vector2{
		X: distance,
	}

	moveForwardWithRotation := rotateVector(moveForwardWithoutRotationVector, object.Rotation)
	newPosition := Vector2{
		X: object.Position.X + moveForwardWithRotation.X,
		Y: object.Position.Y + moveForwardWithRotation.Y,
	}

	object.Position = newPosition
}

type PolygonInWorld interface {
	GetLineSegmentsInWorld() []LineSegment
}