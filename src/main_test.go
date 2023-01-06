package main

import "testing"

func Test_CreateNumList(t *testing.T) {
	tests := []struct {
		name string
		args int
		want []int
	}{
		{
			name: "正常",
			args: 3,
			want: []int{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createNumList(tt.args); len(got) != len(tt.want) {
				t.Errorf("createNumList() = %v, want %v", got, tt.want)
			}
		})
	}

}
