// web/src/lib/guild/loot-view.test.ts
import { describe, expect, it } from 'vitest';
import type { GuildLootEncounter, GuildLootItem } from './api';
import { bossPickerLabel, candidateAction, candidateGainLabel, nextUnkilledEncounter } from './loot-view';

function encounter(overrides: Partial<GuildLootEncounter> = {}): GuildLootEncounter {
  return { encounter_id: 1084, name: 'Onyxia', zone: "Onyxia's Lair", killed: true, ...overrides };
}

function item(overrides: Partial<GuildLootItem> = {}): GuildLootItem {
  return {
    item_id: 16955,
    name: 'Helm of Wrath',
    icon: 'inv_helmet_71',
    quality: 4,
    slot: 'head',
    awarded_to: null,
    candidates: [
      {
        character_key: 'a',
        name: 'Kraggor',
        class: 'warrior',
        spec: 'Protection',
        gain_dps: 40,
        not_sim_checked: false,
        ilvl_delta: null,
        attendance: { present: 8, nights: 8 },
        already_equivalent: false,
      },
      {
        character_key: 'b',
        name: 'Grimtotem',
        class: 'warrior',
        spec: 'Fury',
        gain_dps: 12,
        not_sim_checked: false,
        ilvl_delta: null,
        attendance: { present: 6, nights: 8 },
        already_equivalent: false,
      },
    ],
    ...overrides,
  };
}

describe('bossPickerLabel', () => {
  it('reads "next unkilled: none, farm" for an already-killed boss', () => {
    expect(bossPickerLabel(encounter({ killed: true }))).toBe('Onyxia · next unkilled: none, farm');
  });

  it('reads "not yet killed" for an unkilled boss', () => {
    expect(bossPickerLabel(encounter({ killed: false }))).toBe('Onyxia · not yet killed');
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
  function candidate(overrides: Partial<GuildLootItem['candidates'][number]> = {}) {
    return {
      character_key: 'a',
      name: 'Kraggor',
      class: 'warrior',
      spec: 'Protection',
      gain_dps: null,
      not_sim_checked: true,
      ilvl_delta: null,
      attendance: { present: 8, nights: 8 },
      already_equivalent: false,
      ...overrides,
    };
  }

  it('reads a DPS figure for a tier-0 (BiS-matched) candidate', () => {
    expect(candidateGainLabel(candidate({ gain_dps: 40.4 }))).toBe('+40 DPS');
  });

  // Live-fix round, defect 5: every real candidate on a live roster is tier-1 today
  // (gain_dps null, ilvl_delta set) -- this must read the item level figure, never throw
  // calling a number method on the null gain_dps the way GuildLoot.svelte once did.
  it('reads a signed item-level figure for a tier-1 (fallback) candidate', () => {
    expect(candidateGainLabel(candidate({ gain_dps: null, ilvl_delta: 6 }))).toBe('+6 ilvl');
    expect(candidateGainLabel(candidate({ gain_dps: null, ilvl_delta: -3 }))).toBe('-3 ilvl');
  });

  it('reads an em dash when neither figure is present', () => {
    expect(candidateGainLabel(candidate({ gain_dps: null, ilvl_delta: null }))).toBe('—');
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
          ilvl_delta: null,
          attendance: { present: 8, nights: 8 },
          already_equivalent: true,
        },
      ],
    });
    expect(candidateAction(withEquivalent, 'c')).toEqual({ kind: 'equivalent' });
  });

  it('marks exactly the awardee "awarded" and every other candidate "awarded to {name}"', () => {
    const awarded = item({
      awarded_to: { character_key: 'a', name: 'Kraggor', at: '2026-12-15T00:00:00Z', by_name: 'Kraggor' },
    });
    expect(candidateAction(awarded, 'a')).toEqual({ kind: 'awarded' });
    expect(candidateAction(awarded, 'b')).toEqual({ kind: 'awarded-to-other', label: 'Awarded to Kraggor' });
  });
});
