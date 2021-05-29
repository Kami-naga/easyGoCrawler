package zhenai

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("testPage.html")

	if err != nil {
		panic(err)
	}
	result := ParseCityList(contents)

	const resultSize = 470
	expectedUrls := []string{
		"http://localhost:8080/mock/www.zhenai.com/zhenghun/aba",
		"http://localhost:8080/mock/www.zhenai.com/zhenghun/akesu",
		"http://localhost:8080/mock/www.zhenai.com/zhenghun/alashanmeng",
	}
	expectedCities := []string{
		"City 阿坝", "City 阿克苏", "City 阿拉善盟",
	}
	if len(result.Requests) != resultSize {
		t.Errorf("request size should be: %d, but had %d", resultSize, len(result.Requests))
	}
	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected request of %d should be: %s, but had %s", i, url, result.Requests[i].Url)
		}
	}
	if len(result.Items) != resultSize {
		t.Errorf("item size should be: %d, but had %d", resultSize, len(result.Items))
	}
	for i, city := range expectedCities {
		if result.Items[i].(string) != city {
			t.Errorf("expected city of %d should be: %s, but had %s", i, city, result.Items[i].(string))
		}
	}
}
