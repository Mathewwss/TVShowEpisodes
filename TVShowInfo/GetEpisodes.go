package TVShowInfo

import (
	"github.com/Mathewwss/TVShowEpisodes/GETRequest"
	"fmt"
)

func (s *TVShow) GetEpisodes(show_id string, season int, ch chan <- error) {
	// 1. Get source
	// 2. Get episodes id
	// 3. Get episodes year

	// 1
	url := fmt.Sprint("https://www.imdb.com/title/")
	url = url + show_id + "/episodes/?season=" + fmt.Sprint(season)
	html, err := GETRequest.HtmlSource(url)

	if err != nil {
		ch <- err
		return
	}

	// 2
	err = s.getEpisodesID(&html, season)

	if err != nil {
		ch <- err
		return
	}

	// 3
	err = s.getEpisodesYear(&html, season)

	if err != nil {
		ch <- err
		return
	}

	ch <- nil
}
