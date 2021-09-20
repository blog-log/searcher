package search

import (
	"context"
	"testing"
)

func TestClient_Remove(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		ctx context.Context
		id  string
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
				id:  "fake",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.fields.client.Remove(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Client.Add() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
