package minimum_time

import "testing"

func Test_minimumTime(t *testing.T) {
	type args struct {
		time       []int
		totalTrips int
	}
	tests := []struct {
		name      string
		args      args
		wantCount int64
	}{
		{
			name: "",
			args: args{
				time:       []int{1, 2, 3},
				totalTrips: 5,
			},
			wantCount: 3,
		},
		{
			name: "",
			args: args{
				time:       []int{2},
				totalTrips: 1,
			},
			wantCount: 2,
		},
		{
			name: "",
			args: args{
				time:       []int{5, 10, 10},
				totalTrips: 9,
			},
			wantCount: 25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := minimumTime(tt.args.time, tt.args.totalTrips); gotCount != tt.wantCount {
				t.Errorf("minimumTime() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}
