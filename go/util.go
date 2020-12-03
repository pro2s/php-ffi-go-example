package main

import (
    "./base"
    "./util"
    "C"
    "github.com/elliotchance/phpserialize"
)

type Option = base.Option

//export combine
func combine(id int, combinations [][]Option) *C.char {
    mixed := mix(combinations, base.GoCombine)
    filtered := filter(mixed)
    formater := formatWithId(id)
    out := format(filtered, base.GoFormat, formater)

    ret, err := phpserialize.Marshal(out, nil)
	if err != nil {
		panic(err)
    }

    return C.CString(string(ret))
}

func mix(slice [][]Option, combine base.Combiner) [][]Option {
    if len(slice) == 0 {
        return [][]Option{};
    }

    if len(slice) == 1 {
        return base.Wrap(slice[0])
    }

    combination, tail := slice[0], slice[1:]
    combinations := mix(tail, combine)

    return combine(combination, combinations);
}

func filter(combinations [][]Option) [][]Option {
    res := combinations[:0]
    for _, x := range combinations {
        if base.CorrectLinked(x) {
            res = append(res, x)
        }
    }

    return res;
}

func format(combinations [][]Option, formater base.Formater, format base.OptionsFormater) map[string][]int {
    return formater(format, combinations)
}

func formatWithId(id int) base.OptionsFormater {
    return func (options []Option) (string, []int) {
        ids := base.GetIds(options)

        return util.GetHash(id, ids), ids
    }
}

func main() {}
