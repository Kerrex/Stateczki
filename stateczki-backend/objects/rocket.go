package objects

type Rocket struct {
	GameObject
	Shape    Rectangle
}

func (rocket *Rocket) GetLineSegmentsInWorld() []LineSegment {
	triangleMovedByRocketPosition := rocket.GetShapeInWorld()
	return []LineSegment{
		{
			X: triangleMovedByRocketPosition.X,
			Y: triangleMovedByRocketPosition.Y,
		}, {
			X: triangleMovedByRocketPosition.Y,
			Y: triangleMovedByRocketPosition.Z,
		}, {
			X: triangleMovedByRocketPosition.Z,
			Y: triangleMovedByRocketPosition.W,
		}, {
			X: triangleMovedByRocketPosition.X,
			Y: triangleMovedByRocketPosition.W,
		},
	}
}

func (rocket *Rocket) GetShapeInWorld() Rectangle {
	rotatedRectangle := Rectangle{
		X: rotateVector(rocket.Shape.X, rocket.Rotation),
		Y: rotateVector(rocket.Shape.Y, rocket.Rotation),
		Z: rotateVector(rocket.Shape.Z, rocket.Rotation),
		W: rotateVector(rocket.Shape.W, rocket.Rotation),
	}

	return Rectangle{
		X: Vector2{
			X: rotatedRectangle.X.X + rocket.Position.X,
			Y: rotatedRectangle.X.Y + rocket.Position.Y,
		},
		Y: Vector2{
			X: rotatedRectangle.Y.X + rocket.Position.X,
			Y: rotatedRectangle.Y.Y + rocket.Position.Y,
		},
		Z: Vector2{
			X: rotatedRectangle.Z.X + rocket.Position.X,
			Y: rotatedRectangle.Z.Y + rocket.Position.Y,
		},
		W: Vector2{
			X: rotatedRectangle.W.X + rocket.Position.X,
			Y: rotatedRectangle.W.Y + rocket.Position.Y,
		},
	}
}
