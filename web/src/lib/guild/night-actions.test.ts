// web/src/lib/guild/night-actions.test.ts
// The done raid's officer writes (web layout spec 2026-11-05 §7.2): the attendance fix shows its pick
// while it saves, applies the server's raid and keeps the row with its line on a refusal; Post to
// Discord posts once (again with repost), re-reads on 409 already_posted, leaves the button live on
// 502 post_failed, and flips the plan on 403.
import { describe, expect, it, vi } from 'vitest';
import type { Me } from '../account/api';
import { deferred, deps, me, started, useControllerClock } from '../../test-support/raid-controller-fakes';
import { doneRaid } from '../../test-support/raid-fixtures';
import { attendanceRow, raidReport } from '../../test-support/report-fixtures';
import { GuildApiError, type GuildHome, type GuildRaidView, type RaidAttendanceRow } from './api';
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

function night(override: RaidAttendanceRow['override'] = null): GuildRaidView {
  return doneRaid({
    attendance: [attendanceRow({ character_key: KEY, name: 'Grimtotem', outcome: 'no_show', override })],
    report: raidReport(),
  });
}

const BENCH = { status: 'bench' as const, by_name: 'Kraggor', at: '2026-12-11T06:00:00Z' };

describe('the attendance fix', () => {
  it('shows the pick while it saves, then keeps the server answer', async () => {
    const answer = deferred<GuildRaidView>();
    const putOverride = vi.fn(() => answer.promise);
    const { c, stop } = await started(
      deps({ ...OFFICER, fetchRaid: () => Promise.resolve(night()), putOverride }),
    );
    const pending = c.night.fix(KEY, 'bench');
    expect(c.night.fixing[KEY]).toBe('bench');
    await expect(c.night.fix(KEY, 'present')).resolves.toBe(false);
    answer.resolve(night(BENCH));
    await expect(pending).resolves.toBe(true);
    expect(putOverride).toHaveBeenCalledTimes(1);
    expect(putOverride).toHaveBeenCalledWith(2, '7', KEY, 'bench');
    expect(c.raid?.attendance?.[0].override).toEqual(BENCH);
    expect(c.night.fixing).toEqual({});
    stop();
  });

  it('goes back to the log with a delete', async () => {
    const deleteOverride = vi.fn(() => Promise.resolve(night()));
    const { c, stop } = await started(
      deps({ ...OFFICER, fetchRaid: () => Promise.resolve(night(BENCH)), deleteOverride }),
    );
    await expect(c.night.fix(KEY, null)).resolves.toBe(true);
    expect(deleteOverride).toHaveBeenCalledWith(2, '7', KEY);
    expect(c.raid?.attendance?.[0].override).toBeNull();
    stop();
  });

  it('keeps the row and says it did not save on a refusal', async () => {
    const putOverride = vi.fn(() => Promise.reject(new Error('down')));
    const { c, stop } = await started(
      deps({ ...OFFICER, fetchRaid: () => Promise.resolve(night()), putOverride }),
    );
    await expect(c.night.fix(KEY, 'present')).resolves.toBe(false);
    expect(c.raid?.attendance?.[0].override).toBeNull();
    expect(c.night.fixLines[KEY]).toBe(attendanceCopy.errors.rows);
    stop();
  });

  it('ends the officer tools on guild_plan_required', async () => {
    const refusal = new GuildApiError('plan', 403, 'guild_plan_required');
    const { c, stop } = await started(
      deps({
        ...OFFICER,
        fetchRaid: () => Promise.resolve(night()),
        putOverride: () => Promise.reject(refusal),
      }),
    );
    await c.night.fix(KEY, 'excused');
    expect(c.night.fixLines[KEY]).toBe(attendanceCopy.errors.officerToolsEnded);
    expect(c.officerTools).toBe(false);
    stop();
  });
});

describe('Post to Discord', () => {
  const POSTED_AT = '2026-12-11T05:58:00Z';

  it('posts again with repost', async () => {
    const postReport = vi.fn(() => Promise.resolve({ posted_at: POSTED_AT }));
    const { c, stop } = await started(
      deps({ ...OFFICER, fetchRaid: () => Promise.resolve(night()), postReport }),
    );
    await c.night.post(true);
    expect(postReport).toHaveBeenCalledWith(2, '7', true);
    stop();
  });

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
