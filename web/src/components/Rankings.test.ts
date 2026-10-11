// web/src/components/Rankings.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { rankingsEmptyCopy } from '../lib/rankings/copy';
import Rankings from './Rankings.svelte';

describe('Rankings', () => {
  it('renders the band, then the shared character header, and no spine bar or pill', () => {
    const { body } = render(Rankings, { props: { slug: 'ragnaros' } });
    const bandIndex = body.indexOf('data-testid="rankings-band"');
    const headerIndex = body.indexOf('data-testid="character-header');
    expect(bandIndex).toBeGreaterThan(-1);
    expect(headerIndex).toBeGreaterThan(bandIndex);
    expect(body).not.toContain('current-character-bar');
    expect(body).not.toMatch(/class="[^"]*\bpill\b/);
  });

  it('prints one Updated stamp at most, and names no sign-in control of its own', () => {
    const { body } = render(Rankings, { props: { slug: 'ragnaros' } });
    expect(body.match(/Updated /g) ?? []).toHaveLength(0);
    expect(body).not.toMatch(/Sign in with/i);
  });

  it('heads the page with Rankings and points at the tier list', () => {
    const { body } = render(Rankings, { props: {} });
    expect(body).toContain('>Rankings</h1>');
    expect(body).toContain('href="/tiers"');
  });
});

describe('Rankings empty state copy', () => {
  it('the empty-result copy names what fills it and offers one action, through EmptyState', () => {
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
