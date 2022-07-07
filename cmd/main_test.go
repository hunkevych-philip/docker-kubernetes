package main

import "testing"

func TestDummy(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Some dummy test.",
			want: TestString,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Dummy(); got != tt.want {
				t.Errorf("Test() = %v, want %v", got, tt.want)
			}
		})
	}
}
