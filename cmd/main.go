package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/lulumel0n/arrl-ham-questions-pool-proto/proto"
	pbtext "google.golang.org/protobuf/encoding/prototext"
	pb "google.golang.org/protobuf/proto"
)

const outputDir = "../out/"
const generalText = "2019-2023_general"

func main() {

	data, err := ioutil.ReadFile("../raw-questions/2019-2023_general.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	qp := createPool(string(data))

	out, err := pb.Marshal(qp)
	if err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile(outputDir+generalText+"_pb", out, 0644); err != nil {
		panic(err)
	}

	msout, err := pbtext.MarshalOptions{}.Marshal(qp)
	if err != nil {
		fmt.Println(err)
	}
	if err := ioutil.WriteFile(outputDir+generalText+"_pbtext", msout, 0644); err != nil {
		panic(err)
	}

	fmt.Printf("Generated output to: %s\n", outputDir)
}

func createPool(sourcePool string) *proto.QuestionPool {
	splitted := strings.Split(sourcePool, "\n")

	qpool := &proto.QuestionPool{}

	startr, _ := regexp.Compile("G[0-9][A-Z][0-9][0-9]\\s\\([A-D]\\)")
	endr, _ := regexp.Compile("~~")

	inQ := false

	question := ""
	for _, s := range splitted {

		matchStart := startr.MatchString(s)
		matchEnd := endr.MatchString(s)

		if inQ == true {
			question += s
			question += "\n"
		}

		if matchStart {
			inQ = true
			question += s
			question += "\n"
		} else if matchEnd {
			inQ = false
			qpool.Question = append(qpool.Question, qparse(question))
			question = ""
		}
	}

	return qpool
}
