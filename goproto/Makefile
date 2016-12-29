

.PHONY: all go

all: go gotest


go:     add_person_go     list_people_go
gotest: add_person_gotest list_people_gotest


clean:
	rm -f protoc_middleman addressbook.pb.cc addressbook.pb.h addressbook_pb2.py com/example/tutorial/AddressBookProtos.java
	rm -f *.pyc
	rm -f protoc_middleman_go example/*.pb.go add_person_go list_people_go

protoc_middleman_go: addressbook.proto
	mkdir -p example  # make directory for go package
	protoc --go_out=example addressbook.proto
	@touch protoc_middleman_go

add_person_go: add_person.go protoc_middleman_go
	go build -o add_person_go add_person.go

add_person_gotest: add_person_test.go add_person_go
	go test add_person.go add_person_test.go

list_people_go: list_people.go protoc_middleman_go
	go build -o list_people_go list_people.go

list_people_gotest: list_people.go list_people_go
	go test list_people.go list_people_test.go

