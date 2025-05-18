package analyzer

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"sync"
)

type Analyzer struct {
	mu sync.Mutex
}

var re = regexp.MustCompile(`[\wа-яА-ЯёЁ]+`)

func NewAnalyzer() *Analyzer {
	return &Analyzer{}
}

func (a *Analyzer) AnalyzeFiles(filePaths []string) *Results {
	var wg sync.WaitGroup
	results := NewResults()

	for i, filePath := range filePaths {
		wg.Add(1)
		log.Printf("%d. %s начал обрабатываться\n", i+1, filePath)
		go func(index int, path string) {
			defer wg.Done()
			a.analyzeFiles(index, path, results)
			log.Printf("%d. %s закончил обрабатываться\n", i+1, filePath)
		}(i, filePath)
	}

	wg.Wait()
	return results
}

func (a *Analyzer) analyzeFiles(index int, path string, results *Results) {
	file, err := os.Open(path)
	if err != nil {
		a.mu.Lock()
		defer a.mu.Unlock()

		results.Files[index] = FileAnalysis{
			FileName: path,
			Error:    err.Error(),
		}
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var wordCount, charCount int

	for scanner.Scan() {
		line := scanner.Text()
		charCount += len([]rune(line))
		wordCount += len(re.FindAllString(line, -1))
	}

	if err := scanner.Err(); err != nil {
		a.mu.Lock()
		defer a.mu.Unlock()

		results.Files[index] = FileAnalysis{
			FileName: path,
			Error:    err.Error(),
		}
		return
	}

	a.mu.Lock()
	defer a.mu.Unlock()
	results.Files[index] = FileAnalysis{
		FileName:  path,
		WordCount: wordCount,
		CharCount: charCount,
	}
	results.TotalWords += wordCount
	results.TotalChars += charCount
}
