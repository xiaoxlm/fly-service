// probe action
package strfmt

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"net/url"
	"strconv"
	"strings"
)

var ProbeActionInvalid = "[probe action invalid]"

// http://:80
// tcp://:80
// exec
func ParseAction(str string) (*Action, error) {
	if str == "" {
		return nil, nil
	}

	action := &Action{}

	if strings.HasPrefix(str, "http") || strings.HasPrefix(str, "tcp") {
		u, err := url.Parse(str)
		if err != nil {
			return nil, fmt.Errorf(ProbeActionInvalid + err.Error())
		}

		port, _ := strconv.ParseUint(u.Port(), 10, 64)

		// tcp
		if u.Scheme == "tcp" {
			action.TCPSocket = &corev1.TCPSocketAction{}
			action.TCPSocket.Host = u.Hostname()
			action.TCPSocket.Port = intstr.FromInt(int(port))
			return action, nil
		}

		// http
		action.HTTPGet = &corev1.HTTPGetAction{}
		action.HTTPGet.Port = intstr.FromInt(int(port))
		action.HTTPGet.Host = u.Hostname()
		action.HTTPGet.Path = u.Path
		action.HTTPGet.Scheme = corev1.URIScheme(strings.ToUpper(u.Scheme))

		return action, nil
	}

	// exec
	action.Exec = &corev1.ExecAction{
		Command: []string{"sh", "-c", str},
	}

	return action, nil
}

type Action struct {
	corev1.ProbeHandler
}

func (a *Action) String() string {
	if a.Exec != nil {
		return a.Exec.Command[2]
	}

	if a.HTTPGet != nil {
		u := &url.URL{}
		u.Scheme = strings.ToLower(string(a.HTTPGet.Scheme))
		u.Path = a.HTTPGet.Path
		u.Host = a.HTTPGet.Host + ":" + a.HTTPGet.Port.String()

		if u.Scheme != "" {
			u.Scheme = "http"
		}
		return u.String()
	}

	if a.TCPSocket != nil {
		u := &url.URL{}
		u.Scheme = "tcp"
		u.Host = a.TCPSocket.Host + ":" + a.TCPSocket.Port.String()

		return u.String()
	}

	return ""
}

func (a *Action) MarshalText() ([]byte, error) {
	return []byte(a.String()), nil
}

func (a *Action) UnmarshalText(data []byte) error {
	action, err := ParseAction(string(data))
	if err != nil {
		return err
	}
	if action != nil {
		*a = *action
	}
	return nil
}
