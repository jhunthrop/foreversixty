// web/src/components/guild/GuildReadiness.test.ts
// The Readiness tab rendered with svelte/server: officer-read-only (officer tools off) shows the
// plan sentence once and keeps every Nudge on screen but disabled; a member sees the Read-only
// label instead.
import { gateFactsFixture } from '../../test-support/gate-facts';
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { buildMockReadiness } from '../../fixtures/guild/mock-guild';
import type { GuildReadinessRow } from '../../lib/guild/api';
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
  it('says the plan sentence once, under the h2 and above the rows, in officer-read-only', () => {
    const markup = html(true, false);
    expect(markup.split('data-testid="guild-readiness-plan-gate"').length - 1).toBe(1);
    expect(markup).toContain('id="officer-tools-gate-readiness"');
    const gateAt = markup.indexOf('data-testid="guild-readiness-plan-gate"');
    expect(gateAt).toBeGreaterThan(markup.indexOf('<h2'));
    expect(gateAt).toBeLessThan(markup.indexOf('data-testid="guild-readiness-board"'));
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
    expect(text(html(false, false))).toMatch(new RegExp(`ready ${guildGateCopy.readOnly}`));
    expect(text(html(true, false))).not.toContain(guildGateCopy.readOnly);
    expect(text(html(true, true))).not.toContain(guildGateCopy.readOnly);
  });

  it('disables every Nudge, described by the sentence, with no footnote', () => {
    const markup = html(true, false);
    const tags = nudges(markup);
    expect(tags.length).toBeGreaterThan(0);
    for (const tag of tags) {
      expect(tag).toMatch(/\sdisabled(=""|\s|>)/);
      expect(tag).toContain('aria-describedby="officer-tools-gate-readiness"');
    }
    expect(text(markup)).not.toContain('Nudge copies a message you can paste to them.');
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
        guildTimezone: 'UTC',
        viewerTimezone: 'UTC',
      },
    }).body;
  }

  it('shows the toggle pressed and the raid in the panel line', () => {
    const markup = filtered(true);
    expect(markup).toContain('aria-pressed="true"');
    const copy = text(markup).replace(/&#39;/g, "'");
    expect(copy).toContain('Next raid only');
    expect(copy).toContain('ready for Thu 10 Dec In and Late');
  });

  it('lists a raider with no export by exception, and nothing else about them', () => {
    const markup = filtered(true, [rows[1]]);
    expect(text(markup)).toContain('no export yet');
    expect(text(markup)).not.toContain('upgrade');
  });

  it('shows the toggle unpressed, and says so when nobody signed up', () => {
    expect(filtered(false)).toContain('aria-pressed="false"');
    expect(text(filtered(true, []))).toContain('Nobody signed up for this raid yet.');
  });

  it('shows no toggle without a raid', () => {
    expect(html(true, true)).not.toContain('guild-readiness-next-raid');
  });
});

describe('GuildReadiness by exception', () => {
  const base = buildMockReadiness().rows[0];
  const clean: GuildReadinessRow = {
    ...base,
    no_export: false,
    synced_before_cutoff: false,
    gear_gap: { upgrades: 0, gain_dps: null, not_sim_checked: 0, unit: null },
    enchants: { missing_slots: [], checked: true },
    consumables: { state: 'stocked' as const },
    talent_points_unspent: 0,
    talent_points_read: true,
    failing: 0,
  };
  function board(shown: GuildReadinessRow[], raidAsks = false): string {
    return render(GuildReadiness, {
      props: {
        rows: shown,
        officer: true,
        officerTools: true,
        raidGuildTools: true,
        myCharacterKey: null,
        gateFacts: gateFactsFixture({ settingsHref: SETTINGS_HREF }),
        raid: raidAsks
          ? {
              id: 7,
              title: "Onyxia's Lair",
              starts_at: '2026-12-10T20:00:00Z',
              cutoff_at: '2026-12-10T19:00:00Z',
              attunement: true,
            }
          : null,
        nextRaidOnly: raidAsks,
        guildTimezone: 'UTC',
        viewerTimezone: 'UTC',
      },
    }).body;
  }

  it('names every shortfall under the raider, and the clean raiders as ready', () => {
    const short = {
      ...clean,
      character_key: 'us/pvp/short',
      name: 'Shorty',
      enchants: { missing_slots: ['main_hand', 'chest'], checked: true },
      consumables: { state: 'short' as const },
      talent_points_unspent: 1,
      failing: 3,
    };
    const markup = board([clean, short]);
    expect(text(markup)).toContain('2 enchants missing, consumables short, 1 talent point unspent');
    expect(text(markup)).toContain('1 of 2 ready');
    expect(text(markup)).toContain(`${clean.name} Protection · 68 ilvl Ready`);
    expect(markup.match(/data-testid="guild-readiness-nudge-/g)).toHaveLength(1);
  });

  it('reads everyone ready when nobody is short', () => {
    expect(text(board([clean]))).toContain('1 of 1 ready');
    expect(board([clean])).not.toContain('data-testid="guild-readiness-shortfalls"');
  });

  it('reads the raid asks: attunement, flask and food', () => {
    const asks = { ...clean, attuned: false, flask_food: { flask_or_elixir: false, food: true } };
    expect(text(board([asks], true))).toContain('no flask or elixir, not attuned');
  });

  it('puts the leveling characters on their own panel, never graded', () => {
    const low = { ...clean, character_key: 'us/pvp/low', name: 'Lowbie', level: 52 };
    const markup = board([clean, low]);
    expect(text(markup)).toContain('Leveling Lowbie Level 52');
    expect(markup).not.toContain('guild-readiness-nudge-us/pvp/low');
  });
});
