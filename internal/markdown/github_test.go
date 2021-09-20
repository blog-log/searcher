package markdown

import (
	"reflect"
	"testing"

	"bloglog.com/search/v1/internal/model"
)

func TestGenerateUrl(t *testing.T) {
	type args struct {
		page *model.Page
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				page: &model.Page{
					Repo:   "https://github.com/fake-owner/fake-name",
					Branch: "fake-branch",
					Path:   "fake/path.md",
				},
			},
			want:    "https://raw.githubusercontent.com/fake-owner/fake-name/fake-branch/fake/path.md",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenerateUrl(tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGet(t *testing.T) {
	type args struct {
		page *model.Page
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				page: &model.Page{
					Repo:   "https://github.com/brandoncate-personal/blog-content-2",
					Branch: "main",
					Path:   "src/testing1.md",
				},
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Get(tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.want && got == "" {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
			if tt.wantErr && err == nil {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}
