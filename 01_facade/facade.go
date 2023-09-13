package facade

import "fmt"

func NewAPI() API {
	return &apiImpl{
		a: NewAModuleAPI(),
		b: NewBModuleAPI(),
	}
}

type API interface {
	Test() string
}

type AModuleAPI interface {
	TestA() string
}

type BModuleAPI interface {
	TestB() string
}

type apiImpl struct {
	a AModuleAPI
	b BModuleAPI
}

func (a *apiImpl) Test() string {
	aRet := a.a.TestA()
	bRet := a.b.TestB()
	return fmt.Sprintf("%s\n%s", aRet, bRet)
}

func NewAModuleAPI() AModuleAPI {
	return &aModuleImpl{}
}

func NewBModuleAPI() BModuleAPI {
	return &bModuleImpl{}
}

type aModuleImpl struct{}

type bModuleImpl struct{}

func (*aModuleImpl) TestA() string {
	return "A module running"
}

func (*bModuleImpl) TestB() string {
	return "B module running"
}
