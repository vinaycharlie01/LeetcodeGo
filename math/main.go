package main

func FindDelayedArrivalTime(arrivalTime int, delayedTime int) int {
	totaltime := arrivalTime + delayedTime
	switch {
	case totaltime >= 24:
		return totaltime - 24
	default:
		return totaltime
	}
}
