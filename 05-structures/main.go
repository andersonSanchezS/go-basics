package main

type Coordinate struct {
	X int
	Y int
}

type Rectangle struct {
	TopLeft     Coordinate
	BottomRight Coordinate
}

func width(r Rectangle) int {
	return r.BottomRight.X - r.TopLeft.X
}

func length(r Rectangle) int {
	return r.TopLeft.Y - r.BottomRight.Y
}

func area(r Rectangle) int {
	return width(r) * length(r)
}

func perimeter(r Rectangle) int {
	return (2 * width(r)) + (2 * length(r))
}

func printInfo(r Rectangle) {
	println("area:", area(r))
	println("perimeter:", perimeter(r))
}

func main() {
	r := Rectangle{
		TopLeft:     Coordinate{X: 0, Y: 7},
		BottomRight: Coordinate{X: 10, Y: 0},
	}
	printInfo(r)

	r.TopLeft.Y *= 2
	r.BottomRight.X *= 2
	printInfo(r)
}
