// ----------------------------- Package ---------------------------- //

package GETRequest

// ------------------------------------------------------------------ //

// ----------------------------- Imports ---------------------------- //

import "net/http"
import "io"
import "fmt"

// ------------------------------------------------------------------ //

// ------------------------------ Types ----------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Variables --------------------------- //

// ------------------------------------------------------------------ //

// ---------------------------- Functions --------------------------- //

func HtmlSource (url string) (string, error) {
	// Custom user agent
	// Fix: error 404
	user_agent := fmt.Sprint("Mozilla/5.0 (X11; Linux x86_64)")
	user_agent = user_agent + " " + "AppleWebKit/537.36"
	user_agent = user_agent + " " + "(KHTML, like Gecko)"
	user_agent = user_agent + " " + "Chrome/118.0.0.0 Safari/537.36"

	// Start variables
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	// Check errors
	if err != nil {
		// Stop
		return "", err

	}

	// Set user agent
	// Fix: error 404
	req.Header.Set("User-Agent", user_agent)

	// Try request
	resp, err := client.Do(req)

	// Check errors
	if err != nil {
		// Stop
		return "", err

	}

	// Close body
	defer resp.Body.Close()

	// Get html source
	body, err := io.ReadAll(resp.Body)

	// Check errors
	if err != nil {
		// Stop
		return "", err

	}

	// Regex changes
	html := configureTags(string(body))

	// Show final string
	return html, nil
}

// ------------------------------------------------------------------ //
