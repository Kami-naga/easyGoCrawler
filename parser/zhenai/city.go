package zhenai

import (
	"goCrawler/engine"
	"regexp"
)

const cityRe = `<a href="(http://localhost:8080/mock/album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

//get infos(city name, city url) from the page
func ParseCity(
	contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		name := string(m[2]) // key point! name should be copied! or we will get the name of the same person
		result.Requests = append(
			result.Requests, engine.Request{
				Url: string(m[1]),
				ParseFunc: func(
					c []byte) engine.ParseResult {
					return ParseProfile(c, name)
				},
			})
		result.Items = append(
			result.Items, "User "+string(m[2]))
	}

	return result
}
