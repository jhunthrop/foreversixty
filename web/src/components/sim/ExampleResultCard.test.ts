// web/src/components/sim/ExampleResultCard.test.ts
import { JSDOM } from 'jsdom';
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { EXAMPLE_IMAGE_HEIGHT, EXAMPLE_IMAGE_WIDTH } from '../../lib/billing/example-image';
import { serverRunCaption } from '../../lib/billing/premium-example-copy';
import type { ServerRunExample } from '../../lib/billing/premium-examples';
import { landingCopy } from '../../lib/sim/landing-copy';
import ExampleResultCard from './ExampleResultCard.svelte';

const VALID: ServerRunExample = {
  simId: 'abc123',
  characterName: 'Kraggor',
  level: 60,
  specLabel: 'Fury Warrior',
  preset: true,
  placeholder: false,
  combinations: 4812,
  ranAt: '2026-10-27T18:00:00Z',
};

function renderCard(example?: ServerRunExample): Document {
  const { body } = render(ExampleResultCard, { props: { example } });
  return new JSDOM(body).window.document;
}

describe('ExampleResultCard', () => {
  it.each([true, false])('caption and alt follow the data file (preset=%s)', (preset) => {
    const ex = { ...VALID, preset };
    const doc = renderCard(ex);
    expect(doc.body.textContent).toContain(serverRunCaption(ex));
    expect(doc.querySelector('img')?.getAttribute('alt')).toBe(serverRunCaption(ex));
  });

  it('img carries width, height and both formats', () => {
    const doc = renderCard(VALID);
    const img = doc.querySelector('img');
    expect(img?.getAttribute('width')).toBe(String(EXAMPLE_IMAGE_WIDTH));
    expect(img?.getAttribute('height')).toBe(String(EXAMPLE_IMAGE_HEIGHT));
    const types = [...doc.querySelectorAll('picture source')].map((s) => s.getAttribute('type'));
    expect(types).toEqual(['image/avif', 'image/webp']);
  });

  it('no invented figure and no link', () => {
    const doc = renderCard();
    expect(doc.body.textContent).not.toContain('1,245 DPS');
    expect(doc.querySelector('a')).toBeNull();
  });

  it('shows the B5 label and hint under the testid', () => {
    const doc = renderCard();
    expect(doc.querySelector('[data-testid="sim-example-result"]')).not.toBeNull();
    expect(doc.body.textContent).toContain(landingCopy.exampleCaption);
    expect(doc.body.textContent).toContain(landingCopy.exampleHint);
  });
});
