package parser

import (
	"bytes"
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"project/railroad/engine"
	"project/railroad/model"
	"strings"
)

func goqueryCard(html string, selectorContent string) string {
	buffer := bytes.NewBuffer([]byte(html))
	doc, err := goquery.NewDocumentFromReader(buffer)
	if err != nil {
		panic(err)
	}
	dataIndexs := make(map[string]model.CardSrc)
	doc.Find(selectorContent).Find("li").Each(func(i int, s *goquery.Selection) {
		var data model.CardSrc
		dataIndex, _ := s.Attr("data-index")
		str := s.Text()
		str = strings.TrimSpace(str)
		data.SrcKey = str
		imgSrc, exists := s.Find(".obc-tmpl__switch-item img").Attr("src")
		if exists {
			data.SrcValue = imgSrc
		}
		value, ok := dataIndexs[dataIndex]
		if !ok {
			dataIndexs[dataIndex] = data
		} else {
			if value.SrcKey == "" {
				value.SrcKey = str
			}
			if value.SrcValue == "" {
				value.SrcValue = imgSrc
			}
			dataIndexs[dataIndex] = value
		}
	})
	for _, card := range dataIndexs {
		if card.SrcKey == "角色卡丨横板" {
			return card.SrcValue
		} else if card.SrcKey == "角色卡（横版）" {
			return card.SrcValue
		} else if card.SrcKey == "角色卡" {
			return card.SrcValue
		} else if card.SrcKey == "角色卡丨竖版" {
			return card.SrcValue
		} else if card.SrcKey == "角色立绘" {
			return card.SrcValue
		} else if card.SrcKey == "立绘" {
			return card.SrcValue
		}
	}
	return "暂时没有角色卡"
}

func goqueryCampSelector(html string, selectorContent string) string {
	buffer := bytes.NewBuffer([]byte(html))
	reader, err := goquery.NewDocumentFromReader(buffer)
	if err != nil {
		panic(err)
	}
	Information := make(map[string]string)
	reader.Find(selectorContent).Each(func(i int, selection *goquery.Selection) {
		key := selection.Find(".obc-tmp-character__key").Text()
		value := selection.Find(".obc-tmp-character__value").Text()
		Information[key] = value
	})
	if Information["阵营"] == "" {
		Information["阵营"] = "暂无阵营"
	}
	return Information["阵营"]
}

func goqueryPosterSelector(html string, selectorContent string) (string, string) {
	buffer := bytes.NewBuffer([]byte(html))
	reader, err := goquery.NewDocumentFromReader(buffer)
	if err != nil {
		panic(err)
	}
	divs := reader.Find(selectorContent)
	posters := make(map[int]string)
	divs.Each(func(i int, div *goquery.Selection) {
		styleAttr, _ := div.Attr("style")
		styles := parseStyles(styleAttr)
		for _, value := range styles {
			posters[i] = value[5 : len(value)-2]
		}
	})
	return posters[1], posters[0]
}

func parseStyles(styleAttr string) map[string]string {
	styles := make(map[string]string)
	for _, style := range strings.Split(styleAttr, ";") {
		if style == "" {
			continue
		}
		nameValue := strings.SplitN(style, ":", 2)
		if len(nameValue) != 2 {
			continue
		}
		name := strings.TrimSpace(nameValue[0])
		value := strings.TrimSpace(nameValue[1])
		styles[name] = value
	}
	return styles
}

func GetCampAndPoster(contents []byte, char *model.Character) engine.ParseResult {
	var data model.CharGet
	json.Unmarshal([]byte(contents), &data)
	html := data.Data.Content.Contents[0].Text
	selectorCamp := ".obc-tmp-character__property .obc-tmp-character__list .obc-tmp-character__wrap .obc-tmp-character__item"
	charCamp := goqueryCampSelector(html, selectorCamp)
	selectorPoster := ".obc-tmp-character__box"
	poster1, poster2 := goqueryPosterSelector(html, selectorPoster)
	charCard := "[data-part='painting']"
	card := goqueryCard(html, charCard)
	if card == "暂时没有角色卡" {
		l := len(data.Data.Content.Contents)
		html2 := data.Data.Content.Contents[l-1].Text
		card = goqueryCard(html2, charCard)
	}
	char.Camp = charCamp
	char.Poster1 = poster1
	char.Poster2 = poster2
	char.Card = card
	result := engine.ParseResult{Items: []interface{}{char}}
	return result
}
