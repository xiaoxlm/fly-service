package strfmt

import (
	"fmt"
	"strconv"
	"strings"
)

// 80:8000/TCP, 80/TCP
func ParsePortForward(str string) (*PortForward, error) {
	if str == "" {
		return nil, fmt.Errorf("missing port value")
	}

	segment := strings.Split(str, "/")
	if len(segment) > 2 {
		return nil, fmt.Errorf("parse from '/' is invalid")
	}

	var protocol = "TCP"
	if len(segment) == 2 {
		protocol = strings.ToLower(segment[1])
	}

	var (
		port       uint16 = 80
		targetPort uint16
	)
	{
		portSegment := strings.Split(segment[0], ":")

		// port
		p, err := strconv.Atoi(portSegment[0])
		if err != nil {
			return nil, fmt.Errorf("invalid port %v", portSegment[0])
		}
		port = uint16(p)

		// targetPort
		targetPort = port
		if len(portSegment) == 2 {
			tp, err := strconv.Atoi(portSegment[1])
			if err != nil {
				return nil, fmt.Errorf("invalid targetPort %v", portSegment[1])
			}

			targetPort = uint16(tp)
		}

	}

	return &PortForward{
		Port:       port,
		TargetPort: targetPort,
		Protocol:   strings.ToUpper(protocol),
	}, nil
}

type PortForward struct {
	Port       uint16
	TargetPort uint16
	Protocol   string
}

func (s *PortForward) String() string {
	v := ""

	if s.Port != 0 {
		v += strconv.FormatUint(uint64(s.Port), 10)
	}

	if s.TargetPort != 0 {
		v += ":" + strconv.FormatUint(uint64(s.TargetPort), 10)
	}

	if s.Protocol != "" {
		v += "/" + strings.ToUpper(s.Protocol)
	}

	return v
}

func (s *PortForward) MarshalText() ([]byte, error) {
	return []byte(s.String()), nil
}

func (s *PortForward) UnmarshalText(data []byte) error {
	servicePort, err := ParsePortForward(string(data))
	if err != nil {
		return err
	}
	*s = *servicePort
	return nil
}
