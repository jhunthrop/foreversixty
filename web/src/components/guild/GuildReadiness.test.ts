// web/src/components/guild/GuildReadiness.test.ts
// The Readiness tab rendered with svelte/server: officer-read-only (officer tools off) shows the
// plan sentence once and keeps every Nudge on screen but disabled; a member sees the Read-only
// label instead.
import { gateFactsFixture } from '../../test-support/gate-facts';
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { buildMockReadiness } from '../../fixtures/guild/mock-guild';
import { guildGateCopy } from '../../lib/guild/copy';
import GuildReadiness from './GuildReadiness.svelte';

const SETTINGS_HREF = '/guild/us/hardcore/the-last-watch/settings';

function html(officer: boolean, officerTools: boolean, raidGuildTools = true): string {
  return render(GuildReadiness, {
    props: {
      rows: buildMockReadiness().rows,
      officer,
      officerTools,
      raidGuildTools,
      myCharacterKey: null,
      gateFacts: gateFactsFixture({
        guildTier: officerTools && !raidGuildTools,
        settingsHref: SETTINGS_HREF,
      }),
    },
  }).body;
}

function text(markup: string): string {
  return markup
    .replace(/<!--[\s\S]*?-->/g, '')
    .replace(/<[^>]+>/g, ' ')
    .replace(/&nbsp;/g, ' ')
    .replace(/\s+/g, ' ')
    .trim();
}

function nudges(markup: string): string[] {
  return [...markup.matchAll(/<button[^>]*>/g)]
    .map((m) => m[0])
    .filter((tag) => tag.includes('data-testid="guild-readiness-nudge-'));
}

describe('GuildReadiness', () => {
  it('says the plan sentence once, under the h2 and above the header row, in officer-read-only', () => {
    const markup = html(true, false);
    expect(markup.split('data-testid="guild-readiness-plan-gate"').length - 1).toBe(1);
    expect(markup).toContain('id="officer-tools-gate-readiness"');
    const gateAt = markup.indexOf('data-testid="guild-readiness-plan-gate"');
    expect(gateAt).toBeGreaterThan(markup.indexOf('<h2'));
    expect(gateAt).toBeLessThan(markup.indexOf('data-testid="guild-readiness-header"'));
    const gate = markup.slice(gateAt, markup.indexOf('</p>', gateAt));
    expect(text(gate)).toContain(`${guildGateCopy.lead.readiness} ${guildGateCopy.startTrial}`);
    expect(gate).toContain(`href="${SETTINGS_HREF}?plan=raid_guild#billing"`);
  });

  it('leaves a moderator as before: no sentence, Nudge enabled', () => {
    const markup = html(true, true);
    expect(markup).not.toContain('guild-readiness-plan-gate');
    expect(text(markup)).not.toContain(guildGateCopy.readOnly);
    const nudges = [...markup.matchAll(/<button[^>]*>/g)]
      .map((m) => m[0])
      .filter((tag) => tag.includes('nudge'));
    expect(nudges.length).toBeGreaterThan(0);
    for (const tag of nudges) expect(tag).not.toMatch(/ disabled(=| |>)/);
  });

  it('shows the Read-only label beside the h2 for members only, as Loot does', () => {
    expect(text(html(false, false))).toContain(`Readiness ${guildGateCopy.readOnly}`);
    expect(text(html(true, false))).not.toContain(guildGateCopy.readOnly);
    expect(text(html(true, true))).not.toContain(guildGateCopy.readOnly);
  });

  it('disables every Nudge, described by the sentence, and keeps the footnote', () => {
    const markup = html(true, false);
    const tags = nudges(markup);
    expect(tags.length).toBeGreaterThan(0);
    for (const tag of tags) {
      expect(tag).toMatch(/\sdisabled(=""|\s|>)/);
      expect(tag).toContain('aria-describedby="officer-tools-gate-readiness"');
      expect(tag).toContain('disabled:opacity-50');
      expect(tag).toContain('disabled:cursor-default');
    }
    expect(text(markup)).toContain('Nudge copies a message you can paste to them.');
  });

  it('keeps live Nudge with no sentence while officer tools are on', () => {
    const markup = html(true, true);
    expect(markup).not.toContain('guild-readiness-plan-gate');
    for (const tag of nudges(markup)) {
      expect(tag).not.toMatch(/\sdisabled(=""|\s|>)/);
      expect(tag).not.toContain('aria-describedby');
    }
  });

  it('says the Raid guild sentence, linking to #billing, on a Guild-tier guild', () => {
    const markup = html(true, true, false);
    const gateAt = markup.indexOf('data-testid="guild-readiness-plan-gate"');
    expect(gateAt).toBeGreaterThan(-1);
    const gate = markup.slice(gateAt, markup.indexOf('</p>', gateAt));
    expect(text(gate)).toContain(`${guildGateCopy.lead.readiness} ${guildGateCopy.moveToRaidGuild}`);
    expect(gate).toContain(`href="${SETTINGS_HREF}?upgrade=1#billing"`);
    for (const tag of nudges(markup)) expect(tag).toMatch(/\sdisabled(=""|\s|>)/);
  });

  it('offers no Move to a guild with no plan at all', () => {
    const markup = text(html(true, false, false));
    expect(markup).toContain(guildGateCopy.startTrial);
    expect(markup).not.toContain(guildGateCopy.moveToRaidGuild);
  });

  it('shows a member nothing of the Raid guild gate', () => {
    expect(html(false, true, false)).not.toContain('guild-readiness-plan-gate');
  });

  it('leaves a member exactly as before: no sentence, no Nudge', () => {
    const markup = html(false, false);
    expect(markup).not.toContain('guild-readiness-plan-gate');
    expect(nudges(markup)).toHaveLength(0);
  });
});

describe('GuildReadiness next-raid filter', () => {
  const raid = {
    id: 7,
    title: "Onyxia's Lair",
    starts_at: '2026-12-10T20:00:00Z',
    cutoff_at: '2026-12-03T20:00:00Z',
  };
  const base = buildMockReadiness().rows;
  const rows = [
    { ...base[0], no_export: false, synced_before_cutoff: true, logged_at: '2026-12-01T20:00:00Z' },
    { ...base[1], no_export: true, synced_before_cutoff: null },
  ];

  function filtered(nextRaidOnly: boolean, shown = rows): string {
    return render(GuildReadiness, {
      props: {
        rows: shown,
        officer: true,
        officerTools: true,
        raidGuildTools: true,
        myCharacterKey: null,
        gateFacts: gateFactsFixture({ settingsHref: SETTINGS_HREF }),
        raid,
        nextRaidOnly,
        raidHref: '/guild/x/raids/7',
        guildTimezone: 'UTC',
        viewerTimezone: 'UTC',
      },
    }).body;
  }

  it('shows the toggle pressed, the showing line, stale and no-export rows', () => {
    const markup = filtered(true);
    expect(markup).toContain('aria-pressed="true"');
    const copy = text(markup.replace(/<\/a>/g, '')).replace(/&#39;/g, "'");
    expect(copy).toContain('Next raid \u00b7 Thu 10 Dec');
    expect(copy).toContain("Showing the 2 signed In or Late for Onyxia's Lair, Thu 10 Dec, 8:00 pm.");
    // The page reads the clock it runs on, so only the format is pinned (the words are in readiness-raid.test).
    expect(copy).toMatch(/Synced (today|yesterday|\d days ago|\d{1,2} [A-Z][a-z]{2})/);
    expect(copy).toContain('No export yet');
  });

  it('reads a dash in Gear gap through Item level for a row with no export', () => {
    const markup = filtered(true, [rows[1]]);
    // Five on the desktop row and five on the phone card.
    expect(markup.match(/data-testid="guild-readiness-dash"/g)).toHaveLength(10);
    expect(text(markup)).not.toContain('upgrade');
  });

  it('keeps the space before the Nudge note', () => {
    expect(filtered(true)).toMatch(
      /Sorted by the most checks to fix, then the biggest gear gap\.(<!--[^>]*-->)* Nudge copies/,
    );
  });

  it('keeps the stale pill on one line in a Last synced column wide enough for it', () => {
    const markup = filtered(true);
    const pill = markup.match(/<span class="([^"]*)" data-testid="guild-readiness-stale"/);
    expect(pill?.[1]).toContain('whitespace-nowrap');
    expect(pill?.[1]).not.toContain('whitespace-normal');
    expect(markup).toContain('140px 128px 84px');
  });

  it('draws no column header over the empty state', () => {
    expect(filtered(true, [])).not.toContain('guild-readiness-header');
    expect(filtered(true)).toContain('guild-readiness-header');
  });

  it('prints the raid title without a link when there is no raid page', () => {
    const body = render(GuildReadiness, {
      props: {
        rows,
        officer: true,
        officerTools: true,
        raidGuildTools: true,
        myCharacterKey: null,
        gateFacts: gateFactsFixture({ settingsHref: SETTINGS_HREF }),
        raid,
        nextRaidOnly: true,
        guildTimezone: 'UTC',
        viewerTimezone: 'UTC',
      },
    }).body;
    expect(body).not.toContain('href="/guild/x/raids/7"');
    expect(body).toContain('Onyxia');
  });

  it('shows the toggle unpressed without the line, and says so when nobody signed up', () => {
    expect(filtered(false)).toContain('aria-pressed="false"');
    expect(text(filtered(false))).not.toContain('Showing the');
    expect(text(filtered(true, []))).toContain('Nobody signed up for this raid yet.');
  });

  it('shows no toggle without a raid', () => {
    expect(html(true, true)).not.toContain('guild-readiness-next-raid');
  });

  it('names every missing enchant slot in the pill, wrapped and never clamped (guild loop round 3)', () => {
    const base = buildMockReadiness().rows[0];
    const row = {
      ...base,
      no_export: false,
      enchants: { missing_slots: ['main_hand', 'chest', 'back', 'feet'], checked: true },
    };
    const markup = render(GuildReadiness, {
      props: {
        rows: [row],
        officer: true,
        officerTools: true,
        raidGuildTools: true,
        myCharacterKey: null,
        gateFacts: gateFactsFixture({ settingsHref: SETTINGS_HREF }),
      },
    }).body;
    const at = markup.indexOf('data-testid="guild-readiness-enchants"');
    const pill = markup.slice(markup.lastIndexOf('<span', at), markup.indexOf('</span>', at));
    expect(text(pill)).toContain('Main Hand, Chest, Back, Feet');
    expect(pill).not.toContain('line-clamp');
    expect(pill).not.toContain('overflow:hidden');
    expect(markup).not.toContain('more slots');
  });
});
