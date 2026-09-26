// web/src/components/Rankings.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { rankingsEmptyCopy } from '../lib/rankings/copy';
import Rankings from './Rankings.svelte';

describe('Rankings loading state', () => {
  it('reserves height with a Skeleton instead of a bare "Loading rankings." line', () => {
    const { body } = render(Rankings, { props: { slug: 'warden-kelthas' } });
    expect(body).toContain('data-testid="rankings-skeleton"');
    expect(body).not.toContain('Loading rankings.');
  });
});

describe('Rankings empty state copy', () => {
  it('the empty-result copy names what fills it and offers one action, through EmptyState', () => {
    const { body } = render(Rankings, { props: {} });
    // svelte/server never runs $effect, so the fetch never fires and the board never reaches
    // its loaded-empty state here -- this pins the copy module's own contract instead: the
    // exact string and action this component will render once `status === 'ready'` and the
    // board is empty (the runtime path is e2e's job, rankings-encounter-picker.spec.ts's own
    // sibling coverage plus rankings.spec.ts).
    expect(rankingsEmptyCopy.message).toBe(
      'No ranked fights yet for this filter. Rankings fill in as reports are uploaded.',
    );
    expect(rankingsEmptyCopy.action).toBe('Upload a log');
    expect(rankingsEmptyCopy.href).toBe('/logs');
    void body;
  });
});
