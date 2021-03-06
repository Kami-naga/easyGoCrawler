package zhenai

import (
	"goCrawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://localhost:8080/mock/www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

//get infos(city name, city url) from the page
func ParseCityList(
	contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Requests = append(
			result.Requests, engine.Request{
				Url:       string(m[1]),
				ParseFunc: ParseCity,
			})
		result.Items = append(
			result.Items, "City "+string(m[2]))
	}

	return result
}
