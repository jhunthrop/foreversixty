// web/src/lib/tiers/tier-copy.test.ts
import { describe, expect, it } from 'vitest';
import {
  ROLE_TABS,
  formatMetric,
  gapText,
  raceBasis,
  readColumns,
  tiersCopy,
  type ReadColumn,
} from './tier-copy';
import { ordinal } from '../format/number';

const flat = (column: ReadColumn): string => column.body.map((segment) => segment.text).join('');
const CTX = { faction: 'horde', presetLabel: 'Raid-ready, Phase 1' } as const;
const HEALER_LEAD =
  'HPS is effective healing, overheal left out. A shield counts for what it soaks, not for the death it stopped by being up before the big hit, and buffs on others, like Power Infusion, are not counted.';
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

describe('the title band', () => {
  it('names the gear, then the fight: the target level and its length', () => {
    expect(tiersCopy.title).toBe('Tier list at 60');
    expect(tiersCopy.metaLine({ role: 'dps' })).toBe('Best in slot, raid buffs · vs a level 63 boss · 3 min');
    expect(tiersCopy.metaLine({ role: 'tank' })).toBe('Best in slot, raid buffs · vs a level 63 boss');
    expect(tiersCopy.metaLine({ role: 'healer', healSeconds: 300 })).toBe('Best in slot, raid buffs · 5 min');
    expect(tiersCopy.metaLine({ role: 'healer', healSeconds: 90 })).toBe('Best in slot, raid buffs · 90 s');
  });

  it('says under the healer title what HPS does not count, and says it on no other list', () => {
    expect(tiersCopy.lead('healer')).toBe(HEALER_LEAD);
    expect(tiersCopy.lead('dps')).toBeUndefined();
    expect(tiersCopy.lead('tank')).toBeUndefined();
    expect(HEALER_LEAD).not.toMatch(/[—–]/);
  });

  it('the healer notes say the same thing the healer title does', () => {
    const leavesOut = readColumns({ ...CTX, role: 'healer', healSeconds: 300 })[1]!;
    expect(flat(leavesOut)).toContain(HEALER_LEAD);
  });

  it('fails fast when the healer fight length is missing', () => {
    expect(() => tiersCopy.metaLine({ role: 'healer' })).toThrow(/heal profile duration/);
  });

  it('titles each letter with its band, against #1', () => {
    expect(tiersCopy.tierBand('S')).toBe('Within 5% of #1');
    expect(tiersCopy.tierBand('B')).toBe('10 to 20% behind #1');
    expect(tiersCopy.tierBand('D')).toBe('30% or more behind #1');
  });

  it("names a tied healer's or tank's partners in full, and nothing on a DPS row", () => {
    const mate = { classSlug: 'shaman', specSlug: 'restoration', name: 'Restoration', tier: null };
    const bear = { classSlug: 'druid', specSlug: 'feral-bear', name: 'Feral (Bear)', tier: null };
    expect(tiersCopy.rowTie({ role: 'tank', tiedWith: [bear] })?.title).toMatch(/within 1% of each other/);
    expect(tiersCopy.rowTie({ role: 'healer', tiedWith: [mate] })).toMatchObject({
      text: '≈ tie with Restoration Shaman',
    });
    expect(tiersCopy.rowTie({ role: 'healer', tiedWith: [] })).toBeNull();
    expect(tiersCopy.rowTie({ role: 'dps', tiedWith: [mate] })).toBeNull();
  });

  it('names each role tab for the list it opens', () => {
    expect(ROLE_TABS).toEqual({ dps: 'Damage', tank: 'Tanks', healer: 'Healers' });
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

describe('column heads', () => {
  it("names each figure's unit in its head, short enough for one line, and the gap against #1", () => {
    expect(tiersCopy.heads.dps).toEqual(['#', '', 'Spec', '', 'DPS', 'vs #1', '']);
    expect(tiersCopy.heads.tank.slice(4, 8)).toEqual(['Taken / s', 'Health', 'Threat / s', 'vs #1']);
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
    expect(gapText({ rank: 2, role: 'tank', gapPercent: 15.23, ranked: true })).toBe('+15% taken');
    expect(gapText({ rank: 2, role: 'tank', gapPercent: 52.9, ranked: true })).toBe('+53% taken');
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

describe('the spec panel', () => {
  it("points a spec on another role to that role's list, in the tab's words", () => {
    expect(tiersCopy.card.otherRole('dps')).toBe('On the Damage list');
    expect(tiersCopy.card.signInHref('/tiers/tank')).toBe('/login?next=%2Ftiers%2Ftank');
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
