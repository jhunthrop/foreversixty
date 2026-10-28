// web/src/lib/coach/viewer.ts
// Who is looking at the report, as far as the coach cares: whether they are signed in, whether
// they hold `log_coach` (Premium) or officer views, and which logged unit is one of their own
// characters. Pure over the `/v1/me` answer, so the row strip, the tab gate and the Coach tab
// agree (Ruling 6: the row line fetches only for the viewer's own rows).
import { splitUnitName } from '../characters';
import type { Me } from '../account/me-types';

export interface CoachViewer {
  signedIn: boolean;
  premium: boolean;
  officer: boolean;
  /** Character key by lower-cased display name. */
  ownedKeys: ReadonlyMap<string, string>;
}

export const SIGNED_OUT_VIEWER: CoachViewer = {
  signedIn: false,
  premium: false,
  officer: false,
  ownedKeys: new Map(),
};

const nameKey = (name: string): string => splitUnitName(name).name.toLowerCase();

export function viewerFromMe(me: Me | null): CoachViewer {
  if (me === null) return SIGNED_OUT_VIEWER;
  return {
    signedIn: true,
    premium: me.entitlements?.log_coach === true,
    officer: me.entitlements?.officer_views === true,
    ownedKeys: new Map(me.characters.map((character) => [nameKey(character.name), character.key])),
  };
}

/** The viewer's character key for a logged unit name, or null when the unit is not theirs. */
export function ownedKey(viewer: CoachViewer, unitName: string): string | null {
  return viewer.ownedKeys.get(nameKey(unitName)) ?? null;
}

/** ModeBar lists the Coach tab for an owner of a character in the report, or for an officer. */
export function coachTabListed(viewer: CoachViewer, roster: readonly { name: string }[]): boolean {
  if (!viewer.signedIn) return false;
  return viewer.officer || roster.some((unit) => ownedKey(viewer, unit.name) !== null);
}
