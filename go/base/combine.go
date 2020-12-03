package base
type Combiner func ([]Option, [][]Option) [][]Option

func Combine(combination []Option, combinations [][]Option) [][]Option {
    res := [][]Option{}

    for _, value := range combination {
        res = append(res, value.AddTo(combinations)...)
    }

    return res
}

func chAdd(channel chan<- [][]Option, value Option, combinations [][]Option) {
    channel <- value.AddTo(combinations)
}

func GoCombine(combination []Option, combinations [][]Option) [][]Option {
    res := [][]Option{}
    count := len(combination)
    channel := make(chan [][]Option, count)

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
