// ----------------------------- Package ---------------------------- //

package TVShowInfo

// ------------------------------------------------------------------ //
// ----------------------------- Imports ---------------------------- //

import "github.com/Mathewwss/TVShowEpisodes/GETRequest"
import "fmt"
import "strings"
import "errors"

// ------------------------------------------------------------------ //

// ------------------------------ Types ----------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Variables --------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Functions --------------------------- //

func New (show_id string) (TVShow, error) {
	// Base url
	url := fmt.Sprint("https://www.imdb.com/title/")
	url = url + show_id + "/episodes"

	// Get html code
	html, err := GETRequest.HtmlSource(url)

	// Check errors
	if err != nil {
		// Stop
		return TVShow{}, err

	}

	// Html pattern
	tag := "data-testid=\"tab-season-entry\""

	// Count seasons
	seasons := strings.Count(html, tag)

	// Check value
	if seasons == 0 {
		// Create error
		msg := "[ERROR] -> '" + show_id + "' is probably not a TV"
		msg = msg + " " + "Show IMDB ID!"
		err = errors.New(msg)

		// Stop
		return TVShow{}, err

	}

	// Start struct
	out := TVShow{}
	out.IMDBID = show_id
	out.Seasons = seasons
	out.EpisodesIMDBID = map[int][]string{}
	out.EpisodesYear = map[int][]int{}

	// Get tv show title
	err = out.GetTitle()

	// Check errors
	if err != nil {
		// Stop
		return TVShow{}, err

	}

	// Finish
	return out, nil

}

// ------------------------------------------------------------------ //
