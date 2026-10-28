// web/src/lib/coach/viewer.ts
// Who is looking at the report, as far as the coach cares: whether they are signed in, whether
// they hold `log_coach` (Premium) or officer views, and which logged unit is one of their own
// characters. Pure over the `/v1/me` answer, so the row strip, the tab gate and the Coach tab
// agree (Ruling 6: the row line fetches only for the viewer's own rows).
import { rulesetFromRealm, splitUnitName } from '../characters';

/** The log's realm segment ends in the ruleset word: `ClassicBetaPvP`, `ClassicBetaPvE2`. */
const SEGMENT_RULESETS: readonly (readonly [RegExp, string])[] = [
  [/pvp\d*$/i, 'pvp'],
  [/pve\d*$/i, 'normal'],
  [/(roleplay|rp)\d*$/i, 'rp'],
  [/(hardcore|hc)\d*$/i, 'hardcore'],
];

function segmentRuleset(segment: string): string {
  const hit = SEGMENT_RULESETS.find(([pattern]) => pattern.test(segment));
  return hit?.[1] ?? rulesetFromRealm(segment);
}
import type { Me } from '../account/me-types';

export interface OwnedCharacter {
  key: string;
  ruleset: string;
}

export interface CoachViewer {
  signedIn: boolean;
  premium: boolean;
  officer: boolean;
  /** The viewer's characters by lower-cased display name; two realms can share a name. */
  ownedKeys: ReadonlyMap<string, readonly OwnedCharacter[]>;
}

export const SIGNED_OUT_VIEWER: CoachViewer = {
  signedIn: false,
  premium: false,
  officer: false,
  ownedKeys: new Map(),
};

const nameKey = (name: string): string => splitUnitName(name).name.toLowerCase();

function groupByName(characters: Me['characters']): Map<string, OwnedCharacter[]> {
  const grouped = new Map<string, OwnedCharacter[]>();
  for (const character of characters) {
    const key = nameKey(character.name);
    grouped.set(key, [...(grouped.get(key) ?? []), { key: character.key, ruleset: character.ruleset }]);
  }
  return grouped;
}

export function viewerFromMe(me: Me | null): CoachViewer {
  if (me === null) return SIGNED_OUT_VIEWER;
  return {
    signedIn: true,
    premium: me.entitlements?.log_coach === true,
    officer: me.entitlements?.officer_views === true,
    ownedKeys: groupByName(me.characters),
  };
}

/**
 * The viewer's character key for a logged unit, or null when the unit is not theirs. The unit's
 * realm segment names a ruleset ("ClassicBetaPvP" is pvp): a same-named stranger on another
 * ruleset is not the viewer, and two own characters of one name resolve by ruleset. An ambiguous
 * match is null, so no strip asks the API about the wrong key.
 */
export function ownedKey(viewer: CoachViewer, unitName: string): string | null {
  const candidates = viewer.ownedKeys.get(nameKey(unitName)) ?? [];
  const ruleset = segmentRuleset(splitUnitName(unitName).segment);
  const matching = ruleset === '' ? candidates : candidates.filter((c) => c.ruleset === ruleset);
  return matching.length === 1 ? (matching[0]?.key ?? null) : null;
}

/** ModeBar lists the Coach tab for an owner of a character in the report, or for an officer. */
export function coachTabListed(viewer: CoachViewer, roster: readonly { name: string }[]): boolean {
  if (!viewer.signedIn) return false;
  return viewer.officer || roster.some((unit) => ownedKey(viewer, unit.name) !== null);
}
