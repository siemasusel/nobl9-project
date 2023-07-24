package stddevapi

import (
	"context"
	"stddevapi/mock"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCalculateStdDev(t *testing.T) {
	type args struct {
		requests int
		length   int
	}

	type want struct {
		result StdDevGroupResult
		isErr  func(error) bool
	}

	type fields struct {
		mockRandomBackend func(*mock.MockRandomBackend)
	}

	cases := []struct {
		name   string
		args   args
		want   want
		fields fields
	}{
		{
			name: "get sequence integers",
			args: args{
				requests: 2,
				length:   5,
			},
			want: want{
				result: StdDevGroupResult{
					Results: []StdDevResult{
						{
							StdDev: 1.4142135623730951,
							Data:   []int{1, 2, 3, 4, 5},
						},
						{
							StdDev: 1.4142135623730951,
							Data:   []int{1, 2, 3, 4, 5},
						},
					},
					ResultSum: StdDevResult{
						StdDev: 1.4142135623730951,
						Data:   []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5},
					},
				},
			},
			fields: fields{
				mockRandomBackend: func(m *mock.MockRandomBackend) {
					m.EXPECT().GetRandomIntegers(gomock.Any(), 5).Return([]int{1, 2, 3, 4, 5}, nil).Times(2)
				},
			},
		},
		{
			name: "get random integers",
			args: args{
				requests: 2,
				length:   4,
			},
			want: want{
				result: StdDevGroupResult{
					Results: []StdDevResult{
						{
							StdDev: 266176.81167729,
							Data:   []int{758071, 606843, 43914, 455915},
						},
						{
							StdDev: 255747.16872430866,
							Data:   []int{567707, 308842, 940701, 888198},
						},
					},
					ResultSum: StdDevResult{
						StdDev: 281374.9610548341,
						Data: []int{
							758071,
							606843,
							43914,
							455915,
							567707,
							308842,
							940701,
							888198,
						},
					},
				},
			},
			fields: fields{
				mockRandomBackend: func(m *mock.MockRandomBackend) {
					m.EXPECT().GetRandomIntegers(gomock.Any(), 4).Return([]int{758071, 606843, 43914, 455915}, nil).Times(1)
					m.EXPECT().GetRandomIntegers(gomock.Any(), 4).Return([]int{567707, 308842, 940701, 888198}, nil).Times(1)
				},
			},
		},
		{
			name: "validation error from random backend",
			args: args{
				requests: 2,
				length:   4,
			},
			want: want{
				isErr: IsValidationError,
			},
			fields: fields{
				mockRandomBackend: func(m *mock.MockRandomBackend) {
					m.EXPECT().GetRandomIntegers(gomock.Any(), 4).Return(nil, NewValidationError(nil, "test error")).Times(2)
				},
			},
		},
		{
			name: "negative requests",
			args: args{
				requests: -5,
				length:   4,
			},
			want: want{
				isErr: IsValidationError,
			},
		},
		{
			name: "negative length",
			args: args{
				requests: 2,
				length:   -5,
			},
			want: want{
				isErr: IsValidationError,
			},
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			m := mock.NewMockRandomBackend(gomock.NewController(t))
			if tt.fields.mockRandomBackend != nil {
				tt.fields.mockRandomBackend(m)
			}
			svc := Service{randBackend: m}

			res, err := svc.CalculateStdDev(context.Background(), tt.args.requests, tt.args.length)
			if tt.want.isErr == nil {
				require.NoError(t, err)
				assert.Equal(t, tt.want.result, res)
			}
			if tt.want.isErr != nil && !tt.want.isErr(err) {
				t.Errorf("got wrong err: %v ", err)
			}
		})
	}

	mock.NewMockRandomBackend(gomock.NewController(t))
}

func TestCalcStdDev(t *testing.T) {
	cases := []struct {
		name  string
		input []int
		res   float64
	}{
		{
			name:  "stddev for 10 integers",
			input: []int{3, 5, 9, 1, 8, 6, 58, 9, 4, 10},
			res:   15.8117045254457,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			actual := calcStdDev(tt.input)
			assert.Equal(t, tt.res, actual)
		})
	}
}
