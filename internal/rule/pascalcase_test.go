package rule

import "testing"

func TestPascalCase(t *testing.T) {
	var rule = new(PascalCase).Init()

	var tests = []*ruleTest{
		{value: "pascal", expected: false, err: nil},
		{value: "pascalcase", expected: false, err: nil},
		{value: "pascalCase", expected: false, err: nil},
		{value: "Pascalcase", expected: true, err: nil},
		{value: "PascalCase", expected: true, err: nil},
		{value: "Pascal1Case", expected: true, err: nil},
		{value: "PascalVCase", expected: true, err: nil},
		{value: "PascalCaseForever", expected: true, err: nil},
		{value: "PASCALCASE", expected: false, err: nil},
		{value: "pascal_case", expected: false, err: nil},
		{value: "pascal.case", expected: false, err: nil},
		{value: "pascal-case", expected: false, err: nil},
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
