// The W3-C1 files (spec F9): the generator against the guides on disk and the contract fixtures
// under docs/contracts/fixtures/app. The Go contract test holds the fixtures to api/openapi.yaml;
// this holds the generator to the fixtures. Floors only: no count or id is pinned. The suites
// that read build data skip on a checkout without it.
import { existsSync, readdirSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';
import activeBuild from '../../data/active-build.json';
import { readSpecCatalog } from '../bis/load';
import {
  APP_FILE_BUDGETS,
  appChangelog,
  appGuideIndex,
  appGuides,
  appTiers,
  encodeAppFile,
} from './app-files';
import { markdownBlocks } from './changelog-file';
import { bulletLines, type AppGuideFile } from './guide-file';

const HERE = path.dirname(fileURLToPath(import.meta.url));
const ROOT = path.resolve(HERE, '../../../..');
const FIXTURES = path.join(ROOT, 'docs/contracts/fixtures/app');
const BUILD = activeBuild.build;
const hasBuild = existsSync(path.join(ROOT, 'data/builds', BUILD, 'talents'));

function fixture(name: string): Record<string, unknown> {
  return JSON.parse(readFileSync(path.join(FIXTURES, `${name}.json`), 'utf8')) as Record<string, unknown>;
}

function keysOf(value: unknown): string[] {
  return Object.keys(value as Record<string, unknown>).sort();
}

describe('markdown helpers', () => {
  it('reads bullets as plain text', () => {
    expect(bulletLines('- **Role:** Ranged\n- Uses `Lava Burst`\nprose')).toEqual([
      'Role: Ranged',
      'Uses Lava Burst',
    ]);
  });

  it('splits a changelog body into paragraphs, headings and items', () => {
    expect(markdownBlocks('One\nline.\n\n## Head\n- **a** item\n- b\n\nLast [link](https://x.gg).')).toEqual([
      { type: 'p', text: 'One line.' },
      { type: 'h', text: 'Head' },
      { type: 'li', text: 'a item' },
      { type: 'li', text: 'b' },
      { type: 'p', text: 'Last link.' },
    ]);
  });
});

describe('budgets', () => {
  it('fails a file over its gzipped budget', () => {
    const big = { text: Array.from({ length: 40_000 }, (_, i) => `${i * 7919}`).join('') };
    expect(() => encodeAppFile('guideIndex', 'guides/index.json', big)).toThrow(/over its 10240-byte budget/);
  });
});

describe.skipIf(!hasBuild)('W3-C1 generator', () => {
  const guides = appGuides(BUILD);
  const catalog = readSpecCatalog();

  it('writes one guide per spec guide, every one inside its budget', () => {
    expect(guides.length).toBeGreaterThan(0);
    for (const guide of guides) expect(() => encodeAppFile('guide', guide.spec, guide)).not.toThrow();
    expect(() => encodeAppFile('guideIndex', 'index', appGuideIndex(BUILD))).not.toThrow();
    expect(() => encodeAppFile('tiers', 'tiers', appTiers(BUILD))).not.toThrow();
    expect(APP_FILE_BUDGETS.guide).toBe(60 * 1024);
  });

  it('draws every spell and talent with its icon or an explicit null, and a rank on every rotation spell', () => {
    for (const guide of guides) {
      for (const spell of guide.rotation?.priority ?? []) {
        expect(spell.name, guide.spec).not.toBe('');
        if (spell.spell_id !== null) expect(spell.rank, `${guide.spec} ${spell.name}`).not.toBeUndefined();
      }
      for (const talent of guide.talents.key) expect(talent).toHaveProperty('icon');
    }
  });

  it('sums the build totals from the decoded trees, within 51 points', () => {
    for (const guide of guides) {
      for (const build of [guide.talents.leveling_build, guide.talents.raid_build]) {
        if (build === null) continue;
        const total = build.totals.reduce((sum, n) => sum + n, 0);
        expect(total, guide.spec).toBeGreaterThan(0);
        expect(total, guide.spec).toBeLessThanOrEqual(51);
        expect(build.tree_names).toHaveLength(build.totals.length);
      }
    }
  });

  it('names only the sections the guide has, in reading order', () => {
    const order = ['overview', 'talents', 'rotation', 'stats', 'gear', 'races'];
    for (const guide of guides) {
      expect(guide.sections.length, guide.spec).toBeGreaterThan(0);
      expect([...guide.sections].sort((a, b) => order.indexOf(a) - order.indexOf(b))).toEqual(guide.sections);
      if (guide.rotation === null) expect(guide.sections).not.toContain('rotation');
    }
  });

  it('names the raid-build talents the rotation casts among the key talents', () => {
    const elemental = guides.find((guide) => guide.spec === 'shaman-elemental')!;
    expect(elemental.talents.key.map((talent) => talent.name)).toContain('Lightning Mastery');
  });

  it('never hands the apps the client placeholder icon', () => {
    for (const guide of guides) {
      const spells = [
        ...(guide.rotation?.opener ?? []),
        ...(guide.rotation?.priority ?? []),
        ...(guide.rotation?.cooldowns ?? []),
      ];
      for (const spell of spells)
        expect(spell.icon, `${guide.spec} ${spell.name}`).not.toBe('inv_misc_questionmark');
    }
  });

  it('never repeats a talent text as its why', () => {
    for (const guide of guides) {
      for (const talent of guide.talents.key) {
        if (talent.why !== null && talent.text !== null)
          expect(
            talent.text.startsWith(talent.why),
            `${guide.spec} ${talent.name}: ${talent.why} || ${talent.text}`,
          ).toBe(false);
      }
    }
  });

  it('orders the index by classes.json and each class by its trees', () => {
    const index = appGuideIndex(BUILD);
    expect(index.classes.map((row) => row.class)).toHaveLength(9);
    for (const row of index.classes) {
      const trees = row.specs.map(
        (spec) => catalog.find((entry) => entry.spec === spec.spec)?.tree_index ?? -1,
      );
      expect(trees).toEqual([...trees].sort((a, b) => a - b));
    }
  });

  it('ranks the tier lists exactly as the web does, with the callout sentences on each row', () => {
    const tiers = appTiers(BUILD);
    const dps = tiers.factions.horde.dps.rows;
    expect(dps[0]?.gap_text).toBe('Top');
    expect(dps.map((row) => row.rank)).toEqual(dps.map((_, i) => i + 1));
    for (const row of dps)
      expect(row.title).toMatch(
        new RegExp(`^${row.name} is \\d+(st|nd|rd|th) of ${dps.length} DPS specs\\.$`),
      );
    expect(tiers.factions.horde.tank.rows.every((row) => row.tier === null)).toBe(true);
  });

  // The planner's golden vectors (apps/shared/fixtures/planner-rules.json) replay these codes, so a
  // fixture left on an old guide build tests the apps against a build no guide recommends. The
  // build id is dropped: a client move that keeps the trees leaves the vectors valid.
  it("carries each guide's current build codes in its fixture", () => {
    // Trailing zeros in a tree spend nothing: the move's re-encoder drops them, a hand-written
    // code may keep them (the 70938 move re-encoded Fire's 2301052001000 to 2301052001).
    const talentsOf = (code: string | undefined): string | null => {
      if (code === undefined) return null;
      const [cls, race, trees, ...rest] = code.replace(/^FS1:[^:]+:/, '').split(':');
      const trimmed = (trees ?? '').split('/').map((tree) => tree.replace(/0+$/, ''));
      return [cls, race, trimmed.join('/'), ...rest].join(':');
    };
    const named = readdirSync(FIXTURES)
      .map((file) => /^cdn\.guides\.(?!index\.)(.+)\.json$/.exec(file)?.[1])
      .filter((spec): spec is string => spec !== undefined);
    expect(named.length).toBeGreaterThan(0);
    for (const spec of named) {
      const guide = guides.find((candidate) => candidate.spec === spec);
      expect(guide, spec).toBeDefined();
      const want = fixture(`cdn.guides.${spec}`) as { talents: AppGuideFile['talents'] };
      for (const key of ['leveling_build', 'raid_build'] as const)
        expect(talentsOf(want.talents[key]?.code), `${spec} ${key}`).toBe(
          talentsOf(guide!.talents[key]?.code),
        );
    }
  });

  it('carries every field the fixtures name', () => {
    const guide = guides.find((candidate) => candidate.spec === 'shaman-elemental')!;
    const want = fixture('cdn.guides.shaman-elemental');
    expect(keysOf(guide)).toEqual(keysOf(want));
    expect(keysOf(guide.talents)).toEqual(keysOf(want.talents));
    expect(keysOf(appGuideIndex(BUILD))).toEqual(keysOf(fixture('cdn.guides.index')));
    expect(keysOf(appTiers(BUILD))).toEqual(keysOf(fixture('cdn.tiers')));
    const entry = appChangelog().entries[0]!;
    const wantEntry = (fixture('cdn.changelog').entries as unknown[])[0];
    expect(keysOf(entry)).toEqual(keysOf(wantEntry));
  });
});

describe('changelog', () => {
  it('is newest first', () => {
    const dates = appChangelog().entries.map((entry) => entry.date);
    expect(dates).toEqual([...dates].sort().reverse());
  });
});
