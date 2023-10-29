# TVShowEpisodes

The project aims to automate the collection of data from episodes of
TV Shows available on IMDb, allowing users to access information from
all episodes of a TV Show.

# Tools

## [Golang (Go)](https://go.dev/)

Build simple, secure, scalable systems with Go.

- An open-source programming language supported by Google.
- Easy to learn and great for teams.
- Built-in concurrency and a robust standard library.
- Large ecosystem of partners, communities, and tools.

[Learn More](https://go.dev/)

## Web scraping

Web scraping, web harvesting, or web data extraction is data scraping
used for extracting data from websites.

[Learn More](https://en.wikipedia.org/wiki/Web_scraping)

## Goroutines

Using goroutines and channels to reduce wait time.

### Test: TV Show (22 seasons).

|Goroutines|Time|
|:--------:|:--:|
|🟢|~ 5 seconds|
|🔴|~ 25 seconds|

# Features
- Obtain the TV Show name according to the country of origin of the
web request.
- Retrieve all available episodes.
- Obtain the IMDB code for each episode of the TV Show.
- Retrieve the release year of the episodes.

# Usage

1. Clone the project.

2. Build

Execute within the project folder:

```bash
	go build -o TVShowEpisodes.bin TVShowEpisodes.go
```

3. Running

You need to pass 1 parameter, which is the IMDB code of the TV Show.

```
	# Using go run
	# Chernobyl IMDBID = tt9166672
	go run TVShowEpisodes.go tt9166672

	# Using binary
	# Chernobyl IMDBID = tt9166672
	./TVShowEpisodes.bin tt9166672
```

# Example

![](https://github.com/Mathewwss/TVShowEpisodes/blob/master/example.gif)
