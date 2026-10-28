import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { coachModelReview, coachTeaser } from '../../test-support/coach-fixtures';
import { rowLineFor } from '../../lib/coach/row-line';
import type { StripPhase } from '../../lib/coach/row-line';
import CoachRowStrip from './CoachRowStrip.svelte';

const props = (answer: Parameters<typeof rowLineFor>[0], phase: StripPhase = 'line') => ({
  phase,
  line: rowLineFor(answer),
  href: '?fight=3&tab=coach&source=P-1',
  iconUrl: () => '/data/b/icons/x.webp',
  label: 'label',
  onOpen: () => {},
  onRetry: () => {},
});

describe('CoachRowStrip', () => {
  it('Premium owner: Fix first, one link to the coach tab', () => {
    const { body } = render(CoachRowStrip, { props: props(coachModelReview()) });
    expect(body).toContain('Fix first ·');
    expect(body).toContain('Bloodthirst: 9 casts lost');
    expect(body).toContain('href="?fight=3&amp;tab=coach&amp;source=P-1"');
    expect(body).not.toContain('/premium');
  });
  it('free owner: count, first, and a Premium link to /premium', () => {
    const { body } = render(CoachRowStrip, { props: props(coachTeaser()) });
    expect(body).toContain('3 things to fix · First:');
    expect(body).toContain('Bloodthirst: 9 casts lost');
    expect(body).toContain('href="/premium"');
    expect(body).toContain('Read the full review with Premium');
  });
  it('loading holds the strip height and says so to screen readers', () => {
    const { body } = render(CoachRowStrip, { props: props(null, 'loading') });
    expect(body).toContain('min-h-[88px]');
    expect(body).toContain('Reading the pull…');
  });
  it('error offers Try again', () => {
    const { body } = render(CoachRowStrip, { props: props(null, 'error') });
    expect(body).toContain('Try again');
    expect(body).toContain('just now.');
  });
  it('nothing for a null line', () => {
    const { body } = render(CoachRowStrip, { props: props(null) });
    expect(body).not.toContain('coach-row-line');
  });
});
