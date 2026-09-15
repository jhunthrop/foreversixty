// web/src/lib/report/night.ts
// The whole night in one table: every boss pull's summary folded into one row per player
// and one row per boss. A raid leader opens a log to compare the team across the night,
// not to sum eighteen pulls by hand. Trash is left out, as it is from the rankings: a
// trash pull's damage is mostly a function of how much trash there was.
import type { FightEntry, Summary } from './types';

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
