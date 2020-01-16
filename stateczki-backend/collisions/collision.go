package collisions

import "backend/objects"

func IsColliding(polygonInWorld objects.PolygonInWorld, polygonInWorld2 objects.PolygonInWorld) bool {
	segmentsOne := polygonInWorld.GetLineSegmentsInWorld()
	segmentsTwo := polygonInWorld2.GetLineSegmentsInWorld()

	for _, linesFromPolygonOne := range segmentsOne {
		for _, linesFromPolygonTwo := range segmentsTwo {
			if areLinesIntersecting(linesFromPolygonOne, linesFromPolygonTwo) {
				return true
			}
		}
	}

	return false
}


// Algorithm from https://stackoverflow.com/a/1968345
func areLinesIntersecting(line objects.LineSegment, line2 objects.LineSegment) bool {
	p0_x := float32(line.X.X)
	p0_y := float32(line.X.Y)

	p1_x := float32(line.Y.X)
	p1_y := float32(line.Y.Y)

	p2_x := float32(line2.X.X)
	p2_y := float32(line2.X.Y)

	p3_x := float32(line2.Y.X)
	p3_y := float32(line2.Y.Y)

	var s1_x, s1_y, s2_x, s2_y float32
	s1_x = p1_x - p0_x
	s1_y = p1_y - p0_y
	s2_x = p3_x - p2_x
	s2_y = p3_y - p2_y
	var s, t float32
	s = (-s1_y*(p0_x-p2_x) + s1_x*(p0_y-p2_y)) / (-s2_x*s1_y + s1_x*s2_y)
	t = (s2_x*(p0_y-p2_y) - s2_y*(p0_x-p2_x)) / (-s2_x*s1_y + s1_x*s2_y)

	if s >= 0 && s <= 1 && t >= 0 && t <= 1 {
		// Collision detected
		return true
	}

	return false // No collision
}
