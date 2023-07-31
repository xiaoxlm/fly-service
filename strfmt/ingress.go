package strfmt

import (
	"fmt"
	networkingv1 "k8s.io/api/networking/v1"
	"strconv"
	"strings"
)

var IngressRuleInvalid = "[ingress rule invalid]"

// /path default
// /v1! Exact pathType
// /v1* Prefix pathType
func parsePath(paths string) (pathRules []PathRule) {
	segments := strings.Split(paths, ",")

	for _, p := range segments {
		if p == "" {
			continue
		}

		if strings.HasSuffix(p, "!") {
			pathRules = append(pathRules, PathRule{
				Path:     p[0 : len(p)-1],
				PathType: networkingv1.PathTypeExact,
			})
		} else if strings.HasSuffix(p, "*") {
			pathRules = append(pathRules, PathRule{
				Path:     p[0 : len(p)-1],
				PathType: networkingv1.PathTypePrefix,
			})
		} else {
			pathRules = append(pathRules, PathRule{
				Path:     p,
				PathType: networkingv1.PathTypeImplementationSpecific,
			})
		}
	}

	return
}

// https://www.baidu.com,/path,/v1!,/v2*
func ParseIngress(ingress string) (*Ingress, error) {
	if ingress == "" {
		return nil, fmt.Errorf(IngressRuleInvalid + "ingress content can't be empty")
	}

	var (
		origin = ingress
		paths  = ""
	)
	if i := strings.Index(ingress, ","); i != -1 {
		origin = ingress[0:i]
		paths = ingress[i+1:]
	}

	var schemeHost = strings.Split(origin, "://")
	{
		// check
		if len(schemeHost) != 2 {
			return nil, fmt.Errorf(IngressRuleInvalid)
		}
	}

	var scheme string
	{
		scheme = schemeHost[0]
		if scheme == "" {
			scheme = "http"
		}
	}

	var (
		host string
		port uint16
	)
	{
		hostPort := strings.Split(schemeHost[1], ":")

		host = hostPort[0]

		if len(hostPort) > 1 {
			p, err := strconv.ParseUint(hostPort[1], 10, 16)
			if err != nil {
				return nil, fmt.Errorf(IngressRuleInvalid + fmt.Sprintf("parse port error:%v", err))
			}
			port = uint16(p)

			if port == 0 {
				port = 80
			}
		}
	}

	return &Ingress{
		Scheme: scheme,
		Host:   host,
		Port:   port,
		Paths:  parsePath(paths),
	}, nil
}

type Ingress struct {
	Scheme string
	Host   string
	Port   uint16
	Paths  []PathRule
}

type PathRule struct {
	Path     string
	PathType networkingv1.PathType
}
