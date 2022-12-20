package sms

import mapset "github.com/deckarep/golang-set/v2"

type Strategy func([]string) []string

func Order(gties []string) []string {
	return gties
}

func Random(gties []string) []string {
	sset := mapset.NewSet(gties...)
	return sset.ToSlice()
}
