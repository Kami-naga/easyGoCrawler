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
		result.Requests = append(
			result.Requests, engine.Request{
				Url:       string(m[1]),
				ParseFunc: engine.NilParser,
			})
		result.Items = append(
			result.Items, "User "+string(m[2]))
	}

	return result
}
