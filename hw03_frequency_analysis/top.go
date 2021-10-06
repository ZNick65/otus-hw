package hw03frequencyanalysis

const asterisk bool = true

func Top10(s string) []string {
	// Place your code here.
	return NewMyTopTen(s, asterisk, TopTen).Top10()
}
