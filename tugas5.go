package main

import "fmt"
import "math/rand"

type pelajar struct {
	nama string
	umur int
}

func main() {
	var n1, n2, n3 pelajar
	n1.nama = "Ahmad"
	n1.umur = rand.Intn(100)
	n2.nama = "Abi"
	n2.umur = rand.Intn(100)
	n3.nama = "Zaka"
	n3.umur = rand.Intn(100)

	fmt.Println(n1.nama, "\t", n1.umur)
	fmt.Println(n2.nama, "\t", n2.umur)
	fmt.Println(n3.nama, "\t", n3.umur)

}
