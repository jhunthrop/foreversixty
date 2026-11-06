// web/src/components/Rankings.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { rankingsCopy, rankingsEmptyCopy } from '../lib/rankings/copy';
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

  it('titles a board with its encounter, and the index with Rankings and the encounter list', () => {
    const board = render(Rankings, { props: { slug: 'ragnaros' } }).body;
    expect(board).toMatch(/data-testid="page-title"[^>]*>Ragnaros<\/h1>/);
    const index = render(Rankings, { props: {} }).body;
    expect(index).toMatch(/data-testid="page-title"[^>]*>Rankings<\/h1>/);
    expect(index).toContain(rankingsCopy.jobSignedOut);
    expect(index).toContain('Choose an encounter');
    // The index has no bar: it is a list of boards, not a board.
    expect(index).not.toContain('data-testid="rankings-filters"');
  });

  it('prints one Updated stamp at most, and names no sign-in control of its own', () => {
    const { body } = render(Rankings, { props: { slug: 'ragnaros' } });
    expect(body.match(/Updated /g) ?? []).toHaveLength(0);
    expect(body).not.toMatch(/Sign in with/i);
  });

  it('points at the tier list, and closes with the rank colour legend from gold 100', () => {
    for (const props of [{}, { slug: 'ragnaros' }]) {
      const { body } = render(Rankings, { props });
      expect(body).toContain('href="/tiers"');
      expect(body).toContain(rankingsCopy.tierQuestion);
      expect(body).toContain('data-testid="rankings-legend"');
      expect(body).toMatch(/color: var\(--color-parse-gold\);?">100<\/b>/);
    }
  });
});

describe('Rankings empty state copy', () => {
  it('the empty-result copy names what fills it and offers one action', () => {
    // svelte/server never runs $effect, so the fetch never fires and the board never reaches
    // its loaded-empty state here -- this pins the copy module's own contract instead: the
    // exact string and action this component will render once `status === 'ready'` and the
    // board is empty (the runtime path is e2e's job, rankings.spec.ts).
    expect(rankingsEmptyCopy.message).toBe(
      'No ranked fights yet for this filter. Rankings fill in as reports are uploaded.',
    );
    expect(rankingsEmptyCopy.action).toBe('Upload a log');
    expect(rankingsEmptyCopy.href).toBe('/logs');
  });
});
