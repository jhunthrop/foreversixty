// web/src/components/guides/GuideStatTable.test.ts
// Same AstroContainer render-test convention StatPriorityPills.test.ts and
// BuildActionButtons.test.ts already use for a co-located .astro component. Reads the real
// warrior-fury band-60 BiS file rather than a mock, the same way class-dps.test.ts and
// rail-stats.test.ts do -- the nightly republish moves these numbers, so this only ever
// asserts shape and ordering, never a pinned value (data contracts are floors).
import { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { describe, expect, it } from 'vitest';
import { bandEntry, loadBisFile } from '../../lib/bis/load';
import { guideExpertiseCapLine } from '../../lib/guides/hit-cap-line';
import { guideHitCapLine } from '../../lib/guides/hit-cap-line';
import { railStatRows } from '../../lib/guides/rail-stats';
import GuideStatTable from './GuideStatTable.astro';

const BUILD = '1.60.1.70374';
const STAT_PRIORITY = ['Strength', 'Critical strike', 'Attack power', 'Hit', 'Melee haste', 'Agility'];

describe('GuideStatTable', () => {
  it('renders one row per stat, in the sim’s own order, never the frontmatter order', async () => {
    const file = loadBisFile('warrior-fury', BUILD);
    if (file === null) throw new Error('warrior-fury BiS file missing');
    const band = bandEntry(file, 60, 'alliance');
    if (band === undefined) throw new Error('band 60 alliance missing');
    const expectedRows = railStatRows(STAT_PRIORITY, band.weights, 'warrior-fury', {
      scaleFactor: band.haste_scale_factor ?? null,
      onItems: band.haste_on_items ?? true,
    });

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
      if (row.reading === 'not-on-items') {
        expect(html).toContain('Not on any item');
        expect(html).toContain(
          `${row.label}: ${row.value!.toFixed(1)} DPS per 1%, but no item at level 60 has it.`,
        );
      } else if (row.value === undefined) {
        expect(html).toContain('Not significant');
      } else if (row.perPercent) {
        expect(html).toContain(`${row.value.toFixed(1)} per 1% haste`);
      } else {
        expect(html).toContain(row.value.toFixed(2));
      }
    }
    // Row order comes straight from railStatRows (the sim's own tiering), not from a
    // second sort this component invents -- the row-by-row label/value assertions above
    // already pin that order; this just names the row count matches exactly.
    expect(expectedRows).toHaveLength(STAT_PRIORITY.length);
  });

  it('has the weight and DPS-per-point columns, and no "Updated" line (the fact line carries the date)', async () => {
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

    // Fury prices hit and crit per rating point, so the column says so (round 20).
    expect(html).toContain('DPS per point or rating');
    expect(html).toMatch(/Weight \([A-Za-z ]+ = 1\.00\)/);
    expect(html).not.toContain('guide-stat-table-updated');
  });

  it('sets the unit and rating lines apart from the muted notes (round 20)', async () => {
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

    expect(html).toMatch(/class="stat-unit" data-testid="guide-stat-table-unit"/);
    expect(html).toMatch(/class="stat-unit" data-testid="guide-stat-table-ratings"/);
    expect(html).not.toMatch(/class="stat-note" data-testid="guide-stat-table-unit"/);
  });

  it('puts the hit cap on the headline and the breakdown on its own detail line (round 20)', async () => {
    const file = loadBisFile('warrior-fury', BUILD);
    if (file === null) throw new Error('warrior-fury BiS file missing');
    const expected = guideHitCapLine(bandEntry(file, 60, 'alliance')?.hit_to_cap);
    if (expected === undefined) throw new Error('the fixture Fury band publishes a hit cap');

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

    expect(html).toContain('data-testid="guide-stat-table-hit-cap"');
    expect(html).toContain(expected.cap);
    expect(html).not.toMatch(/data-testid="guide-stat-table-hit-cap">[^<]*·/);
    if (expected.detail === undefined) {
      expect(html).not.toContain('guide-stat-table-hit-cap-detail');
    } else {
      expect(html).toContain('data-testid="guide-stat-table-hit-cap-detail"');
      expect(html).toContain(expected.detail);
    }
    // The split form, with the breakdown, is what the fixture carries for a mage.
    const arcane = guideHitCapLine({ kind: 'spell', baseline: 6, spell: 10 }, [
      { name: 'Precision', scope: 'spell', percent: 6 },
    ]);
    expect(arcane?.detail).toBe('6% from Precision, 10% still to find');
  });

  it('gives a tank its tank-score unit, never DPS, and no Expertise when nothing gives it', async () => {
    const c = await AstroContainer.create();
    const html = await c.renderToString(GuideStatTable, {
      props: {
        build: BUILD,
        spec: 'warrior-protection',
        classSlug: 'warrior',
        recommendedRaces: ['human'],
        statPriority: ['Expertise', 'Stamina', 'Defense', 'Dodge', 'Parry', 'Agility'],
      },
    });
    expect(html).toContain('Tank score per point');
    expect(html).not.toContain('DPS');
    // The level 60 Protection band's character has no Expertise (no Forever item carries it),
    // so the weights, the table and the cap line all leave it out together.
    expect(html).not.toMatch(/per 1% expertise/);
    expect(html).not.toMatch(/>Expertise</);
    expect(html).not.toContain('Expertise cap');
  });

  it('gives a healer HPS, and prints no negative weight', async () => {
    const c = await AstroContainer.create();
    const html = await c.renderToString(GuideStatTable, {
      props: {
        build: BUILD,
        spec: 'priest-holy',
        classSlug: 'priest',
        recommendedRaces: ['human'],
        statPriority: ['Healing power', 'Intellect', 'Spirit', 'MP5', 'Crit', 'Spell haste'],
      },
    });
    expect(html).toContain('HPS per point');
    expect(html).not.toContain('DPS');
    expect(html).not.toMatch(/>\s*-\d/);
  });

  it('prints the hit-to-cap line exactly when the band publishes the key', async () => {
    const file = loadBisFile('warrior-fury', BUILD);
    if (file === null) throw new Error('warrior-fury BiS file missing');
    const band = bandEntry(file, 60, 'alliance');
    const expected = guideHitCapLine(band?.hit_to_cap);

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

  it('prints the expertise-to-cap line exactly when the band publishes the key', async () => {
    const file = loadBisFile('warrior-fury', BUILD);
    if (file === null) throw new Error('warrior-fury BiS file missing');
    const expected = guideExpertiseCapLine(bandEntry(file, 60, 'alliance')?.expertise_to_cap);

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
      expect(html).not.toContain('guide-stat-table-expertise-cap');
    } else {
      expect(html).toContain('data-testid="guide-stat-table-expertise-cap"');
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
