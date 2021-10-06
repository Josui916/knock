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

	//ここちょっと理解浅いがstringの単語ごとのスライスに直してる。
	for _, str := range k {
		//fmt.Printf("[%s]", str)
		m = append(m, str)
	}

	//上で単語ごとのスライスにしたものをそれぞれ長さ出力したのをそれぞれアペンドでnというスライスに格納している。
	var n []int
	for i := 0; i < 15; i++ {
		p := len(m[i])
		n = append(n, p)
	}
	fmt.Println(n)
}
func knock4() {
	s := "Hi He Lied Because Boron Could Not Oxidize Fluorine. New Nations Might Also Sign Peace Security Clause. Arthur King Can."
	k := strings.Split(s, " ")
	var m []string
	//fmt.Println(k)

	//ここちょっと理解浅いがstringの単語ごとのスライスに直してる。
	for _, str := range k {
		//fmt.Printf("[%s]", str)
		m = append(m, str)
	}
	n := make(map[int]string)

	for i := 0; i < 20; i++ {
		p := m[i]
		if i == 0 {
			n[i] = p[0:1]
		} else if i < 4 {
			n[i] = p[0:2]
		} else if i < 9 {
			n[i] = p[0:1]
		} else if i < 14 {
			n[i] = p[0:2]
		} else if i < 16 {
			n[i] = p[0:1]
		} else if i < 18 {
			n[i] = p[0:2]
		} else if i == 18 {
			n[i] = p[0:1]
		} else {
			n[i] = p[0:2]
		}

	}
	fmt.Println(n)

}

func main() {
	knock0()
	knock1()
	knock2()
	knock3()
	knock4()
}
