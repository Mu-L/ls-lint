package rule

import "testing"

func TestKebabCase(t *testing.T) {
	var rule = new(KebabCase).Init()

	var tests = []*ruleTest{
		{value: "kebab", expected: true, err: nil},
		{value: "kebabcase", expected: true, err: nil},
		{value: "kebabCase", expected: false, err: nil},
		{value: "Kebabcase", expected: false, err: nil},
		{value: "KebabCase", expected: false, err: nil},
		{value: "KEBABCASE", expected: false, err: nil},
		{value: "kebab-case", expected: true, err: nil},
		{value: "kebab-case-test", expected: true, err: nil},
		{value: "kebab-123-test", expected: true, err: nil},
		{value: "kebab.test", expected: false, err: nil},
		{value: "kebab_test", expected: false, err: nil},
	}

	var i = 0
	for _, test := range tests {
		res, err := rule.Validate(test.value)

		if err != nil && err != test.err {
			t.Errorf("Test %d failed with unmatched error - %s", i, err.Error())
			return
		}

		if res != test.expected {
			t.Errorf("Test %d failed with unmatched return value - %+v", i, res)
			return
		}

		i++
	}
}
