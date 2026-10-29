// web/src/lib/billing/coach-sample-icons.ts
// Resolves the sample review's spell icons at build time, through the build's own icon table
// (data/builds/<build>/spellicons.json), to the same `/data/<build>/icons/<stem>.webp` urls the
// coach panel uses in the browser (lib/coach/names.ts). A spell the build has no icon for
// resolves to null, and the sample draws nothing and leaves no gap.
import activeBuild from '../../data/active-build.json';
import { buildIconExists, spellIcons } from '../data/build-files';
import { dataUrl } from '../planner/load';
import type { IconOf } from './coach-sample';

export function buildSpellIconOf(build: string = activeBuild.build): IconOf {
  const icons = spellIcons(build);
  return (spellId) => {
    const stem = spellId === null ? undefined : icons[String(spellId)];
    return stem !== undefined && buildIconExists(build, stem) ? dataUrl(build, `icons/${stem}.webp`) : null;
  };
}
