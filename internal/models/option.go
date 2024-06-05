package models

import "sort"

type Option struct {
	Key   string
	Value string
}

type Options []Option

func (arr Options) GetKey(index int) string {
	return arr[index].Key
}

func (arr Options) GetValue(index int) string {
	return arr[index].Value
}

func (arr Options) Size() int {
	return len(arr)
}

func (arr Options) Get() []Option {
	return arr
}

func (arr *Options) AppendUnique(data string) []Option {
	uniqueSet := make(map[string]struct{})
	result := make(Options, arr.Size())
	for i, obj := range arr.Get() {
		uniqueSet[obj.Value] = struct{}{}
		result[i] = obj
	}

	if _, exists := uniqueSet[data]; !exists {
		result = append(arr.Get(), Option{Key: data, Value: data})
	}

	return result
}

func (arr *Options) Sort() Options {
	sort.Slice(*arr, func(i, j int) bool {
		return arr.GetKey(i) < arr.GetKey(j)
	})

	return *arr
}
