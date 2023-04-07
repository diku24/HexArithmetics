package api

import "HexArithmatics/internal/ports"

//Application Implements the APIPort Interface
type Application struct {
	db    ports.DBPort
	arith Arithmetic
}

//NewApplication creates a new applications
func NewApplication(db ports.DBPort, arith Arithmetic) *Application {
	return &Application{db: db, arith: arith}
}

//GetAddition gets the result of addting parameters a and b
func (apia Application) GetAddtion(a, b int32) (int32, error) {
	answer, err := apia.arith.Addition(a, b)
	if err != nil {
		return 0, err
	}
	err = apia.db.AddToHistory(answer, "addtion")
	if err != nil {
		return 0, err
	}
	return answer, nil
}

//GetSubstraction gets the result of subtracting parameters a and b
func (apia Application) GetSubstraction(a, b int32) (int32, error) {
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

//GetMultiplication gets the result of subtracting parameters a and b
func (apia Application) GetMultiplication(a, b int32) (int32, error) {
	answer, err := apia.arith.Multiplication(a, b)
	if err != nil {
		return 0, err
	}
	err = apia.db.AddToHistory(answer, "Multiplication")
	if err != nil {
		return 0, err
	}
	return answer, nil
}

//GetDivision gets the result of subtracting parameters a and b
func (apia Application) GetDivision(a, b int32) (int32, error) {
	answer, err := apia.arith.Division(a, b)
	if err != nil {
		return 0, err
	}
	err = apia.db.AddToHistory(answer, "Division")
	if err != nil {
		return 0, err
	}
	return answer, nil
}
