package reflectx

import (
	"github.com/frankee/sqlx/types"
)

type VirtualField interface {
	types.Extension
	SetDelegateField(field interface{})
}
