package strfmt

import (
	"fmt"
	"strings"
)

// SecretName:host1,host2,host3
func ParseIngressTLS(s string) (*IngressTLS, error) {
	if s == "" {
		return nil, fmt.Errorf("invalid ingress tls")
	}

	segments := strings.Split(s, ":")

	var (
		secretName string
		hosts      []string
	)
	{
		secretName = segments[0]

		if len(segments) > 1 {
			hosts = strings.Split(segments[1], ",")
		}
	}

	return &IngressTLS{
		SecretName: secretName,
		Hosts:      hosts,
	}, nil
}

type IngressTLS struct {
	SecretName string
	Hosts      []string
}

func (r *IngressTLS) String() string {
	return fmt.Sprintf("%s:%s", r.SecretName, strings.Join(r.Hosts, ","))
}

func (r *IngressTLS) MarshalText() ([]byte, error) {
	return []byte(r.String()), nil
}

func (r *IngressTLS) UnmarshalText(data []byte) error {
	ir, err := ParseIngressTLS(string(data))
	if err != nil {
		return err
	}
	*r = *ir
	return nil
}
