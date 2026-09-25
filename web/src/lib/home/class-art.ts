// web/src/lib/home/class-art.ts
// The home hero's backdrop for a character whose realm Blizzard serves no render for: one of
// the class's talent-tree backgrounds the site already ships per build (scripts/sync-data.mjs,
// /data/<build>/trees/<file>.webp). Every class has one, so every hero has a picture; the
// Blizzard render, when it exists, sits on top of it.
import activeBuild from '../../data/active-build.json';
import { dataUrl } from '../planner/load';

/** The tree whose art stands for the class: the one players picture first. */
const TREE_FOR_CLASS: Readonly<Record<string, string>> = {
  warrior: 'warriorarms',
  paladin: 'paladinholy',
  hunter: 'huntermarksmanship',
  rogue: 'roguecombat',
  priest: 'priestshadow',
  shaman: 'shamanelementalcombat',
  mage: 'magefire',
  warlock: 'warlockdestruction',
  druid: 'druidferalcombat',
};

/** The class art's url for `classSlug` (any case, spaces tolerated), or `undefined` for a
 *  class the build has no tree for. */
export function classArtUrl(
  classSlug: string | undefined,
  build: string = activeBuild.build,
): string | undefined {
  if (classSlug === undefined) return undefined;
  const tree = TREE_FOR_CLASS[classSlug.trim().toLowerCase()];
  return tree === undefined ? undefined : dataUrl(build, `trees/${tree}.webp`);
}
