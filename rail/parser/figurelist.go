package parser

import (
	"encoding/json"
	"project/railroad/engine"
	"project/railroad/model"
	"strconv"
	"strings"
)

const url = "https://api-static.mihoyo.com/common/blackboard/sr_wiki/v1/content/info?app_sn=sr_wiki&content_id="

func ChangeToChar(ch model.ChGet) model.Character {
	var test model.Character
	test.Cid = ch.ContentID
	test.Name = ch.Title
	//test.icon = ch.Icon
	var ext model.ExtGet
	json.Unmarshal([]byte(ch.Ext), &ext)
	n := len(ext.C18.Filter.Text)
	split := strings.Split(ext.C18.Filter.Text[1:n-1], ",")
	extmap := make(map[string]string)
	for _, s := range split {
		m := len(s)
		th := strings.Split(s[1:m-1], "/")
		extmap[th[0]] = th[1]
	}
	test.Attribute = extmap["属性"]
	test.Fate = extmap["命途"]
	test.Star = extmap["星级"]
	return test
}

func CharNumber(contents []byte) int {
	var data model.UrlGet
	json.Unmarshal([]byte(contents), &data)
	lens := len(data.Data.List[0].Children[0].List)
	return lens
}

func ParseFigureList(contents []byte) engine.ParseResult {
	var data model.UrlGet
	json.Unmarshal([]byte(contents), &data)
	lens := len(data.Data.List[0].Children[0].List)
	characters := make([]model.Character, lens)
	result := engine.ParseResult{}
	for i, ch := range data.Data.List[0].Children[0].List {
		characters[i] = ChangeToChar(ch)
		char := characters[i]
		result.Requests = append(result.Requests, engine.Request{
			Url: url + strconv.Itoa(characters[i].Cid),
			ParserFunc: func(c []byte) engine.ParseResult {
				return GetCampAndPoster(c, &char)
			},
		})
	}
	return result
}
