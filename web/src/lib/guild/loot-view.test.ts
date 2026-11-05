// web/src/lib/guild/loot-view.test.ts
import { describe, expect, it } from 'vitest';
import type { GuildLootCandidate, GuildLootEncounter, GuildLootItem, GuildLootPage } from './api';
import {
  ALREADY_AWARDED,
  LOOT_FOOTNOTE,
  attendanceLabel,
  awardQuestion,
  awardedLine,
  bossPickerLabel,
  candidateAction,
  candidateGainLabel,
  candidateSubline,
  itemsWonLabel,
  nextUnkilledEncounter,
  overviewLootLine,
  rankedDrops,
  showsAttendance,
  topPick,
  unrankedDrops,
  unrankedLabel,
  unrankedLine,
} from './loot-view';

function encounter(overrides: Partial<GuildLootEncounter> = {}): GuildLootEncounter {
  return { encounter_id: 1084, name: 'Onyxia', zone: "Onyxia's Lair", killed: true, ...overrides };
}

function candidateRow(overrides: Partial<GuildLootCandidate> = {}): GuildLootCandidate {
  return {
    character_key: 'a',
    name: 'Kraggor',
    class: 'warrior',
    spec: 'Protection',
    gain_dps: null,
    gain_unit: null,
    tier: 1,
    not_sim_checked: true,
    ilvl_delta: null,
    attendance: { present: 8, nights: 8 },
    already_equivalent: false,
    replaces: '',
    replaces_item: null,
    items_won: 0,
    ...overrides,
  };
}

function item(overrides: Partial<GuildLootItem> = {}): GuildLootItem {
  return {
    item_id: 16955,
    name: 'Helm of Wrath',
    icon: 'inv_helmet_71',
    quality: 4,
    slot: 'head',
    inventory_type: 1,
    item_class: 4,
    item_subclass: 4,
    kind: 'gear',
    quest_label: '',
    eligible: [],
    copies: 1,
    awarded_to: null,
    candidates: [
      candidateRow({ gain_dps: 40, gain_unit: 'dps', tier: 0, not_sim_checked: false }),
      candidateRow({
        character_key: 'b',
        name: 'Grimtotem',
        spec: 'Fury',
        gain_dps: 12,
        gain_unit: 'dps',
        tier: 0,
        not_sim_checked: false,
        attendance: { present: 6, nights: 8 },
      }),
    ],
    ...overrides,
  };
}

describe('bossPickerLabel', () => {
  it('reads "on farm" for a boss that is down, never filter jargon', () => {
    expect(bossPickerLabel(encounter({ killed: true }))).toBe('Onyxia · on farm');
  });

  it('reads "not killed yet" for an unkilled boss', () => {
    expect(bossPickerLabel(encounter({ killed: false }))).toBe('Onyxia · not killed yet');
  });

  it('reads the honest empty line when nothing is selected', () => {
    expect(bossPickerLabel(undefined)).toBe('No bosses killed yet this tier');
  });
});

describe('nextUnkilledEncounter', () => {
  it('finds the first unkilled encounter', () => {
    const rows = [
      encounter({ encounter_id: 1, killed: true }),
      encounter({ encounter_id: 2, killed: false }),
    ];
    expect(nextUnkilledEncounter(rows)?.encounter_id).toBe(2);
  });

  it('is undefined when every encounter is killed', () => {
    expect(nextUnkilledEncounter([encounter({ killed: true })])).toBeUndefined();
  });
});

describe('candidateGainLabel', () => {
  it('reads Fills an empty slot when the candidate wears nothing there, never the raw item level', () => {
    expect(candidateGainLabel(candidateRow({ tier: 1, ilvl_delta: 76, fills_empty_slot: true }))).toBe(
      'Fills an empty slot',
    );
    expect(candidateGainLabel(candidateRow({ tier: 1, ilvl_delta: 14, fills_empty_slot: false }))).toBe(
      '+14 item levels',
    );
  });

  it("reads a tier-0 gain to one decimal, in the BiS file's unit", () => {
    expect(candidateGainLabel(candidateRow({ tier: 0, gain_dps: 40.44, gain_unit: 'dps' }))).toBe(
      '+40.4\u00a0DPS',
    );
    expect(candidateGainLabel(candidateRow({ tier: 0, gain_dps: 12.4, gain_unit: 'hps' }))).toBe(
      '+12.4\u00a0HPS',
    );
    expect(candidateGainLabel(candidateRow({ tier: 0, gain_dps: 9, gain_unit: 'tps' }))).toBe(
      '+9.0\u00a0TPS',
    );
  });

  it('prints no figure for a tier-0 row with no unit or no measurement', () => {
    expect(candidateGainLabel(candidateRow({ tier: 0, gain_dps: null, gain_unit: null }))).toBe(
      'On their BiS list',
    );
    expect(candidateGainLabel(candidateRow({ tier: 0, gain_dps: 4, gain_unit: null }))).toBe(
      'On their BiS list',
    );
  });

  // Live-fix round, defect 5: a tier-1 row has gain_dps null, and this must read the item level
  // figure, never throw calling a number method on it.
  it('reads the item-level gap for a tier-1 candidate', () => {
    expect(candidateGainLabel(candidateRow({ ilvl_delta: 14 }))).toBe('+14\u00a0item\u00a0levels');
    expect(candidateGainLabel(candidateRow({ ilvl_delta: 1 }))).toBe('+1\u00a0item\u00a0level');
  });
});

describe('which drops are drawn as cards', () => {
  const ranked = item();
  const nobody = item({ item_id: 1, name: 'Judgement Crown', candidates: [] });
  const quest = item({ item_id: 2, name: 'Head of Onyxia', kind: 'quest', candidates: [] });
  const bag = item({ item_id: 3, name: 'Onyxia Hide Backpack', kind: 'bag', candidates: [] });
  const scale = item({ item_id: 4, name: 'Scale of Onyxia', kind: 'skinned', candidates: [] });
  const all = [scale, ranked, nobody, quest, bag];

  it('keeps a slotted drop with a candidate, in the boss order', () => {
    expect(rankedDrops(all)).toEqual([ranked]);
  });

  it('collapses everything else', () => {
    expect(unrankedDrops(all).map((row) => row.item_id)).toEqual([4, 1, 2, 3]);
  });

  it('counts the collapsed line, singular at one', () => {
    expect(unrankedLine(6)).toBe('6 more drops nobody here ranks for');
    expect(unrankedLine(1)).toBe('1 more drop nobody here ranks for');
  });

  it('labels each kind of collapsed drop', () => {
    expect(unrankedLabel({ ...quest, quest_label: 'Quest: A Proper String (Hunter)' })).toBe(
      'Quest: A Proper String (Hunter)',
    );
    expect(unrankedLabel(bag)).toBe('Bag. Anyone can use it.');
    expect(unrankedLabel(scale)).toBe('Skinned from Onyxia. Not a drop.');
    expect(unrankedLabel(nobody)).toBe('Nobody here can use it yet.');
  });
});

describe('showsAttendance', () => {
  const zero = { attendance: { present: 0, nights: 0 } };

  it('hides the cell until a character has a counted raid night', () => {
    expect(showsAttendance([item({ candidates: [candidateRow(zero)] })])).toBe(false);
    expect(showsAttendance([item({ candidates: [] })])).toBe(false);
    expect(
      showsAttendance([
        item({ candidates: [candidateRow(zero)] }),
        item({ candidates: [candidateRow({ attendance: { present: 1, nights: 1 } })] }),
      ]),
    ).toBe(true);
  });
});

describe('the overview line', () => {
  function page(overrides: Partial<GuildLootPage> = {}): GuildLootPage {
    return {
      encounters: [encounter({ killed: false })],
      selected: 1084,
      items: [],
      at_level_cap: 0,
      ...overrides,
    };
  }

  it('says nobody is 60 before a kill with nobody at the cap', () => {
    expect(overviewLootLine(page())).toEqual({
      kind: 'text',
      text: "Onyxia's drops need level 60. Nobody here is 60 yet.",
    });
  });

  it('names the top pick before a kill', () => {
    const line = overviewLootLine(page({ at_level_cap: 3, items: [item()] }));
    expect(line).toMatchObject({ kind: 'pick', lead: 'Onyxia not killed yet. Top pick today: ' });
    expect(line.kind === 'pick' ? line.pick.candidate.name : '').toBe('Kraggor');
  });

  it('names the next open drop after a kill, skipping an awarded one', () => {
    const awarded = item({
      awarded_to: { award_id: 1, character_key: 'a', name: 'Kraggor', at: '', by_name: '' },
    });
    const open = item({ item_id: 2, name: 'Deathbringer' });
    const line = overviewLootLine(
      page({ encounters: [encounter({ killed: true })], at_level_cap: 3, items: [awarded, open] }),
    );
    expect(line).toMatchObject({ kind: 'pick', lead: 'Next open drop: ' });
    expect(line.kind === 'pick' ? line.pick.item.name : '').toBe('Deathbringer');
  });

  it('never offers the skinned scale, and reads the boss word when nothing is open', () => {
    const scale = item({ item_id: 15410, kind: 'skinned', candidates: [] });
    expect(overviewLootLine(page({ encounters: [encounter()], at_level_cap: 3, items: [scale] }))).toEqual({
      kind: 'text',
      text: 'Onyxia · on farm',
    });
  });
});

describe('topPick', () => {
  it('takes the highest tier-0 gain across every drop, over any tier-1 row', () => {
    const tierOne = item({
      item_id: 1,
      candidates: [candidateRow({ name: 'Delta', ilvl_delta: 60 })],
    });
    const small = item({
      item_id: 2,
      candidates: [candidateRow({ name: 'Small', tier: 0, gain_dps: 2, gain_unit: 'dps' })],
    });
    const big = item({
      item_id: 3,
      candidates: [candidateRow({ name: 'Big', tier: 0, gain_dps: 9, gain_unit: 'dps' })],
    });
    expect(topPick([tierOne, small, big])?.candidate.name).toBe('Big');
  });

  it('takes the highest tier-1 row when no tier-0 row exists', () => {
    const low = item({ item_id: 1, candidates: [candidateRow({ name: 'Low', ilvl_delta: 2 })] });
    const high = item({ item_id: 2, candidates: [candidateRow({ name: 'High', ilvl_delta: 20 })] });
    expect(topPick([low, high])?.candidate.name).toBe('High');
  });

  it('skips a candidate who already holds the item', () => {
    const held = item({
      candidates: [
        candidateRow({ name: 'Holds', tier: 0, already_equivalent: true }),
        candidateRow({ name: 'Wants', ilvl_delta: 3 }),
      ],
    });
    expect(topPick([held])?.candidate.name).toBe('Wants');
  });

  it('is null with nothing open', () => {
    expect(topPick([])).toBeNull();
  });
});

describe('candidateAction', () => {
  it('reads Award for an unawarded item with a real gain', () => {
    expect(candidateAction(item(), 'a')).toEqual({ kind: 'award' });
  });

  it('reads "equivalent" for a candidate who already holds the item', () => {
    const withEquivalent = item({
      candidates: [
        {
          character_key: 'c',
          name: 'Sunderfel',
          class: 'warrior',
          spec: 'Arms',
          gain_dps: 0,
          not_sim_checked: false,
          gain_unit: 'dps',
          tier: 0,
          ilvl_delta: null,
          attendance: { present: 8, nights: 8 },
          already_equivalent: true,
          replaces: '',
          replaces_item: null,
          items_won: 0,
        },
      ],
    });
    expect(candidateAction(withEquivalent, 'c')).toEqual({ kind: 'equivalent' });
  });

  it('marks exactly the awardee "awarded" and every other candidate "Already awarded"', () => {
    const awarded = item({
      awarded_to: {
        award_id: 7,
        character_key: 'a',
        name: 'Kraggor',
        at: '2026-12-15T00:00:00Z',
        by_name: 'Kraggor',
      },
    });
    expect(candidateAction(awarded, 'a')).toEqual({ kind: 'awarded' });
    expect(candidateAction(awarded, 'b')).toEqual({ kind: 'awarded-to-other', label: ALREADY_AWARDED });
  });
});

describe('the words around a candidate', () => {
  it('asks the award as a question naming the item and the candidate', () => {
    expect(awardQuestion('Stormrage Cover', 'Rootbound')).toBe('Award Stormrage Cover to Rootbound?');
  });

  it('counts raid nights as a player says it', () => {
    expect(attendanceLabel(1, 1)).toBe('1 of 1 raid night');
    expect(attendanceLabel(6, 8)).toBe('6 of 8 raid nights');
  });

  it('prints the award once, with the day and who gave it', () => {
    const now = new Date(2026, 9, 20, 21, 0);
    expect(
      awardedLine(
        { name: 'Rootbound', at: new Date(2026, 9, 20, 20, 30).toISOString(), by_name: 'Bow Jackzon' },
        now,
      ),
    ).toBe('Awarded to Rootbound · today · by\u00a0Bow Jackzon');
    expect(
      awardedLine({ name: 'Rootbound', at: new Date(2026, 9, 19, 20, 30).toISOString(), by_name: '' }, now),
    ).toBe('Awarded to Rootbound · yesterday');
  });

  it('names an officer when the account has no character in the guild', () => {
    const at = new Date(2026, 9, 20, 20, 30).toISOString();
    const now = new Date(2026, 9, 20, 21, 0);
    expect(awardedLine({ name: 'Rootbound', at, by_name: 'an officer' }, now)).toBe(
      'Awarded to Rootbound · today · by\u00a0an officer',
    );
    // A no-break space after "by", so the name never orphans from it.
    expect(awardedLine({ name: 'Rootbound', at, by_name: 'Bow Jackzon' }, now)).toContain(
      'by\u00a0Bow Jackzon',
    );
  });

  it('keeps the one closing line a player can use', () => {
    expect(LOOT_FOOTNOTE).toBe(
      'Candidates are characters who can wear the item at their level. Each is ranked by how much it adds over what they wear now. A drop stays open until an officer awards it.',
    );
  });
});

describe('candidateSubline', () => {
  const base = candidateRow({ name: 'Earthhoof', class: 'druid', spec: 'Restoration' });

  it('reads the spec with its class', () => {
    expect(candidateSubline(base)).toBe('Restoration Druid');
    expect(candidateSubline({ ...base, class: 'hunter', spec: 'BeastMastery' })).toBe('Beast Mastery Hunter');
  });

  it('adds the tier count only when the guild has awarded them something', () => {
    expect(candidateSubline({ ...base, items_won: 1 })).toBe('Restoration Druid · 1 item won this tier');
    expect(candidateSubline({ ...base, items_won: 0 })).not.toContain('won');
    expect(itemsWonLabel(3)).toBe('3 items won this tier');
  });
});

describe('raid mode loot line', () => {
  it('reads 7 of 8 raids', () => {
    expect(attendanceLabel(7, 8, 'raids')).toBe('7 of 8 raids');
    expect(attendanceLabel(1, 1, 'raids')).toBe('1 of 1 raid');
  });

  it('keeps report mode on raid nights', () => {
    expect(attendanceLabel(7, 8, 'reports')).toBe('7 of 8 raid nights');
  });
});
