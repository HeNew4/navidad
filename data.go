package main

import (
	"time"
)

// This function return true if now is christmas.
// Sadly, today is not Christmas, but maybe tomorrow (It won't be Christmas either).
func IsChristmas() bool {

	dateNow := time.Now()

	christmas := time.Date(dateNow.Year(), time.December, 25, 0, 0, 0, 0, dateNow.Location())

	if dateNow.Year() == christmas.Year() && dateNow.Month() == christmas.Month() && dateNow.Day() == christmas.Day() {
		return true
	}

	return false

}

// # CalculateChristmasWeather
//
// This function calculates the time remaining until Christmas.
//
// return in this order:
//
// - Weeks completed
//
// - Remaining weeks
func CalculateChristmasWeather() (int, int) {

	// Done :D
	done := make(chan bool, 2)

	// This save date now
	dateNow := time.Now()

	// This variable save start date
	startDate := time.Date(dateNow.Year(), time.January, 0, 0, 0, 0, 0, dateNow.Location())
	endDay := startDate.AddDate(0, 12, -1)

	// This variable save christmas date
	christmas := time.Date(dateNow.Year(), time.December, 25, 0, 0, 0, 0, dateNow.Location())

	// Init variables for weeks completed and remaining weeks
	var weeksCompleted int
	var remainingWeeks int

	// First Go rutine
	go func() {

		var initCount int = 0

		// Ya me dio flojera traducir mis comentarios al ingles
		// Bucle para poder calcular las semanas transcrurridas en la fecha actual
		for date := startDate; date.Before(endDay.AddDate(0, 0, 1)); date = date.AddDate(0, 0, 1) {
			if dateNow.Day() == date.Day() && dateNow.Month() == date.Month() {
				weeksCompleted = initCount / 7
				break
			}

			initCount++
		}

		done <- true

	}()

	// Secod Go rutine
	go func() {

		var initCount int = 0

		// Ya me dio flojera traducir mis comentarios al ingles
		// Bucle para poder calcular las semanas que faltan para navidad
		for date := dateNow; date.Before(christmas.AddDate(0, 0, 1)); date = date.AddDate(0, 0, 1) {
			if christmas.Day() == date.Day() && christmas.Month() == date.Month() {
				remainingWeeks = initCount / 7
			}

			initCount++
		}

		done <- true

	}()

	<-done
	<-done

	return weeksCompleted, remainingWeeks
}

// # CalculateChristmasWeather
//
// This function calculates the time remaining until Christmas. (Funcion con la que me papeo chatGPT :c Lo pongo namas)
//
// return in this order:
//
// - Weeks completed
//
// - Remaining weeks
func CalculateChristmasWeatherAI() (int, int) {

	dateNow := time.Now()
	startOfYear := time.Date(dateNow.Year(), time.January, 1, 0, 0, 0, 0, dateNow.Location())
	christmas := time.Date(dateNow.Year(), time.December, 25, 0, 0, 0, 0, dateNow.Location())

	daysCompleted := dateNow.Sub(startOfYear).Hours() / 24
	weeksCompleted := int(daysCompleted) / 7

	daysRemaining := christmas.Sub(dateNow).Hours() / 24
	remainingWeeks := int(daysRemaining) / 7

	return weeksCompleted, remainingWeeks

}
