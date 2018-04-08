package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"time"
)

func createFolder() (todayDir string) {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	currentTime := time.Now()
	todayDir = fmt.Sprintf("%s/.journal/%s", usr.HomeDir, currentTime.Format("2006-01-02"))

	// Create todays directory
	err = os.MkdirAll(todayDir, 0755)
	if err != nil {
		log.Fatal(err)
	}

	return todayDir
}

// CreateJournal ...
func (j *Journal) CreateJournal(text string) (todayFile string) {
	todaysDirectory := createFolder()

	// Check amount of files add journal[num]++
	files, err := ioutil.ReadDir(todaysDirectory)
	if err != nil {
		log.Fatal(err)
	}

	todayFile = fmt.Sprintf("%s/%s-%d.txt", todaysDirectory, "journal", len(files)+1)

	// Create the content file for today
	f, err := os.Create(todayFile)
	defer f.Close()

	if err != nil {
		log.Fatal(err)
	}

	f.Write([]byte(text))

	return todayFile
}

// CreateJournalTextFile ..
func (j *Journal) CreateJournalTextFile(text string) {
	todaysTxtFile := j.CreateJournal(text)
	printJournalCreated([]string{todaysTxtFile})
}

// CreateJournalImageFile ..
func (j *Journal) CreateJournalImageFile(text string) {
	todaysTxtFile := j.CreateJournal(text)
	todaysImgFile := screenShot(todaysTxtFile)
	printJournalCreated([]string{todaysTxtFile, todaysImgFile})
}

func printJournalCreated(fileNames []string) {
	for _, fileName := range fileNames {
		fmt.Print(fmt.Sprintf("Created -> %s\n", fileName))
	}
}
