// companion/internal/app/addons.go
// The Addon row's version lines and the Install the addon button: the
// window's side of internal/addonupdater, which does the downloading and the
// replacing. The words are copy.go's.
package app

import (
	"context"
	"errors"
	"slices"
	"time"

	"github.com/jhunthrop/foreversixty/companion/internal/addonupdater"
	"github.com/jhunthrop/foreversixty/companion/internal/wow"
)

// AddonState is what a version line says about its version.
type AddonState string

// The three states of a version line.
const (
	// AddonCurrent: the newest, or at least what the install has.
	AddonCurrent AddonState = "current"
	// AddonInstalled: put in by the companion, no login seen since.
	AddonInstalled AddonState = "installed"
	// AddonReady: downloaded and verified, waiting for the game to close.
	AddonReady AddonState = "ready"
)

// AddonVersionLine is one line under the Addon row's sync line.
type AddonVersionLine struct {
	Text string `json:"text"`
	// Tone is "muted" for a line that is a fact and not a fault.
	Tone string `json:"tone,omitempty"`
}

// AddonOffer is an Install the addon into <folder> button.
type AddonOffer struct {
	Path  string `json:"path"`
	Label string `json:"label"`
}

// AddonUpdateView is the Addon row's version lines and install buttons.
type AddonUpdateView struct {
	// Enabled is the Keep the addon up to date switch.
	Enabled bool               `json:"enabled"`
	Lines   []AddonVersionLine `json:"lines"`
	Offers  []AddonOffer       `json:"offers"`
}

// addonUpdatesOn is the switch, read for the updater on its own goroutine.
func (a *App) addonUpdatesOn() bool {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.cfg.AddonUpdates()
}

// installPaths are the game installs the updater keeps current.
func (a *App) installPaths() []string {
	a.mu.Lock()
	defer a.mu.Unlock()
	out := make([]string, 0, len(a.installs))
	for _, in := range a.installs {
		out = append(out, in.Path)
	}
	return out
}

// tickAddons gives the updater its turn without holding up the tail: a
// download is minutes on a bad line, and the log must keep being read.
func (a *App) tickAddons(now time.Time) {
	if !a.addons.Due(now) || !a.addonsBusy.CompareAndSwap(false, true) {
		return
	}
	a.bgWG.Add(1)
	go func() {
		defer a.bgWG.Done()
		defer a.addonsBusy.Store(false)
		a.addons.Tick(a.bg, now)
	}()
}

// familyLine words one family in one install, or reports that there is
// nothing to say.
func familyLine(name string, v addonupdater.FamilyView) (string, bool) {
	switch {
	case v.Ready != "":
		return CopyVersionLine(name, v.Ready, AddonReady), true
	case v.Version == "":
		return "", false
	case v.Fresh:
		return CopyVersionLine(name, v.Version, AddonInstalled), true
	}
	return CopyVersionLine(name, v.Version, AddonCurrent), true
}

// addonInstall is one install as the Addon row sees it.
type addonInstall struct {
	Path string
	// Folder is the install's name in the player's words ("Classic Era").
	Folder string
	View   addonupdater.InstallView
}

// addonUpdateView reads every install's addon folders and words them.
func (a *App) addonUpdateView() AddonUpdateView {
	a.mu.Lock()
	installs := slices.Clone(a.installs)
	on := a.cfg.AddonUpdates()
	a.mu.Unlock()
	in := make([]addonInstall, len(installs))
	for i, ins := range installs {
		at, _ := wow.LastActivity(ins)
		in[i] = addonInstall{Path: ins.Path, Folder: wow.FlavorName(ins.Flavor), View: a.addons.View(ins.Path, at)}
	}
	return buildAddonView(on, a.addons.Keyed(), in)
}

// buildAddonView words the installs. Where more than one install has the
// addon each line names its install; an install without it says so, once,
// and offers the button (when this build can verify a zip at all).
func buildAddonView(on, keyed bool, installs []addonInstall) AddonUpdateView {
	view := AddonUpdateView{Enabled: on, Lines: []AddonVersionLine{}, Offers: []AddonOffer{}}
	with := 0
	for _, in := range installs {
		if in.View.Addon.Present {
			with++
		}
	}
	for _, in := range installs {
		if !in.View.Addon.Present {
			view.Lines = append(view.Lines, AddonVersionLine{Text: CopyAddonNotFound(in.Folder), Tone: "muted"})
			if keyed {
				view.Offers = append(view.Offers, AddonOffer{Path: in.Path, Label: CopyInstallAddon(in.Folder)})
			}
			continue
		}
		for _, l := range []struct {
			name string
			v    addonupdater.FamilyView
		}{{CopyAddonName, in.View.Addon}, {CopyDataName, in.View.Data}} {
			text, ok := familyLine(l.name, l.v)
			if !ok {
				continue
			}
			if with > 1 {
				text = CopyForInstall(in.Folder, text)
			}
			view.Lines = append(view.Lines, AddonVersionLine{Text: text})
		}
	}
	return view
}

// InstallAddon is the Install the addon button: it puts the addon, and the
// data addon with it, into one game install, from the newest signed release.
// A failure is a Fault the window words; the cause goes to the log.
func (a *App) InstallAddon(ctx context.Context, path string) error {
	a.mu.Lock()
	known := slices.ContainsFunc(a.installs, func(in wow.Install) bool { return in.Path == path })
	a.mu.Unlock()
	if !known {
		return fault(CodeBadRequest, errors.New("that is not one of the game folders"))
	}
	err := a.addons.Install(ctx, path)
	switch {
	case err == nil:
		a.log.Info("the addon was installed from the button", "component", "app", "install", path)
		return nil
	case errors.Is(err, addonupdater.ErrGameRunning):
		return fault(CodeGameRunning, err)
	case errors.Is(err, addonupdater.ErrDownload), errors.Is(err, addonupdater.ErrNoRelease),
		errors.Is(err, addonupdater.ErrNoKey):
		a.log.Warn("the addon could not be fetched", "component", "app", "install", path, "err", err.Error())
		return fault(CodeAddonDownload, err)
	}
	a.log.Warn("the addon could not be installed", "component", "app", "install", path, "err", err.Error())
	return fault(CodeAddonInstall, err)
}
