package main

import (
	"fmt"

	hamquestions "github.com/jkl73/arrl-ham-questions-pool-proto/ham-questions"
)

const outputDir = "../out/"
const generalText = "2019-2023_general"
const rawQuestionsGeneralFile = "../raw-questions/2019-2023_general.txt"
const rawQuestionsTechnicianFile = "../raw-questions/2018-2022_technician.txt"
const rawQuestionsExtraFile = "../raw-questions/2020-2024_extra.txt"

func main() {

	generalqpb, gentitles, err := hamquestions.NewHamQuestionsAndTitles("", rawQuestionsGeneralFile, hamquestions.General)
	if err != nil {
		fmt.Println("General")
		panic(err)
	}
	fmt.Println(generalqpb)
	fmt.Println(gentitles)

	techqpb, techtitles, err := hamquestions.NewHamQuestionsAndTitles("", rawQuestionsTechnicianFile, hamquestions.Tech)
	if err != nil {
		fmt.Println("Tech")
		panic(err)
	}
	fmt.Println(techqpb)
	fmt.Println(techtitles)

	extraqpb, extratitles, err := hamquestions.NewHamQuestionsAndTitles("", rawQuestionsExtraFile, hamquestions.Extra)
	if err != nil {
		fmt.Println("EXTRA")
		panic(err)
	}
	fmt.Println(extraqpb)
	fmt.Println(extratitles)

	// out, err := pb.Marshal(qpb.SubelementMap)
	// if err != nil {
	// 	panic(err)
	// }
	// if err := ioutil.WriteFile(outputDir+generalText+"_pb", out, 0644); err != nil {
	// 	panic(err)
	// }

	// msout, err := pbtext.MarshalOptions{}.Marshal(qpb.Pool)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// if err := ioutil.WriteFile(outputDir+generalText+"_pbtext", msout, 0644); err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("Generated output to: %s\n", outputDir)
}
