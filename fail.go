package main

//go:generate msgp

type Flobbity struct {
	A Flobs `msg:"a,omitempty"`
	B Flobs `msg:"b,omitempty"`
}

type Flobs []Flob

type Flob struct {
	X Numberwang `msg:"x"`
	Y int8       `msg:"y"`
}

type Numberwang int8
