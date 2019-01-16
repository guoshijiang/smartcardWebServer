package types

type SqlBase interface {
	Query()
	QueryByCondition(condition string)
	Delete(condition string)
	Clear()
	Update()
	Insert()
}
