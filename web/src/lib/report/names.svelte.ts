// web/src/lib/report/names.svelte.ts
// How a logged unit name reads on the report. The log writes the normalised realm after the
// name ("Elyra-ClassicBetaPvP2"); a player reads "Elyra". The realm is shown only for a
// player from another realm than the report's own, as "Elyra · Classic Beta PvP 2", with the
// full logged name on hover.
import { splitUnitName } from '../characters';

/** The realm the most players in the report carry: the report's own. Empty when none writes one. */
export function homeRealmOf(names: Iterable<string>): string {
  const counts: Record<string, number> = Object.create(null) as Record<string, number>;
  for (const logged of names) {
    const { segment } = splitUnitName(logged);
    if (segment !== '') counts[segment] = (counts[segment] ?? 0) + 1;
  }
  let best = '';
  let bestCount = 0;
  for (const [segment, count] of Object.entries(counts)) {
    if (count > bestCount) {
      best = segment;
      bestCount = count;
    }
  }
  return best;
}

/** "ClassicBetaPvP2" as "Classic Beta PvP 2"; `known` (segment to display name) wins when it has the realm. */
export function realmLabel(segment: string, known?: ReadonlyMap<string, string>): string {
  const named = known?.get(segment);
  if (named !== undefined && named !== '') return named;
  return (segment.match(/PvP|PvE|[A-Z][a-z]+|[A-Z]+(?![a-z])|[a-z]+|\d+/g) ?? [segment]).join(' ');
}

export interface PlayerLabel {
  /** The name alone, for a link or a character URL. */
  name: string;
  /** The realm to show after the name, empty for a player of the report's own realm. */
  realm: string;
  /** What the cell shows: the name, plus " · Realm Name" for a cross-realm player. */
  text: string;
  /** The hover: the full logged "Name-Realm", or the bare name when the log wrote no realm. */
  title: string;
}

export function playerLabel(logged: string, home: string, known?: ReadonlyMap<string, string>): PlayerLabel {
  const { name, segment } = splitUnitName(logged);
  const crossRealm = segment !== '' && home !== '' && segment !== home;
  const realm = crossRealm ? realmLabel(segment, known) : '';
  return { name, realm, text: crossRealm ? `${name} · ${realm}` : name, title: logged };
}

/** The report's own realm, set by the report island once it has the units. */
export const realmContext = $state({ home: '' });

export function setHomeRealm(segment: string): void {
  realmContext.home = segment;
}

/** `playerLabel` against the report's own realm. */
export function unitLabel(logged: string): PlayerLabel {
  return playerLabel(logged, realmContext.home);
}
