package zhenai

import (
	"goCrawler/model"
	"io/ioutil"
	"testing"
)

func TestParseProfile(t *testing.T) {
	contents, err := ioutil.ReadFile("profile_test_page.html")
	if err != nil {
		panic(err)
	}

	result := ParseProfile(contents, "逆天飞翔莓哒")

	if len(result.Items) != 1 {
		t.Errorf("Items should contain 1 element, "+
			"but was %v", result.Items)
	}

	profile := result.Items[0].(model.Profile)
	profile.Name = "逆天飞翔莓哒"
	expected := model.Profile{
		Name:       "逆天飞翔莓哒",
		Gender:     "男",
		Age:        41,
		Height:     15,
		Weight:     235,
		Income:     "1-2000元",
		Marriage:   "未婚",
		Education:  "小学",
		Occupation: "销售",
		Hukou:      "苏州市",
		Xinzuo:     "双鱼座",
		House:      "租房",
		Car:        "有豪车",
	}

	if profile != expected {
		t.Errorf("expected %v, "+
			"but was %v", expected, profile)
	}
}
