package main

import (
	"reflect"
	"testing"
)

func TestRuneDecode(t *testing.T) {
	type args struct {
		char rune
		dial int
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
		{"t0", args{'A', 0}, 'A'},
		{"t1", args{'A', 10}, 'K'},
		{"t2", args{'A', 26}, '0'},
		{"t3", args{'9', 0}, '9'},
		{"t4", args{'9', 1}, 'A'},
		{"t5", args{'A', -1}, '9'},
		{"t6", args{'a', 0}, 'A'},
		{"t7", args{'a', 10}, 'K'},
		{"t8", args{'=', 10}, '='},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RuneDecode(tt.args.char, tt.args.dial); got != tt.want {
				t.Errorf("RuneDecode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRuneDistance(t *testing.T) {
	type args struct {
		r1 rune
		r2 rune
	}
	tests := []struct {
		name           string
		args           args
		wantDistance   int
		wantSafeIgnore bool
	}{
		{"t1", args{'A', 'A'}, 0, false},
		{"t2", args{'A', 'K'}, 10, false},
		{"t3", args{'A', '9'}, 35, false},
		{"t4", args{'9', 'A'}, 1, false},
		{"t5", args{'=', '='}, 0, true},
		{"t6", args{'=', '1'}, -1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDistance, gotSafeIgnore := RuneDistance(tt.args.r1, tt.args.r2)
			if gotDistance != tt.wantDistance {
				t.Errorf("RuneDistance() gotDistance = %v, want %v", gotDistance, tt.wantDistance)
			}
			if gotSafeIgnore != tt.wantSafeIgnore {
				t.Errorf("RuneDistance() gotSafeIgnore = %v, want %v", gotSafeIgnore, tt.wantSafeIgnore)
			}
		})
	}
}

func TestDecodeLine(t *testing.T) {
	type args struct {
		line  string
		dials []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"t1", args{"AAAA", []int{0, 1}}, "ABAB"},
		{"t2", args{"AAAA", []int{0, 1, 2}}, "ABCA"},
		{"t3", args{"AAAABBBB", []int{0, 1, 2}}, "ABCACDBC"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecodeLine(tt.args.line, tt.args.dials); got != tt.want {
				t.Errorf("DecodeLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTryMatching(t *testing.T) {
	type args struct {
		line     string
		sample   string
		numDials int
	}
	tests := []struct {
		name        string
		args        args
		wantDials   []int
		wantSuccess bool
	}{
		{"t1", args{"AAAA", "ABAB", 2}, []int{0, 1}, true},
		{"t2", args{"S7MFSSM A1M2V 9SI", "AVERAGE SPEED ", 4}, []int{18, 24, 28, 12}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDials, gotSuccess := TryMatching(tt.args.line, tt.args.sample, tt.args.numDials)
			if !reflect.DeepEqual(gotDials, tt.wantDials) {
				t.Errorf("TryMatching() gotDials = %v, want %v", gotDials, tt.wantDials)
			}
			if gotSuccess != tt.wantSuccess {
				t.Errorf("TryMatching() gotSuccess = %v, want %v", gotSuccess, tt.wantSuccess)
			}
		})
	}
}
