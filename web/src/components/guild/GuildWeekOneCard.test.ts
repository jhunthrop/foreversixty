// web/src/components/guild/GuildWeekOneCard.test.ts
// The Before raids open card, and the Progression tab, drawn
// with svelte/server: what a guild in week one reads, and what it must never read.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { buildMockProgression } from '../../fixtures/guild/mock-guild';
import type { GuildProgressionPage, GuildRosterEntry } from '../../lib/guild/api';
import GuildProgression from './GuildProgression.svelte';
import GuildWeekOneCard from './GuildWeekOneCard.svelte';

const NOW = new Date('2026-11-05T00:00:00Z');
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
  it('draws the countdown, the spread as a Tally, professions and one primary button', () => {
    const markup = card();
    const copy = text(markup);
    expect(copy).toContain('First raids open · Wed 9 Dec');
    expect(markup).toContain('data-region="countdowns"');
    expect(copy).toContain('Before raids open');
    expect(markup).toContain('data-testid="guild-week-one-spread"');
    expect(copy).toContain('At 60 1/2');
    expect(copy).toContain('50 to 59 1/2');
    expect(copy).toContain('1 with no level on file yet.');
    expect(copy).toContain('Professions');
    expect(copy).toContain('Alchemy: B');
    expect(copy).not.toContain('Cooking');
    expect(copy).toContain('Nothing on the calendar yet.');
    expect(markup.match(/data-testid="guild-week-one-action"/g)).toHaveLength(1);
    expect(markup).toContain('href="/guild/settings#invite"');
  });

  it('draws no box, no caps title and no stat tiles: an open Section of the ui primitives', () => {
    const markup = card();
    expect(markup).not.toMatch(/card-surface|rounded-plate|uppercase|tracking-\[/);
    expect(markup).toContain('data-region="section-before-raids-open"');
  });

  it('counts down to nothing once the raids are open, and says what the first night starts', () => {
    const copy = text(card({ fixedNow: new Date('2026-12-15T00:00:00Z') }));
    expect(copy).toContain('Raids open now');
    expect(copy).toContain('Your first raid');
    expect(copy).toContain('Your first logged night starts the progression board.');
  });

  it('holds its height while it loads and offers a retry when the levels fail', () => {
    expect(card({ status: 'loading' })).toContain('data-testid="guild-week-one-loading"');
    const failed = card({ status: 'error' });
    expect(text(failed)).toContain("The guild's levels did not load.");
    expect(text(failed)).toContain('Try again');
  });

  it('swaps the Tally for one line when nobody has a level, and keeps the step', () => {
    const markup = card({ entries: [{ name: 'A', state: 'approved', rank: 'member' }] });
    expect(markup).not.toContain('guild-week-one-spread');
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
        const nights = rest.pulls_by_night.map((n) => ({ ...n, killed: false, kill_ms: null }));
        return { ...rest, kills: 0, pulls_by_night: nights };
      }),
    };
    const copy = text(progressionHtml(noKill));
    expect(copy).toContain('0 of 1 boss down');
    expect(copy).toContain('0 kills');
    expect(copy).not.toContain('NaN');
    expect(copy).not.toContain('We do not track');
    expect(copy).not.toContain('first kill');
  });

  it('heads each boss with its kills, best kill and deaths a pull, and marks the first kill', () => {
    const copy = text(progressionHtml(buildMockProgression()));
    expect(copy).toMatch(/\d+ kills? · best \d+:\d{2}.* · [\d.]+ deaths a pull/);
    expect(copy).toContain('first kill');
  });

  it('labels the role bests as output, never a parse, rounded as the roster writes them', () => {
    const copy = text(progressionHtml(buildMockProgression()));
    expect(copy).toContain('Lightbrand');
    expect(copy).toContain('302 HPS');
    expect(copy).toContain('Pyrewisp');
    expect(copy).toContain('358 DPS');
    expect(copy).not.toMatch(/parse/i);
  });
});
