package base

type Option struct {
	id int
	linked_id int
	empty bool
}
type Combination []Option
type Combinations [][]Option

func (options Combination) GetIds() []int {
	ids := []int{}
    for _, option := range options {
        if !option.empty {
            ids = append(ids, option.id)
        }
	}

	return ids;
}

var EmptyOption = Option{empty: true}

func NewOption(id int, linked_id int) Option {
	return Option{id: id, linked_id: linked_id}
}