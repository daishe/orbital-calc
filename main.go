package main

import (
	"flag"
	"fmt"
	"math"
	"time"
)

func pow2(x float64) float64 {
	return x * x
}

func floatTimeToString(x float64) string {
	return fmt.Sprintf("%f (%s)", x, time.Duration(x*float64(time.Second)))
}

func angleToString(x float64) string {
	return fmt.Sprintf("%f (%f res / %f deg)", x, x/(math.Pi*2), x*360/(math.Pi*2))
}

func main() {
	var orbitPeriod time.Duration
	var numberOfSatelites uint
	var sateliteNumber uint

	var baseLong float64
	var baseLat float64
	var nLong float64
	var nLat float64

	flag.DurationVar(&orbitPeriod, "orbit-period", 0, "time of single orbit")

	flag.UintVar(&numberOfSatelites, "satelite-count", 0, "total statelites in ring")
	flag.UintVar(&sateliteNumber, "satelite-number", 0, "number of target satelite in ring")

	flag.Float64Var(&baseLong, "first-satelite-longditude", 0, "longditude of first satelite in ring")
	flag.Float64Var(&baseLat, "first-satelite-latitude", 0, "latitude of first satelite in ring")
	flag.Float64Var(&nLong, "satelite-longditude", 0, "longditude of target satelite in ring")
	flag.Float64Var(&nLat, "satelite-latitude", 0, "latitude of target satelite in ring")
	flag.Parse()

	sateliteNumber--                          // make zero for first sat
	baseLong = baseLong / 360 * (2 * math.Pi) // to radians
	baseLat = baseLat / 360 * (2 * math.Pi)   // to radians
	nLong = nLong / 360 * (2 * math.Pi)       // to radians
	nLat = nLat / 360 * (2 * math.Pi)         // to radians

	dLong := baseLong - nLong // delta long
	dLat := baseLat - nLat    // delta lat

	expectedAngle := 2 * math.Pi / float64(numberOfSatelites)
	nExpectedAngle := expectedAngle * float64(sateliteNumber)

	orbitPeriodInSeconds := float64(orbitPeriod) / float64(time.Second)
	expectedAngleTime := orbitPeriodInSeconds / float64(numberOfSatelites)
	nExpectedAngleTime := expectedAngleTime * float64(sateliteNumber)

	fmt.Printf("Delta longditude: %s\n", angleToString(dLong))
	fmt.Printf("Delta latitude: %s\n", angleToString(dLat))

	rCon2Long := 1 / pow2(math.Cos(dLong))
	rCon2Lat := 1 / pow2(math.Cos(dLat))

	nActualAngle := math.Acos(math.Sqrt(1 / (rCon2Long + rCon2Lat - 1)))

	fmt.Printf("Expected separation angle: %s\n", angleToString(nExpectedAngle))
	fmt.Printf("Actual separation angle: %s\n", angleToString(nActualAngle))

	nActualAngleTime := orbitPeriodInSeconds * nActualAngle / (2 * math.Pi)

	fmt.Printf("Expected separation time: %s\n", floatTimeToString(nExpectedAngleTime))
	fmt.Printf("Actual separation time: %s\n", floatTimeToString(nActualAngleTime))

	nAngleTimeDelta := nExpectedAngleTime - nActualAngleTime
	fmt.Printf("Separation time delta: %s\n", floatTimeToString(nAngleTimeDelta))
}
