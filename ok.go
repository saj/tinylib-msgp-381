package main

//go:generate msgp

type Flibbity struct {
	A Flibs `msg:"a,omitempty"`
	B Flibs `msg:"b,omitempty"`
}

type Flibs []Flib

type Flib struct {
	X int8 `msg:"x"`
	Y int8 `msg:"y"`
}
