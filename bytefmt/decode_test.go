package bytefmt

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestByteDecoder_DecodeBinary(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr require.ErrorAssertionFunc
	}{
		{"999", args{"999"}, 999, require.NoError},
		{"-100", args{"-100"}, -100, require.NoError},
		{"100.1", args{"100.1"}, 100, require.NoError},
		{"12.25KiB", args{"12.25KiB"}, 12544, require.NoError},
		{"12KiB", args{"12KiB"}, 12288, require.NoError},
		{"12Ki", args{"12Ki"}, 12288, require.NoError},
		{"12.25kib", args{"12.25kib"}, 12544, require.NoError},
		{"12kib", args{"12kib"}, 12288, require.NoError},
		{"12ki", args{"12ki"}, 12288, require.NoError},
		{"12.25 KiB", args{"12.25 KiB"}, 12544, require.NoError},
		{"12 KiB", args{"12 KiB"}, 12288, require.NoError},
		{"12 Ki", args{"12 Ki"}, 12288, require.NoError},
		{"2MiB", args{"2MiB"}, 2097152, require.NoError},
		{"2Mi", args{"2Mi"}, 2097152, require.NoError},
		{"6 GiB", args{"6 GiB"}, 6442450944, require.NoError},
		{"6 Gi", args{"6 Gi"}, 6442450944, require.NoError},
		{"6GiB", args{"6GiB"}, 6442450944, require.NoError},
		{"6Gi", args{"6Gi"}, 6442450944, require.NoError},
		{"5TiB", args{"5TiB"}, 5497558138880, require.NoError},
		{"5Ti", args{"5Ti"}, 5497558138880, require.NoError},
		{"5 TiB", args{"5 TiB"}, 5497558138880, require.NoError},
		{"5 Ti", args{"5 Ti"}, 5497558138880, require.NoError},
		{"9PiB", args{"9PiB"}, 10133099161583616, require.NoError},
		{"9Pi", args{"9Pi"}, 10133099161583616, require.NoError},
		{"9 PiB", args{"9 PiB"}, 10133099161583616, require.NoError},
		{"9 Pi", args{"9 Pi"}, 10133099161583616, require.NoError},
		{"empty ", args{""}, 0, require.Error},
		{"no number", args{"MB"}, 0, require.Error},
		{"no number or suffix", args{"test"}, 0, require.Error},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewDecoder().DecodeBinary(tt.args.value)
			tt.wantErr(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestByteDecoder_DecodeDecimal(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr require.ErrorAssertionFunc
	}{
		{"999", args{"999"}, 999, require.NoError},
		{"-100", args{"-100"}, -100, require.NoError},
		{"100.1", args{"100.1"}, 100, require.NoError},
		{"515B", args{"515B"}, 515, require.NoError},
		{"515 B", args{"515 B"}, 515, require.NoError},
		{"12.25KB", args{"12.25KB"}, 12250, require.NoError},
		{"12KB", args{"12KB"}, 12000, require.NoError},
		{"12K", args{"12K"}, 12000, require.NoError},
		{"12.25kb", args{"12.25kb"}, 12250, require.NoError},
		{"12kb", args{"12kb"}, 12000, require.NoError},
		{"12k", args{"12k"}, 12000, require.NoError},
		{"12.25 KB", args{"12.25 KB"}, 12250, require.NoError},
		{"12 KB", args{"12 KB"}, 12000, require.NoError},
		{"12 K", args{"12 K"}, 12000, require.NoError},
		{"2MB", args{"2MB"}, 2000000, require.NoError},
		{"2M", args{"2M"}, 2000000, require.NoError},
		{"6 GB", args{"6 GB"}, 6000000000, require.NoError},
		{"6 G", args{"6 G"}, 6000000000, require.NoError},
		{"6GB", args{"6GB"}, 6000000000, require.NoError},
		{"6G", args{"6G"}, 6000000000, require.NoError},
		{"5TB", args{"5TB"}, 5000000000000, require.NoError},
		{"5T", args{"5T"}, 5000000000000, require.NoError},
		{"5 TB", args{"5 TB"}, 5000000000000, require.NoError},
		{"5 T", args{"5 T"}, 5000000000000, require.NoError},
		{"9PB", args{"9PB"}, 9000000000000000, require.NoError},
		{"9P", args{"9P"}, 9000000000000000, require.NoError},
		{"9 PB", args{"9 PB"}, 9000000000000000, require.NoError},
		{"9 P", args{"9 P"}, 9000000000000000, require.NoError},
		{"8EB", args{"8EB"}, 8000000000000000000, require.NoError},
		{"8E", args{"8E"}, 8000000000000000000, require.NoError},
		{"8 EB", args{"8 EB"}, 8000000000000000000, require.NoError},
		{"8 E", args{"8 E"}, 8000000000000000000, require.NoError},
		{"empty ", args{""}, 0, require.Error},
		{"no number", args{"MB"}, 0, require.Error},
		{"no number or suffix", args{"test"}, 0, require.Error},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewDecoder().DecodeDecimal(tt.args.value)
			tt.wantErr(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_splitBytes(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name       string
		args       args
		wantBytes  float64
		wantSuffix string
		wantErr    require.ErrorAssertionFunc
	}{
		{"empty", args{""}, 0, "", require.Error},
		{"1 MiB", args{"1 MiB"}, 1, "MiB", require.NoError},
		{"1MiB", args{"1MiB"}, 1, "MiB", require.NoError},
		{"1 MB", args{"1 MB"}, 1, "MB", require.NoError},
		{"1MB", args{"1MB"}, 1, "MB", require.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBytes, gotSuffix, err := split(tt.args.val)
			tt.wantErr(t, err)
			if tt.wantBytes != 0 {
				assert.InEpsilon(t, tt.wantBytes, gotBytes, 0.01)
			} else {
				assert.Zero(t, gotBytes)
			}
			assert.Equal(t, tt.wantSuffix, gotSuffix)
		})
	}
}

func TestByteDecoder_Decode(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr require.ErrorAssertionFunc
	}{
		{"1 MiB", args{"1 MiB"}, 1048576, require.NoError},
		{"1MiB", args{"1MiB"}, 1048576, require.NoError},
		{"1 MB", args{"1 MB"}, 1000000, require.NoError},
		{"1MB", args{"1MB"}, 1000000, require.NoError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewDecoder().Decode(tt.args.text)
			tt.wantErr(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func BenchmarkByteDecoder_Decode(b *testing.B) {
	b.Run("binary input", func(b *testing.B) {
		d := NewDecoder()
		for i := range b.N {
			_, _ = d.Decode(strconv.Itoa(i) + " KiB")
		}
	})

	b.Run("decimal input", func(b *testing.B) {
		d := NewDecoder()
		for i := range b.N {
			_, _ = d.Decode(strconv.Itoa(i) + " KB")
		}
	})
}

func BenchmarkByteDecoder_DecodeBinary(b *testing.B) {
	d := NewDecoder()
	for i := range b.N {
		_, _ = d.DecodeBinary(strconv.Itoa(i) + " KiB")
	}
}

func BenchmarkByteDecoder_DecodeDecimal(b *testing.B) {
	d := NewDecoder()
	for i := range b.N {
		_, _ = d.DecodeDecimal(strconv.Itoa(i) + " KB")
	}
}
