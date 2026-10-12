// web/src/lib/character-profile.ts
// Small reads off the character page's GET /v1/characters/... answer (rankings/api.ts's
// CharacterPage) that the profile header and the Forever Score card share.
import { guildHref, rulesetLabel } from './characters';
import type { CharacterFight, CharacterPage } from './rankings/api';

/** The spec of the character's newest fight: `history[0]`, else the first best row that has
 *  one. Null when no fight names a spec (the header then shows the class alone). */
export function newestSpec(page: Pick<CharacterPage, 'history' | 'best'>): string | null {
  const spec = page.history[0]?.spec ?? page.best.find((row) => row.spec !== undefined)?.spec;
  return spec === undefined || spec.trim() === '' ? null : spec;
}

/** "Holy Priest" when the spec is known, else "Priest"; empty with no class. */
export function specClassLabel(className: string | undefined, spec: string | null): string {
  if (className === undefined || className === '') return '';
  return spec === null ? className : `${spec} ${className}`;
}

function slug(text: string): string {
  return text.trim().toLowerCase().replace(/\s+/g, '-');
}

/** The best-in-slot page for a class and spec (`/bis/priest/holy`), or null when either is
 *  unknown. */
export function bisHref(className: string | undefined, spec: string | null): string | null {
  if (className === undefined || className === '' || spec === null) return null;
  return `/bis/${slug(className)}/${slug(spec)}`;
}

/** The two fight tables' column headers (fix round 2 item 6). */
export const characterTableCopy = {
  boss: 'Boss',
  spec: 'Spec',
  parse: 'Parse %',
  date: 'Date',
  execution: 'Execution',
} as const;

/** The amount column's header: "HPS" or "DPS" when every row is one metric, else both. */
export function amountHeader(rows: readonly Pick<CharacterFight, 'metric'>[]): string {
  const metrics = new Set(rows.map((row) => row.metric));
  if (metrics.size === 1 && metrics.has('hps')) return 'HPS';
  if (metrics.size === 1 && metrics.has('dps')) return 'DPS';
  return 'DPS or HPS';
}

/** Whether "Best per encounter" would only repeat "Every ranked fight": the same fights, by
 *  report and fight index, in both. */
export function bestRepeatsHistory(
  best: readonly Pick<CharacterFight, 'report_id' | 'fight_index'>[],
  history: readonly Pick<CharacterFight, 'report_id' | 'fight_index'>[],
): boolean {
  if (best.length !== history.length) return false;
  const key = (row: Pick<CharacterFight, 'report_id' | 'fight_index'>): string =>
    `${row.report_id}:${row.fight_index}`;
  const seen = new Set(history.map(key));
  return best.every((row) => seen.has(key(row)));
}

/** The profile header's words (Task L8.4 ruling 1). */
export const characterHeaderCopy = {
  separator: ' · ',
  /** The second line: "Level 60 Night Elf", or whichever half is known. */
  levelRace: (level: number | null, race: string | null): string | null => {
    const parts = [level === null ? '' : `Level ${level}`, race ?? ''].filter((part) => part !== '');
    return parts.length === 0 ? null : parts.join(' ');
  },
  /** The trailing chip: "Hardcore US". */
  chip: (ruleset: string, region: string): string => `${rulesetLabel(ruleset)} ${region.toUpperCase()}`,
} as const;

export interface ProfileHeaderPiece {
  key: 'spec' | 'guild' | 'realm';
  text: string;
  /** The guild's page; the other pieces are plain text. */
  href?: string;
}

export interface ProfileHeader {
  /** "Holy Priest", "<Guild>", "<Realm>": each present only when known, never a placeholder. */
  pieces: ProfileHeaderPiece[];
  /** "Level 60 Night Elf", or null when neither level nor race is known. */
  levelRace: string | null;
  /** The ruleset and region, "Hardcore US", always present (it is the page's address). */
  chip: string;
}

function known(text: string | null | undefined): string | null {
  return typeof text === 'string' && text.trim() !== '' ? text.trim() : null;
}

/**
 * The profile header (Task L8.4 ruling 1): "Holy Priest · <Guild> · <Realm>", then "Level 60
 * Night Elf" when known, with the ruleset and region as a muted trailing chip. A piece the API
 * sends as null (or omits) is left out. The guild links to its page: a verified membership
 * means the site holds the guild, and a guild lives on its members' region and ruleset.
 */
export function profileHeader(
  page: Pick<CharacterPage, 'character' | 'history' | 'best'>,
  path: { region: string; ruleset: string },
): ProfileHeader {
  const { character } = page;
  const pieces: ProfileHeaderPiece[] = [];
  const specClass = specClassLabel(character.class, newestSpec(page));
  if (specClass !== '') pieces.push({ key: 'spec', text: specClass });
  const guild = known(character.guild?.name);
  if (guild !== null)
    pieces.push({ key: 'guild', text: guild, href: guildHref(path.region, path.ruleset, guild) });
  const realm = known(character.realm);
  if (realm !== null) pieces.push({ key: 'realm', text: realm });
  const level =
    typeof character.level === 'number' && Number.isFinite(character.level) ? character.level : null;
  return {
    pieces,
    levelRace: characterHeaderCopy.levelRace(level, known(character.race)),
    chip: characterHeaderCopy.chip(path.ruleset, path.region),
  };
}
