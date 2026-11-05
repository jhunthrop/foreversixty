// web/src/lib/guild/raid-comp.test.ts
// The raid comp rules (design/specs/2026-11-05-guild-raid-comp.md): the class order, the spec and
// class line, the picker's pairs, the count clauses, the class row and the copied list.
import { describe, expect, it } from 'vitest';
import { specCell, specString } from '../../test-support/spec-table';
import type { GuildRaidSignup, GuildRaidView } from './api';
import { compCopy } from './raid-copy';
import {
  CLASS_ORDER,
  classRow,
  compBreakdownParts,
  pickFor,
  pickOf,
  raidBreakdownLine,
  raidBreakdownParts,
  rowGroup,
  signupListText,
  specClassLine,
  specPicks,
  whoLine,
} from './raid-comp';

const SPEC = new URL('../../../../design/specs/2026-11-05-guild-raid-comp.md', import.meta.url);
const HEADING = '## 7. Copy, verbatim';
const str = (label: string, index = 0, fill: Record<string, string> = {}): string =>
  specString(SPEC, HEADING, label, index, fill);
const CHICAGO = 'America/Chicago';

function signup(
  name: string,
  cls: string,
  specKey: string | null,
  role: GuildRaidSignup['role'],
  state: GuildRaidSignup['state'],
  group: GuildRaidSignup['group'],
): GuildRaidSignup {
  return {
    character_key: `us/pvp/${name.toLowerCase()}`,
    name,
    class: cls,
    spec: specKey ?? '',
    spec_key: specKey,
    role,
    state,
    note: '',
    own: false,
    group,
  };
}

describe('the class order', () => {
  it('is the game’s, Warrior to Druid (spec 3.1)', () => {
    expect(CLASS_ORDER).toEqual([
      'warrior',
      'paladin',
      'hunter',
      'rogue',
      'priest',
      'shaman',
      'mage',
      'warlock',
      'druid',
    ]);
  });
});

describe('specClassLine and whoLine', () => {
  it('reads spec then class, the catalogue name, the class alone when unresolved', () => {
    expect(specClassLine('paladin', 'paladin-holy')).toBe(
      str('Who line', 0, { Spec: 'Holy', Class: 'Paladin' }),
    );
    expect(specClassLine('hunter', 'hunter-beast-mastery')).toBe('Beast Mastery Hunter');
    expect(specClassLine('paladin', null)).toBe('Paladin');
    expect(specClassLine('paladin', 'priest-holy')).toBe('Paladin');
  });

  it('marks an off-tank pick with its role', () => {
    expect(whoLine('warrior', 'warrior-fury', 'tank')).toBe(str('Who line', 3));
    expect(whoLine('warrior', 'warrior-fury', 'dps')).toBe('Fury Warrior');
    expect(whoLine('warrior', 'warrior-protection', 'tank')).toBe('Protection Warrior');
  });
});

describe('specPicks', () => {
  it('lists one pair per spec, Arms and Fury adding a tank pair (spec 3.2.1)', () => {
    expect(specPicks('warrior').map((pick) => pick.label)).toEqual([
      'Arms · DPS',
      'Arms · Tank',
      'Fury · DPS',
      'Fury · Tank',
      'Protection · Tank',
    ]);
    expect(specPicks('druid').map((pick) => pick.label)).toEqual([
      'Balance · DPS',
      'Feral · DPS',
      'Feral Bear · Tank',
      'Restoration · Healer',
    ]);
    expect(specPicks('paladin').map((pick) => pick.label)).toEqual([
      'Holy · Healer',
      'Protection · Tank',
      'Retribution · DPS',
    ]);
  });

  it('keeps the spec key on the off-tank pair, only the role differs', () => {
    const offTank = pickOf('warrior', 'warrior-fury:tank');
    expect(offTank?.spec.spec).toBe('warrior-fury');
    expect(offTank?.role).toBe('tank');
    expect(pickOf('warrior', 'warrior-protection:dps')).toBeNull();
    expect(pickOf('mage', 'warrior-fury')).toBeNull();
  });

  it('starts a returning off-tank on Fury · Tank, and any other row on the spec’s own pair', () => {
    expect(pickFor('warrior', 'warrior-fury', 'tank')?.label).toBe('Fury · Tank');
    expect(pickFor('warrior', 'warrior-fury', 'dps')?.label).toBe('Fury · DPS');
    expect(pickFor('priest', 'priest-holy', 'dps')?.label).toBe('Holy · Healer');
    expect(pickFor('priest', null, 'healer')).toBeNull();
  });
});

describe('the count clauses', () => {
  it('reads melee and ranged in place of DPS, a zero clause kept, the late roles named', () => {
    const counts = { tanks: 3, healers: 7, dps: 12, melee: 7, ranged: 5, late: 1 };
    expect(compBreakdownParts(counts, { tanks: 0, healers: 1 }).join(' · ')).toBe(
      '3 tanks · 7 healers · 7 melee · 5 ranged · 1 late (healer)',
    );
    expect(compBreakdownParts({ ...counts, tanks: 0, late: 0 }, { tanks: 0, healers: 0 })).toEqual([
      '0 tanks',
      '7 healers',
      '7 melee',
      '5 ranged',
    ]);
  });

  it('takes the late roles from the rows’ groups', () => {
    const raid = {
      counts: { in: 1, late: 2, bench: 0, out: 0, tanks: 1, healers: 0, dps: 0, melee: 0, ranged: 0 },
      signups: [
        signup('Kraggor', 'warrior', 'warrior-protection', 'tank', 'in', 'tank'),
        signup('Brokk', 'warrior', 'warrior-fury', 'tank', 'late', 'tank'),
        signup('Mendara', 'priest', 'priest-holy', 'healer', 'late', 'healer'),
      ],
    };
    expect(raidBreakdownParts(raid).at(-1)).toBe('2 late (1 tank, 1 healer)');
  });

  it('glues each shown clause with no-break spaces, so it wraps only between clauses', () => {
    const raid = {
      counts: { in: 1, late: 1, bench: 0, out: 0, tanks: 0, healers: 1, dps: 0, melee: 0, ranged: 0 },
      signups: [signup('Mendara', 'priest', 'priest-holy', 'healer', 'late', 'healer')],
    };
    const line = raidBreakdownLine(raid);
    expect(line.split(' · ')).toEqual(['0 tanks', '1 healer', '0 melee', '0 ranged', '1 late (healer)']);
  });

  it('reads a row with no group by its role, DPS as ranged', () => {
    expect(rowGroup({ role: 'healer' })).toBe('healer');
    expect(rowGroup({ role: 'dps' })).toBe('ranged');
    expect(rowGroup({ role: 'dps', group: 'melee' })).toBe('melee');
  });
});

const ONYXIA = {
  id: 31,
  title: "Onyxia's Lair",
  starts_at: '2026-12-11T02:00:00Z',
  size: 40,
  counts: {
    in: 5,
    late: 1,
    bench: 1,
    out: 1,
    tanks: 2,
    healers: 1,
    dps: 2,
    melee: 1,
    ranged: 1,
    classes: { warrior: 2, priest: 2, hunter: 2, mage: 0 },
  },
  signups: [
    signup('Kraggor', 'warrior', 'warrior-protection', 'tank', 'in', 'tank'),
    signup('Brokk', 'warrior', 'warrior-fury', 'tank', 'in', 'tank'),
    signup('Mendara', 'priest', 'priest-holy', 'healer', 'in', 'healer'),
    signup('Kelthar', 'hunter', 'hunter-survival', 'dps', 'in', 'melee'),
    signup('Zulmara', 'hunter', 'hunter-marksmanship', 'dps', 'in', 'ranged'),
    signup('Ysolde', 'priest', 'priest-holy', 'healer', 'late', 'healer'),
    signup('Pyrewisp', 'mage', 'mage-frost', 'dps', 'bench', 'ranged'),
    signup('Corvane', 'rogue', null, 'dps', 'out', 'melee'),
  ],
  not_signed: [],
} as unknown as GuildRaidView;

describe('classRow', () => {
  it('lists the API’s classes in the game’s order, zeros kept, with each class’s specs', () => {
    const row = classRow(ONYXIA);
    expect(row.map((item) => `${item.name} ${item.count}`)).toEqual([
      'Warrior 2',
      'Hunter 2',
      'Priest 2',
      'Mage 0',
    ]);
    expect(row[0]?.specs).toBe('Fury (tank) 1, Protection 1');
    expect(row[2]?.accessibleName).toBe('Priest 2: Holy 2');
    expect(row[3]?.accessibleName).toBe('Mage 0');
  });

  it('puts the most played spec first, then catalogue order, an unresolved spec as Unknown', () => {
    const raid = {
      ...ONYXIA,
      counts: { ...ONYXIA.counts, classes: { priest: 5 } },
      signups: [
        signup('A', 'priest', 'priest-shadow', 'dps', 'in', 'ranged'),
        signup('B', 'priest', 'priest-holy', 'healer', 'in', 'healer'),
        signup('C', 'priest', 'priest-holy', 'healer', 'late', 'healer'),
        signup('D', 'priest', 'priest-discipline', 'healer', 'in', 'healer'),
        signup('E', 'priest', null, 'healer', 'in', 'healer'),
      ],
    } as GuildRaidView;
    expect(classRow(raid)[0]?.specs).toBe('Holy 2, Discipline 1, Shadow 1, Unknown 1');
  });

  it('tallies from the rows for an API without the class counts, and is empty with no roster', () => {
    const { classes: _dropped, ...counts } = ONYXIA.counts;
    const items = classRow({ ...ONYXIA, counts });
    expect(items.map((item) => `${item.classSlug} ${item.count}`)).toEqual([
      'warrior 2',
      'hunter 2',
      'rogue 0',
      'priest 2',
      'mage 0',
    ]);
    expect(classRow({ ...ONYXIA, counts, signups: [], not_signed: [] })).toEqual([]);
  });

  it('names the list and an item in the spec’s words', () => {
    expect(specCell(SPEC, HEADING, 'Class row')[1]).toBe(compCopy.classRowLabel);
  });
});

describe('signupListText (spec 4.2)', () => {
  it('writes the count line, the class line and one line per group, names with spec and class', () => {
    expect(signupListText(ONYXIA, CHICAGO).split('\n')).toEqual([
      "Onyxia's Lair · Thu 10 Dec, 8:00 pm CST · <t:1796954400:F>",
      '6 of 40: 2 tanks, 1 healer, 1 melee, 1 ranged, 1 late (healer)',
      'Classes: Warrior 2, Hunter 2, Priest 2',
      'Tanks: Brokk (Fury Warrior), Kraggor (Protection Warrior)',
      'Healers: Mendara (Holy Priest)',
      'Melee: Kelthar (Survival Hunter)',
      'Ranged: Zulmara (Marksmanship Hunter)',
      'Late: Ysolde (Holy Priest)',
      'Bench: Pyrewisp (Frost Mage)',
      'Out: Corvane (Rogue)',
    ]);
  });

  it('leaves the class line out with nobody in, and a null size reads signed up', () => {
    const none = {
      ...ONYXIA,
      size: null,
      counts: { ...ONYXIA.counts, in: 0, late: 0, classes: { mage: 0 } },
      signups: [],
    } as unknown as GuildRaidView;
    expect(signupListText(none, CHICAGO).split('\n').slice(1)).toEqual([
      '0 signed up: 2 tanks, 1 healer, 1 melee, 1 ranged',
    ]);
  });
});
