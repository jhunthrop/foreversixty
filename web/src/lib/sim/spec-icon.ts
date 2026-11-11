// web/src/lib/sim/spec-icon.ts
// A spec's lead image on /sim/specs: its own talent-tab icon, or null when the build has not
// published one (or the icon file fails to load), and the caller draws the class crest (ui/Crest)
// in its place. Damage rows (SpecCard) and the healer and tank rows (UnsimulatedSpecCard) read it
// here, so both lists lead with the same art.
import activeBuild from '../../data/active-build.json';
import { SPEC_BY_KEY } from './specs';
import { iconUrl } from '../ui/assets';

/** `failed`: the spec icon's file did not load, so the crest takes its place (null). */
export function specLeadIcon(spec: string, failed = false): string | null {
  return failed ? null : iconUrl(activeBuild.build, SPEC_BY_KEY[spec]?.icon);
}
