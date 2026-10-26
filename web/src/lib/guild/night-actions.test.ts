// web/src/lib/guild/night-actions.test.ts
// The done raid's officer writes (attendance spec 3.4, 3.5, 4): Excused is optimistic (row and summary
// move at once) and reverts with the row line on a refusal; Post to Discord posts once, re-reads on
// 409 already_posted, leaves the button live on 502 post_failed, and flips the plan on 403.
import { describe, expect, it, vi } from 'vitest';
import type { Me } from '../account/api';
import { deferred, deps, me, started, useControllerClock } from '../../test-support/raid-controller-fakes';
import { doneRaid } from '../../test-support/raid-fixtures';
import { attendanceRow, raidReport } from '../../test-support/report-fixtures';
import { GuildApiError, type GuildHome, type GuildRaidView } from './api';
import { attendanceCopy } from './attendance-copy';

useControllerClock();

const KEY = 'us-pvp-grimtotem';
const officerHome = (): Promise<GuildHome> =>
  Promise.resolve({
    viewer: { role: 'officer', character_key: null, verified: true, officer_tools: true },
  } as unknown as GuildHome);
const OFFICER = {
  loadMe: (): Promise<Me> => Promise.resolve(me({ rank: 'officer' })),
  loadHome: officerHome,
};

function night(excused = false): GuildRaidView {
  return doneRaid({
    attendance: [attendanceRow({ character_key: KEY, name: 'Grimtotem', outcome: 'no_show', excused })],
    report: raidReport(),
  });
}

describe('Excused', () => {
  it('moves the row and the summary at once, then keeps the server answer', async () => {
    const answer = deferred<GuildRaidView>();
    const putExcused = vi.fn(() => answer.promise);
    const { c, stop } = await started(
      deps({ ...OFFICER, fetchRaid: () => Promise.resolve(night()), putExcused }),
    );
    const pending = c.night.excuse(KEY, true);
    expect(c.raid?.attendance?.[0].excused).toBe(true);
    expect(c.raid?.night?.excused).toBe(1);
    expect(c.night.excuseBusy[KEY]).toBe(true);
    answer.resolve(night(true));
    await expect(pending).resolves.toBe(true);
    expect(putExcused).toHaveBeenCalledWith(2, '7', KEY, true);
    expect(c.night.excuseBusy[KEY]).toBeUndefined();
    stop();
  });

  it('reverts and says it did not save on a refusal', async () => {
    const putExcused = vi.fn(() => Promise.reject(new Error('down')));
    const { c, stop } = await started(
      deps({ ...OFFICER, fetchRaid: () => Promise.resolve(night()), putExcused }),
    );
    await expect(c.night.excuse(KEY, true)).resolves.toBe(false);
    expect(c.raid?.attendance?.[0].excused).toBe(false);
    expect(c.raid?.night?.excused).toBe(0);
    expect(c.night.excuseLines[KEY]).toBe(attendanceCopy.errors.rows);
    stop();
  });

  it('ends the officer tools on guild_plan_required', async () => {
    const refusal = new GuildApiError('plan', 403, 'guild_plan_required');
    const { c, stop } = await started(
      deps({
        ...OFFICER,
        fetchRaid: () => Promise.resolve(night()),
        putExcused: () => Promise.reject(refusal),
      }),
    );
    await c.night.excuse(KEY, true);
    expect(c.night.excuseLines[KEY]).toBe(attendanceCopy.errors.officerToolsEnded);
    expect(c.officerTools).toBe(false);
    stop();
  });
});

describe('Post to Discord', () => {
  const POSTED_AT = '2026-12-11T05:58:00Z';

  it('posts once and stamps the report', async () => {
    const answer = deferred<{ posted_at: string }>();
    const postReport = vi.fn(() => answer.promise);
    const { c, stop } = await started(
      deps({ ...OFFICER, fetchRaid: () => Promise.resolve(night()), postReport }),
    );
    const first = c.night.post();
    void c.night.post();
    expect(c.night.posting).toBe(true);
    answer.resolve({ posted_at: POSTED_AT });
    await first;
    expect(postReport).toHaveBeenCalledTimes(1);
    expect(c.raid?.report?.posted_at).toBe(POSTED_AT);
    expect(c.night.posting).toBe(false);
    stop();
  });

  it('shows the posted line after the re-read on 409 already_posted', async () => {
    const posted = doneRaid({ attendance: [], report: raidReport({ posted_at: POSTED_AT }) });
    const refetchRaid = vi.fn(() => Promise.resolve(posted));
    const postReport = vi.fn(() => Promise.reject(new GuildApiError('again', 409, 'already_posted')));
    const { c, stop } = await started(
      deps({ ...OFFICER, fetchRaid: () => Promise.resolve(night()), refetchRaid, postReport }),
    );
    await c.night.post();
    expect(c.raid?.report?.posted_at).toBe(POSTED_AT);
    expect(c.night.postLine).toBeNull();
    stop();
  });

  it('says it could not post on 502 and leaves the button live', async () => {
    const postReport = vi.fn(() => Promise.reject(new GuildApiError('bad', 502, 'post_failed')));
    const { c, stop } = await started(
      deps({ ...OFFICER, fetchRaid: () => Promise.resolve(night()), postReport }),
    );
    await c.night.post();
    expect(c.night.postLine).toBe(attendanceCopy.errors.post);
    expect(c.night.posting).toBe(false);
    expect(c.raid?.report?.posted_at).toBeNull();
    stop();
  });
});
