// web/src/lib/character-profile.ts
// Small reads off the character page's GET /v1/characters/... answer (rankings/api.ts's
// CharacterPage) that the profile header and the Forever Score card share.
import { bandForLevel } from './bis/hover';
import { specLinkHref } from './bis/landing-links';
import type { MeCharacter } from './account/api';
import { characterSpecName, specFromPoints } from './character-spec';
import { guildHref, rulesetLabel } from './characters';
import { specName } from './guild/spec-name';
import { factionOfRace } from './race-faction';
import type { BuildSeen, CharacterFight, CharacterPage } from './rankings/api';
import { scoreCopy } from './score/copy';
import { shortDate } from './score/score';

/** The spec of the character's newest fight: `history[0]`, else the first best row that has
 *  one. Null when no fight names a spec (the header then shows the class alone). */
export function newestSpec(page: Pick<CharacterPage, 'history' | 'best'>): string | null {
  const spec = page.history[0]?.spec ?? page.best.find((row) => row.spec !== undefined)?.spec;
  return spec === undefined || spec.trim() === '' ? null : spec;
}

/** "hunter" -> "Hunter", "death knight" -> "Death Knight": the API sends the class as it keys it. */
export function titleCase(text: string): string {
  return text.replace(/\b\p{L}/gu, (letter) => letter.toUpperCase());
}

/** "Holy Priest" when the spec is known, else "Priest"; empty with no class. */
export function specClassLabel(className: string | undefined, spec: string | null): string {
  if (className === undefined || className === '') return '';
  const label = titleCase(className);
  return spec === null ? label : `${spec} ${label}`;
}

function slug(text: string): string {
  return text.trim().toLowerCase().replace(/\s+/g, '-');
}

/**
 * The Gear section's footer link, "Every pick for Holy at 60" (L8.6 ruling 2; "Most to gain:
 * Gear" lands on that section now, and comes here only for an API that sends no `gear`; L8.4 fix
 * round 4 item 2 first pointed it here):
 * the spec's best-in-slot page at this character's band, and its faction when the race gives
 * one -- `/bis/priest/holy?faction=alliance#band-alliance-60`, the route and anchor the /bis
 * landing's spec links open (bis/landing-links.ts). With no level, the spec's page as it opens
 * (`/bis/priest/holy`). Null when the class or spec is unknown.
 */
export function bisHref(
  className: string | undefined,
  spec: string | null,
  level: number | null | undefined,
  race: string | null | undefined,
): string | null {
  if (className === undefined || className === '' || spec === null) return null;
  const classSlug = slug(className);
  const specSlug = slug(spec);
  if (typeof level !== 'number' || !Number.isFinite(level)) return `/bis/${classSlug}/${specSlug}`;
  const faction = factionOfRace(race) ?? undefined;
  return specLinkHref({ classSlug, specSlug, band: bandForLevel(level), faction });
}

/** The two fight tables' column headers (fix round 2 item 6). */
export const characterTableCopy = {
  boss: 'Boss',
  spec: 'Spec',
  parse: 'Parse %',
  date: 'Date',
  execution: 'Execution',
  /** What the Execution column is (fix round 3 item 7): a muted 12px caption under each
   *  table's title, shown at every width (fix round 4 items 6 and 7). Damage only (fix round 5
   *  item 1), as api/openapi.yaml's `ExecutionScore` says since L6.7: "the fight's damage as a
   *  share of what the same gear and talents do in the simulator", null for healers and tanks
   *  (the simulator scores the dps role alone). */
  executionDefinition:
    'Execution: the fight’s damage as a share of what the same gear and talents do in the simulator; healers and tanks have none.',
  /** The "—" cell's title and aria-label (fix round 5 item 7). */
  executionUnscored: 'Not scored yet',
  /** Under the latest rated fight's parse (L8.5 fix round 2 item 3): "rated 65". */
  ratedNote: (overall: number): string => `rated ${Math.round(overall)}`,
} as const;

/**
 * One Execution cell style for both tables and both states (fix round 3 items 4 and 6): a
 * scored cell (a link) and an unscored one (a dash) take the same vertically centred,
 * right-aligned flex box, so both sit on their row's centre line. Both tables keep the column
 * on a phone (fix round 4 items 3 and 7): the box is at least 44px high and wide, and 4px of
 * margin on top of the 12px column gap puts it 16px after the amount there.
 */
export function executionCellClass(scored: boolean): string {
  return `flex min-h-11 min-w-11 max-md:ml-1 items-center justify-end text-right numeric text-[13px]${scored ? '' : ' text-muted'}`;
}

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

type FightRef = Pick<CharacterFight, 'report_id' | 'fight_index'>;

/** A row of "Every ranked fight" by its report and fight (L8.5 ruling 1): "fight-fixture2abcd-3". */
export function fightRowId(row: FightRef): string {
  return `fight-${row.report_id}-${row.fight_index}`;
}

/**
 * Where the rating card's "Weakest part of a fight" goes (L8.5 ruling 1): its fight's row in
 * "Every ranked fight", or null when that fight is not a row there (or there is no trend).
 *
 * The rating API names the weakest part from its newest card alone (api/internal/rating/
 * dto.go's bestWorstComponent), and neither the trend's points nor the history rows carry a
 * per-fight part score, so the fight the part was read from is the newest rated one.
 */
export function weakestFightHref(
  trend: readonly (FightRef & { fought_at: string })[],
  history: readonly Pick<CharacterFight, 'report_id' | 'fight_index' | 'encounter'>[],
): string | null {
  const latest = latestRatedFight(trend, history);
  return latest === null || latest.rowId === null ? null : `#${latest.rowId}`;
}

/** The latest rated fight (L8.5 fix round 2 items 1 and 3). */
export interface LatestRatedFight {
  report_id: string;
  fight_index: number;
  /** Its encounter, from its row in "Every ranked fight"; null when it is not a row there. */
  encounter: string | null;
  /** "9 Dec", the site's day-month form, in UTC; null for an unreadable time. */
  date: string | null;
  /** Its row's id in "Every ranked fight", or null. */
  rowId: string | null;
}

/**
 * The latest rated fight: the trend's newest point, which is the card the rating endpoint's
 * `latest` is built from (dto.go's buildCharacterRatingDTO); `latest` itself carries no
 * encounter, date, report or fight index, so the trend point names the fight and "Every ranked
 * fight" its encounter. Null with no trend.
 */
export function latestRatedFight(
  trend: readonly (FightRef & { fought_at: string })[],
  history: readonly Pick<CharacterFight, 'report_id' | 'fight_index' | 'encounter'>[],
): LatestRatedFight | null {
  const newest = [...trend].sort(
    (a, b) => new Date(b.fought_at).getTime() - new Date(a.fought_at).getTime(),
  )[0];
  if (newest === undefined) return null;
  const row = history.find(
    (fight) => fight.report_id === newest.report_id && fight.fight_index === newest.fight_index,
  );
  const time = new Date(newest.fought_at);
  return {
    report_id: newest.report_id,
    fight_index: newest.fight_index,
    encounter: row?.encounter ?? null,
    date: Number.isNaN(time.getTime()) ? null : shortDate(time.toISOString().slice(0, 10)),
    rowId: row === undefined ? null : fightRowId(row),
  };
}

/** The note under "Builds seen at pull" (L8.4 fix round 4 item 4), in plain words, naming
 *  every distinct split seen, newest first (fix round 5 item 4): "a 0/31/20 split" for one,
 *  "0/31/20 and 20/31/0 splits" for more. Null with no split to name. */
export const characterBuildsCopy = {
  footnote: (splits: readonly string[]): string | null => {
    if (splits.length === 0) return null;
    const named =
      splits.length === 1
        ? `a ${splits[0]} split`
        : `${splits.slice(0, -1).join(', ')} and ${splits[splits.length - 1]} splits`;
    return `Seen in logs as ${named}. Logs don’t record the order points were spent, so there is no saved build to open.`;
  },
} as const;

/** The distinct talent splits of the builds seen at pull, newest first, blank ones left out
 *  (fix round 5 item 4). */
export function seenSplits(builds: readonly BuildSeen[] | null | undefined): string[] {
  const newestFirst = [...(builds ?? [])].sort((a, b) => b.first_seen.localeCompare(a.first_seen));
  return [...new Set(newestFirst.map((build) => build.talent_split.trim()).filter((split) => split !== ''))];
}

/** The page's one line with JavaScript off (fix round 3 item 9): the island fetches the
 *  character in the browser, so without JavaScript there is nothing to show; the Planner works
 *  without this page. Rendered "{lead} {link}{tail}" with the link to `href`. */
export const characterNoScriptCopy = {
  lead: 'Forever Sixty needs JavaScript to show this character. Turn it on and reload, or open the',
  link: 'Planner',
  href: '/planner',
  tail: '.',
} as const;

/** The profile header's words (Task L8.4 ruling 1). */
export const characterHeaderCopy = {
  separator: ' · ',
  /** The end of every segment of the first line but the last (fix round 5 item 6, replacing
   *  fix round 3 item 10's leading dot): a no-break space, then the dot, inside the segment's
   *  nowrap unit, so the dot stays with the segment before it and a wrap never starts a line
   *  with one: "Holy Priest · 0/31/20 · <Olympus XXVII> ·" then "Classic Beta PvP". */
  segmentTail: '\u00a0\u00b7',
  /** Between two segments: a plain space outside the units, the only place line 1 breaks. */
  segmentBreak: ' ',
  /** Between the last word and the chip ending the first line (fix round 3 item 2): a word
   *  joiner, zero width, inside a nowrap span, so the chip never wraps away from that word while
   *  the text before it (a long guild) still wraps. Not words: layout only. */
  chipGlue: '\u2060',
  /** The second line: "Level 60 Night Elf", or whichever half is known. */
  levelRace: (level: number | null, race: string | null): string | null => {
    const parts = [level === null ? '' : `Level ${level}`, race ?? ''].filter((part) => part !== '');
    return parts.length === 0 ? null : parts.join(' ');
  },
  /** The realm field as the character carries it, and nothing else (P.1): the badge prints the ruleset, so
   *  the line never has to explain itself. */
  realm: (realm: string): string => realm,
  /** The trailing chip: "Hardcore US". */
  chip: (ruleset: string, region: string): string => `${rulesetLabel(ruleset)} ${region.toUpperCase()}`,
  /** The guild the way the client prints it, "<Olympus XXVII>" (fix round 2 item 4): the same
   *  words as the tooltip mock's guild line. */
  guild: scoreCopy.tooltipGuild,
} as const;

export interface ProfileHeaderPiece {
  key: 'spec' | 'guild' | 'realm';
  text: string;
  /** The guild's page; the other pieces are plain text. */
  href?: string;
  /** The spec piece only: the newest build's talent split, "31/20/0" (fix round 2 item 4). */
  split?: string;
}

export interface ProfileHeader {
  /** "Holy Priest", "<Olympus XXVII>", "Classic Beta PvP": each present only when known, never
   *  a placeholder. */
  pieces: ProfileHeaderPiece[];
  /** "Level 60 Night Elf", or null when neither level nor race is known. */
  levelRace: string | null;
  /** The ruleset and region, "Hardcore US", always present (it is the page's address). */
  chip: string;
  /** The class as the API keys it ("hunter"), for the descriptor's class colour; null when unknown. */
  className: string | null;
  /** The line the chip ends (fix round 2 item 3): the last text line, so the level line when
   *  there is one, else the first. */
  chipOn: 'levelRace' | 'pieces';
}

/** The newest build seen at pull (the latest `first_seen`), or null with none. */
export function newestBuild(builds: readonly BuildSeen[] | null | undefined): BuildSeen | null {
  let newest: BuildSeen | null = null;
  for (const build of builds ?? []) {
    if (newest === null || build.first_seen > newest.first_seen) newest = build;
  }
  return newest;
}

/** The talent split of the newest build seen at pull, or null. */
export function newestSplit(builds: readonly BuildSeen[] | null | undefined): string | null {
  const split = newestBuild(builds)?.talent_split.trim() ?? '';
  return split === '' ? null : split;
}

/**
 * The spec the profile names (fix round 3 item 8): the newest build's own spec when it carries
 * one, so the split and the spec word beside it come from one record; else the character's
 * spec from its fights (`newestSpec`). A Priest's 31/20/0 is Discipline whatever the newest
 * fight row says.
 */
export function profileSpec(
  page: Pick<CharacterPage, 'history' | 'best'> & Partial<Pick<CharacterPage, 'builds_seen'>>,
): string | null {
  return known(newestBuild(page.builds_seen)?.spec) ?? newestSpec(page);
}

/** The record the viewer owns, when it is theirs: the pieces the one spec rule reads. */
type OwnRecord = Pick<MeCharacter, 'class' | 'spec' | 'build'> | null | undefined;

/**
 * The page's spec by the one rule (lib/character-spec.ts): the owner's synced talent ranks when
 * the viewer owns the character, else the newest build seen at pull when it carries a split, else
 * the spec the fights or the addon named -- always as a player says it ("Beast Mastery").
 */
export function resolvedProfileSpec(
  page: Pick<CharacterPage, 'character' | 'history' | 'best'> & Partial<Pick<CharacterPage, 'builds_seen'>>,
  own?: OwnRecord,
): string | null {
  if (own !== null && own !== undefined) {
    const fromOwn = characterSpecName({ ...own, class: own.class ?? page.character.class, spec: undefined });
    if (own.build?.talents !== undefined && fromOwn !== undefined) return fromOwn;
  }
  const split = newestSplit(page.builds_seen);
  if (split !== null && page.character.class !== undefined) {
    const derived = specFromPoints(
      page.character.class,
      split.split('/').map((count) => Number.parseInt(count, 10) || 0),
    );
    if (derived !== null) return derived;
  }
  const named = profileSpec(page);
  return named === null ? null : specName(named);
}

/** A build seen at pull as a player says it: the spec its split leans to, else the spec it was filed under. */
export function seenBuildSpec(
  className: string | undefined,
  build: Pick<BuildSeen, 'talent_split' | 'spec'>,
): string {
  const derived =
    className === undefined || className === ''
      ? null
      : specFromPoints(
          className,
          build.talent_split.split('/').map((count) => Number.parseInt(count, 10) || 0),
        );
  return derived ?? specName(build.spec ?? '');
}

/** The talent split beside the spec: the owner's synced ranks, else the newest build seen. */
export function resolvedProfileSplit(
  page: Partial<Pick<CharacterPage, 'builds_seen'>>,
  own?: OwnRecord,
): string | null {
  const points = own?.build?.talents?.points;
  if (points !== undefined && points.some((count) => count > 0)) return points.join('/');
  return newestSplit(page.builds_seen);
}

function known(text: string | null | undefined): string | null {
  return typeof text === 'string' && text.trim() !== '' ? text.trim() : null;
}

/**
 * The profile header (Task L8.4 ruling 1, fix round 2 items 3 and 4): "Holy Priest · 0/31/20 ·
 * <Olympus XXVII> · Classic Beta PvP", then "Level 60 Night Elf" when known, with the ruleset
 * and region as a muted chip ending the last line. A piece the API sends as null (or omits) is
 * left out. The guild links to its page: a verified membership means the site holds the guild,
 * and a guild lives on its members' region and ruleset.
 */
export function profileHeader(
  page: Pick<CharacterPage, 'character' | 'history' | 'best'> & Partial<Pick<CharacterPage, 'builds_seen'>>,
  path: { region: string; ruleset: string },
  own?: OwnRecord,
): ProfileHeader {
  const { character } = page;
  const pieces: ProfileHeaderPiece[] = [];
  const specClass = specClassLabel(character.class, resolvedProfileSpec(page, own));
  const split = resolvedProfileSplit(page, own);
  if (specClass !== '')
    pieces.push(split === null ? { key: 'spec', text: specClass } : { key: 'spec', text: specClass, split });
  const guild = known(character.guild?.name);
  if (guild !== null)
    pieces.push({
      key: 'guild',
      text: characterHeaderCopy.guild(guild),
      href: guildHref(path.region, path.ruleset, guild),
    });
  const realm = known(character.realm);
  if (realm !== null) pieces.push({ key: 'realm', text: characterHeaderCopy.realm(realm) });
  const level =
    typeof character.level === 'number' && Number.isFinite(character.level) ? character.level : null;
  const levelRace = characterHeaderCopy.levelRace(level, known(character.race));
  return {
    pieces,
    levelRace,
    chip: characterHeaderCopy.chip(path.ruleset, path.region),
    className: known(character.class),
    chipOn: levelRace === null ? 'pieces' : 'levelRace',
  };
}
