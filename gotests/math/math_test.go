package math

import "testing"

func TestSum(t *testing.T) {
	result := sum(1, 2)

	if result != 3 {
		t.Errorf("sum was incorrect, got: %d, want: %d.", result, 3)
	}
}

func TestSumStr(t *testing.T) {
	// Table Drive Test
	testCases := []struct {
		Name     string
		A        string
		B        string
		Expected int
	}{
		{
			Name:     "Sum two numbers",
			A:        "1",
			B:        "2",
			Expected: 3,
		},
		{
			Name:     "Incorrect A parameter",
			A:        "a",
			B:        "2",
			Expected: 0,
		},
		{
			Name:     "Incorrect B parameter",
			A:        "1",
			B:        "b",
			Expected: 0,
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			result := sumStr(tc.A, tc.B)

			t.Log("Result:", result)

			if result != tc.Expected {
				t.Errorf("sum was incorrect, got: %d, want: %d.", result, tc.Expected)
			}
		})
	}

}
