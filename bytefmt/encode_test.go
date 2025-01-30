package bytefmt

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestByteEncoder_EncodeBinary(t *testing.T) {
	type args struct {
		value     int64
		precision int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"0", args{value: 0, precision: 2}, "0"},
		{"515 B", args{value: 515, precision: 2}, "515 B"},
		{"30.59 KiB", args{value: 31323, precision: 2}, "30.59 KiB"},
		{"12.62 MiB", args{value: 13231323, precision: 2}, "12.62 MiB"},
		{"6.82 GiB", args{value: 7323232398, precision: 2}, "6.82 GiB"},
		{"6.66 TiB", args{value: 7323232398434, precision: 2}, "6.66 TiB"},
		{"8.81 PiB", args{value: 9923232398434432, precision: 2}, "8.81 PiB"},
		{"8.00 EiB", args{value: math.MaxInt64, precision: 2}, "8 EiB"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewEncoder().EncodeBinary(tt.args.value))
		})
	}
}

func TestByteEncoder_EncodeDecimal(t *testing.T) {
	type args struct {
		value     int64
		precision int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"0", args{value: 0, precision: 2}, "0"},
		{"515 B", args{value: 515, precision: 2}, "515 B"},
		{"31.32 KB", args{value: 31323, precision: 2}, "31.32 KB"},
		{"13.23 MB", args{value: 13231323, precision: 2}, "13.23 MB"},
		{"7.32 GB", args{value: 7323232398, precision: 2}, "7.32 GB"},
		{"7.32 TB", args{value: 7323232398434, precision: 2}, "7.32 TB"},
		{"9.92 PB", args{value: 9923232398434432, precision: 2}, "9.92 PB"},
		{"9.92 EB", args{value: math.MaxInt64, precision: 2}, "9.22 EB"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, NewEncoder().EncodeDecimal(tt.args.value))
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
