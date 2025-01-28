package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var (
	currentInput string
	result       float64
	operator     string
	fullEquation string // Stores the full equation for display
)

func main() {
	// Create a new Fyne application
	myApp := app.New()
	myWindow := myApp.NewWindow("Calculator")

	// Display for the calculator
	display := widget.NewLabel("0")
	display.Alignment = fyne.TextAlignTrailing

	// Function to update the display with the full equation
	updateDisplay := func(text string) {
		fullEquation += text
		display.SetText(fullEquation)
	}

	// Function to handle digit buttons
	handleDigit := func(digit string) {
		currentInput += digit
		updateDisplay(digit)
	}

	// Function to handle operator buttons
	handleOperator := func(op string) {
		if currentInput != "" {
			num, _ := strconv.ParseFloat(currentInput, 64)
			if operator == "" {
				result = num
			} else {
				switch operator {
				case "+":
					result += num
				case "-":
					result -= num
				case "*":
					result *= num
				case "/":
					result /= num
				}
			}
			operator = op
			updateDisplay(" " + op + " ") // Add operator to the equation
			currentInput = ""
		}
	}

	// Function to handle the equals button
	handleEquals := func() {
		if currentInput != "" && operator != "" {
			num, _ := strconv.ParseFloat(currentInput, 64)
			switch operator {
			case "+":
				result += num
			case "-":
				result -= num
			case "*":
				result *= num
			case "/":
				result /= num
			}
			updateDisplay(" = " + fmt.Sprintf("%.2f", result)) // Add result to the equation
			currentInput = ""
			operator = ""
		}
	}

	// Function to clear the calculator
	handleClear := func() {
		currentInput = ""
		result = 0
		operator = ""
		fullEquation = ""
		display.SetText("0")
	}

	// Create buttons for digits 0-9
	digitButtons := make([]fyne.CanvasObject, 10)
	for i := 0; i < 10; i++ {
		digit := strconv.Itoa(i)
		digitButtons[i] = widget.NewButton(digit, func() {
			handleDigit(digit)
		})
	}

	// Create buttons for operators
	operatorButtons := map[string]*widget.Button{
		"+": widget.NewButton("+", func() { handleOperator("+") }),
		"-": widget.NewButton("-", func() { handleOperator("-") }),
		"*": widget.NewButton("*", func() { handleOperator("*") }),
		"/": widget.NewButton("/", func() { handleOperator("/") }),
		"=": widget.NewButton("=", handleEquals),
		"C": widget.NewButton("C", handleClear),
	}

	// Layout for number buttons (0-9)
	numberPanel := container.NewGridWithColumns(3,
		digitButtons[7], digitButtons[8], digitButtons[9],
		digitButtons[4], digitButtons[5], digitButtons[6],
		digitButtons[1], digitButtons[2], digitButtons[3],
		digitButtons[0],
	)

	// Layout for operator buttons
	operatorPanel := container.NewVBox(
		operatorButtons["C"],
		operatorButtons["/"],
		operatorButtons["*"],
		operatorButtons["-"],
		operatorButtons["+"],
		operatorButtons["="],
	)

	// Create a top bar with the display
	topBar := container.NewBorder(nil, nil, nil, nil, display)

	// Combine the top bar, number panel, and operator panel
	content := container.NewBorder(
		topBar,        // Top: Display
		nil,           // Bottom: Nothing
		nil,           // Left: Nothing
		operatorPanel, // Right: Operator buttons
		numberPanel,   // Center: Number buttons
	)

	// Set the window content and size
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(200, 300))
	myWindow.ShowAndRun()
}
