// web/src/lib/guides/talent-cells.ts
// A guide's decoded build as the site's one talent tree draws it (ui/TalentTrees through
// `viewTrees`): the class's own layout, art and arrows at the build's ranks. Build time only (the
// talent file is read from disk). The cells carry no tooltip text: each names its tooltip by key
// and the guide's tooltip host serves it from the page's bundle on first open.
import { viewTrees } from '../planner/editor-trees';
import { indexTalents } from '../planner/rules';
import type { TalentTreeView } from '../ui/talent-tree';
import { loadTalentFile, type DecodedTreeView } from './build-tree';

export function guideTalentTrees(
  build: string,
  classSlug: string,
  trees: readonly DecodedTreeView[],
): TalentTreeView[] {
  const file = trees.length === 0 ? undefined : loadTalentFile(build, classSlug);
  if (file === undefined) return [];
  return viewTrees({
    index: indexTalents(file),
    ranks: new Map(
      trees.flatMap((tree) => tree.talents.map((talent): [number, number] => [talent.id, talent.rank])),
    ),
    classSlug,
    build,
    withText: false,
  });
}
