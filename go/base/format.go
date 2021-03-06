package base

type Formater func (Combination) (string, []int)

type Result struct {
	ids []int
    hash string
}

func (combinations Combinations) Format(format Formater) map[string][]int {
    res := make(map[string][]int, len(combinations))

    for _, combination := range combinations {
        hash, ids := format(combination)
        res[hash] = ids
    }

    return res
}

func worker(format Formater, jobs <-chan Combination, results chan<- Result) {
    for j := range jobs {
        hash, ids := format(j)
        results <- Result{ids, hash}
    }
}

func (combinations Combinations) GoFormat(format Formater) map[string][]int {
    res := map[string][]int{}
    numJobs := len(combinations)

    jobs := make(chan Combination, numJobs)
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
