package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
--- Day 2: I Was Told There Would Be No Math ---

The elves are running low on wrapping paper, and so they need to submit an order for more. They have a list of the dimensions (length l, width w, and height h) of each present, and only want to order exactly as much as they need.

Fortunately, every present is a box (a perfect right rectangular prism), which makes calculating the required wrapping paper for each gift a little easier: find the surface area of the box, which is 2*l*w + 2*w*h + 2*h*l. The elves also need a little extra paper for each present: the area of the smallest side.

For example:

A present with dimensions 2x3x4 requires 2*6 + 2*12 + 2*8 = 52 square feet of wrapping paper plus 6 square feet of slack, for a total of 58 square feet.
A present with dimensions 1x1x10 requires 2*1 + 2*10 + 2*10 = 42 square feet of wrapping paper plus 1 square foot of slack, for a total of 43 square feet.
All numbers in the elves' list are in feet. How many total square feet of wrapping paper should they order?
*/

type Parcel struct {
	length int
	width  int
	height int
}

type Materials struct {
	wrappingArea int
	ribbonLength int
	bowLength    int
}

func (p *Parcel) String() string {
	return fmt.Sprintf("(%dx%dx%d)", p.length, p.width, p.height)
}

func (p *Parcel) LArea() int {
	return p.length * p.width
}

func (p *Parcel) WArea() int {
	return p.width * p.height
}

func (p *Parcel) HArea() int {
	return p.height * p.length
}

func (p *Parcel) LWPerimeter() int {
	return 2*p.length + 2*p.width
}

func (p *Parcel) LHPerimeter() int {
	return 2*p.length + 2*p.height
}

func (p *Parcel) WHPerimeter() int {
	return 2*p.width + 2*p.height
}

func (p *Parcel) FullSurfaceArea() int {
	return 2*p.LArea() + 2*p.WArea() + 2*p.HArea()
}

func (p *Parcel) SmallestSideArea() int {
	return min(p.LArea(), min(p.WArea(), p.HArea()))
}

func (p *Parcel) getTotalWrappingArea() int {
	return p.FullSurfaceArea() + p.SmallestSideArea()
}

func (p *Parcel) getRibbonLength() int {
	return min(p.LWPerimeter(), min(p.LHPerimeter(), p.WHPerimeter()))
}

func (p *Parcel) getBowLength() int {
	return p.length * p.width * p.height
}

func (p *Parcel) getTotalWrappingMaterials() Materials {
	return Materials{p.getTotalWrappingArea(), p.getRibbonLength(), p.getBowLength()}
}

type Line string

func (l *Line) parse() *Parcel {
	var length, width, height int
	_, err := fmt.Sscanf(string(*l), "%dx%dx%d", &length, &width, &height)
	if err != nil {
		return nil
	}
	return &Parcel{length, width, height}
}

func readLines(filename string) []Line {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []Line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, Line(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return lines
}

func main() {
	totalWrappingPaper := 0
	totalRibbonLength := 0

	for _, line := range readLines("./2/puzzle.txt") {
		if parcel := line.parse(); parcel != nil {
			materials := parcel.getTotalWrappingMaterials()
			totalWrappingPaper += materials.wrappingArea
			totalRibbonLength += materials.ribbonLength + materials.bowLength

		}
	}

	fmt.Printf("Result: %d paper, %d ribbon+bow", totalWrappingPaper, totalRibbonLength)
}
