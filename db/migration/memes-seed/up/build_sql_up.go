package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"time"
)

var urlRegex = regexp.MustCompile(`\((https?:[^)]*)\)`)

func main() {
	f, err := os.Open("./raw_data.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")
	fmt.Println(`-- Auto-generated by ./memes-seed/build_sql_up.go`)
	fmt.Println(`-- Using memes from https://old.reddit.com/r/image_linker_bot/comments/2znbrg/image_suggestion_thread_20/`)
	for _, line := range lines {
		lineParts := strings.Split(line, "|")

		memeNames := []string{}
		for _, name := range strings.Split(lineParts[1], ",") {
			memeNames = append(memeNames, strings.ToLower(strings.TrimSpace(name)))
		}

		memeURLs := []string{}
		for _, match := range urlRegex.FindAllStringSubmatch(lineParts[2], -1) {
			memeURLs = append(memeURLs, match[1])
		}

		sql := bytes.NewBufferString("")
		fmt.Fprint(sql, `
WITH meme AS (
	INSERT INTO memes DEFAULT VALUES RETURNING *
)`)

		for i, name := range memeNames {
			fmt.Fprintf(sql, `, meme_names_insert_%d as (
	INSERT INTO meme_names(name, timestamp, author, meme_id)
	VALUES('%s', '%s', 'seed', (SELECT id FROM meme))
	ON CONFLICT DO NOTHING
)`, i, name, time.Now().Format(time.RFC3339))
		}

		for i, url := range memeURLs {
			fmt.Fprintf(sql, `, meme_urls_insert_%d as (
	INSERT INTO meme_urls(url, timestamp, author, meme_id)
	VALUES('%s', '%s', 'seed', (SELECT id FROM meme))
	ON CONFLICT DO NOTHING
)`, i, url, time.Now().Format(time.RFC3339))
		}

		fmt.Fprint(sql, "\nSELECT 1;")

		fmt.Println(sql.String())
	}
}