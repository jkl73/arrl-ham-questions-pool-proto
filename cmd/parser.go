package main

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/lulumel0n/arrl-ham-questions-pool-proto/proto"
)

const generalFigG7_1 = "2019-2023_general-G7-1.png"

func qparse(q string) *proto.Question {

	lines := strings.Split(q, "\n")

	// title
	sblmnt := lines[0][0:3]
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

	// figure handle (only 1 figure in the General exam)
	// Figure G7-1
	fig := ""
	containFigure, err := regexp.MatchString(".*[Ff]igure [Gg]7-1.*", stem)
	if err != nil {
		panic(err)
	}
	if containFigure {
		fig = generalFigG7_1
	}

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
			}
		}
		if string(ans[0]) == keyChoice {
			key = ans[3:]

		} else {
			distractor = append(distractor, ans[3:])

		}
	}

	res := proto.Question{
		Sublement:   sblmnt,
		Sequence:    seqnumber,
		Chapter:     chapter,
		Stem:        stem,
		Key:         key,
		Distractors: distractor,
		Figure:      fig,
	}

	return &res
}