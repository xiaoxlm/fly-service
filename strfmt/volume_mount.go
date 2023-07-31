package strfmt

import (
	"bytes"
	"fmt"
	"strings"
)

var VolumeMountInvalid = "[volume mount invalid]"

// data[/subPath]:mountPath[:ro]
func ParseVolumeMount(str string) (*VolumeMount, error) {
	if str == "" {
		return nil, fmt.Errorf(VolumeMountInvalid + "str can't be empty")
	}

	segments := strings.Split(str, ":")
	vm := &VolumeMount{}

	if len(segments) > 3 {
		return nil, fmt.Errorf(VolumeMountInvalid + "formatting error")
	}

	if len(segments) == 3 {
		n := len(segments)
		if segments[n-1] != "ro" {
			return nil, fmt.Errorf(VolumeMountInvalid + `"ro" invalid`)
		}
		vm.ReadOnly = true
		segments = segments[0 : n-1]
	}

	if len(segments) != 2 {
		return nil, fmt.Errorf(VolumeMountInvalid + "formatting len != 2")
	}

	vm.MountPath = segments[1]

	volumeFrom := strings.Split(segments[0], "/")

	vm.Name = volumeFrom[0]

	if len(volumeFrom) == 2 {
		vm.SubPath = volumeFrom[1]
	}

	return vm, nil
}

type VolumeMount struct {
	Name      string
	MountPath string
	SubPath   string
	ReadOnly  bool
}

func (v *VolumeMount) String() string {
	buf := bytes.NewBufferString(v.Name)
	if v.SubPath != "" {
		buf.WriteByte('/')
		buf.WriteString(v.SubPath)
	}

	buf.WriteByte(':')
	buf.WriteString(v.MountPath)

	if v.ReadOnly {
		buf.WriteString(":ro")
	}

	return buf.String()
}

func (v *VolumeMount) MarshalText() ([]byte, error) {
	return []byte(v.String()), nil
}

func (v *VolumeMount) UnmarshalText(data []byte) error {
	vm, err := ParseVolumeMount(string(data))
	if err != nil {
		return err
	}
	*v = *vm
	return nil
}
