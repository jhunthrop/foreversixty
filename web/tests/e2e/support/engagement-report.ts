// web/tests/e2e/support/engagement-report.ts
// A sample owner's report for /admin/activity (GET /v1/admin/engagement): 28 days, 12 weeks and 8
// signup weeks, live since 12 Nov so the older figures are floors. tests/e2e/admin-activity.spec.ts
// and tests/captures/admin-activity share it. Review data only.
import type { EngagementPeriod, EngagementReport } from '../../../src/lib/engagement/api';

const LIVE_SINCE = '2026-11-12';

/** n periods back from start, stepping days days, partial before LIVE_SINCE. */
function periods(start: string, n: number, days: number): EngagementPeriod[] {
  return Array.from({ length: n }, (_, i) => {
    const d = new Date(`${start}T00:00:00Z`);
    d.setUTCDate(d.getUTCDate() - i * days);
    const day = d.toISOString().slice(0, 10);
    const active = Math.max(0, 30 - i * 2);
    const fresh = Math.max(0, 8 - i);
    return {
      start: day,
      active,
      new: fresh,
      returning: active - fresh,
      site: Math.max(0, active - 3),
      addon: Math.round(active / 2),
      app: i % 3,
      partial: day < LIVE_SINCE,
    };
  });
}

export const SAMPLE_REPORT: EngagementReport = {
  today: '2026-11-13',
  live_since: LIVE_SINCE,
  internal_accounts: 6,
  days: periods('2026-11-13', 28, 1),
  weeks: periods('2026-11-09', 12, 7),
  cohorts: [
    '2026-09-21',
    '2026-09-28',
    '2026-10-05',
    '2026-10-12',
    '2026-10-19',
    '2026-10-26',
    '2026-11-02',
    '2026-11-09',
  ].map((week, c) => ({
    week,
    joined: 10 + c * 3,
    weeks: Array.from({ length: 7 - c }, (_, k) => ({ active: Math.max(1, 6 - k), partial: k < 6 - c })),
  })),
};
