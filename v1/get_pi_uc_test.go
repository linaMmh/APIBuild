package v1

import (
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGetPi_DeletePi(t *testing.T) {
	type fields struct {
		keepPiInterface    KeepPiInterface
		maxRandomPrecision int
		redisEnabled       bool
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &GetPi{
				keepPiInterface:    tt.fields.keepPiInterface,
				maxRandomPrecision: tt.fields.maxRandomPrecision,
				redisEnabled:       tt.fields.redisEnabled,
			}
			uc.DeletePi(tt.args.c)
		})
	}
}

func TestGetPi_GetPi(t *testing.T) {
	type fields struct {
		keepPiInterface    KeepPiInterface
		maxRandomPrecision int
		redisEnabled       bool
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &GetPi{
				keepPiInterface:    tt.fields.keepPiInterface,
				maxRandomPrecision: tt.fields.maxRandomPrecision,
				redisEnabled:       tt.fields.redisEnabled,
			}
			uc.GetPi(tt.args.c)
		})
	}
}

func TestGetPi_GetPiRandom(t *testing.T) {
	type fields struct {
		keepPiInterface    KeepPiInterface
		maxRandomPrecision int
		redisEnabled       bool
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := &GetPi{
				keepPiInterface:    tt.fields.keepPiInterface,
				maxRandomPrecision: tt.fields.maxRandomPrecision,
				redisEnabled:       tt.fields.redisEnabled,
			}
			uc.GetPiRandom(tt.args.c)
		})
	}
}

func TestNewGetPi(t *testing.T) {
	type args struct {
		keepPiInterface    KeepPiInterface
		maxRandomPrecision int
		redisEnabled       bool
	}
	tests := []struct {
		name string
		args args
		want *GetPi
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGetPi(tt.args.keepPiInterface, tt.args.maxRandomPrecision, tt.args.redisEnabled); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGetPi() = %v, want %v", got, tt.want)
			}
		})
	}
}
