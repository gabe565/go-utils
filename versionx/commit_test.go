package versionx

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommit_Long(t *testing.T) {
	type fields struct {
		SHA      string
		Modified bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "clean", fields: fields{SHA: "deadbeefdeadbeef"}, want: "deadbeefdeadbeef"},
		{name: "modified", fields: fields{SHA: "deadbeefdeadbeef", Modified: true}, want: "*deadbeefdeadbeef"},
		{name: "empty", fields: fields{}, want: ""},
		{name: "empty modified", fields: fields{Modified: true}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Commit{
				SHA:      tt.fields.SHA,
				Modified: tt.fields.Modified,
			}
			assert.Equal(t, tt.want, c.Long())
		})
	}
}

func TestCommit_Short(t *testing.T) {
	type fields struct {
		SHA      string
		Modified bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "clean", fields: fields{SHA: "deadbeefdeadbeef"}, want: "deadbeef"},
		{name: "modified", fields: fields{SHA: "deadbeefdeadbeef", Modified: true}, want: "*deadbeef"},
		{name: "empty", fields: fields{}, want: ""},
		{name: "empty modified", fields: fields{Modified: true}, want: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Commit{
				SHA:      tt.fields.SHA,
				Modified: tt.fields.Modified,
			}
			assert.Equal(t, tt.want, c.Short())
		})
	}
}
