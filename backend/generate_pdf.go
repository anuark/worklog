package main

import (
	"time"

	"cloud.google.com/go/datastore"
	"github.com/jung-kurt/gofpdf"
)

var pdf *gofpdf.Fpdf

// GeneratePdf .
//
// Related: convert pdf to jpg with imagick https://stackoverflow.com/questions/47492837/golang-convert-pdf-to-image-by-bimg
func GeneratePdf(inputDate time.Time) {
	pdf = gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(100, 10, "Anuar Kilgore")
	pdf.Ln(10)

	// TODO: make this data dynamic coded.
	// address, phone and email.
	pdf.SetFont("Arial", "", 10)
	pdf.Cell(0, 10, "col. los alamos")
	pdf.Ln(5)
	pdf.Cell(0, 10, "san pedro sula, cortes")
	pdf.Ln(5)
	pdf.Cell(0, 10, "21101")
	pdf.Ln(5)
	pdf.Cell(0, 10, "Honduras")
	pdf.Ln(5)
	pdf.Cell(0, 10, "Phone: +504 99467213")
	pdf.Ln(5)
	pdf.Cell(0, 10, "jaicof@gmail.com")
	pdf.Ln(5)
	pdf.Ln(10)

	pdf.SetFont("Arial", "", 12)
	// pdf.Cell(0, 10, "Week 25 - 29 December of 2017")
	// pdf.Ln(10)

	//pdf.Rect(12, 70, 50, 20, "D")
	//pdf.Rect(12, 70, 100, 20, "D")

	//pdf.Rect(10, 70, 30, 10, "D")
	//pdf.Rect(40, 70, 30, 10, "D")
	//pdf.Rect(70, 70, 30, 10, "D")
	//pdf.Rect(100, 70, 30, 10, "D")
	//pdf.Rect(130, 70, 30, 10, "D")

	columns := []Column{
		{Content: "Description", Width: 60, Height: 10},
		{Content: "Day", Width: 30, Height: 10},
		{Content: "Hours", Width: 30, Height: 10},
		{Content: "Rate", Width: 30, Height: 10},
		{Content: "Amount", Width: 30, Height: 10},
	}

	q := datastore.NewQuery("Task").
		// Filter("created >", inputDate).
		Order("created")

	// fmt.Println(inputDate)

	var prevTask Task
	nextH := 0.0
	var allTasks []Task
	var table *Table

	n := 0

	tasks := []Task{
		{Description: "Desc 1", Created: time.Date(2018, 5, 1, 12, 30, 0, 0, time.UTC)},
		{Description: "Desc 2", Created: time.Date(2018, 5, 2, 12, 30, 0, 0, time.UTC)},
		{Description: "Desc 3", Created: time.Date(2018, 5, 7, 12, 30, 0, 0, time.UTC)},
	}

	_ = q
	_ = allTasks
	// for t := dsClient.Run(dsCtx, q); ; {
	// var task Task
	// _, err := t.Next(&task)
	// allTasks = append(allTasks, task)
	// if err == iterator.Done {
	// 	// pdf.Ln(10)
	// 	// pdf.CellFormat(10, 220, "Total Hours:		80", "", 0, "L", false, 0, "")
	// 	// pdf.Ln(5)
	// 	// pdf.CellFormat(10, 220, "Hourly Rate:		$25", "", 0, "L", false, 0, "")
	// 	// pdf.Ln(5)
	// 	// pdf.CellFormat(10, 220, "Total Amount:	$2000.00", "", 0, "L", false, 0, "")
	// 	break
	// }
	// if err != nil {
	// 	log.Fatal(err)
	// }

	for _, task := range tasks {
		newWeek := false
		switch true {
		case prevTask.Description == "":
			pdf.Cell(0, 10, "Week 25 - 29 December of 2017")
			pdf.Ln(10)
			table = NewTable(10, 70+nextH, columns)
			break
		case prevTask.Created.Weekday() > task.Created.Weekday():
			newWeek = true
			break
		}

		_ = table

		if newWeek {
			pdf.AddPage()
			// pdf.CellFormat(10, 0, "New week 1 - 2 Jan", "", 0, "L", false, 0, "")
			// table(10, 70+nextH, 38, 10, content)
			// nextH = 10 * float64(len(content)+1)
			table = NewTable(10, 10+nextH, columns)
		}

		// fmt.Println(newWeek)

		table.AddRow([]string{task.Description, task.Created.Format("Monday"), "8", "$25.00", "$200.00"})
		// content = append(content, []string{task.Description, task.Created.Format("Monday"), "8", "$25.00", "$200"})
		// content = append(content, []Column{Content: task.Description, Width: 50, Height: 10}, {Content: task.Created.Format("Monday"), Width: 30, Height: 10}, {Content: "8", 30, Height: 10}, {Content: "$25.00", Width: 30, Height: 10}, {Content: "$200", Width: 30, Height: 10}})
		prevTask = task

		if n == 1 {
			break
		}

		n++
	}

	// fmt.Println(allTasks)

	pdf.OutputFileAndClose("hello.pdf")
}
