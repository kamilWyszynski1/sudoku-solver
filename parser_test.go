package sudoku_solver

import (
	"fmt"
	"testing"
)

func TestParser_Parse(t *testing.T) {
	type args struct {
		arrangement string
	}
	tests := []struct {
		name    string
		args    args
		want    *SudokuBoard
		wantErr bool
	}{
		{
			name:    "basic",
			args:    args{arrangement: "310004069000000200008005040000000005006000017807030000590700006600003050000100002"},
			want:    nil,
			wantErr: false,
		},
		{
			name:    "one digit",
			args:    args{arrangement: "000000000000000000000000000000000000000010000000000000000000000000000000000000000"},
			want:    nil,
			wantErr: false,
		},
		{
			name:    "wiki",
			args:    args{arrangement: "530070000600195000098000060800060003400803001700020006060000280000419005000080079"},
			want:    nil,
			wantErr: false,
		},
		//{
		//	name:    "anit bruteforce",
		//	args:    args{arrangement: "000000000000003085001020000000507000004000100090000000500000073002010000000040009"},
		//	want:    nil,
		//	wantErr: false,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Parser{}
			got, _ := p.Parse(tt.args.arrangement)
			fmt.Println(got)
			if err := got.Solve(); err != nil {
				fmt.Println(err)
			}
			fmt.Println(got)
		})
	}
}
