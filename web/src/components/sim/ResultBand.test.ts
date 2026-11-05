// web/src/components/sim/ResultBand.test.ts
// Spec 2026-11-04-sim-results 3.3: the band, on a static render, for both fixture shapes.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import full from '../../fixtures/sim/result-full.json';
import older from '../../fixtures/sim/result.json';
import type { SimResult } from '../../lib/sim/types';
import ResultBand from './ResultBand.svelte';

const band = (result: unknown): string =>
  render(ResultBand, { props: { result: result as SimResult, onrerun: () => {} } }).body;

describe('ResultBand', () => {
  const body = band(full);

  it('states the figure to one decimal, in one primary element', () => {
    expect(/data-testid="sim-dps"[^>]*>([^<]*)</.exec(body)?.[1]).toBe('904.4');
    expect(body).toContain('0.05%');
  });

  it('draws the circular class crest, not a square class icon', () => {
    expect(body).toContain('/icons/hd/crests/warrior');
    expect(body).toContain('crest-ring');
  });

  it('heads with the spec, past the ranker placeholder name', () => {
    expect(/data-testid="sim-saved-title"[^>]*>([^<]*)</.exec(body)?.[1]).toBe('Fury Warrior');
    expect(body).not.toContain('breakdown');
  });

  it('holds the two actions and the next step', () => {
    expect(body).toContain('Open in the sim');
    expect(body).toContain('Copy link');
    expect(body).toContain('Next: find the upgrade that adds the most.');
    expect(body).toContain('href="/sim/gear"');
  });

  it('an older result leaves the Saved fact out and shows no engine sha', () => {
    const old = band(older);
    expect(old).not.toContain('sim-fact-saved');
    expect(old).toContain('sim-fact-runs');
    expect(old).not.toContain('d5c7088');
  });
});
