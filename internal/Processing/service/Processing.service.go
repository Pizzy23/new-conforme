package Processing

import (
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

type Result struct {
	TotalWords    int
	NotFoundCount int
	TotalConfirm  int
	WordsFound    []string
	NotFoundWords []string
}

func NonConformity(data map[string]interface{}, words []string, typePdf string) Result {
	search := data["search"].([]string)
	notFoundWords := make([]string, 0)
	wordsFound := make([]string, 0)
	var wg sync.WaitGroup

	resultChan := make(chan string)
	notFoundChan := make(chan string)

	processWord := func(wordItem string) {
		defer wg.Done()
		word := strings.ToLower(wordItem)
		found := false
		for _, s := range search {
			if strings.Contains(strings.ToLower(s), word) {
				found = true
				break
			}
		}
		if found {
			resultChan <- word
		} else {
			notFoundChan <- word
		}
	}
	for _, wordItem := range words {
		wg.Add(1)
		go processWord(wordItem)
	}

	go func() {
		wg.Wait()
		close(resultChan)
		close(notFoundChan)
	}()

	for word := range resultChan {
		wordsFound = append(wordsFound, word)
	}

	for word := range notFoundChan {
		notFoundWords = append(notFoundWords, word)
	}

	totalWords := len(words)
	notFoundCount := len(notFoundWords)
	switch typePdf {
	case "Sector":
		return Result{
			WordsFound:    wordsFound,
			NotFoundWords: notFoundWords,
		}
	case "Word":
		return Result{
			TotalWords:    totalWords,
			NotFoundCount: notFoundCount,
			TotalConfirm:  totalWords - notFoundCount,
			WordsFound:    wordsFound,
			NotFoundWords: notFoundWords,
		}
	default:
		return Result{}
	}
}

func GetPDF(c *gin.Context, id string) {
	_, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.Set("Response", "Invalid Idpdf header")
		c.Status(http.StatusBadRequest)
		return
	}
}
