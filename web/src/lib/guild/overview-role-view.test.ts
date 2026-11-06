import { describe, expect, it } from 'vitest';
import type { GuildRaidCounts } from './api';
import type { GuildChecklist, GuildOfficerActions } from './overview-api';
import {
  actionRows,
  attendanceGridRows,
  checklistRows,
  compTallyRows,
  compTargetLine,
  verdictLine,
} from './overview-role-view';
import { tallyTone } from '../ui/cells';

const checklist = (
  items: GuildChecklist['items'],
  verdict: GuildChecklist['verdict'] = 'almost',
): GuildChecklist => ({
  character: {
    character_key: 'us/pvp/hollowmend',
    name: 'Hollowmend',
    class: 'Priest',
    spec: 'Holy',
    level: 60,
  },
  raid: { id: 7, zone_key: 'raid:onyxias-lair', zone_name: "Onyxia's Lair" },
  verdict,
  items,
});

describe('checklistRows', () => {
  it('drops the rows nothing tracks yet, never drawing a stand-in', () => {
    const rows = checklistRows(
      checklist([
        { key: 'level', state: 'ok', level: 60, required_level: 60 },
        { key: 'flask_food', state: 'not_tracked', reason: 'consumables_not_named' },
        { key: 'fire_resistance', state: 'not_tracked', reason: 'resistance_not_in_export' },
      ]),
    );
    expect(rows.map((row) => row.key)).toEqual(['level']);
  });

  it('names the gear gap in the client’s slot words and links the BiS list', () => {
    const [row] = checklistRows(
      checklist([{ key: 'gear_gap', state: 'fail', upgrades: 2, slots: ['hands', 'finger_1'], unit: 'dps' }]),
    );
    expect(row).toMatchObject({
      tone: 'warn',
      title: '2 slots are not best in slot',
      sub: 'Hands and Finger 1.',
    });
    expect(row?.fix?.href).toMatch(/^\/bis\/priest\/holy/);
  });

  it('marks a level below the raid’s as missing, with the zone named', () => {
    const [row] = checklistRows(checklist([{ key: 'level', state: 'fail', level: 52, required_level: 60 }]));
    expect(row).toMatchObject({
      tone: 'miss',
      title: "Level 60 for Onyxia's Lair",
      sub: 'You are level 52.',
    });
  });

  it('lists the slots without an enchant', () => {
    const [row] = checklistRows(
      checklist([{ key: 'enchants', state: 'fail', missing_slots: ['main_hand', 'feet'] }]),
    );
    expect(row).toMatchObject({ title: '2 slots without an enchant', sub: 'Main Hand and Feet.' });
  });
});

describe('verdictLine (D12)', () => {
  const open = checklistRows(
    checklist([
      { key: 'gear_gap', state: 'fail', upgrades: 1, slots: ['hands'] },
      { key: 'enchants', state: 'fail', missing_slots: ['chest'] },
    ]),
  );

  it('counts the open checks in words', () => {
    expect(verdictLine('almost', open, 'Thursday')).toBe('Almost there: two things before Thursday.');
  });

  it('never says floor, and reads ready when nothing is open', () => {
    expect(verdictLine('ready', [], 'Thursday')).toBe('Ready to raid on Thursday.');
    expect(verdictLine('getting_started', open, null)).toBe('Getting started: two things before the raid.');
    expect(verdictLine('almost', open, null)).not.toMatch(/floor/i);
  });
});

describe('actionRows', () => {
  const base: GuildOfficerActions = {
    approvals: { count: 0, names: [] },
    not_signed: { raid_id: 7, count: 0, names: [] },
    drops_to_award: {
      tracked: false,
      count: null,
      raid_id: 3,
      label: "Onyxia's Lair, Tue 8 Dec",
      awarded: 0,
    },
  };

  it('shows only what is waiting, and hides untracked drops', () => {
    expect(actionRows(base, 'Thursday')).toEqual([]);
  });

  it('words approvals and non-signers, naming three then a count', () => {
    const rows = actionRows(
      {
        ...base,
        approvals: { count: 1, names: ['Nightlatch'] },
        not_signed: { raid_id: 7, count: 4, names: ['Voidcall', 'Skyarcana', 'Starnettle', 'Ashveil'] },
      },
      'Thursday',
    );
    expect(rows).toEqual([
      { key: 'approvals', title: '1 raider is waiting for approval', sub: 'Nightlatch' },
      {
        key: 'not_signed',
        title: '4 have not signed up for Thursday',
        sub: 'Voidcall, Skyarcana, Starnettle and 1 more',
      },
    ]);
  });
});

describe('comp against the target', () => {
  const counts: GuildRaidCounts = {
    in: 23,
    late: 2,
    bench: 2,
    out: 0,
    tanks: 1,
    healers: 6,
    dps: 16,
    melee: 5,
    ranged: 11,
    late_groups: { tanks: 0, healers: 1, melee: 1, ranged: 0 },
  };
  const target = { tanks: 2, healers: 8, melee: 10, ranged: 20, source: 'default' as const };

  it('reads In per role against the seats, Late beside it', () => {
    expect(compTallyRows(counts, target).map((r) => [r.label, r.have, r.late, r.want, r.tone])).toEqual([
      ['Tanks', 1, 0, 2, 'miss'],
      ['Healers', 6, 1, 8, 'warn'],
      ['Melee', 5, 1, 10, 'miss'],
      ['Ranged', 11, 0, 20, 'miss'],
    ]);
  });

  it('tones at, near and short of the target', () => {
    expect([tallyTone(8, 8), tallyTone(6, 8), tallyTone(5, 8)]).toEqual(['ok', 'warn', 'miss']);
  });

  it('says where the target came from', () => {
    expect(compTargetLine("Onyxia's Lair", target)).toBe("Onyxia's Lair default: 40 players");
    expect(compTargetLine("Onyxia's Lair", { ...target, source: 'raid', ranged: 18 })).toBe(
      'Set for this raid: 38 players',
    );
  });
});

describe('attendanceGridRows', () => {
  it('puts slipping raiders first and maps each night to a mark', () => {
    const rows = attendanceGridRows({
      raids: [],
      rows: [
        {
          name: 'Cinderglass',
          class: 'mage',
          account_key: 'a',
          cells: ['bench', 'present'],
          slipping: false,
        },
        {
          name: 'Pinefeather',
          class: 'druid',
          account_key: 'b',
          cells: ['present', 'absent'],
          slipping: true,
        },
      ],
    });
    expect(rows.map((r) => [r.label, r.sub, r.cells.map((c) => c.text).join('')])).toEqual([
      ['Pinefeather', 'slipping', '✓×'],
      ['Cinderglass', 'bench', 'B✓'],
    ]);
  });
});
