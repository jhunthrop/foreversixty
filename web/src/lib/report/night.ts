// web/src/lib/report/night.ts
// The whole night in one table: every boss pull's summary folded into one row per player
// and one row per boss. A raid leader opens a log to compare the team across the night,
// not to sum eighteen pulls by hand. Trash is left out, as it is from the rankings: a
// trash pull's damage is mostly a function of how much trash there was.
import type {
  Actor,
  AuraTrack,
  CastRow,
  CombatantRow,
  Death,
  ExchangeRow,
  FightEntry,
  ResourceTrack,
  Summary,
  ThreatRow,
} from './types';

export interface NightPlayerFight {
  index: number;
  boss: string;
  kill: boolean;
  duration_ms: number;
  dps: number;
  hps: number;
  damage_taken: number;
  deaths: number;
}

export interface NightPlayer {
  guid: string;
  name: string;
  class?: string;
  spec?: string;
  role: string;
  /** Pulls this player was in. */
  fights: number;
  /** Their pulls' durations added up: what the per-second figures divide by. */
  time_ms: number;
  active_ms: number;
  damage_done: number;
  healing_done: number;
  damage_taken: number;
  deaths: number;
  dps: number;
  hps: number;
  dtps: number;
  activity_pct: number;
  by_fight: NightPlayerFight[];
}

export interface NightBoss {
  name: string;
  pulls: number;
  kills: number;
  wipes: number;
  time_ms: number;
  deaths: number;
  /** The quickest kill's index and length, when there was a kill. */
  best?: { index: number; duration_ms: number };
  /** The lowest the boss was brought to on a wipe, as a percentage; absent with no wipe seen. */
  lowest_wipe_pct?: number;
  fights: number[];
}

export interface Night {
  /** Boss pulls that had a summary to fold in. */
  fights: number;
  /** Boss pulls report.json lists, loaded or not. */
  expected: number;
  time_ms: number;
  deaths: number;
  players: NightPlayer[];
  bosses: NightBoss[];
}

function perSecond(total: number, ms: number): number {
  return ms <= 0 ? 0 : total / (ms / 1000);
}

/**
 * Folds the loaded boss summaries into the night. `summaries` may be missing fights
 * (one is still loading, one failed); those are counted in `expected` and left out of
 * every total, so a partial night is an honest partial rather than a wrong whole.
 */
export function aggregateNight(
  fights: readonly FightEntry[],
  summaries: ReadonlyMap<number, Summary>,
): Night {
  const encounters = fights.filter((fight) => fight.kind === 'encounter' && !fight.in_progress);
  const players = new Map<string, NightPlayer>();
  const bosses = new Map<string, NightBoss>();
  let time = 0;
  let deaths = 0;
  let loaded = 0;

  for (const fight of encounters) {
    const boss = bosses.get(fight.name) ?? {
      name: fight.name,
      pulls: 0,
      kills: 0,
      wipes: 0,
      time_ms: 0,
      deaths: 0,
      fights: [],
    };
    boss.pulls += 1;
    boss.kills += fight.kill ? 1 : 0;
    boss.wipes += fight.kill ? 0 : 1;
    boss.time_ms += fight.duration_ms;
    boss.deaths += fight.deaths;
    boss.fights.push(fight.index);
    if (fight.kill && (boss.best === undefined || fight.duration_ms < boss.best.duration_ms))
      boss.best = { index: fight.index, duration_ms: fight.duration_ms };
    const health = fight.boss_health_pct;
    if (!fight.kill && health !== undefined && health >= 0)
      boss.lowest_wipe_pct = Math.min(boss.lowest_wipe_pct ?? 100, health);
    bosses.set(fight.name, boss);

    const summary = summaries.get(fight.index);
    if (summary === undefined) continue;
    loaded += 1;
    time += summary.duration_ms;
    deaths += summary.deaths.length;
    for (const row of summary.roster) {
      const player = players.get(row.guid) ?? {
        guid: row.guid,
        name: row.name,
        class: row.class,
        spec: row.spec,
        role: row.role,
        fights: 0,
        time_ms: 0,
        active_ms: 0,
        damage_done: 0,
        healing_done: 0,
        damage_taken: 0,
        deaths: 0,
        dps: 0,
        hps: 0,
        dtps: 0,
        activity_pct: 0,
        by_fight: [],
      };
      player.fights += 1;
      player.time_ms += summary.duration_ms;
      player.active_ms += row.active_ms;
      player.damage_done += row.damage_done;
      player.healing_done += row.healing_done;
      player.damage_taken += row.damage_taken;
      player.deaths += row.deaths;
      // The latest spec wins: a player who respecced mid-night is filed under what they
      // finished as, which is what the next raid will see.
      if (row.spec) player.spec = row.spec;
      if (row.class) player.class = row.class;
      player.by_fight.push({
        index: fight.index,
        boss: fight.name,
        kill: fight.kill,
        duration_ms: summary.duration_ms,
        dps: row.dps,
        hps: row.hps,
        damage_taken: row.damage_taken,
        deaths: row.deaths,
      });
      players.set(row.guid, player);
    }
  }

  const rows = [...players.values()].map((player) => ({
    ...player,
    dps: perSecond(player.damage_done, player.time_ms),
    hps: perSecond(player.healing_done, player.time_ms),
    dtps: perSecond(player.damage_taken, player.time_ms),
    activity_pct: player.time_ms <= 0 ? 0 : (player.active_ms / player.time_ms) * 100,
  }));
  rows.sort((a, b) => b.damage_done - a.damage_done || a.name.localeCompare(b.name));

  return {
    fights: loaded,
    expected: encounters.length,
    time_ms: time,
    deaths,
    players: rows,
    bosses: [...bosses.values()],
  };
}

/**
 * Every loaded boss pull folded into one Summary-shaped object, so the fight tabs can
 * show the night the way they show a pull: damage by ability across every boss, deaths
 * in the order they happened with the pull they happened in, interrupts and the casts
 * that went through, buff uptime over the night's combat time. Timestamps are offset by
 * the pulls before them, so the night reads as one long fight; the per-second series are
 * left empty, since nothing draws a chart over a night.
 */
export function nightSummary(
  fights: readonly FightEntry[],
  summaries: ReadonlyMap<number, Summary>,
): Summary {
  const encounters = fights.filter((fight) => fight.kind === 'encounter' && !fight.in_progress);
  const pulls = pullNumbersOf(encounters);
  const actorsBy = {
    damage_done: new Map<string, Actor>(),
    damage_taken: new Map<string, Actor>(),
    healing: new Map<string, Actor>(),
    healing_taken: new Map<string, Actor>(),
  };
  const deaths: Death[] = [];
  const auras = new Map<string, AuraTrack>();
  const casts = new Map<string, CastRow>();
  const exchanges = new Map<string, ExchangeRow>();
  const threat = new Map<string, ThreatRow>();
  const resources = new Map<string, ResourceTrack>();
  const combatants = new Map<string, CombatantRow>();
  let offset = 0;
  let engine = '';

  for (const fight of encounters) {
    const summary = summaries.get(fight.index);
    if (summary === undefined) continue;
    engine = summary.engine_version;
    const label = `${fight.name} · pull ${pulls.get(fight.index) ?? 1}`;
    for (const table of ['damage_done', 'damage_taken', 'healing', 'healing_taken'] as const) {
      for (const actor of summary[table]) mergeActor(actorsBy[table], actor);
    }
    for (const death of summary.deaths) deaths.push({ ...death, at_ms: death.at_ms + offset, label });
    // Folded by the target's name, not its GUID: a boss is a new GUID every pull, and
    // three "Nalthor" rows at a third of the uptime each are one row at the whole. The
    // denominator is the time that name was in a pull, counted once per pull.
    const pullsCounted = new Set<string>();
    for (const track of summary.auras) {
      const key = `${track.target_name}|${track.spell_id}`;
      const found = auras.get(key);
      const shifted = track.segments.map((segment) => ({
        ...segment,
        start_ms: segment.start_ms + offset,
        end_ms: segment.end_ms + offset,
      }));
      // Counted once per pull per track, not per target: two spells on one boss each
      // need the pull's time added to their own denominator.
      const countPull = !pullsCounted.has(key);
      pullsCounted.add(key);
      if (found === undefined) auras.set(key, { ...track, segments: shifted, time_ms: summary.duration_ms });
      else
        auras.set(key, {
          ...found,
          applications: found.applications + track.applications,
          max_stacks: Math.max(found.max_stacks, track.max_stacks),
          uptime_ms: found.uptime_ms + track.uptime_ms,
          segments: [...found.segments, ...shifted],
          appliers: [...new Set([...found.appliers, ...track.appliers])],
          time_ms: (found.time_ms ?? 0) + (countPull ? summary.duration_ms : 0),
        });
    }
    for (const row of summary.casts) {
      const key = `${row.guid}|${row.spell_id}`;
      const found = casts.get(key);
      const sequence = row.sequence.map((at) => at + offset);
      if (found === undefined) casts.set(key, { ...row, sequence });
      else
        casts.set(key, {
          ...found,
          started: found.started + row.started,
          succeeded: found.succeeded + row.succeeded,
          failed: found.failed + row.failed,
          cast_time_ms: found.cast_time_ms + row.cast_time_ms,
          sequence: [...found.sequence, ...sequence],
        });
    }
    for (const row of [...summary.interrupts, ...summary.dispels]) {
      const key = `${row.kind}|${row.source_guid}|${row.target_guid}|${row.spell_id}|${row.extra_spell_id}`;
      const found = exchanges.get(key);
      exchanges.set(key, found === undefined ? { ...row } : { ...found, count: found.count + row.count });
    }
    for (const row of summary.threat) {
      const found = threat.get(row.guid);
      threat.set(
        row.guid,
        found === undefined ? { ...row } : { ...found, threat: found.threat + row.threat },
      );
    }
    for (const row of summary.combatants) combatants.set(row.guid, row);
    for (const track of summary.resources) {
      const rkey = `${track.guid}|${track.power_type}`;
      const found = resources.get(rkey);
      // The night's series is the pulls' series end to end, with the gaps between pulls
      // (a player absent from a pull) padded with the last reading so the line holds.
      const padTo = offset / 1000;
      if (found === undefined) {
        const lead = new Array<number>(Math.max(0, Math.round(padTo))).fill(0);
        resources.set(rkey, { ...track, series: [...lead, ...track.series] });
      } else {
        const last = found.series[found.series.length - 1] ?? 0;
        const gap = Math.max(0, Math.round(padTo) - found.series.length);
        resources.set(rkey, {
          ...found,
          series: [...found.series, ...new Array<number>(gap).fill(last), ...track.series],
          gained: found.gained + track.gained,
          spent: found.spent + track.spent,
          zero_ms: found.zero_ms + track.zero_ms,
        });
      }
    }
    offset += summary.duration_ms;
  }

  const night = aggregateNight(fights, summaries);
  // A player who left after two pulls is measured over two pulls, not the night: the
  // Summary tab already does, and the two tabs must agree.
  const timeOf = new Map(night.players.map((player) => [player.guid, player.time_ms]));
  const sorted = (table: Map<string, Actor>): Actor[] =>
    [...table.values()]
      .map((actor) => ({ ...actor, time_ms: timeOf.get(actor.guid) }))
      .sort((a, b) => b.effective - a.effective || a.guid.localeCompare(b.guid));
  return {
    engine_version: engine,
    fight_index: 0,
    duration_ms: offset,
    damage_done: sorted(actorsBy.damage_done),
    damage_taken: sorted(actorsBy.damage_taken),
    healing: sorted(actorsBy.healing),
    healing_taken: sorted(actorsBy.healing_taken),
    deaths,
    auras: [...auras.values()],
    casts: [...casts.values()],
    interrupts: [...exchanges.values()].filter((row) => row.kind === 'interrupt'),
    dispels: [...exchanges.values()].filter((row) => row.kind === 'dispel'),
    resources: [...resources.values()],
    threat: [...threat.values()],
    combatants: [...combatants.values()],
    roster: night.players.map((player) => ({
      guid: player.guid,
      name: player.name,
      class: player.class,
      spec: player.spec,
      role: player.role,
      active_ms: player.active_ms,
      activity_pct: player.activity_pct,
      deaths: player.deaths,
      damage_done: player.damage_done,
      healing_done: player.healing_done,
      damage_taken: player.damage_taken,
      dps: player.dps,
      hps: player.hps,
      dtps: player.dtps,
    })),
  };
}

/** Pull numbers per boss, in fight order, for the death labels. */
function pullNumbersOf(encounters: readonly FightEntry[]): Map<number, number> {
  const counts = new Map<string, number>();
  const out = new Map<number, number>();
  for (const fight of encounters) {
    const pull = (counts.get(fight.name) ?? 0) + 1;
    counts.set(fight.name, pull);
    out.set(fight.index, pull);
  }
  return out;
}

/** Adds an actor's totals, abilities and targets into the table's row for that GUID. */
function mergeActor(table: Map<string, Actor>, actor: Actor): void {
  const found = table.get(actor.guid);
  if (found === undefined) {
    table.set(actor.guid, {
      ...actor,
      abilities: actor.abilities.map((a) => ({ ...a })),
      targets: actor.targets.map((t) => ({ ...t })),
      series: [],
    });
    return;
  }
  const abilities = new Map(found.abilities.map((ability) => [ability.spell_id, { ...ability }]));
  for (const ability of actor.abilities) {
    const have = abilities.get(ability.spell_id);
    if (have === undefined) abilities.set(ability.spell_id, { ...ability });
    else
      abilities.set(ability.spell_id, {
        ...have,
        total: have.total + ability.total,
        effective: have.effective + ability.effective,
        overheal: sumOptional(have.overheal, ability.overheal),
        overkill: sumOptional(have.overkill, ability.overkill),
        absorbed: sumOptional(have.absorbed, ability.absorbed),
        blocked: sumOptional(have.blocked, ability.blocked),
        hits: have.hits + ability.hits,
        crits: have.crits + ability.crits,
        ticks: have.ticks + ability.ticks,
        min: have.min === 0 ? ability.min : ability.min === 0 ? have.min : Math.min(have.min, ability.min),
        max: Math.max(have.max, ability.max),
        misses: mergeCounts(have.misses, ability.misses),
      });
  }
  const targets = new Map(found.targets.map((target) => [target.guid, { ...target }]));
  for (const target of actor.targets) {
    const have = targets.get(target.guid);
    targets.set(
      target.guid,
      have === undefined ? { ...target } : { ...have, total: have.total + target.total },
    );
  }
  table.set(actor.guid, {
    ...found,
    total: found.total + actor.total,
    effective: found.effective + actor.effective,
    overheal: sumOptional(found.overheal, actor.overheal),
    absorbed: sumOptional(found.absorbed, actor.absorbed),
    active_ms: found.active_ms + actor.active_ms,
    abilities: [...abilities.values()],
    targets: [...targets.values()],
  });
}

function sumOptional(a: number | undefined, b: number | undefined): number | undefined {
  if (a === undefined && b === undefined) return undefined;
  return (a ?? 0) + (b ?? 0);
}

function mergeCounts(
  a: Record<string, number> | undefined,
  b: Record<string, number> | undefined,
): Record<string, number> | undefined {
  if (a === undefined && b === undefined) return undefined;
  const out: Record<string, number> = { ...(a ?? {}) };
  for (const [key, count] of Object.entries(b ?? {})) out[key] = (out[key] ?? 0) + count;
  return out;
}
