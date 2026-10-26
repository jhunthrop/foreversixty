// web/src/lib/guild/planned-reserve.ts
// The Raids tab's loading reserve (raids spec 4, "Raids tab loading"): the planned rows last
// drawn for this guild, kept in localStorage so the skeleton holds their height. Tiny on purpose:
// Guild.svelte's entry chunk imports it.
import type { GuildRaidsPage } from './api';
import { GUILD_PLANNED_KEY_PREFIX } from './reserve-keys';

/** The reserve holds at most the four rows `Planned` shows before `Show all`. */
const RESERVE_MAX_ROWS = 4;
const RESERVE_DEFAULT_ROWS = 1;

function reserveKey(guildId: number): string {
  return `${GUILD_PLANNED_KEY_PREFIX}${guildId}`;
}

/** The planned rows to reserve while the tab loads: the count last rendered (0–4), else one. */
export function plannedReserveRows(guildId: number, storage: Pick<Storage, 'getItem'> | null): number {
  try {
    const raw = storage?.getItem(reserveKey(guildId)) ?? null;
    const rows = raw === null ? Number.NaN : Number(raw);
    return Number.isInteger(rows) && rows >= 0 && rows <= RESERVE_MAX_ROWS ? rows : RESERVE_DEFAULT_ROWS;
  } catch {
    return RESERVE_DEFAULT_ROWS;
  }
}

/** Stores the rows a loaded page draws; a page without `planned` (a visitor) stores nothing. */
export function rememberPlannedRows(
  guildId: number,
  page: GuildRaidsPage,
  storage: Pick<Storage, 'setItem'> | null,
): void {
  if (page.planned === undefined) return;
  try {
    storage?.setItem(reserveKey(guildId), String(Math.min(page.planned.length, RESERVE_MAX_ROWS)));
  } catch {
    // A blocked store only costs the next visit its warm reserve.
  }
}
