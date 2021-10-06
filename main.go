package main

import (
	"fmt"
	"strings"
)

func knock0() {
	s := string("stressed")
	k := []string{}
	for i := 7; i >= 0; i-- {
		//fmt.Println(string(s[i]))
		k = append(k, string(s[i]))
	}
	fmt.Println(strings.Join(k, ""))
}

func knock1() {
	s := "パタトクカシーー"
	//fmt.Println(string([]rune(s)))
	k := []rune(s)
	m := []rune{k[0], k[2], k[4], k[6]}
	fmt.Println(string(m))

	/*
		var k []string
		k = append(k, string(s[2]), string(s[4]), string(s[6]), string(s[8]))
		fmt.Println(string(s[2]))
		fmt.Println(strings.Join(k, ""))
	*/
}

func knock2() {
	s := "パタトクカシーー"
	//fmt.Println(string([]rune(s)))
	k := []rune(s)
	m := []rune{k[0], k[2], k[4], k[6]}
	M := string(m)
	n := []rune{k[1], k[3], k[5], k[7]}
	N := string(n)
	fmt.Println(M + N)

	/*
		var k []string
		k = append(k, string(s[2]), string(s[4]), string(s[6]), string(s[8]))
		fmt.Println(string(s[2]))
		fmt.Println(strings.Join(k, ""))
	*/
}

func main() {
	knock0()
	knock1()
	knock2()
}
