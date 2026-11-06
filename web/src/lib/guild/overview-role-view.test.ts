import { describe, expect, it } from 'vitest';
import type { GuildRaidCounts } from './api';
import type { GuildChecklist, GuildOfficerActions } from './overview-api';
import {
  actionRows,
  attendanceGridRows,
  checklistRows,
  compTallyRows,
  compTargetLine,
  coverGridRows,
  nextRaidFact,
  raidTimeLine,
  raidWeekday,
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

  it('reads Onyxia attunement as the quest chain, flask and food by name, and the fire total', () => {
    const rows = checklistRows(
      checklist([
        {
          key: 'attunement',
          state: 'fail',
          attunement_kind: 'quest',
          key_item: { id: 16309, name: 'Drakefire Amulet', quality: 3 },
        },
        {
          key: 'flask_food',
          state: 'fail',
          supplies: [{ id: 13510, name: 'Flask of the Titans', kind: 'flask', count: 1 }],
          missing: ['food'],
        },
        { key: 'fire_resistance', state: 'ok', resistance: 85 },
      ]),
    );
    expect(rows).toEqual([
      expect.objectContaining({
        tone: 'miss',
        title: "Attuned to Onyxia's Lair",
        sub: "Finish the attunement chain for Onyxia's Lair and keep the Drakefire Amulet in your bags.",
      }),
      expect.objectContaining({
        tone: 'warn',
        title: 'No buff food in your bags',
        sub: 'You have Flask of the Titans.',
      }),
      expect.objectContaining({ tone: 'ok', title: 'Fire resistance 85' }),
    ]);
  });

  it('asks the member to share gear the checklist cannot see, so the verdict never reads ready for it', () => {
    const rows = checklistRows(
      checklist([{ key: 'gear_gap', state: 'not_tracked', reason: 'gear_not_shared' }]),
    );
    expect(rows).toEqual([
      expect.objectContaining({
        tone: 'warn',
        title: 'Share your gear',
        fix: { label: 'How', href: '/setup' },
      }),
    ]);
    expect(verdictLine('almost', rows, 'Thursday')).toBe('Almost there: one thing before Thursday.');
  });

  it('asks the member to share quests and bags the export cannot see, each linking the setup', () => {
    const rows = checklistRows(
      checklist([
        { key: 'attunement', state: 'not_tracked', reason: 'quests_not_in_export', attunement_kind: 'quest' },
        { key: 'flask_food', state: 'not_tracked', reason: 'bags_not_shared' },
      ]),
    );
    expect(rows).toEqual([
      expect.objectContaining({
        key: 'attunement',
        tone: 'warn',
        title: 'Share your quests',
        fix: { label: 'How', href: '/setup' },
      }),
      expect.objectContaining({
        key: 'flask_food',
        tone: 'warn',
        title: 'Share your bags',
        fix: { label: 'How', href: '/setup' },
      }),
    ]);
  });

  it('says a done attunement keeps the amulet on you, and names what the bags carry', () => {
    const rows = checklistRows(
      checklist([
        {
          key: 'attunement',
          state: 'ok',
          attunement_kind: 'quest',
          key_item: { id: 16309, name: 'Drakefire Amulet', quality: 3 },
        },
        {
          key: 'flask_food',
          state: 'ok',
          supplies: [
            { id: 13510, name: 'Flask of the Titans', kind: 'flask', count: 1 },
            { id: 13931, name: 'Nightfin Soup', kind: 'food', count: 5 },
          ],
        },
      ]),
    );
    expect(rows.map((row) => row.sub)).toEqual([
      'Attunement chain done. Keep the Drakefire Amulet on you.',
      'You carry Flask of the Titans and Nightfin Soup.',
    ]);
  });

  it('never words a row with an em dash or a floor', () => {
    const rows = checklistRows(
      checklist([
        {
          key: 'attunement',
          state: 'fail',
          attunement_kind: 'quest',
          key_item: { id: 16309, name: 'Drakefire Amulet', quality: 3 },
        },
        { key: 'flask_food', state: 'fail', missing: ['flask_or_elixir', 'food'] },
        { key: 'fire_resistance', state: 'ok', resistance: 0 },
        { key: 'attunement', state: 'not_tracked', reason: 'quests_not_in_export' },
        { key: 'flask_food', state: 'not_tracked', reason: 'bags_not_shared' },
      ]),
    );
    for (const row of rows) expect(`${row.title} ${row.sub}`).not.toMatch(/\u2014|floor/i);
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

  it('counts the pending applications after approvals, and none when an older API sends none', () => {
    const rows = actionRows(
      {
        ...base,
        approvals: { count: 1, names: ['Nightlatch'] },
        applications: { count: 2, names: ['Ashenveil', 'Quillbright'] },
      },
      null,
    );
    expect(rows.map((row) => row.key)).toEqual(['approvals', 'applications']);
    expect(rows[1]).toEqual({ key: 'applications', title: '2 applications', sub: 'Ashenveil, Quillbright' });
    expect(actionRows({ ...base, applications: { count: 1, names: ['Ashenveil'] } }, null)[0].title).toBe(
      '1 application',
    );
    expect(actionRows({ ...base, applications: { count: 0, names: [] } }, null)).toEqual([]);
  });

  it("counts the loot council's open drops of its newest night, and none once all are awarded", () => {
    const drops = { tracked: true, count: 3, raid_id: 3, label: "Onyxia's Lair, Tue 8 Dec", awarded: 2 };
    expect(actionRows({ ...base, drops_to_award: drops }, null)).toEqual([
      { key: 'drops', title: '3 drops to award', sub: "Onyxia's Lair, Tue 8 Dec" },
    ]);
    expect(actionRows({ ...base, drops_to_award: { ...drops, count: 1 } }, null)[0].title).toBe(
      '1 drop to award',
    );
    expect(actionRows({ ...base, drops_to_award: { ...drops, count: 0 } }, null)).toEqual([]);
    expect(actionRows({ ...base, drops_to_award: null }, null)).toEqual([]);
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

describe('raid time and cover rows', () => {
  it('writes the guild time, and the viewer’s when it differs', () => {
    const at = '2026-12-11T02:00:00Z';
    expect(raidTimeLine(at, 'America/Chicago', 'America/Chicago', 'en-US')).toBe('Thu 10 Dec · 8 pm CST');
    expect(raidTimeLine(at, 'America/Chicago', 'America/New_York', 'en-US')).toBe(
      'Thu 10 Dec · 8 pm CST · 9 pm EST your time',
    );
    expect(raidWeekday(at, 'America/Chicago', 'en-US')).toBe('Thursday');
  });

  it('shows the need beside the signed providers, toned, with the bench and talents named', () => {
    const [row] = coverGridRows([
      {
        key: 'trueshot',
        name: 'Trueshot Aura',
        kind: 'buff',
        spell_id: 20906,
        icon: 'ability_trueshot',
        need: 1,
        per: 'raid',
        providers_in: 1,
        providers_late: 0,
        providers_bench: 1,
        likely: 1,
        state: 'ok',
      },
    ]);
    expect(row?.sub).toBe('1 on a talent we cannot see · 1 on the bench');
    expect(row?.cells.map((c) => [c.tone, c.text])).toEqual([
      ['none', '1'],
      ['warn', '1'],
    ]);
  });
});

describe('nextRaidFact', () => {
  it('names the zone and the day on the guild’s calendar', () => {
    expect(
      nextRaidFact({
        zone_name: "Onyxia's Lair",
        starts_at: '2026-12-11T02:00:00Z',
        guild_timezone: 'America/Chicago',
      }),
    ).toBe("Next raid: Onyxia's Lair, Thu 10 Dec");
  });
});
