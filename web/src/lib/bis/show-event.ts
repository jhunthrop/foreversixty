// web/src/lib/bis/show-event.ts
// The one event the header's character follower sends the BiS page's island: show this faction
// and band. A module of its own so the island never imports the follower's session code.
import type { Faction } from './types';

export const BIS_SHOW_EVENT = 'bis:show';

export interface BisShowDetail {
  faction: Faction;
  band: number;
}
