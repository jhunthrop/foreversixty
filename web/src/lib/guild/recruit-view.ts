// web/src/lib/guild/recruit-view.ts
// The public recruiting card and the application form as plain data (web layout spec 2026-11-05
// §7.2 Recruitment, "Should I apply?"): the card's lines, which of the viewer's characters can
// apply, which state the form is in, and the API's refusals in words. Pure, so a test pins them.
import type { Me } from '../account/me-types';
import { rulesetLabel } from '../characters';
import { classDisplayName } from '../sim/spec-label';
import { GuildApiError } from './api';
import { clockText, dateLabel } from './raid-time';
import { recruitCopy as copy } from './recruit-copy';
import type { GuildApplication, RecruitCard, RecruitOfficer, RecruitWanted } from './recruit-api';

const SEP = ' · ';

/** "US · PvP · Alliance · 24 members": the card's one fact line. */
export function cardFacts(card: RecruitCard): string[] {
  const { guild } = card;
  const facts = [guild.region.toUpperCase(), rulesetLabel(guild.ruleset)];
  if (guild.faction !== null) facts.push(guild.faction === 'alliance' ? 'Alliance' : 'Horde');
  facts.push(copy.card.members(card.members));
  return facts;
}

/** `Thu 10 Dec, 8 pm` in the viewer's own clock, one per planned raid. */
export function upcomingNights(card: RecruitCard, viewerTz: string, locale: string): string[] {
  return card.raid_nights.upcoming.map((at) => {
    const d = new Date(at);
    return `${dateLabel(d, viewerTz, locale)}, ${clockText(d, viewerTz, locale)}`;
  });
}

/** "1 of 1 named bosses down, Onyxia's Lair", or none before anything is down. */
export function progressLine(card: RecruitCard): string {
  const { down, named_encounters: named, name } = card.progress;
  if (named === 0) return '';
  const head = down === 0 ? 'No boss down yet' : `${down} of ${named} named bosses down`;
  return name === '' ? head : `${head}${SEP}${name}`;
}

export interface WantedRow {
  key: string;
  classSlug: string;
  /** "Holy Paladin", or "Paladin, any spec". */
  title: string;
  note: string;
}

export function wantedRows(wanted: readonly RecruitWanted[]): WantedRow[] {
  return wanted.map((line, i) => ({
    key: `${line.class}:${line.spec ?? ''}:${i}`,
    classSlug: line.class,
    title: line.spec_name ?? `${classDisplayName(line.class)}, ${copy.card.anySpec}`,
    note: line.note,
  }));
}

/** "1 tank, 2 healers" is the officers' own note; the card's summary is the classes it names. */
export function wantedSummary(wanted: readonly RecruitWanted[]): string {
  if (wanted.length === 0) return copy.card.nobodyListed;
  return wanted.map((line) => line.spec_name ?? classDisplayName(line.class)).join(', ');
}

export interface OfficerRow {
  key: string;
  name: string;
  classSlug: string;
  sub: string;
}

/** "Holy Paladin · Guild master": the spec word joins the class; the rank in words. */
export function officerRows(officers: readonly RecruitOfficer[]): OfficerRow[] {
  return officers.map((officer) => {
    const className = officer.class === null ? '' : classDisplayName(officer.class);
    const spec = [officer.spec ?? '', className].filter((part) => part !== '').join(' ');
    const rank = officer.rank === 'leader' ? copy.card.leader : copy.card.officer;
    return {
      key: officer.name,
      name: officer.name,
      classSlug: officer.class ?? '',
      sub: spec === '' ? rank : `${spec}${SEP}${rank}`,
    };
  });
}

export interface CharacterChoice {
  value: string;
  label: string;
  sub: string;
  classSlug: string;
}

/** The viewer's characters on the guild's region and ruleset: the only ones that may apply. */
export function eligibleCharacters(me: Me, card: RecruitCard): CharacterChoice[] {
  return me.characters
    .filter((c) => c.region === card.guild.region && c.ruleset === card.guild.ruleset)
    .map((c) => ({
      value: c.key,
      label: c.name,
      sub: [c.level === undefined ? '' : `Level ${c.level}`, c.class ? classDisplayName(c.class) : '']
        .filter((part) => part !== '')
        .join(' '),
      classSlug: c.class ?? '',
    }));
}

/** "US PvP": the realm a character must be on to apply. */
export const realmLabel = (card: RecruitCard): string =>
  `${card.guild.region.toUpperCase()} ${rulesetLabel(card.guild.ruleset)}`;

export type RecruitState =
  | { kind: 'signed-out' }
  | { kind: 'member' }
  | { kind: 'application'; application: GuildApplication }
  | { kind: 'closed' }
  | { kind: 'no-character' }
  | { kind: 'form'; characters: CharacterChoice[] };

/** Is the viewer a verified member of the card's guild? */
function isMember(me: Me, card: RecruitCard): boolean {
  return me.guilds.some((g) => g.id === card.guild.id && g.verified);
}

/**
 * What the page offers under the card: sign in; you are in; your open (or just decided)
 * application; the guild is closed; no character to apply with; the form.
 */
export function recruitState(
  card: RecruitCard,
  me: Me | null,
  mine: readonly GuildApplication[],
): RecruitState {
  if (me === null) return card.open ? { kind: 'signed-out' } : { kind: 'closed' };
  if (isMember(me, card)) return { kind: 'member' };
  const latest = mine[0];
  if (latest !== undefined && (latest.status === 'pending' || !card.open)) {
    return { kind: 'application', application: latest };
  }
  if (!card.open) return { kind: 'closed' };
  const characters = eligibleCharacters(me, card);
  return characters.length === 0 ? { kind: 'no-character' } : { kind: 'form', characters };
}

/** Your application's status in a sentence. */
export function applicationLine(app: GuildApplication, guildName: string): string {
  const name = app.character.name;
  if (app.status === 'pending') return copy.mine.pending(name);
  if (app.status === 'accepted') return copy.mine.accepted(name, guildName);
  if (app.status === 'declined') return copy.mine.declined(name);
  return copy.mine.withdrawn(name);
}

/** "Sent Thu 10 Dec" or "Decided Fri 11 Dec", in the viewer's clock. */
export function applicationStamp(app: GuildApplication, viewerTz: string, locale: string): string {
  if (app.decided_at !== null && app.status !== 'pending')
    return copy.mine.decided(dateLabel(new Date(app.decided_at), viewerTz, locale));
  return copy.mine.sent(dateLabel(new Date(app.created_at), viewerTz, locale));
}

/** The API's refusal of an application in words; a field error names the field's rule. */
export function applyErrorLine(error: unknown): string {
  if (!(error instanceof GuildApiError)) return copy.apply.failed;
  const known = copy.apply.errors[error.code];
  if (known !== undefined) return known;
  const field = error.fields === null ? undefined : Object.values(error.fields)[0];
  return field ?? (error.status === 0 ? copy.apply.failed : error.message);
}

export const MESSAGE_MAX = 1000;
export const AVAILABILITY_MAX = 200;
