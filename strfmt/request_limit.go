package strfmt

import (
	"fmt"
	"regexp"
	"strconv"
)

var RequestLimitInvalid = "[request limit invalid]"

var requestLimitReg = regexp.MustCompile(`^([+\-]?[0-9.]+)?(\/([+\-]?[0-9.]+))?([eEinumkKMGTP]*[-+]?[0-9]*)$`)

func ParseRequestLimit(s string) (*RequestLimit, error) {
	if s == "" || !requestLimitReg.MatchString(s) {
		return nil, fmt.Errorf(RequestLimitInvalid + "missing request and limit")
	}

	parts := requestLimitReg.FindAllStringSubmatch(s, 1)[0]

	rl := &RequestLimit{}

	if parts[1] != "" {
		i, err := strconv.ParseInt(parts[1], 10, 64)
		if err == nil {
			rl.Request = int(i)
		}
	}

	if parts[3] != "" {
		i, err := strconv.ParseInt(parts[3], 10, 64)
		if err == nil {
			rl.Limit = int(i)
		}
	}

	if parts[4] != "" {
		rl.Unit = parts[4]
	}

	return rl, nil
}

type RequestLimit struct {
	Request int
	Limit   int
	Unit    string
}

func (rl *RequestLimit) String() string {
	v := ""
	if rl.Request != 0 {
		v = strconv.FormatInt(int64(rl.Request), 10)
	}
	if rl.Limit != 0 {
		v += "/" + strconv.FormatInt(int64(rl.Limit), 10)
	}
	if rl.Unit != "" {
		v += rl.Unit
	}
	return v
}

func (rl *RequestLimit) RequestString() string {
	v := ""
	if rl.Request != 0 {
		v = strconv.FormatInt(int64(rl.Request), 10)
	}
	if rl.Unit != "" {
		v += rl.Unit
	}
	return v
}

func (rl *RequestLimit) LimitString() string {
	v := ""
	if rl.Limit != 0 {
		v = strconv.FormatInt(int64(rl.Limit), 10)
	}
	if rl.Unit != "" {
		v += rl.Unit
	}
	return v
}

func (rl *RequestLimit) MarshalText() ([]byte, error) {
	return []byte(rl.String()), nil
}

func (rl *RequestLimit) UnmarshalText(data []byte) error {
	servicePort, err := ParseRequestLimit(string(data))
	if err != nil {
		return err
	}
	*rl = *servicePort
	return nil
}
