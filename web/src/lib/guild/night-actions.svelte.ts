// web/src/lib/guild/night-actions.svelte.ts
// The raid page's two officer writes on a done raid (attendance spec 3.4, 3.5, 4 "States"): the
// Excused tick (optimistic, reverting with the row's line on a refusal) and `Post to Discord` (once;
// 409 `already_posted` re-reads the raid so the posted line shows; 502 `post_failed` leaves the
// button live). Both run through the controller's officer-action pattern: `host.write` (the writeSeq
// guard that drops a spanning poll) and `host.refusalLine` (the plan flip).
import { GuildApiError, type GuildRaidView } from './api';
import { attendanceCopy } from './attendance-copy';
import { withExcused, withPosted } from './night-patch';
import type { RaidPageDeps, RaidRoute } from './raid-page-types';

const ALREADY_POSTED = 'already_posted';

type NightDeps = Pick<RaidPageDeps, 'putExcused' | 'postReport' | 'refetchRaid'>;

/** What the controller lends: its raid, its guarded write, its raid setter and the plan flip. */
export interface NightActionHost {
  readonly raid: GuildRaidView | null;
  write<T>(run: (guildId: number) => Promise<T>): Promise<T>;
  applyRaid(raid: GuildRaidView): void;
  refusalLine(thrown: unknown, ended: string, failed: string): string;
}

export class NightActions {
  /** Characters whose Excused write is in flight. */
  excuseBusy = $state<Readonly<Record<string, boolean>>>({});
  /** A refused Excused write's line, by character. */
  excuseLines = $state<Readonly<Record<string, string>>>({});
  posting = $state(false);
  /** `Could not post. Try again.` (or the plan sentence) under the header buttons. */
  postLine = $state<string | null>(null);

  private readonly host: NightActionHost;
  private readonly deps: NightDeps;
  private readonly raidId: string;

  constructor(host: NightActionHost, deps: NightDeps, route: RaidRoute) {
    this.host = host;
    this.deps = deps;
    this.raidId = route.raidId;
  }

  private setBusy(key: string, busy: boolean): void {
    const { [key]: _dropped, ...rest } = this.excuseBusy;
    this.excuseBusy = busy ? { ...rest, [key]: true } : rest;
  }

  private setLine(key: string, line: string | null): void {
    const { [key]: _dropped, ...rest } = this.excuseLines;
    this.excuseLines = line === null ? rest : { ...rest, [key]: line };
  }

  /** Ticks or unticks Excused: the row and the summary move at once; a refusal moves them back. */
  excuse = async (characterKey: string, excused: boolean): Promise<boolean> => {
    this.setLine(characterKey, null);
    this.setBusy(characterKey, true);
    try {
      const fresh = await this.host.write((id) => {
        const { raid } = this.host;
        if (raid !== null) this.host.applyRaid(withExcused(raid, characterKey, excused));
        return this.deps.putExcused(id, this.raidId, characterKey, excused);
      });
      this.host.applyRaid(fresh);
      return true;
    } catch (thrown) {
      const { raid } = this.host;
      if (raid !== null) this.host.applyRaid(withExcused(raid, characterKey, !excused));
      this.setLine(
        characterKey,
        this.host.refusalLine(thrown, attendanceCopy.errors.officerToolsEnded, attendanceCopy.errors.rows),
      );
      return false;
    } finally {
      this.setBusy(characterKey, false);
    }
  };

  /** Posts the report once; a press while one is in flight does nothing. */
  post = async (): Promise<void> => {
    if (this.posting) return;
    this.posting = true;
    this.postLine = null;
    try {
      const { posted_at: postedAt } = await this.host.write((id) => this.deps.postReport(id, this.raidId));
      const { raid } = this.host;
      if (raid !== null) this.host.applyRaid(withPosted(raid, postedAt));
    } catch (thrown) {
      await this.postRefused(thrown);
    } finally {
      this.posting = false;
    }
  };

  /** Another officer posted first: re-read, so the posted line shows. Anything else is a line. */
  private async postRefused(thrown: unknown): Promise<void> {
    if (thrown instanceof GuildApiError && thrown.code === ALREADY_POSTED) {
      try {
        this.host.applyRaid(await this.host.write((id) => this.deps.refetchRaid(id, this.raidId)));
        return;
      } catch {
        // The re-read failed too: fall through to the retry line.
      }
    }
    this.postLine = this.host.refusalLine(
      thrown,
      attendanceCopy.errors.officerToolsEnded,
      attendanceCopy.errors.post,
    );
  }
}
