// web/src/lib/guild/raid-page-controller.svelte.ts
// The raid page's state (design/specs/2026-10-26-guild-raids.md 4 "Page states", "Poll", "Races",
// R1): who may see the raid, its load, the 10 s poll while visible and open, the optimistic sign-up,
// the race codes, the officer actions and the plan gate (officer-tools gate spec 4). Every network
// call and clock comes in through `RaidPageDeps`, so the tests drive it with fakes and fake timers.
import type { Me } from '../account/api';
import { characterSlug, type CharacterPath } from '../characters';
import type { GuildPage } from '../rankings/api';
import {
  GuildApiError,
  type GuildHome,
  type GuildRaidView,
  type GuildSignupState,
  type GuildViewerRole,
  type GuildZoneView,
} from './api';
import { guildMarkSrc } from './mark';
import { raidCopy } from './raid-copy';
import type { RaidSubmitResult } from './raid-form';
import { POLL_FAILURES_TO_PAUSE, RAID_POLL_MS, pollOpen, rememberStripRows } from './raid-page';
import type { RosterRow } from './raid-roster';
import type { RaidInput, RaidStatusAction, RaidStatusBody, SignupInput, SignupResult } from './raids-api';
import { submitRefusal } from './raids-tab-controller';
import type { SignupHandler, SignupOutcome } from './signup-strip';

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
  putSignup: (guildId: number, raidId: string, key: string, body: SignupInput) => Promise<SignupResult>;
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
  /** The edit form's Raid list (the raids page carries it). */
  loadZones: (guildId: number) => Promise<GuildZoneView[]>;
  isVisible: () => boolean;
  now: () => Date;
  storage: Pick<Storage, 'getItem' | 'setItem'> | null;
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
}

const PLAN_REQUIRED = 'guild_plan_required';
const RACE_STATUS: Readonly<Record<string, GuildRaidView['status']>> = {
  raid_locked: 'locked',
  raid_cancelled: 'cancelled',
};

function codeOf(thrown: unknown): string {
  return thrown instanceof GuildApiError ? thrown.code : '';
}

function rankRole(rank: string | undefined): RaidPageRole {
  return rank === 'officer' || rank === 'leader' ? 'officer' : 'member';
}

function homeRole(role: GuildViewerRole | undefined, fallback: RaidPageRole): RaidPageRole {
  if (role === undefined) return fallback;
  return role === 'officer' ? 'officer' : 'member';
}

export class RaidPageController {
  guild = $state<RaidBandGuild | null>(null);
  access = $state<RaidAccess>('loading');
  load = $state<RaidLoad>('loading');
  raid = $state<GuildRaidView | null>(null);
  role = $state<RaidPageRole>('member');
  officerToolsField = $state<boolean | undefined>(undefined);
  officerToolsEnded = $state(false);
  currentKey = $state<string | null>(null);
  pollPaused = $state(false);
  /** The officer row's action in flight. */
  busy = $state<RaidStatusAction | null>(null);
  /** The roster row whose officer select is saving. */
  rosterBusy = $state<string | null>(null);
  line = $state<OfficerLine | null>(null);
  /** The edit form is open in place of the band's lines 3 to 6. */
  editing = $state(false);
  zones = $state<GuildZoneView[]>([]);
  /** `This raid was cancelled while you were signing up.`, shown above the cancelled strip. */
  raceLine = $state<string | null>(null);

  private readonly deps: RaidPageDeps;
  private readonly route: RaidRoute;
  private timer: ReturnType<typeof setInterval> | undefined;
  private polling = false;
  private pollFailures = 0;
  /** Bumped at each write's start and end: a poll that spans one is stale. */
  private writeSeq = 0;
  private writes = 0;
  private stopped = false;

  constructor(deps: RaidPageDeps, route: RaidRoute, seed: Partial<RaidPageSeed> = {}) {
    this.deps = deps;
    this.route = route;
    Object.assign(this, seed);
  }

  /** Officer-read-only once the API says so or refused a write with the plan code. */
  get officerTools(): boolean {
    return this.role !== 'officer' || ((this.officerToolsField ?? true) && !this.officerToolsEnded);
  }

  /** Loads the page and starts the poll; answers the stop. */
  start(): () => void {
    clearInterval(this.timer);
    this.stopped = false;
    void this.loadPage();
    this.timer = setInterval(() => void this.poll(), RAID_POLL_MS);
    return () => {
      this.stopped = true;
      clearInterval(this.timer);
      this.timer = undefined;
    };
  }

  retry(): void {
    void this.loadPage();
  }

  private async loadPage(): Promise<void> {
    this.load = 'loading';
    const [page, me] = await Promise.allSettled([this.deps.loadGuild(), this.deps.loadMe()]);
    if (this.stopped) return;
    // A failed /v1/me is not a signed-out answer (`null`): the page fails and offers a retry.
    if (page.status === 'rejected' || me.status === 'rejected') {
      this.load = 'failed';
      return;
    }
    const { guild } = page.value;
    this.guild = {
      id: guild.id,
      name: guild.name,
      faction: guild.faction ?? null,
      mark: guildMarkSrc(guild),
    };
    this.resolveAccess(me.value);
    if (this.access !== 'member') {
      this.load = 'ready';
      return;
    }
    void this.deps
      .loadHome(guild.id)
      .then((home) => {
        if (!this.stopped) this.applyHome(home);
      })
      .catch(() => {
        // The rank from /v1/me stands; the server's 403 is the real gate.
      });
    try {
      const raid = await this.deps.fetchRaid(guild.id, this.route.raidId);
      if (this.stopped) return;
      this.applyRaid(raid);
      this.load = 'ready';
    } catch {
      if (!this.stopped) this.load = 'failed';
    }
  }

  private resolveAccess(me: Me | null): void {
    const { region, ruleset, slug } = this.route.guild;
    const membership = me?.guilds.find(
      (g) => g.region === region && g.ruleset === ruleset && characterSlug(g.name) === slug,
    );
    if (me === null) this.access = 'signed-out';
    else if (membership === undefined) this.access = 'not-member';
    else if (!membership.verified) this.access = 'unverified';
    else {
      this.access = 'member';
      this.role = rankRole(membership.rank);
    }
  }

  private applyHome(home: GuildHome): void {
    this.role = homeRole(home.viewer?.role, this.role);
    this.officerToolsField = home.viewer?.officer_tools;
    this.currentKey = home.viewer?.character_key ?? null;
  }

  private applyRaid(raid: GuildRaidView): void {
    this.raid = raid;
    rememberStripRows(raid, this.deps.storage);
  }

  private canPoll(): boolean {
    return (
      !this.stopped &&
      !this.pollPaused &&
      !this.polling &&
      this.raid !== null &&
      this.guild !== null &&
      this.deps.isVisible() &&
      pollOpen(this.raid, this.deps.now())
    );
  }

  private async poll(): Promise<void> {
    if (!this.canPoll() || this.guild === null) return;
    this.polling = true;
    const seq = this.writeSeq;
    try {
      const fresh = await this.deps.refetchRaid(this.guild.id, this.route.raidId);
      this.pollFailures = 0;
      if (seq === this.writeSeq && this.writes === 0 && !this.stopped) this.applyRaid(fresh);
    } catch {
      this.pollFailures += 1;
      if (this.pollFailures >= POLL_FAILURES_TO_PAUSE) this.pollPaused = true;
    } finally {
      this.polling = false;
    }
  }

  /** Runs one write; a poll that spans it is dropped, so a saving row keeps its own answer. */
  private async write<T>(run: (guildId: number) => Promise<T>): Promise<T> {
    if (this.guild === null) throw new Error('the raid page wrote before its guild loaded');
    this.writes += 1;
    this.writeSeq += 1;
    try {
      return await run(this.guild.id);
    } finally {
      this.writes -= 1;
      this.writeSeq += 1;
    }
  }

  private planRefused(thrown: unknown): boolean {
    if (codeOf(thrown) !== PLAN_REQUIRED) return false;
    this.officerToolsEnded = true;
    return true;
  }

  /** The strip's handler: the answer is applied before the row hears `ok`. */
  signup: SignupHandler = async (key, state, role, note) => {
    try {
      const result = await this.write((id) =>
        this.deps.putSignup(id, this.route.raidId, key, { state, role, note }),
      );
      this.applyRaid(result.raid);
      return { ok: true, moved: result.moved };
    } catch (thrown) {
      return this.signupRefused(thrown);
    }
  };

  private signupRefused(thrown: unknown): SignupOutcome {
    const code = codeOf(thrown);
    const status = RACE_STATUS[code];
    if (status !== undefined && this.raid !== null) this.raid = { ...this.raid, status };
    // The cancelled strip drops its rows, so the race line moves up to the page (spec 4 "Races").
    if (status === 'cancelled') this.raceLine = raidCopy.races.cancelled;
    this.planRefused(thrown);
    return { ok: false, code, message: raidCopy.stripStatus.failed };
  }

  /** An officer's inline state select: the row's spec role, its note kept. */
  async setRosterState(row: RosterRow, state: GuildSignupState): Promise<boolean> {
    this.rosterBusy = row.character_key;
    this.line = null;
    try {
      const result = await this.write((id) =>
        this.deps.putSignup(id, this.route.raidId, row.character_key, {
          state,
          role: row.role,
          note: row.note,
        }),
      );
      this.applyRaid(result.raid);
      return true;
    } catch (thrown) {
      this.line = this.planRefused(thrown)
        ? { control: `roster:${row.character_key}`, message: raidCopy.races.officerToolsEnded }
        : { control: 'roster', message: raidCopy.officerRow.failed };
      return false;
    } finally {
      this.rosterBusy = null;
    }
  }

  /** Lock, Unlock or Cancel; a cancel carries the Discord choice. */
  async runStatus(action: RaidStatusAction, postToDiscord?: boolean): Promise<boolean> {
    this.busy = action;
    this.line = null;
    const body = postToDiscord === undefined ? undefined : { post_to_discord: postToDiscord };
    try {
      this.applyRaid(await this.write((id) => this.deps.setStatus(id, this.route.raidId, action, body)));
      return true;
    } catch (thrown) {
      if (this.planRefused(thrown)) {
        this.line = { control: action, message: raidCopy.races.officerToolsEnded };
        return false;
      }
      this.line = { control: action, message: raidCopy.officerRow.failed };
      if (codeOf(thrown) === 'conflict') this.refreshAfterConflict();
      return false;
    } finally {
      this.busy = null;
    }
  }

  /** Someone else moved the raid on: show what it is now. */
  private refreshAfterConflict(): void {
    if (this.guild === null) return;
    void this.deps
      .refetchRaid(this.guild.id, this.route.raidId)
      .then((raid) => this.applyRaid(raid))
      .catch(() => {
        // The failure line already says it did not save.
      });
  }

  /** Opens the edit form: the raid's own zone at once, the guild's list once it answers. */
  openEdit(): void {
    this.line = null;
    this.editing = true;
    const { raid, guild } = this;
    if (raid === null || guild === null) return;
    if (this.zones.length === 0) this.zones = [{ key: raid.zone, name: raid.zone_name, size: raid.size }];
    void this.deps
      .loadZones(guild.id)
      .then((zones) => {
        if (zones.length > 0) this.zones = zones;
      })
      .catch(() => {
        // The raid's own zone stays the one choice.
      });
  }

  closeEdit(): void {
    this.editing = false;
  }

  /** The edit form's submit; a saved edit closes the form. */
  saveEdit = async (payload: RaidInput): Promise<RaidSubmitResult> => {
    const { repeat_weeks: _create, ...body } = payload;
    try {
      this.applyRaid(await this.write((id) => this.deps.updateRaid(id, this.route.raidId, body)));
      this.editing = false;
      return { ok: true };
    } catch (thrown) {
      this.planRefused(thrown);
      return submitRefusal(thrown);
    }
  };
}
