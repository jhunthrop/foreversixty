import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import IntervalSegments from './IntervalSegments.svelte';

describe('IntervalSegments', () => {
  it('draws Yearly then Monthly as native radios with the given one checked', () => {
    const body = render(IntervalSegments, {
      props: { value: 'yearly', name: 'iv', onChange: () => {} },
    }).body;
    expect(body.indexOf('iv-yearly')).toBeLessThan(body.indexOf('iv-monthly'));
    expect(body).toMatch(/<input[^>]*type="radio"/);
    expect(body).toMatch(/<input[^>]*value="yearly"[^>]*checked|<input[^>]*checked[^>]*value="yearly"/);
    expect(body).toContain('role="radiogroup"');
  });
});
