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

func (s *TVShow) getEpisodesIMDBID (source string, season int) error {
	// String to slice
	html := strings.Split(source, fmt.Sprint("\n"))

	// Html pattern
	pattern := "class=\"ipc-title-link-wrapper\" tabindex=\"0\">"

	// Start slice
	out := []string{}

	// View lines
	for a := 0; a < len(html); a++ {
		// Check size
		if len(pattern) > len(html[a]) {
			// Next
			continue

		}

		// String to slice
		line := strings.Split(html[a], " ")

		// Check size
		if len(line) != 4 {
			// Next
			continue

		}

		// Check value
		if line[2] + " " + line[3] != pattern  {
			// Next
			continue

		}

		// Append slice
		out = append(out, strings.Split(line[1], "/")[2])

	}

	// Update map
	s.EpisodesIMDBID[season] = out

	// Finish
	return nil
}

// ------------------------------------------------------------------ //
