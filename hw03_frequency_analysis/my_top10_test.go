package hw03frequencyanalysis

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidMyTopTenFreqWords(t *testing.T) {
	tests := []struct {
		input    string
		expected map[string]int
	}{
		{"cat and dog, one dog,two cats and one man", map[string]int{
			"and":     2,
			"one":     2,
			"cat":     1,
			"cats":    1,
			"dog,":    1,
			"dog,two": 1,
			"man":     1}},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			mtt := NewMyTopTen(tc.input, false, 10)
			mtt.freqWords()
			require.Equal(t, tc.expected, mtt.freq)
		})
	}
}

func TestValidMyTopTenSortFreq(t *testing.T) {
	tests := []struct {
		input    string
		expected Kvs
	}{
		{"cat and dog, one dog,two cats and one man", []Kv{
			{"and", 2},
			{"one", 2},
			{"cat", 1},
			{"cats", 1},
			{"dog,", 1},
			{"dog,two", 1},
			{"man", 1},
		},
		}}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			mtt := NewMyTopTen(tc.input, false, 10)
			mtt.freqWords()
			kvs := mtt.sortFreq()
			require.Equal(t, tc.expected, *kvs)
		})
	}
}

func TestValidMyTopTenGetTop(t *testing.T) {
	tests := []struct {
		input    string
		top      int
		expected []string
	}{
		{"cat and dog, one dog,two cats and one man", 10, []string{"and", "one", "cat", "cats", "dog,", "dog,two", "man"}},
		{"cat and dog, one dog,two cats and one man", 7, []string{"and", "one", "cat", "cats", "dog,", "dog,two", "man"}},
		{"cat and dog, one dog,two cats and one man", 3, []string{"and", "one", "cat"}},
		{"cat and dog, one dog,two cats and one man", 2, []string{"and", "one"}},
		{"cat and dog, one dog,two cats and one man", 1, []string{"and"}},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			mtt := NewMyTopTen(tc.input, false, tc.top)
			mtt.freqWords()
			kvs := mtt.sortFreq()
			result := mtt.getTop(kvs)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestValidMyTopTenIsEmpty(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"", true},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result := NewMyTopTen(tc.input, false, 10).isEmpty()
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestInvalidMyTopTenIsEmpty(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"почем опиум для народа", false},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result := NewMyTopTen(tc.input, false, 10).isEmpty()
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestValidMyTopTenMakeAsteriks(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			`Если вас интересует мое мнение – я выскажусь: настоящая дружба (именно дружба, а не шапочное знакомство или приятельские отношения) проверяется в радости; умение разделить радость другого человека – этим сегодня могут похвастаться немногие…очень немногие «друзья»!`,
			`если вас интересует мое мнение я выскажусь настоящая дружба именно дружба а не шапочное знакомство или приятельские отношения проверяется в радости умение разделить радость другого человека этим сегодня могут похвастаться немногие очень немногие друзья`,
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.input, func(t *testing.T) {
			result := NewMyTopTen(tc.input, false, 10).makeAsteriks(tc.input)
			require.Equal(t, tc.expected, result)
		})
	}
}
