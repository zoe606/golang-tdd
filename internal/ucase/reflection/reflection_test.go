package reflection

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         any
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Nolan"},
			[]string{"Nolan"},
		},
		{
			"struct with two string field",
			struct {
				Name string
				City string
			}{"Nolan", "Bandung"},
			[]string{"Nolan", "Bandung"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Nolan", 20},
			[]string{"Nolan"},
		},
		{
			"nested struct",
			Person{
				"Nolan",
				Profile{20, "Bandung"}},
			[]string{"Nolan", "Bandung"},
			//struct {
			//	Name    string
			//	Profile struct {
			//		Age  int
			//		City string
			//	}
			//}{"Nolan", struct {
			//	Age  int
			//	City string
			//}{20, "Bandung"}},
			//[]string{"Nolan", "Bandung"},
		},
		{
			"Pointer to things",
			&Person{"Nolan", Profile{20, "Bandung"}},
			[]string{"Nolan", "Bandung"},
		},
		{
			"slices",
			[]Profile{{20, "Bandung"}, {21, "Jogja"}},
			[]string{"Bandung", "Jogja"},
		},
		{
			"arrays",
			[2]Profile{{20, "Bandung"}, {21, "Yogya"}},
			[]string{"Bandung", "Yogya"},
		},

		// maps in Go do not guarantee order move test case
		//{
		//	"maps",
		//	map[string]string{"Foo": "Bar", "Baz": "Boz"},
		//	[]string{"Bar", "Boz"},
		//},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})

	t.Run("With channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{20, "Bandung"}
			aChannel <- Profile{21, "Jogja"}
			close(aChannel)
		}()

		var got []string
		expected := []string{"Bandung", "Jogja"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got %v, want %v", got, expected)
		}
	})

	t.Run("With functin", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{20, "Bandung"}, Profile{21, "Jogja"}
		}

		var got []string
		expected := []string{"Bandung", "Jogja"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got %v, want %v", got, expected)
		}
	})
}

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
	}
}
