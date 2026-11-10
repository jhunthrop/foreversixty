// web/src/lib/guides/stat-lane.test.ts
// The guide's one stat model (the At a glance column and the Stat priority lane read it). Ported
// from the deleted GuideStatTable.test.ts: it reads the real band 60 BiS files rather than a
// mock, so it asserts shape and order, never a pinned figure (data contracts are floors).
import { describe, expect, it } from 'vitest';
import activeBuild from '../../data/active-build.json';
import { bandEntry, loadBisFile } from '../bis/load';
import { guideExpertiseCapLine, guideHitCapLine } from './hit-cap-line';
import { railStatRows } from './rail-stats';
import { guideStatLane, perPointText } from './stat-lane';
import { bandTalentHits } from './talent-hit';

const BUILD = '1.60.1.70374';
const STAT_PRIORITY = ['Strength', 'Critical strike', 'Attack power', 'Hit', 'Melee haste', 'Agility'];

function furyLane() {
  const lane = guideStatLane(BUILD, 'warrior', 'warrior-fury', ['human', 'troll'], STAT_PRIORITY);
  if (lane === undefined) throw new Error('the Fury band 60 weights are missing');
  return lane;
}

function furyBand() {
  const file = loadBisFile('warrior-fury', BUILD);
  if (file === null) throw new Error('warrior-fury BiS file missing');
  const band = bandEntry(file, 60, 'alliance');
  if (band === undefined) throw new Error('band 60 alliance missing');
  return band;
}

describe('perPointText', () => {
  it('prints two decimals under 10, one at 10 or more, and nothing without a figure', () => {
    expect(perPointText(0.456)).toBe('0.46');
    expect(perPointText(12.34)).toBe('12.3');
    expect(perPointText(undefined)).toBe('');
  });
});

describe('guideStatLane', () => {
  it('lists one row per stat in the sim’s own order, never the frontmatter order', () => {
    const band = furyBand();
    const expected = railStatRows(STAT_PRIORITY, band.weights, 'warrior-fury', {
      scaleFactor: band.haste_scale_factor ?? null,
      onItems: band.haste_on_items ?? true,
    });
    const lane = furyLane();
    expect(expected).toHaveLength(STAT_PRIORITY.length);
    expect(lane.rows.map((row) => row.label)).toEqual(expected.map((row) => row.label));
    for (const row of lane.rows) {
      if (row.reading === 'not-on-items') expect(row.valueText).toBe('Not on any item');
      else if (row.value === undefined) expect(['Not significant', 'No gain']).toContain(row.valueText);
      else if (row.perPercent) expect(row.valueText).toContain(`${row.value.toFixed(1)} per 1%`);
      else expect(row.valueText).toBe(row.value.toFixed(2));
      expect(row.perPointText).toBe(perPointText(row.perPoint));
    }
  });

  it('heads the weight column on the anchor stat and the per-point column on the role rate', () => {
    const lane = furyLane();
    expect(lane.role).toBe('dps');
    expect(lane.weightHead).toMatch(/^Weight \([A-Za-z ]+ = 1\.00\)$/);
    // Fury prices hit and crit per rating point (round 20).
    expect(lane.perPointHead).toBe('DPS per point');
    expect(lane.ratingNote).toMatch(/per rating point: .*hit rating is 1% hit/);
  });

  it('names the top per-point stat in the unit line', () => {
    expect(furyLane().unitLine).toMatch(
      /^1 (?:point of [a-z][a-z ]*|[a-z][a-z ]* rating) = (?:\d\.\d\d|\d{2,}\.\d) DPS at level 60\.?$/,
    );
  });

  it('carries the hit cap line exactly when the band publishes it, with the talents that add hit', () => {
    const band = furyBand();
    const expected = guideHitCapLine(band.hit_to_cap, bandTalentHits(BUILD, 'warrior', band.talents));
    expect(furyLane().hitCap).toEqual(expected);
    const expertise = guideExpertiseCapLine(
      band.expertise_to_cap,
      bandTalentHits(BUILD, 'warrior', band.talents),
    );
    expect(furyLane().expertiseCap).toEqual(expertise);
  });

  it('gives a tank its tank-score unit and no Expertise when nothing gives it', () => {
    const lane = guideStatLane(
      BUILD,
      'warrior',
      'warrior-protection',
      ['human'],
      ['Expertise', 'Stamina', 'Defense', 'Dodge', 'Parry', 'Agility'],
    );
    expect(lane?.role).toBe('tank');
    expect(lane?.perPointHead).toMatch(/^Tank score per point/);
    expect(lane?.unitLine ?? '').not.toContain('DPS');
    expect(lane?.rows.map((row) => row.label)).not.toContain('Expertise');
    expect(lane?.expertiseCap).toBeUndefined();
  });

  it('gives a healer HPS and no negative figure', () => {
    const lane = guideStatLane(
      BUILD,
      'priest',
      'priest-holy',
      ['human'],
      ['Healing power', 'Intellect', 'Spirit', 'MP5', 'Crit', 'Spell haste'],
    );
    expect(lane?.perPointHead).toMatch(/^HPS per point/);
    for (const row of lane?.rows ?? []) expect(row.valueText).not.toMatch(/^-\d/);
  });

  it('notes every row with no figure, and keeps "Not on any item" out of the plate’s rows', () => {
    const lane = guideStatLane(
      activeBuild.build,
      'mage',
      'mage-frost',
      ['gnome', 'undead'],
      [
        'Hit',
        'Spell power',
        'Frost power',
        'Critical strike',
        'Intellect',
        'Spell haste',
        'Spell penetration',
      ],
    );
    if (lane === undefined) throw new Error('the Frost band 60 weights are missing');
    const offItems = lane.rows.filter((row) => row.reading === 'not-on-items');
    const quiet = lane.rows.filter((row) => row.reading === 'not-significant' || row.reading === 'no-gain');
    expect(lane.notes.length).toBeGreaterThanOrEqual(offItems.length + (quiet.length > 0 ? 1 : 0));
    for (const row of offItems) {
      expect(lane.notes).toContain(
        `${row.label}: ${row.value!.toFixed(1)} DPS per 1%, but no item at level 60 has it.`,
      );
    }
  });

  it('is undefined for a spec with no ranked BiS file', () => {
    expect(guideStatLane(BUILD, 'warrior', 'warrior-nonexistent', ['dwarf'], ['Defense'])).toBeUndefined();
  });
});
