package keeper

import (
	"newblog/x/newblog/types"
)

var _ types.QueryServer = Keeper{}
