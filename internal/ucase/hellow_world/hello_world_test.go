package main

import "testing"

//func TestHello(t *testing.T) {
//	type args struct {
//		subject string
//	}
//
//	tests := []struct {
//		name string
//		args args
//		want string
//	}{
//		{
//			name: "first hello",
//			args: args{
//				"Moms",
//			},
//			want: "Hello, Moms",
//		},
//		{
//			name: "hello moms with empty string",
//			args: args{
//				subject: "",
//			},
//			want: "Hello, Moms",
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := Hello(tt.args.subject); got != tt.want {
//				t.Errorf("Hello() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Branz", "")
		want := "Hello, Branz"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, Moms!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
