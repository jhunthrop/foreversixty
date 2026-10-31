// web/src/components/chart/OverviewBrush.test.ts
// The brush's handles are range inputs. A step that does not divide the fight's length makes the
// browser snap the end handle down to the last whole step (a 9.9 s pull drew its whole-fight
// handle at 9 s), so the inputs take any value and the end sits on the fight's exact end.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import OverviewBrush from './OverviewBrush.svelte';

/** Fights whose rate chart takes 1 s, 2 s, 5 s and 10 s points (lib/chart/buckets.ts). */
const FIGHTS = [
  { durationMs: 9_879, label: '0:09.9' },
  { durationMs: 80_261, label: '1:20' },
  { durationMs: 763_140, label: '12:43' },
  { durationMs: 1_800_437, label: '30:00' },
];

function inputs(body: string): string[] {
  return body.match(/<input[^>]*>/g) ?? [];
}

describe('OverviewBrush', () => {
  for (const { durationMs, label } of FIGHTS) {
    it(`puts the whole-fight end handle on the exact end of a ${durationMs} ms fight`, () => {
      const { body } = render(OverviewBrush, {
        props: {
          domainMs: durationMs,
          window: { startMs: 0, endMs: durationMs },
          shape: [1, 2, 3],
          onWindow: () => {},
        },
      });
      const [start, end] = inputs(body);
      expect(start).toContain('step="any"');
      expect(end).toContain('step="any"');
      expect(end).toContain(`max="${durationMs}"`);
      expect(end).toContain(`value="${durationMs}"`);
      expect(end).toContain(`aria-valuetext="${label}"`);
      expect(body).toContain(`>${label}</span>`);
    });
  }

  it('reads a zoomed end that stops short of the fight on the whole-second clock', () => {
    const { body } = render(OverviewBrush, {
      props: { domainMs: 9_879, window: { startMs: 2_000, endMs: 6_000 }, shape: [], onWindow: () => {} },
    });
    expect(inputs(body)[1]).toContain('aria-valuetext="0:06"');
  });
});
