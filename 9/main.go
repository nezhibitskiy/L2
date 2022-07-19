package main

// Реализовать утилиту wget с возможностью скачивать сайты целиком.

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"path"
	"regexp"
	"strings"

	colly "github.com/gocolly/colly"
)

func getCleanArgs() (urls []string) {
	args := os.Args
	if len(args) < 2 {
		log.Fatal("usage: wget example.com example2.com ...")
	}

	return args[1:]
}

func CreateFolder(folderName string) {
	_, ok := os.Stat(folderName)
	if os.IsExist(ok) {
		return
	} else if os.IsNotExist(ok) {
		if ok = os.MkdirAll(folderName, os.ModePerm); ok != nil {
			log.Fatal(ok)
		}
	}
}

func main() {
	var (
		folder   = "download"
		mainlink *url.URL
		reg      *regexp.Regexp
		ok       error
		col      *colly.Collector
		ulSet    = make(map[string]struct{})
	)
	CreateFolder(folder)
	urls := getCleanArgs()
	for _, link := range urls {
		link = strings.TrimRight(link, "/")
		if mainlink, ok = url.ParseRequestURI(link); ok != nil {
			log.Fatal(ok)
		}
		if reg, ok = regexp.Compile("https?://([a-z0-9]+[.])*" + mainlink.Host); ok != nil {
			log.Fatal(ok)
		}
		CreateFolder(folder + "/" + mainlink.Host)
		col = colly.NewCollector(colly.URLFilters(reg))
		col.OnHTML("a[href]", func(el *colly.HTMLElement) {
			ul := el.Request.AbsoluteURL(el.Attr("href"))
			if _, isExist := ulSet[ul]; !isExist {
				ulSet[ul] = struct{}{}
				_ = col.Visit(ul)
			}
		})

		col.OnResponse(func(r *colly.Response) {
			pth := r.Request.URL.Path
			full := folder + "/" + mainlink.Hostname() + pth
			if _, ok := ulSet[full]; !ok {
				ulSet[full] = struct{}{}
			} else {
				return
			}
			if path.Ext(full) == "" {
				CreateFolder(full)
			} else {
				CreateFolder(full[:strings.LastIndexByte(full, '/')])
			}
			if path.Ext(pth) == "" {
				if full[len(full)-1] != '/' {
					full += "/"
				}
				full += "index.html"
				if _, ok := os.Create(full); ok != nil {
					fmt.Println("err filecreation", ok)
				}
			}
			fmt.Println("saved:", mainlink.Hostname()+pth)
			if ok = r.Save(full); ok != nil {
				log.Fatal(ok)
			}
		})

		if ok = col.Visit(mainlink.String()); ok != nil {
			log.Fatal("err: visit: " + ok.Error())
		}
		col.Wait()
	}
}
