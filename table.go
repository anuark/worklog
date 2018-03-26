package main

import (
	"fmt"

	"github.com/pkg/errors"
)

// Table .
type Table struct {
	X, Y, relHeight float64
	CurrentLine     int
	Columns         [][]Column
	TextAlignment   string
}

// NewTable create new Table.
func NewTable(x, y float64, columns []Column) *Table {
	cols := make([][]Column, 0)
	cols = append(cols, columns)
	table := &Table{
		X:       x,
		Y:       y,
		Columns: cols,
	}

	table.printRow(columns)

	return table
}

// AddRow to Table.
func (t *Table) AddRow(content []string) error {
	var err error
	if len(content) != len(t.Columns[0]) {
		return errors.Wrap(err, "Not same column count")
	}

	columns := make([]Column, 0)
	for i, col := range t.Columns[len(t.Columns)-1] {
		columns = append(columns, Column{Content: content[i], Width: col.Width, Height: col.Height})
	}

	t.Columns = append(t.Columns, columns)
	t.printRow(columns)

	return err
}

func (t *Table) printRow(columns []Column) {
	lastHeight := 0.0
	nextX := t.X
	for _, col := range columns {
		pdf.Rect(nextX, t.Y+t.relHeight, col.Width, col.Height, "D")
		// fmt.Println(i)
		// fmt.Println(nextX, relW, col.Width, float64(i))
		// fmt.Println("")

		// Add Bold to header's font.
		if t.CurrentLine == 0 {
			pdf.SetFont("Arial", "B", 12)
		} else {
			pdf.SetFont("Arial", "", 12)
		}

		calcDimensions(col)

		pdf.CellFormat(col.Width, col.Height, col.Content, "", 0, "L", false, 0, "")

		// if i == 0 {
		// 	break
		// }

		lastHeight = col.Height
		nextX = nextX + col.Width
	}

	pdf.Ln(10)
	t.relHeight = lastHeight * float64(len(t.Columns))
	t.CurrentLine++
}

const letterWidth = 1
const descriptionWidth = 60

func calcDimensions(col Column) {
	totalLength := len(col.Content) * letterWidth
	if totalLength > descriptionWidth {
		fmt.Println("greater")
	}
}

// Column of Table.
type Column struct {
	Width, Height float64
	Content       string
}
