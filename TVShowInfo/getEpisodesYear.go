package TVShowInfo

import (
	"github.com/Mathewwss/TVShowEpisodes/utils"
	"fmt"
	"strconv"
	"regexp"
	"errors"
)

func (s *TVShow) getEpisodesYear(html *[]string, season int) error {
	// 1. Find pattern
	// 2. String to int
	// 3. Update values

	// 1
	pattern := fmt.Sprint("^.*f2169d65-10 bYaARM.*$")
	tags := utils.SliceFilter((*html), func(str string) bool {
		if regexp.MustCompile(pattern).MatchString(str) {
			return true
		}
		return false
	})

	// 2
	years := make([]int, 0)
	var err error
	_ = utils.SliceMap(tags, func(str string) string {

		if err != nil {
			return ""
		}

		n, r := strconv.Atoi(str[len(str) - 4:])

		if r != nil {
			err = errors.New("Failed convert to int!")
		}

		years = append(years, n)

		return ""

	})

	if err != nil {
		return err
	}

	// 3
	s.EpisodesYear[season - 1] = years
	return nil
}
