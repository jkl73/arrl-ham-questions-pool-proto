package hamquestions

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	pb "google.golang.org/protobuf/proto"

	"github.com/lulumel0n/arrl-ham-questions-pool-proto/proto"
)

// NewHamQuestionsAndTitles returns a struct with all questions in one proto and the titles in another
func NewHamQuestionsAndTitles(cached string, rawQuestions string) (*proto.CompleteQuestionPool, *proto.AllTitles, error) {
	qpb := &proto.CompleteQuestionPool{}
	titles := &proto.AllTitles{}

	// load from cache
	cachedpb, err := ioutil.ReadFile(cached)

	if err != nil {
		data, err := ioutil.ReadFile(rawQuestions)
		if err != nil {
			fmt.Println(err)
			return nil, nil, err
		}
		qpb, titles = CreatePool(string(data))
	} else {
		if err = pb.Unmarshal(cachedpb, qpb); err != nil {
			fmt.Println("Fail to unmarshal cached proto")
			return nil, nil, err
		}
	}

	return qpb, titles, nil
}

// CreatePool creates a Ham quesitons pool from a formated txt questions pool, and a titles only proto
func CreatePool(sourcePool string) (*proto.CompleteQuestionPool, *proto.AllTitles) {
	lines := strings.Split(sourcePool, "\n")
	qpool := &proto.CompleteQuestionPool{}
	qpool.SubelementMap = make(map[string]*proto.Subelement)

	alltitles := &proto.AllTitles{}

	startr, _ := regexp.Compile("G[0-9][A-Z][0-9][0-9]\\s\\([A-D]\\)")
	endr, _ := regexp.Compile("~~")
	subelementr, _ := regexp.Compile("SUBELEMENT G.*")
	groupr, _ := regexp.Compile("G[0-9][A-Z] –.*")
	inQ := false

	var subelement *proto.Subelement
	var group *proto.Group
	var question string

	var subelementTitle *proto.SubelementTitle
	var groupTitle *proto.GroupTitle

	for _, s := range lines {
		if s == "" {
			continue
		}

		matchStart := startr.MatchString(s)
		matchEnd := endr.MatchString(s)
		matchSubelement := subelementr.MatchString(s)
		matchGroup := groupr.MatchString(s)

		if inQ == true {
			question += s
			question += "\n"
		} else {
			if matchSubelement {
				subelement = &proto.Subelement{}
				subelement.Id = s[11:13]
				subelement.Title = s[18 : len(s)-32]
				subelement.GroupMap = make(map[string]*proto.Group)

				subelementTitle = &proto.SubelementTitle{}
				subelementTitle.Id = s[11:13]
				subelementTitle.Title = subelement.Title
				alltitles.Subelements = append(alltitles.Subelements, subelementTitle)

				qpool.SubelementMap[subelement.Id] = subelement
			} else if matchGroup {
				group = &proto.Group{}
				group.Id = string(s[2])
				group.Title = s[8:]

				subelement.GroupMap[group.Id] = group

				groupTitle = &proto.GroupTitle{}
				groupTitle.Id = group.Id
				groupTitle.Title = group.Title

				subelementTitle.Groups = append(subelementTitle.Groups, groupTitle)
			}
		}

		if matchStart {
			inQ = true
			question += s
			question += "\n"
		} else if matchEnd {
			q := qparse(question)

			if subelement.GroupMap == nil {
				subelement.GroupMap = make(map[string]*proto.Group)
			}
			group.Questions = append(group.Questions, q)

			// flush question buffer
			question = ""
			inQ = false
		}
	}
	return qpool, alltitles
}
