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
		{
			name: "引数が1",
			args: 1,
			want: []int{1},
		},
		{
			name: "引数が0以下",
			args: 0,
			want: nil,
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
