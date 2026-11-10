//go:build !darwin

// companion/internal/install/system_other.go
// Only the Mac moves itself: the Windows setup program installs the
// companion, and Linux has no Applications folder.
package install

// Supported reports whether this platform moves itself into Applications.
const Supported = false

// NewSystem is nil off the Mac; nothing calls it there.
func NewSystem() System { return nil }

// Detect is the bare Env off the Mac.
func Detect(home string) Env { return Env{Home: home} }
