package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	wkhtmltopdf "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

type week string

type InvoiceData struct {
	Header      []string
	TasksByWeek map[week][]Task
}

type User struct {
	gorm.Model
	Email, Name, AuthKey string
}

type Task struct {
	gorm.Model
	Description, Day, Hours, Rate, Amount string
	UserID                                uint
}

func getInvoiceData() InvoiceData {
	header := []string{
		"Anuar Kilgore",
		"Col. Los Alamos",
		"San Pedro Sula, Cortez",
		"21101",
		"Honduras",
		"Phone: +504 99467213",
		"jaicof@gmail.com",
	}

	tasks := make(map[week][]Task, 0)
	tasks["25-29 December of 2017"] = []Task{
		{Description: "First commit", Day: "Monday", Hours: "8", Rate: "25.00", Amount: "200.00"},
		{Description: "Second commit", Day: "Monday", Hours: "8", Rate: "25.00", Amount: "200.00"},
	}

	invoiceData := InvoiceData{
		Header:      header,
		TasksByWeek: tasks,
	}

	return invoiceData
}

func generateInvoice() {
	tmpl := template.Must(template.ParseFiles("template.html"))
	file, err := os.Create("result.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(file, getInvoiceData())
	file.Close()

	// page := wkhtmltopdf.NewPage("https://godoc.org/github.com/SebastiaanKlippert/go-wkhtmltopdf")
	// page.FooterRight.Set("[page]")
	// page.FooterFontSize.Set(10)
	// page.Zoom.Set(95.50)
	// pdfg.AddPage(page)
	// pdfg.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(html)))
	// pdfg.AddPage(wkhtmltopdf.NewPageReader(file))

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	pdfg.Dpi.Set(300)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdfg.Grayscale.Set(true)
	page := wkhtmltopdf.NewPage("./result.html")
	page.FooterRight.Set("[page]")
	page.FooterFontSize.Set(10)
	pdfg.AddPage(page)

	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	err = pdfg.WriteFile("./invoice.pdf")
	if err != nil {
		log.Fatal(err)
	}
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

var db *gorm.DB

func main() {
	// flags
	runMigration := flag.Bool("migrate", false, "Run auto migration")
	flag.Parse()

	// Connect to postgres
	var err error
	db, err = gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=worklog password=asd123 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Run migrations with --migrate flag
	if *runMigration {
		Migrate()
		return
	}

	r := mux.NewRouter()
	r.HandleFunc("/", root)
	http.ListenAndServe("localhost:80", r)
}
