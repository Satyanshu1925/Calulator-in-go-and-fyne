package main

import (
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/knetic/govaluate"
)

func main() {
	a := app.New()                 // create an app
	w := a.NewWindow("Calculator by SATYANSHU") // create a window
	// 	w.Resize(fyne.NewSize(500,280));	Resize not working for now
	output := ""
	input := widget.NewLabel(output) // Input
	historyString := ""
	isHistory := false
	history := widget.NewLabel(historyString)
	var historyArray []string
	// ROW 1
	historyBtn := widget.NewButton("history", func() { //history button
		if isHistory {
			historyString = ""
		} else {
			for i := len(historyArray) - 1; i >= 0; i-- {
				historyString += historyArray[i]
				historyString += "\n"
			}
		}
		isHistory = !isHistory
		input.SetText(historyString)
	})

	backBtn := widget.NewButton("Back", func() { //back button
		if len(output) > 0 {
			output = output[:len(output)-1]
		}
		input.SetText(output)
	})

	//ROW 2
	clearBtn := widget.NewButton("Clear", func() { //clear button
		output = ""
		input.SetText(output)
	})

	openBracketBtn := widget.NewButton("(", func() { // open bracket button
		output = output + ")"
		input.SetText(output)
	})

	closeBracketBtn := widget.NewButton(")", func() { // close bracket button
		output = output + ")"
		input.SetText(output)
	})

	divideBtn := widget.NewButton("/", func() { //divide button
		output = output + "/"
		input.SetText(output)
	})

	//ROW 3
	sevenBtn := widget.NewButton("7", func() { // 7 button
		output = output + "7"
		input.SetText(output)
	})

	eightBtn := widget.NewButton("8", func() { // 8 button
		output = output + "8"
		input.SetText(output)
	})

	nineBtn := widget.NewButton("9", func() { // 9 button
		output = output + "9"
		input.SetText(output)
	})

	multiBtn := widget.NewButton("*", func() { // + button
		output = output + "*"
		input.SetText(output)
	})

	//ROW 4
	fourBtn := widget.NewButton("4", func() { // 4 button
		output = output + "4"
		input.SetText(output)
	})

	fiveBtn := widget.NewButton("5", func() { // 5 button
		output = output + "5"
		input.SetText(output)
	})

	sixBtn := widget.NewButton("6", func() { // 6 button
		output = output + "6"
		input.SetText(output)
	})

	minBtn := widget.NewButton("-", func() { // - button
		output = output + "-"
		input.SetText(output)
	})
	//ROW 5
	oneBtn := widget.NewButton("1", func() { // 1 button
		output = output + "1"
		input.SetText(output)
	})

	twoBtn := widget.NewButton("2", func() { // 2 button
		output = output + "2"
		input.SetText(output)
	})

	threeBtn := widget.NewButton("3", func() { // 3 button
		output = output + "3"
		input.SetText(output)
	})

	addBtn := widget.NewButton("+", func() { // + button
		output = output + "+"
		input.SetText(output)
	})

	//Row 4
	zeroBtn := widget.NewButton("0", func() { // + button
		output = output + "0"
		input.SetText(output)
	})
	dotBtn := widget.NewButton(".", func() { // . button
		output = output + "."
		input.SetText(output)
	})
	equalBtn := widget.NewButton("=", func() { // = button

		expression, err := govaluate.NewEvaluableExpression(output)
		if err == nil {
			result, err := expression.Evaluate(nil)
			if err == nil {
				ans := strconv.FormatFloat(result.(float64), 'f', -1, 64)
				strToAppend := output + " = " + ans
				historyArray = append(historyArray, strToAppend)
				output = ans
			} else {
				output = "error"
			}
		} else {
			output = "error"
		}
		// result is now set to "true", the bool value.
		input.SetText(output)
	})
	w.SetContent(container.NewVBox( // It is a container
		input,
		history,
		container.NewGridWithColumns(1,
			container.NewGridWithColumns(2,
				historyBtn,
				backBtn,
			),
			container.NewGridWithColumns(4,
				clearBtn,
				openBracketBtn,
				closeBracketBtn,
				divideBtn,
			),
			container.NewGridWithColumns(4,
				sevenBtn,
				eightBtn,
				nineBtn,
				multiBtn,
			),
			container.NewGridWithColumns(4,
				fourBtn,
				fiveBtn,
				sixBtn,
				minBtn,
			),
			container.NewGridWithColumns(4,
				oneBtn,
				twoBtn,
				threeBtn,
				addBtn,
			),
			container.NewGridWithColumns(2,
				container.NewGridWithColumns(2,
					zeroBtn,
					dotBtn),
				equalBtn,
			),
		),
	))

	w.ShowAndRun()
}
