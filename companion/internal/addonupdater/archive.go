// companion/internal/addonupdater/archive.go
package addonupdater

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"slices"
	"strings"
)

// Ceilings on what a zip may unpack to. The real addon is well under a
// few megabytes; these only stop a hostile or corrupt archive from
// filling a disk.
const (
	maxUnpacked = 64 << 20
	maxFiles    = 4000
)

// extract unpacks a zip into dest, which must exist. The zip may hold only
// the family's own folders: a file at the top level, a folder the family
// does not ship, a path that climbs out, a symlink or an oversized archive
// refuses the whole zip before anything has been written to a game folder.
// Every folder must carry its own TOC.
func extract(body []byte, dest string, folders []string) error {
	zr, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		return fmt.Errorf("the zip is unreadable: %w", err)
	}
	if len(zr.File) > maxFiles {
		return fmt.Errorf("the zip holds %d entries, over the %d ceiling", len(zr.File), maxFiles)
	}
	var total uint64
	for _, f := range zr.File {
		name := f.Name
		if name == "" || strings.ContainsAny(name, "\\:\x00") || strings.HasPrefix(name, "/") {
			return fmt.Errorf("the zip holds an unsafe path %q", name)
		}
		clean := path.Clean(strings.TrimSuffix(name, "/"))
		parts := strings.Split(clean, "/")
		if slices.Contains(parts, "..") || clean == "." {
			return fmt.Errorf("the zip holds an unsafe path %q", name)
		}
		if !slices.Contains(folders, parts[0]) {
			return fmt.Errorf("the zip holds %q, which is not one of %s", name, strings.Join(folders, ", "))
		}
		mode := f.Mode()
		if mode&(fs.ModeSymlink|fs.ModeDevice|fs.ModeNamedPipe|fs.ModeSocket|fs.ModeIrregular) != 0 {
			return fmt.Errorf("the zip holds a link or special file %q", name)
		}
		target := filepath.Join(dest, filepath.FromSlash(clean))
		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(target, 0o755); err != nil {
				return err
			}
			continue
		}
		if len(parts) == 1 {
			return fmt.Errorf("the zip holds a file %q at the top level", name)
		}
		total += f.UncompressedSize64
		if total > maxUnpacked {
			return errors.New("the zip unpacks to more than the ceiling")
		}
		if err := writeEntry(f, target); err != nil {
			return err
		}
	}
	for _, folder := range folders {
		if _, err := os.Stat(filepath.Join(dest, folder, folder+".toc")); err != nil {
			return fmt.Errorf("the zip has no %s/%s.toc", folder, folder)
		}
	}
	return nil
}

func writeEntry(f *zip.File, target string) error {
	if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
		return err
	}
	rc, err := f.Open()
	if err != nil {
		return err
	}
	defer rc.Close()
	out, err := os.OpenFile(target, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	// The declared size is not trusted: one byte past it is a lie.
	n, err := io.Copy(out, io.LimitReader(rc, int64(f.UncompressedSize64)+1))
	if cerr := out.Close(); err == nil {
		err = cerr
	}
	if err != nil {
		return err
	}
	if uint64(n) != f.UncompressedSize64 {
		return fmt.Errorf("%s is not the size the zip says", f.Name)
	}
	return nil
}

// copyTree copies a directory of regular files and folders to dst. Anything
// else (a link, a device) fails it.
func copyTree(src, dst string) error {
	return filepath.WalkDir(src, func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(src, p)
		if err != nil {
			return err
		}
		target := filepath.Join(dst, rel)
		switch {
		case d.IsDir():
			return os.MkdirAll(target, 0o755)
		case d.Type().IsRegular():
			return copyFile(p, target)
		}
		return fmt.Errorf("%s is not a plain file", p)
	})
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.OpenFile(dst, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0o644)
	if err != nil {
		return err
	}
	if _, err := io.Copy(out, in); err != nil {
		out.Close()
		return err
	}
	return out.Close()
}
