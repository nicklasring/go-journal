package main

import (
	"flag"
)

// AddImage ..
func (j *Journal) AddImage(imageText *string) {
	j.CreateJournalImageFile(*imageText)
}

// AddText ..
func (j *Journal) AddText(text *string) {
	j.CreateJournalTextFile(*text)
}

func main() {
	journal := &Journal{}

	text := flag.String("t", "", "Text")
	imageText := flag.String("i", "", "Image")

	flag.Parse()

	var flagSet bool

	if *text != "" {
		journal.AddText(text)
		flagSet = true
	}

	if *imageText != "" {
		journal.AddImage(imageText)
		flagSet = true
	}

	if !flagSet {
		flag.Usage()
	}
}
