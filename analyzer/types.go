package analyzer

type FileAnalysis struct {
	FileName  string
	WordCount int
	CharCount int
	Error     string
}

type Results struct {
	Files      map[int]FileAnalysis
	TotalWords int
	TotalChars int
}

func NewResults() *Results {
	return &Results{
		Files: make(map[int]FileAnalysis),
	}
}
