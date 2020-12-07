package main

import (
    "./base"
    "./util"
    "C"
    "github.com/elliotchance/phpserialize"
)

type Option = base.Option
type Combinations = base.Combinations

//export combine
func combine(id int, combinations [][]Option) *C.char {
    mixed := mix(combinations, base.GoCombine)
    filtered := filter(mixed)
    formater := formatWithId(id)
    out := filtered.GoFormat(formater)

    ret, err := phpserialize.Marshal(out, nil)
	if err != nil {
		panic(err)
    }

    return C.CString(string(ret))
}

func mix(slice Combinations, combine base.Combiner) Combinations {
    if len(slice) == 0 {
        return Combinations{};
    }

    if len(slice) == 1 {
        return base.Wrap(slice[0])
    }

    combination, tail := slice[0], slice[1:]
    combinations := mix(tail, combine)

    return combine(combination, combinations);
}

func filter(combinations Combinations) Combinations {
    res := combinations[:0]
    for _, x := range combinations {
        if base.CorrectLinked(x) {
            res = append(res, x)
        }
    }

    return res;
}

func formatWithId(id int) base.OptionsFormater {
    return func (options []Option) (string, []int) {
        ids := base.GetIds(options)

        return util.GetHash(id, ids), ids
    }
}

func main() {}
