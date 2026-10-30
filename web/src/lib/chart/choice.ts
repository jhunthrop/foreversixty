// web/src/lib/chart/choice.ts
// The Players control's choice (controls and charts spec §B5.3, §B6.4): null keeps the default
// top 8; a set is exactly who is drawn beside the always-drawn players. Pure, so the control, the
// legend sheet and the url (lib/report/url.ts, `players=`) all agree.
import type { PlayerEntry } from './types';

/** Toggles one player, starting from who is drawn now when the default still holds. */
export function togglePlayer(
  chosen: ReadonlySet<string> | null,
  entries: readonly PlayerEntry[],
  id: string,
): Set<string> {
  const base =
    chosen ?? new Set(entries.filter((entry) => entry.drawn && !entry.pinned).map((entry) => entry.id));
  const next = new Set(base);
  if (next.has(id)) next.delete(id);
  else next.add(id);
  return next;
}

/** Case- and accent-insensitive substring match on a label. */
export function matchesFilter(label: string, query: string): boolean {
  const fold = (text: string): string => text.normalize('NFD').replace(/\p{M}/gu, '').toLowerCase();
  return fold(label).includes(fold(query.trim()));
}
