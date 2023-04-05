package main

import "testing"

func Test_bfsTreeSearch(t *testing.T) {
	type args struct {
		start State
		goalP GoalP
		succ  SuccessorsP
	}
	tests := []struct {
		name string
		args args
		want State
	}{
		{
			name: "test1",
			args: args{
				start: 1,
				goalP: stateIsGoal(4),
				succ:  binaryTree,
			},
			want: 4,
		},
		{
			name: "test2",
			args: args{
				start: 1,
				goalP: stateIsGoal(5),
				succ:  binaryTree,
			},
			want: 5,
		},
		{
			name: "test3",
			args: args{

				start: 1,
				goalP: stateIsGoal(6),
				succ:  binaryTree,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bfsTreeSearch(tt.args.start, tt.args.goalP, tt.args.succ); got != tt.want {
				t.Errorf("bfsTreeSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
