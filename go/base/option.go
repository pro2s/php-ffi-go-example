package base

type Option struct {
	id int
	linked_id int
	empty bool
}

type Combinations [][]Option

func Wrap(slice []Option) Combinations {
    res := [][]Option{}
    for _, value := range slice {
        res = append(res, []Option{value})
    }

    return res
}

func CorrectLinked(options []Option) bool {
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

func GetIds(options []Option) []int {
	ids := []int{}
    for _, option := range options {
        if !option.empty {
            ids = append(ids, option.id)
        }
	}

	return ids;
}

func (option Option) AddTo (combinations Combinations) Combinations {
    res := make(Combinations, len(combinations))

    for i, combine := range combinations {
        res[i] = append([]Option{option}, combine...)
    }

    return res
}

