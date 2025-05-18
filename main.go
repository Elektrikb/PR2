package main

import (
	"file_analyzer/analyzer"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("использование: go run main.go <file1> <file2> ... <filen>")
		return
	}

	fileAnalyzer := analyzer.NewAnalyzer()
	results := fileAnalyzer.AnalyzeFiles(os.Args[1:])

	printResults(results)
}

func printResults(results *analyzer.Results) {
	fmt.Println(`-------------------------------------------------------------`)
	fmt.Println("Результаты анализа:")
	for i := 0; i < len(results.Files); i++ {
		file := results.Files[i]
		if file.Error != "" {
			fmt.Printf(
				"%d. %s: ОШИБКА: %s\n",
				i+1,
				file.FileName,
				file.Error,
			)
		} else {
			fmt.Printf(
				"%d. %s: %d слов, %d символов\n",
				i+1,
				file.FileName,
				file.WordCount,
				file.CharCount,
			)
		}
	}

	if len(results.Files) > 0 {
		fmt.Printf(
			"\nИтог: %d слов, %d символов",
			results.TotalWords,
			results.TotalChars,
		)
	}
}
