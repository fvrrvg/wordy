package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func dbPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	if os.PathSeparator == '\\' {
		return filepath.Join(home, "AppData", "Local", "wordy")
	} else {
		return filepath.Join(home, ".config", "wordy")
	}
}

func createDir() {

	if _, err := os.Stat(dbPath()); os.IsNotExist(err) {
		os.Mkdir(dbPath(), 0755)
	}

}

func createFile() {
	if os.PathSeparator == '\\' {
		if _, err := os.Stat(dbPath() + "\\db.txt"); os.IsNotExist(err) {
			os.Create(dbPath() + "\\db.txt")
		}
	} else {
		if _, err := os.Stat(dbPath() + "/db.txt"); os.IsNotExist(err) {
			os.Create(dbPath() + "/db.txt")
		}
	}
}

func writeScore(score int) {
	createDir()
	createFile()
	if os.PathSeparator == '\\' {
		f, err := os.OpenFile(dbPath()+"\\db.txt", os.O_APPEND|os.O_WRONLY, 0644)

		if err != nil {
			panic(err)
		}
		defer f.Close()

		_, err = fmt.Fprintf(f, "%d\n", score)
		if err != nil {
			panic(err)
		}

	} else {
		f, err := os.OpenFile(dbPath()+"/db.txt", os.O_APPEND|os.O_WRONLY, 0644)

		if err != nil {
			panic(err)
		}
		defer f.Close()

		_, err = fmt.Fprintf(f, "%d\n", score)
		if err != nil {
			panic(err)
		}
	}

}

func readScore() int {
	createDir()
	createFile()

	if os.PathSeparator == '\\' {

		content, err := os.ReadFile(dbPath() + "\\db.txt")
		if err != nil {
			panic(err)
		}

		if len(content) == 0 {
			return 0
		}

		lines := strings.Split(string(content), "\n")
		var sum int

		for _, line := range lines {
			line = strings.TrimSpace(line)

			if line == "" {
				continue
			}

			score, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println("Error parsing line:", err)
				continue
			}
			sum += score
		}
		return sum

	} else {
		content, err := os.ReadFile(dbPath() + "/db.txt")
		if err != nil {
			panic(err)
		}

		if len(content) == 0 {
			return 0
		}

		lines := strings.Split(string(content), "\n")
		var sum int

		for _, line := range lines {
			line = strings.TrimSpace(line)
			if line == "" {
				continue
			}

			score, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println("Error parsing line:", err)
				continue
			}
			sum += score
		}
		return sum

	}

}
