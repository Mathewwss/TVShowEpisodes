package GETRequest

import (
	"net/http"
	"io"
	"fmt"
	"github.com/Mathewwss/TVShowEpisodes/utils"
)

func HtmlSource(url string) ([]string, error) {
	// 1. Query settings
	// 2. Get source
	// 3. Show results

	// 1
	// Fix: error 404
	user_agent := "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/118.0.0.0 Safari/537.36"

	// 2
	client := &http.Client{}
	req, err := http.NewRequest(fmt.Sprint("GET"), url, nil)

	if err != nil {
		return []string{}, err
	}

	req.Header.Set("User-Agent", user_agent)
	resp, err := client.Do(req)

	if err != nil {
		return []string{}, err
	}

	defer resp.Body.Close()

	// 3
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return []string{}, err

	}

	return utils.HtmlTrim(string(body)), nil
}
