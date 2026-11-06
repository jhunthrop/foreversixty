// web/src/lib/guild/recruit-api.ts
// Recruitment, applications, the member profile and private officer notes (web layout spec
// 2026-11-05 §7.2, api/openapi.yaml /recruit, /applications, /members/{account}/profile|notes).
// Same transport, cache and error class as api.ts; every write drops the guild's cached reads.
import { query } from '../data/query';
import { API_BASE_URL } from '../planner/config';
import {
  GUILD_TTL_MS,
  call,
  invalidateGuild,
  type GuildReadinessRow,
  type GuildRosterAttendance,
  type GuildRosterRow,
} from './api';

export type ApplicationStatus = 'pending' | 'accepted' | 'declined' | 'withdrawn';
export type ApplicationVote = 'yes' | 'no';

export interface RecruitWanted {
  /** A class slug ("paladin"). */
  class: string;
  /** A specs.json key of the class ("paladin-holy"); null for the class as a whole. */
  spec: string | null;
  /** "Holy Paladin"; null without a spec. */
  spec_name: string | null;
  note: string;
}

export interface RecruitOfficer {
  name: string;
  class: string | null;
  /** The spec word of the kill that names them ("Holy"). */
  spec: string | null;
  rank: 'leader' | 'officer';
}

export interface RecruitCard {
  guild: {
    id: number;
    name: string;
    region: string;
    ruleset: string;
    faction: 'alliance' | 'horde' | null;
    crest_url: string | null;
    path: string;
  };
  /** Verified accounts. */
  members: number;
  open: boolean;
  pitch: string;
  raid_nights: {
    /** The officers' words ("Tue and Thu, 8 to 11 pm server"); null when never set. */
    text: string | null;
    /** The planned raids of the next 14 days (at most 8), only while open. */
    upcoming: string[];
    timezone: string | null;
  };
  progress: { name: string; raids: string[]; named_encounters: number; down: number; pulls_total: number };
  wanted: RecruitWanted[];
  officers: RecruitOfficer[];
  updated_at: string | null;
}

export interface RecruitSettings {
  card: RecruitCard;
  discord_role_id: string | null;
}

export interface RecruitWantedInput {
  class: string;
  spec: string | null;
  note: string;
}

/** An officer's edit; an absent field is left as it is. `discord_role_id: ''` clears it. */
export interface RecruitPatch {
  open?: boolean;
  pitch?: string;
  raid_nights?: string;
  wanted?: RecruitWantedInput[];
  discord_role_id?: string;
}

export interface ApplicantBestParse {
  metric: 'dps' | 'hps';
  value: number;
  percentile: number | null;
  encounter: string;
  report_id: string;
  fight_index: number;
}

export interface ApplicantCharacter {
  character_key: string;
  name: string;
  class: string | null;
  spec: string | null;
  spec_name: string | null;
  level: number | null;
  item_level: number | null;
  best_parse: ApplicantBestParse | null;
}

export interface GuildApplication {
  id: number;
  guild_id: number;
  status: ApplicationStatus;
  created_at: string;
  decided_at: string | null;
  message: string;
  availability: string;
  discord_dm: boolean;
  character: ApplicantCharacter;
}

export interface ApplicationVotes {
  yes: number;
  no: number;
  /** How many verified officers (the guild master included) can vote. */
  voters: number;
  mine: ApplicationVote | null;
}

export interface GuildCharacterRef {
  name: string;
  class: string | null;
  character_key: string;
}

export interface OfficerApplication extends GuildApplication {
  votes: ApplicationVotes;
  decided_by: GuildCharacterRef | null;
  /** What the bot did on the decision: `granted`, `sent`, or a DiscordOutcome reason; null before one. */
  role_outcome: string | null;
  dm_outcome: string | null;
}

export interface OfficerApplications {
  rows: OfficerApplication[];
  pending: number;
  discord_role_id: string | null;
}

export type DiscordReason =
  'no_role_set' | 'no_bot' | 'not_linked' | 'not_opted_in' | 'not_accepted' | 'discord_refused' | 'failed';

export interface DiscordOutcome {
  done: boolean;
  reason: DiscordReason | null;
}

export interface ApplicationDecision {
  application: OfficerApplication;
  discord_role: DiscordOutcome;
  discord_dm: DiscordOutcome;
}

export interface ApplicationInput {
  character_key: string;
  message: string;
  availability: string;
  discord_dm: boolean;
}

export interface MemberNote {
  id: number;
  body: string;
  author: GuildCharacterRef | null;
  created_at: string;
  /** The viewer's own note, or any for the guild master. */
  may_delete: boolean;
}

export type ProfileOutcome = 'present' | 'late' | 'absent' | 'no_show';

export interface ProfileNight {
  raid_id: number;
  starts_at: string;
  outcome: ProfileOutcome | null;
  excused: boolean;
}

export interface ProfileAttendance {
  source: 'raids' | 'reports';
  last_4: GuildRosterAttendance;
  last_30: GuildRosterAttendance | null;
  nights: ProfileNight[];
}

export interface ProfileLootItem {
  item_id: number;
  name: string | null;
  icon: string | null;
  quality: number | null;
  raid_id: number | null;
  awarded_at: string;
}

export interface MemberProfile {
  account_key: string;
  member: GuildRosterRow | null;
  characters: { character_key: string; name: string; class: string | null; is_main: boolean }[];
  attendance: ProfileAttendance | null;
  loot: { count: number; items: ProfileLootItem[] };
  readiness: GuildReadinessRow | null;
  /** Officers only: a member's own read never carries the field. */
  notes?: MemberNote[];
}

const guildPath = (guildId: number): string => `/v1/guilds/${guildId}`;
const appPath = (guildId: number, appId: number): string => `${guildPath(guildId)}/applications/${appId}`;
const memberPath = (guildId: number, account: string): string =>
  `${guildPath(guildId)}/members/${encodeURIComponent(account)}`;

function read<T>(path: string, apiBase: string, scope: 'public' | 'private'): Promise<T> {
  return query<T>(`${apiBase}${path}`, () => call<T>(path, apiBase), {
    scope,
    ttlMs: GUILD_TTL_MS,
    revalidateOnRead: true,
  });
}

async function write<T>(
  guildId: number,
  path: string,
  apiBase: string,
  init: { method: string; body?: unknown },
): Promise<T> {
  const data = await call<T>(path, apiBase, init);
  invalidateGuild(guildId, apiBase);
  return data;
}

/** Anyone: the public recruiting card. */
export function fetchRecruitCard(guildId: number, apiBase: string = API_BASE_URL): Promise<RecruitCard> {
  return read<RecruitCard>(`${guildPath(guildId)}/recruit`, apiBase, 'public');
}

/** Officers: edit the card and the Discord role an accepted applicant gets. */
export function patchRecruit(
  guildId: number,
  patch: RecruitPatch,
  apiBase: string = API_BASE_URL,
): Promise<RecruitSettings> {
  return write<RecruitSettings>(guildId, `${guildPath(guildId)}/recruit`, apiBase, {
    method: 'PATCH',
    body: patch,
  });
}

/** Signed in: apply with one of your own characters. */
export function applyToGuild(
  guildId: number,
  input: ApplicationInput,
  apiBase: string = API_BASE_URL,
): Promise<GuildApplication> {
  return write<GuildApplication>(guildId, `${guildPath(guildId)}/applications`, apiBase, {
    method: 'POST',
    body: input,
  });
}

/** Signed in: your own applications to the guild, newest first. */
export function fetchMyApplications(
  guildId: number,
  apiBase: string = API_BASE_URL,
): Promise<{ rows: GuildApplication[] }> {
  return read<{ rows: GuildApplication[] }>(`${guildPath(guildId)}/applications/mine`, apiBase, 'private');
}

export function withdrawApplication(
  guildId: number,
  appId: number,
  apiBase: string = API_BASE_URL,
): Promise<GuildApplication> {
  return write<GuildApplication>(guildId, `${appPath(guildId, appId)}/withdraw`, apiBase, { method: 'POST' });
}

/** Officers: pending (oldest first), then the latest decisions. */
export function fetchOfficerApplications(
  guildId: number,
  apiBase: string = API_BASE_URL,
): Promise<OfficerApplications> {
  return read<OfficerApplications>(`${guildPath(guildId)}/applications`, apiBase, 'private');
}

/** Officers: a yes or a no; null takes the viewer's vote back. Answers the tally. */
export function voteApplication(
  guildId: number,
  appId: number,
  vote: ApplicationVote | null,
  apiBase: string = API_BASE_URL,
): Promise<ApplicationVotes> {
  const path = `${appPath(guildId, appId)}/vote`;
  return vote === null
    ? write<ApplicationVotes>(guildId, path, apiBase, { method: 'DELETE' })
    : write<ApplicationVotes>(guildId, path, apiBase, { method: 'PUT', body: { vote } });
}

/** Officers: accept or decline; the answer says what the bot did on Discord. */
export function decideApplication(
  guildId: number,
  appId: number,
  decision: 'accept' | 'decline',
  apiBase: string = API_BASE_URL,
): Promise<ApplicationDecision> {
  return write<ApplicationDecision>(guildId, `${appPath(guildId, appId)}/${decision}`, apiBase, {
    method: 'POST',
  });
}

/** Officer or self: `account` is the roster's account_key ("u:42") or "me". */
export function fetchMemberProfile(
  guildId: number,
  account: string,
  apiBase: string = API_BASE_URL,
): Promise<MemberProfile> {
  return read<MemberProfile>(`${memberPath(guildId, account)}/profile`, apiBase, 'private');
}

export function addMemberNote(
  guildId: number,
  account: string,
  body: string,
  apiBase: string = API_BASE_URL,
): Promise<MemberNote> {
  return write<MemberNote>(guildId, `${memberPath(guildId, account)}/notes`, apiBase, {
    method: 'POST',
    body: { body },
  });
}

export function deleteMemberNote(
  guildId: number,
  account: string,
  noteId: number,
  apiBase: string = API_BASE_URL,
): Promise<{ status: 'removed' }> {
  return write<{ status: 'removed' }>(guildId, `${memberPath(guildId, account)}/notes/${noteId}`, apiBase, {
    method: 'DELETE',
  });
}
