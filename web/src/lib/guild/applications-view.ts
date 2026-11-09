// web/src/lib/guild/applications-view.ts
// The officers' applications page as plain data (web layout spec 2026-11-05 §7.2 Applications,
// "Do we take them?"): the applicant's facts line, the decision record, what the bot did on
// Discord in plain words, and the recruiting card's edit (wanted lines, the role id). Pure.
import { classDisplayName } from '../sim/spec-label';
import { SPECS } from '../sim/specs';
import { outputFigure, topOutputLabel } from './output-figure';
import { applicationsCopy as copy } from './recruit-copy';
import type {
  ApplicantCharacter,
  ApplicationDecision,
  OfficerApplication,
  OfficerApplications,
  RecruitPatch,
  RecruitWanted,
  RecruitWantedInput,
} from './recruit-api';

const SEP = ' · ';

/** "Top HPS 465", the roster's own words for the figure; "91st percentile" when the log ranks it. */
export function bestParseLine(character: ApplicantCharacter): string {
  const parse = character.best_parse;
  if (parse === null) return '';
  return `${topOutputLabel(parse.metric)} ${outputFigure(parse.value)}`;
}

/** "Holy Paladin · Level 60 · Item level 63 · Top HPS 465": every figure with its unit named. */
export function applicantFacts(character: ApplicantCharacter): string {
  const who = character.spec_name ?? (character.class === null ? '' : classDisplayName(character.class));
  return [
    who,
    character.level === null ? '' : `Level ${character.level}`,
    character.item_level === null ? '' : `Item level ${character.item_level}`,
    bestParseLine(character),
  ]
    .filter((part) => part !== '')
    .join(SEP);
}

/** Pending applications, oldest first, and the decided ones, newest first, as the API orders them. */
export function splitApplications(list: OfficerApplications): {
  pending: OfficerApplication[];
  decided: OfficerApplication[];
} {
  return {
    pending: list.rows.filter((row) => row.status === 'pending'),
    decided: list.rows.filter((row) => row.status !== 'pending'),
  };
}

/** "Ashenveil accepted", "... declined", "... withdrew". */
export function decisionTitle(app: OfficerApplication): string {
  const name = app.character.name;
  if (app.status === 'accepted') return copy.accepted(name);
  if (app.status === 'declined') return copy.declined(name);
  return copy.withdrawn(name);
}

/** What the bot did, as words: the role (accepts only) and the message, each only when there is one. */
export function discordLines(
  app: Pick<OfficerApplication, 'status' | 'role_outcome' | 'dm_outcome'>,
): string[] {
  const lines: string[] = [];
  if (app.status === 'accepted' && app.role_outcome !== null) {
    lines.push(copy.role[app.role_outcome] ?? copy.role.failed);
  }
  if (app.status !== 'withdrawn' && app.dm_outcome !== null) {
    lines.push(copy.dm[app.dm_outcome] ?? copy.dm.failed);
  }
  return lines;
}

/** "3 yes, 0 no of 3 officers · by Kraggor · Discord role granted": the decision record's sub line. */
export function decisionSub(app: OfficerApplication): string {
  const parts: string[] = [];
  if (app.status !== 'withdrawn') parts.push(copy.votes(app.votes.yes, app.votes.no, app.votes.voters));
  if (app.decided_by !== null && app.status !== 'withdrawn') parts.push(copy.by(app.decided_by.name));
  return [...parts, ...discordLines(app)].join(SEP);
}

/** A decision's answer folded into its row: the outcome words the API stores on the row. */
export function appliedDecision(decision: ApplicationDecision): OfficerApplication {
  const { application, discord_role: role, discord_dm: dm } = decision;
  return {
    ...application,
    role_outcome: role.done ? 'granted' : role.reason,
    dm_outcome: dm.done ? 'sent' : dm.reason,
  };
}

/** The list with one row replaced (a vote or a decision), never mutated. */
export function withRow(list: OfficerApplications, row: OfficerApplication): OfficerApplications {
  const rows = list.rows.map((existing) => (existing.id === row.id ? row : existing));
  return { ...list, rows, pending: rows.filter((r) => r.status === 'pending').length };
}

/** The officer's edit form, one string per field, as the card stands. */
export interface CardDraft {
  open: boolean;
  pitch: string;
  nights: string;
  wanted: RecruitWantedInput[];
  roleId: string;
}

export function draftOf(
  card: { open: boolean; pitch: string; raid_nights: { text: string | null }; wanted: RecruitWanted[] },
  roleId: string | null,
): CardDraft {
  return {
    open: card.open,
    pitch: card.pitch,
    nights: card.raid_nights.text ?? '',
    wanted: card.wanted.map((line) => ({ class: line.class, spec: line.spec, note: line.note })),
    roleId: roleId ?? '',
  };
}

const ROLE_ID = /^[0-9]{17,20}$/;

/** A role id is 17 to 20 digits, or empty to clear it. */
export function roleIdError(value: string): string {
  const trimmed = value.trim();
  return trimmed === '' || ROLE_ID.test(trimmed) ? '' : copy.card.roleInvalid;
}

/** The PATCH body: the whole draft, trimmed; an empty role id clears the role. */
export function patchOf(draft: CardDraft): RecruitPatch {
  return {
    open: draft.open,
    pitch: draft.pitch.trim(),
    raid_nights: draft.nights.trim(),
    wanted: draft.wanted.map((line) => ({ class: line.class, spec: line.spec, note: line.note.trim() })),
    discord_role_id: draft.roleId.trim(),
  };
}

export const WANTED_MAX = 12;
export const PITCH_MAX = 500;
export const NIGHTS_MAX = 120;
export const WANTED_NOTE_MAX = 80;

/** The nine classes in the client's order, for the wanted line's class picker. */
export const CLASS_ORDER = [
  'warrior',
  'paladin',
  'hunter',
  'rogue',
  'priest',
  'shaman',
  'mage',
  'warlock',
  'druid',
] as const;

export function classOptions(): { value: string; label: string }[] {
  return CLASS_ORDER.map((slug) => ({ value: slug, label: classDisplayName(slug) }));
}

/** "Any spec" then the class's specs, keyed as the API wants them ("paladin-holy"). */
export function specOptions(classSlug: string): { value: string; label: string }[] {
  return [
    { value: '', label: copy.card.anySpec },
    ...SPECS.filter((spec) => spec.class_slug === classSlug).map((spec) => ({
      value: spec.spec,
      label: spec.name,
    })),
  ];
}

/** "Holy Paladin" or "Paladin": a wanted line's name for its Remove button. */
export function wantedLabel(line: RecruitWantedInput): string {
  const spec = line.spec === null ? undefined : SPECS.find((s) => s.spec === line.spec);
  const className = classDisplayName(line.class);
  return spec === undefined ? className : spec.full_name;
}
