// web/src/components/guild/GuildRaids.test.ts
// The Raids tab rendered with svelte/server (raids spec 3.1, 4, R1): Plan an event for officers
// only, the read-only gate and the cap line, Planned with its first four rows and Show all, the
// empty lines, Raid nights under it, and nothing planned for a visitor.
import { gateFactsFixture } from '../../test-support/gate-facts';
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { GuildRaidRow, GuildRaidView } from '../../lib/guild/api';
import { attendanceCopy } from '../../lib/guild/attendance-copy';
import { guildGateCopy } from '../../lib/guild/copy';
import { raidCopy } from '../../lib/guild/raid-copy';
import { NEW_YORK, NOW, plannedRaid, text } from '../../test-support/raid-fixtures';
import GuildRaids from './GuildRaids.svelte';
import GuildRaidNights from './GuildRaidNights.svelte';

/** A `disabled` attribute, not the `disabled:` class variants. */
const DISABLED = / disabled(=| |>)/;
const SETTINGS_HREF = '/guild/us/pvp/olympus-xxvii/settings';

interface TabProps {
  planned?: GuildRaidView[];
  officer?: boolean;
  officerTools?: boolean;
  memberView?: boolean;
}

function html({ planned, officer = false, officerTools = true, memberView = true }: TabProps = {}): string {
  return render(GuildRaids, {
    props: {
      rows: [],
      planned,
      zones: [],
      guildTimezone: 'America/Chicago',
      officer,
      officerTools,
      memberView,
      settingsHref: SETTINGS_HREF,
      gateFacts: gateFactsFixture({ settingsHref: SETTINGS_HREF }),
      guildPath: '/guild/us/pvp/olympus-xxvii',
      now: NOW,
      viewerTimezone: NEW_YORK,
    },
  }).body;
}

function weeks(n: number): GuildRaidView[] {
  return Array.from({ length: n }, (_, i) =>
    plannedRaid({
      id: i + 1,
      starts_at: new Date(Date.parse('2026-12-11T02:00:00Z') + i * 7 * 86_400_000).toISOString(),
    }),
  );
}

function planButton(markup: string): string {
  return /<button[^>]*data-testid="guild-raids-plan"[^>]*>/.exec(markup)?.[0] ?? '';
}

describe('GuildRaids planned section', () => {
  it('gives an officer Plan an event in the header row, enabled', () => {
    const markup = html({ planned: weeks(2), officer: true });
    const button = planButton(markup);
    expect(button).not.toBe('');
    expect(button).not.toMatch(DISABLED);
    expect(text(markup)).toMatch(/^Planned 2 Plan an event/);
  });

  it('gives a member no Plan an event', () => {
    expect(planButton(html({ planned: weeks(2) }))).toBe('');
  });

  it('shows the first four rows and Show all for more, sorted by start', () => {
    const markup = html({ planned: weeks(7).reverse() });
    expect(markup.match(/data-testid="planned-raid-\d+"/g)).toEqual(
      [1, 2, 3, 4].map((id) => `data-testid="planned-raid-${id}"`),
    );
    expect(text(markup)).toContain('Show all 7 planned raids');
    expect(text(markup)).toContain(`${raidCopy.subheads.planned} 7`);
    expect(html({ planned: weeks(4) })).not.toContain('guild-planned-show-all');
  });

  it('leaves the count out of the sub-head with none planned', () => {
    const markup = html({ planned: [] });
    expect(markup).not.toContain('text-small text-muted');
    expect(text(markup)).toContain(`${raidCopy.subheads.planned} ${raidCopy.emptyPlanned.member}`);
  });

  it('says the member and the officer empty lines', () => {
    expect(text(html({ planned: [] }))).toContain(raidCopy.emptyPlanned.member);
    expect(text(html({ planned: [], officer: true }))).toContain(raidCopy.emptyPlanned.officer);
  });

  it('draws no Planned rows for a gated officer with nothing planned: the head, the gate, then Raid nights', () => {
    const markup = html({ planned: [], officer: true, officerTools: false });
    expect(markup).not.toContain('data-testid="guild-planned-rows"');
    expect(text(markup)).not.toContain(raidCopy.emptyPlanned.officer);
    const gate = markup.indexOf('guild-raids-plan-gate');
    expect(gate).toBeGreaterThan(markup.indexOf('guild-raids-plan"'));
    expect(markup.indexOf('guild-raid-nights')).toBeGreaterThan(gate);
  });

  it('keeps the Planned rows of a gated officer whose plan lapsed with raids still planned', () => {
    const markup = html({ planned: weeks(1), officer: true, officerTools: false });
    expect(markup).toContain('data-testid="guild-planned-rows"');
  });

  it('puts Raid nights under Planned', () => {
    const markup = html({ planned: weeks(1) });
    expect(markup.indexOf('guild-raid-nights')).toBeGreaterThan(markup.indexOf('guild-planned'));
  });

  it('shows a visitor nothing planned (R1)', () => {
    const markup = html({ planned: undefined, memberView: false });
    expect(markup).not.toContain('guild-planned');
    expect(planButton(markup)).toBe('');
  });

  it('keeps Plan an event disabled under the gate sentence in officer-read-only', () => {
    const markup = html({ planned: weeks(1), officer: true, officerTools: false });
    const button = planButton(markup);
    expect(button).toMatch(DISABLED);
    expect(button).toContain('aria-describedby="officer-tools-gate-raids"');
    expect(markup).toContain('id="officer-tools-gate-raids"');
    const gate = markup.slice(
      markup.indexOf('guild-raids-plan-gate'),
      markup.indexOf('</p>', markup.indexOf('guild-raids-plan-gate')),
    );
    expect(text(gate)).toContain(`${guildGateCopy.lead.raids} ${guildGateCopy.startTrial}`);
    expect(gate).toContain(`href="${SETTINGS_HREF}#billing"`);
  });

  it('keeps Plan an event live with the cap line at 12 raids: the cap holds raids only', () => {
    const markup = html({ planned: weeks(12), officer: true });
    expect(planButton(markup)).not.toMatch(DISABLED);
    expect(planButton(markup)).toContain('aria-describedby="guild-raids-cap"');
    expect(text(markup)).toContain(raidCopy.cap);
    expect(html({ planned: weeks(11), officer: true })).not.toContain('data-testid="guild-raids-cap"');
  });

  it('does not count cancelled raids toward the cap', () => {
    const list = [...weeks(11), plannedRaid({ id: 99, status: 'cancelled' })];
    expect(planButton(html({ planned: list, officer: true }))).not.toMatch(DISABLED);
  });
});

describe('GuildRaids raid nights linked to raids', () => {
  const night = (id: string, raidId: number | null): GuildRaidRow => ({
    id,
    title: "Onyxia's Lair",
    zone: "Onyxia's Lair",
    created_at: '2026-12-11T05:00:00Z',
    duration_ms: 9_660_000,
    fight_count: 4,
    kill_count: 1,
    wipe_count: 3,
    raiders: 23,
    deaths: 14,
    top_hps: null,
    top_dps: null,
    fights: [],
    present: [],
    raid_id: raidId,
  });

  function nights(rows: GuildRaidRow[]): string {
    return render(GuildRaids, {
      props: {
        rows,
        guildPath: '/guild/us/pvp/olympus-xxvii',
        now: NOW,
        viewerTimezone: NEW_YORK,
        gateFacts: gateFactsFixture(),
      },
    }).body;
  }

  function opened(rows: GuildRaidRow[], expanded: string): string {
    return render(GuildRaidNights, {
      props: { rows, bossIds: null, guildPath: '/guild/us/pvp/olympus-xxvii', expanded, onToggle: () => {} },
    }).body;
  }

  function cell(body: string, testid: string): string {
    const at = body.indexOf(`data-testid="${testid}"`);
    expect(at).toBeGreaterThan(-1);
    const start = body.indexOf('>', at) + 1;
    // The cell's own closing tag: skip the spans nested inside it (the character name).
    let depth = 1;
    const tags = /<span\b|<\/span>/g;
    tags.lastIndex = start;
    for (let tag = tags.exec(body); tag !== null; tag = tags.exec(body)) {
      depth += tag[0] === '</span>' ? -1 : 1;
      if (depth === 0) return text(body.slice(start, tag.index));
    }
    return text(body.slice(start));
  }

  it('opens a night in place: its raid page by raid_id, with the id the deep link targets', () => {
    expect(nights([night('abc123def', 7)])).not.toContain('guild-raid-expanded-');
    const body = opened([night('abc123def', 7)], 'abc123def');
    expect(body).toContain('href="/guild/us/pvp/olympus-xxvii/raids/7"');
    expect(text(body)).toContain(attendanceCopy.match.openRaid);
    expect(body).toContain('id="guild-raid-night-abc123def"');
  });

  it('reads the night as Top HPS then Top DPS, never a parse, rounded as the roster writes it', () => {
    const row = {
      ...night('abc123def', null),
      top_hps: { name: 'Rootbound', class: 'druid', metric: 'hps', value: 449.4 },
      top_dps: { name: 'Bow Jackzon', class: 'hunter', metric: 'dps', value: 272.2 },
    };
    const body = nights([row]);
    expect(cell(body, 'guild-raid-top-hps-abc123def')).toBe('Top HPS Rootbound 449');
    expect(cell(body, 'guild-raid-top-dps-abc123def')).toBe('Top DPS Bow Jackzon 272');
    expect(body.indexOf('guild-raid-top-hps-')).toBeLessThan(body.indexOf('guild-raid-top-dps-'));
  });

  it('names only the figures a night has', () => {
    const dpsOnly = nights([
      {
        ...night('abc123def', null),
        top_dps: { name: 'Bow Jackzon', class: 'hunter', metric: 'dps', value: 272 },
      },
    ]);
    expect(cell(dpsOnly, 'guild-raid-top-dps-abc123def')).toBe('Top DPS Bow Jackzon 272');
    expect(dpsOnly).not.toContain('guild-raid-top-hps-');
  });

  it('draws no link for a night matched to no raid', () => {
    expect(nights([night('abc123def', null)])).not.toContain('guild-raid-link-');
  });
});
