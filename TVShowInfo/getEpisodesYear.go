// ----------------------------- Package ---------------------------- //

package TVShowInfo

// ------------------------------------------------------------------ //

// ----------------------------- Imports ---------------------------- //

import "fmt"
import "strings"

// ------------------------------------------------------------------ //

// ------------------------------ Types ----------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Variables --------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Functions --------------------------- //

func (s *TVShow) getEpisodesYear (source string, season int) error {
	// String to int
	html := strings.Split(source, "\n")

	// Html pattern
	pattern := "<span class=\"sc-f2169d65-10 bYaARM\">"

	// Start slice
	s.EpisodesYear[season] = []int{}

	// View lines
	for a := 0; a < len(html); a++ {

		// Check size
		if len(pattern) > len(html[a]) {
			// Next
			continue

		}

		// Check size
		if len(strings.Split(html[a], " ")) != 3 {
			// Next
			continue

		}

		// Check value
		if html[a] != pattern {
			// Next
			continue

		}

		// String to slice
		next := strings.Split(html[a + 1], " ")

		// Position
		index := 0

		// Check size
		if len(next) == 4 {
			// Update value
			index = 3

		} else {
			// Update value
			index = 0

		}

		// Start variable
		num := -1

		// String to int
		_, err := fmt.Sscanf(next[index], "%d", &num)

		// Check errors
		if err != nil {
			// Stop
			return err

		}

		// Append slice
		s.EpisodesYear[season] = append(s.EpisodesYear[season], num)

	}

	// Finish
	return nil
}

// ------------------------------------------------------------------ //
