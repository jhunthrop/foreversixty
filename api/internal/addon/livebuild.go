package addon

import (
	"fmt"

	"github.com/jhunthrop/foreversixty/api/internal/trees"
)

// MinLiveBuild is the first client build of live WoW Forever, the one it
// launched on (4 November 2026), and the single source of truth for which
// FS1 exports are live ones. An export stamped with an older client build
// was made on the Classic beta or a pre-launch client: its characters are
// not live characters, and on 2026-11-10 a companion that still read the
// beta folder uploaded six of them (1.60.1.70374 and 1.60.1.70450) beside
// the owner's live ones. The client watcher's build manifests name a
// product, not the launch (1.60.1.70532 to 1.60.1.70755 were "wowf" builds
// before it), so the floor is set here, by hand, once.
const MinLiveBuild = "1.60.1.70790"

// preLiveProblem is the player's sentence for an FS1 export made before
// launch, or "" when the export is live or its build cannot be read as a
// client build. A head with no client build ("FS1:aaa", a named data set)
// is left to the rest of the pipeline, which treats the export as opaque.
func preLiveProblem(name, export string) string {
	build, ok := ParseFS1Build(export)
	if !ok || !trees.IsClientBuild(build) || trees.CompareVersions(build, MinLiveBuild) >= 0 {
		return ""
	}
	return fmt.Sprintf("%s's export is from the beta (build %s), not live WoW Forever. "+
		"Log in to WoW Forever and export again.", name, build)
}
