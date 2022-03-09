package main

type TestStruct struct {
	Field1      int
	TestField2  string
	ExistingTag string
	Embed
}

type Embed struct {
}
