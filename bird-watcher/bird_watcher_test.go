package birdwatcher

import (
	"reflect"
	"testing"
)

func TestTotalBirdCount(t *testing.T) {
	tests := []struct {
		name       string
		birdCounts []int
		want       int
	}{
		{
			name:       "calculates the correct total number of birds",
			birdCounts: []int{9, 0, 8, 4, 5, 1, 3},
			want:       30,
		},
		{
			name:       "works for a short bird count list",
			birdCounts: []int{2},
			want:       2,
		},
		{
			name:       "works for a long bird count list",
			birdCounts: []int{2, 8, 4, 1, 3, 5, 0, 4, 1, 6, 0, 3, 0, 1, 5, 4, 1, 1, 2, 6},
			want:       57,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TotalBirdCount(tt.birdCounts); got != tt.want {
				t.Errorf("TotalBirdCount(%v) = %v; want %v", tt.birdCounts, got, tt.want)
			}
		})
	}
}

func TestBirdsInWeek(t *testing.T) {
	tests := []struct {
		name       string
		birdCounts []int
		week       int
		want       int
	}{
		{
			name:       "calculates the number of birds in the first week",
			birdCounts: []int{3, 0, 5, 1, 0, 4, 1, 0, 3, 4, 3, 0, 8, 0},
			week:       1,
			want:       14,our bird watcher intuition also tells you that the bird was in your garden on the first day that you tracked in your list.


			want:       12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BirdsInWeek(tt.birdCounts, tt.week); got != tt.want {
				t.Errorf("BirdsInWeek(%v) = %v; want %v", tt.birdCounts, got, tt.want)
			}
		})
	}
}

func TestFixBirdCount(t *testing.T) {
	tests := []struct {
		name       string
		birdCounts []int
		week       int
		want       []int
	}{
		{
			name:       "returns a bird count list with the corrected values",
			birdCounts: []int{3, 0, 5, 1, 0, 4, 1, 0, 3, 4, 3, 0},
			want:       []int{4, 0, 6, 1, 1, 4, 2, 0, 4, 4, 4, 0},
		},
		{
			name:       "works for a short bird count list",
			birdCounts: []int{4, 2},
			want:       []int{5, 2},
		},
		{
			name:       "works for a long bird count list",
			birdCounts: []int{2, 8, 4, 1, 3, 5, 0, 4, 1, 6, 0, 3, 0, 1, 5, 4, 1, 1, 2, 6},
			want:       []int{3, 8, 5, 1, 4, 5, 1, 4, 2, 6, 1, 3, 1, 1, 6, 4, 2, 1, 3, 6},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := make([]int, len(tt.birdCounts))
			copy(input, tt.birdCounts)

			if got := FixBirdCountLog(tt.birdCounts); !reflect.DeepEqual(tt.want, got) {
				t.Errorf("FixBirdCountLog(%v) = %v; want %v", input, got, tt.want)
			}
		})
	}

}

func TestFixBirdCountDoesNotCreateNewSlice(t *testing.T) {
	counts := []int{4, 0, 6, 1, 1, 4, 2, 0, 4, 4, 4, 0}
	got := FixBirdCountLog(counts)
	if reflect.ValueOf(got).Pointer() != reflect.ValueOf(counts).Pointer() {
		t.Error("it looks like that you are creating a new slice in the function FixBirdCountLog - " +
			"please make sure you are modifying the slice passed as argument")
	}
}
