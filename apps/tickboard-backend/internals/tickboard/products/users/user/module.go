package user

import (
	v1GetUser "github.com/halimdotnet/tickboard-backend/internals/tickboard/products/users/user/features/get_user/v1"
	"go.uber.org/fx"
)

var Module = fx.Options(
	v1GetUser.Module,
)
