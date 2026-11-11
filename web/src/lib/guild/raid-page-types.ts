// web/src/lib/guild/raid-page-types.ts
// The raid page controller's contract: its route, the deps every network call and clock come in
// through, and the state vocabulary the components read (who may see the raid, its load, where an
// officer failure line sits, and the fields a component test seeds).
import type { Me } from '../account/api';
import type { CharacterPath } from '../characters';
import type { GuildPage } from '../rankings/api';
import type { GuildHome, GuildRaidView, GuildZoneView } from './api';
import type { AttendanceOverrideStatus, MatchCandidates, PostReportResult } from './attendance-contract';
import type { AssignmentPostResult } from './assignments-api';
import type { GuildCompRole, GuildPingResult } from './overview-api';
import type { RaidInput, RaidStatusAction, RaidStatusBody, SignupInput, SignupResult } from './raids-api';

export interface RaidRoute {
  guild: CharacterPath;
  raidId: string;
}

export interface RaidPageDeps {
  loadGuild: () => Promise<GuildPage>;
  loadMe: () => Promise<Me | null>;
  loadHome: (guildId: number) => Promise<GuildHome>;
  fetchRaid: (guildId: number, raidId: string) => Promise<GuildRaidView>;
  /** Asks the network now, whatever the cache holds; rejects on a failure. */
  refetchRaid: (guildId: number, raidId: string) => Promise<GuildRaidView>;
  /** The one sign-up save (lib/guild/signup-save.ts sendSignup). */
  saveSignup: (guildId: number, raidId: string, key: string, body: SignupInput) => Promise<SignupResult>;
  setStatus: (
    guildId: number,
    raidId: string,
    action: RaidStatusAction,
    body?: RaidStatusBody,
  ) => Promise<GuildRaidView>;
  updateRaid: (
    guildId: number,
    raidId: string,
    body: Omit<RaidInput, 'repeat_weeks'>,
  ) => Promise<GuildRaidView>;
  /** The match picker's read (officers; never cached). */
  loadMatchCandidates: (guildId: number, raidId: string) => Promise<MatchCandidates>;
  putMatch: (guildId: number, raidId: string, reportId: string) => Promise<GuildRaidView>;
  deleteMatch: (guildId: number, raidId: string) => Promise<GuildRaidView>;
  /** An officer's attendance fix over the log; answers the raid view. */
  putOverride: (
    guildId: number,
    raidId: string,
    characterKey: string,
    status: AttendanceOverrideStatus,
  ) => Promise<GuildRaidView>;
  /** Back to what the log said; answers the raid view. */
  deleteOverride: (guildId: number, raidId: string, characterKey: string) => Promise<GuildRaidView>;
  /** Post to Discord (again, with `repost`): 409 already_posted, 502 post_failed. */
  postReport: (guildId: number, raidId: string, repost: boolean) => Promise<PostReportResult>;
  /** The planner's comp target (officers); each answers the raid view. */
  putCompTarget: (
    guildId: number,
    raidId: string,
    target: Readonly<Record<GuildCompRole, number>>,
  ) => Promise<GuildRaidView>;
  deleteCompTarget: (guildId: number, raidId: string) => Promise<GuildRaidView>;
  /** Ping on Discord: one message under the raid's post naming the non-signers. */
  pingRaid: (guildId: number, raidId: string) => Promise<GuildPingResult>;
  /** Post now (unposted) or Repost (posted): answers the raid view. */
  postRaidNow: (guildId: number, raidId: string) => Promise<GuildRaidView>;
  /** The assignments (officers): a slot's names in order, a cleared slot, the Discord post. */
  putAssignment: (
    guildId: number,
    raidId: string,
    encounterId: number,
    slot: string,
    characterKeys: readonly string[],
  ) => Promise<GuildRaidView>;
  deleteAssignment: (
    guildId: number,
    raidId: string,
    encounterId: number,
    slot: string,
  ) => Promise<GuildRaidView>;
  postAssignments: (guildId: number, raidId: string) => Promise<AssignmentPostResult>;
  /** The edit form's Raid list (the raids page carries it). */
  loadZones: (guildId: number) => Promise<GuildZoneView[]>;
  isVisible: () => boolean;
  now: () => Date;
  storage: Pick<Storage, 'getItem' | 'setItem' | 'removeItem'> | null;
}

/** R1: only a verified member reads the raid; everyone else gets the not-visible page. */
export type RaidAccess = 'loading' | 'signed-out' | 'not-member' | 'unverified' | 'member';
export type RaidLoad = 'loading' | 'ready' | 'failed';
export type RaidPageRole = 'officer' | 'member';

export interface RaidBandGuild {
  id: number;
  name: string;
  faction: 'alliance' | 'horde' | null;
  mark: string | null;
}

/** Where a failure line sits: under the officer row's button, under the roster, or under a row. */
export type OfficerControl = RaidStatusAction | 'roster' | `roster:${string}`;

export interface OfficerLine {
  control: OfficerControl;
  message: string;
}

/** The fields a component test seeds, since a server render runs no effects. */
export interface RaidPageSeed {
  guild: RaidBandGuild | null;
  access: RaidAccess;
  load: RaidLoad;
  raid: GuildRaidView | null;
  role: RaidPageRole;
  officerToolsField: boolean | undefined;
  officerToolsEnded: boolean;
  currentKey: string | null;
  pollPaused: boolean;
  line: OfficerLine | null;
  raceLine: string | null;
  logMatched: boolean;
}
