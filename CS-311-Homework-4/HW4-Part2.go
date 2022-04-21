// coordinates for a 2-dimensional point 
type Point struct { x, y float64 } 
// Produces "direction" of parameter point from this point 
func (pt *Point) dir(other *Point) (result *Point) { 
	result = new(Point) 
	xdiff := other.x - pt.x 
	ydiff := other.y - pt.y 
	if xdiff > 0 { result.x = 1 } 
	else if xdiff < 0 { result.x = -1 } 
	if ydiff > 0 { result.y = 1 } 
	else if ydiff < 0 { result.y = -1 } 
	return 
} 

func (pt *Point) printSomeDir(other *Point) { 
	dirPt := pt.dir(other) 
	
	// Q3: second time you reach the line below 
	if dirPt.x > 0.9 && math.Abs(dirPt.y) < 0.1 { 
		fmt.Println("East") 
	} else if dirPt.x < -0.9 && math.Abs(dirPt.y) < 0.1 { 
		fmt.Println("West") 
	} 
} 

func main() { 
	point1 := new(Point) 
	point1.x, point1.y = 0, 0 
	point2 := new(Point) 
	point2.x, point2.y = 3, 0 
	point1.printSomeDir(point2) 
	
	point2 = new(Point)
	point2.x, point2.y = -7, 0 
	point1.printSomeDir(point2) 
	// Q4: first time you reach this line 
	fmt.Println("done") 
} 