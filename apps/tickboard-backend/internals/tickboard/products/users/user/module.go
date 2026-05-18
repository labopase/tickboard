package user

import (
	v1GetUser "github.com/halimdotnet/tickboard-backend/internals/tickboard/products/users/user/features/get_user/v1"
	"go.uber.org/fx"
)

var Module = fx.Module("user_data",
	fx.Provide(v1GetUser.NewGetUserEndpoint),
	fx.Invoke(func(lc fx.Lifecycle,
		getUser v1GetUser.GetUserEndpoint,
	) {
		getUser.RegisterRoute()
	}),
)
