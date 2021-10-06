package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

const (
	TopTen int = 10
)

// Kvs - Kv collection.
type Kvs []Kv

// Kv - key value struct.
type Kv struct {
	Key   string
	Value int
}

// MyTopTen - main struct for this task.
type MyTopTen struct {
	withAsteriks bool
	countTop     int
	text         string
	freq         map[string]int
}

// NewMyTopTen - create new MyTopTen.
func NewMyTopTen(text string, withAsteriks bool, top int) *MyTopTen {
	return &MyTopTen{
		withAsteriks: withAsteriks,
		countTop:     top,
		text:         text,
		freq:         make(map[string]int, 2),
	}
}

func (mtt *MyTopTen) Top10() []string {
	if mtt.isEmpty() {
		return nil
	}

	mtt.freqWords()
	kvs := mtt.sortFreq()
	return mtt.getTop(kvs)
}

// isEmpty - checking if the text is empty.
func (mtt *MyTopTen) isEmpty() (empty bool) {
	if len(mtt.text) == 0 {
		empty = true
	}
	return
}

// freqWords - splitting into words and
// counting the frequency of occurrence of each word.
func (mtt *MyTopTen) freqWords() {
	text := mtt.text

	if mtt.withAsteriks {
		text = mtt.makeAsteriks(text)
	}

	words := strings.Fields(text)
	for _, w := range words {
		if _, ok := mtt.freq[w]; !ok {
			mtt.freq[w] = 1
		} else {
			mtt.freq[w]++
		}
	}
}

// makeAsteriks - doing the lowercase and remove punctuation
func (mtt *MyTopTen) makeAsteriks(text string) string {
	text = strings.ToLower(text)
	text = strings.ReplaceAll(text, ",", "")
	text = strings.ReplaceAll(text, ".", "")
	text = strings.ReplaceAll(text, "!", "")
	text = strings.ReplaceAll(text, "?", "")
	text = strings.ReplaceAll(text, ":", "")
	text = strings.ReplaceAll(text, ";", "")
	text = strings.ReplaceAll(text, "(", "")
	text = strings.ReplaceAll(text, ")", "")
	text = strings.ReplaceAll(text, "\"", "")
	text = strings.ReplaceAll(text, "«", "")
	text = strings.ReplaceAll(text, "»", "")
	text = strings.ReplaceAll(text, "…", " ")
	text = strings.ReplaceAll(text, " – ", " ")
	text = strings.ReplaceAll(text, " - ", " ")
	return text
}

// sortFreq - sorting the obtaind sequense.
func (mtt *MyTopTen) sortFreq() *Kvs {
	i := 0
	_freq := make(Kvs, len(mtt.freq))
	for k, v := range mtt.freq {
		_freq[i].Key = k
		_freq[i].Value = v
		i++
	}

	sort.Slice(_freq, func(i, j int) bool {
		if _freq[i].Value != _freq[j].Value {
			return _freq[i].Value > _freq[j].Value
		}
		return _freq[i].Key < _freq[j].Key
	})
	return &_freq
}

// getTop - get top of freq words.
func (mtt *MyTopTen) getTop(kvs *Kvs) []string {
	count := mtt.countTop
	if len(*kvs) < mtt.countTop {
		count = len(*kvs)
	}

	ranked := make([]string, count)
	for i, kv := range *kvs {
		if i == count {
			break
		}
		ranked[i] = kv.Key
	}
	return ranked
}
