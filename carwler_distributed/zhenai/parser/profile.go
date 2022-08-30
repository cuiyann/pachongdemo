package parser

import (
	"log"
	"pachong/carwler_distributed/engine"
	"pachong/carwler_distributed/model"
	"regexp"
	"strconv"
)

var ageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>([\d]+)岁</div>`)
var genderRe = regexp.MustCompile(`<a href="http://www.zhenai.com/zhenghun/[a-zA-Z0-9]+/([^>]+)/">[^<]+</a>`)
var heightRe = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">([\d]+)cm</div>`)
var weightRe = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">([\d]+)kg</div>`)
var incomeRe = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">月收入:([^<]+)</div>`)
var marriageRe = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">离异</div>`)
var occupationRe = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn purple">工作地:([^<]+)</div>`)
var hukouRe = regexp.MustCompile(`<div data-v-8b1eac0c="" class="m-btn pink">籍贯:([^<]+)</div>`)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	var profile model.Profile

	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err != nil {
		profile.Age = age
	}
	profile.Age = age
	profile.Name = name
	s2 := extractString(contents, genderRe)
	log.Println(s2)
	profile.Gender = s2
	s3 := extractString(contents, heightRe)
	atoi, _ := strconv.Atoi(s3)
	profile.Height = atoi
	s4 := extractString(contents, weightRe)
	i, _ := strconv.Atoi(s4)
	log.Println(s4)
	profile.Weight = i
	s5 := extractString(contents, incomeRe)
	profile.Income = s5
	s6 := extractString(contents, marriageRe)
	profile.Marriage = s6
	s7 := extractString(contents, occupationRe)
	profile.Occupation = s7
	s8 := extractString(contents, hukouRe)
	profile.Occupation = s8
	result := engine.ParseResult{
		Item: []interface{}{profile},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	submatch := re.FindSubmatch(contents)
	log.Println(string(submatch[1]))
	if len(submatch) >= 2 {
		return string(submatch[1])
	} else {
		return ""
	}
}
