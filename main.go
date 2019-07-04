package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

const (
	// File
	fp = "./grades_site.htm"

	// Selectors
	gradeSelector         = `span[title="7-trins-skalaen"]`
	pointSelector         = `td:nth-last-child(2)`
	validTableRowSelector = `tr:has(td:has(span[title="7-trins-skalaen"]))`

	// Regular expressions
	gradeRegExp = `(-3|00|02|4|7|10|12)`
	pointRegExp = `^[0-9]+(\.[0-9])?`
)

// Compile regexes at init time
var matchGrade = regexp.MustCompile(gradeRegExp)
var matchPoint = regexp.MustCompile(pointRegExp)

func main() {
	// Read file
	f, err := os.Open(fp)
	if err != nil {
		fmt.Println(err)
		return
	}

	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	f.Close()

	grades := []int{}
	points := []float64{}

	// Select all rows that have a "7-step scale" grade.
	doc.Find(validTableRowSelector).Each(func(i int, s *goquery.Selection) {
		// Find grades
		s.Find(gradeSelector).Each(func(i int, s *goquery.Selection) {
			text := matchGrade.FindString(s.Text())
			grade, err := strconv.Atoi(text)
			if err != nil {
				return
			}

			grades = append(grades, grade)
		})

		// Find points
		s.Find(pointSelector).Each(func(i int, s *goquery.Selection) {
			parsed := matchPoint.FindString(s.Text())
			point, err := strconv.ParseFloat(parsed, 64)
			if err != nil {
				return
			}

			points = append(points, point)
		})
	})

	// Convert grades to floats
	gf := make([]float64, len(grades))
	for i, g := range grades {
		gf[i] = float64(g)
	}

	// Calculate weighted average
	result := WeightedArithmeticMean(gf, points)

	fmt.Println(result)
}
