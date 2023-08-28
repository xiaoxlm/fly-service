package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	"serving.kubefly.tech/fly-service/strfmt"
)

type Pod struct {
	InitContainers []Container `json:"initContainers,omitempty"`
	Containers     []Container `json:"containers"`

	RestartPolicy                 string                             `json:"restartPolicy,omitempty"`
	TerminationGracePeriodSeconds *int64                             `json:"terminationGracePeriodSeconds,omitempty"`
	ActiveDeadlineSeconds         *int64                             `json:"activeDeadlineSeconds,omitempty"`
	DNSConfig                     *DNSConfig                         `json:"dnsConfig,omitempty"`
	DNSPolicy                     string                             `json:"dnsPolicy,omitempty"`
	NodeSelector                  map[string]string                  `json:"nodeSelector,omitempty"`
	Hosts                         []*strfmt.HostAlias                `json:"hosts,omitempty"`
	ServiceAccountName            string                             `json:"serviceAccountName,omitempty"`
	Volumes                       Volumes                            `json:"volumes,omitempty"`
	TopologySpreadConstraints     []*corev1.TopologySpreadConstraint `json:"topologySpreadConstraints,omitempty" `
	Affinity                      *corev1.Affinity                   `json:"affinity,omitempty"`
	Tolerations                   []*corev1.Toleration               `json:"tolerations,omitempty"`
}

type Volumes map[string]*corev1.VolumeSource

type DNSConfig struct {
	Nameservers []string      `json:"nameservers,omitempty" yaml:"nameservers,omitempty"`
	Searches    []string      `json:"searches,omitempty" yaml:"searches,omitempty"`
	Options     []*KubeOption `json:"options,omitempty" yaml:"options,omitempty"`
}

type KubeOption struct {
	Name  string `yaml:"name" json:"name" toml:"name"`
	Value string `yaml:"value,omitempty" json:"value,omitempty" toml:"value,omitempty"`
}
