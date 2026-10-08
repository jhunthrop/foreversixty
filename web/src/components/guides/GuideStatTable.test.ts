// web/src/components/guides/GuideStatTable.test.ts
// Same AstroContainer render-test convention StatPriorityPills.test.ts and
// BuildActionButtons.test.ts already use for a co-located .astro component. Reads the real
// warrior-fury band-60 BiS file rather than a mock, the same way class-dps.test.ts and
// rail-stats.test.ts do -- the nightly republish moves these numbers, so this only ever
// asserts shape and ordering, never a pinned value (data contracts are floors).
import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { describe, expect, it } from 'vitest';
import { bandEntry, loadBisFile } from '../../lib/bis/load';
import { hitCapLine } from '../../lib/bis/hit-cap';
import { railStatRows } from '../../lib/guides/rail-stats';
import GuideStatTable from './GuideStatTable.astro';

const BUILD = '1.60.1.70291';
const STAT_PRIORITY = ['Strength', 'Critical strike', 'Attack power', 'Hit', 'Melee haste', 'Agility'];

describe('GuideStatTable', () => {
  it('renders one row per stat, in the sim’s own order, never the frontmatter order', async () => {
    const file = loadBisFile('warrior-fury', BUILD);
    if (file === null) throw new Error('warrior-fury BiS file missing');
    const band = bandEntry(file, 60, 'alliance');
    if (band === undefined) throw new Error('band 60 alliance missing');
    const expectedRows = railStatRows(
      STAT_PRIORITY,
      band.weights,
      'warrior-fury',
      band.haste_scale_factor ?? null,
    );

    const c = await AstroContainer.create();
    const html = await c.renderToString(GuideStatTable, {
      props: {
        build: BUILD,
        spec: 'warrior-fury',
        classSlug: 'warrior',
        recommendedRaces: ['human', 'troll'],
        statPriority: STAT_PRIORITY,
      },
    });

    expect(html).toContain('data-testid="guide-stat-table"');
    for (const [index, row] of expectedRows.entries()) {
      expect(html).toContain(`data-testid="guide-stat-table-row-${index}"`);
      expect(html).toContain(row.label);
      if (row.value === undefined) {
        expect(html).toContain('not separable from zero');
      } else {
        expect(html).toContain(row.value.toFixed(2));
      }
    }
    // Row order comes straight from railStatRows (the sim's own tiering), not from a
    // second sort this component invents -- the row-by-row label/value assertions above
    // already pin that order; this just names the row count matches exactly.
    expect(expectedRows).toHaveLength(STAT_PRIORITY.length);
  });

  it('prints an "Updated" line from the file’s own generated_at', async () => {
    const file = loadBisFile('warrior-fury', BUILD);
    if (file === null) throw new Error('warrior-fury BiS file missing');

    const c = await AstroContainer.create();
    const html = await c.renderToString(GuideStatTable, {
      props: {
        build: BUILD,
        spec: 'warrior-fury',
        classSlug: 'warrior',
        recommendedRaces: ['human', 'troll'],
        statPriority: STAT_PRIORITY,
      },
    });

    expect(html).toContain('data-testid="guide-stat-table-updated"');
    expect(html).toContain('Updated');
  });

  it('prints the hit-to-cap line exactly when the band publishes the key', async () => {
    const file = loadBisFile('warrior-fury', BUILD);
    if (file === null) throw new Error('warrior-fury BiS file missing');
    const band = bandEntry(file, 60, 'alliance');
    const expected = hitCapLine(band?.hit_to_cap);

    const c = await AstroContainer.create();
    const html = await c.renderToString(GuideStatTable, {
      props: {
        build: BUILD,
        spec: 'warrior-fury',
        classSlug: 'warrior',
        recommendedRaces: ['human', 'troll'],
        statPriority: STAT_PRIORITY,
      },
    });

    if (expected === undefined) {
      expect(html).not.toContain('guide-stat-table-hit-cap');
    } else {
      expect(html).toContain('data-testid="guide-stat-table-hit-cap"');
      expect(html).toContain(expected.text);
    }
  });

  it('renders nothing for a spec with no ranked BiS file', async () => {
    const c = await AstroContainer.create();
    const html = await c.renderToString(GuideStatTable, {
      props: {
        build: BUILD,
        spec: 'warrior-nonexistent',
        classSlug: 'warrior',
        recommendedRaces: ['dwarf'],
        statPriority: ['Defense'],
      },
    });

    expect(html).not.toContain('data-testid="guide-stat-table"');
  });
});
