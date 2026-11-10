// web/src/components/icon-guard.test.ts
// One shared component per game icon and gear fact (component audit 2026-11-10, clusters 1, 6, 9,
// 25 and 27; owner ruling: the copies deleted, and a guard that fails if one reappears). It scans
// every .astro, .svelte and .ts file under src/ (tests, fixtures and generated data aside) and
// fails on:
//   - a hand-built game icon url (`icons/${stem}`, `/data/<build>/icons/`) outside lib/ui/assets.ts
//     (iconUrl, with the placeholder check built in);
//   - a game icon <img> outside components/ui/ (draw ui/SpellIcon, its .astro twin, ItemIcon or
//     ZoneIcon);
//   - a spell icon component outside ui/ (the deleted report, coach and sim copies), or an import
//     of one;
//   - the shared item tooltip's hook (`data-bis-item`) outside ItemName and the primitives that
//     own it;
//   - `raidZoneIcon(` outside ui/ZoneIcon;
//   - a spelled-out gear slot list, or a second average item level, outside lib/gear.
// Each rule's self-check below proves it catches the pattern it was written for.
import { readdirSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';

const SRC = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '..');

/** lane/talent-tree converts the talent tree's icons at the same time (2026-11-10); these three
 *  files are its own and keep their urls until it lands. Delete each entry as that lane does. */
const TALENT_TREE_HELD = ['components/ui/TalentTrees.svelte', 'components/planner/TalentCell.svelte'];
const TALENT_TREE_HELD_LIB = ['lib/planner/editor-trees.ts'];
/** The talent tree's own <img> (tree art, talent cells, tree heads), held for the same lane. */
const TALENT_TREE_HELD_IMG = [
  ...TALENT_TREE_HELD,
  'components/planner/TreeGrid.svelte',
  'components/planner/Planner.svelte',
];

interface Rule {
  /** What the offence is and what to use instead. */
  readonly name: string;
  readonly pattern: RegExp;
  /** Files (relative to src/) the pattern is allowed in, each for the reason beside it. */
  readonly allowed: (file: string) => boolean;
  /** Which files the rule reads. */
  readonly scans: (file: string) => boolean;
}

const MARKUP = (file: string): boolean => /\.(astro|svelte)$/.test(file);
const ANY = (): boolean => true;

const RULES: readonly Rule[] = [
  {
    name: 'a hand-built game icon url: use iconUrl or catalogueIconUrl (lib/ui/assets.ts)',
    pattern: /icons\/\$\{|\/data\/[^'"`\s]*\/icons\//g,
    allowed: (file) =>
      file === 'lib/ui/assets.ts' || [...TALENT_TREE_HELD, ...TALENT_TREE_HELD_LIB].includes(file),
    scans: ANY,
  },
  {
    name: 'a game icon <img> outside components/ui/: draw ui/SpellIcon, SpellIcon.astro, ItemIcon or ZoneIcon',
    pattern: /<img\b[^>]*?\bsrc=\{[^}]*icon[^}]*\}/gis,
    allowed: (file) => file.startsWith('components/ui/') || TALENT_TREE_HELD_IMG.includes(file),
    scans: MARKUP,
  },
  {
    name: 'a spell icon copy: import ui/SpellIcon (report/SpellIcon, coach/SpellIcon and sim/ActionIcon are deleted)',
    pattern:
      /import\s+\w+\s+from\s+'[^']*\/(?:report|coach)\/SpellIcon\.svelte'|import\s+\w+\s+from\s+'[^']*ActionIcon\.svelte'/g,
    allowed: () => false,
    scans: ANY,
  },
  {
    name: "the item tooltip's hook outside ItemName: draw ui/ItemName, ItemRow or ItemIcon in ItemHover",
    pattern: /data-bis-item=/g,
    allowed: (file) =>
      [
        // The primitives that own the hook: the name, the row, the doll, the icon's own hover.
        'components/ui/ItemName.svelte',
        'components/ui/ItemRow.svelte',
        'components/ui/PaperdollFrame.svelte',
        'components/ui/SpellIcon.svelte',
        // The host that opens the tooltip from the hook.
        'components/BisTooltipHost.svelte',
        // A sim action is a spell or an item: one keyboard stop opens whichever tooltip it is.
        'components/sim/ActionMark.svelte',
        'components/sim/ActionStatRow.svelte',
        // A guide consumable row's title (a static page, no quality to colour it by).
        'components/guides/GuideExtrasLane.astro',
      ].includes(file),
    scans: MARKUP,
  },
  {
    name: "a raid's icon outside ui/ZoneIcon: draw ZoneIcon",
    pattern: /\braidZoneIcon\(/g,
    allowed: (file) => file === 'components/ui/ZoneIcon.svelte' || file === 'lib/first-raids.ts',
    scans: ANY,
  },
  {
    name: 'a spelled-out gear slot list: read lib/gear/slots.ts',
    pattern: /'finger1',\s*'finger2',\s*'trinket1'/g,
    allowed: (file) =>
      file === 'lib/gear/slots.ts' ||
      // The combat log's own equipped-item layouts (retail and Forever), positions not an order.
      file === 'lib/report/planner-link.ts',
    scans: ANY,
  },
  {
    name: 'a second average item level: use lib/gear/paperdoll averageItemLevel (no shirt or tabard, no partial mean)',
    pattern: /function\s+average\w*ItemLevel\b/g,
    allowed: (file) => file === 'lib/gear/paperdoll.ts',
    scans: ANY,
  },
];

function relative(file: string): string {
  return path.relative(SRC, file).split(path.sep).join('/');
}

function scannedFiles(dir: string): string[] {
  return readdirSync(dir, { withFileTypes: true }).flatMap((entry) => {
    const full = path.join(dir, entry.name);
    if (entry.isDirectory())
      return entry.name === 'data' || entry.name === 'fixtures' || entry.name === 'test-support'
        ? []
        : scannedFiles(full);
    if (entry.name.includes('.test.')) return [];
    return /\.(astro|svelte|ts)$/.test(entry.name) ? [full] : [];
  });
}

function withoutComments(text: string): string {
  return text
    .replace(/\/\*[\s\S]*?\*\//g, '')
    .replace(/<!--[\s\S]*?-->/g, '')
    .replace(/^\s*\/\/.*$/gm, '');
}

/** The offences one file's text holds, by rule name. */
export function offencesIn(file: string, text: string): string[] {
  const code = withoutComments(text);
  return RULES.filter((rule) => rule.scans(file) && !rule.allowed(file))
    .filter((rule) => (code.match(rule.pattern)?.length ?? 0) > 0)
    .map((rule) => rule.name);
}

/** A spell icon component outside ui/ by its file name. */
export function strayIconComponent(file: string): boolean {
  return (
    /(?:^|\/)(?:SpellIcon|ActionIcon)\.(?:svelte|astro)$/.test(file) && !file.startsWith('components/ui/')
  );
}

describe('the icon guard', () => {
  const files = scannedFiles(SRC).map((full) => ({ file: relative(full), full }));

  it('no file builds its own game icon, item name, zone icon, slot list or item level', () => {
    const found = files.flatMap(({ file, full }) =>
      offencesIn(file, readFileSync(full, 'utf8')).map((name) => `${file}: ${name}`),
    );
    expect(found).toEqual([]);
  });

  it('the one spell icon is ui/SpellIcon and its .astro twin', () => {
    expect(files.map(({ file }) => file).filter(strayIconComponent)).toEqual([]);
  });

  it('catches every pattern it was written for', () => {
    const page = 'components/report/Example.svelte';
    const lib = 'lib/report/example.ts';
    expect(offencesIn(lib, 'return dataUrl(build, `icons/${stem}.webp`);')).toHaveLength(1);
    expect(offencesIn(page, '<img src={`/data/${build}/icons/x.webp`} alt="" />')).toHaveLength(1);
    expect(offencesIn(page, '<img\n  src={row.iconUrl}\n  alt=""\n/>')).toHaveLength(1);
    expect(offencesIn(page, "import SpellIcon from './../coach/SpellIcon.svelte';")).toHaveLength(1);
    expect(offencesIn(page, "import ActionIcon from '../sim/ActionIcon.svelte';")).toHaveLength(1);
    expect(offencesIn(page, '<span class="q4" data-bis-item={id}>{name}</span>')).toHaveLength(1);
    expect(offencesIn(page, '<ItemIcon icon={raidZoneIcon(zone)} />')).toHaveLength(1);
    expect(offencesIn(lib, "const ORDER = ['finger1', 'finger2', 'trinket1'];")).toHaveLength(1);
    expect(offencesIn(lib, "const BOTH = { finger: ['finger1', 'finger2'] };")).toEqual([]);
    expect(offencesIn(lib, 'export function averageItemLevel(gear) {}')).toHaveLength(1);
    expect(strayIconComponent('components/report/SpellIcon.svelte')).toBe(true);
    // What the shared components themselves write is allowed.
    expect(offencesIn('lib/ui/assets.ts', 'dataUrl(build, `icons/${stem}.webp`)')).toEqual([]);
    expect(offencesIn('components/ui/SpellIcon.svelte', '<img src={shown} alt="" />')).toEqual([]);
    expect(offencesIn(page, '<SpellIcon url={iconUrl(build, stem)} size={24} />')).toEqual([]);
    expect(strayIconComponent('components/ui/SpellIcon.astro')).toBe(false);
  });
});
