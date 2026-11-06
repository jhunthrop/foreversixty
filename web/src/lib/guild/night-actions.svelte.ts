// web/src/lib/guild/night-actions.svelte.ts
// The raid page's officer writes on a done raid (web layout spec 2026-11-05 §7.2 "Raid report"):
// the attendance fix (present, absent, bench or excused over the log, or back to the log; the
// control shows the pick while it saves and a refusal puts the row back with its line) and the
// Discord summary (`Post to Discord`, then `Post again`; 409 `already_posted` re-reads the raid so
// the posted line shows; 502 `post_failed` leaves the button live). Both run through the
// controller's officer-action pattern: `host.write` (the writeSeq guard that drops a spanning
// poll) and `host.refusalLine` (the plan flip).
import { GuildApiError, type AttendanceOverrideStatus, type GuildRaidView } from './api';
import { attendanceCopy } from './attendance-copy';
import { withPosted } from './night-patch';
import type { RaidPageDeps, RaidRoute } from './raid-page-types';

const ALREADY_POSTED = 'already_posted';

type NightDeps = Pick<RaidPageDeps, 'putOverride' | 'deleteOverride' | 'postReport' | 'refetchRaid'>;

/** An officer's choice for one row: a status over the log, or `null` for back to the log. */
export type AttendanceFix = AttendanceOverrideStatus | null;

/** What the controller lends: its raid, its guarded write, its raid setter and the plan flip. */
export interface NightActionHost {
  raid(): GuildRaidView | null;
  write<T>(run: (guildId: number) => Promise<T>): Promise<T>;
  applyRaid(raid: GuildRaidView): void;
  refusalLine(thrown: unknown, ended: string, failed: string): string;
}

export class NightActions {
  /** The fix each row is saving, by character (`null` is back to the log). */
  fixing = $state<Readonly<Record<string, AttendanceFix>>>({});
  /** A refused fix's line, by character. */
  fixLines = $state<Readonly<Record<string, string>>>({});
  posting = $state(false);
  /** `Could not post. Try again.` (or the plan sentence) under the Discord buttons. */
  postLine = $state<string | null>(null);

  private readonly host: NightActionHost;
  private readonly deps: NightDeps;
  private readonly raidId: string;

  constructor(host: NightActionHost, deps: NightDeps, route: RaidRoute) {
    this.host = host;
    this.deps = deps;
    this.raidId = route.raidId;
  }

  private setFixing(key: string, fix: AttendanceFix | undefined): void {
    const { [key]: _dropped, ...rest } = this.fixing;
    this.fixing = fix === undefined ? rest : { ...rest, [key]: fix };
  }

  private setLine(key: string, line: string | null): void {
    const { [key]: _dropped, ...rest } = this.fixLines;
    this.fixLines = line === null ? rest : { ...rest, [key]: line };
  }

  /** Saves one row's fix; the server's raid (every count moved) replaces the page's. */
  fix = async (characterKey: string, fix: AttendanceFix): Promise<boolean> => {
    if (characterKey in this.fixing) return false;
    this.setLine(characterKey, null);
    this.setFixing(characterKey, fix);
    try {
      const fresh = await this.host.write((id) =>
        fix === null
          ? this.deps.deleteOverride(id, this.raidId, characterKey)
          : this.deps.putOverride(id, this.raidId, characterKey, fix),
      );
      this.host.applyRaid(fresh);
      return true;
    } catch (thrown) {
      this.setLine(
        characterKey,
        this.host.refusalLine(thrown, attendanceCopy.errors.officerToolsEnded, attendanceCopy.errors.rows),
      );
      return false;
    } finally {
      this.setFixing(characterKey, undefined);
    }
  };

  /** Posts the summary (again, with `repost`); a press while one is in flight does nothing. */
  post = async (repost = false): Promise<void> => {
    if (this.posting) return;
    this.posting = true;
    this.postLine = null;
    try {
      const { posted_at: postedAt } = await this.host.write((id) =>
        this.deps.postReport(id, this.raidId, repost),
      );
      const raid = this.host.raid();
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
