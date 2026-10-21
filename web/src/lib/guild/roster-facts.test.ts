import { describe, expect, it } from 'vitest';
import { buildMockHome, VIEWER_MEMBER } from '../../fixtures/guild/mock-guild';
import { entryFacts, gearSetLabel, inGameFacts, rosterSpecLine, rosterStatChips } from './roster-facts';
import { specClassName } from './spec-name';

const NOW = new Date(2026, 11, 22, 21, 46);

function rowNamed(name: string) {
  const row = buildMockHome(VIEWER_MEMBER).roster.find((r) => r.name === name);
  if (row === undefined) throw new Error(`no fixture row ${name}`);
  return row;
}

describe('rosterStatChips', () => {
  it('names the unit first on every chip', () => {
    const chips = rosterStatChips(rowNamed('Kraggor'));
    expect(chips[0]).toMatch(/^Item level \d+$/);
    expect(chips.some((c) => /^\d+ of \d+ nights$/.test(c))).toBe(true);
    expect(chips.some((c) => /^Best (DPS|HPS) \d+$/.test(c))).toBe(true);
    expect(chips.some((c) => /^Rating \d+$/.test(c))).toBe(true);
    expect(chips.join(' ')).not.toMatch(/\d+ rating|ilvl|best DPS/);
  });
});

describe('inGameFacts', () => {
  it('always runs Level, Gear gap, gear set, seen: the spec has its own line', () => {
    const keys = inGameFacts(rowNamed('Zulmara'), NOW).map((f) => f.key);
    expect(keys).toEqual(['level', 'gap', 'build', 'seen']);
    expect(inGameFacts(rowNamed('Kraggor'), NOW).map((f) => f.key)).toEqual([
      'level',
      'gap',
      'build',
      'seen',
    ]);
  });

  it('says the gear set with its kind in brackets, and leaves the spec to its own line', () => {
    const texts = inGameFacts(rowNamed('Zulmara'), NOW).map((f) => f.text);
    expect(texts.some((t) => t.endsWith('in game'))).toBe(false);
    expect(texts).toContain('Gear set: Survival (raid)');
    expect(gearSetLabel('Protection raid')).toBe('Gear set: Protection (raid)');
    expect(gearSetLabel('Frost')).toBe('Gear set: Frost');
  });

  it('drops the Gear gap fact when it is 0, and keeps it when it is a real gap', () => {
    const row = rowNamed('Zulmara');
    const withGap = (gap: number) =>
      inGameFacts({ ...row, addon: { ...(row.addon as NonNullable<typeof row.addon>), gear_gap: gap } }, NOW);
    expect(withGap(0).map((f) => f.key)).not.toContain('gap');
    expect(withGap(12).map((f) => f.text)).toContain('Gear gap 12 DPS');
    expect(withGap(-3).map((f) => f.text)).toContain('Gear gap -3 DPS');
  });

  it('a member who withholds everything shows only seen', () => {
    expect(inGameFacts(rowNamed('Duskstrider'), NOW).map((f) => f.key)).toEqual(['seen']);
  });

  it('a name only the addon saw reads Level then seen', () => {
    const entry = buildMockHome(VIEWER_MEMBER).members?.find((m) => m.state === 'no_account');
    if (entry === undefined) return;
    expect(entryFacts(entry, NOW).map((f) => f.key)).toEqual(
      entry.level === undefined ? ['seen'] : ['level', 'seen'],
    );
  });
});

describe('specClassName', () => {
  it('carries the class on the spec', () => {
    expect(specClassName('Holy', 'priest')).toBe('Holy Priest');
    expect(specClassName('Restoration', 'druid')).toBe('Restoration Druid');
    expect(specClassName('BeastMastery', 'hunter')).toBe('Beast Mastery Hunter');
    expect(specClassName('Holy Priest', 'priest')).toBe('Holy Priest');
    expect(specClassName('', 'priest')).toBe('Priest');
    expect(specClassName('Holy', undefined)).toBe('Holy');
  });
});

// One spec per character, labelled with the record it comes from (lifecycle round 8, P8-1).
describe('rosterSpecLine', () => {
  it('says "in game" alone when the snapshot and the site agree', () => {
    expect(rosterSpecLine({ siteSpec: 'BeastMastery', addonSpec: 'BeastMastery', className: 'hunter' })).toBe(
      'Beast Mastery Hunter in game',
    );
    expect(rosterSpecLine({ addonSpec: 'Holy', className: 'priest' })).toBe('Holy Priest in game');
  });

  it('adds the sites spec when it disagrees with the snapshot', () => {
    expect(rosterSpecLine({ siteSpec: 'Survival', addonSpec: 'BeastMastery', className: 'hunter' })).toBe(
      'Beast Mastery Hunter in game · Survival on the site',
    );
  });

  it('reads the sites spec, labelled, when the snapshot carries none', () => {
    expect(rosterSpecLine({ siteSpec: 'Survival', addonSpec: null, className: 'hunter' })).toBe(
      'Survival Hunter on the site',
    );
    expect(rosterSpecLine({ siteSpec: 'Survival', className: 'hunter' })).toBe('Survival Hunter on the site');
    expect(rosterSpecLine({ addonSpec: null, className: 'hunter' })).toBe('');
  });
});
