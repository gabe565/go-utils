package versionx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVersion_String(t *testing.T) {
	type fields struct {
		Version string
		Commit  Commit
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "empty", fields: fields{}, want: ""},
		{name: "version only", fields: fields{Version: "1.2.3"}, want: "1.2.3"},
		{name: "commit only", fields: fields{Commit: Commit{SHA: "deadbeefdeadbeef"}}, want: "deadbeef"},
		{
			name:   "version and commit",
			fields: fields{Version: "1.2.3", Commit: Commit{SHA: "deadbeefdeadbeef"}},
			want:   "1.2.3 (deadbeef)",
		},
		{
			name:   "version and modified commit",
			fields: fields{Version: "1.2.3", Commit: Commit{SHA: "deadbeefdeadbeef", Modified: true}},
			want:   "1.2.3 (*deadbeef)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := Version{
				Version: tt.fields.Version,
				Commit:  tt.fields.Commit,
			}
			assert.Equal(t, tt.want, v.String())
		})
	}
}
