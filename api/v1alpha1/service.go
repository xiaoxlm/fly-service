package v1alpha1

import "serving.kubefly.tech/fly-service/strfmt"

type ServiceType string

const (
	ServiceTypeClusterIP    ServiceType = "ClusterIP"
	ServiceTypeNodePort     ServiceType = "NodePort"
	ServiceTypeLoadBalancer ServiceType = "LoadBalancer"
	ServiceTypeExternalName ServiceType = "ExternalName"
)

type Service struct {
	Name         string                `json:"name"`
	Type         ServiceType           `json:"type,omitempty"`
	PortForwards []*strfmt.PortForward `json:"portForwards,omitempty"`
}
