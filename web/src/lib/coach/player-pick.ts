// web/src/lib/coach/player-pick.ts
// Which player the Coach tab reviews (spec 3.1, "Player"). The review is for the one unit `source`
// names. With `friendlies` the viewer's first owned character is chosen, and several owned
// characters get crest buttons. A unit that is not the viewer's is reviewable only by an officer.
// With none chosen, an officer is told to open a row; anyone else sees nothing.
import { SOURCE_FRIENDLIES } from '../report/url';
import { ownedKey, type CoachViewer } from './viewer';

export interface RosterUnit {
  guid: string;
  name: string;
  class?: string;
  /** The log's spec label (`Fury`), for the panel's title before an answer names it. */
  spec?: string;
}

export interface PlayerPick {
  /** The unit the panel reviews, or null when none is chosen. */
  unit: RosterUnit | null;
  /** The viewer's characters in this fight, for the crest buttons (shown when more than one). */
  owned: readonly RosterUnit[];
  /** The key the API is asked about when the unit is the viewer's own; null for an officer's pick. */
  ownKey: string | null;
}

export function pickPlayer(viewer: CoachViewer, roster: readonly RosterUnit[], source: string): PlayerPick {
  const owned = roster.filter((unit) => ownedKey(viewer, unit.name) !== null);
  const named = roster.find((unit) => unit.guid === source);
  const unit = named !== undefined ? named : source === SOURCE_FRIENDLIES ? (owned[0] ?? null) : null;
  if (unit === null) return { unit: null, owned, ownKey: null };
  const ownKey = ownedKey(viewer, unit.name);
  if (ownKey === null && !viewer.officer) return { unit: null, owned, ownKey: null };
  return { unit, owned, ownKey };
}

/** "Open a player's row and choose Coach review" is for an officer with nobody chosen (spec 14). */
export function showsOfficerPick(viewer: CoachViewer, pick: PlayerPick): boolean {
  return pick.unit === null && viewer.signedIn && viewer.officer;
}
