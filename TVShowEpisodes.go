package main

import (
	"github.com/Mathewwss/TVShowEpisodes/TVShowInfo"
	"fmt"
	"os"
)

func main () {
	// 1. Check args
	// 2. Create object
	// 3. Get info
	// 4. Show data

	// 1
	if len(os.Args) != 2 {
		msg := "[USAGE] -> " + os.Args[0] + " <IMDB ID (TV Show)>"
		fmt.Println(msg)

		return
	}

	// 2
	tvshow, err := TVShowInfo.New(os.Args[1])

	if err != nil {
		fmt.Println(err)
		return
	}

	ch := make(chan error)

	func(i, end int) {

		var fn func(int, int)
		fn = func(i, end int) {
			if i > end {
				return
			}

			go tvshow.GetEpisodes(tvshow.IMDBID, i, ch)

			fn(i + 1, end)
		}

		fn(i, end)

	}(1, tvshow.Seasons)

	err = func(i, end int) error {

		var fn func(int, int) error
		fn = func(i, end int) error {

			if i > end {
				return nil
			}

			ch_read := <- ch

			if ch_read != nil {
				return ch_read
			}

			return fn(i + 1, end)
		}

		return fn(i, end)

	}(1, tvshow.Seasons)

	if err != nil {
		fmt.Println(err)
		return
	}

	var season func(int)
	var eps func(int, int)

	season = func(i int) {

		if i == len(tvshow.EpisodesID) || i == len(tvshow.EpisodesYear) {
			return
		}

		if i > 1 {
			fmt.Printf("\n")
		}

		fmt.Printf("# Season %02d\n", i)

		eps(i - 1, 0)

		season(i + 1)
	}

	eps = func(i, j int) {

		if j >= len(tvshow.EpisodesYear[i]) || j >= len(tvshow.EpisodesID[i]) {
			return
		}

		fmt.Printf(
			// Title.sXeY.Year.ID
			"%v.s%02de%02d.%v.%v\n",
			tvshow.Title, i + 1, j + 1,
			tvshow.EpisodesYear[i][j],
			tvshow.EpisodesID[i][j],
		)

		eps(i, j + 1)
	}

	season(1)
}
