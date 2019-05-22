package helper

import (
	"testing"

	"github.com/irisnet/irishub-sync/module/logger"
)

func TestGetValidator(t *testing.T) {
	type args struct {
		valAddr string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test get validator",
			args: args{
				valAddr: "faa15lpdxlk0hwkewmncdhlyfle8jc3k9xzhh75txs",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := GetValidator(tt.args.valAddr)
			if err != nil {
				logger.Error.Fatalln(err)
			}
			logger.Info.Println(ToJson(res))
		})
	}
}

func TestGetDelegation(t *testing.T) {
	type args struct {
		delAddr string
		valAddr string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test get delegation",
			args: args{
				delAddr: "faa15p4n0uqafr7udgw59g3fq3dwj2kdww5q6p4znd",
				valAddr: "faa15lpdxlk0hwkewmncdhlyfle8jc3k9xzhh75txs",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := GetDelegation(tt.args.delAddr, tt.args.valAddr)
			if err != nil {
				logger.Error.Fatalln(err)
			}
			logger.Info.Println(ToJson(res))
		})
	}
}

func TestGetUnbondingDelegation(t *testing.T) {
	type args struct {
		delAddr string
		valAddr string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test get unbonding delegation",
			args: args{
				delAddr: "faa19tyxwyj7y2sld8qy2m2wgv7cekfep229schqnn",
				valAddr: "faa15lpdxlk0hwkewmncdhlyfle8jc3k9xzhh75txs",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := GetUnbondingDelegation(tt.args.delAddr, tt.args.valAddr)
			if err != nil {
				logger.Error.Fatalln(err)
			}
			logger.Info.Println(ToJson(res))
		})
	}
}
