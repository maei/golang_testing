package service

var HelperService HelperServiceInterface = NewHelperService()

type HelperServiceInterface interface {
	DoSomething(string) bool
}

type helperService struct{}

func NewHelperService() HelperServiceInterface {
	return &helperService{}
}

func (*helperService) DoSomething(data string) bool {
	if data == "No Good" {
		return true
	}
	return false

}
