package main

import (
    "./base"
    "./util"
    "C"
    "github.com/elliotchance/phpserialize"
)

type Option = base.Option
type Combination = base.Combination
type Combinations = base.Combinations

//export combine
func combine(id int, combinations [][]Option) *C.char {
    formater := formatWithId(id)

    out := GoProccess(combinations, formater)

    ret, err := phpserialize.Marshal(out, nil)
	if err != nil {
		panic(err)
    }

    return C.CString(string(ret))
}

func GoProccess(in Combinations, formater base.Formater) map[string][]int {
    return in.GoMix().Filter().GoFormat(formater)
}

func Proccess(in Combinations, formater base.Formater) map[string][]int {
    return in.Mix().Filter().Format(formater)
}

func formatWithId(id int) base.Formater {
    return func (options Combination) (string, []int) {
        ids := options.GetIds()

        return util.GetHash(id, ids), ids
    }
}

func main() {}
