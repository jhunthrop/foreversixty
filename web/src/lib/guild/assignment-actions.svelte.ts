// web/src/lib/guild/assignment-actions.svelte.ts
// The raid assignments' officer writes (web layout spec 2026-11-05 §7.2): set a slot's names
// (an add or a remove sends the slot whole), clear it, and Post to Discord (one message under the
// raid's post). They run through the controller's guarded write, as PlannerActions does, so a
// poll that spans one is dropped. Copy as raid warning is the component's: no write.
import { GuildApiError } from './api';
import type { AssignmentSlot } from './assignments-api';
import { assignmentsCopy, withAdded, withRemoved } from './assignments-view';
import type { NightActionHost } from './night-actions.svelte';
import { raidCopy } from './raid-copy';
import type { RaidPageDeps, RaidRoute } from './raid-page-types';

type AssignmentDeps = Pick<RaidPageDeps, 'putAssignment' | 'deleteAssignment' | 'postAssignments'>;

/** The post's refusals whose API sentence the officer reads as is (raider voice, Go raid_assign_post.go). */
const SAID_BY_API: ReadonlySet<string> = new Set([
  'discord_not_connected',
  'nothing_assigned',
  'assignments_throttled',
  'discord_refused',
]);

/** `${encounterId}:${slotKey}`: which slot a busy state or a line belongs to. */
export const slotId = (encounterId: number, slotKey: string): string => `${encounterId}:${slotKey}`;

export class AssignmentActions {
  /** The slot whose write is in flight. */
  busySlot = $state<string | null>(null);
  /** A refused write's line, under its slot. */
  slotLine = $state<{ slot: string; message: string } | null>(null);
  posting = $state(false);
  /** The post's answer, or why it failed. */
  postAnswer = $state<string | null>(null);

  private readonly host: NightActionHost;
  private readonly deps: AssignmentDeps;
  private readonly raidId: string;

  constructor(host: NightActionHost, deps: AssignmentDeps, route: RaidRoute) {
    this.host = host;
    this.deps = deps;
    this.raidId = route.raidId;
  }

  add = (encounterId: number, slot: AssignmentSlot, key: string): Promise<boolean> =>
    this.setKeys(encounterId, slot.key, withAdded(slot, key));

  remove = (encounterId: number, slot: AssignmentSlot, key: string): Promise<boolean> =>
    this.setKeys(encounterId, slot.key, withRemoved(slot, key));

  clear = (encounterId: number, slot: AssignmentSlot): Promise<boolean> =>
    this.slotWrite(slotId(encounterId, slot.key), (id) =>
      this.deps.deleteAssignment(id, this.raidId, encounterId, slot.key),
    );

  private setKeys(encounterId: number, slotKey: string, keys: string[]): Promise<boolean> {
    return this.slotWrite(slotId(encounterId, slotKey), (id) =>
      this.deps.putAssignment(id, this.raidId, encounterId, slotKey, keys),
    );
  }

  private async slotWrite(
    slot: string,
    run: (guildId: number) => ReturnType<AssignmentDeps['putAssignment']>,
  ): Promise<boolean> {
    if (this.busySlot !== null) return false;
    this.busySlot = slot;
    this.slotLine = null;
    try {
      this.host.applyRaid(await this.host.write(run));
      return true;
    } catch (thrown) {
      const field = thrown instanceof GuildApiError ? thrown.fields?.character_keys : undefined;
      this.slotLine = {
        slot,
        message:
          field ??
          this.host.refusalLine(thrown, raidCopy.races.officerToolsEnded, assignmentsCopy.saveFailed),
      };
      return false;
    } finally {
      this.busySlot = null;
    }
  }

  /** Posts every slot to Discord; the raid's posted_at moves to the answer's. */
  post = async (): Promise<boolean> => {
    if (this.posting) return false;
    this.posting = true;
    try {
      const result = await this.host.write((id) => this.deps.postAssignments(id, this.raidId));
      this.postAnswer = assignmentsCopy.posted;
      const raid = this.host.raid();
      if (raid?.assignments)
        this.host.applyRaid({ ...raid, assignments: { ...raid.assignments, posted_at: result.posted_at } });
      return true;
    } catch (thrown) {
      this.postAnswer =
        thrown instanceof GuildApiError && SAID_BY_API.has(thrown.code)
          ? thrown.message
          : this.host.refusalLine(thrown, raidCopy.races.officerToolsEnded, assignmentsCopy.postFailed);
      return false;
    } finally {
      this.posting = false;
    }
  };
}
