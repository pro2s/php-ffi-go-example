package base

type Option struct {
	id int
	linked_id int
	empty bool
}
type Combination []Option
type Combinations [][]Option

func Wrap(slice Combination) Combinations {
    res := Combinations{}
    for _, value := range slice {
        res = append(res, Combination{value})
    }

    return res
}

func isCorrectLinked(options Combination) bool {
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

func (options Combination) GetIds() []int {
	ids := []int{}
    for _, option := range options {
        if !option.empty {
            ids = append(ids, option.id)
        }
	}

	return ids;
}

func (option Option) AddTo(combinations Combinations) Combinations {
    res := make(Combinations, len(combinations))

    for i, combine := range combinations {
        res[i] = append(Combination{option}, combine...)
    }

    return res
}

func (combinations Combinations) Filter() Combinations {
    res := combinations[:0]
    for _, x := range combinations {
        if isCorrectLinked(x) {
            res = append(res, x)
        }
    }

    return res;
}
