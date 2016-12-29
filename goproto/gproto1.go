package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/hangshiisi/gohello/goproto/example"
	"log"
)

func examMain() {
	test := &example.Test{
		Label: proto.String("hello"),
		Type:  proto.Int32(17),
		Reps:  []int64{1, 2, 3},
		Optionalgroup: &example.Test_OptionalGroup{
			RequiredField: proto.String("good bye"),
		},
	}
	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	newTest := &example.Test{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	// Now test and newTest contain the same data.
	if test.GetLabel() != newTest.GetLabel() {
		log.Fatalf("data mismatch %q != %q", test.GetLabel(), newTest.GetLabel())
	} else {
		fmt.Printf("data match %q == %q from new and old test data \n",
			test.GetLabel(), newTest.GetLabel())
	}
	// etc.
}
