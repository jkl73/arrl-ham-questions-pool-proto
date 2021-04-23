package hamquestions

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/jkl73/arrl-ham-questions-pool-proto/proto"
)

const generalFigG7_1 = "2019-2023_general-G7-1"
const techFigT1 = "2018-2022_technician-T-1"
const techFigT2 = "2018-2022_technician-T-2"
const techFigT3 = "2018-2022_technician-T-3"
const extraFigE5_1 = "2020_Extra-E5-1"
const extraFigE6_1 = "2020_Extra-E6-1"
const extraFigE6_2 = "2020_Extra-E6-2"
const extraFigE6_3 = "2020_Extra-E6-3"
const extraFigE7_1 = "2020_Extra-E7-1"
const extraFigE7_2 = "2020_Extra-E7-2"
const extraFigE7_3 = "2020_Extra-E7-3"
const extraFigE9_1 = "2020_Extra-E9-1"
const extraFigE9_2 = "2020_Extra-E9-2"
const extraFigE9_3 = "2020_Extra-E9-3"

// parse a single question
//
// G0A01 (A)
// What is one way that RF energy can affect human body tissue?
// A. It heats body tissue
// B. It causes radiation poisoning
// C. It causes the blood count to reach a dangerously low level
// D. It cools body tissue
// ~~
func qparse(q string) *proto.Question {

	lines := strings.Split(q, "\n")

	// title
	sblmnt := lines[0][0:2]
	section := string(lines[0][2])
	seqnum, err := strconv.Atoi(lines[0][3:5])
	if err != nil {
		panic(err)
	}
	seqnumber := int32(seqnum)

	keyChoice := lines[0][7:8]

	chapter := ""
	if len(lines[0]) > 10 {
		chapter = lines[0][10:]
	}

	// stem
	stem := ""
	i := 1
	for i = 1; i < len(lines); i++ {
		if lines[i][0:2] == "A." {
			break
		} else {
			stem += lines[i]
			stem += " "
		}
	}
	stem = strings.TrimSpace(stem)
	fig := checkContaintedFigure(stem)

	// answers
	key := ""
	var distractor []string
	for _, t := range []string{"B.", "C.", "D.", "~~"} {
		ans := ""
		for ; i < len(lines); i++ {
			if lines[i][0:2] == t {
				break
			} else {
				ans += lines[i]
				ans += " "
			}
		}
		ans = strings.TrimSpace(ans)

		if string(ans[0]) == keyChoice {
			key = ans[3:]

		} else {
			distractor = append(distractor, ans[3:])

		}
	}

	res := proto.Question{
		Subelement:  sblmnt,
		Group:       section,
		Sequence:    seqnumber,
		Chapter:     chapter,
		Stem:        stem,
		Key:         key,
		Distractors: distractor,
		Figure:      fig,
	}

	return &res
}

func checkContaintedFigure(stem string) string {
	containFigure, err := regexp.MatchString(".*[Ff]igure [Gg]7-1.*", stem)
	if err != nil {
		panic(err)
	}
	if containFigure {
		return generalFigG7_1
	}
	containFigure, err = regexp.MatchString(".*[Ff]igure T1.*", stem)
	if err != nil {
		panic(err)
	}
	if containFigure {
		return generalFigG7_1
	}
	containFigure, err = regexp.MatchString(".*[Ff]igure T2.*", stem)
	if err != nil {
		panic(err)
	}
	if containFigure {
		return generalFigG7_1
	}
	containFigure, err = regexp.MatchString(".*[Ff]igure T3.*", stem)
	if err != nil {
		panic(err)
	}
	if containFigure {
		return generalFigG7_1
	}
	containFigure, err = regexp.MatchString(".*[Ff]igure T3.*", stem)
	if err != nil {
		panic(err)
	}
	if containFigure {
		return generalFigG7_1
	}

	containFigure, err = regexp.MatchString(".*[Ff]igure E5-1.*", stem)
	if err != nil {
		panic(err)
	}
	if containFigure {
		return extraFigE5_1
	}
	containFigure, err = regexp.MatchString(".*[Ff]igure E6-1.*", stem)
	if err != nil {
		panic(err)
	}
	if containFigure {
		return extraFigE6_1
	}
	containFigure, err = regexp.MatchString(".*[Ff]igure E6-2.*", stem)
	if err != nil {
		panic(err)
	}
	if containFigure {
		return extraFigE6_2
	}
	containFigure, err = regexp.MatchString(".*[Ff]igure E6-3.*", stem)
	if err != nil {
		panic(err)
	}
	if containFigure {
		return extraFigE6_3
	}
	containFigure, err = regexp.MatchString(".*[Ff]igure E7-1.*", stem)
	if err != nil {
		panic(err)
	}
	if containFigure {
		return extraFigE7_1
	}
	containFigure, err = regexp.MatchString(".*[Ff]igure E7-2.*", stem)
	if err != nil {
		panic(err)
	}
	if containFigure {
		return extraFigE7_2
	}
	containFigure, err = regexp.MatchString(".*[Ff]igure E7-3.*", stem)
	if err != nil {
		panic(err)
	}
	if containFigure {
		return extraFigE7_3
	}
	containFigure, err = regexp.MatchString(".*[Ff]igure E9-1.*", stem)
	if err != nil {
		panic(err)
	}
	if containFigure {
		return extraFigE9_1
	}
	containFigure, err = regexp.MatchString(".*[Ff]igure E9-2.*", stem)
	if err != nil {
		panic(err)
	}
	if containFigure {
		return extraFigE9_2
	}
	containFigure, err = regexp.MatchString(".*[Ff]igure E9-3.*", stem)
	if err != nil {
		panic(err)
	}
	if containFigure {
		return extraFigE9_3
	}
	return ""
}
