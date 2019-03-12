//+build ignore
package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strings"

	"github.com/wzshiming/goquery"
	"github.com/wzshiming/requests"
	ffmt "gopkg.in/ffmt.v1"
)

// http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/

var ua = `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.119 Safari/537.36`
var host = `http://www.stats.gov.cn/tjsj/tjbz/tjyqhdmhcxhfdm/`
var cli = requests.NewClient().
	SetCache(requests.FileCacheDir("./tmp/")).
	SetLogLevel(requests.LogInfo).
	NewRequest().
	// SetTimeout(time.Second).
	SetUserAgent(ua)

func main() {
	a, err := start()
	if err != nil {
		ffmt.Mark(err)
		return
	}
	data, err := json.Marshal(a)
	if err != nil {
		ffmt.Mark(err)
		return
	}
	ioutil.WriteFile("pcctv.json", data, 0666)
}

func start() ([]*Area, error) {
	index, err := getIndex(host)
	if err != nil {
		return nil, err
	}
	areas, err := getProvincetr(index)
	if err != nil {
		return nil, err
	}
	for _, area := range areas {
		areas, err := getCity(area.SourceURL)
		if err != nil {
			ffmt.Mark(err)
		}
		area.Children = areas

		for _, area := range areas {
			areas, err := getCounty(area.SourceURL)
			if err != nil {
				ffmt.Mark(err)
			}
			area.Children = areas

			for _, area := range areas {
				areas, err := getTown(area.SourceURL)
				if err != nil {
					ffmt.Mark(err)
				}
				area.Children = areas

				for _, area := range areas {
					// time.Sleep(time.Second / 10)
					areas, err := getVillage(area.SourceURL)
					if err != nil {
						ffmt.Mark(err)
					}
					area.Children = areas

				}
			}
		}
	}
	return areas, nil
}

func getVillage(url string) (areas []*Area, err error) {
	cli := cli.Clone()
	resp, err := cli.
		SetMethod(requests.MethodGet).
		SetURLByStr(url).
		Do()
	if err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromReader(resp.RawBody())
	if err != nil {
		return nil, err
	}

	a := &Area{}
	doc.Find("table.villagetable tbody tr.villagetr td").EachWithBreak(func(i int, v *goquery.Selection) bool {
		switch i % 3 {
		case 0:
			a = &Area{
				AreaID: v.Text(),
			}
		case 2:
			a.Name = v.Text()
			areas = append(areas, a)
		}
		return true
	})
	return areas, nil
}

func getTown(url string) (areas []*Area, err error) {
	cli := cli.Clone()
	resp, err := cli.
		SetMethod(requests.MethodGet).
		SetURLByStr(url).
		Do()
	if err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromReader(resp.RawBody())
	if err != nil {
		return nil, err
	}

	doc.Find("table.towntable tbody tr td a").EachWithBreak(func(i int, v *goquery.Selection) bool {
		if i%2 == 0 {
			return true
		}
		href := v.AttrOr("href", "")
		text := v.Text()
		id := strings.TrimSuffix(href, ".html")
		id = id[strings.Index(id, "/")+1:]
		areas = append(areas, &Area{
			Name:      text,
			AreaID:    id,
			SourceURL: cli.GetURL(href).String(),
		})
		return true
	})
	return areas, nil
}

func getCounty(url string) (areas []*Area, err error) {
	cli := cli.Clone()
	resp, err := cli.
		SetMethod(requests.MethodGet).
		SetURLByStr(url).
		Do()
	if err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromReader(resp.RawBody())
	if err != nil {
		return nil, err
	}

	doc.Find("table.countytable tbody tr td a").EachWithBreak(func(i int, v *goquery.Selection) bool {
		if i%2 == 0 {
			return true
		}
		href := v.AttrOr("href", "")
		text := v.Text()
		id := strings.TrimSuffix(href, ".html")
		id = id[strings.Index(id, "/")+1:]
		areas = append(areas, &Area{
			Name:      text,
			AreaID:    id,
			SourceURL: cli.GetURL(href).String(),
		})
		return true
	})
	return areas, nil
}

func getCity(url string) (areas []*Area, err error) {
	cli := cli.Clone()
	resp, err := cli.
		SetMethod(requests.MethodGet).
		SetURLByStr(url).
		Do()
	if err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromReader(resp.RawBody())
	if err != nil {
		return nil, err
	}

	doc.Find("table.citytable tbody tr td a").EachWithBreak(func(i int, v *goquery.Selection) bool {
		if i%2 == 0 {
			return true
		}
		href := v.AttrOr("href", "")
		text := v.Text()
		id := strings.TrimSuffix(href, ".html")
		id = id[strings.Index(id, "/")+1:]
		areas = append(areas, &Area{
			Name:      text,
			AreaID:    id,
			SourceURL: cli.GetURL(href).String(),
		})
		return true
	})
	return areas, nil
}

func getProvincetr(url string) (areas []*Area, err error) {
	cli := cli.Clone()
	resp, err := cli.
		SetMethod(requests.MethodGet).
		SetURLByStr(url).
		Do()
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(resp.RawBody())
	if err != nil {
		return nil, err
	}

	doc.Find(".provincetr td a").EachWithBreak(func(i int, v *goquery.Selection) bool {
		href := v.AttrOr("href", "")
		text := v.Text()
		areas = append(areas, &Area{
			Name:      text,
			AreaID:    strings.TrimSuffix(href, ".html"),
			SourceURL: cli.GetURL(href).String(),
		})
		return true
	})
	return areas, nil
}

func getIndex(url string) (string, error) {
	cli := cli.Clone()
	resp, err := cli.
		SetMethod(requests.MethodGet).
		SetURLByStr(url).
		Do()
	if err != nil {
		return "", err
	}
	doc, err := goquery.NewDocumentFromReader(resp.RawBody())
	if err != nil {
		return "", err
	}
	href := ""
	doc.Find(".center_list ul li a").EachWithBreak(func(i int, v *goquery.Selection) bool {
		href = v.AttrOr("href", "")
		return false
	})
	if href == "" {
		return "", errors.New("Not found")
	}

	return cli.GetURL(href).String(), nil
}

// Area
type Area struct {
	Name      string  `json:"name"`
	AreaID    string  `json:"area_id"`
	Children  []*Area `json:"children,omitempty"`
	SourceURL string  `json:"-"`
}
