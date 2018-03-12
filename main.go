package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/jung-kurt/gofpdf"
	"google.golang.org/api/iterator"
)

var _ = datastore.ErrNoSuchEntity
var _ = context.Background()
var _ = gofpdf.CnProtectAnnotForms

type Task struct {
	Description string
	Created     time.Time
}

var dsClient *datastore.Client
var clientCtx context.Context
var pdf *gofpdf.Fpdf

func main() {
	inputDate := flag.String("since", time.Now().Format("2006-01-02"), "Generate pdf data since the input date.")
	flag.Parse()

	clientCtx = context.Background()

	var err error
	dsClient, err = datastore.NewClient(clientCtx, "worklog-191500")
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", root)
	http.HandleFunc("/create", actionCreate)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))

	//dateRange, err := time.Parse("2006-01-02", *inputDate)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//_ = dateRange

	//generatePdf(dateRange)
}

func table2(x, y float64, cols [][]Column) {

}

func generatePdf(inputDate time.Time) {
	pdf = gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(100, 10, "Anuar Kilgore")
	pdf.Ln(10)

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
		Filter("Created >", inputDate).
		Order("Created")

	var prevTask Task
	nextH := 0.0
	var allTasks []Task
	var table *Table

	for t := dsClient.Run(clientCtx, q); ; {
		var task Task
		_, err := t.Next(&task)
		allTasks = append(allTasks, task)
		if err == iterator.Done {
			// pdf.Ln(10)
			// pdf.CellFormat(10, 220, "Total Hours:		80", "", 0, "L", false, 0, "")
			// pdf.Ln(5)
			// pdf.CellFormat(10, 220, "Hourly Rate:		$25", "", 0, "L", false, 0, "")
			// pdf.Ln(5)
			// pdf.CellFormat(10, 220, "Total Amount:	$2000.00", "", 0, "L", false, 0, "")
			break
		}
		if err != nil {
			log.Fatal(err)
		}
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
			// 	// table(10, 70+nextH, 38, 10, content)
			// 	nextH = 10 * float64(len(content)+1)
			// 	content = [][]string{{"Description", "Day", "Hours", "Rate", "Amount"}}
			// 	table := NewTable(10, nextH, columns)
			// 	pdf.AddPage()
			// 	pdf.CellFormat(10, 0, "New week 1 - 2 Jan", "", 0, "L", false, 0, "")
		}

		fmt.Println(newWeek)

		table.AddRow([]string{task.Description, task.Created.Format("Monday"), "8", "$25.00", "$200.00"})
		// content = append(content, []string{task.Description, task.Created.Format("Monday"), "8", "$25.00", "$200"})
		// content = append(content, []Column{Content: task.Description, Width: 50, Height: 10}, {Content: task.Created.Format("Monday"), Width: 30, Height: 10}, {Content: "8", 30, Height: 10}, {Content: "$25.00", Width: 30, Height: 10}, {Content: "$200", Width: 30, Height: 10}})
		prevTask = task
	}

	fmt.Println(allTasks)

	pdf.OutputFileAndClose("hello.pdf")
}

// func table(x, y, cellW, cellH float64, content [][]string) {
// 	relH, rowCount := 0.0, 0.0
// 	for i, row := range content {
// 		for j, col := range row {
// 			// relW := cellW * float64(j-1)
// 			// if j == 0 {
// 			// 	relW = 0.0
// 			// }

// 			bonus, bonus2 := 0.0, 0.0
// 			_, _ = bonus, bonus2
// 			if j == 0 {
// 				bonus = float64(cellW) * 1.2
// 			} else {
// 				bonus = float64(cellW) - float64(cellW)*1.5
// 				bonus2 = float64(cellW) * 1.2
// 			}
// 			// relW := bonus2 * float64(j)
// 			relW := bonus * float64(j)
// 			// relW := 0.0

// 			// if i != 0 {
// 			// 	bonus = float64(cellW) * 0.8
// 			// }

// 			pdf.Rect(x+relW+bonus2, y+relH, cellW+bonus, cellH, "D")
// 			fmt.Println(j)
// 			fmt.Println(x)
// 			fmt.Println(relW)
// 			fmt.Println(bonus)
// 			fmt.Println(bonus2)
// 			fmt.Println(x + relW + bonus2)
// 			fmt.Println("\n")

// 			// Add Bold to header's font.
// 			if rowCount == 0 {
// 				pdf.SetFont("Arial", "B", 12)
// 			} else {
// 				pdf.SetFont("Arial", "", 12)
// 			}

// 			pdf.CellFormat(cellW+bonus, cellH, col, "", 0, "L", false, 0, "")

// 			if j == 2 {
// 				break
// 			}
// 		}

// 		_ = i

// 		if i == 0 {
// 			break
// 		}

// 		rowCount++
// 		relH = rowCount * cellH
// 		pdf.Ln(cellH)
// 	}

// 	//pdf.Ln(cellH*iterated)
// 	fmt.Println("tbl called")
// }

func actionCreate(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024)
	task := newTask()
	val := r.PostFormValue("desc")
	if len(val) > 1 {
		task.Description = r.PostFormValue("desc")
		task.Save()
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", "Hello world")
}

func newTask() *Task {
	return &Task{
		Created: time.Now(),
	}
}

func (e *Task) Save() {
	k := datastore.IncompleteKey("Task", nil)

	if _, err := dsClient.Put(clientCtx, k, e); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("New Task with description %q\n", e.Description)
}
