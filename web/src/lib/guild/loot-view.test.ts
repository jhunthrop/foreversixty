// web/src/lib/guild/loot-view.test.ts
import { describe, expect, it } from 'vitest';
import type { GuildLootCandidate } from './api';
import { candidateGainLabel } from './loot-view';

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
