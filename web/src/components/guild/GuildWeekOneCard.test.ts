// web/src/components/guild/GuildWeekOneCard.test.ts
// The Before raids open card, the zero-log Progression tab and the leveling rows on Readiness, drawn
// with svelte/server: what a guild in week one reads, and what it must never read.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { buildMockProgression, buildMockReadiness } from '../../fixtures/guild/mock-guild';
import type { GuildProgressionPage, GuildReadinessRow, GuildRosterEntry } from '../../lib/guild/api';
import GuildProgression from './GuildProgression.svelte';
import GuildReadiness from './GuildReadiness.svelte';
import GuildWeekOneCard from './GuildWeekOneCard.svelte';
import { gateFactsFixture } from '../../test-support/gate-facts';

const NOW = new Date('2026-11-05T00:00:00Z');
const QUIET_RAID = {
  id: 7,
  zone: 'deadmines',
  zone_name: 'The Deadmines',
  title: 'Deadmines run',
  starts_at: '2026-11-08T02:00:00Z',
  size: 5,
  guild_timezone: 'America/Chicago',
  counts: { in: 2, late: 1, tanks: 1, healers: 1, dps: 1 },
  own: null,
};
const noop = (): void => {};

function text(markup: string): string {
  return markup
    .replace(/<!--[\s\S]*?-->/g, '')
    .replace(/<[^>]+>/g, ' ')
    .replace(/&#39;/g, "'")
    .replace(/&nbsp;/g, ' ')
    .replace(/\s+/g, ' ')
    .trim();
}

function card(overrides: Record<string, unknown> = {}): string {
  const entries: GuildRosterEntry[] = [
    { name: 'A', state: 'approved', rank: 'member', level: 60 },
    { name: 'B', state: 'approved', rank: 'member', level: 52 },
    { name: 'C', state: 'waiting', rank: 'member' },
  ];
  return render(GuildWeekOneCard, {
    props: {
      status: 'ready',
      officer: true,
      waiting: 0,
      entries,
      rows: [{ name: 'B', level: 52, professions: ['alchemy', 'cooking'] }],
      character: { name: 'B', class: 'hunter', level: 52, spec: null },
      nextRaid: null,
      raidHref: null,
      inviteHref: '/guild/settings#invite',
      onRetry: noop,
      onApprove: noop,
      onPlan: noop,
      fixedNow: NOW,
      ...overrides,
    },
  }).body;
}

describe('GuildWeekOneCard', () => {
  it('draws the spread, the raids line, professions and one primary button', () => {
    const markup = card();
    const copy = text(markup);
    expect(copy).toContain('Before raids open');
    expect(copy).toContain('1 AT 60');
    expect(copy).toContain('1 50 TO 59');
    expect(copy).toContain('1 with no level on file yet.');
    expect(copy).toContain("Barrow Deeps, Hyjal Summit and Onyxia's Lair open Wed 9 Dec, in 34 days");
    expect(copy).toContain('Professions');
    expect(copy).toContain('Alchemy: B');
    expect(copy).not.toContain('Cooking');
    expect(copy).toContain('Nothing on the calendar yet.');
    expect(markup.match(/data-testid="guild-week-one-action"/g)).toHaveLength(1);
    expect(markup).toContain('href="/guild/settings#invite"');
  });

  it('leads with the next step and ends with professions below 768, so the button clears the fold', () => {
    const markup = card();
    const step = /<div class="([^"]*)"[^>]*data-testid="guild-week-one-step"/.exec(markup)?.[1] ?? '';
    expect(step).toContain('max-md:order-first');
    expect(step).toContain('md:border-t');
    const professions =
      /<p class="([^"]*)"[^>]*data-testid="guild-week-one-professions"/.exec(markup)?.[1] ?? '';
    expect(professions).toContain('max-md:order-last');
    const action =
      /<button[^>]*class="([^"]*)"[^>]*data-testid="guild-week-one-action"/.exec(markup)?.[1] ?? '';
    expect(action).toContain('max-md:order-first');
  });

  it('holds its height while it loads and offers a retry when the levels fail', () => {
    expect(card({ status: 'loading' })).toContain('data-testid="guild-week-one-loading"');
    const failed = card({ status: 'error' });
    expect(text(failed)).toContain("The guild's levels did not load.");
    expect(text(failed)).toContain('Try again');
  });

  it('swaps the tiles for one line when nobody has a level, and keeps the step', () => {
    const markup = card({ entries: [{ name: 'A', state: 'approved', rank: 'member' }] });
    expect(markup).not.toContain('guild-week-one-tiles');
    expect(text(markup)).toContain("Nobody's level is on file yet.");
    expect(markup).toContain('guild-week-one-action');
  });

  it('tells a raider with no export to import, with no invite link', () => {
    const markup = card({ officer: false, character: null });
    expect(text(markup)).toContain('Import your character');
    expect(markup).not.toContain('guild-week-one-invite');
  });
});

function progression(overrides: Partial<GuildProgressionPage['tier']>): GuildProgressionPage {
  const page = buildMockProgression();
  return { ...page, tier: { ...page.tier, ...overrides } };
}

function progressionHtml(page: GuildProgressionPage, fixedNow: Date = NOW): string {
  return render(GuildProgression, { props: { progression: page, region: 'us', ruleset: 'pvp', fixedNow } })
    .body;
}

describe('GuildProgression at zero pulls', () => {
  it('draws one empty state naming the date, with no bar, no card and no roster bests', () => {
    const markup = progressionHtml(progression({ pulls_total: 0, down: 0 }));
    expect(text(markup)).toContain(
      "Barrow Deeps, Hyjal Summit and Onyxia's Lair open Wed 9 Dec. Progression starts with your first logged pull.",
    );
    expect(text(markup)).toContain('How to log a raid');
    // The empty state sits in the tab's card, not as a bare line under the h2 (guild round 2).
    const cardAt = markup.indexOf('data-testid="guild-progression-empty-card"');
    expect(cardAt).toBeGreaterThan(-1);
    expect(markup.indexOf('data-testid="guild-progression-empty"')).toBeGreaterThan(cardAt);
    expect(markup.split('<h2').length - 1).toBe(1);
    expect(markup).not.toContain('guild-progression-tier');
    expect(markup).not.toContain('guild-progression-encounter');
  });

  it('asks for a log from 9 Dec', () => {
    const markup = progressionHtml(progression({ pulls_total: 0 }), new Date('2026-12-10T00:00:00Z'));
    expect(text(markup)).toContain('No pulls logged yet. Upload your first raid night and it shows here.');
    expect(text(markup)).toContain('Upload a log');
  });
});

describe('GuildProgression with pulls', () => {
  it('reads down of named, and never prints NaN for a field the API omitted', () => {
    const page = buildMockProgression();
    const noKill: GuildProgressionPage = {
      ...page,
      tier: { ...page.tier, down: 0 },
      encounters: page.encounters.map((encounter) => {
        const { first_kill_at: _firstKill, best_kill_ms: _bestKill, ...rest } = encounter;
        return { ...rest, kills: 0 };
      }),
    };
    const copy = text(progressionHtml(noKill));
    expect(copy).toContain('0 of 1 Named bosses down');
    expect(copy).toContain('pulls, no kill yet.');
    expect(copy).not.toContain('NaN');
    expect(copy).not.toContain('We do not track');
    expect(copy).not.toContain('First kill');
  });

  it('shows the kill line with both figures', () => {
    expect(text(progressionHtml(buildMockProgression()))).toMatch(/First kill .* · best kill \d+:\d\d/);
  });

  it('labels the role bests as output, never a parse, rounded as the roster writes them', () => {
    const copy = text(progressionHtml(buildMockProgression()));
    expect(copy).toContain('Top HPS Lightbrand 302');
    expect(copy).toContain('Top DPS Pyrewisp 358');
    expect(copy).not.toMatch(/parse/i);
  });
});

function board(rows: GuildReadinessRow[], officer = true, myCharacterKey: string | null = null): string {
  return render(GuildReadiness, {
    props: {
      rows,
      officer,
      officerTools: true,
      raidGuildTools: true,
      myCharacterKey,
      gateFacts: gateFactsFixture(),
    },
  }).body;
}

describe('GuildReadiness for leveling characters and healers', () => {
  const [sixty] = buildMockReadiness().rows;
  const leveling: GuildReadinessRow = {
    ...sixty,
    character_key: 'us/pvp/low',
    name: 'Lowbie',
    level: 52,
    gear_gap: null,
    enchants: { missing_slots: [], checked: false },
    consumables: { state: 'unknown' },
    talent_points_unspent: 0,
    item_level_delta: null,
    failing: 0,
    nudge_text: undefined,
  };

  it('a level 52 row has no Nudge, no gear figure, and sits under the LEVELING rule', () => {
    const markup = board([leveling, sixty]);
    const copy = text(markup);
    expect(copy).toContain('Level 52 · 8 levels to 60');
    expect(copy).toContain('checks start at 60');
    expect(copy).toContain('LEVELING');
    expect(markup).not.toContain('guild-readiness-nudge-us/pvp/low');
    expect(markup).toContain(`guild-readiness-nudge-${sixty.character_key}`);
    expect(markup.indexOf(`guild-readiness-row-${sixty.character_key}`)).toBeLessThan(
      markup.indexOf('guild-readiness-leveling-group'),
    );
  });

  it('draws no LEVELING rule when only one group exists', () => {
    expect(board([leveling])).not.toContain('guild-readiness-leveling-group');
    expect(board([sixty])).not.toContain('guild-readiness-leveling-group');
  });

  it('a healer row never prints DPS', () => {
    const healer: GuildReadinessRow = {
      ...sixty,
      gear_gap: { upgrades: 4, gain_dps: 31, not_sim_checked: 2, unit: null },
    };
    const copy = text(board([healer]));
    expect(copy).toContain('4 upgrades');
    expect(copy).not.toContain('DPS');
    expect(copy).not.toContain('not simmed yet');
  });

  it('labels the item level against the raid median', () => {
    const copy = text(board([{ ...sixty, item_level: 60, item_level_delta: -1 }]));
    expect(copy).toContain('Avg item level');
    expect(copy).toContain('1 below the raid median');
    expect(copy).toContain('Item level 60 · 1 below median');
    expect(copy).not.toContain('60 (-1)');
  });

  it('tags the member own clean row Ready, and only it', () => {
    const clean: GuildReadinessRow = {
      ...sixty,
      gear_gap: { upgrades: 0, gain_dps: 0, not_sim_checked: 0, unit: 'dps' },
      enchants: { missing_slots: [], checked: true },
      consumables: { state: 'stocked' },
      talent_points_unspent: 0,
      failing: 0,
    };
    expect(board([clean], false, clean.character_key)).toContain('guild-readiness-ready');
    expect(board([clean], false, null)).not.toContain('guild-readiness-ready');
  });

  it('says who to approve with nobody on the board', () => {
    const officerBoard = board([]);
    expect(text(officerBoard)).toContain('No raiders on the board yet. Approve members on the Roster tab.');
    expect(officerBoard).toContain('href="#roster"');
    expect(board([], false)).not.toContain('Open the roster');
  });

  it('draws the raider step as a quiet line with no button when the band signs up (guild round 3)', () => {
    const markup = card({
      officer: false,
      nextRaid: QUIET_RAID,
      raidHref: '/r/7',
      bandSignsUp: true,
    });
    expect(markup).not.toContain('data-testid="guild-week-one-action"');
    expect(markup).toMatch(/class="text-muted text-\[14px\]" data-testid="guild-week-one-sentence"/);
  });
});
