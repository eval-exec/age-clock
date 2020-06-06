package main

import (
	"time"
)

type Person struct {
	BirthDay time.Time
}

func (p *Person) Age() float64 {
	ageSec := time.Now().Sub(p.BirthDay)
	age := ageSec.Seconds() / 60 / 60 / 24 / 365
	return age
}
