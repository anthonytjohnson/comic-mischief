package main

import "fmt"

func main() {
	var publisher, writer, artist, title string
	var year, pageNumber int
	var grade float32
	title = "Mr GoToSleep"
	writer = "Tracey Hatchet"
	artist = "Jewel Tampson"
	publisher = "Dizzy Books Publishing Inc."
	year = 1997
	pageNumber = 14
	grade = 6.5
	fmt.Println(title, "published by", publisher, "written by", writer, "artist:", artist, "Year:", year, "pages:", pageNumber, "grade:", grade)
}
