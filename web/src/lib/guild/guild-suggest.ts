// web/src/lib/guild/guild-suggest.ts
// "Did you mean" for a guild address that 404s (guild loop round 3): "lifecycle" names no guild,
// but "Lifecycle Test" is one of the account's guilds, on the Top guilds board, or found by the
// guild search. Pure, so the matching and the order are pinned without a page.
import { characterKey, guildHref, rulesetLabel, type CharacterPath } from '../characters';

export interface GuildCandidate {
  name: string;
  region: string;
  ruleset: string;
}

export interface GuildSuggestion {
  name: string;
  href: string;
  /** `US PvP`, the way the lookup's selects write it. */
  where: string;
}

/** Fewer typed letters than this match too much to be worth a guess (or a search). */
const MIN_TYPED = 3;
/** The line names at most this many guilds. */
export const MAX_SUGGESTIONS = 3;

/** A name or slug as words: lower case, hyphens, underscores and runs of spaces as one space. */
function words(text: string): string {
  let decoded = text;
  try {
    decoded = decodeURIComponent(text);
  } catch {
    // A malformed escape is matched as typed.
  }
  return decoded
    .toLowerCase()
    .replace(/[-_\s]+/g, ' ')
    .trim();
}

/** 0 for the typed place, 1 elsewhere: a match where the player looked comes first. */
function placeRank(candidate: GuildCandidate, path: CharacterPath): number {
  return candidate.region === path.region && candidate.ruleset === path.ruleset ? 0 : 1;
}

/** 0 when the name starts with the typed words (the whole name typed on the wrong ruleset
 *  included), 1 when it contains them, null when neither. */
function matchRank(name: string, typed: string): number | null {
  if (name.startsWith(typed)) return 0;
  return name.includes(typed) ? 1 : null;
}

/** The typed words to search for, or null under three letters (no request is made then). */
export function searchTerm(path: CharacterPath): string | null {
  const typed = words(path.slug);
  return typed.length < MIN_TYPED ? null : typed;
}

function suggestionOf(candidate: GuildCandidate): GuildSuggestion {
  return {
    name: candidate.name,
    href: guildHref(candidate.region, candidate.ruleset, candidate.name),
    where: `${candidate.region.toUpperCase()} ${rulesetLabel(candidate.ruleset)}`,
  };
}

/** The guilds the page already reads that match the typed words: a name that starts with them
 *  before one that contains them, the typed region and ruleset first, then the order given (the
 *  account's guilds before the board). */
function rankedKnown(path: CharacterPath, typed: string, known: readonly GuildCandidate[]): GuildCandidate[] {
  return known
    .map((candidate, order) => ({ candidate, order, match: matchRank(words(candidate.name), typed) }))
    .filter((entry): entry is typeof entry & { match: number } => entry.match !== null)
    .sort(
      (a, b) =>
        a.match - b.match || placeRank(a.candidate, path) - placeRank(b.candidate, path) || a.order - b.order,
    )
    .map((entry) => entry.candidate);
}

/**
 * The guilds the typed address most likely meant, at most three: the matching guilds the page
 * already reads first (`known`), then the guild search's rows in its own order (`searched`: the
 * API matched them, a typo included). Never the address that just failed, never one guild twice.
 * Empty when nothing matches or the typed text is too short.
 */
export function guildSuggestions(
  path: CharacterPath,
  known: readonly GuildCandidate[],
  searched: readonly GuildCandidate[] = [],
): GuildSuggestion[] {
  const typed = searchTerm(path);
  if (typed === null) return [];
  const here = `${path.region}/${path.ruleset}/${path.slug}`;
  const keyOf = (c: GuildCandidate): string => characterKey(c.region, c.ruleset, c.name);
  const all = [...rankedKnown(path, typed, known), ...searched];
  return all
    .filter((c, index) => keyOf(c) !== here && all.findIndex((other) => keyOf(other) === keyOf(c)) === index)
    .slice(0, MAX_SUGGESTIONS)
    .map(suggestionOf);
}
