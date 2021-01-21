package main

import (
	"fmt"

	"github.com/imsilence/radix"
)

func main() {
	radix := radix.New()
	radix.Add("abcd", "abcd")
	radix.Add("abc", "abc")
	radix.Add("abde", "abde")
	radix.Add("abdd", "abdd")
	radix.Add("bcd", "bcd")
	radix.Add("bcef", "bcef")
	radix.Add("bcee", "bcee")

	fmt.Println(radix)

	fmt.Println(radix.Get(""))
	fmt.Println(radix.Get("ab"))
	fmt.Println(radix.Get("abc"))
	fmt.Println(radix.Get("abd"))
	fmt.Println(radix.Get("abde"))
	fmt.Println(radix.Get("abdd"))
	fmt.Println(radix.Get("bc"))
	fmt.Println(radix.Get("bcd"))
	fmt.Println(radix.Get("bce"))
	fmt.Println(radix.Get("bcef"))
	fmt.Println(radix.Get("bcee"))
	fmt.Println(radix.Get("c"))
	fmt.Println(radix.Get("bcf"))
	fmt.Println(radix.Get("bceg"))
	fmt.Println(radix.Get("bcefg"))

	fmt.Println(radix.Delete("bcefg"))
	fmt.Println(radix.Delete("abc"))
	fmt.Println(radix.Delete("bcef"))
	fmt.Println(radix.Delete("abcd"))
	fmt.Println(radix)

	// fmt.Println(radix.GetValue("bc"))
	// fmt.Println(radix.GetValue("bcd"))
	// fmt.Println(radix.GetValue("bce"))
	// fmt.Println(radix.GetValue("bcef"))
	// fmt.Println(radix.GetValue("bcee"))
	// fmt.Println(radix.GetValue("abcd"))
}
