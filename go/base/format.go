package base

type OptionsFormater func ([]Option) (string, []int)
type Formater func (OptionsFormater, [][]Option) map[string][]int

type Result struct {
	ids []int
    hash string
}

func Format(format OptionsFormater, combinations [][]Option) map[string][]int {
    res := make(map[string][]int, len(combinations))

    for _, combination := range combinations {
        hash, ids := format(combination)
        res[hash] = ids
    }

    return res
}

func worker(format OptionsFormater, jobs <-chan []Option, results chan<- Result) {
    for j := range jobs {
        hash, ids := format(j)
        results <- Result{ids, hash}
    }
}

func GoFormat(format OptionsFormater, combinations [][]Option) map[string][]int {
    res := map[string][]int{}
    numJobs := len(combinations)

    jobs := make(chan []Option, numJobs)
    results := make(chan Result, numJobs)

    for w := 1; w <= 10; w++ {
        go worker(format, jobs, results)
    }

    for _, j := range combinations {
        jobs <- j
    }

    close(jobs)

    for numJobs > 0 {
        r := <- results
        res[r.hash] = r.ids
        numJobs -= 1
    }

    return res
}
