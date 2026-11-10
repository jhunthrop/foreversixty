//go:build windows

// companion/internal/updater/icons_windows.go
package updater

import "golang.org/x/sys/windows"

const (
	shcneAssocChanged = 0x08000000 // SHCNE_ASSOCCHANGED
	shcnfIDList       = 0x0000     // SHCNF_IDLIST
)

var shChangeNotify = windows.NewLazySystemDLL("shell32.dll").NewProc("SHChangeNotify")

// shellNotifyIconChange asks Explorer to drop its icon cache for the
// swapped exe: SHChangeNotify(SHCNE_ASSOCCHANGED, SHCNF_IDLIST, 0, 0). The
// call returns nothing; only a missing shell32 export is an error.
func shellNotifyIconChange() error {
	if err := shChangeNotify.Find(); err != nil {
		return err
	}
	shChangeNotify.Call(shcneAssocChanged, shcnfIDList, 0, 0)
	return nil
}
