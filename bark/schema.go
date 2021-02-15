package bark

import (
	"fmt"
	"strings"
)

// Device represents a registered device with Bark app installed that receives notifications.
type Device struct {
	endpointURL string
	deviceKey   string
}

func (d Device) String() string {
	return fmt.Sprintf("{Device:%s@%s}", d.deviceKey, d.endpointURL)
}

// Options represents options for notification request.
type Options struct {
	Ringtone     RingtoneName
	OpenURL      string
	CopyText     string
	ForceArchive bool
	ForceCopy    bool
}

func (o Options) String() string {
	var parts []string
	if isNotBlank(string(o.Ringtone)) {
		parts = append(parts, fmt.Sprintf("Ringtone:%s", o.Ringtone))
	}
	if isNotBlank(o.OpenURL) {
		parts = append(parts, fmt.Sprintf("URL:%s", o.OpenURL))
	}
	if isNotBlank(o.CopyText) {
		parts = append(parts, fmt.Sprintf("Copy:%q", o.CopyText))
	}
	if o.ForceArchive {
		parts = append(parts, "Archive")
	}
	if o.ForceCopy {
		parts = append(parts, "AutoCopy")
	}
	return "{" + strings.Join(parts, ",") + "}"
}
