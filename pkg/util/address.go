package util

import (
	"fmt"
	"math/rand"
)

type Address struct {
	StreetAddress string
	City          string
	Lat           float64
	Long          float64
}

var (
	sheikhZayedLat   = 30.0444
	sheikhZayedLong  = 31.2357
	mohandseenLat    = 30.0587
	mohandseenLong   = 31.2048
	latDiff          = 0.02
	longDiff         = 0.02
	sheikhZayedRange = [2][2]float64{
		{sheikhZayedLat - latDiff, sheikhZayedLong - longDiff},
		{sheikhZayedLat + latDiff, sheikhZayedLong + longDiff},
	}
	mohandseenRange = [2][2]float64{
		{mohandseenLat - latDiff, mohandseenLong - longDiff},
		{mohandseenLat + latDiff, mohandseenLong + longDiff},
	}
)

func GenerateRandomAddress() Address {
	var (
		streetAddress string
		city          string
		latitude      float64
		longitude     float64
	)

	if rand.Intn(2) == 0 {
		streetAddress = fmt.Sprintf("%d %s street", rand.Intn(100), "sheikh zayed")
		city = "Giza"
		latitude = randFloatInRange(sheikhZayedRange[0][0], sheikhZayedRange[1][0])
		longitude = randFloatInRange(sheikhZayedRange[0][1], sheikhZayedRange[1][1])
	} else {
		streetAddress = fmt.Sprintf("%d %s street", rand.Intn(100), "mohandseen")
		city = "Giza"
		latitude = randFloatInRange(mohandseenRange[0][0], mohandseenRange[1][0])
		longitude = randFloatInRange(mohandseenRange[0][1], mohandseenRange[1][1])
	}

	return Address{
		StreetAddress: streetAddress,
		City:          city,
		Lat:           latitude,
		Long:          longitude,
	}
}

func randFloatInRange(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

// func RandomAddress() string {
// 	addresses := []string{
// 		"Sheikh Zayed, Giza, Egypt",
// 		"Sheikh Zayed, 6th of October, Egypt",
// 		"Mohandseen, Giza, Egypt",
// 		"Mohandseen, Dokki, Giza, Egypt",
// 	}

// 	return addresses[rand.Intn(len(addresses))]
// }
