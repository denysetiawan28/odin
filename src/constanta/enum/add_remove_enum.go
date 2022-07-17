package enum

type AddRemoveEnum string

func (e AddRemoveEnum) string() string {
	return string(e)
}

var (
	Add     = AddRemoveEnum("add")
	Removed = AddRemoveEnum("removed")
)
