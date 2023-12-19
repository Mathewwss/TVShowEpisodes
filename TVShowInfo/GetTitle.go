// ----------------------------- Package ---------------------------- //

package TVShowInfo

// ------------------------------------------------------------------ //

// ----------------------------- Imports ---------------------------- //

import "github.com/Mathewwss/TVShowEpisodes/GETRequest"
import "fmt"
import "strings"

// ------------------------------------------------------------------ //

// ------------------------------ Types ----------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Variables --------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Functions --------------------------- //

func (s *TVShow) GetTitle () (error) {
	// Base url
	url := fmt.Sprint("https://www.imdb.com/title/") + s.IMDBID

	// Get html code
	html, err := GETRequest.HtmlSource(url)

	// Check errors
	if err != nil {
		// Stop
		return err

	}

	// Html pattern
	pattern := "class=\"hero__primary-text\""

	// String to slice
	src := strings.Split(html, "\n")

	// View lines
	for a := 0; a < len(src); a++ {

		// Check size
		if len(src[a]) < len(pattern) {
			// Next
			continue

		}

		// String to slice
		line := strings.Split(src[a], " ")

		// Check size
		if len(line) != 3 {
			// Next
			continue

		}

		// Check value
		if line[1] != pattern {
			// Next
			continue

		}

		// Update struct
		s.Title = src[a + 1]

		// Stop loop
		break

	}

	// Finish
	return nil
}

// ------------------------------------------------------------------ //
