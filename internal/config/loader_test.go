package config

import (
	"reflect"
	"testing"
)

func TestLoad(t *testing.T) {
	type args struct {
		vars map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    *Algolia
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				vars: map[string]string{
					"APP_ID":       "fake",
					"API_KEY":      "fake",
					"SEARCH_INDEX": "fake",
				},
			},
			want: &Algolia{
				AppId:     "fake",
				ApiKey:    "fake",
				IndexName: "fake",
			},
			wantErr: false,
		},
		{
			name: "failure required fields",
			args: args{
				vars: map[string]string{
					"SEARCH_INDEX": "fake",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		if err := SetEnv(tt.args.vars); err != nil {
			t.Errorf("SetEnv() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
		t.Run(tt.name, func(t *testing.T) {
			got, err := Load()
			if (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Load() = %v, want %v", got, tt.want)
			}
		})
		if err := UnsetEnv(tt.args.vars); err != nil {
			t.Errorf("UnsetEnv() error = %v, wantErr %v", err, tt.wantErr)
			return
		}
	}
}
