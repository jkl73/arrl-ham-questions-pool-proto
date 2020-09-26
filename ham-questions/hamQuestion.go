package hamquestions

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	pb "google.golang.org/protobuf/proto"

	"github.com/lulumel0n/arrl-ham-questions-pool-proto/proto"
)

// HamQuestion contains the complete question set
type HamQuestion struct {
	Pool *proto.CompleteQuestionPool
}

// NewHamQuestion returns a struct with all questions in proto
func NewHamQuestion(cached string, rawQuestions string) (*HamQuestion, error) {
	qpb := &proto.CompleteQuestionPool{}

	// load from cache
	cachedpb, err := ioutil.ReadFile(cached)

	if err != nil {
		data, err := ioutil.ReadFile(rawQuestions)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		qpb = CreatePool(string(data))
	} else {
		if err = pb.Unmarshal(cachedpb, qpb); err != nil {
			fmt.Println("Fail to unmarshal cached proto")
			return nil, err
		}
	}

	res := &HamQuestion{qpb}
	return res, nil
}

// CreatePool creates a Ham quesitons pool from a formated txt questions pool
func CreatePool(sourcePool string) *proto.CompleteQuestionPool {
	lines := strings.Split(sourcePool, "\n")
	qpool := &proto.CompleteQuestionPool{}

	startr, _ := regexp.Compile("G[0-9][A-Z][0-9][0-9]\\s\\([A-D]\\)")
	endr, _ := regexp.Compile("~~")
	sublr, _ := regexp.Compile("SUBELEMENT G.*")
	inQ := false

	var subl *proto.Sublement
	var question string

	for _, s := range lines {
		if s == "" {
			continue
		}

		matchStart := startr.MatchString(s)
		matchEnd := endr.MatchString(s)
		matchSub := sublr.MatchString(s)

		if inQ == true {
			question += s
			question += "\n"
		} else {
			if matchSub {
				subl = &proto.Sublement{}
				subl.SublementId = s
				qpool.Subl = append(qpool.Subl, subl)
			}
		}

		if matchStart {
			inQ = true
			question += s
			question += "\n"
		} else if matchEnd {
			inQ = false
			q := qparse(question)

			if subl.GroupMap == nil {
				subl.GroupMap = make(map[string]*proto.QuestionList)
			}
			if subl.GroupMap[q.GetSection()] == nil {
				subl.GroupMap[q.GetSection()] = &proto.QuestionList{}
			}
			subl.GroupMap[q.GetSection()].Questions = append(subl.GroupMap[q.GetSection()].Questions, q)
			question = ""
		}
	}
	return qpool
}
