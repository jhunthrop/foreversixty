// web/src/lib/bis/show-event.ts
// The one event the header's character follower sends the BiS page's island: show this faction
// and band. A module of its own so the island never imports the follower's session code.
import type { Faction } from './types';

export const BIS_SHOW_EVENT = 'bis:show';

export interface BisShowDetail {
  faction: Faction;
  band: number;
}

/** The root attribute a `bis:show` leaves behind, so an island that mounts after it still shows it. */
const PENDING_KEY = 'bisShow';

/** Sends `bis:show` to the page's root, and leaves it pending for an island not yet mounted. */
export function sendBisShow(root: HTMLElement, detail: BisShowDetail): void {
  root.dataset[PENDING_KEY] = `${detail.faction}:${detail.band}`;
  root.dispatchEvent(new CustomEvent<BisShowDetail>(BIS_SHOW_EVENT, { detail }));
}

/** The show request sent before the island mounted (cleared once read), or undefined. */
export function takePendingShow(root: HTMLElement): BisShowDetail | undefined {
  const raw = root.dataset[PENDING_KEY];
  delete root.dataset[PENDING_KEY];
  const match = /^(alliance|horde):(\d+)$/.exec(raw ?? '');
  if (match === null) return undefined;
  return { faction: match[1] as Faction, band: Number(match[2]) };
}
