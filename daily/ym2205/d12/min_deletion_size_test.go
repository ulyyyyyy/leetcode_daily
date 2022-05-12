package d12

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

func Test_minDeletionSize(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{
			name: "删列造序-测试用例1",
			args: args{
				strs: []string{"cba", "daf", "ghi"},
			},
			wantAns: 1,
		},
		{
			name: "删列造序-测试用例2",
			args: args{
				strs: []string{"a", "b"},
			},
			wantAns: 0,
		},
		{
			name: "删列造序-测试用例3",
			args: args{
				strs: []string{"zyx", "wvu", "tsr"},
			},
			wantAns: 3,
		},
		{
			name: "删列造序-测试用例4-error",
			args: args{
				strs: []string{"rrjk", "furt", "guzm"},
			},
			wantAns: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := minDeletionSize(tt.args.strs); gotAns != tt.wantAns {
				t.Errorf("minDeletionSize() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}

const (
	Len       = 5
	LetterLen = 5
)

func Test_constructLetters(t *testing.T) {
	bTotal := make([]string, Len)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < Len; i++ {
		bStr := "\""
		for j := 0; j < LetterLen; j++ {
			b := 'a' + byte(rand.Int()%26)
			bStr += string(b)
		}
		bTotal[i] = bStr + "\""
	}
	fmt.Println(strings.Join(bTotal, ","))
}
