package engagement

import "github.com/jhunthrop/foreversixty/api/internal/auth"

// Surface is one bit of user_active_days.surfaces: what the account was using.
type Surface int16

// The surfaces, as the table's bits.
const (
	SurfaceSite  Surface = 1 // a browser session's API call
	SurfaceAddon Surface = 2 // the desktop companion's device token (the addon's syncs)
	SurfaceApp   Surface = 4 // the phone app's device token
)

// SurfaceOf names the surface an authenticated actor came through.
func SurfaceOf(a auth.Actor) Surface {
	switch {
	case a.IsApp():
		return SurfaceApp
	case a.IsDevice():
		return SurfaceAddon
	default:
		return SurfaceSite
	}
}

// Has reports whether the bitmask s holds surface.
func (s Surface) Has(surface Surface) bool { return s&surface != 0 }
