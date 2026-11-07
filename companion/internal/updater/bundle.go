// companion/internal/updater/bundle.go
// The macOS companion is an app bundle, signed and notarized as a whole:
// the bundle's signature lives in its main executable and seals the
// Info.plist and the resources beside it. Swapping the bare executable
// inside a bundle would break that seal, so a Mac update replaces the whole
// bundle: the release carries the signed, stapled bundle as a zip, and the
// swap renames the running bundle aside and the new one into its place.
package updater

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// BundleAssetSuffix ends the name of an asset that is a zipped app bundle.
const BundleAssetSuffix = ".app.zip"

// maxBundleBytes caps what a bundle zip may unpack to: a zip is a few
// megabytes of compressed Go binary, and nothing it says should fill a disk.
const maxBundleBytes = 512 << 20

// zipMagic opens every zip file; no executable starts with it.
var zipMagic = []byte("PK\x03\x04")

// BundleRoot is the app bundle the executable at exe runs from,
// "/Applications/Forever Sixty.app" for
// "/Applications/Forever Sixty.app/Contents/MacOS/foreversixty-companion",
// or "" when exe is not a bundle's main executable.
func BundleRoot(exe string) string {
	macos := filepath.Dir(exe)
	contents := filepath.Dir(macos)
	bundle := filepath.Dir(contents)
	if filepath.Base(macos) != "MacOS" || filepath.Base(contents) != "Contents" ||
		!strings.HasSuffix(bundle, ".app") {
		return ""
	}
	return bundle
}

// applyBundle swaps the verified bundle zip in body over the bundle that
// holds exe. The new bundle is unpacked into a hidden folder beside the
// running one first, so both renames stay on one volume, and the folder
// takes the old bundle with it when it is removed.
func applyBundle(body []byte, bundle, exe string) error {
	stage, err := os.MkdirTemp(filepath.Dir(bundle), "."+filepath.Base(bundle)+".update-*")
	if err != nil {
		return fmt.Errorf("stage the new app next to the running one: %w", err)
	}
	defer os.RemoveAll(stage)

	fresh, err := unpackBundle(body, filepath.Join(stage, "new"))
	if err != nil {
		return err
	}
	// The relaunch starts the same path the running build has, so the new
	// bundle must carry an executable of the same name.
	if fi, err := os.Stat(filepath.Join(fresh, "Contents", "MacOS", filepath.Base(exe))); err != nil || !fi.Mode().IsRegular() {
		return fmt.Errorf("the staged app has no Contents/MacOS/%s", filepath.Base(exe))
	}

	previous := filepath.Join(stage, "previous.app")
	if err := rename(bundle, previous); err != nil {
		return fmt.Errorf("move the running app aside: %w", err)
	}
	if err := rename(fresh, bundle); err != nil {
		if restoreErr := rename(previous, bundle); restoreErr != nil {
			return fmt.Errorf("swap in the new app: %w; restoring the previous app also failed: %w", err, restoreErr)
		}
		return fmt.Errorf("swap in the new app: %w", err)
	}
	return nil
}

// unpackBundle unpacks a zip holding one .app folder at its top level
// (ditto -c -k --keepParent makes exactly that) into dest and returns the
// unpacked bundle's path. An entry that would land outside dest, a link
// that points outside the bundle, or a zip of anything else is refused.
func unpackBundle(body []byte, dest string) (string, error) {
	zr, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		return "", fmt.Errorf("the staged app is not a readable zip: %w", err)
	}
	top := ""
	var total uint64
	for _, f := range zr.File {
		name := f.Name
		if strings.HasPrefix(name, "__MACOSX/") {
			continue // Finder's resource-fork shadows; the bundle never needs them
		}
		first, _, _ := strings.Cut(name, "/")
		if top == "" {
			top = first
		}
		if first != top || !strings.HasSuffix(top, ".app") {
			return "", fmt.Errorf("the staged zip holds %q, not one app bundle", name)
		}
		total += f.UncompressedSize64
		if total > maxBundleBytes {
			return "", errors.New("the staged app unpacks to more than any companion could be")
		}
		if err := unpackEntry(f, dest); err != nil {
			return "", err
		}
	}
	if top == "" {
		return "", errors.New("the staged zip is empty")
	}
	return filepath.Join(dest, top), nil
}

// unpackEntry writes one zip entry under dest.
func unpackEntry(f *zip.File, dest string) error {
	if !filepath.IsLocal(filepath.FromSlash(f.Name)) {
		return fmt.Errorf("the staged zip names %q, outside the app", f.Name)
	}
	path := filepath.Join(dest, filepath.FromSlash(f.Name))
	mode := f.Mode()
	switch {
	case mode.IsDir():
		return os.MkdirAll(path, 0o755)
	case mode&fs.ModeSymlink != 0:
		return unpackLink(f, path, dest)
	case mode.IsRegular():
		return unpackFile(f, path, mode.Perm()|0o600)
	}
	return fmt.Errorf("the staged zip holds %q, which is neither a file, a folder nor a link", f.Name)
}

func unpackFile(f *zip.File, path string, perm fs.FileMode) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	rc, err := f.Open()
	if err != nil {
		return fmt.Errorf("read %s from the staged zip: %w", f.Name, err)
	}
	defer rc.Close()
	out, err := os.OpenFile(path, os.O_CREATE|os.O_EXCL|os.O_WRONLY, perm)
	if err != nil {
		return err
	}
	if _, err := io.Copy(out, io.LimitReader(rc, maxBundleBytes)); err != nil {
		out.Close()
		return fmt.Errorf("unpack %s: %w", f.Name, err)
	}
	return out.Close()
}

// unpackLink recreates a symbolic link whose target stays inside the
// unpacked tree, as a framework's Versions/Current does.
func unpackLink(f *zip.File, path, dest string) error {
	rc, err := f.Open()
	if err != nil {
		return fmt.Errorf("read %s from the staged zip: %w", f.Name, err)
	}
	target, err := io.ReadAll(io.LimitReader(rc, 4096))
	rc.Close()
	if err != nil {
		return fmt.Errorf("read %s from the staged zip: %w", f.Name, err)
	}
	t := string(target)
	resolved, err := filepath.Rel(dest, filepath.Join(filepath.Dir(path), t))
	if filepath.IsAbs(t) || err != nil || !filepath.IsLocal(resolved) {
		return fmt.Errorf("the staged zip links %q outside the app", f.Name)
	}
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	return os.Symlink(t, path)
}
