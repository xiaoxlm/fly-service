package v1alpha1

import "serving.kubefly.tech/fly-service/strfmt"

type Ingress struct {
	Ingresses []*strfmt.Ingress    `json:"ingresses,omitempty"`
	TLS       []*strfmt.IngressTLS `json:"tls,omitempty"`
}
