package ports

type APIPort interface {
	GetAddtion(a, b int32) (int32, error)
	GetSubstraction(a, b int32) (int32, error)
	GetMultiplication(a, b int32) (int32, error)
	GetDivision(a, b int32) (int32, error)
}
