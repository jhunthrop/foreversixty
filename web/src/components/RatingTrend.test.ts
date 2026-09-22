// web/src/components/RatingTrend.test.ts
// The brief for this task (task-7-brief.md) wrote this test against
// @testing-library/svelte's render()/screen/container.querySelector, but neither
// @testing-library/svelte nor jsdom is installed in this repo (checked package.json and
// node_modules -- absent), the same gap RatingPanel.test.ts and HelpNote.test.ts already
// document. RatingTrend has no `$effect` at all -- it is pure, synchronous, prop-derived
// markup -- so it renders identically under `svelte/server`'s render(), which returns
// `{ head, body }` strings rather than a DOM. This file asserts against those strings
// instead of DOM queries: the intent of the brief's two cases is unchanged, only the
// query mechanism is adapted.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import RatingTrend from './RatingTrend.svelte';

describe('RatingTrend', () => {
  it('draws one point per fight as an inline svg polyline, oldest first', () => {
    const { body } = render(RatingTrend, {
      props: {
        points: [
          { fought_at: '2026-12-01T00:00:00Z', overall: 40 },
          { fought_at: '2026-12-05T00:00:00Z', overall: 70 },
          { fought_at: '2026-12-09T00:00:00Z', overall: 55 },
        ],
      },
    });
    expect(body).toContain('<svg');
    const match = body.match(/<polyline[^>]*\spoints="([^"]*)"/);
    expect(match).not.toBeNull();
    const pairs = (match?.[1] ?? '').trim().split(' ').filter(Boolean);
    expect(pairs).toHaveLength(3);
    // Oldest first: each pair is `x,y` -- the first fight (lowest overall, 40) sorts to
    // the leftmost x (smallest), the last fight (2026-12-09) sorts to the rightmost x.
    const xs = pairs.map((pair) => Number(pair.split(',')[0]));
    expect(xs[0]).toBeLessThan(xs[xs.length - 1]);
  });

  it('renders nothing (an empty fragment) for fewer than two points', () => {
    const { body } = render(RatingTrend, {
      props: { points: [{ fought_at: '2026-12-01T00:00:00Z', overall: 40 }] },
    });
    expect(body).not.toContain('<svg');
  });
});
