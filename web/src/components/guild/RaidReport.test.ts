// web/src/components/guild/RaidReport.test.ts
// The Report and its blocks rendered with svelte/server (attendance spec 3.4, 4, R-A1, R-A3):
// header buttons by rank and webhook state, the two-column and one-column block order, kill-order
// rows, deaths and consumables with Show all, the no-buff line, loot grouped with empty lines by
// rank, and the changes block with and without a previous raid.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { ReportCopy } from '../../lib/guild/report-copy.svelte';
import { between, doneRaid, text } from '../../test-support/raid-fixtures';
import { raidReport } from '../../test-support/report-fixtures';
import type { GuildRaidView, RaidDeathRow, RaidNight, RaidReport as Report } from '../../lib/guild/api';
import RaidReport from './RaidReport.svelte';

const SETTINGS = '/guild/us/pvp/olympus-xxvii#settings';

function view(
  raid: GuildRaidView,
  extra: { officer?: boolean; officerTools?: boolean; copy?: ReportCopy; postLine?: string | null } = {},
  deathsExpanded = false,
): string {
  const report = raid.report as Report;
  return render(RaidReport, {
    props: {
      raid,
      report,
      night: raid.night as RaidNight,
      reportId: raid.match?.report_id as string,
      guildName: 'Olympus XXVII',
      guildPath: '/guild/us/pvp/olympus-xxvii',
      settingsHref: SETTINGS,
      officer: extra.officer ?? true,
      officerTools: extra.officerTools ?? true,
      gateId: 'officer-tools-gate-raids',
      posting: false,
      postLine: extra.postLine ?? null,
      onPost: () => {},
      viewerTimezone: 'America/Chicago',
      locale: 'en-US',
      copy: extra.copy,
      deathsExpanded,
    },
  }).body;
}

const done = (
  report = raidReport(),
  discord = {
    available: true,
    connected: true,
    post_default: false,
    remind_default: true,
    channel_name: null,
  },
) => doneRaid({ report, discord });

describe('the header', () => {
  it('gives an officer Copy and Post when a webhook is connected', () => {
    const body = view(done());
    expect(text(body)).toContain('Copy for Discord');
    expect(body).toContain('data-testid="raid-report-post"');
    expect(body).not.toContain('raid-report-posted');
  });

  it('replaces Post with the posted line once posted, and never offers it twice', () => {
    const body = view(done(raidReport({ posted_at: '2026-12-11T05:58:00Z' })));
    expect(text(body)).toContain('Posted to Discord at 11:58 pm.');
    expect(body).not.toContain('data-testid="raid-report-post"');
  });

  it('says to connect a channel when there is no webhook', () => {
    const body = view(
      done(raidReport(), {
        available: true,
        connected: false,
        post_default: false,
        remind_default: true,
        channel_name: null,
      }),
    );
    expect(text(body)).toContain('Connect a Discord channel in Settings to post raids there.');
    expect(body).toContain(`href="${SETTINGS}"`);
    expect(body).not.toContain('data-testid="raid-report-post"');
  });

  it('disables Post in officer-read-only, names the gate and keeps Copy live', () => {
    const body = view(done(), { officerTools: false });
    expect(body).toMatch(/data-testid="raid-report-post"/);
    expect(body).toMatch(
      /disabled[^>]*aria-describedby="officer-tools-gate-raids"|aria-describedby="officer-tools-gate-raids"[^>]*disabled/,
    );
    expect(text(body)).toContain('Subscribe the guild');
    expect(body).not.toMatch(
      /disabled=""[^>]*data-testid="raid-report-copy"|data-testid="raid-report-copy"[^>]*disabled/,
    );
  });

  it('shows no buttons to a member', () => {
    const body = view(done(), { officer: false });
    expect(body).not.toContain('raid-report-copy');
    expect(body).not.toContain('raid-report-post');
  });

  it('shows Copied and the post failure line', () => {
    const copy = new ReportCopy(() => Promise.resolve());
    copy.state = 'copied';
    expect(text(view(done(), { copy }))).toContain('Copied');
    const body = view(done(), { postLine: 'Could not post. Try again.' });
    expect(body).toContain('data-testid="raid-report-post-error"');
  });

  it('shows the fallback textarea of eight rows on a failed copy', () => {
    const copy = new ReportCopy(() => Promise.reject(new Error('no')));
    copy.state = 'failed';
    copy.text = 'the report';
    const body = view(done(), { copy });
    expect(text(body)).toContain('Could not copy. Select the report and copy it.');
    expect(body).toContain('rows="8"');
    expect(body).toContain('readonly');
  });
});

describe('the block order', () => {
  const body = view(done());

  it('numbers the blocks in the one-column order below 1024', () => {
    const rank = (id: string): number => {
      const found = new RegExp(`order-(\\d)[^"]*"[^>]*>(?:<!--.*?-->)*<div[^>]*data-testid="${id}"`).exec(
        body,
      );
      return found === null ? -1 : Number(found[1]);
    };
    const ids = ['raid-kill-order', 'raid-loot', 'raid-deaths', 'raid-moves', 'raid-consumables'];
    expect(ids.map(rank)).toEqual([1, 2, 3, 4, 5]);
  });

  it('puts kill order and deaths in the left column and loot and changes in the right from 1024', () => {
    const left = body.slice(body.indexOf('lg:min-w-0'), body.indexOf('data-testid="raid-loot"'));
    expect(left).toContain('raid-kill-order');
    expect(left).toContain('raid-deaths');
    expect(left).not.toContain('raid-moves');
    expect(body).toContain('lg:col-span-2');
  });
});

describe('kill order', () => {
  it('links each row to its fight, marks a first pull and a boss never killed', () => {
    const report = raidReport({
      kill_order: [
        {
          encounter_id: 1,
          name: 'Lucifron',
          fight_index: 1,
          killed: true,
          kill_ms: 61_000,
          wipes_before: 0,
          deaths: 0,
        },
        {
          encounter_id: 2,
          name: 'Magmadar',
          fight_index: 3,
          killed: false,
          kill_ms: null,
          wipes_before: 4,
          deaths: 5,
        },
      ],
      unnamed: { pulls: 7, kills: 5, wipes: 2 },
    });
    const body = view(done(report));
    expect(body).toContain('href="/reports/abc123def?fight=1"');
    expect(text(body)).toContain('Lucifron 1:01 first pull');
    expect(body).toContain('text-kill');
    expect(text(body)).toContain('Magmadar Not killed · 4 wipes');
    expect(text(body)).toContain('7 pulls with no named encounter published: 5 kills, 2 wipes.');
  });
});

describe('deaths', () => {
  const rows = (n: number): RaidDeathRow[] =>
    Array.from({ length: n }, (_, i) => ({
      character_key: `k${i}`,
      name: `Raider${i}`,
      class: 'mage',
      on_wipes: 1,
      on_kills: 0,
    }));

  it('shows the first ten and Show all n raiders', () => {
    const body = view(done(raidReport({ deaths: rows(14) })));
    expect(body.match(/data-testid="raid-death-row"/g)).toHaveLength(10);
    expect(text(body)).toContain('Show all 14 raiders');
  });

  it('shows every raider once expanded, and no button', () => {
    const body = view(done(raidReport({ deaths: rows(14) })), {}, true);
    expect(body.match(/data-testid="raid-death-row"/g)).toHaveLength(14);
    expect(body).not.toContain('raid-deaths-show-all');
  });

  it('has no Show all at ten or fewer and says Nobody died when none', () => {
    expect(view(done(raidReport({ deaths: rows(10) })))).not.toContain('raid-deaths-show-all');
    expect(text(view(done(raidReport({ deaths: [] }))))).toContain('Nobody died.');
  });

  it('never colours the bar with the wipe colour and shows the split', () => {
    const body = view(done());
    expect(text(body)).toContain('3 on wipes · 1 on kills');
    expect(body.slice(body.indexOf('raid-deaths'), body.indexOf('raid-loot'))).not.toContain('bg-wipe');
  });
});

describe('consumables', () => {
  it('draws the In the log chips when the log has buffs', () => {
    const body = view(done());
    expect(text(body)).toContain('Short');
    expect(text(body)).toContain('Potions 3');
    expect(text(body)).toContain('No flask');
  });

  it('drops the In the log column and shows its line when the log has no buff data', () => {
    const base = raidReport();
    const body = view(done(raidReport({ consumables: { ...base.consumables, log_has_buffs: false } })));
    expect(text(body)).toContain('This log carries no buff data, so only the readiness check shows.');
    expect(body).not.toContain('Potions 3');
    expect(text(body)).not.toContain('In the log');
    // The three columns stay, so Before keeps its x (UX 14).
    expect(body).toContain('1.5fr');
  });

  it('heads the columns Before and In the log', () => {
    const heads = between(view(done()), 'raid-consumables-heads', 'raid-consumable-row');
    expect(text(heads)).toBe('Before In the log');
  });
});

describe('loot', () => {
  it('groups awards under their boss with a count and shows who awarded', () => {
    const body = view(done());
    expect(text(body)).toContain('Onyxia 1');
    expect(text(body)).toContain('Head of Onyxia');
    expect(text(body)).toContain('by Bow Jackzon');
    expect(body).toContain('data-testid="item-hover-18423"');
  });

  it('tells an officer to award it on the Loot tab, with a link', () => {
    const body = view(done(raidReport({ loot: [] })));
    expect(text(body)).toContain('No loot awarded for this raid. Award it on the Loot tab');
    expect(body).toContain('href="/guild/us/pvp/olympus-xxvii#loot"');
  });

  it('tells a member that none was recorded', () => {
    expect(text(view(done(raidReport({ loot: [] })), { officer: false }))).toContain(
      'No loot recorded for this raid.',
    );
  });
});

describe('biggest changes', () => {
  it('heads the block with the previous raid date and shows the move and its rise', () => {
    const body = view(done());
    expect(text(body)).toContain('Biggest changes since Thu 3 Dec');
    expect(text(body)).toContain('DPS 238 → 291');
    expect(text(body)).toContain('+53');
  });

  it('says it is the first raid logged when there is no previous one', () => {
    const base = raidReport();
    const body = view(done(raidReport({ moves: { ...base.moves, previous: null, rows: [] } })));
    expect(text(body)).toContain("First Onyxia's Lair raid logged. Changes show from the next one.");
  });

  it('never prints the word no-show', () => {
    expect(text(view(done())).toLowerCase()).not.toContain('no-show');
  });
});
