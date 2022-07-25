package v1

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"test3/common"
	"testing"
)

var resp = common.Response{
	Param:  5,
	Random: 5,
	PiCalc: "3.14",
}

func TestKeepPi_deletePi(t *testing.T) {
	db, mock := common.NewCreateConnectionMock()
	type fields struct {
		ClientDB common.ConnectDBMock
	}
	type args struct {
		index string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		mock    func(f fields, a args)
	}{
		{
			name:    "delete success",
			fields:  fields{ClientDB: db},
			args:    args{index: fmt.Sprintf(IndexRedis, 5)},
			wantErr: false,
			mock: func(f fields, a args) {
				mock.ExpectDel(a.index).SetVal(1)
			},
		},
		{
			name:    "delete fail",
			fields:  fields{ClientDB: db},
			args:    args{index: fmt.Sprintf(IndexRedis, 5)},
			wantErr: true,
			mock: func(f fields, a args) {
				mock.ExpectDel(a.index).SetErr(errors.New("fail connect redis"))
			},
		},
	}
	for _, tt := range tests {
		tt.mock(tt.fields, tt.args)
		t.Run(tt.name, func(t *testing.T) {
			r := &KeepPi{
				ClientDB: db.GetClient(),
			}
			if err := r.deletePi(tt.args.index); (err != nil) != tt.wantErr {
				t.Errorf("deletePi() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestKeepPi_getPi(t *testing.T) {
	db, mock := common.NewCreateConnectionMock()
	type fields struct {
		ClientDB common.ConnectDBMock
	}
	type args struct {
		index string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    common.Response
		wantErr bool
		mock    func(f fields, a args)
	}{
		{
			name: "get pi success",
			fields: fields{
				ClientDB: db,
			},
			args: args{
				index: fmt.Sprintf(IndexRedis, 5),
			},
			want:    resp,
			wantErr: false,
			mock: func(f fields, a args) {
				mock.ExpectHGet(a.index, CacheField).
					SetVal("{\"param\":5,\"random\":5,\"PiCalc\":\"3.14\"}")
			},
		},
		{
			name: "get pi error",
			fields: fields{
				ClientDB: db,
			},
			args: args{
				index: fmt.Sprintf(IndexRedis, 5),
			},
			wantErr: true,
			mock: func(f fields, a args) {
				mock.ExpectHGet(a.index, CacheField).SetErr(errors.New("fail connect redis"))
			},
		},
	}
	for _, tt := range tests {
		tt.mock(tt.fields, tt.args)
		t.Run(tt.name, func(t *testing.T) {
			r := &KeepPi{
				ClientDB: db.GetClient(),
			}
			got, err := r.getPi(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("getPi() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getPi() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeepPi_setPi(t *testing.T) {
	db, mock := common.NewCreateConnectionMock()
	type fields struct {
		ClientDB common.ConnectDBMock
	}
	type args struct {
		index    string
		response common.Response
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		mock    func(f fields, a args)
	}{
		{
			name: "success set Pi",
			fields: fields{
				ClientDB: db,
			},
			args: args{
				index:    fmt.Sprintf(IndexRedis, 5),
				response: resp,
			},
			wantErr: false,
			mock: func(f fields, a args) {
				jsonSend, _ := json.Marshal(a.response)
				mock.Regexp().ExpectHSet(fmt.Sprintf(IndexRedis, 5), CacheField, jsonSend).SetVal(1)
			},
		},
		{
			name: "fail set Pi",
			fields: fields{
				ClientDB: db,
			},
			args: args{
				index:    fmt.Sprintf(IndexRedis, 5),
				response: resp,
			},
			wantErr: true,
			mock: func(f fields, a args) {
				jsonSend, _ := json.Marshal(a.response)
				mock.Regexp().ExpectHSet(fmt.Sprintf(IndexRedis, 5), CacheField, jsonSend).SetErr(errors.New("fail connect redis"))
			},
		},
	}
	for _, tt := range tests {
		tt.mock(tt.fields, tt.args)
		t.Run(tt.name, func(t *testing.T) {
			r := &KeepPi{
				ClientDB: db.GetClient(),
			}
			if err := r.setPi(tt.args.index, tt.args.response); (err != nil) != tt.wantErr {
				t.Errorf("setPi() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewKeepPiRepository(t *testing.T) {
	db, _ := common.NewCreateConnectionMock()
	redisRepository := NewKeepPiRepository(db.GetClient())

	type args struct {
		client common.ConnectDBMock
	}
	tests := []struct {
		name string
		args args
		want *KeepPi
	}{
		{
			name: "Success",
			args: args{
				client: db,
			},
			want: redisRepository,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewKeepPiRepository(tt.args.client.GetClient()); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewKeepPiRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}
