// Raid-ready rendered with svelte/server (web layout spec 2026-11-05 §7.1): the officer rail's
// counts and the signed tanks' fire resistance, never a list of who is behind.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import RaidReady from './RaidReady.svelte';
import type { RaidReadyView } from '../../../lib/guild/raid-ready-view';

const html = (ready: RaidReadyView): string =>
  render(RaidReady, { props: { ready } }).body.replace(/<!--[^>]*-->/g, '');

const ready: RaidReadyView = {
  attuned: { yes: 18, known: 21, unknown: 3 },
  flaskFood: { yes: 12, known: 20, unknown: 4 },
  tanks: [
    { key: 'a', name: 'Thornhide', class: 'druid', fire: 142 },
    { key: 'b', name: 'Kraggor', class: 'warrior', fire: null },
  ],
};

describe('RaidReady', () => {
  it('draws Attuned and Flask and food as X of N with the unshared counted apart, and the tanks', () => {
    const body = html(ready);
    expect(body).toContain('Raid-ready');
    expect(body).toContain('18 of 21');
    expect(body).toContain('3 have not shared their quests');
    expect(body).toContain('12 of 20');
    expect(body).toContain('4 have not shared their bags');
    expect(body).toContain('Fire resistance on tanks');
    expect(body).toContain('142 fire resistance');
    expect(body).toContain('Not shared yet');
    expect(body).toContain('var(--color-class-druid)');
    expect(body).not.toMatch(/—|below floor/i);
  });

  it('leaves out attunement and fire resistance for a raid that asks for neither', () => {
    const body = html({ ...ready, attuned: null, tanks: null });
    expect(body).not.toContain('guild-raid-ready-attuned');
    expect(body).not.toContain('Fire resistance on tanks');
    expect(body).toContain('Flask and food');
  });
});
