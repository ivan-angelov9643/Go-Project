package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/api/object1", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("object1"))
	})
	http.HandleFunc("/api/object2", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("object2"))
	})
	http.ListenAndServe(":8080", nil)

	//a := 2
	//b := 3
	//if a > b {
	//	fmt.Println("a > b")
	//} else if a == b {
	//	fmt.Println("a == b")
	//} else {
	//	fmt.Println("a < b")
	//}
	//switch a + b {
	//case 1:
	//	fmt.Println("a + b = 1")
	//case 2:
	//	fmt.Println("a + b = 2")
	//case 3:
	//	fmt.Println("a + b = 3")
	//case 5:
	//	fmt.Println("a + b = 5")
	//	fallthrough
	//default:
	//	fmt.Println("a + b")
	//}

	//nums := []int{1, 3, 4, 5, 6, 234, 634}
	//for i, d := range nums {
	//	fmt.Println(i, d)
	//}
	//
	//n := make([]int, len(nums), 20)
	//n[2] = 4
	//fmt.Println(n)

	//m := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
	//fmt.Println(m["a"])

	//for t := 10; t >= 0; t-- {
	//	if t == 0 {
	//		fmt.Println("boom")
	//		break
	//	}
	//	fmt.Println(t)
	//}
}

func max1(finishes ...int) int {
	best := finishes[0]
	for _, i := range finishes {
		if i > best {
			best = i
		}
	}
	return best
}
