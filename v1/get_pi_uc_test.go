package v1

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"pi-api/common"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetPi_DeletePi(t *testing.T) {
	type fields struct {
		keepPiInterface    *KeepPiMock
		maxRandomPrecision int
		redisEnabled       bool
	}

	commonFields := fields{
		keepPiInterface:    &KeepPiMock{},
		maxRandomPrecision: 200,
		redisEnabled:       true,
	}

	tests := []struct {
		name   string
		fields fields
		code   int
		random int
		mock   func(f fields)
	}{
		{
			name:   "delete pi success",
			fields: commonFields,
			code:   http.StatusOK,
			random: 5,
			mock: func(f fields) {
				f.keepPiInterface.On("getPi",
					fmt.Sprintf(IndexRedis, 5)).Return(Generalresp, nil).Once()

				f.keepPiInterface.On("deletePi",
					fmt.Sprintf(IndexRedis, 5)).Return(nil).Once()
			},
		},
		{
			name:   "error not found",
			fields: commonFields,
			code:   http.StatusConflict,
			random: 5,
			mock: func(f fields) {
				f.keepPiInterface.On("getPi",
					fmt.Sprintf(IndexRedis, 5)).Return(common.Response{}, errors.New("error consulting redis")).Once()
			},
		},
		{
			name:   "error with data",
			fields: commonFields,
			code:   http.StatusBadRequest,
			random: 0,
			mock: func(f fields) {
			},
		},
	}
	for _, tt := range tests {
		tt.mock(tt.fields)
		t.Run(tt.name, func(t *testing.T) {
			resp := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			c, _ := gin.CreateTestContext(resp)
			c.Request = httptest.NewRequest("DELETE", fmt.Sprintf("/deletePi?random_number=%d", tt.random), nil)

			uc := &GetPi{
				keepPiInterface:    tt.fields.keepPiInterface,
				maxRandomPrecision: tt.fields.maxRandomPrecision,
				redisEnabled:       tt.fields.redisEnabled,
			}
			uc.DeletePi(c)
			if !reflect.DeepEqual(resp.Result().StatusCode, tt.code) {
				t.Errorf("getPi() got = %v, want %v", resp.Result().StatusCode, tt.code)
			}
		})
	}
}

func TestGetPi_GetPi(t *testing.T) {
	type fields struct {
		keepPiInterface    *KeepPiMock
		maxRandomPrecision int
		redisEnabled       bool
	}

	commonFields := fields{
		keepPiInterface:    &KeepPiMock{},
		maxRandomPrecision: 200,
		redisEnabled:       true,
	}

	tests := []struct {
		name   string
		fields fields
		code   int
		random int
		mock   func(f fields)
	}{
		{
			name:   "get pi success",
			fields: commonFields,
			code:   http.StatusOK,
			random: 5,
			mock: func(f fields) {
				f.keepPiInterface.On("getPi",
					fmt.Sprintf(IndexRedis, 5)).Return(Generalresp, nil).Once()

				f.keepPiInterface.On("setPi",
					fmt.Sprintf(IndexRedis, 5), Generalresp).Return(nil).Once()
			},
		},
		{
			name:   "get pi error data",
			fields: commonFields,
			random: 0,
			code:   http.StatusBadRequest,
			mock: func(f fields) {
			},
		},
		{
			name:   "get pi error redis",
			fields: commonFields,
			random: 5,
			code:   http.StatusConflict,
			mock: func(f fields) {
				f.keepPiInterface.On("getPi",
					fmt.Sprintf(IndexRedis, 5)).Return(common.Response{}, errors.New("error consulting redis")).Once()
			},
		},
	}
	for _, tt := range tests {
		tt.mock(tt.fields)
		t.Run(tt.name, func(t *testing.T) {
			resp := httptest.NewRecorder()
			gin.SetMode(gin.TestMode)
			c, _ := gin.CreateTestContext(resp)
			c.Request = httptest.NewRequest("GET", fmt.Sprintf("/getPi?random_number=%d", tt.random), nil)
			uc := &GetPi{
				keepPiInterface:    tt.fields.keepPiInterface,
				maxRandomPrecision: tt.fields.maxRandomPrecision,
				redisEnabled:       tt.fields.redisEnabled,
			}

			uc.GetPi(c)
			if !reflect.DeepEqual(resp.Result().StatusCode, tt.code) {
				t.Errorf("getPi() got = %v, want %v", resp.Result().StatusCode, tt.code)
			}
		})

	}
}

func TestNewGetPi(t *testing.T) {
	type args struct {
		keepPiInterface    *KeepPiMock
		maxRandomPrecision int
		redisEnabled       bool
	}
	tests := []struct {
		name string
		args args
		want func(a args) *GetPi
	}{
		{
			name: "success",
			args: args{
				keepPiInterface:    &KeepPiMock{},
				maxRandomPrecision: 200,
				redisEnabled:       true,
			},
			want: func(a args) *GetPi {
				return NewGetPi(
					a.keepPiInterface,
					a.maxRandomPrecision,
					a.redisEnabled,
				)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGetPi(tt.args.keepPiInterface, tt.args.maxRandomPrecision, tt.args.redisEnabled); !reflect.DeepEqual(got, tt.want(tt.args)) {
				t.Errorf("NewGetPi() = %v, want %v", got, tt.want(tt.args))
			}
		})
	}
}
