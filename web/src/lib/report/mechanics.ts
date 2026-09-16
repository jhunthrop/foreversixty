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
  /**
   * Avoidable hits on them, most damage first. An entry with `others` is not a hit on
   * them: it is an ability their role was meant to take landing on that many other
   * players, which is their problem as much as the victims' (a cleave the tank faced
   * into the raid).
   */
  hits: { row: MechanicRow; hit: MechanicHit; others?: number }[];
  /** Their avoidable damage: what landed on them. */
  damage: number;
  /** Damage their role's ability put on other players (the `others` entries): not theirs taken, theirs placed. */
  placed: number;
  /** What the fight did to them regardless: the table's unavoidable rows they are on. */
  unavoidable: { damage: number; names: string[] };
  /** Avoidable mechanics that never touched them, by name. */
  avoided: string[];
}

interface PlayerTally {
  guid: string;
  name: string;
  hits: { row: MechanicRow; hit: MechanicHit; others?: number }[];
  damage: number;
  placed: number;
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
  roster: readonly { guid: string; name: string; role?: string }[],
): PlayerMechanics[] {
  const tallies = new Map<string, PlayerTally>();
  const tally = (guid: string, name: string): PlayerTally =>
    tallies.get(guid) ?? {
      guid,
      name,
      hits: [],
      damage: 0,
      placed: 0,
      unavoidableDamage: 0,
      unavoidableNames: [],
    };
  const roleOf = new Map(roster.map((player) => [player.guid, player.role]));
  for (const row of rows) {
    if (row.kind !== 'avoidable' && row.kind !== 'unavoidable') continue;
    for (const hit of row.players ?? []) {
      const found = tally(hit.guid, hit.name);
      // An ability the table assigns to a role is that role's to take: on the tank, the
      // cleave is unavoidable damage, not a failure; on anyone else it is avoidable.
      const avoidable =
        row.kind === 'avoidable' && (row.role === undefined || roleOf.get(hit.guid) !== row.role);
      tallies.set(
        hit.guid,
        avoidable
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

  // An ability one role is meant to take, landing on anyone else, goes on that role's
  // card as a ranked entry rather than under "never hit by" as praise: a tank who faced
  // the cleave into four people did not avoid it.
  for (const placement of rolePlacements(rows, roster)) {
    const found = tally(placement.hit.guid, placement.hit.name);
    tallies.set(placement.hit.guid, {
      ...found,
      hits: [...found.hits, placement],
      placed: found.placed + placement.hit.damage,
    });
  }

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
        placed: player.placed,
        unavoidable: {
          damage: player.unavoidableDamage,
          names: [...new Set(player.unavoidableNames)],
        },
        avoided: avoidableNames.filter((name) => !hitNames.has(name)),
      };
    })
    .sort((a, b) => b.damage + b.placed - (a.damage + a.placed) || a.name.localeCompare(b.name));
}

/** A role's ability placed on other players, as an entry on the role holder's account. */
export interface RolePlacement {
  row: MechanicRow;
  /** The holder's guid and name, with the strays' hits, damage, span and deaths folded together. */
  hit: MechanicHit;
  /** How many players not of the role it landed on. */
  others: number;
}

/**
 * For every avoidable row the table assigns to a role, what it did to players not of
 * that role, charged to each holder of the role: the cards and the problems list both
 * read it from here, so the two never disagree about whose the cleave was.
 */
export function rolePlacements(
  rows: readonly MechanicRow[],
  roster: readonly { guid: string; name: string; role?: string }[],
): RolePlacement[] {
  const out: RolePlacement[] = [];
  for (const row of rows) {
    if (row.kind !== 'avoidable' || row.role === undefined) continue;
    const owners = roster.filter((player) => player.role === row.role);
    const ownerGuids = new Set(owners.map((player) => player.guid));
    const strays = (row.players ?? []).filter((hit) => !ownerGuids.has(hit.guid));
    if (strays.length === 0) continue;
    const folded = {
      hits: strays.reduce((sum, stray) => sum + stray.hits, 0),
      damage: strays.reduce((sum, stray) => sum + stray.damage, 0),
      first_ms: Math.min(...strays.map((stray) => stray.first_ms)),
      last_ms: Math.max(...strays.map((stray) => stray.last_ms)),
      killed: strays.some((stray) => stray.killed),
    };
    for (const owner of owners) {
      out.push({ row, hit: { ...folded, guid: owner.guid, name: owner.name }, others: strays.length });
    }
  }
  return out;
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
 * the pull it came from -- so the caller says so beside the list. A spell a roster player
 * dealt damage with is a player's, not a mechanic: friendly fire (a totem hit, a mind
 * controlled raider's spell) stays off the list, since a damage-taken ability carries no
 * source to tell it apart by.
 */
/** What the enemies' melee swings did to players: the fight's baseline, never a mechanic. */
export interface MeleeBucket {
  /** Effective damage of every enemy swing on a roster player. */
  damage: number;
  /** How many players it hit. */
  players: number;
  /** Who took the most of it, by name, with their share of it. */
  most?: { name: string; damage: number };
  /** What landed on everyone else: the swings the tank did not take. */
  others: { damage: number; players: number };
}

/**
 * Every melee swing on a player, in one bucket: the table never lists it, since a swing
 * is what a boss does to whoever holds it, but leaving it out made the mode account for
 * a third of what landed and read as if the rest never happened.
 */
export function meleeBucket(summary: Summary): MeleeBucket {
  const players = new Set(summary.roster.map((row) => row.guid));
  let damage = 0;
  let hit = 0;
  let most: MeleeBucket['most'];
  for (const actor of summary.damage_taken) {
    if (!players.has(actor.guid)) continue;
    const own = actor.abilities
      .filter((ability) => ability.spell_id === MELEE_SPELL_ID)
      .reduce((sum, ability) => sum + ability.effective, 0);
    if (own <= 0) continue;
    damage += own;
    hit += 1;
    if (most === undefined || own > most.damage) most = { name: actor.name, damage: own };
  }
  return {
    damage,
    players: hit,
    most,
    others: {
      damage: damage - (most?.damage ?? 0),
      players: Math.max(hit - (most === undefined ? 0 : 1), 0),
    },
  };
}

/** A death whose killing blow is on no row of the table: the mode cannot judge it, so it says so. */
export interface UnjudgedDeath {
  guid: string;
  name: string;
  at_ms: number;
  /** The death's instant on the summary's own clock (the night's, over a night): the Deaths tab's key. */
  night_ms: number;
  /** What killed them, as the death card names it. */
  by: string;
  label?: string;
}

/**
 * The deaths the problems list leaves out: killed by a swing, by an ability the table
 * does not list, or by nothing the log named. Listed rather than dropped, since "three
 * of nine deaths" reads as six that went fine.
 */
export function unjudgedDeaths(summary: Summary): UnjudgedDeath[] {
  const classified = new Set((summary.mechanics?.rows ?? []).map((row) => row.spell_id));
  const players = new Set(summary.roster.map((row) => row.guid));
  return summary.deaths
    .filter((death) => players.has(death.guid))
    .filter((death) => death.killing_blow === undefined || !classified.has(death.killing_blow.spell_id))
    .map((death) => {
      const blow = death.killing_blow;
      const spell = blow === undefined ? '' : blow.spell_id === MELEE_SPELL_ID ? 'melee' : blow.spell_name;
      const source = blow?.source_name ?? '';
      const by = blow === undefined ? 'nothing the log named' : [source, spell].filter(Boolean).join(' · ');
      // On the night the death's instant is on the night's clock; the pull's own clock is
      // what the Deaths tab shows, so the two agree.
      const pull = (summary.pulls ?? []).find(
        (mark) => mark.start_ms <= death.at_ms && death.at_ms < mark.end_ms,
      );
      return {
        guid: death.guid,
        name: death.name,
        at_ms: death.at_ms - (pull?.start_ms ?? 0),
        night_ms: death.at_ms,
        by,
        label: death.label,
      };
    });
}

export function unclassifiedAbilities(summary: Summary): UnclassifiedAbility[] {
  const classified = new Set((summary.mechanics?.rows ?? []).map((row) => row.spell_id));
  const players = new Set(summary.roster.map((row) => row.guid));
  const playerSpells = new Set(
    summary.damage_done
      .filter((actor) => players.has(actor.guid))
      .flatMap((actor) => actor.abilities.map((ability) => ability.spell_id)),
  );
  const damage = new Map<number, number>();
  const names = new Map<number, string>();
  const hit = new Map<number, Set<string>>();
  for (const actor of summary.damage_taken) {
    if (!players.has(actor.guid)) continue;
    for (const ability of actor.abilities) {
      if (ability.spell_id === MELEE_SPELL_ID || classified.has(ability.spell_id)) continue;
      if (playerSpells.has(ability.spell_id)) continue;
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
