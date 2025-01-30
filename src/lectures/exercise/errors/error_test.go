package timeparse

import "testing"

type TestUnit struct {
	time   string
	result bool
}

func TestTimeParser(t *testing.T) {
	test_units := []TestUnit{
		{"12:00:12", true},
		{"23:59:59", true},
		{"00:00:00", true},
		{"ab:00:00", false},
		{"00:ab:00", false},
		{"00:00:ab", false},
		{"  :00:00", false},
		{"00:  :00", false},
		{"24:00:00", false},
		{"12:60:00", false},
		{"12:00:60", false},
	}

	for _, i := range test_units {
		_, err := TimeParser(i.time)
		if (err != nil) == i.result {
			t.Errorf("unvalid input \n%v", err)
		}
	}
}
