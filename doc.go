/*
Package c80 provides an ‘engine’ for creating a fantasy console computer geared towards video games
that is inspired by arcade game machines, home video game consoles, handheld game consoles, and (other) computers from the 1980s.

Example

	import "github.com/reiver/go-c80"
	
	// ...
	
	c80.Begin()
	
	c80.Draw(Map(mapX, mapY, width,height), 0, 0)
	c80.Draw(Sprite("32x32", id),           x, y)
	
	c80.End()



*/
package c80
