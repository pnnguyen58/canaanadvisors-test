package activities

import (
	"canaanadvisors-test/proto/management"
	"context"
)

func GetMenu(ctx context.Context, input *management.MenuGetRequest) (*management.MenuGetResponse, error) {
	// TODO: implement get menu
	res := []*management.Restaurant{
		{
			Id: 1,
			Name: "Restaurant 1",
			Categories: []*management.Category{
				{
					Id: 1,
				},
			},
		},
	}
	return &management.MenuGetResponse{
		Data: res,
	},  nil
}
