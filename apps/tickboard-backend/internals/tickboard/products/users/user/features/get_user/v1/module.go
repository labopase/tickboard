package v1

import "go.uber.org/fx"

var Module = fx.Module("get_user_v1",
	fx.Provide(NewGetUserEndpoint),
	fx.Invoke(func(r GetUserEndpoint) {
		r.RegisterRoute()
	}),
)
