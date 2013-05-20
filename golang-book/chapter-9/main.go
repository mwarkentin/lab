package main

import ("fmt"; "math")

type Shape interface {
    area() float64
    perimiter() float64
}

type Circle struct {
    x, y, r float64
}

type Rectangle struct {
    x1, y1, x2, y2 float64
}

func distance(x1, y1, x2, y2 float64) float64 {
    a := x2 - x1
    b := y2 - y1
    return math.Sqrt(a*a + b*b)
}

func (r *Rectangle) area() float64 {
    l := distance(r.x1, r.y1, r.x1, r.y2)
    w := distance(r.x1, r.y1, r.x2, r.y1)
    return l * w
}

func (r *Rectangle) perimiter() float64 {
    l := distance(r.x1, r.y1, r.x1, r.y2)
    w := distance(r.x1, r.y1, r.x2, r.y1)
    return l*2 + w*2
}

func (c *Circle) area() float64 {
    return math.Pi * c.r*c.r
}

func (c *Circle) perimiter() float64 {
    return 2 * math.Pi * c.r
}

func totalArea(shapes ...Shape) float64 {
    var area float64
    for _, s := range shapes {
        area += s.area()
    }
    return area
}

type Person struct {
    Name string
}

type Android struct {
    Person
    Model string
}

func (p *Person) Talk() {
    fmt.Println("Bleep bloop.")
}

func main() {
    r := Rectangle{0, 0, 10, 10}
    c := Circle{0, 0, 5}

    fmt.Println(r.area(), r.perimiter())
    fmt.Println(c.area(), c.perimiter())
    fmt.Println(totalArea(&c, &r))

    a := new(Android)
    a.Talk()
}
