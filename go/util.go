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

    out := Combinations(combinations).GoMix().Filter().GoFormat(formater)

    ret, err := phpserialize.Marshal(out, nil)
	if err != nil {
		panic(err)
    }

    return C.CString(string(ret))
}

func formatWithId(id int) base.OptionsFormater {
    return func (options Combination) (string, []int) {
        ids := options.GetIds()

        return util.GetHash(id, ids), ids
    }
}

func main() {}
