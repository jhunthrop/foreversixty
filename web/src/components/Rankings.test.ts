// web/src/components/Rankings.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { encounterPickerCopy, rankingsEmptyCopy } from '../lib/rankings/copy';
import Rankings from './Rankings.svelte';

describe('Rankings', () => {
  it('a board is the PageHeader, then the ControlBar; no shared band, spine bar or pill', () => {
    const { body } = render(Rankings, { props: { slug: 'ragnaros' } });
    const titleIndex = body.indexOf('data-testid="page-title"');
    const barIndex = body.indexOf('data-testid="rankings-filters"');
    expect(titleIndex).toBeGreaterThan(-1);
    expect(barIndex).toBeGreaterThan(titleIndex);
    expect(body).toContain('aria-label="Ranking choices"');
    expect(body).not.toContain('data-testid="rankings-band"');
    expect(body).not.toContain('data-testid="character-header');
    expect(body).not.toContain('current-character-bar');
    expect(body).not.toMatch(/class="[^"]*\bpill\b/);
  });

  it('titles every view Rankings: the encounter is the bar choice, the index the encounter list', () => {
    const board = render(Rankings, { props: { slug: 'ragnaros' } }).body;
    expect(board).toMatch(/data-testid="page-title"[^>]*>Rankings<\/h1>/);
    const index = render(Rankings, { props: {} }).body;
    expect(index).toMatch(/data-testid="page-title"[^>]*>Rankings<\/h1>/);
    expect(index).toContain(encounterPickerCopy.heading);
    // The index's bar holds the board tabs alone: it is a list of boards, not a board.
    expect(index).toContain('data-testid="control-board-guild"');
    expect(index).not.toContain('data-testid="control-encounter"');
  });

  it('sets the board in the page lead panel and the rail beside it, the rail first on a phone', () => {
    const { body } = render(Rankings, { props: { slug: 'ragnaros' } });
    expect(body).toContain('data-panels="true"');
    expect(body).toMatch(/class="ui-page-rail[^"]*"[^>]*data-first="true"/);
  });

  it('prints one Updated stamp at most, and names no sign-in control of its own', () => {
    const { body } = render(Rankings, { props: { slug: 'ragnaros' } });
    // No date is printed before the board answers; the stamp follows the tier link, so it moves nothing.
    expect(body).not.toMatch(/Updated /);
    expect(body).not.toMatch(/data-testid="rankings-updated"/);
    expect(body).not.toMatch(/Sign in with/i);
  });

  it('points at the tier list from the band on every view', () => {
    for (const props of [{}, { slug: 'ragnaros' }]) {
      const { body } = render(Rankings, { props });
      expect(body).toMatch(/href="\/tiers"[^>]*data-testid="rankings-tier-link"/);
    }
  });
});

describe('Rankings empty state copy', () => {
  it('names what is missing, then what fills it, then one action', () => {
    expect(rankingsEmptyCopy.board('Fury Warrior', 'Onyxia')).toBe('No Fury Warrior kills on Onyxia yet.');
    expect(rankingsEmptyCopy.guilds('Onyxia')).toBe('No guild has a ranked kill on Onyxia yet.');
    expect(rankingsEmptyCopy.action).toBe('Upload a log');
    expect(rankingsEmptyCopy.href).toBe('/logs');
    expect(encounterPickerCopy.where("Onyxia's Lair", 40)).toBe("Onyxia's Lair · 40-player");
    expect(encounterPickerCopy.where(null, 5)).toBe('5-player');
    expect(encounterPickerCopy.where(null, null)).toBe('');
    expect(encounterPickerCopy.ranked(1212)).toBe('1,212 kills');
    expect(encounterPickerCopy.ranked(1)).toBe('1 kill');
  });
});
