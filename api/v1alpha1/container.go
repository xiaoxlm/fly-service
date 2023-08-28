package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	"serving.kubefly.tech/fly-service/strfmt"
)

type Container struct {
	Image           string                  `json:"image,omitempty"`
	ImagePullSecret string                  `json:"imagePullSecret,omitempty"`
	ImagePullPolicy string                  `json:"imagePullPolicy,omitempty"`
	WorkingDir      string                  `json:"workingDir,omitempty"`
	Command         []string                `json:"command,omitempty"`
	Args            []string                `json:"args,omitempty"`
	Mounts          []*strfmt.VolumeMount   `json:"mounts,omitempty"`
	Ports           []string                `json:"ports,omitempty"`
	Envs            Envs                    `json:"envs,omitempty"`
	Quota           Quota                   `json:"quota,omitempty"`
	TTY             bool                    `json:"tty,omitempty"`
	ReadinessProbe  *Probe                  `json:"readinessProbe,omitempty"`
	LivenessProbe   *Probe                  `json:"livenessProbe,omitempty"`
	Lifecycle       *Lifecycle              `json:"lifecycle,omitempty"`
	SecurityContext *corev1.SecurityContext `json:"securityContext,omitempty"`
}

type Lifecycle struct {
	PostStart *strfmt.Action `json:"postStart,omitempty"`
	PreStop   *strfmt.Action `json:"preStop,omitempty"`
}

type Probe struct {
	Action    strfmt.Action `json:"action"`
	ProbeOpts `json:",inline"`
}

type ProbeOpts struct {
	InitialDelaySeconds int32 `json:"initialDelaySeconds,omitempty"`
	TimeoutSeconds      int32 `json:"timeoutSeconds,omitempty"`
	PeriodSeconds       int32 `json:"periodSeconds,omitempty"`
	SuccessThreshold    int32 `json:"successThreshold,omitempty"`
	FailureThreshold    int32 `json:"failureThreshold,omitempty"`
}

type Quota map[string]*strfmt.RequestLimit

type Envs map[string]string

func (envs Envs) Merge(srcEnvs Envs) Envs {
	es := Envs{}
	for k, v := range envs {
		es[k] = v
	}
	for k, v := range srcEnvs {
		es[k] = v
	}
	return es
}
