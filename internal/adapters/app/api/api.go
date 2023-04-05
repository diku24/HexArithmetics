package api

import "HexArithmatics/internal/ports"

type Adapter struct {
	db    ports.DBPort
	arith ports.ArthimeticsPort
}

func NewAdapter(db ports.DBPort, arith ports.ArthimeticsPort) *Adapter {
	return &Adapter{arith: arith, db: db}

}

func (apia Adapter) GetAddtion(a, b int32) (int32, error) {
	answer, err := apia.arith.Addtion(a, b)
	if err != nil {
		return 0, err
	}
	err = apia.db.AddToHistory(answer, "addtion")
	if err != nil {
		return 0, err
	}
	return answer, nil
}

func (apia Adapter) GetSubstraction(a, b int32) (int32, error) {
	answer, err := apia.arith.Subtraction(a, b)
	if err != nil {
		return 0, err
	}
	err = apia.db.AddToHistory(answer, "Substraction")
	if err != nil {
		return 0, err
	}
	return answer, nil
}

func (apia Adapter) GetMultiplication(a, b int32) (int32, error) {
	answer, err := apia.arith.Multiplication(a, b)
	if err != nil {
		return 0, err
	}
	err = apia.db.AddToHistory(answer, "multiplication")
	if err != nil {
		return 0, err
	}
	return answer, nil
}

func (apia Adapter) GetDivision(a, b int32) (int32, error) {
	answer, err := apia.arith.Divison(a, b)
	if err != nil {
		return 0, err
	}
	err = apia.db.AddToHistory(answer, "division")
	if err != nil {
		return 0, err
	}
	return answer, nil
}
