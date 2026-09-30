// web/src/lib/bis/spec-icon.ts
// `SpecTabs`' own 28px icon (design/DESIGN-SYSTEM.md's Class page header, bis rebuild spec
// §4.A) -- verified against real data that this build carries NO dedicated talent-TAB icon
// field anywhere (`data/curated/specs.json`, `data/builds/<build>/talents.json` and the
// pipeline's own `data/pipeline/normalize/talent_trees.py` were all checked; the per-class
// `talents/<class>.json` only resolves an icon per TALENT, via `pipeline.icons.resolve_icon`
// on that talent's own first rank). This reads the tree's own first talent's icon as the
// best available per-spec stand-in -- a real, per-spec-tree icon grounded in this build's
// own client data, not an invented image -- flagged in this lane's own report as a gap a
// future data lane could close by capturing the client's actual tab icon.
import { existsSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';

function repoRoot(): string {
  return path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../../../..');
}

interface TalentTreeFile {
  trees: { talents: { icon: string }[] }[];
}

/** `undefined` when the class has no talent file for this build, or `treeIndex` is out of
 *  range (should not happen on real data -- `specs.json`'s own `tree_index` always indexes
 *  a real tree) -- the caller renders a plain neutral square rather than guess. */
export function specTabIcon(build: string, classSlug: string, treeIndex: number): string | undefined {
  const file = path.join(repoRoot(), 'data/builds', build, 'talents', `${classSlug}.json`);
  if (!existsSync(file)) return undefined;
  const data = JSON.parse(readFileSync(file, 'utf8')) as TalentTreeFile;
  return data.trees[treeIndex]?.talents[0]?.icon;
}
