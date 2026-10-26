// web/src/lib/guild/band-hold.ts
// The guild band's member lines (standing, raid line, needs line, settings link, and the claim line
// `/v1/me` adds) arrive after the band itself has painted. A viewer last seen verified gets the
// band header's last height, per guild and per layout, held while `/home` is in flight (raids spec
// 4 "Page states", the guild band), so the tabs and everything under them do not move.
// Tiny on purpose: Guild.svelte's entry chunk imports it.

/** Below md the band stacks its lines; from md they run wider, so each layout keeps its own height. */
export type BandLayout = 'phone' | 'desktop';

/** No member block is taller than this; a larger stored value is not trusted. */
const MAX_HOLD_PX = 1200;

export function bandHoldKey(region: string, ruleset: string, slug: string): string {
  return `fs.guildBand.${region}.${ruleset}.${slug}`;
}

function readAll(storage: Pick<Storage, 'getItem'> | null, key: string): Partial<Record<BandLayout, number>> {
  const raw = storage?.getItem(key) ?? null;
  if (raw === null) return {};
  const parsed: unknown = JSON.parse(raw);
  return typeof parsed === 'object' && parsed !== null ? (parsed as Partial<Record<BandLayout, number>>) : {};
}

/** The member lines' last height for this layout, or `null` when this viewer was not last seen
 *  verified here (or the store is blocked or holds nonsense). */
export function readBandHold(
  storage: Pick<Storage, 'getItem'> | null,
  key: string,
  layout: BandLayout,
): number | null {
  try {
    const height = readAll(storage, key)[layout];
    return typeof height === 'number' && Number.isFinite(height) && height > 0 && height <= MAX_HOLD_PX
      ? Math.round(height)
      : null;
  } catch {
    return null;
  }
}

/** Stores the member lines' height after a verified viewer's `/home` draws them. */
export function rememberBandHold(
  storage: Pick<Storage, 'getItem' | 'setItem'> | null,
  key: string,
  layout: BandLayout,
  height: number,
): void {
  try {
    const all = readAll(storage, key);
    if (all[layout] === Math.round(height)) return;
    storage?.setItem(key, JSON.stringify({ ...all, [layout]: Math.round(height) }));
  } catch {
    // A blocked store only costs the next visit its hold.
  }
}

/** Forgets the height once the viewer is no longer a verified member here. */
export function forgetBandHold(storage: Pick<Storage, 'removeItem'> | null, key: string): void {
  try {
    storage?.removeItem(key);
  } catch {
    // Nothing held is the fallback anyway.
  }
}
