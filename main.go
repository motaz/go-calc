package main

import (
	"fmt"
	"strconv"
	"strings"

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
		if digit == "." && strings.Contains(currentInput, ".") {
			return // Prevent multiple decimal points
		}
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
					if num == 0 {
						display.SetText("Error: Division by zero")
						return
					}
					result /= num
				}
			}
			operator = op
			updateDisplay(" " + op + " ") // Add operator to the equation
			currentInput = ""
		} else if fullEquation != "" {
			// If there's no new input but a result is available, use the result as the starting value
			operator = op
			fullEquation = fmt.Sprintf("%.2f", result) + " " + op + " "
			display.SetText(fullEquation)
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
				if num == 0 {
					display.SetText("Error: Division by zero")
					return
				}
				result /= num
			}
			updateDisplay(" = " + fmt.Sprintf("%.2f", result)) // Add result to the equation
			currentInput = ""
			operator = ""
			fullEquation = fmt.Sprintf("%.2f", result) // Store the result for further calculations
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

	// Create buttons for digits 0-9 and decimal point
	digitButtons := make([]fyne.CanvasObject, 11)
	for i := 0; i < 10; i++ {
		digit := strconv.Itoa(i)
		digitButtons[i] = widget.NewButton(digit, func() {
			handleDigit(digit)
		})
	}
	digitButtons[10] = widget.NewButton(".", func() {
		handleDigit(".")
	})

	// Create buttons for operators
	operatorButtons := map[string]*widget.Button{
		"+": widget.NewButton("+", func() { handleOperator("+") }),
		"-": widget.NewButton("-", func() { handleOperator("-") }),
		"*": widget.NewButton("*", func() { handleOperator("*") }),
		"/": widget.NewButton("/", func() { handleOperator("/") }),
		"=": widget.NewButton("=", handleEquals),
		"C": widget.NewButton("C", handleClear),
	}

	// Layout for number buttons (0-9 and decimal point)
	numberPanel := container.NewGridWithColumns(3,
		digitButtons[7], digitButtons[8], digitButtons[9],
		digitButtons[4], digitButtons[5], digitButtons[6],
		digitButtons[1], digitButtons[2], digitButtons[3],
		digitButtons[0], digitButtons[10],
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

	// Add keyboard event listener for typed runes (characters)
	myWindow.Canvas().SetOnTypedRune(func(r rune) {
		switch r {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			handleDigit(string(r))
		case '.':
			handleDigit(".")
		case '+', '-', '*', '/':
			handleOperator(string(r))
		case '=':
			handleEquals()
		case 'c', 'C':
			handleClear()
		}
	})

	// Add keyboard event listener for special keys (Enter, Backspace, etc.)
	myWindow.Canvas().SetOnTypedKey(func(event *fyne.KeyEvent) {
		switch event.Name {
		case fyne.KeyEnter, fyne.KeyReturn:
			handleEquals()
		case fyne.KeyBackspace, fyne.KeyDelete:
			handleClear()
		}
	})

	// Show and run the application
	myWindow.ShowAndRun()
}
