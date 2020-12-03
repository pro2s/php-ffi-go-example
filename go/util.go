package main

import (
    "hash/crc32"
    "C"
    "fmt"
    "strings"
    "strconv"
    "github.com/elliotchance/phpserialize"
)

type Option struct {
	id int
	linked_id int
	empty bool
}

//export combine
func combine(id int, combinations [][]Option) *C.char {
    mixed := mix(combinations)
    filtered := filter(mixed)
    out := format(id, filtered)

    ret, err := phpserialize.Marshal(out, nil)
	if err != nil {
		panic(err)
    }

    return C.CString(string(ret))
}

func wrap(slice []Option) [][]Option {
    res := [][]Option{}
    for _, value := range slice {
        res = append(res, []Option{value})
    }

    return res
}

func mix(slice [][]Option) [][]Option {
    if len(slice) == 0 {
        return [][]Option{};
    }

    if len(slice) == 1 {
        return wrap(slice[0])
    }

    combination, tail := slice[0], slice[1:]
    combinations := mix(tail)

    return goCombine(combination, combinations);
}

func simpleCombine(combination []Option, combinations [][]Option) [][]Option {
    res := [][]Option{}

    for _, value := range combination {
        res = append(res, add(value, combinations)...)
    }

    return res
}

func goCombine(combination []Option, combinations [][]Option) [][]Option {
    res := [][]Option{}
    count := len(combination)
    channel := make(chan [][]Option, count)

    for _, value := range combination {
        go chAdd(channel, value, combinations)
    }

    for r := range channel {
        count -= 1
        res = append(res, r...)

        if count == 0 {
            break
        }
    }

    return res
}

func chAdd(channel chan<-[][]Option, value Option, combinations [][]Option) {
    channel <- add(value, combinations)
}

func add(value Option, combinations [][]Option) [][]Option {
    res := make([][]Option, len(combinations))

    for i, combine := range combinations {
        res[i] = append([]Option{value}, combine...)
    }

    return res
}

func gethash(id int, ids []int) string {
    stringIds := make([]string, len(ids) + 1)
    stringIds[0] = strconv.Itoa(id)
    for i, id := range ids {
        stringIds[i+1] = strconv.Itoa(id)
    }

    data := strings.Join(stringIds, "-")
    hash := crc32.ChecksumIEEE([]byte(data))

    return fmt.Sprintf("%x", hash)
}

func format(id int, combinations [][]Option) map[string][]int {
    res := make(map[string][]int, len(combinations))

    for _, combination :=range combinations {
        ids := []int{}
        for _, option := range combination {
            if !option.empty {
                ids = append(ids, option.id)
            }
        }

        res[gethash(id, ids)] = ids
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
    ids := make(map[int]bool, len(options))
    linked := []int{}

    for _, option := range options {
        if !option.empty {
            ids[option.id] = true
            if option.linked_id != 0 {
                linked = append(linked, option.linked_id)
            }
        }
    }

    res := true

    for _, key := range linked {
        _, ok := ids[key]
        res = res && ok
    }

    return res
}

func main() {}
