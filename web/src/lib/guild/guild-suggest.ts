// web/src/lib/guild/guild-suggest.ts
// "Did you mean" for a guild address that 404s (guild loop round 3): "lifecycle" names no guild,
// but "Lifecycle Test" is one of the account's guilds or on the Top guilds board. No search
// endpoint: the candidates are what the page can already read. Pure, so the matching is pinned
// without a page.
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

/** Fewer typed letters than this match too much to be worth a guess. */
const MIN_TYPED = 3;

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

/**
 * The one guild the typed address most likely meant: a name that starts with the typed words
 * before one that contains them, the typed region and ruleset first, then the order given (the
 * account's guilds before the board). Null when nothing matches or the typed text is too short.
 */
export function guildSuggestion(
  path: CharacterPath,
  candidates: readonly GuildCandidate[],
): GuildSuggestion | null {
  const typed = words(path.slug);
  if (typed.length < MIN_TYPED) return null;
  const here = `${path.region}/${path.ruleset}/${path.slug}`;
  const ranked = candidates
    .filter((candidate) => characterKey(candidate.region, candidate.ruleset, candidate.name) !== here)
    .map((candidate, order) => ({ candidate, order, match: matchRank(words(candidate.name), typed) }))
    .filter((entry): entry is typeof entry & { match: number } => entry.match !== null)
    .sort(
      (a, b) =>
        a.match - b.match || placeRank(a.candidate, path) - placeRank(b.candidate, path) || a.order - b.order,
    );
  const best = ranked[0]?.candidate;
  if (best === undefined) return null;
  return {
    name: best.name,
    href: guildHref(best.region, best.ruleset, best.name),
    where: `${best.region.toUpperCase()} ${rulesetLabel(best.ruleset)}`,
  };
}
