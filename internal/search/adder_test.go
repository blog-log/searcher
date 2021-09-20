package search

import (
	"context"
	"testing"

	"bloglog.com/search/v1/internal/model"
)

func TestClient_Add(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		ctx  context.Context
		page *model.Page
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				client: NewClient("", "", ""),
			},
			args: args{
				ctx: context.TODO(),
				page: &model.Page{
					Id:     "fake",
					Repo:   "fake",
					Branch: "fake",
					Path:   "fake",
					Title:  "fake",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.fields.client.Add(tt.args.ctx, tt.args.page); (err != nil) != tt.wantErr {
				t.Errorf("Client.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
