package main

import "fmt"

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev := Reverse(input)
	doubleRev := Reverse(rev)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q\n", rev)
	fmt.Printf("reversed again: %q\n", doubleRev)
}

func Reverse(s string) string {
	bytes := []byte(s)
	for i, j := 0, len(bytes)-1; i < len(bytes)/2; i, j = i+1, j-1 {
		bytes[i], bytes[j] = bytes[j], bytes[i]
	}
	return string(bytes)
}
