package main

import (
	"fmt"
	"io/ioutil"

	hamquestions "github.com/lulumel0n/arrl-ham-questions-pool-proto/ham-questions"
	pbtext "google.golang.org/protobuf/encoding/prototext"
	pb "google.golang.org/protobuf/proto"
)

const outputDir = "../out/"
const generalText = "2019-2023_general"
const rawQuestionsTxt = "../raw-questions/2019-2023_general.txt"

func main() {

	qpb, err := hamquestions.NewHamQuestion("", rawQuestionsTxt)

	if err != nil {
		panic(err)
	}

	out, err := pb.Marshal(qpb.Pool)
	if err != nil {
		panic(err)
	}
	if err := ioutil.WriteFile(outputDir+generalText+"_pb", out, 0644); err != nil {
		panic(err)
	}

	msout, err := pbtext.MarshalOptions{}.Marshal(qpb.Pool)
	if err != nil {
		fmt.Println(err)
	}
	if err := ioutil.WriteFile(outputDir+generalText+"_pbtext", msout, 0644); err != nil {
		panic(err)
	}

	fmt.Printf("Generated output to: %s\n", outputDir)
}
