// web/src/lib/sim/spec-icon.ts
// A spec's lead image on /sim/specs: its own talent-tab icon, or the class crest when the build
// has not published one (or the icon file fails to load). Damage rows (SpecCard) and the healer
// and tank rows (UnsimulatedSpecCard) read it here, so both lists lead with the same art.
import activeBuild from '../../data/active-build.json';
import { classCrestSrc } from '../class-crest';
import { dataUrl } from '../planner/load';
import { classOfSpec } from './spec-label';
import { SPEC_BY_KEY } from './specs';

export interface SpecLeadIcon {
  src: string;
  /** The class crest stands in: a round image, not the square spell icon. */
  crest: boolean;
}

/** `failed`: the spec icon's file did not load, so the crest takes its place. */
export function specLeadIcon(spec: string, failed = false): SpecLeadIcon {
  const icon = SPEC_BY_KEY[spec]?.icon;
  return icon === undefined || failed
    ? { src: classCrestSrc(classOfSpec(spec)), crest: true }
    : { src: dataUrl(activeBuild.build, `icons/${icon}.webp`), crest: false };
}
