// web/src/lib/sim/stats.ts
// The stat vocabulary: the fork's `proto.Stat` enum names in snake case, pinned verbatim
// by contract 10.8 and generated into IDS.md's Stats section by the sim module (A7).
//
// Two things about this list are easy to get wrong and are wrong everywhere else:
//
//   * the engine carries ONE `hit` and ONE `crit`. There is no `melee_hit`, `spell_hit`,
//     `melee_crit` or `spell_crit`, and a page that offered either pair would be asking
//     for a weight the engine cannot compute;
//   * haste IS split -- `spell_haste` and `melee_haste` -- so there is no bare `haste`.
//
// `MP5` is spelled `mp5`. 10.8 also fixes the reference defaults the weights page starts
// on: `attack_power` for melee and hunters, `spell_power` for casters, served per spec as
// `reference_stat` on GET /v1/specs, so no page hard-codes one.
//
// This is here rather than in buffs.ts because a stat is not a buff: buffs.ts owns the
// panel's grouping and nothing else. Part B's /sim/weights is the only reader.
import generated from '../../data/generated/sim-ids.json';
import { simCopy } from './copy';

/** Contract 10.8's list, verbatim and in its order. */
export const PINNED_STATS: readonly string[] = [
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
];

/**
 * The generated section when IDS.md has one, the pinned list until it does -- the same
 * rule `buffs.ts` uses for the world buffs, and for the same reason: the generator is the
 * long-term source of truth and the pinning is what makes the page correct today.
 * stats.test.ts asserts the two agree whenever both exist.
 */
const published = (generated as { stats: string[] }).stats;

export const SIM_STATS: readonly string[] = published.length > 0 ? published : PINNED_STATS;

/** "attack_power" reads as "Attack power". The copy table first, then plain casing. */
export function statLabel(id: string): string {
  const named = simCopy.statLabel[id];
  if (named !== undefined) return named;
  const words = id.split('_').join(' ');
  return words.charAt(0).toUpperCase() + words.slice(1);
}
