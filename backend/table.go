package main

import (
	"math"

	"github.com/pkg/errors"
)

// Table .
type Table struct {
	X, Y, relHeight float64
	CurrentLine     int
	Columns         [][]Column
	TextAlignment   string
	LineHeight      float64
}

// NewTable create new Table.
func NewTable(x, y float64, columns []Column) *Table {
	cols := make([][]Column, 0)
	cols = append(cols, columns)
	table := &Table{
		X:          x,
		Y:          y,
		Columns:    cols,
		LineHeight: 5.5,
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
	maxHeight := 0.0
	for i, col := range t.Columns[len(t.Columns)-1] {
		lines := pdf.SplitLines([]byte(content[i]), col.Width)
		maxHeight = math.Max(float64(len(lines))*t.LineHeight+col.Height, maxHeight)
		columns = append(columns, Column{Content: content[i], Width: col.Width, Height: maxHeight})
	}

	t.Columns = append(t.Columns, columns)
	t.printRow(columns)

	return err
}

func (t *Table) printRow(columns []Column) {
	lastHeight := 0.0
	nextX := t.X
	for _, col := range columns {
		// pdf.Rect(nextX, t.Y+t.relHeight, col.Width, col.Height, "D")
		// fmt.Println(i)
		// fmt.Println(nextX, relW, col.Width, float64(i))
		// fmt.Println("")

		// Add Bold to header's font.
		alignStr := "L"
		cellYCentering := 0.0
		if t.CurrentLine == 0 {
			pdf.SetFont("Arial", "B", 12)
		} else {
			alignStr = "L"
			pdf.SetFont("Arial", "", 12)
			cellYSpacing = 1.1
			cellYCentering = t.LineHeight
		}

		lines := pdf.SplitLines([]byte(col.Content), col.Width)
		pdf.SetXY(nextX, t.Y+t.relHeight)
		if len(lines) > 1 {
			nextCellH := 0.0
			for _, v := range lines {
				pdf.SetXY(nextX, t.Y)
				pdf.CellFormat(col.Width, col.Height-cellYCentering, string(v), "1", 0, alignStr, false, 0, "")
				nextCellH += t.LineHeight
			}
		} else {
			pdf.CellFormat(col.Width, col.Height, col.Content, "1", 0, alignStr, false, 0, "")
		}

		// if i == 0 {
		// 	break
		// }

		lastHeight = col.Height
		nextX = nextX + col.Width
	}

	pdf.Ln(10)
	t.relHeight = t.relHeight + lastHeight
	t.CurrentLine++
}

// Column of Table.
type Column struct {
	Width, Height float64
	Content       string
}
