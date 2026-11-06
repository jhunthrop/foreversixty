// web/src/lib/tiers/tier-copy.test.ts
import { describe, expect, it } from 'vitest';
import {
  caption,
  formatMetric,
  gapText,
  ordinal,
  raceBasis,
  readColumns,
  rowName,
  rowSub,
  tieText,
  tieTitle,
  tiersCopy,
  type ReadColumn,
} from './tier-copy';
import type { TierRow, TierTieMate } from './tier-list';

const flat = (column: ReadColumn): string => column.body.map((segment) => segment.text).join('');
const CTX = { faction: 'horde', presetLabel: 'Raid-ready, Phase 1' } as const;
const mate = (name: string, classSlug: string, tier: TierTieMate['tier'] = 'B'): TierTieMate => ({
  classSlug,
  specSlug: name.toLowerCase(),
  name,
  tier,
});
const dpsRow = (tier: TierRow['tier'], tiedWith: TierTieMate[]) => ({ role: 'dps' as const, tier, tiedWith });

describe('ordinal', () => {
  it.each([
    [1, '1st'],
    [2, '2nd'],
    [3, '3rd'],
    [4, '4th'],
    [11, '11th'],
    [12, '12th'],
    [21, '21st'],
    [112, '112th'],
  ])('%i is %s', (n, expected) => {
    expect(ordinal(n)).toBe(expected);
  });
});

describe("the race label (the race is the guide's, not simmed)", () => {
  it("names the guide's recommended race of the faction, never the best one", () => {
    expect(raceBasis('horde')).toBe("the guide's recommended Horde race");
    for (const role of ['dps', 'tank', 'healer'] as const) {
      const text = readColumns({ ...CTX, role, healSeconds: 300, bossSwingSeconds: 2 })
        .map(flat)
        .join(' ');
      expect(text).toContain("as the guide's recommended Horde race, named under it.");
      expect(text).not.toMatch(/best (Horde|Alliance) race/);
    }
  });
});

describe('readColumns', () => {
  it('gives each role three labelled columns, linking the leveling bands', () => {
    for (const role of ['dps', 'tank', 'healer'] as const) {
      const columns = readColumns({ ...CTX, role, healSeconds: 300, bossSwingSeconds: 2 });
      expect(columns).toHaveLength(3);
      expect(columns.flatMap((c) => c.body).some((s) => s.href === '/bis')).toBe(true);
    }
    expect(readColumns({ ...CTX, role: 'dps' }).map((c) => c.label)).toEqual([
      'What the number is',
      'What it leaves out',
      'Letters and ties',
    ]);
    expect(readColumns({ ...CTX, role: 'tank', bossSwingSeconds: 2 }).map((c) => c.label)).toEqual([
      'What the number is',
      'The boss',
      'The gear',
    ]);
  });

  it('links the log rankings from what the DPS number leaves out', () => {
    const leavesOut = readColumns({ ...CTX, role: 'dps' })[1]!;
    expect(leavesOut.body.find((s) => s.href === '/rankings')?.text).toBe('the log rankings');
  });

  it('states the letters and the tie rule from the rules, with no "box below"', () => {
    const text = readColumns({ ...CTX, role: 'dps' })
      .map(flat)
      .join(' ');
    expect(text).toContain('S is within 5% of the top, A within 10%, B within 20%, C within 30%, D beyond.');
    expect(text).toContain('≈ tie names the specs within 1% of a spec');
    expect(text).not.toContain('box below');
  });

  it('renders the healer duration and the tank swing from the data it is given', () => {
    expect(flat(readColumns({ ...CTX, role: 'healer', healSeconds: 300 })[0]!)).toContain('over 300 seconds');
    expect(flat(readColumns({ ...CTX, role: 'tank', bossSwingSeconds: 2 })[1]!)).toContain('every 2 seconds');
  });

  it('says the healer order has no level-60 log behind it yet, and how a shared rank counts', () => {
    const healer = readColumns({ ...CTX, role: 'healer', healSeconds: 300 })
      .map(flat)
      .join(' ');
    expect(healer).toContain('No level-60 log is behind this order yet');
    expect(healer).toContain('Tied healers share a rank: =2, =2, then 4.');
    const unmeasured = readColumns({ ...CTX, role: 'healer', healSeconds: 300, healerOrderMeasured: false })
      .map(flat)
      .join(' ');
    expect(unmeasured).toContain('Ranks appear once the next ranking publishes the measured error.');
    expect(unmeasured).not.toContain('≈ tie');
  });

  it('fails fast when a profile figure is missing', () => {
    expect(() => readColumns({ ...CTX, role: 'healer' })).toThrow(/heal profile duration/);
    expect(() => readColumns({ ...CTX, role: 'tank' })).toThrow(/boss swing speed/);
  });
});

describe('caption', () => {
  it('says what each role is measured on, from the constants and the data', () => {
    expect(caption({ ...CTX, role: 'dps' })).toBe(
      "Damage per second, one target for 180 seconds. The bar is the share of the top spec's damage; ≈ tie is within 1%.",
    );
    expect(caption({ ...CTX, role: 'tank' })).toContain('stand-in level 63 boss, lower is better');
  });

  it('explains ≈ tie on the healer list only when a tie exists', () => {
    expect(caption({ ...CTX, role: 'healer', healSeconds: 300, hasTie: true })).toBe(
      'Effective healing per second over 300 seconds of a stand-in Phase 1 fight. ≈ tie: the sim cannot order the pair.',
    );
    expect(caption({ ...CTX, role: 'healer', healSeconds: 300, hasTie: false })).not.toContain('≈ tie');
  });
});

describe('row names', () => {
  it('names a DPS spec alone and a healer or tank in full, with the class moved off the sub line', () => {
    expect(rowName({ name: 'Fury', classSlug: 'warrior', role: 'dps' })).toBe('Fury');
    expect(rowName({ name: 'Holy', classSlug: 'priest', role: 'healer' })).toBe('Holy Priest');
    expect(rowName({ name: 'Restoration', classSlug: 'shaman', role: 'healer' })).toBe('Restoration Shaman');
    expect(rowSub({ classSlug: 'warrior', role: 'dps', race: 'Night Elf' })).toBe('Warrior · Night Elf');
    expect(rowSub({ classSlug: 'priest', role: 'healer', race: 'Dwarf' })).toBe('Dwarf');
  });
});

describe('tie text (owner MUST 1)', () => {
  it('names the one partner: "≈ tie with Arms"', () => {
    expect(tieText(dpsRow('B', [mate('Arms', 'warrior')]))).toBe('≈ tie with Arms');
  });

  it('names every partner, and nothing on an untied row', () => {
    expect(
      tieText(
        dpsRow('B', [mate('Survival', 'hunter'), mate('Beast Mastery', 'hunter'), mate('Feral', 'druid')]),
      ),
    ).toBe('≈ tie with Survival, Beast Mastery and Feral');
    expect(tieText(dpsRow('B', []))).toBeNull();
    expect(tieTitle(dpsRow('B', []))).toBeUndefined();
  });

  it('reads correctly across a letter line: the partner carries its letter', () => {
    const survival = dpsRow('A', [mate('Beast Mastery', 'hunter', 'B')]);
    expect(tieText(survival)).toBe('≈ tie with Beast Mastery (B)');
    expect(tieTitle(survival)).toBe(
      'Within 1% of Beast Mastery Hunter: a tie; the letters differ by less than the sim can tell',
    );
    expect(tieTitle(dpsRow('B', [mate('Arms', 'warrior')]))).toBe('Within 1% of Arms Warrior: a tie');
  });

  it('names a healer partner in full, with the healer reason', () => {
    const row = { role: 'healer' as const, tier: null, tiedWith: [mate('Holy', 'priest', null)] };
    expect(tieText(row)).toBe('≈ tie with Holy Priest');
    expect(tieTitle(row)).toContain('The sim cannot order this healer and Holy Priest');
  });
});

describe('ranks', () => {
  it('writes a tied healer’s rank with an equals sign, an untied one plainly, and an unranked one not at all', () => {
    expect(tiersCopy.rankLabel({ rank: 2, role: 'healer', tie: 'above', ranked: true })).toBe('=2');
    expect(tiersCopy.rankLabel({ rank: 4, role: 'healer', tie: null, ranked: true })).toBe('4');
    expect(tiersCopy.rankLabel({ rank: 3, role: 'dps', tie: 'below', ranked: true })).toBe('3');
    expect(tiersCopy.rankLabel({ rank: 1, role: 'healer', tie: null, ranked: false })).toBe('');
  });

  it('shows a gap figure on a healer tied for the top or unranked, and Top only on a ranked leader', () => {
    expect(gapText({ rank: 1, role: 'healer', gapPercent: 0.3, ranked: true })).toBe('−0.3%');
    expect(gapText({ rank: 1, role: 'healer', gapPercent: 0, ranked: true })).toBe(tiersCopy.topGap);
    expect(gapText({ rank: 1, role: 'healer', gapPercent: 0, ranked: false })).toBe('0.0%');
  });
});

describe('column heads (owner MUST 2)', () => {
  it('keeps the tank heads short enough for one line', () => {
    expect(tiersCopy.heads.tank.slice(3, 7)).toEqual([
      'Survival vs the best',
      'Taken / s',
      'Health',
      'Threat / s',
    ]);
    for (const role of ['dps', 'tank', 'healer'] as const) {
      for (const head of tiersCopy.heads[role].slice(4)) expect(head.length).toBeLessThanOrEqual(10);
    }
  });
});

describe('gapText', () => {
  it('reads Top or Least on the leader, a signed gap below', () => {
    expect(gapText({ rank: 1, role: 'dps', gapPercent: 0, ranked: true })).toBe(tiersCopy.topGap);
    expect(gapText({ rank: 1, role: 'tank', gapPercent: 0, ranked: true })).toBe(tiersCopy.leastGap);
    expect(gapText({ rank: 2, role: 'dps', gapPercent: 4.84, ranked: true })).toBe('−4.8%');
    expect(gapText({ rank: 2, role: 'tank', gapPercent: 15.23, ranked: true })).toBe('+15.2%');
  });
});

describe('formatMetric', () => {
  it('writes tank figures whole and damage and healing rates to one decimal', () => {
    expect(formatMetric('tank', 322.14)).toBe('322');
    expect(formatMetric('tank', 35711.4)).toBe('35,711');
    expect(formatMetric('dps', 921.44)).toBe('921.4');
    expect(formatMetric('healer', 1204)).toBe('1,204.0');
  });
});

describe('tier bands', () => {
  it('words each band from its bounds', () => {
    expect(tiersCopy.tierBand({ from: 0, to: 5 })).toBe('Within 5% of the top');
    expect(tiersCopy.tierBand({ from: 5, to: 10 })).toBe('5 to 10% behind');
    expect(tiersCopy.tierBand({ from: 30, to: null })).toBe('30% or more behind');
  });

  it('counts specs in the singular and the plural', () => {
    expect(tiersCopy.tierCount(1)).toBe('1 spec');
    expect(tiersCopy.tierCount(4)).toBe('4 specs');
  });
});

describe('tier page head', () => {
  it('gives each role its own title and description', () => {
    const roles = ['dps', 'tank', 'healer'] as const;
    expect(new Set(roles.map(tiersCopy.pageTitle)).size).toBe(3);
    expect(new Set(roles.map(tiersCopy.metaDescription)).size).toBe(3);
    expect(tiersCopy.pageTitle('tank')).toBe('Tank tier list');
  });

  it('keeps every description inside a search snippet', () => {
    for (const role of ['dps', 'tank', 'healer'] as const) {
      expect(tiersCopy.metaDescription(role).length).toBeLessThanOrEqual(160);
    }
  });
});
