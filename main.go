package main

import (
	"fmt"
	"math/rand"
	"strings"

	mapset "github.com/deckarep/golang-set"
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

	m = append(m, k...)

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

	m = append(m, k...)
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
	m = append(m, k...)
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
	var X, Y []interface{}
	for i := 0; i < len(x)-1; i++ {
		X = append(X, x[i:i+n])
	}
	for i := 0; i < len(y)-1; i++ {
		Y = append(Y, y[i:i+n])
	}
	//ここまででn-gramの行ごとにスライスに格納するのはできた。
	//X,Yを[]interface{}に格納すれば、union関数が使えそう。

	XX := mapset.NewSetFromSlice(X)
	YY := mapset.NewSetFromSlice(Y)

	Z1 := XX.Union(YY)
	Z2 := XX.Intersect(YY)
	Z3 := XX.Difference(YY)

	fmt.Println(Z1)
	fmt.Println(Z2)
	fmt.Println(Z3)
}

func knock7(x int, y string, z float64) {
	fmt.Printf("%v時の%vは%v \n", x, y, z)
}

func cipher(v string) {
	//この時すでにv[i]はbyteコード？になってる
	//スライスに格納した文字を1文字ずつrangeで処理して、小文字はasciiの219であるブロックに置き換えたい

	//var s []string
	var t []string
	for i := 0; i <= len(v)-1; i++ { //vの文字数をうまく出力したい
		//fmt.Println(v[i])
		//ここまでの出力は正常なので、このifのところが問題ありそう

		if v[i] >= 97 && v[i] <= 122 {
			t = append(t, string(219))
			//v[i] := 219 //文字コードを文字に起こす処理
		} else {
			t = append(t, string(v[i]))
		}
	}
	fmt.Println(strings.Join(t, ""))

}

func knock9(v string) {
	k := strings.Split(v, " ")
	var m []string
	m = append(m, k...)
	//ここまででとりあえず与えられた文字列を単語ごとに分けてmというスライスに格納した

	for i := 0; i < len(m); i++ {
		var n string = string(m[i])
		if len(n) > 4 {

			var s int = int(len(n) - 2)
			var N string = string(n[1 : s+1])
			//shuffle関数自体はおそらくbyteとしてまとまってる単語列にしか対応してない。
			//なので、各単語を文章化してスペース開けたものをシャッフルに処理してもらう

			Ns := strings.Split(N, "")
			u := len(Ns)
			for c := 0; c < u+2; c++ {
			}
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
			//これを1文字に直したい
			m[i] = strings.Join(w, "")
			//joinを使って、元の文字列と結合した
		}
	}
	fmt.Println(m)

}

func main() {
	knock0()
	knock1()
	knock2()
	knock3()
	knock4()
	knock5("I am an NLPer", 2)
	knock6("paraparaparadise", "paragraph", 2)
	knock7(12, "気温", 22.4)
	cipher("parAparDSDASGadisGe")
	knock9("I couldn't believe that I could actually understand what I was reading : the phenomenal power of the human mind.")
}
