package strfmt

import (
	"bytes"
	"strings"
)

type HostAlias struct {
	IP        string
	HostNames []string
}

// 127.0.0.1 test1.com,test2.com
func ParseHostAlias(str string) (*HostAlias, error) {
	if str == "" {
		return nil, nil
	}

	t := &HostAlias{}

	segments := strings.Split(str, ":")
	if len(segments) < 2 {
		segments = strings.Split(str, " ")
		if len(segments) < 2 {
			return nil, nil
		}
	}

	t.IP = segments[0]

	kv := strings.Split(segments[1], ",")

	if len(kv) > 0 {
		t.HostNames = append(t.HostNames, kv...)
	}

	return t, nil
}

func (t *HostAlias) UnmarshalText(text []byte) error {
	to, err := ParseHostAlias(string(text))
	if err != nil {
		return err
	}
	*t = *to
	return nil
}

func (t *HostAlias) MarshalText() (text []byte, err error) {
	return []byte(t.String()), nil
}

func (t *HostAlias) String() string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString(t.IP)
	buf.WriteString(" ")

	if len(t.HostNames) != 0 {
		for index, host := range t.HostNames {
			buf.WriteString(host)
			if index != len(t.HostNames)-1 {
				buf.WriteRune(',')
			}
		}
	}

	return buf.String()
}
