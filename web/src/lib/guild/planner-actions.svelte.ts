// web/src/lib/guild/planner-actions.svelte.ts
// The raid planner's officer writes beyond the sign-up state (web layout spec 2026-11-05 §7.2):
// the comp target's Save and Back to the default (a 422 keeps its per-field lines), and Ping on
// Discord (one message under the raid's post, owner decision D16). They run through the
// controller's guarded write, as NightActions does, so a poll that spans one is dropped.
import { GuildApiError, type GuildRaidView } from './api';
import type { NightActionHost } from './night-actions.svelte';
import type { GuildCompRole } from './overview-api';
import { pingLine } from './overview-role-view';
import { raidCopy, raidPostCopy } from './raid-copy';
import type { RaidPageDeps, RaidRoute } from './raid-page-types';
import { withRetry } from '../ui/copy';

type PlannerDeps = Pick<
  RaidPageDeps,
  'putCompTarget' | 'deleteCompTarget' | 'pingRaid' | 'postRaidNow' | 'now'
>;

const PING_FAILED = withRetry('The ping did not go out.');
const COMP_FAILED = withRetry('The target did not save.');

export class PlannerActions {
  compBusy = $state<'save' | 'reset' | null>(null);
  /** A 422's lines by the API's field (`tanks`, `total`). */
  compFields = $state<Readonly<Record<string, string>>>({});
  compLine = $state<string | null>(null);
  pinging = $state(false);
  /** The ping's answer, or why it failed. */
  pingAnswer = $state<string | null>(null);
  posting = $state(false);
  /** Post now / Repost's answer, or why it failed. */
  postAnswer = $state<string | null>(null);

  private readonly host: NightActionHost;
  private readonly deps: PlannerDeps;
  private readonly raidId: string;

  constructor(host: NightActionHost, deps: PlannerDeps, route: RaidRoute) {
    this.host = host;
    this.deps = deps;
    this.raidId = route.raidId;
  }

  /** Field lines the browser found before the API (a blank or a letter). */
  refuseComp(fields: Readonly<Record<string, string>>): void {
    this.compFields = fields;
    this.compLine = null;
  }

  saveComp = (target: Readonly<Record<GuildCompRole, number>>): Promise<boolean> =>
    this.compWrite('save', (id) => this.deps.putCompTarget(id, this.raidId, target));

  resetComp = (): Promise<boolean> =>
    this.compWrite('reset', (id) => this.deps.deleteCompTarget(id, this.raidId));

  private async compWrite(
    busy: 'save' | 'reset',
    run: (guildId: number) => Promise<GuildRaidView>,
  ): Promise<boolean> {
    this.compBusy = busy;
    this.compFields = {};
    this.compLine = null;
    try {
      this.host.applyRaid(await this.host.write(run));
      return true;
    } catch (thrown) {
      if (thrown instanceof GuildApiError && thrown.fields !== null) this.compFields = thrown.fields;
      else this.compLine = this.host.refusalLine(thrown, raidCopy.races.officerToolsEnded, COMP_FAILED);
      return false;
    } finally {
      this.compBusy = null;
    }
  }

  /** Posts the raid to Discord now, or reposts it as a new message; the view comes back. */
  postNow = async (): Promise<boolean> => {
    if (this.posting) return false;
    const repost = this.host.raid()?.discord_post?.posted_at != null;
    this.posting = true;
    try {
      this.host.applyRaid(await this.host.write((id) => this.deps.postRaidNow(id, this.raidId)));
      this.postAnswer = repost ? raidPostCopy.reposted : raidPostCopy.posted;
      return true;
    } catch (thrown) {
      this.postAnswer =
        thrown instanceof GuildApiError && thrown.code === 'no_channel'
          ? raidPostCopy.noChannel
          : this.host.refusalLine(thrown, raidCopy.races.officerToolsEnded, raidPostCopy.failed);
      return false;
    } finally {
      this.posting = false;
    }
  };

  /** Pings the non-signers; the raid's throttle moves to the answer's next allowed time. */
  ping = async (): Promise<boolean> => {
    if (this.pinging) return false;
    this.pinging = true;
    try {
      const result = await this.host.write((id) => this.deps.pingRaid(id, this.raidId));
      this.pingAnswer = pingLine(result);
      const raid = this.host.raid();
      if (raid !== null)
        this.host.applyRaid({
          ...raid,
          ping: { last_pinged_at: this.deps.now().toISOString(), next_allowed_at: result.next_allowed_at },
        });
      return true;
    } catch (thrown) {
      this.pingAnswer = this.host.refusalLine(thrown, raidCopy.races.officerToolsEnded, PING_FAILED);
      return false;
    } finally {
      this.pinging = false;
    }
  };
}
