package main

import (
	"github.com/nsf/termbox-go"
)

// Dibuja el titulo
func DrawTitle(title string, x, y int) {
	for i, ch := range title {
		termbox.SetCell(x+i, y, ch, termbox.ColorWhite, termbox.ColorBlack)
	}
}

// Dibuja el checkbox (no se como se dice en español)
func DrawCheckbox(x, y int, checked bool) {

	if checked {
		termbox.SetCell(x, y, '[', termbox.ColorGreen, termbox.ColorBlack)

		termbox.SetCell(x+1, y, 'X', termbox.ColorGreen, termbox.ColorBlack)

		termbox.SetCell(x+2, y, ']', termbox.ColorGreen, termbox.ColorBlack)

		return
	}

	termbox.SetCell(x, y, '[', termbox.ColorWhite, termbox.ColorBlack)

	termbox.SetCell(x+1, y, ' ', termbox.ColorWhite, termbox.ColorBlack)

	termbox.SetCell(x+2, y, ']', termbox.ColorWhite, termbox.ColorBlack)

}

// Dibuja el checkbox (no se como se dice en español)
func DrawCheckboxSpecial(x, y int) {
	termbox.SetCell(x, y, '[', termbox.ColorRed, termbox.ColorBlack)

	termbox.SetCell(x+1, y, '*', termbox.ColorRed, termbox.ColorBlack)

	termbox.SetCell(x+2, y, ']', termbox.ColorRed, termbox.ColorBlack)

}

// This function draw in terminal... OK?
// Ya me dio flojera traducir mis comentarios... nuevamente
func DrawInTerminalWeeks() {

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	weeksCompleted, remainingWeeks := CalculateChristmasWeather()

	// Obtener el tamaño de la terminal
	width, height := termbox.Size()

	// Título
	title := "Semanas que faltan para Navidad"
	DrawTitle(title, (width-len(title))/2, 1)

	// Calcular cuántos checkboxes caben horizontalmente
	numCheckboxesPerRow := width / 4 // Cada checkbox ocupa 4 caracteres (cuadro + espacio)

	// Dibujar checkboxes
	for i := 0; i < weeksCompleted+remainingWeeks; i++ { // Dibujar un máximo de 20 checkboxes
		row := i / numCheckboxesPerRow
		col := i % numCheckboxesPerRow
		x := col * 4   // Posición horizontal para cada checkbox
		y := row*2 + 3 // Posición vertical para cada fila de checkboxes, empezando después del título y separados por 1 línea

		if y >= height {
			break // Si superamos el tamaño vertical de la terminal, salimos del bucle
		}

		// Si menor a las semanas completadas entoces dibuja un checkbox marcado
		if i < weeksCompleted {
			DrawCheckbox(x, y, true)
		}

		// Si mayor a las semanas completadas entoces dibuja un checkbox desmarcado y tambien difrente al ultmo
		if i >= weeksCompleted && i != (weeksCompleted+remainingWeeks)-1 {
			DrawCheckbox(x, y, false)
		}

		// Si la semana es la de navidad :D
		if i == (weeksCompleted+remainingWeeks)-1 {
			DrawCheckboxSpecial(x, y)
		}

	}

	termbox.Flush()

	// Esperar a que el usuario presione una tecla para salir
	termbox.PollEvent()

}

func DrawInTerminalChristmas() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	// Obtener el tamaño de la terminal
	width, height := termbox.Size()

	if !IsChristmas() {

		// Título
		title := "No es Navidad :C"
		titleX := (width - len(title)) / 2
		titleY := height / 2

		// Dibujar el título centrado
		DrawTitle(title, titleX, titleY)

	} else {

		title := "Si es Navidad :D"
		titleX := (width - len(title)) / 2
		titleY := height / 2

		for i, ch := range title {
			termbox.SetCell(titleX+i, titleY, ch, termbox.ColorRed, termbox.ColorBlack)
		}
	}

	termbox.Flush()

	// Esperar a que el usuario presione una tecla para salir
	termbox.PollEvent()
}
