package bytefmt

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestByteEncoder_EncodeBinary(t *testing.T) {
	type args struct {
		value int64
	}

	tests := []struct {
		name    string
		args    args
		encoder *Encoder
		want    string
	}{
		{"0", args{value: 0}, nil, "0"},
		{"515 B", args{value: 515}, nil, "515 B"},
		{"30.59 KiB", args{value: 31323}, nil, "30.59 KiB"},
		{"12.62 MiB", args{value: 13231323}, nil, "12.62 MiB"},
		{"6.82 GiB", args{value: 7323232398}, nil, "6.82 GiB"},
		{"6.66 TiB", args{value: 7323232398434}, nil, "6.66 TiB"},
		{"8.81 PiB", args{value: 9923232398434432}, nil, "8.81 PiB"},
		{"8.00 EiB", args{value: math.MaxInt64}, nil, "8 EiB"},
		{"no space", args{value: 31323}, NewEncoder().SetUseSpace(false), "30.59KiB"},
		{"decimal not trimmed", args{value: 31323}, NewEncoder().SetTrimIntDecimal(true), "30.59 KiB"},
		{"int decimal trimmed", args{value: 1024}, NewEncoder().SetTrimIntDecimal(true), "1 KiB"},
		{
			"int decimal trimmed without space",
			args{value: 1024},
			NewEncoder().SetUseSpace(false).SetTrimIntDecimal(true),
			"1KiB",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := tt.encoder
			if e == nil {
				e = NewEncoder()
			}
			assert.Equal(t, tt.want, e.EncodeBinary(tt.args.value))
		})
	}
}

func TestByteEncoder_EncodeDecimal(t *testing.T) {
	type args struct {
		value int64
	}
	tests := []struct {
		name    string
		args    args
		encoder *Encoder
		want    string
	}{
		{"0", args{value: 0}, nil, "0"},
		{"515 B", args{value: 515}, nil, "515 B"},
		{"31.32 KB", args{value: 31323}, nil, "31.32 KB"},
		{"13.23 MB", args{value: 13231323}, nil, "13.23 MB"},
		{"7.32 GB", args{value: 7323232398}, nil, "7.32 GB"},
		{"7.32 TB", args{value: 7323232398434}, nil, "7.32 TB"},
		{"9.92 PB", args{value: 9923232398434432}, nil, "9.92 PB"},
		{"9.92 EB", args{value: math.MaxInt64}, nil, "9.22 EB"},
		{"no space", args{value: 31323}, NewEncoder().SetUseSpace(false), "31.32KB"},
		{"decimal not trimmed", args{value: 31323}, NewEncoder().SetTrimIntDecimal(true), "31.32 KB"},
		{"int decimal trimmed", args{value: 1000}, NewEncoder().SetTrimIntDecimal(true), "1 KB"},
		{
			"int decimal trimmed without space",
			args{value: 1000},
			NewEncoder().SetUseSpace(false).SetTrimIntDecimal(true),
			"1KB",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := tt.encoder
			if e == nil {
				e = NewEncoder()
			}
			assert.Equal(t, tt.want, e.EncodeDecimal(tt.args.value))
		})
	}
}

func BenchmarkByteEncoder_EncodeBinary(b *testing.B) {
	e := NewEncoder()
	for i := range b.N {
		_ = e.EncodeBinary(int64(i * i))
	}
}

func BenchmarkByteEncoder_EncodeDecimal(b *testing.B) {
	e := NewEncoder()
	for i := range b.N {
		_ = e.EncodeDecimal(int64(i * i))
	}
}
