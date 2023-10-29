// ----------------------------- Package ---------------------------- //

package TVShowInfo

// ------------------------------------------------------------------ //

// ----------------------------- Imports ---------------------------- //

import "github.com/Mathewwss/TVShowEpisodes/GETRequest"
import "fmt"

// ------------------------------------------------------------------ //

// ------------------------------ Types ----------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Variables --------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Functions --------------------------- //

func (s *TVShow) GetEpisodes (show_id string, season int, ch chan <- error) {
	// Base url
	url := fmt.Sprint("https://www.imdb.com/title/")
	url = url + show_id + "/episodes/?season=" + fmt.Sprint(season)

	// Get html source
	html, err := GETRequest.HtmlSource(url)

	// Check error
	if err != nil {
		// Update channel
		ch <- err

		// Stop
		return

	}

	// Get IMDB ID
	err = s.getEpisodesIMDBID(html, season)

	// Check errors
	if err != nil {
		// Update channel
		ch <- err

		// Stop
		return

	}

	// Get episodes year
	err = s.getEpisodesYear(html, season)

	// Check errors
	if err != nil {
		// Update channel
		ch <- err

		// Stop
		return

	}

	// Update channel
	ch <- nil
}

// ------------------------------------------------------------------ //
