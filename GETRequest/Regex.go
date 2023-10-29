// ----------------------------- Package ---------------------------- //

package GETRequest

// ------------------------------------------------------------------ //

// ----------------------------- Imports ---------------------------- //

import "regexp"

// ------------------------------------------------------------------ //

// ------------------------------ Types ----------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Variables --------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Functions --------------------------- //

func configureTags (input string) string {
	// Regex patterns
	changes := [][]string{
		[]string{">", ">\n"},
		[]string{"<", "\n<"},
		[]string{"\t", " "},
		[]string{"  *", " "},
		[]string{"^ ", ""},
		[]string{" $", ""},
	}

	// View patterns
	for a := 0; a < len(changes); a++ {
		// Update string
		re := regexp.MustCompile(changes[a][0])
		input = re.ReplaceAllString(input, changes[a][1])

	}

	// Show new string
	return input
}

// ------------------------------------------------------------------ //
