package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func knock0() {
	s := string("stressed")
	k := []string{}
	for i := 7; i >= 0; i-- {
		k = append(k, string(s[i]))
	}
	fmt.Println(strings.Join(k, ""))
}

func knock1() {
	s := "パタトクカシーー"
	k := []rune(s)
	m := []rune{k[0], k[2], k[4], k[6]}
	fmt.Println(string(m))
}

func knock2() {
	s := "パタトクカシーー"
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

func knock5(v string, n int) {

	//まず単語n-gram
	k := strings.Split(v, " ")
	var m []string
	//ここちょっと理解浅いがstringの単語ごとのスライスに直してる。
	for _, str := range k {
		m = append(m, str)
	}
	//ここまででとりあえず与えられた文字列を単語ごとに分けてmというスライスに格納した
	//以下で、このスライスをnずつ束ねた
	for i := 0; i < len(m)-1; i++ {
		fmt.Println(m[i : i+n])
	}
	//ここから文字n-gram
	//文字列にあるスペースの除去
	w := strings.Replace(v, " ", "", -1)
	fmt.Println(w)
	//単語ごとのスライスに直してたものを文字ごとのスライスに直したい。
	for i := 0; i < len(w)-1; i++ {
		fmt.Println(w[i : i+n])
	}
}

func knock6(x, y string, n int) {
	//1単語に対する文字n-gram
	var X, Y []string
	for i := 0; i < len(x)-1; i++ {
		X = append(X, x[i:i+n])
	}
	for i := 0; i < len(y)-1; i++ {
		Y = append(Y, y[i:i+n])
	}
	//ここまででn-gramの行ごとにスライスに格納するのはできた。
	fmt.Println(X, Y) //↑の確認
	//ここから和集合Z1、積集合Z2、差集合Z3

	//一旦パス

}

func knock7(x int, y string, z float64) {
	fmt.Printf("%v時の%vは%v", x, y, z)
}

func cipher(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Println( /*func219的な何か*/ (v))
		//一旦パス

	default:
		fmt.Println(v)
	}
}

func knock8(v string) {
	k := strings.Split(v, " ")
	var m []string
	//ここちょっと理解浅いがstringの単語ごとのスライスに直してる。
	for _, str := range k {
		m = append(m, str)
	}
	//ここまででとりあえず与えられた文字列を単語ごとに分けてmというスライスに格納した

	for i := 0; i < len(m)-1; i++ {
		var n string = string(m[i])
		if len(n) > 4 {

			var s int = int(len(n) - 2)
			var N string = string(n[1:s])
			//shuffle関数自体はおそらくbyteとしてまとまってる単語列にしか対応してない。
			//なので、各単語を文章化してスペース開けたものをシャッフルに処理してもらうか、単語レベルで処理してくれる関数を見つけるか
			//あるいは順番をあてがうか

			Ns := strings.Split(N, "")
			u := len(Ns)
			for c := 0; c < u-1; c++ {
				//fmt.Printf("%s ", slice[c])
				//fmt.Println(reflect.TypeOf(slice[c])) //取り出した文字の型を示す。
			}
			fmt.Println("")
			//ここまでで単語をスライスに格納するところまでは行けた。
			//以下でそれを、1.ランダムに動かす

			rand.Shuffle(u, func(i, j int) {
				Ns[i], Ns[j] = Ns[j], Ns[i]
			})

			v = strings.Join(Ns, "")
			//fmt.Printf("%T, %v", v, v)
			//一旦v含めてsliceにぶちこむ
			w := []string{}
			w = append(w, string(n[0]))
			w = append(w, v)
			w = append(w, string(n[s+1]))
			//ここまでで、混ぜた文字と元の文字をwっていうスライスにぶちこむことに成功

			fmt.Println(w)
			//joinを使って、元の文字列と結合
			//v0 := []string{string(n[0])}
			//vi := []string{string(n[s+1])}
			//w0 := strings.Join(v0, v)
			//w1 := strings.Join(w0,vi)
			//fmt.Println(w0)
			//fmt.Println(w1)
			//すでにランダムに動かせてるから、単語に結合したい。　2.元の単語と結合して単語にする

		}

		//fmt.Println(n)
	}

}

func main() {
	/*
		knock0()
		knock1()
		knock2()
		knock3()
		knock4()
		knock5("I am an NLPer", 2)
		knock6("paraparaparadise", "paragraph", 2)
		knock7(12, "気温", 22.4)
		cipher("paraparaparadise")
	*/
	knock8("I couldn't believe that I could actually understand what I was reading : the phenomenal power of the human mind.")
}
