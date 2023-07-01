package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type WaifuPics struct {
	Url string `json:"url"`
}

type NekosBest struct {
	Results []NBResults `json:"results"`
}

type NBResults struct {
	Url string `json:"url"`
}

func scrape(wCategory, res string) {
	url := res + wCategory
	apiResp, eGet := http.Get(url)
	if eGet != nil {
		log.Fatal(eGet)
	}
	defer apiResp.Body.Close()

	picUrl := ""

	switch res {
	case "https://api.waifu.pics/":
		pic := WaifuPics{}
		eJSON := json.NewDecoder(apiResp.Body).Decode(&pic)
		if eJSON != nil {
			log.Fatal(eJSON)
		}
		picUrl = pic.Url
	case "https://nekos.best/api/v2/":
		pic := NekosBest{}
		eJSON := json.NewDecoder(apiResp.Body).Decode(&pic)
		if eJSON != nil {
			log.Fatal(eJSON)
		}
		picUrl = pic.Results[0].Url
	case "https://api.catboys.com/":
		pic := WaifuPics{} // same layout
		eJSON := json.NewDecoder(apiResp.Body).Decode(&pic)
		if eJSON != nil {
			log.Fatal(eJSON)
		}
		picUrl = pic.Url
	default:
		picUrl = ""
	}

	urlArr := strings.Split(picUrl, "/")
	fName := urlArr[len(urlArr)-1]
	if strings.HasPrefix(wCategory, "sfw/") {
		wCategory = wCategory[4:]
	}
	dirPath := "./pics/" + wCategory + "/"

	ePath := os.MkdirAll(dirPath, os.ModePerm)
	if ePath != nil {
		log.Fatal(ePath)
	}

	if _, err := os.Stat(dirPath + fName); err == nil {
		return
	}

	picResp, eGet := http.Get(picUrl)
	if eGet != nil {
		log.Fatal(eGet)
	}
	defer picResp.Body.Close()

	file, eFile := os.Create(dirPath + fName)
	if eFile != nil {
		log.Fatal(eFile)
	}
	defer file.Close()

	_, eWrite := io.Copy(file, picResp.Body)
	if eWrite != nil {
		log.Fatal(eWrite)
	}
}

func wrap(resource string, rd, cd map[string]string) {
	num := 100

	if resource == "catboys" {
		num = 500
		for i := 0; i < num; i++ {
			scrape("img", rd[resource])
		}
		fmt.Println(resource, "img", "scraped")
		return
	}

	if resource == "waifuim" {
		num = 1000
		for i := 0; i < num; i++ {
			scrape("", rd[resource]) // <----------------------------------- no category, TBD
		}
		fmt.Println(resource, "img", "scraped")
		return
	}

	file, err := os.Open(cd[resource])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var eParse error
		waifuCategory := ""

		switch resource {
		case "waifupics":
			waifuCategory = "sfw/" + scanner.Text()
		case "nekosbest":
			arr := strings.Split(scanner.Text(), ",")
			waifuCategory = arr[0]
			num, eParse = strconv.Atoi(arr[1])
			if eParse != nil {
				log.Fatal(eParse)
			}
		default:
			waifuCategory = "neko"
		}

		for i := 0; i < num; i++ {
			scrape(waifuCategory, rd[resource])
		}
		fmt.Println(resource, waifuCategory, "scraped")
	}
}

func main() {
	resourceDict := map[string]string{
		"nekosbest": "https://nekos.best/api/v2/",
		"waifupics": "https://api.waifu.pics/",
		"catboys":   "https://api.catboys.com/",
		"waifuim":   "https://api.waifu.im/search/?excluded_tags=oppai&is_nsfw=false/", // exclude explicit content
	}

	catDict := map[string]string{
		"nekosbest": "./catNekosBest.txt",
		"waifupics": "./catWaifuPics.txt",
	}

	wrap("catboys", resourceDict, catDict)
}
