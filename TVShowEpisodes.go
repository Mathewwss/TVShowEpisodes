// ----------------------------- Package ---------------------------- //

package main

// ------------------------------------------------------------------ //

// ----------------------------- Imports ---------------------------- //

import "github.com/Mathewwss/TVShowEpisodes/TVShowInfo"
import "fmt"
import "os"

// ------------------------------------------------------------------ //

// ------------------------------ Types ----------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Variables --------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Functions --------------------------- //

func main () {
	// Check os arguments
	if len(os.Args) != 2 {
		// Show error
		msg := "[USAGE] -> " + os.Args[0] + " <IMDB ID (TV Show)>"
		fmt.Println(msg)

		// Stop
		return

	}

	// Generate struct
	tvshow, err := TVShowInfo.New(os.Args[1])

	// Check errors
	if err != nil {
		// Show error
		fmt.Println(err)

		// Stop
		return

	}

	// Create channel
	ch := make(chan error)

	// View seasons
	for a := 1; a <= tvshow.Seasons; a++ {
		// Start go routines
		go tvshow.GetEpisodes(tvshow.IMDBID, a, ch)

	}

	// View seasons
	for a := 1; a <= tvshow.Seasons; a++ {

		// Check channel values
		err := <- ch
		if err != nil {
			// Show error
			fmt.Println(err)

			// Stop
			return

		}
	}

	// View seasons
	for a := 1; a <= tvshow.Seasons; a++ {
		// Start variable
		separator := ""

		// Check value
		if a < 10 {
			// Update value
			separator = "# Season 0" + fmt.Sprint(a)

		} else {
			// Update value
			separator = "# Season " + fmt.Sprint(a)

		}

		// Show separator
		fmt.Println(separator)

		// View episodes
		for b := 0; b < len(tvshow.EpisodesIMDBID[a]); b++ {
			// Episode final name
			ep := tvshow.Title
			ep = ep + "."

			// Check value
			if a < 10 {
				// Update final name
				ep = ep + "s0" + fmt.Sprint(a)

			} else {
				// Update final name
				ep = ep + "s" + fmt.Sprint(a)

			}

			// Check value
			if b + 1 < 10 {
				// Update final name
				ep = ep + "e0" + fmt.Sprint(b + 1)
				ep = ep + "."

			} else {

				// Update final name
				ep = ep + "e" + fmt.Sprint(b + 1)
				ep = ep + "."

			}

			// Check size
			if b < len(tvshow.EpisodesYear[a]) {
				// Update final name
				ep = ep + fmt.Sprint(tvshow.EpisodesYear[a][b])
				ep = ep + "."

			}

			// Update final name
			ep = ep + tvshow.EpisodesIMDBID[a][b]

			// Show final name
			fmt.Println(ep)

		}

		// Check loop
		if a != tvshow.Seasons {
			// Skip line
			fmt.Println()

		}
	}
}

// ------------------------------------------------------------------ //
