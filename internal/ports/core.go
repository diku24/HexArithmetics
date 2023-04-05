package ports

type ArthimeticsPort interface {
	Addtion(a int32, b int32) (int32, error)
	Subtraction(a int32, b int32) (int32, error)
	Multiplication(a int32, b int32) (int32, error)
	Divison(a int32, b int32) (int32, error)
}
