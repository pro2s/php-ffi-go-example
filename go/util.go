package main

import (
    "hash/crc32"
    "C"
    "fmt"
    "strings"
    "strconv"
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

//export serialize
func serialize(data string) *C.char {
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

type Option struct {
	id int
	linked_id int
	empty bool
}

//export combine
func combine(id int, combinations [][]Option) *C.char {
    mixed := mix(combinations)
    fmt.Println("Init: ", len(mixed))

    filtered := filter(mixed)
    fmt.Println("Filter: ", len(filtered))

    out := format(id, filtered)
    fmt.Println("Format: ", len(out))

    ret, err := phpserialize.Marshal(out, nil)
	if err != nil {
		panic(err)
    }

    return C.CString(string(ret))
}

func mix(slice [][]Option) [][]Option {
    res := [][]Option{}

    if len(slice) == 0 {
        return res;
    }

    if len(slice) == 1 {
        for _, value := range slice[0] {
            res = append(res, []Option{value})
        }

        return res
    }

    combination, tail := slice[0], slice[1:]
    combinations := mix(tail)

    for _, value := range combination {
        for _, combine := range combinations {
            res = append(res, append([]Option{value}, combine...))
        }
    }


    return res;
}

func format(id int, combinations [][]Option) map[string]map[string]interface{} {
    res := make(map[string]map[string]interface{})

    for _, combination :=range combinations {
        ids := []int{}
        idsString := []string{strconv.Itoa(id)}
        for _, option := range combination {
            if !option.empty {
                ids = append(ids, option.id)
                idsString = append(idsString, strconv.Itoa(option.id))
            }
        }
        data := strings.Join(idsString, "-")
        hash := crc32.ChecksumIEEE([]byte(data))
        hashString := fmt.Sprintf("%x", hash)
        res[hashString] = map[string]interface{}{"ids": ids}
    }

    return res
}

func filter(combinations [][]Option) [][]Option {
    res := combinations[:0]
    for _, x := range combinations {
        if filterLinked(x) {
            res = append(res, x)
        }
    }

    return res;
}

func filterLinked(options []Option) bool {
    id := map[int]bool{}
    linked := []int{}
    test := []int{}

    for _, option := range options {
        if !option.empty {
            id[option.id] = true
            test = append(test, option.id)
            if option.linked_id != 0 {
                linked = append(linked, option.linked_id)
            }
        }
    }

    res := true

    for _, key := range linked {
        _, ok := id[key]
        res = res && ok
    }

    return res
}

func main() {}
