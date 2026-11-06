// web/src/lib/guild/raid-ready-view.ts
//
// The officer Overview's Raid-ready (web layout spec 2026-11-05 §7.1), officer-only: for the next
// raid's In and Late raiders (GET .../readiness?raid=), how many are attuned (only when the zone
// has an attunement), each signed tank's fire resistance (only for a raid that asks for it) and
// how many carry a flask or raid elixir and buff food. Counts, never a worst-first list of
// people; a raider whose export cannot say is counted apart. Nothing tracked hides the module.

import { counted } from '../plural';
import type { GuildRaidSignup, GuildReadinessPage, GuildReadinessRow } from './api';

/** "X of N", with how many of the signed raiders the export cannot say for. */
export interface ReadyCount {
  yes: number;
  known: number;
  unknown: number;
}

/** One signed tank's fire resistance; `fire` null when their export does not carry it. */
export interface TankFire {
  key: string;
  name: string;
  class: string;
  fire: number | null;
}

export interface RaidReadyView {
  /** Null when the raid's zone has no attunement. */
  attuned: ReadyCount | null;
  /** Null when the raid does not ask for fire resistance; else every signed In or Late tank. */
  tanks: TankFire[] | null;
  flaskFood: ReadyCount;
}

const SIGNED = new Set(['in', 'late']);

function count<T>(
  rows: readonly GuildReadinessRow[],
  read: (row: GuildReadinessRow) => T | null | undefined,
  ok: (value: T) => boolean,
): ReadyCount {
  let yes = 0;
  let known = 0;
  for (const row of rows) {
    const value = read(row);
    if (value === null || value === undefined) continue;
    known += 1;
    if (ok(value)) yes += 1;
  }
  return { yes, known, unknown: rows.length - known };
}

function tankFire(rows: readonly GuildReadinessRow[], signups: readonly GuildRaidSignup[]): TankFire[] {
  const byKey = new Map(rows.map((row) => [row.character_key, row]));
  return signups
    .filter((s) => s.role === 'tank' && SIGNED.has(s.state))
    .map((s) => ({
      key: s.character_key,
      name: s.name,
      class: s.class,
      fire: byKey.get(s.character_key)?.resistances?.fire ?? null,
    }));
}

/**
 * Raid-ready for the next raid, or null when nothing is tracked: no signed raider's export
 * answers attunement, fire resistance or flask and food (exports from before the raid-prep
 * sections), or the page is not cut to a raid.
 */
export function raidReadyView(
  page: GuildReadinessPage,
  signups: readonly GuildRaidSignup[],
): RaidReadyView | null {
  if (page.raid === undefined) return null;
  const rows = page.rows;
  const attuned =
    page.raid.attunement === true
      ? count(
          rows,
          (r) => r.attuned,
          (v) => v,
        )
      : null;
  const tanks = page.raid.fire_resistance === true ? tankFire(rows, signups) : null;
  const flaskFood = count(
    rows,
    (r) => r.flask_food,
    (v) => v.flask_or_elixir && v.food,
  );
  const tracked =
    (attuned?.known ?? 0) > 0 || (tanks ?? []).some((t) => t.fire !== null) || flaskFood.known > 0;
  return tracked ? { attuned, tanks, flaskFood } : null;
}

/** "18 of 24". */
export function ofLine(c: ReadyCount): string {
  return `${c.yes} of ${c.known}`;
}

/** "3 have not shared their quests", "" when every raider's export answers. */
export function unknownLine(c: ReadyCount, what: 'quests' | 'bags'): string {
  if (c.unknown === 0) return '';
  return `${c.unknown} ${c.unknown === 1 ? 'has' : 'have'} not shared their ${what}`;
}

/** A tank's fire resistance line: "142 fire resistance" or "Not shared yet". */
export function tankFireLine(t: TankFire): string {
  return t.fire === null ? 'Not shared yet' : `${t.fire} fire resistance`;
}

/** The fire resistance block's meta: "2 tanks signed", or the empty raid's line. */
export function tanksMeta(tanks: readonly TankFire[]): string {
  return tanks.length === 0 ? 'No tank signed yet' : `${counted(tanks.length, 'tank')} signed`;
}
