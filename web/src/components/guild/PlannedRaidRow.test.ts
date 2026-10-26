// web/src/components/guild/PlannedRaidRow.test.ts
// One planned row rendered with svelte/server (raids spec 3.1, 3.6): the desktop cells and the
// phone stack, the guild-time label only when the offsets differ, the count and bar or the bare
// count without a size, and the chips.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { GuildRaidView } from '../../lib/guild/api';
import {
  CHICAGO,
  NEW_YORK,
  NOW,
  between,
  plannedRaid,
  text,
  viewerCharacter,
} from '../../test-support/raid-fixtures';
import PlannedRaidRow from './PlannedRaidRow.svelte';

const HREF = '/guild/us/pvp/olympus-xxvii/raids/7';

function html(raid: GuildRaidView = plannedRaid(), viewerTimezone = NEW_YORK, now = NOW): string {
  return render(PlannedRaidRow, { props: { raid, href: HREF, now, viewerTimezone } }).body;
}

const desktop = (markup: string): string =>
  text(between(markup, 'planned-raid-7-desktop', 'planned-raid-7-phone'));
const phone = (markup: string): string => text(between(markup, 'planned-raid-7-phone', '__end__'));

describe('PlannedRaidRow', () => {
  it('gives the link its date in words, since the date block is hidden from screen readers', () => {
    const markup = html();
    expect(markup).toMatch(/<span class="sr-only">Thursday 10 December, <\/span>/);
    expect(markup).toMatch(/border-line rounded-control[^"]*" aria-hidden="true"/);
  });

  it('is one link to the raid page', () => {
    const markup = html();
    expect(markup.match(/<a /g)?.length).toBe(1);
    expect(markup).toContain(`href="${HREF}"`);
  });

  it('draws the desktop cells in order, in the viewer’s zone', () => {
    const cells = desktop(html());
    expect(cells).toBe(
      "Thu 10 Dec Onyxia's Lair Thursday 10 December, 9:00 pm 8:00 pm CST guild time in 2 days 8 h 23/40 3 tanks · 7 healers · 13 DPS In · DPS ›",
    );
  });

  it('runs the phone stack as three lines beside the date block', () => {
    expect(phone(html())).toBe(
      "Thu 10 Dec Onyxia's Lair In · DPS Thursday 10 December, 9:00 pm · 8:00 pm CST guild time · in 2 days 8 h 23/40 3 tanks · 7 healers · 13 DPS",
    );
  });

  it('drops the guild-time label when the offsets match, and shows the zone under a custom title', () => {
    const cells = desktop(html(plannedRaid({ title: 'Ony farm' }), CHICAGO));
    expect(cells).not.toContain('guild time');
    expect(cells).toContain("Ony farm Onyxia's Lair Thursday 10 December, 8:00 pm");
  });

  it('reads `23 signed up` with no bar while the size is unknown', () => {
    const markup = html(plannedRaid({ size: null }));
    expect(desktop(markup)).toContain('23 signed up');
    expect(markup).not.toContain('bg-gold block h-full');
    expect(html()).toContain('width: 57%');
  });

  it('marks an unsigned viewer with the gold chip and mutes the countdown', () => {
    const markup = html(plannedRaid({ viewer: { characters: [viewerCharacter({ state: null })] } }));
    expect(desktop(markup)).toContain('Not signed up');
    expect(markup).toMatch(/text-muted">in 2 days/);
    expect(html()).toMatch(/text-gold">in 2 days/);
  });

  it('keeps a cancelled row muted with its chip, and derives Ended past the end', () => {
    const cancelled = html(plannedRaid({ status: 'cancelled' }));
    expect(desktop(cancelled)).toContain('Cancelled');
    expect(cancelled).toMatch(/truncate text-\[15px\] font-semibold[^"]*text-muted/);
    const ended = html(plannedRaid(), NEW_YORK, new Date('2026-12-11T06:00:00Z'));
    expect(desktop(ended)).toContain('Ended');
    expect(desktop(html())).not.toContain('Planned');
  });
});
