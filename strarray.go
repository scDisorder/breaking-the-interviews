package main

type StringArray []string

func (s StringArray) Len() int           { return len(s) }
func (s StringArray) Less(i, j int) bool { return s[i] < s[j] }
func (s StringArray) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

type Uint8Array []uint8

func (s Uint8Array) Len() int           { return len(s) }
func (s Uint8Array) Less(i, j int) bool { return s[i] < s[j] }
func (s Uint8Array) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
