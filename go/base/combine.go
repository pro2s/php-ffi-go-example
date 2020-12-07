package base
type Combiner func ([]Option, Combinations) Combinations

func (slice Combinations) Mix() Combinations {
	return slice.mix(combine)
}

func (slice Combinations) GoMix() Combinations {
	return slice.mix(goCombine)
}

func (slice Combinations) mix(combine Combiner) Combinations {
    switch len(slice) {
	case 0:
		return Combinations{};
	case 1:
		return Wrap(slice[0])
	default:
	    combination, tail := slice[0], slice[1:]
    	combinations := tail.mix(combine)

		return combine(combination, combinations)
	}
}

func combine(combination []Option, combinations Combinations) Combinations {
    res := Combinations{}

    for _, value := range combination {
        res = append(res, value.AddTo(combinations)...)
    }

    return res
}

func chAdd(channel chan<- Combinations, value Option, combinations Combinations) {
    channel <- value.AddTo(combinations)
}

func goCombine(combination []Option, combinations Combinations) Combinations {
    res := Combinations{}
    count := len(combination)
    channel := make(chan Combinations, count)

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
