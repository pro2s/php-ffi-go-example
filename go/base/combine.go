package base
type Combiner func (Combination, Combinations) Combinations

func (slice *Combinations) Mix() Combinations {
	return slice.mix(combine)
}

func (slice *Combinations) GoMix() Combinations {
	return slice.mix(goCombine)
}

func wrap(slice Combination) Combinations {
    res := Combinations{}
    for _, value := range slice {
        res = append(res, Combination{value})
    }

    return res
}

func (option Option) addTo(combinations Combinations) Combinations {
    res := make(Combinations, len(combinations))

    for i, combine := range combinations {
        res[i] = append(Combination{option}, combine...)
    }

    return res
}

func (slice Combinations) mix(combine Combiner) Combinations {
    switch len(slice) {
	case 0:
		return Combinations{};
	case 1:
		return wrap(slice[0])
	default:
	    combination, tail := slice[0], slice[1:]
    	combinations := tail.mix(combine)

		return combine(combination, combinations)
	}
}

func combine(combination Combination, combinations Combinations) Combinations {
    res := Combinations{}

    for _, value := range combination {
        res = append(res, value.addTo(combinations)...)
    }

    return res
}

func chAdd(channel chan<- Combinations, value Option, combinations Combinations) {
    channel <- value.addTo(combinations)
}

func goCombine(combination Combination, combinations Combinations) Combinations {
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
