package main

import "strconv"
import "strings"

import "math"

func solution(aircraftEmissions []string, flightItineraries []string, origin string, destination string) float64 {

	for _, itin := range flightItineraries {

		slice := strings.Split(itin, "-")

		km := slice[2]
		aircraftType := slice[3]
		origen := slice[0]
		dest := slice[1]

		for _, airplane := range aircraftEmissions {

			sliceAirplane := strings.Split(airplane, "-")

			airpl := sliceAirplane[0]
			emis := sliceAirplane[1]

			if airpl == aircraftType && dest == destination && origen == origin {

				convertedEmis, _ := strconv.ParseFloat(emis, 64)
				convertedKm, _ := strconv.ParseFloat(km, 64)
				consumption := convertedEmis * convertedKm

				return (math.Floor(consumption*100) / 100)

			}

		}

	}

	return 0

}

func main() {}
