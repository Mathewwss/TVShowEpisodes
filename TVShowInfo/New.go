package TVShowInfo

import (
	"github.com/Mathewwss/TVShowEpisodes/GETRequest"
	"github.com/Mathewwss/TVShowEpisodes/utils"
	"fmt"
	"errors"
	"regexp"
)

func New(show_id string) (TVShow, error) {
	// 1. Get source
	// 2. Find seasons
	// 3. Create object

	// 1
	url := fmt.Sprint("https://www.imdb.com/title/")
	url = url + show_id + "/episodes"
	html, err := GETRequest.HtmlSource(url)

	if err != nil {
		return TVShow{}, err
	}

	// 2
	tag := "^.*tab-season-entry.*$"
	seasons := utils.SliceFilter(html, func(str string) bool {
		if regexp.MustCompile(tag).MatchString(str) {
			return true
		}
		return false
	})

	if len(seasons) == 0 {
		return TVShow{}, errors.New("Not found seasons!")
	}

	// 3
	out := TVShow{}
	out.IMDBID = show_id
	out.Seasons = len(seasons)
	out.EpisodesID = make([][]string, len(seasons))
	out.EpisodesYear = make([][]int, len(seasons))
	err = out.GetTitle()

	if err != nil {
		return TVShow{}, err
	}

	return out, nil
}
