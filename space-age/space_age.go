package space

import "fmt"

type Planet string


func Age(seconds int , planet string) int {

	var secInEarth: float64 = seconds / 31557600

	if planet == "Mercury" {
		fmt.Println(secInEarth * 0.2408467)
	}
/*
Earth: orbital period 365.25 Earth days, or 31557600 seconds
Mercury: orbital period 0.2408467 Earth years
Venus: orbital period 0.61519726 Earth years
Mars: orbital period 1.8808158 Earth years
Jupiter: orbital period 11.862615 Earth years
Saturn: orbital period 29.447498 Earth years
Uranus: orbital period 84.016846 Earth years
Neptune: orbital period 164.79132 Earth years*/
}
