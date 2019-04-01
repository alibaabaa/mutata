package main

import (
	"crypto/rand"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
)

type flagData struct {
	infile     string
	input      string
	outfile    string
	showInput  bool
	transforms []string
}

func getFlagData() flagData {
	infile := flag.String("infile", "", "file input data")
	input := flag.String("input", "", "string input data")
	outfile := flag.String("outfile", "", "output file")
	showInput := flag.Bool("show-input", false, "prints input data to stdout")

	flag.Parse()
	return flagData{*infile, *input, *outfile, *showInput, flag.Args()}
}

func getRandFromInput(input string) []byte {
	re := regexp.MustCompile(`rand\((\d+)\)`)
	var match string
	if re.MatchString(input) {
		match = re.ReplaceAllString(input, "$1")
	}

	if len(match) == 0 {
		return nil
	}

	byteCount, err := strconv.Atoi(match)
	if err != nil {
		return nil
	}

	randomBytes := make([]byte, byteCount)
	_, err = rand.Read(randomBytes)
	if err != nil {
		return nil
	}

	return randomBytes
}

func getInputData(fData flagData) ([]byte, error) {
	if (fData.infile == "" && fData.input == "") ||
		(fData.infile != "" && fData.input != "") {
		return nil, errors.New("Exactly one of infile or input must be provided")
	}

	if randInput := getRandFromInput(fData.input); len(randInput) > 0 {
		return randInput, nil
	}

	if fData.input != "" {
		return []byte(fData.input), nil
	}

	info, err := os.Stat(fData.infile)
	if os.IsNotExist(err) || info.IsDir() {
		return nil, errors.New("infile not found")
	}

	fileData, err := ioutil.ReadFile(fData.infile)
	if err != nil {
		return nil, err
	}

	return fileData, nil
}

func runMutations(input []byte, mutations []string) []byte {
	result := input
	mutationRegister := getMutations()
	for i := 0; i < len(mutations); i++ {
		mutation := mutationRegister[mutations[i]]
		result = mutation(result)
	}
	return result
}

func main() {
	fData := getFlagData()
	inputData, err := getInputData(fData)
	if err != nil {
		panic(err)
	}

	if fData.showInput {
		fmt.Println("Input (hex):")
		fmt.Println(string(runMutations(inputData, []string{"hex"})))
		fmt.Println()
	}

	result := runMutations(inputData, fData.transforms)

	if fData.outfile != "" {
		err := ioutil.WriteFile(fData.outfile, result, 0644)
		if err != nil {
			panic(err)
		}
	} else {
		fmt.Print(string(result))
	}
}
