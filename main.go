package main

import (
	"flag"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"
)

var (
	client ssmiface.SSMAPI
)

func init() {
	session := session.Must(session.NewSession())
	client = ssm.New(session)
}

type parameter struct {
	name    string
	value   string
	element string
}

func main() {
	file := flag.String("file", "", "Name of the template file")
	output := flag.String("output", "", "Name for the output file")
	flag.Parse()
	if *file == "" || *output == "" {
		flag.Usage()
		os.Exit(1)
	}

	text := readFile(*file)
	createFile(fetchParameters(text), text, *output)
}

func readFile(templatePath string) string {
	dat, err := ioutil.ReadFile(templatePath)
	check(err)
	return string(dat)
}

func createFile(parameters []parameter, text string, output string) {
	for _, parameter := range parameters {
		text = strings.Replace(text, parameter.element, parameter.value, -1)
	}
	ioutil.WriteFile(output, []byte(text), 0644)
}

func fetchParameters(text string) []parameter {
	var parameters []parameter
	paramRegex, _ := regexp.Compile("{{.*}}")
	removeBracketsRegex, _ := regexp.Compile("\\{{2}|\\}{2}")
	for _, token := range paramRegex.FindAllString(text, -1) {
		name := removeBracketsRegex.ReplaceAllString(token, "")
		if !retrieved(name, parameters) {
			value := fetchParametersFromSSM(name)
			parameters = append(parameters, parameter{name, value, token})
		}
	}
	return parameters
}

func fetchParametersFromSSM(name string) string {
	parameter, err := client.GetParameter(&ssm.GetParameterInput{
		Name:           &name,
		WithDecryption: aws.Bool(true),
	})
	check(err)
	return *parameter.Parameter.Value
}

func retrieved(value string, array []parameter) bool {
	for _, element := range array {
		if element.name == value && element.value != "" {
			return true
		}
	}
	return false
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
