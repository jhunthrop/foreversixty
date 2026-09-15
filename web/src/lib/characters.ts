// web/src/lib/characters.ts
// Forever has no realms. The Deep Dive panel replaced them with four rulesets per region,
// so a character is keyed <region>/<ruleset>/<name-slug> and names are two parts, first and
// last. This module owns every place the site turns one of those into a URL, and the one
// place it turns a logged unit name back into a displayable name.
//
// The combat log still writes units as `Name-Realm`. What that second segment carries for
// Forever is settled by the Sept 17 beta log, so nothing here treats it as a ruleset: it is
// split off for display and otherwise shown as the log wrote it.

export const RULESETS = [
  { id: 'normal', label: 'Normal' },
  { id: 'pvp', label: 'PvP' },
  { id: 'rp', label: 'Roleplay' },
  { id: 'hardcore', label: 'Hardcore' },
] as const;

export type Ruleset = (typeof RULESETS)[number]['id'];

export const REGIONS = ['us', 'eu', 'kr', 'tw', 'cn'] as const;
export type Region = (typeof REGIONS)[number];

export function isRuleset(value: string): value is Ruleset {
  return RULESETS.some((ruleset) => ruleset.id === value);
}

export function isRegion(value: string): value is Region {
  return (REGIONS as readonly string[]).includes(value);
}

/** A ruleset the site has not heard of is echoed rather than hidden: the API is the truth. */
export function rulesetLabel(id: string): string {
  return RULESETS.find((ruleset) => ruleset.id === id)?.label ?? id;
}

export interface UnitName {
  /** What the report shows. */
  name: string;
  /** The log's trailing segment, empty when the log did not write one. */
  segment: string;
}

/**
 * Splits on the last hyphen, which is the game's own rule: a hyphenated first name is
 * common and a hyphenated realm segment is not.
 */
export function splitUnitName(logged: string): UnitName {
  const cut = logged.lastIndexOf('-');
  if (cut <= 0) return { name: logged, segment: '' };
  return { name: logged.slice(0, cut), segment: logged.slice(cut + 1) };
}

/** `Elyra Duskvale` becomes `elyra-duskvale`; runs of whitespace collapse to one hyphen. */
export function characterSlug(name: string): string {
  return name.trim().toLowerCase().replace(/\s+/g, '-');
}

export function characterKey(region: string, ruleset: string, name: string): string {
  return `${region}/${ruleset}/${characterSlug(name)}`;
}

export function characterHref(region: string, ruleset: string, name: string): string {
  return `/character/${characterKey(region, ruleset, name)}`;
}

export function guildHref(region: string, ruleset: string, name: string): string {
  return `/guild/${characterKey(region, ruleset, name)}`;
}

export interface CharacterPath {
  region: Region;
  ruleset: Ruleset;
  slug: string;
}

/** Longer than any name the game allows, and a bound on what is pasted into an API path. */
const MAX_SLUG_LENGTH = 64;

/** The same bound before decoding, where one character can be several `%xx` escapes. */
const MAX_ENCODED_SLUG_LENGTH = MAX_SLUG_LENGTH * 4;

/** A path separator, a dot segment, a query or fragment marker, a stray escape, a space. */
const UNSAFE_IN_SLUG = /[/\\.?#%\s]/;

/**
 * Whether a slug names a character rather than somewhere else.
 *
 * The slug is interpolated straight into `/v1/characters/<region>/<ruleset>/<slug>` and
 * into the Worker's own upstream url, so a dot segment there does not 404 -- the URL
 * parser resolves it, and `/character/us/normal/..` becomes a request to a different
 * endpoint entirely. Anything that can traverse or re-target is refused: a path
 * separator, a dot, a query or fragment marker, whitespace, a control character.
 *
 * Deliberately not `[a-z0-9-]`, which is what /rankings/<slug> allows: a pathname reaches
 * here percent-encoded, and kr, tw and cn are regions this site serves, so real character
 * pages carry `%`-escapes for names that are not Latin at all. The escapes are decoded
 * before the check instead, so `%2e%2e` and `%2f` are refused as the `..` and `/` they
 * are -- the WHATWG URL parser resolves those forms too -- while a Hangul name is not.
 */
function isCharacterSlug(slug: string): boolean {
  if (slug === '' || slug.length > MAX_ENCODED_SLUG_LENGTH) return false;
  let decoded: string;
  try {
    decoded = decodeURIComponent(slug);
  } catch {
    // A malformed escape. Nothing legitimate produces one, and it is not worth guessing at.
    return false;
  }
  if (decoded.length > MAX_SLUG_LENGTH || UNSAFE_IN_SLUG.test(decoded)) return false;
  // Control characters, separately: writing them as a range in the expression above is
  // what no-control-regex exists to stop, and this reads better than escaping past it.
  for (const character of decoded) if ((character.codePointAt(0) ?? 0) < 0x20) return false;
  return true;
}

function parsePath(prefix: string, pathname: string): CharacterPath | null {
  const parts = pathname
    .replace(/\/+$/, '')
    .split('/')
    .filter((part) => part !== '');
  if (parts.length !== 4 || parts[0] !== prefix) return null;
  const [, region, ruleset, slug] = parts;
  if (!isRegion(region) || !isRuleset(ruleset) || !isCharacterSlug(slug)) return null;
  return { region, ruleset, slug };
}

export function parseCharacterPath(pathname: string): CharacterPath | null {
  return parsePath('character', pathname);
}

export function parseGuildPath(pathname: string): CharacterPath | null {
  return parsePath('guild', pathname);
}

/**
 * A ranking row's `player.key` (contract: `GET /v1/rankings` and `GET /v1/rankings/guilds`)
 * is `<region>/<ruleset>/<name-slug>` -- exactly `parseCharacterPath`'s last three segments
 * -- so this reuses that parser's validation rather than trusting the key's shape and
 * guessing at a fallback when a row happens to carry no guild. A row with no guild is not a
 * row with no region or ruleset: `player.key` always has both, and a caller that defaulted
 * them (`us`/`normal`, say) would link a real player at a link for someone else's server. A
 * key this does not recognise -- a future format change, a row the API answered oddly --
 * returns null rather than a guess a link would make look correct.
 */
export function parseCharacterKey(key: string): { region: Region; ruleset: Ruleset } | null {
  const path = parseCharacterPath(`/character/${key}`);
  return path === null ? null : { region: path.region, ruleset: path.ruleset };
}
