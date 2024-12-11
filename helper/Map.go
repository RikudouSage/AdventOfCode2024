package helper

func Map[TInput any, TOutput any](input []TInput, mapFunc func(item TInput) TOutput) []TOutput {
	output := make([]TOutput, len(input))
	for i, item := range input {
		output[i] = mapFunc(item)
	}

	return output
}
