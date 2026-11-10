// web/src/components/ui/talent-tree-guard.test.ts
// One talent tree site-wide (talent-tree lane 2026-11-10, design/reviews/2026-11-10-talent-tree):
// every tree on every page is ui/TalentTrees, so it looks, opens its tooltip and takes the keyboard
// the same everywhere. No component outside the shared three draws a talent grid, a talent cell
// or a prerequisite arrow, and none mounts the cell or the single tree past TalentTrees. The
// allowlist is empty: a page that needs a tree uses TalentTrees in one of its modes.
import { readFileSync, readdirSync, statSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';

const SRC = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '../..');

/** The shared tree itself. */
const SHARED = new Set([
  'components/ui/TalentTrees.svelte',
  'components/ui/TalentTree.svelte',
  'components/ui/TalentButton.svelte',
  'components/ui/TalentTreesWithTips.svelte',
]);
/** Components allowed to draw a talent grid of their own: none. */
const ALLOWLIST = new Set<string>();

/** What drawing a talent tree by hand looks like: the tree's grid or cell classes, a talent's
 *  grid named for its talents, a talent icon built in markup, the arrow geometry, the grid layout, the inner parts. */
const TREE_PATTERNS: readonly [string, RegExp][] = [
  ['the shared cell or grid classes', /class=["{][^"}]*\btt-(?:cell|g|wrap|rk|arrows|art)\b/],
  ['a talent grid role', /role="grid"[^>]*talents|aria-label=\{`\$\{[^}]+\} talents`\}/],
  ['the prerequisite arrows', /\bconnectorsFor\s*\(|\bconnectorPath\s*\(/],
  ['the tree grid layout', /\bgridCells\s*\(|\bgridRows\s*\(/],
  ['a talent icon built in markup', /icons\/\$\{(?:talent|cell)\.icon\}/],
  ['the single tree or cell mounted directly', /import\s+Talent(?:Tree|Button)\s+from/],
];

function* componentFiles(dir: string): Generator<string> {
  for (const name of readdirSync(dir)) {
    const full = path.join(dir, name);
    if (statSync(full).isDirectory()) yield* componentFiles(full);
    else if (/\.(svelte|astro)$/.test(name)) yield full;
  }
}

function offenders(): string[] {
  return [...componentFiles(SRC)]
    .map((file) => path.relative(SRC, file))
    .filter((file) => !SHARED.has(file) && !ALLOWLIST.has(file))
    .flatMap((file) => {
      const source = readFileSync(path.join(SRC, file), 'utf8');
      return TREE_PATTERNS.filter(([, pattern]) => pattern.test(source)).map(([what]) => `${file}: ${what}`);
    });
}

describe('one talent tree', () => {
  it('no component outside ui/TalentTrees draws a talent grid', () => {
    expect(offenders()).toEqual([]);
  });

  it('keeps the allowlist empty', () => {
    expect([...ALLOWLIST]).toEqual([]);
  });

  it('catches a hand-drawn tree', () => {
    const drawn =
      '<div role="grid" aria-label={`${tree.name} talents`}><span class="tt-cell max"><img src={`icons/${talent.icon}.webp`} />';
    const caught = TREE_PATTERNS.filter(([, pattern]) => pattern.test(drawn)).map(([what]) => what);
    expect(caught).toEqual([
      'the shared cell or grid classes',
      'a talent grid role',
      'a talent icon built in markup',
    ]);
  });
});
