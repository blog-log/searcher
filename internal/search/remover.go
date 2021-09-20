package search

import (
	"context"
)

func (client *Client) Remove(ctx context.Context, id string) error {
	_, err := client.index.DeleteObject(id)
	if err != nil {
		return err
	}

	return nil
}
