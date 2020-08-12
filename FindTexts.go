package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func scan_recursive(dir_path string, ignore []string) ([]string, []string) {

	folders := []string{}
	files := []string{}

	// Scan
	filepath.Walk(dir_path, func(path string, f os.FileInfo, err error) error {

		_continue := false

		// Loop : Ignore Files & Folders
		for _, i := range ignore {

			// If ignored path
			if strings.Index(path, i) != -1 {

				// Continue
				_continue = true
			}
		}

		if _continue == false {

			f, err = os.Stat(path)

			// If no error
			if err != nil {
				log.Fatal(err)
			}

			// File & Folder Mode
			f_mode := f.Mode()

			// Is folder
			if f_mode.IsDir() {

				// Append to Folders Array
				folders = append(folders, path)

				// Is file
			} else if f_mode.IsRegular() {

				// Append to Files Array
				files = append(files, path)
			}
		}

		return nil
	})

	return folders, files
}

func main() {

	mydir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(mydir)

	folders, files := scan_recursive(mydir, []string{".git", "/.git", "/.git/", ".gitignore", ".DS_Store", ".idea", "/.idea/", "/.idea"})

	// Files
	for _, i := range files {
		fmt.Println(i)
	}

	// Folders
	for _, i := range folders {
		fmt.Println(i)
	}

	fmt.Println("next: \n")
	foundedTextCount := 0
	for i := 0; i <= (len(files) - 1); i++ {
		file, err := os.Open(files[i])
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			match, _ := regexp.MatchString(".text += +\".+\"( +)?", scanner.Text())
			if match {
				foundedTextCount++
				fmt.Println(scanner.Text())
			}
		}
		file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Founded Text " + strconv.Itoa(foundedTextCount) + " match")

	//fmt.Println(prettyCrop(string(content), 5))

}

func prettyCrop(in string, cropLength int) string {
	in2 := []rune(in)
	if len(in2) < cropLength {
		return in
	} else {
		in2 = in2[:cropLength]
		in = strings.TrimRightFunc(string(in2), func(r rune) bool {
			if r == ' ' {
				return true
			}
			return false
		})
		return in + "…"
	}
}
