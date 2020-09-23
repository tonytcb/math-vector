package vector

import (
	"testing"
)

func TestVector_DotProduct(t *testing.T) {
	type args struct {
		v1 Vector
		v2 Vector
	}
	tests := []struct {
		name   string
		args   args
		want   int
	}{
		{
			name:   "dot product should be 0",
			args:   args{
				v1:       New(0,0,0),
				v2:       New(0,0,0),
			},
			want:   0,
		},
		{
			name:   "dot product should be -4",
			args:   args{
				v1:       New(3, -5, 4),
				v2:       New(2, 6, 5),
			},
			want:   -4,
		},
		{
			name:   "dot product should be 61",
			args:   args{
				v1:       New(3, 5, 4),
				v2:       New(2, 7, 5),
			},
			want:   61,
		},
		{
			name:   "dot product should be 122",
			args:   args{
				v1:       New(9,2,7),
				v2:       New(4,8,10),
			},
			want:   122,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.v1.DotProduct(tt.args.v2)

			if got != tt.want {
				t.Errorf("DotProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_CrossProduct(t *testing.T) {
	type args struct {
		v1 Vector
		v2 Vector
	}

	tests := []struct {
		name   string
		args   args
		want   Vector
	}{
		{
			name:   "cross product should be 0",
			args:   args{
				v1:       New(0,0,0),
				v2:       New(0,0,0),
			},
			want:   New(0, 0, 0),
		},
		{
			name:   "cross product should be {-49, -7, 28}",
			args:   args{
				v1:       New(3, -5, 4),
				v2:       New(2, 6, 5),
			},
			want:   New(-49,-7,28),
		},
		{
			name:   "cross product should be {-3, 6, -3}",
			args:   args{
				v1:       New(2, 3, 4),
				v2:       New(5, 6,7),
			},
			want:   New(-3, 6, -3),
		},
		{
			name:   "cross product should be {-3, -7, 11}",
			args:   args{
				v1:       New(3, 5, 4),
				v2:       New(2, 7, 5),
			},
			want:   New(-3, -7, 11),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.v1.CrossProduct(tt.args.v2).Format()
			expected := tt.want.Format()

			if got != tt.want.Format() {
				t.Errorf("CrossProduct() = %v, want %v", got, expected)
			}
		})
	}
}

func TestVector_Add(t *testing.T) {
	type args struct {
		v1 Vector
		v2 Vector
	}

	tests := []struct {
		name   string
		args   args
		want   Vector
	}{
		{
			name:   "addition result should be {0, 0, 0}",
			args:   args{
				v1:       New(0,0,0),
				v2:       New(0,0,0),
			},
			want:   New(0, 0, 0),
		},
		{
			name:   "addition result should be {2, 2, 2}",
			args:   args{
				v1:       New(1,1,1),
				v2:       New(1,1,1),
			},
			want:   New(2,2,2),
		},

		{
			name:   "addition result should be {5, 1, 9}",
			args:   args{
				v1:       New(3, -5, 4),
				v2:       New(2, 6, 5),
			},
			want:   New(5,1,9),
		},
		{
			name:   "addition result should be {7, 9, 11}",
			args:   args{
				v1:       New(2, 3, 4),
				v2:       New(5, 6,7),
			},
			want:   New(7, 9, 11),
		},
		{
			name:   "addition result should be {5, 12, 9}",
			args:   args{
				v1:       New(3, 5, 4),
				v2:       New(2, 7, 5),
			},
			want:   New(5, 12, 9),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.v1.Add(tt.args.v2).Format()
			expected := tt.want.Format()

			if got != tt.want.Format() {
				t.Errorf("Add() = %v, want %v", got, expected)
			}
		})
	}
}

func TestVector_Sub(t *testing.T) {
	type args struct {
		v1 Vector
		v2 Vector
	}

	tests := []struct {
		name   string
		args   args
		want   Vector
	}{
		{
			name:   "subtraction result should be 0",
			args:   args{
				v1:       New(0,0,0),
				v2:       New(0,0,0),
			},
			want:   New(0, 0, 0),
		},
		{
			name:   "subtraction result should be {0, 0, 0}",
			args:   args{
				v1:       New(1,1,1),
				v2:       New(1,1,1),
			},
			want:   New(0, 0, 0),
		},
		{
			name:   "subtraction result should be {1, 1, 1}",
			args:   args{
				v1:       New(2,2,2),
				v2:       New(1,1,1),
			},
			want:   New(1, 1, 1),
		},

		{
			name:   "subtraction result should be {1, -11, -1}",
			args:   args{
				v1:       New(3, -5, 4),
				v2:       New(2, 6, 5),
			},
			want:   New(1, -11, -1),
		},
		{
			name:   "subtraction result should be {-3, -3, -3}",
			args:   args{
				v1:       New(2, 3, 4),
				v2:       New(5, 6,7),
			},
			want:   New(-3, -3, -3),
		},
		{
			name:   "subtraction result should be {1, -2, -1}",
			args:   args{
				v1:       New(3, 5, 4),
				v2:       New(2, 7, 5),
			},
			want:   New(1, -2, -1),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.v1.Sub(tt.args.v2).Format()
			expected := tt.want.Format()

			if got != tt.want.Format() {
				t.Errorf("Sub() = %v, want %v", got, expected)
			}
		})
	}
}