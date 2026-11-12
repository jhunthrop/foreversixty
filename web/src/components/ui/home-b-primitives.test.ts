// Home B's primitives (spec 2026-11-07): Countdown, FeedRow, GradeTile.
import { render } from 'svelte/server';
import { createRawSnippet } from 'svelte';
import { describe, expect, it } from 'vitest';
import Countdown from './Countdown.svelte';
import FeedRow from './FeedRow.svelte';
import GradeTile from './GradeTile.svelte';

const text = (value: string) => createRawSnippet(() => ({ render: () => `<span>${value}</span>` }));
const html = (body: string) => body.replace(/<!--[^>]*-->/g, '');

describe('Countdown', () => {
  it('reads its clock aloud as a timer and draws the parts', () => {
    const body = html(
      render(Countdown, {
        props: {
          eyebrow: 'First raids open',
          target: '2026-12-09T00:00:00Z',
          now: new Date('2026-11-06T18:00:00Z'),
          live: true,
          children: text('raids'),
        },
      }).body,
    );
    expect(body).toContain('role="timer"');
    expect(body).toContain('aria-label="32 days, 6 hours"');
    expect(body).toContain('ui-countdown-live');
    expect(body).toContain('<span>raids</span>');
  });

  it('shows only its content with no instant', () => {
    const body = html(render(Countdown, { props: { eyebrow: 'Find a guild', target: null } }).body);
    expect(body).not.toContain('role="timer"');
  });
});

describe('FeedRow', () => {
  it('is one link with the time and the kind at the right', () => {
    const body = html(
      render(FeedRow, {
        props: {
          href: '/r/1',
          when: '4 min ago',
          kind: 'Logs',
          lead: { crest: 'mage', icon: null, quality: null },
          build: '70532',
          children: text('Nightfall uploaded'),
        },
      }).body,
    );
    expect(body).toMatch(/^<a class="ui-feed-row[^"]*"[^>]*href="\/r\/1"/);
    expect(body).toContain('crest');
    expect(body).toContain('4 min ago<span class="ui-feed-kind');
    expect(body).toContain('Nightfall uploaded');
  });
});

describe('GradeTile', () => {
  it('is the letter on its tier colour, read with its meaning', () => {
    const body = html(render(GradeTile, { props: { letter: 'A' } }).body);
    expect(body).toContain('aria-label="Tier A"');
    expect(body).toContain('var(--color-rarity-epic-text)');
  });

  it('bare: the letter in its tier colour, no tile', () => {
    const body = html(render(GradeTile, { props: { letter: 'A', bare: true } }).body);
    expect(body).toContain('color: var(--color-rarity-epic-text)');
    expect(body).not.toContain('background');
  });
});
