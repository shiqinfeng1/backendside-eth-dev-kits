package common

//import "github.com/go-playground/validator"
import "gopkg.in/go-playground/validator.v9"

//SimpleValidator 参数验证器
type SimpleValidator struct {
	Validator *validator.Validate
}

//Validate 实现Validate接口
func (cv *SimpleValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
