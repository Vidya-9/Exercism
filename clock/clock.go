package clock

import "fmt"

//Clock is a structure which holds the different type variables
type Clock struct {
	h, m int
}

var c Clock

//the definition of the function New
func New(h, m int) Clock {
	h += m / 60
	m = m % 60
	h %= 24
	if m < 0 {
		h--
		m += 60
	}
	if h < 0 {
		h += 24
	}
	c.h, c.m = h, m
	return c
}

//the definition of the function Clock
func (Clock) String() string {
	return fmt.Sprintf("%.2d:%.2d", c.h, c.m)
}

//Definition  of the function Add which adds the given minutes
func (Clock) Add(m int) Clock {
	return New(c.h, c.m+m)
}

//Definition of the function Sub which subtracts the given minutes
func (Clock) Subtract(m int) Clock {
	return New(c.h, c.m-m)
}
