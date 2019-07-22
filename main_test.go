package main

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"
)

type mockSSMClient struct {
	ssmiface.SSMAPI
}

var mockedParameter = parameter{
	name:    "/dev/example/value",
	value:   "myvalue",
	element: "{{/dev/example/value}}",
}

var mockedParameters = []parameter{
	mockedParameter,
}

var mockedText = "value={{/dev/example/value}}"

func (m *mockSSMClient) GetParameter(input *ssm.GetParameterInput) (*ssm.GetParameterOutput, error) {
	output := ssm.GetParameterOutput{
		Parameter: &ssm.Parameter{
			Name:  &mockedParameter.name,
			Value: &mockedParameter.value,
		},
	}
	return &output, nil
}
func TestFetchParametersFromSSM(t *testing.T) {
	client = &mockSSMClient{}
	value := fetchParametersFromSSM(mockedParameter.name)
	if value != mockedParameter.value {
		t.Errorf("Value did not match, expected %s, found %s", mockedParameter.value, value)
	}
}

func TestFetchParameters(t *testing.T) {
	parameters := fetchParameters(mockedText)
	if len(parameters) == 0 {
		t.Error("Parameters were not retrieved correctly, expected >0, found =0")
	}
	for _, parameter := range parameters {
		if parameter.name != mockedParameter.name {
			t.Error("Name did not match")
		}
		if parameter.element != mockedParameter.element {
			t.Error("Element did not match")
		}
		if parameter.value != mockedParameter.value {
			t.Error("Value did not match")
		}
	}
}

func TestRetrieved(t *testing.T) {
	if !retrieved(mockedParameter.name, mockedParameters) {
		t.Error("Expected retrieved value")
	}
}
