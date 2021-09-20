package search

import (
	"context"
	"reflect"
	"testing"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
)

func TestClient_Search(t *testing.T) {
	type fields struct {
		client *Client
	}
	type args struct {
		ctx   context.Context
		query string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *search.QueryRes
		wantErr bool
	}{
		{
			name: "success",
			fields: fields{
				client: NewClient("RJBON5EHU1", "4e4698fc17de6054c3a81bd26697a205", "pages"),
			},
			args: args{
				ctx:   context.TODO(),
				query: "work",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.fields.client.Search(tt.args.ctx, tt.args.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
