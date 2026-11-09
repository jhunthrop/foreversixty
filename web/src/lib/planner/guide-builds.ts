// web/src/lib/planner/guide-builds.ts
// The planner rail's Guide builds (lane P7, info-density MUST 2): every spec guide's published
// builds for the class on screen, its levelling build and its raid build where it has one, read
// at build time from the guides' frontmatter (`build`, `raidBuild`). Each row names the spec, the
// build, its split and opens it in the guide's own levelling order (lib/guides/build-links.ts).
import { loadBuildHref } from '../guides/build-links';
import { guidesCopy } from '../guides/copy';
import { SPECS } from '../sim/specs';

/** One published guide build, as planner.astro reads it from the guides collection. */
export interface GuideBuild {
  classSlug: string;
  /** The guide's spec slug (`survival`). */
  specSlug: string;
  kind: 'leveling' | 'raid';
  /** The FS1 code (`FS1:<build>:<class>:<race>:<trees>:<gear>`). */
  code: string;
}

export interface GuideBuildRow {
  key: string;
  /** The spec as a player names it (`Survival`). */
  name: string;
  /** The spec's icon file name (no extension), when the catalogue has one. */
  icon?: string;
  /** `Level 60 build` or `Raid build`. */
  label: string;
  /** Points per tree, in tree order. */
  split: number[];
  /** Opens the build in the planner in the guide's levelling order. */
  href: string;
  /** The tree ranks segment, each tree's trailing zeros trimmed: what `sameTrees` compares. */
  trees: string;
}

const TREES_PART = 4;

/** A code's tree ranks segment, each tree's trailing zeros trimmed; null when it has none. */
export function codeTrees(code: string): string | null {
  const part = code.split(':')[TREES_PART];
  if (part === undefined || !/^[0-9/]+$/.test(part)) return null;
  return part
    .split('/')
    .map((tree) => tree.replace(/0+$/, ''))
    .join('/');
}

/** The same normal form for a build's ranks per tree (the planner's `treeRanksFor`). */
export function ranksTrees(ranks: readonly (readonly number[])[]): string {
  return ranks.map((tree) => tree.join('').replace(/0+$/, '')).join('/');
}

/** Points per tree of a tree ranks segment. */
export function treesSplit(trees: string): number[] {
  return trees.split('/').map((tree) => [...tree].reduce((sum, digit) => sum + Number(digit), 0));
}

/** The class's guide builds as rail rows: by spec in tree order, the levelling build first. */
export function guideBuildRows(builds: readonly GuideBuild[], classSlug: string): GuideBuildRow[] {
  const specs = SPECS.filter((spec) => spec.class_slug === classSlug);
  const rank = (slug: string): number => {
    const at = specs.findIndex((spec) => spec.spec_slug === slug);
    return at === -1 ? specs.length : at;
  };
  return builds
    .filter((build) => build.classSlug === classSlug && codeTrees(build.code) !== null)
    .sort((a, b) => rank(a.specSlug) - rank(b.specSlug) || (a.kind === 'leveling' ? -1 : 1))
    .map((build) => {
      const spec = specs.find((row) => row.spec_slug === build.specSlug);
      const trees = codeTrees(build.code) ?? '';
      const fullSpec = `${build.classSlug}-${build.specSlug}`;
      const row: GuideBuildRow = {
        key: `${fullSpec}-${build.kind}`,
        name: spec?.name ?? build.specSlug,
        label: build.kind === 'leveling' ? guidesCopy.levelingBuild : guidesCopy.raidBuild,
        split: treesSplit(trees),
        href: loadBuildHref(build.code, fullSpec),
        trees,
      };
      if (spec !== undefined) row.icon = spec.icon;
      return row;
    });
}
