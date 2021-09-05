package sgomock

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestTRun(t *testing.T) {
	for i := 0; i < 5; i++ {
		name := fmt.Sprintf("test number = %d", i)
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			fmt.Println(name)
		})
	}
	t.Error()
}

func TestCalculatorImpl_Calc_GlobalMock(t *testing.T) {
	ctrl := gomock.NewController(t)
	mock := NewMockArithmetic(ctrl)
	mock.EXPECT().Add(gomock.Any(), gomock.Any()).Return(int64(2)).Times(2)
	mock.EXPECT().Sub(gomock.Any(), gomock.Any()).Return(int64(-8)).Times(2)

	type fields struct {
		ArithmeticSerivce Arithmetic
	}
	type args struct {
		x int64
		y int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int64
	}{
		{
			name:   "test case 1",
			fields: fields{ArithmeticSerivce: mock},
			args:   args{x: 1, y: 1},
			want:   -8,
		},
		{
			name:   "test case 2",
			fields: fields{ArithmeticSerivce: mock},
			args:   args{x: 1, y: 1},
			want:   -8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &CalculatorImpl{ArithmeticService: tt.fields.ArithmeticSerivce}
			if got := a.Calc(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("CalculatorImpl.Calc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculatorImpl_Calc_GlobalController(t *testing.T) {
	ctrl := gomock.NewController(t)
	type fields struct {
		ArithmeticSerivce Arithmetic
	}
	type args struct {
		x int64
		y int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int64
	}{
		{
			name: "test case 1",
			fields: fields{
				ArithmeticSerivce: func() Arithmetic {
					mock := NewMockArithmetic(ctrl)
					mock.EXPECT().Add(gomock.Any(), gomock.Any()).Return(int64(2))
					mock.EXPECT().Sub(gomock.Any(), gomock.Any()).Return(int64(-8))
					return mock
				}(),
			},
			args: args{x: 1, y: 1},
			want: -8,
		},
		{
			name: "test case 2",
			fields: fields{
				ArithmeticSerivce: func() Arithmetic {
					mock := NewMockArithmetic(ctrl)
					mock.EXPECT().Add(gomock.Any(), gomock.Any()).Return(int64(2))
					mock.EXPECT().Sub(gomock.Any(), gomock.Any()).Return(int64(-8))
					return mock
				}(),
			},
			args: args{x: 1, y: 1},
			want: -8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &CalculatorImpl{
				ArithmeticService: tt.fields.ArithmeticSerivce,
			}
			if got := a.Calc(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("CalculatorImpl.Calc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculatorImpl_Calc(t *testing.T) {
	type fields struct {
		ArithmeticSerivce func(*gomock.Controller) Arithmetic
	}
	type args struct {
		x int64
		y int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int64
	}{
		{
			name: "test case 1",
			fields: fields{
				ArithmeticSerivce: func(ctrl *gomock.Controller) Arithmetic {
					mock := NewMockArithmetic(ctrl)
					mock.EXPECT().Add(gomock.Any(), gomock.Any()).Return(int64(2))
					mock.EXPECT().Sub(gomock.Any(), gomock.Any()).Return(int64(-8))
					return mock
				},
			},
			args: args{x: 1, y: 1},
			want: -8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			a := &CalculatorImpl{
				ArithmeticService: tt.fields.ArithmeticSerivce(ctrl),
			}
			if got := a.Calc(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("CalculatorImpl.Calc() = %v, want %v", got, tt.want)
			}
		})
	}
}
