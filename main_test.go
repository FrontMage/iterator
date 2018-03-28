package iterator

import (
	"fmt"
	"reflect"
	"testing"
)

func TestChunk(t *testing.T) {
	type args struct {
		s     []interface{}
		chunk int
	}
	tests := []struct {
		name string
		args args
		want [][]interface{}
	}{
		{
			name: "chunk with zero length slice",
			args: args{
				s:     []interface{}{},
				chunk: 10,
			},
			want: [][]interface{}{[]interface{}{}},
		},
		{
			name: "chunk with just divided length slice",
			args: args{
				s:     []interface{}{1, 2, 3, 4, 5, 6},
				chunk: 2,
			},
			want: [][]interface{}{
				[]interface{}{1, 2, 3},
				[]interface{}{4, 5, 6},
			},
		},
		{
			name: "chunk with chunk is greater than slice length",
			args: args{
				s:     []interface{}{1, 2, 3, 4, 5, 6},
				chunk: 10,
			},
			want: [][]interface{}{
				[]interface{}{1, 2, 3, 4, 5, 6},
			},
		},
		{
			name: "chunk with not just divided length slice",
			args: args{
				s:     []interface{}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
				chunk: 3,
			},
			want: [][]interface{}{
				[]interface{}{1, 2, 3},
				[]interface{}{4, 5, 6},
				[]interface{}{7, 8, 9},
				[]interface{}{10},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Chunk(tt.args.s, tt.args.chunk); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chunk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIter(t *testing.T) {
	s := []interface{}{}
	for i := 0; i < 10*10000; i++ {
		s = append(s, i)
	}
	type args struct {
		s           []interface{}
		callback    func(idx int, item interface{})
		workerCount int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test iter with 10 workers",
			args: args{
				s: s,
				callback: func(idx int, item interface{}) {
				},
				workerCount: 10,
			},
		},
		{
			name: "test iter with 100 workers",
			args: args{
				s: s,
				callback: func(idx int, item interface{}) {
				},
				workerCount: 100,
			},
		},
		{
			name: "test iter with 0 workers",
			args: args{
				s: s,
				callback: func(idx int, item interface{}) {
				},
				workerCount: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Iter(tt.args.s, tt.args.callback, tt.args.workerCount)
		})
	}
}

func BenchmarkIter(b *testing.B) {
	s := []interface{}{}
	for i := 0; i < 10*100; i++ {
		s = append(s, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Iter(s, func(idx int, item interface{}) {
			fmt.Println(item)
		}, 10)
	}
}

func BenchmarkNormal(b *testing.B) {
	s := []interface{}{}
	for i := 0; i < 10*100; i++ {
		s = append(s, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, item := range s {
			fmt.Println(item)
		}
	}
}
