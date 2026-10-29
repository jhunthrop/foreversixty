// web/src/components/guild/RaidMatchPicker.test.ts
// The picker rendered with svelte/server from a seeded MatchPicker (attendance spec 3.2, 4): the
// helper in its two forms, loading, error, empty, the radiogroup with `Closest` preselected, the
// matched-elsewhere row (disabled, its sentence, the title linked), Match this log's states and the
// status line. The arrow-key model is in match-picker.test.ts; real focus movement needs a browser and
// no spec covers it yet (the attendance e2e, Task 10, owns it).
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { MatchCandidate } from '../../lib/guild/attendance-contract';
import { attendanceCopy, pickerEmpty, pickerHelper } from '../../lib/guild/attendance-copy';
import { MatchPicker, type MatchPickerSeed } from '../../lib/guild/match-picker.svelte';
import { CHICAGO, NEW_YORK, text } from '../../test-support/raid-fixtures';
import RaidMatchPicker from './RaidMatchPicker.svelte';

const DISABLED = / disabled(=| |>)/;
const ZONE = "Onyxia's Lair";
const GUILD_PATH = '/guild/us/pvp/olympus-xxvii';

function row(id: string, extra: Partial<MatchCandidate> = {}): MatchCandidate {
  return {
    report_id: id,
    zone_name: ZONE,
    uploaded_at: '2026-12-11T05:42:00Z',
    uploaded_by_name: 'Kraggor',
    pulls: 9,
    kills: 1,
    duration_ms: 9_660_000,
    closest: false,
    matched_raid: null,
    ...extra,
  };
}

function html(seed: Partial<MatchPickerSeed> = {}, viewer = CHICAGO): string {
  const picker = new MatchPicker(
    { load: () => new Promise(() => {}), match: () => Promise.resolve({ ok: true }) },
    { load: 'ready', ...seed },
  );
  return render(RaidMatchPicker, {
    props: {
      picker,
      zoneName: ZONE,
      guildPath: GUILD_PATH,
      viewerTimezone: viewer,
      guildTimezone: CHICAGO,
      locale: 'en-US',
      onMatched: () => {},
      onCancel: () => {},
    },
  }).body;
}

describe('RaidMatchPicker', () => {
  it('heads the panel and says what the list is, in both helper forms', () => {
    const known = html({ rows: [row('a')] });
    expect(text(known)).toContain(`${attendanceCopy.picker.head} ${pickerHelper(ZONE, true)}`);
    expect(text(html({ rows: [row('a')], zoneKnown: false }))).toContain(pickerHelper(ZONE, false));
  });

  it('draws a focus target in every state: the heading, tabindex -1', () => {
    const seeds: Partial<MatchPickerSeed>[] = [
      { load: 'loading' },
      { load: 'failed' },
      { rows: [] },
      { rows: [row('a')] },
    ];
    for (const seed of seeds) {
      expect(html(seed)).toMatch(/<h2[^>]*tabindex="-1"[^>]*data-testid="raid-match-heading"/);
    }
  });

  it('keeps each phone row to three lines: the zone, the upload time, then who, the figures and the duration', () => {
    const body = html({ rows: [row('a')] });
    expect(text(body)).toContain(
      `${ZONE} Uploaded Thu 10 Dec, 11:42 pm by Kraggor · 9 pulls · 1 kill · 2 h 41 min`,
    );
    expect(body).toMatch(/numeric whitespace-nowrap">(<!--[^>]*-->)*Thu 10 Dec, 11:42 pm/);
    expect(body).toContain('text-strong text-[14px] font-semibold whitespace-nowrap');
  });

  it('drops the helper beside the empty line, which says the same', () => {
    const body = html({ rows: [] });
    expect(body).toContain('raid-match-empty');
    expect(body).not.toContain('raid-match-helper');
  });

  it('draws three skeleton rows while loading, and no Match this log', () => {
    const body = html({ load: 'loading' });
    expect(body).toContain('raid-match-skeleton');
    expect(body).not.toContain('raid-match-submit');
    expect(body).toContain('raid-match-cancel');
  });

  it('says the logs did not load, with Try again', () => {
    const body = html({ load: 'failed' });
    expect(text(body)).toContain(attendanceCopy.picker.loadFailed);
    expect(body).toContain('raid-match-load-error-retry');
  });

  it('says there are no logs, with Cancel only', () => {
    const body = html({ rows: [] });
    expect(text(body)).toContain(pickerEmpty(ZONE));
    expect(body).not.toContain('raid-match-submit');
    expect(body).toContain('raid-match-cancel');
  });

  it('draws one radiogroup, Closest preselected and tabbable, with the row figures', () => {
    const body = html({
      rows: [row('a', { closest: true }), row('b', { pulls: 1, kills: 0 })],
      selected: 'a',
    });
    expect(body).toContain(`role="radiogroup" aria-label="${attendanceCopy.picker.radiogroupLabel}"`);
    expect(body).toMatch(
      /role="radio" aria-checked="true"[^>]*tabindex="0"[^>]*data-testid="raid-match-row-a"/,
    );
    expect(body).toMatch(
      /role="radio" aria-checked="false"[^>]*tabindex="-1"[^>]*data-testid="raid-match-row-b"/,
    );
    expect(body.match(/raid-match-closest/g)).toHaveLength(1);
    const first = text(body.slice(body.indexOf('raid-match-row-a'), body.indexOf('raid-match-row-b')));
    expect(first).toContain(
      `${ZONE} ${attendanceCopy.picker.closest} Uploaded Thu 10 Dec, 11:42 pm by Kraggor`,
    );
    // The desktop cells, after the phone lines.
    expect(first).toContain(
      'by Kraggor · 9 pulls · 1 kill · 2 h 41 min Uploaded Thu 10 Dec, 11:42 pm 9 pulls · 1 kill 2 h 41 min',
    );
    expect(text(body)).toContain('1 pull · 0 kills');
  });

  it('puts the guild time muted beside the upload time when the offsets differ', () => {
    const body = html({ rows: [row('a')] }, NEW_YORK);
    expect(text(body)).toContain('Fri 11 Dec, 12:42 am (11:42 pm guild time)');
  });

  it('draws a log another raid holds disabled, with its sentence and the title linked', () => {
    const matched = row('a', {
      matched_raid: { id: 3, title: 'Ony farm', starts_at: '2026-12-04T02:00:00Z' },
    });
    const body = html({ rows: [matched, row('b', { closest: true })], selected: 'b' });
    expect(body).toMatch(/aria-disabled="true"[^>]*data-testid="raid-match-row-a"/);
    // The sentence and its link sit outside the radio element.
    expect(between(body, 'raid-match-row-a')).not.toContain('<a ');
    expect(body).toMatch(/aria-describedby="raid-match-elsewhere-a"/);
    expect(body).toContain(`href="${GUILD_PATH}/raids/3"`);
    expect(text(between(body, 'raid-match-elsewhere'))).toBe(
      'Matched to Ony farm · Thu 3 Dec. Clear it there first.',
    );
    expect(body).not.toMatch(/aria-disabled="true"[^>]*data-testid="raid-match-row-b"/);
  });

  it('enables Match this log with a selection and disables it without one', () => {
    const on = html({ rows: [row('a', { closest: true })], selected: 'a' });
    expect(/<button[^>]*data-testid="raid-match-submit"/.exec(on)?.[0]).not.toMatch(DISABLED);
    expect(text(on)).toContain(attendanceCopy.picker.match);
    const off = html({ rows: [row('a')] });
    expect(/<button[^>]*data-testid="raid-match-submit"/.exec(off)?.[0]).toMatch(DISABLED);
  });

  it('says Matching… in flight', () => {
    const body = html({ rows: [row('a')], selected: 'a', busy: true });
    expect(text(body)).toContain(attendanceCopy.picker.matching);
    expect(/<button[^>]*data-testid="raid-match-submit"/.exec(body)?.[0]).toMatch(DISABLED);
  });

  it('shows the status line as an alert, focusable', () => {
    const body = html({ rows: [row('a')], line: attendanceCopy.picker.matchedElsewhere });
    expect(body).toMatch(/role="alert" tabindex="-1"[^>]*data-testid="raid-match-status"/);
    expect(text(between(body, 'raid-match-status'))).toBe(attendanceCopy.picker.matchedElsewhere);
  });
});

/** The markup inside the span or paragraph carrying `testid`. */
function between(markup: string, testid: string): string {
  const marker = markup.indexOf(`data-testid="${testid}"`);
  if (marker < 0) throw new Error(`no ${testid}`);
  const from = markup.indexOf('>', marker) + 1;
  return markup.slice(from).split(/<\/(?:span|p)>/)[0];
}
