package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"slices"

	"github.com/jhunthrop/foreversixty/api/internal/httpx"
)

// AddonFolders are the folders a companion may report under Interface/AddOns: the addon, its
// inbox and the nightly data addon, the three the releases ship.
var AddonFolders = []string{"ForeverSixty", "ForeverSixtyInbox", "ForeverSixtyData"}

// maxAddonVersion bounds one TOC version a companion reports.
const maxAddonVersion = 32

// addonVersionPattern is what a TOC's "## Version" may say once the companion has normalized it:
// a semver or a date version, or empty when the TOC has no version line.
var addonVersionPattern = regexp.MustCompile(`^[0-9A-Za-z.+-]*$`)

// DeviceAddon is one Forever Sixty addon a companion found installed, with its TOC's version
// (empty when the TOC has none).
type DeviceAddon struct {
	Folder  string `json:"folder"`
	Version string `json:"version"`
}

type deviceAddonsRequest struct {
	Addons []DeviceAddon `json:"addons"`
}

// validDeviceAddons checks a report: known folders, each at most once, versions short and plain.
func validDeviceAddons(addons []DeviceAddon) error {
	if addons == nil {
		return errors.New("addons is required; send [] when none are installed")
	}
	seen := map[string]bool{}
	for _, a := range addons {
		if !slices.Contains(AddonFolders, a.Folder) {
			return fmt.Errorf("unknown addon folder %q", a.Folder)
		}
		if seen[a.Folder] {
			return fmt.Errorf("addon folder %q is listed twice", a.Folder)
		}
		seen[a.Folder] = true
		if len(a.Version) > maxAddonVersion || !addonVersionPattern.MatchString(a.Version) {
			return fmt.Errorf("addon %s has an unreadable version", a.Folder)
		}
	}
	return nil
}

// SetDeviceAddons stores what a companion found installed, replacing its last report.
func (s *Store) SetDeviceAddons(ctx context.Context, deviceID string, addons []DeviceAddon) error {
	body, err := json.Marshal(addons)
	if err != nil {
		return fmt.Errorf("auth: encode device addons: %w", err)
	}
	if _, err := s.Pool.Exec(ctx,
		`update devices set addons = $2 where id = $1 and revoked_at is null`, deviceID, body); err != nil {
		return fmt.Errorf("auth: set device addons: %w", err)
	}
	return nil
}

// putDeviceAddons is PUT /v1/devices/self/addons: the calling companion's AddOns scan.
func (s *Service) putDeviceAddons(w http.ResponseWriter, r *http.Request) {
	var in deviceAddonsRequest
	if err := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxAuthBody)).Decode(&in); err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", "body must be JSON with an addons list", nil)
		return
	}
	if err := validDeviceAddons(in.Addons); err != nil {
		httpx.WriteError(w, r, http.StatusBadRequest, "invalid", err.Error(), nil)
		return
	}
	if err := s.Store.SetDeviceAddons(r.Context(), ActorFrom(r.Context()).DeviceID, in.Addons); err != nil {
		s.fail(w, r, "device_addons", err, "could not save the addons just now")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
