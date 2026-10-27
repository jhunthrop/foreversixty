// web/src/lib/guild/reserve-keys.ts
// The guild pages' loading reserves kept in localStorage (raids spec 4 "Page states"): the band
// hold, the Raids tab's planned rows and the raid strip's rows. They describe a signed-in member's
// view, so sign-out forgets them all. Tiny on purpose: the account module and Guild.svelte's entry
// chunk import it.

/** `fs.guildBand.{region}.{ruleset}.{slug}`: the band header's last height per layout. */
export const GUILD_BAND_KEY_PREFIX = 'fs.guildBand.';
/** `fs.guildPlanned.{guildId}`: the planned rows the Raids tab last drew. */
export const GUILD_PLANNED_KEY_PREFIX = 'fs.guildPlanned.';
/** The own-character count the raid strip last drew. */
export const RAID_ROWS_KEY = 'fs.guildRaid.rows';
/** `fs.guildRaidDone.{raidId}`: the raid was last drawn `done`, so its loading holds the report. */
export const RAID_DONE_KEY_PREFIX = 'fs.guildRaidDone.';

function isGuildReserve(key: string): boolean {
  return (
    key.startsWith(GUILD_BAND_KEY_PREFIX) ||
    key.startsWith(GUILD_PLANNED_KEY_PREFIX) ||
    key.startsWith(RAID_DONE_KEY_PREFIX) ||
    key === RAID_ROWS_KEY
  );
}

/** The browser's localStorage, or `null` where it is missing or throws on access. */
export function browserStorage(): Storage | null {
  try {
    return typeof localStorage === 'undefined' ? null : localStorage;
  } catch {
    return null;
  }
}

/** Sign-out: forgets every guild reserve, so the next viewer on this browser holds nothing. */
export function forgetGuildReserves(storage: Storage | null = browserStorage()): void {
  try {
    if (storage === null) return;
    const keys = Array.from({ length: storage.length }, (_, index) => storage.key(index));
    for (const key of keys) if (key !== null && isGuildReserve(key)) storage.removeItem(key);
  } catch {
    // A blocked store holds nothing to forget.
  }
}
