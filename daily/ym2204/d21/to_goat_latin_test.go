package d21

import "testing"

func Test_toGoatLatin(t *testing.T) {
	type args struct {
		sentence string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "山羊拉丁文-测试用例1",
			args: args{
				sentence: "I speak Goat Latin",
			},
			want: "Imaa peaksmaaa oatGmaaaa atinLmaaaaa",
		},
		{
			name: "山羊拉丁文-测试用例2",
			args: args{
				sentence: "The quick brown fox jumped over the lazy dog",
			},
			want: "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa",
		},
		{
			name: "山羊拉丁文-测试用例3",
			args: args{
				sentence: "Each word consists of lowercase and uppercase letters only",
			},
			want: "Eachmaa ordwmaaa onsistscmaaaa ofmaaaaa owercaselmaaaaaa andmaaaaaaa uppercasemaaaaaaaa etterslmaaaaaaaaa onlymaaaaaaaaaa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toGoatLatin(tt.args.sentence); got != tt.want {
				t.Errorf("toGoatLatin() = %v, want %v", got, tt.want)
			}
		})
	}
}
