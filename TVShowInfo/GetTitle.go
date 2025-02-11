package TVShowInfo

import (
	"github.com/Mathewwss/TVShowEpisodes/GETRequest"
	"github.com/Mathewwss/TVShowEpisodes/utils"
	"fmt"
	"regexp"
	"errors"
)

func (s *TVShow) GetTitle() (error) {
	// 1. Get html source
	// 2. Find title
	// 3. Update value

	// 1
	url := fmt.Sprint("https://www.imdb.com/title/") + s.IMDBID
	html, err := GETRequest.HtmlSource(url)

	if err != nil {
		return err
	}

	// 2
	pattern := `<span class="hero__primary-text" data-testid="hero__primary-text">`
	res := utils.SliceFilter(html, func(str string) bool {
		if regexp.MustCompile(pattern).MatchString(str) {
			return true
		}
		return false
	})

	if len(res) == 0 {
		return errors.New("Not found TV Show!")
	}

	// 3
	s.Title = res[0][len(pattern):]
	return nil
}
