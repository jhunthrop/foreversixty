//go:build windows

package addonupdater

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

// ListProcesses is the system's process list. Windows is the truth for
// "is the game running": a snapshot of every process, with the full image
// path read for the ones named like a game client.
func ListProcesses() ([]Proc, error) {
	snap, err := windows.CreateToolhelp32Snapshot(windows.TH32CS_SNAPPROCESS, 0)
	if err != nil {
		return nil, err
	}
	defer windows.CloseHandle(snap)
	var pe windows.ProcessEntry32
	pe.Size = uint32(unsafe.Sizeof(pe))
	var out []Proc
	err = windows.Process32First(snap, &pe)
	for err == nil {
		name := windows.UTF16ToString(pe.ExeFile[:])
		p := Proc{Name: name}
		if IsClientName(name) {
			p.Path = imagePath(pe.ProcessID)
		}
		out = append(out, p)
		err = windows.Process32Next(snap, &pe)
	}
	if err != windows.ERROR_NO_MORE_FILES {
		return out, err
	}
	return out, nil
}

// imagePath is the full path of a process's executable, empty when the
// system will not say.
func imagePath(pid uint32) string {
	h, err := windows.OpenProcess(windows.PROCESS_QUERY_LIMITED_INFORMATION, false, pid)
	if err != nil {
		return ""
	}
	defer windows.CloseHandle(h)
	buf := make([]uint16, 32768)
	n := uint32(len(buf))
	if err := windows.QueryFullProcessImageName(h, 0, &buf[0], &n); err != nil {
		return ""
	}
	return windows.UTF16ToString(buf[:n])
}
