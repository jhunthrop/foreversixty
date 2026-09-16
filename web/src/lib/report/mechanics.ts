// web/src/lib/report/mechanics.ts
// The arithmetic behind Mechanics mode, kept out of the component so it can be read and
// tested on its own. The mode's job is to say what the encounter's curated table judges
// about a fight -- and, where the table is silent, to say that it is silent rather than
// let the reader mistake the table for the whole of what happened.
import type { MechanicHit, MechanicRow, MechanicsBlock, Summary } from './types';

/**
 * A boss's identity inside one mechanics block. A night folds several bosses and two of
 * them can list the same spell id, so nothing here may be keyed on the spell alone. The
 * encounter id is the real key; the name stands in for a pull the report gave no id.
 */
export function encounterKey(encounterId: number | undefined, name: string | undefined): string {
  return encounterId === undefined ? `name:${name ?? ''}` : `encounter:${encounterId}`;
}

/** One row's identity: its boss and its spell. On a single fight the boss half is empty. */
export function mechanicRowKey(row: Pick<MechanicRow, 'spell_id' | 'encounter_id' | 'encounter'>): string {
  return `${encounterKey(row.encounter_id, row.encounter)}|${row.spell_id}`;
}

/**
 * Who a mechanic caught most often over the night, or undefined when nobody leads.
 * Sorted by pulls and then by damage, so the answer is the record rather than the order
 * the fold happened to merge the players in; a genuine tie names nobody.
 */
export function mostOftenHit(players: readonly MechanicHit[] | undefined): MechanicHit | undefined {
  const sorted = [...(players ?? [])].sort((a, b) => (b.pulls ?? 0) - (a.pulls ?? 0) || b.damage - a.damage);
  const [first, second] = sorted;
  if (first === undefined) return undefined;
  if (second !== undefined && (second.pulls ?? 0) === (first.pulls ?? 0) && second.damage === first.damage)
    return undefined;
  return first;
}

/** The rows nothing went wrong on: every row whose key is not among the failures. */
export function cleanMechanicRows(
  rows: readonly MechanicRow[],
  failedKeys: readonly string[],
): MechanicRow[] {
  const failed = new Set(failedKeys);
  return rows.filter((row) => !failed.has(mechanicRowKey(row)));
}

/** One boss's avoidable rows over the night, with the pulls they are counted out of. */
export interface NightBossMechanics {
  key: string;
  name: string;
  /** That boss's pulls folded into the night: the denominator of "6 of 18 pulls". */
  pulls: number;
  rows: MechanicRow[];
}

/**
 * The night's avoidable rows that hit somebody, grouped under the boss whose table lists
 * them, in the order the night met the bosses. A mechanic that hit nobody all night is
 * not a "0 pulls" line here; it belongs under "also in the table".
 */
export function nightBossGroups(block: MechanicsBlock): NightBossMechanics[] {
  const groups = new Map<string, NightBossMechanics>();
  for (const boss of block.bosses ?? []) {
    const key = encounterKey(boss.encounter_id, boss.name);
    groups.set(key, { key, name: boss.name, pulls: boss.pulls, rows: [] });
  }
  for (const row of block.rows) {
    if (row.kind !== 'avoidable' || (row.pulls_hit ?? 0) === 0) continue;
    const key = encounterKey(row.encounter_id, row.encounter);
    const found = groups.get(key) ?? { key, name: row.encounter ?? '', pulls: 0, rows: [] };
    groups.set(key, { ...found, rows: [...found.rows, row] });
  }
  return [...groups.values()].filter((group) => group.rows.length > 0);
}

/** One player's record against the table: what caught them, and what never did. */
export interface PlayerMechanics {
  guid: string;
  name: string;
  /** Avoidable hits on them, most damage first. */
  hits: { row: MechanicRow; hit: MechanicHit }[];
  /** Their avoidable damage. */
  damage: number;
  /** What the fight did to them regardless: the table's unavoidable rows they are on. */
  unavoidable: { damage: number; names: string[] };
  /** Avoidable mechanics that never touched them, by name. */
  avoided: string[];
}

interface PlayerTally {
  guid: string;
  name: string;
  hits: { row: MechanicRow; hit: MechanicHit }[];
  damage: number;
  unavoidableDamage: number;
  unavoidableNames: string[];
}

/**
 * A card's worth of record per player: their avoidable hits, the unavoidable damage the
 * table accounts for, and the mechanics they stayed out of. Both halves are the point --
 * a card with only the failures reads as an accusation rather than a record -- so every
 * roster player gets one whether or not anything landed on them.
 */
export function playerMechanics(
  rows: readonly MechanicRow[],
  roster: readonly { guid: string; name: string }[],
): PlayerMechanics[] {
  const tallies = new Map<string, PlayerTally>();
  const tally = (guid: string, name: string): PlayerTally =>
    tallies.get(guid) ?? { guid, name, hits: [], damage: 0, unavoidableDamage: 0, unavoidableNames: [] };
  for (const row of rows) {
    if (row.kind !== 'avoidable' && row.kind !== 'unavoidable') continue;
    for (const hit of row.players ?? []) {
      const found = tally(hit.guid, hit.name);
      tallies.set(
        hit.guid,
        row.kind === 'avoidable'
          ? { ...found, hits: [...found.hits, { row, hit }], damage: found.damage + hit.damage }
          : {
              ...found,
              unavoidableDamage: found.unavoidableDamage + hit.damage,
              unavoidableNames: [...found.unavoidableNames, row.name],
            },
      );
    }
  }
  for (const player of roster)
    if (!tallies.has(player.guid)) tallies.set(player.guid, tally(player.guid, player.name));

  const avoidableNames = [...new Set(rows.filter((row) => row.kind === 'avoidable').map((row) => row.name))];
  return [...tallies.values()]
    .map((player) => {
      // By name, not by row: over the night one mechanic is one row per boss, and
      // "never hit by Wicked Gash" beside "hit by Wicked Gash" is not an answer.
      const hitNames = new Set(player.hits.map((entry) => entry.row.name));
      return {
        guid: player.guid,
        name: player.name,
        hits: [...player.hits].sort((a, b) => b.hit.damage - a.hit.damage),
        damage: player.damage,
        unavoidable: {
          damage: player.unavoidableDamage,
          names: [...new Set(player.unavoidableNames)],
        },
        avoided: avoidableNames.filter((name) => !hitNames.has(name)),
      };
    })
    .sort((a, b) => b.damage - a.damage || a.name.localeCompare(b.name));
}

/** An enemy ability that hit a player and the encounter's table says nothing about. */
export interface UnclassifiedAbility {
  spell_id: number;
  name: string;
  /** Effective damage it did to players over the whole fight, or the whole night. */
  damage: number;
  /** How many players it hit. */
  players: number;
}

/** The melee swing carries no spell, so there is no mechanic for a table to classify. */
const MELEE_SPELL_ID = 0;

/**
 * Every ability that hit a player and is on no row of the table. The spec's honesty
 * rule: a table is one person's list, not the whole of what a boss does, and an ability
 * missing from it must be listed as unjudged rather than quietly dropped. Over the night
 * the rows come from every boss at once -- a folded damage-taken ability does not carry
 * the pull it came from -- so the caller says so beside the list.
 */
export function unclassifiedAbilities(summary: Summary): UnclassifiedAbility[] {
  const classified = new Set((summary.mechanics?.rows ?? []).map((row) => row.spell_id));
  const players = new Set(summary.roster.map((row) => row.guid));
  const damage = new Map<number, number>();
  const names = new Map<number, string>();
  const hit = new Map<number, Set<string>>();
  for (const actor of summary.damage_taken) {
    if (!players.has(actor.guid)) continue;
    for (const ability of actor.abilities) {
      if (ability.spell_id === MELEE_SPELL_ID || classified.has(ability.spell_id)) continue;
      damage.set(ability.spell_id, (damage.get(ability.spell_id) ?? 0) + ability.effective);
      if (ability.name) names.set(ability.spell_id, ability.name);
      const guids = hit.get(ability.spell_id) ?? new Set<string>();
      guids.add(actor.guid);
      hit.set(ability.spell_id, guids);
    }
  }
  return [...damage.entries()]
    .map(([spell_id, total]) => ({
      spell_id,
      name: names.get(spell_id) ?? `Spell #${spell_id}`,
      damage: total,
      players: hit.get(spell_id)?.size ?? 0,
    }))
    .sort((a, b) => b.damage - a.damage || a.spell_id - b.spell_id);
}
