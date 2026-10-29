// web/src/components/guild/RaidDoneBand.test.ts
// The band's log lines rendered with svelte/server (attendance spec 3.1): the night line with its
// zero-clause rule and bar, the match line on a done raid, `Ended · no log yet` with its hint
// inside and after the 6 h window, `Match a log` for officers (disabled read-only), and a seeded
// open picker.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { GuildRaidView, MatchCandidates } from '../../lib/guild/api';
import { attendanceCopy, cameRatio, nightLine } from '../../lib/guild/attendance-copy';
import { MatchPicker } from '../../lib/guild/match-picker.svelte';
import { CHICAGO, doneRaid, plannedRaid, raidNight, text } from '../../test-support/raid-fixtures';
import RaidDoneBand from './RaidDoneBand.svelte';

const DISABLED = / disabled(=| |>)/;
/** Fri 11 Dec 2026, 06:00 UTC: the Thursday raid ended at 05:00 UTC. */
const JUST_ENDED = new Date('2026-12-11T06:00:00Z');
const LONG_AFTER = new Date('2026-12-11T12:00:00Z');

interface Options {
  now?: Date;
  officer?: boolean;
  readOnly?: boolean;
  picker?: MatchPicker | null;
}

function html(
  raid: GuildRaidView,
  { now = JUST_ENDED, officer = false, readOnly = false, picker = null }: Options = {},
): string {
  return render(RaidDoneBand, {
    props: {
      raid,
      now,
      guildPath: '/guild/us/pvp/olympus-xxvii',
      viewerTimezone: CHICAGO,
      guildTimezone: CHICAGO,
      locale: 'en-US',
      officer,
      readOnly,
      gateId: 'gate',
      loadCandidates: () => new Promise<MatchCandidates>(() => {}),
      onMatch: () => Promise.resolve({ ok: true as const }),
      onClear: () => Promise.resolve({ ok: true as const }),
      picker,
    },
  }).body;
}

describe('the night line', () => {
  it('reads the ratio and the facts, figures in tabular numerals', () => {
    const body = html(doneRaid());
    expect(text(between(body, 'raid-night-line', 'raid-match-line'))).toBe(
      '22 of 23 came · 2 joined · 1 kill · 3 wipes · 14 deaths · 2 h 41 min',
    );
    expect(body).toContain('numeric text-[15px] whitespace-nowrap">22 of 23 came');
    expect(body).toMatch(/<span class="numeric">41<\/span>(<!--.*?-->)* min/);
  });

  it('leaves a zero clause out and fills the bar by came over signed', () => {
    const body = html(
      doneRaid({ night: raidNight({ came: 10, signed: 20, joined: 0, kills: 0, wipes: 0, deaths: 0 }) }),
    );
    expect(text(between(body, 'raid-night-line', 'raid-match-line'))).toBe(
      `${cameRatio({ came: 10, signed: 20 })} · 2 h 41 min`,
    );
    expect(body).toContain('width: 50%');
  });

  it('draws no bar before anyone signed, and reads like nightLine', () => {
    const night = raidNight({ signed: 0, came: 3 });
    const body = html(doneRaid({ night }));
    expect(body).not.toContain('absolute inset-y-0');
    expect(text(between(body, 'raid-night-line', 'raid-match-line'))).toBe(nightLine(night));
  });

  it('is absent without a night', () => {
    expect(html(doneRaid({ night: null }))).not.toContain('raid-night-line');
  });
});

describe('the match and ended lines', () => {
  it('shows the match line, and no ended line, on a done raid', () => {
    const body = html(doneRaid());
    expect(body).toContain('raid-match-line');
    expect(body).not.toContain('raid-ended-line');
  });

  it('says Ended · no log yet, promising the match within 6 h of the end', () => {
    const body = html(plannedRaid());
    expect(text(between(body, 'raid-ended-line', 'never'))).toBe(
      `${attendanceCopy.ended.noLog}${attendanceCopy.ended.matchSoon}`,
    );
    expect(body).not.toContain('raid-match-open');
  });

  it('asks for the upload after the window, linking Upload the log', () => {
    const body = html(plannedRaid(), { now: LONG_AFTER });
    expect(text(between(body, 'raid-ended-line', 'never')).replace(' ,', ',')).toBe(
      'Ended · no log yet Upload the log, then match it here.',
    );
    expect(body).toMatch(/<a href="\/upload"[^>]*>Upload the log<\/a>/);
  });

  it('draws nothing for a raid that has not ended with no log, or a cancelled one', () => {
    expect(html(plannedRaid(), { now: new Date('2026-12-08T18:00:00Z') })).not.toContain('raid-ended-line');
    expect(html(plannedRaid({ status: 'cancelled' }))).not.toContain('raid-ended-line');
  });

  it('gives officers Match a log, disabled read-only with the gate id', () => {
    const on = html(plannedRaid(), { officer: true });
    expect(text(on)).toContain(attendanceCopy.ended.matchALog);
    expect(/<button[^>]*data-testid="raid-match-open"/.exec(on)?.[0]).not.toMatch(DISABLED);
    const off = /<button[^>]*data-testid="raid-match-open"/.exec(
      html(plannedRaid(), { officer: true, readOnly: true }),
    )?.[0];
    expect(off).toMatch(DISABLED);
    expect(off).toContain('aria-describedby="gate"');
  });

  it('opens the picker in place under the row for an officer, and never for a member', () => {
    const picker = new MatchPicker(
      { load: () => new Promise(() => {}), match: () => Promise.resolve({ ok: true }) },
      { load: 'ready' },
    );
    expect(html(plannedRaid(), { officer: true, picker })).toContain('raid-match-picker');
    expect(html(plannedRaid(), { picker })).not.toContain('raid-match-picker');
    expect(html(plannedRaid(), { officer: true, picker })).not.toContain('raid-match-open');
  });
});

/** The markup inside the element carrying `testid`, up to the sibling marker `until`. */
function between(markup: string, testid: string, until: string): string {
  const marker = markup.indexOf(`data-testid="${testid}"`);
  if (marker < 0) throw new Error(`no ${testid}`);
  const from = markup.indexOf('>', marker) + 1;
  const to = markup.indexOf(`data-testid="${until}"`, from);
  return markup.slice(from, to < 0 ? undefined : markup.lastIndexOf('<', to));
}
