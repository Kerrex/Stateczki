package objects

type Ship struct {
	GameObject
	Shape    Triangle
}

func (ship *Ship) GetLineSegmentsInWorld() []LineSegment {
	triangleMovedByShipPosition := ship.GetShapeInWorld()
	return []LineSegment{
		{
			X: triangleMovedByShipPosition.X,
			Y: triangleMovedByShipPosition.Y,
		}, {
			X: triangleMovedByShipPosition.X,
			Y: triangleMovedByShipPosition.Z,
		}, {
			X: triangleMovedByShipPosition.Y,
			Y: triangleMovedByShipPosition.Z,
		},
	}
}

func (ship *Ship) GetShapeInWorld() Triangle {
	rotatedTriangle := Triangle{
		X: rotateVector(ship.Shape.X, ship.Rotation),
		Y: rotateVector(ship.Shape.Y, ship.Rotation),
		Z: rotateVector(ship.Shape.Z, ship.Rotation),
	}

	return rotatedTriangle.GetMovedTriangle(ship.Position)
}
