package Logger

import (
	"fmt"
	"os"
	"time"
)

// Info appends provided log line to a txt file
func Info(line string) {
	var time string = time.Now().Format(time.RFC850)
	var formattedLine string = "[" + time + "] " + line

	fmt.Println(formattedLine)
	appendToFile(formattedLine)
}

// appendToFile appends provided line to log file.
func appendToFile(line string) {
	var f *os.File = openFile("history.txt")

	// Print log line into the log file.
	_, err := fmt.Fprintln(f, line)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	// Finally close the File.
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

// openFile opens file or create new one in case it doesn't exists.
func openFile(path string) *os.File {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		f, err = os.Create(path)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
			return nil
		}
	}
	return f
}
