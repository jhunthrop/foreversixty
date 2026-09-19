import { describe, expect, it } from 'vitest';
import generated from '../../data/generated/sim-ids.json';
import { simCopy } from './copy';
import { PINNED_STATS, SIM_STATS, statLabel } from './stats';

describe('PINNED_STATS', () => {
  it('is contract 10.8’s list, verbatim and in its order', () => {
    expect(PINNED_STATS).toEqual([
      'strength',
      'agility',
      'stamina',
      'intellect',
      'spirit',
      'spell_power',
      'arcane_power',
      'fire_power',
      'frost_power',
      'holy_power',
      'nature_power',
      'shadow_power',
      'mp5',
      'hit',
      'crit',
      'spell_haste',
      'spell_penetration',
      'attack_power',
      'melee_haste',
      'armor_penetration',
      'expertise',
      'mana',
      'energy',
      'rage',
      'armor',
      'ranged_attack_power',
      'defense',
      'block',
      'block_value',
      'dodge',
      'parry',
      'health',
      'arcane_resistance',
      'fire_resistance',
      'frost_resistance',
      'nature_resistance',
      'shadow_resistance',
      'bonus_armor',
      'healing_power',
      'spell_damage',
      'feral_attack_power',
    ]);
  });

  it('carries one hit and one crit, and splits haste, which is 10.8’s whole point', () => {
    expect(PINNED_STATS).toContain('hit');
    expect(PINNED_STATS).toContain('crit');
    expect(PINNED_STATS).toContain('spell_haste');
    expect(PINNED_STATS).toContain('melee_haste');
    expect(PINNED_STATS).toContain('mp5');
    for (const forbidden of ['melee_hit', 'spell_hit', 'melee_crit', 'spell_crit', 'haste']) {
      expect(PINNED_STATS, forbidden).not.toContain(forbidden);
    }
  });

  it('names every one of them, so no weights row ever renders a raw id', () => {
    for (const stat of PINNED_STATS) {
      expect(simCopy.statLabel[stat], stat).toBeTruthy();
    }
  });

  it('names nothing 10.8 does not list, so a stale label cannot outlive its stat', () => {
    for (const named of Object.keys(simCopy.statLabel)) {
      expect(PINNED_STATS, named).toContain(named);
    }
  });
});

describe('SIM_STATS', () => {
  it('is the pinned list until IDS.md publishes a Stats section', () => {
    expect(SIM_STATS).toEqual(generated.stats.length > 0 ? generated.stats : PINNED_STATS);
  });

  it('agrees with the pinned list whenever IDS.md does publish one', () => {
    // The generated section and 10.8's pinning are the same vocabulary from two
    // directions; if they ever disagree, one of them is wrong and this says so loudly
    // rather than letting the page offer a stat the engine cannot weigh.
    if (generated.stats.length > 0) expect([...generated.stats].sort()).toEqual([...PINNED_STATS].sort());
  });
});

describe('statLabel', () => {
  it('prefers the copy table’s name', () => {
    expect(statLabel('attack_power')).toBe(simCopy.statLabel.attack_power);
    expect(statLabel('spell_haste')).toBe(simCopy.statLabel.spell_haste);
    expect(statLabel('mp5')).toBe(simCopy.statLabel.mp5);
  });

  it('humanises an id the copy table does not name, rather than showing the id raw', () => {
    // Not simCopy.statLabel here: 'some_new_stat' is deliberately absent from the copy
    // table, so this literal is the only thing pinning the humanisation rule. Asserting
    // it against the copy table instead would make this tautological and let the
    // humanisation logic rot unnoticed.
    expect(statLabel('some_new_stat')).toBe('Some new stat');
  });
});
