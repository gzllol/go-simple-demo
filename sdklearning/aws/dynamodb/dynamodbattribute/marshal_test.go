package test

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Person struct {
	Name string
	Age  int32
}

func TestMarshalBasic(t *testing.T) {
	p := &Person{
		Name: "Tony",
		Age:  19,
	}
	av, _ := dynamodbattribute.MarshalMap(p)
	fmt.Println(av)
}

type PersonWithTag struct {
	Name string `dynamodbav:"my_name"`
	Age  int32  `dynamodbav:"my_age"`
}

func TestMarshalTag(t *testing.T) {
	p := &PersonWithTag{
		Name: "Tony",
		Age:  20,
	}
	av, _ := dynamodbattribute.MarshalMap(p)
	fmt.Println(av)
}
