package strfmt

import (
	networkingv1 "k8s.io/api/networking/v1"
	"reflect"
	"testing"
)

func TestParseIngress(t *testing.T) {
	type args struct {
		ingress string
	}

	tests := []struct {
		name    string
		args    args
		want    *Ingress
		wantErr bool
	}{
		{
			name: "#default",
			args: args{ingress: "https://www.baidu.com:8080,/aaa"},
			want: &Ingress{
				Scheme: "https",
				Host:   "www.baidu.com",
				Port:   8080,
				Paths: []PathRule{
					{
						Path:     "/aaa",
						PathType: networkingv1.PathTypeImplementationSpecific,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "#excact path type",
			args: args{ingress: "https://www.baidu.com:8080,/aaa!"},
			want: &Ingress{
				Scheme: "https",
				Host:   "www.baidu.com",
				Port:   8080,
				Paths: []PathRule{
					{
						Path:     "/aaa",
						PathType: networkingv1.PathTypeExact,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "#prefix path type",
			args: args{ingress: "https://www.baidu.com:8080,/aaa*"},
			want: &Ingress{
				Scheme: "https",
				Host:   "www.baidu.com",
				Port:   8080,
				Paths: []PathRule{
					{
						Path:     "/aaa",
						PathType: networkingv1.PathTypePrefix,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "#mix",
			args: args{ingress: "https://www.baidu.com:8080,/aaa,/bbb!,/ccc*"},
			want: &Ingress{
				Scheme: "https",
				Host:   "www.baidu.com",
				Port:   8080,
				Paths: []PathRule{
					{
						Path:     "/aaa",
						PathType: networkingv1.PathTypeImplementationSpecific,
					},
					{
						Path:     "/bbb",
						PathType: networkingv1.PathTypeExact,
					},
					{
						Path:     "/ccc",
						PathType: networkingv1.PathTypePrefix,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseIngress(tt.args.ingress)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseIngress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseIngress() got = %v, want %v", got, tt.want)
			}
		})
	}
}
