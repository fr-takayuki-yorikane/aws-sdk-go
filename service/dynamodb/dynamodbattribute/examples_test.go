package dynamodbattribute_test

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"reflect"
)

func ExampleConvertTo() {
	type Record struct {
		MyField string
		Letters []string
		Numbers []int
	}

	r := Record{
		MyField: "MyFieldValue",
		Letters: []string{"a", "b", "c", "d"},
		Numbers: []int{1, 2, 3},
	}
	av, err := dynamodbattribute.ConvertTo(r)
	fmt.Println(av, err)

	// Output:
	// {
	//   M: {
	//     MyField: {
	//       S: "MyFieldValue"
	//     },
	//     Letters: {
	//       L: [
	//         {
	//           S: "a"
	//         },
	//         {
	//           S: "b"
	//         },
	//         {
	//           S: "c"
	//         },
	//         {
	//           S: "d"
	//         }
	//       ]
	//     },
	//     Numbers: {
	//       L: [{
	//           N: "1"
	//         },{
	//           N: "2"
	//         },{
	//           N: "3"
	//         }]
	//     }
	//   }
	// } <nil>
}

func ExampleConvertFrom() {
	type Record struct {
		MyField string
		Letters []string
		A2Num   map[string]int
	}

	r := Record{
		MyField: "MyFieldValue",
		Letters: []string{"a", "b", "c", "d"},
		A2Num:   map[string]int{"a": 1, "b": 2, "c": 3},
	}
	av, err := dynamodbattribute.ConvertTo(r)

	r2 := Record{}
	err = dynamodbattribute.ConvertFrom(av, &r2)
	fmt.Println(err, reflect.DeepEqual(r, r2))

	// Output:
	// <nil> true
}