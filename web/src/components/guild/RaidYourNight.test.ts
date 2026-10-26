// web/src/components/guild/RaidYourNight.test.ts
// Your night rendered with svelte/server (attendance spec 3.3): the header character first, the
// outcome sentence through the copy module, the muted facts, the loot and its empty line, the
// no-show variant (never the word "no-show"), and the waiting-for-approval row.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { GuildRaidView } from '../../lib/guild/api';
import { yourNightSentence } from '../../lib/guild/night';
import { doneRaid, text, viewerCharacter } from '../../test-support/raid-fixtures';
import { attendanceRow, raidReport } from '../../test-support/report-fixtures';
import RaidYourNight from './RaidYourNight.svelte';

function view(raid: GuildRaidView, currentKey: string | null = null): string {
  return render(RaidYourNight, { props: { raid, currentKey, locale: 'en-US' } }).body;
}

const KRAGGOR = attendanceRow({ own: true });
const MENDA = attendanceRow({
  character_key: 'us-pvp-menda',
  name: 'Menda',
  class: 'priest',
  spec: 'Holy',
  role: 'healer',
  own: true,
});

describe('Your night', () => {
  it('reads the sentence, the death count, the move and the loot for an own character', () => {
    const body = view(doneRaid({ attendance: [KRAGGOR], report: raidReport() }));
    const flat = text(body);
    expect(flat).toContain(yourNightSentence(KRAGGOR));
    expect(flat).toContain('Signed In as DPS. Present from the first pull, 4 of 4 pulls.');
    expect(flat).toContain('2 deaths');
    expect(flat).toContain('DPS 291 on Onyxia, up 53 since Thu 3 Dec');
    expect(body).toContain('data-testid="item-hover-18423"');
  });

  it('draws the rise in the kill colour and a fall without it', () => {
    const rise = view(doneRaid({ attendance: [KRAGGOR], report: raidReport() }));
    expect(rise).toContain('text-kill');
    const base = raidReport();
    const fall = view(
      doneRaid({
        attendance: [KRAGGOR],
        report: raidReport({
          moves: { ...base.moves, rows: [{ ...base.moves.rows[0], before: 300, after: 288 }] },
        }),
      }),
    );
    expect(text(fall)).toContain('down 12');
    expect(fall).not.toContain('text-kill');
  });

  it('says No loot this raid. and shows Excused when set', () => {
    const body = view(
      doneRaid({ attendance: [{ ...KRAGGOR, excused: true }], report: raidReport({ loot: [] }) }),
    );
    expect(text(body)).toContain('No loot this raid.');
    expect(text(body)).toContain('Excused');
  });

  it('puts the header character first', () => {
    const body = view(doneRaid({ attendance: [KRAGGOR, MENDA], report: raidReport() }), 'us-pvp-menda');
    expect(body.indexOf('Menda')).toBeLessThan(body.indexOf('Kraggor'));
  });

  it('explains a missing character without shaming it', () => {
    const missing = attendanceRow({ own: true, outcome: 'no_show', pulls_present: 0 });
    const body = view(doneRaid({ attendance: [missing], report: raidReport() }));
    expect(text(body)).toContain(
      'You are not in this log. If you were there under another name, tell an officer.',
    );
    expect(text(body).toLowerCase()).not.toContain('no-show');
  });

  it('keeps the muted waiting row for a character awaiting approval', () => {
    const raid = doneRaid({
      attendance: [KRAGGOR],
      report: raidReport(),
      viewer: { characters: [viewerCharacter({ name: 'Grimtotem', verified: false, character_key: 'g' })] },
    });
    const body = view(raid);
    expect(body).toContain('data-testid="raid-your-night-waiting"');
    expect(text(body)).toContain('Waiting for approval.');
  });

  it('shows no facts line when there is nothing to say', () => {
    const body = view(
      doneRaid({
        attendance: [KRAGGOR],
        report: raidReport({ deaths: [], moves: { previous: null, rows: [] } }),
      }),
    );
    expect(body).not.toContain('raid-your-night-facts');
  });
});
