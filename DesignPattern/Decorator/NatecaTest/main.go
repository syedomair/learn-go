package main

import (
	"fmt"
	"strconv"
)

type deco interface {
	decorateNumber(i int) (string, int)
}

type noDecoration struct {
}

func (p *noDecoration) decorateNumber(i int) (string, int) {
	return strconv.Itoa(i), i
}

type num3Decoration struct {
	deco deco
}

func (c *num3Decoration) decorateNumber(i int) (string, int) {
	formatedNumber, num := c.deco.decorateNumber(i)
	if num%3 == 0 {
		return "Fizz", num
	}
	return formatedNumber, i
}

type num5Decoration struct {
	deco deco
}

func (c *num5Decoration) decorateNumber(i int) (string, int) {
	formatedNumber, num := c.deco.decorateNumber(i)
	if num%5 == 0 {
		return "Buzz", num
	}
	return formatedNumber, i
}

type num3and5Decoration struct {
	deco deco
}

func (c *num3and5Decoration) decorateNumber(i int) (string, int) {
	formatedNumber, num := c.deco.decorateNumber(i)
	if (num%5) == 0 && (num%3 == 0) {
		return "FizzBuzz", num
	}
	return formatedNumber, i
}
func main() {

	noDeco := &noDecoration{}

	deco5 := &num5Decoration{
		deco: noDeco,
	}
	deco3 := &num3Decoration{
		deco: deco5,
	}
	deco3and5 := &num3and5Decoration{
		deco: deco3,
	}
	for i := 1; i <= 100; i++ {
		formatedNumber, _ := deco3and5.decorateNumber(i)
		fmt.Println(formatedNumber)
	}

}
