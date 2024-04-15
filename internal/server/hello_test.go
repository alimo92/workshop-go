package server

import "testing"

func TestHello(t *testing.T) {

	tests := map[string]struct {
		input string
		want  string
	}{
		"greet someone": {input: "Ali", want: "Hello Ali"},
		"greet world":   {input: "", want: "Hello World"},
	}

	for i, tc := range tests {
		t.Run(i, func(t *testing.T) {
			assertCorrectMessage(t, print(tc.input), tc.want)
		})
	}
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
