package v1

import (
	"pi-api/common"

	"github.com/stretchr/testify/mock"
)

type KeepPiMock struct {
	mock.Mock
}

func (mock *KeepPiMock) setPi(index string, response common.Response) error {
	args := mock.Called(index, response)
	return args.Error(0)
}

func (mock *KeepPiMock) getPi(index string) (common.Response, error) {
	args := mock.Called(index)
	return args.Get(0).(common.Response), args.Error(1)
}

func (mock *KeepPiMock) deletePi(index string) error {
	args := mock.Called(index)
	return args.Error(0)
}
