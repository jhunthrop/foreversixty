// web/src/components/TierStrip.test.ts
// Fix round 2 items 8 and 9: from sm, five aligned segments with the current one outlined (no
// shift); on a phone, one 12px strip of five segments with only the current tier labelled.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import labels from '../../../docs/contracts/fixtures/score-labels.json';
import TierStrip from './TierStrip.svelte';

function phoneOf(body: string): string {
  return /<div[^>]*data-testid="tier-strip-phone"[\s\S]*?<\/div>\s*(?=<ol)/.exec(body)?.[0] ?? '';
}

describe('TierStrip', () => {
  it('on a phone: one 12px strip of five segments, and only the current tier named', () => {
    const phone = phoneOf(render(TierStrip, { props: { current: 'epic' } }).body);
    expect(phone).toContain('class="flex h-3 ');
    expect(phone.match(/background-color: #/g)).toHaveLength(5);
    expect(phone).toContain('>Epic<');
    expect(phone).toContain('1500–1999');
    for (const tier of ['Common', 'Uncommon', 'Rare', 'Legendary']) expect(phone).not.toContain(`>${tier}<`);
    expect(phone).toContain('aria-label="Forever Score tiers: Epic 1500–1999, this character’s tier"');
  });

  it('on a phone with no current tier, names all five small under their segments', () => {
    const phone = phoneOf(render(TierStrip, { props: {} }).body);
    for (const tier of Object.values(labels.tiers)) expect(phone).toContain(`>${tier.label}<`);
  });

  it('from sm: every segment the same padding and bar height; the current one outlined, not resized', () => {
    const { body } = render(TierStrip, { props: { current: 'epic' } });
    const items = [...body.matchAll(/<li class="([^"]*)"[^>]*data-testid="tier-strip-(\w+)"/g)];
    expect(items).toHaveLength(5);
    for (const [, cls] of items) expect(cls).toContain('p-1.5');
    const epic = items.find(([, , tier]) => tier === 'epic')?.[1] ?? '';
    expect(epic).toContain('outline');
    expect(epic).not.toContain('ring');
    expect(body.match(/class="block h-2 rounded-\[2px\]/g)).toHaveLength(5);
  });
});
