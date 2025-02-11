package TVShowInfo

import (
	"github.com/Mathewwss/TVShowEpisodes/utils"
	"fmt"
	"errors"
	"regexp"
)

func (s *TVShow) getEpisodesID(html *[]string, season int) error {
	// 1. Find pattern
	// 2. Get ids
	// 3. Update values

	// 1
	pattern := "^.*/tt.*ipc-title-link-wrapper.*$"
	tags := utils.SliceFilter((*html), func(str string) bool {

		if regexp.MustCompile(pattern).MatchString(str) {
			return true
		}

		return false
	})

	if len(tags) == 0 {
		return errors.New(fmt.Sprint("Not found episodes!"))
	}

	// 2.
	codes := utils.SliceMap(tags, func(str string) string {
		re := regexp.MustCompile("/.*$")
		return re.ReplaceAllString(str[16:], "")
	})

	// 3
	s.EpisodesID[season - 1] = codes
	return nil
}
