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
}
func knock3() {
	s := "Now I need a drink, alcoholic of course, after the heavy lectures involving quantum mechanics."
	k := strings.Split(s, " ")
	var m []string
	//fmt.Println(k)

	//ここちょっと理解浅い
	for _, str := range k {
		//fmt.Printf("[%s]", str)
		m = append(m, str)
	}

	//fmt.Println(m)

	var n []int

	for i := 0; i < 15; i++ {
		p := len(m[i])
		n = append(n, p)
	}
	fmt.Println(n)
}

func main() {
	knock0()
	knock1()
	knock2()
	knock3()
}
