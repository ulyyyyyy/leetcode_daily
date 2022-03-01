package prefix_count

import "testing"

func Test_prefixCount(t *testing.T) {
	type args struct {
		words []string
		pref  string
	}
	tests := []struct {
		name      string
		args      args
		wantCount int
	}{
		{
			name: "",
			args: args{
				words: []string{"pay", "attention", "practice", "attend"},
				pref:  "at",
			},
			wantCount: 2,
		},
		{
			name: "",
			args: args{
				words: []string{"leetcode", "win", "loops", "success"},
				pref:  "code",
			},
			wantCount: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := prefixCount(tt.args.words, tt.args.pref); gotCount != tt.wantCount {
				t.Errorf("prefixCount() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}
