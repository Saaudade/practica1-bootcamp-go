package main

import "fmt"

var temperatura_grados int
var humedad_porcentaje int
var presion_hPa int

func main() {
	temperatura_grados = 19
	humedad_porcentaje = 94
	presion_hPa = 1014
	fmt.Println("En Palmira se tiene la siguiente información: ")
	fmt.Println("Temperatura:", temperatura_grados, "°")
	fmt.Println("Humedad:", humedad_porcentaje, "%")
	fmt.Println("Presion:", presion_hPa, "hPa")
}
