package main

import (
	"syscall/js"
	"./base"
    "./util"
)

func main() {
	done := make(chan struct{}, 0)
	global := js.Global()
	global.Set("wasmCombine", js.FuncOf(combine))
	<-done
}

type Ids []int
type Option = base.Option
type Combination = base.Combination
type Combinations = base.Combinations

func GoProccess(in Combinations, formater base.Formater) map[string][]int {
    return in.GoMix().Filter().GoFormat(formater)
}

func (ids Ids) toJsValue() js.Value {
    arr := make([]interface{}, len(ids))
    for i, id := range ids {
        arr[i] = id
    }

    return js.ValueOf(arr)
}

func getCombinations(value js.Value) Combinations {
    // return "ERROR: Array accepted [[Option, null],[...]]"
    combinations := Combinations{}

	for i := 0; i < value.Length(); i++ {
		combination := Combination{}
		subValue := value.Index(i)
		//	return "ERROR: Array accepted [Option, null, ...]"
		for j := 0; j < subValue.Length(); j++ {
			if (subValue.Index(j).IsNull()) {
				combination = append(combination, base.EmptyOption)
			} else {
                id := subValue.Index(j).Get("id").Int()
                linked_id := subValue.Index(j).Get("linked_id").Int()
				combination = append(combination, base.NewOption(id, linked_id))
			}
        }
        combinations = append(combinations, combination)
    }

    return combinations
}

func combine(this js.Value, args []js.Value) interface{} {
	if len(args) != 2 {
		return "ERROR: number of arguments doesn't match"
    }

	combinations := getCombinations(args[1])
	formater := formatWithId(args[0].Int())
    out := GoProccess(combinations, formater)

    outJs := js.ValueOf(map[string]interface{}{})
    outJs.Set("total", len(out))
    for k, v := range out {
        outJs.Set(k, Ids(v).toJsValue())
    }

	return outJs
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
