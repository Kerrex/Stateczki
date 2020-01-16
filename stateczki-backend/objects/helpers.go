package objects

import "math"

var One = Vector2{
	X: 1,
}

var Zero = Vector2{}

type Vector2 struct {
	X float32
	Y float32
}

type Triangle struct {
	X Vector2
	Y Vector2
	Z Vector2
}

func (triangle Triangle) GetMovedTriangle(vector2 Vector2) Triangle {
	return Triangle{
		X: Vector2{
			X: triangle.X.X + vector2.X,
			Y: triangle.X.Y + vector2.Y,
		},
		Y: Vector2{
			X: triangle.Y.X + vector2.X,
			Y: triangle.Y.Y + vector2.Y,
		},
		Z: Vector2{
			X: triangle.Z.X + vector2.X,
			Y: triangle.Z.Y + vector2.Y,
		},
	}
}

type Rectangle struct {
	X Vector2
	Y Vector2
	Z Vector2
	W Vector2
}

func (rectangle Rectangle) GetMovedRectangle(vector2 Vector2) Rectangle {
	return Rectangle{
		X: Vector2{
			X: rectangle.X.X + vector2.X,
			Y: rectangle.X.Y + vector2.Y,
		},
		Y: Vector2{
			X: rectangle.Y.X + vector2.X,
			Y: rectangle.Y.Y + vector2.Y,
		},
		Z: Vector2{
			X: rectangle.Z.X + vector2.X,
			Y: rectangle.Z.Y + vector2.Y,
		},
		W: Vector2{
			X: rectangle.W.X + vector2.X,
			Y: rectangle.W.Y + vector2.Y,
		},
	}
}

type LineSegment struct {
	X Vector2
	Y Vector2
}

func rotateVector(vectorToRotate Vector2, degreesToRotate float32) Vector2 {
	angleDiffInRadian := degreesToRotate * math.Pi / 180
	x := math.Cos(float64(angleDiffInRadian)) * float64(vectorToRotate.X) - math.Sin(float64(angleDiffInRadian)) * float64(vectorToRotate.Y)
	y := math.Sin(float64(angleDiffInRadian)) * float64(vectorToRotate.X) + math.Cos(float64(angleDiffInRadian)) * float64(vectorToRotate.Y)

	return Vector2{
		X: float32(x),
		Y: float32(y),
	}
}
