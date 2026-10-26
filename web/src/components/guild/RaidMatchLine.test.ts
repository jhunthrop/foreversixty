// web/src/components/guild/RaidMatchLine.test.ts
// The match line rendered with svelte/server (attendance spec 3.1): the two lines, the night link,
// guild time after the upload time when the offsets differ, `matched by an officer`, `Not this log`
// for officers only (disabled read-only, with the gate sentence's id), and the open confirm.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { attendanceCopy } from '../../lib/guild/attendance-copy';
import { CHICAGO, NEW_YORK, raidMatch, text } from '../../test-support/raid-fixtures';
import RaidMatchLine from './RaidMatchLine.svelte';

/** A `disabled` attribute, not the `disabled:` class variants. */
const DISABLED = / disabled(=| |>)/;
const GUILD_PATH = '/guild/us/pvp/olympus-xxvii';

interface Options {
  officer?: boolean;
  readOnly?: boolean;
  confirming?: boolean;
  viewer?: string;
  how?: 'auto' | 'hand';
}

function html({
  officer = false,
  readOnly = false,
  confirming = false,
  viewer = CHICAGO,
  how = 'auto',
}: Options = {}): string {
  return render(RaidMatchLine, {
    props: {
      match: raidMatch({ how }),
      guildPath: GUILD_PATH,
      viewerTimezone: viewer,
      guildTimezone: CHICAGO,
      locale: 'en-US',
      officer,
      readOnly,
      gateId: 'gate',
      onClear: () => Promise.resolve({ ok: true as const }),
      confirming,
    },
  }).body;
}

describe('RaidMatchLine', () => {
  it('reads line 1 with the link to the night, expanded', () => {
    const body = html();
    expect(text(body)).toContain(`${attendanceCopy.match.loggedAsNight} ${attendanceCopy.match.openNight}`);
    expect(body).toContain(`href="${GUILD_PATH}#raids-abc123def"`);
    expect(body).toMatch(/class="text-gold font-bold"/);
  });

  it('reads line 2 with the uploader and `matched by time`, no guild time in the same zone', () => {
    expect(text(html())).toContain('Log uploaded Thu 10 Dec, 11:42 pm by Kraggor · matched by time');
    expect(text(html())).not.toContain('guild time');
  });

  it('adds the guild-time suffix when the offsets differ', () => {
    expect(text(html({ viewer: NEW_YORK }))).toContain(
      'Log uploaded Fri 11 Dec, 12:42 am (11:42 pm guild time) by Kraggor · matched by time',
    );
  });

  it('says a hand match was matched by an officer', () => {
    expect(text(html({ how: 'hand' }))).toContain('· matched by an officer');
  });

  it('gives Not this log to officers only', () => {
    expect(html()).not.toContain('raid-not-this-log');
    const officer = html({ officer: true });
    expect(text(officer)).toContain(attendanceCopy.match.notThisLog);
    expect(/<button[^>]*data-testid="raid-not-this-log"/.exec(officer)?.[0]).not.toMatch(DISABLED);
  });

  it('disables it for an officer-read-only guild, described by the gate sentence', () => {
    const body = html({ officer: true, readOnly: true });
    const tag = /<button[^>]*data-testid="raid-not-this-log"/.exec(body)?.[0];
    expect(tag).toMatch(DISABLED);
    expect(tag).toContain('aria-describedby="gate"');
  });

  it('swaps Not this log for the confirm when it is open', () => {
    const body = html({ officer: true, confirming: true });
    expect(body).not.toContain('raid-not-this-log');
    expect(text(body)).toContain(attendanceCopy.clearConfirm.prompt);
  });
});
