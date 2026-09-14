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

function parsePath(prefix: string, pathname: string): CharacterPath | null {
  const parts = pathname.replace(/\/+$/, '').split('/').filter((part) => part !== '');
  if (parts.length !== 4 || parts[0] !== prefix) return null;
  const [, region, ruleset, slug] = parts;
  if (!isRegion(region) || !isRuleset(ruleset) || slug === '') return null;
  return { region, ruleset, slug };
}

export function parseCharacterPath(pathname: string): CharacterPath | null {
  return parsePath('character', pathname);
}

export function parseGuildPath(pathname: string): CharacterPath | null {
  return parsePath('guild', pathname);
}
