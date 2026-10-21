import { describe, expect, it } from 'vitest';
import { buildMockHome, VIEWER_MEMBER } from '../../fixtures/guild/mock-guild';
import { entryFacts, gearSetLabel, inGameFacts, rosterStatChips } from './roster-facts';
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
  it('always runs Level, Gear gap, spec, gear set, seen', () => {
    const keys = inGameFacts(rowNamed('Zulmara'), NOW).map((f) => f.key);
    expect(keys).toEqual(['level', 'gap', 'spec', 'build', 'seen']);
    expect(inGameFacts(rowNamed('Kraggor'), NOW).map((f) => f.key)).toEqual([
      'level',
      'gap',
      'build',
      'seen',
    ]);
  });

  it('says the in-game spec as "<Spec> in game" and the gear set with its kind in brackets', () => {
    const texts = inGameFacts(rowNamed('Zulmara'), NOW).map((f) => f.text);
    expect(texts).toContain('Survival in game');
    expect(texts).toContain('Gear set: Survival (raid)');
    expect(gearSetLabel('Protection raid')).toBe('Gear set: Protection (raid)');
    expect(gearSetLabel('Frost')).toBe('Gear set: Frost');
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
