package main

import (
    "C"
    "fmt"
    "github.com/elliotchance/phpserialize"
)

//export print
func print(out *C.char) {
    fmt.Println("[GO print] " + C.GoString(out))
}

//export sum
func sum(a C.int, b C.int) C.int {
    return a + b
}

//export combine
func combine(data string) *C.char {
	fmt.Println(data)

	var in map[interface {}]interface {}
    err := phpserialize.Unmarshal([]byte(data), &in)
    if err != nil {
		panic(err)
    }

    fmt.Println(in)

    out := map[int]map[string]interface{}{
        0: map[string]interface{}{"id": 1, "link_id": nil},
        1: map[string]interface{}{"id": 2, "link_id": 1},
    }

    fmt.Println(out)

    ret, err := phpserialize.Marshal(out, nil)
	if err != nil {
		panic(err)
    }

    return C.CString(string(ret))
}

func main() {}
