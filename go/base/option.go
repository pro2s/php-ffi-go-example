package base

func (combinations Combinations) Filter() Combinations {
    res := combinations[:0]
    for _, x := range combinations {
        if isCorrectLinked(x) {
            res = append(res, x)
        }
    }

    return res;
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

    for _, key := range linked {
		_, ok := ids[key]
		if !ok {
			return false;
		}
    }

    return true
}
