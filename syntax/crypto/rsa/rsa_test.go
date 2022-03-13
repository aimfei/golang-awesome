package rsa

import (
	"reflect"
	"testing"
)

func TestGenRsaKey(t *testing.T) {
	type args struct {
		bits int
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GenRsaKey(tt.args.bits)
			if got != tt.want {
				t.Errorf("GenRsaKey() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GenRsaKey() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestRSAPriDecrypt(t *testing.T) {
	type args struct {
		planText   string
		privateKey string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RSAPriDecrypt(tt.args.planText, tt.args.privateKey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RSAPriDecrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRSAPubEncrypt(t *testing.T) {
	type args struct {
		planText  string
		publicKey string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RSAPubEncrypt(tt.args.planText, tt.args.publicKey); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RSAPubEncrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
