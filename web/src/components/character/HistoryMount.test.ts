import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { parseCharacterPath, type CharacterPath } from '../../lib/characters';
import HistoryMount from './HistoryMount.svelte';

function fixturePath(): CharacterPath {
  const parsed = parseCharacterPath('/character/us/hardcore/elyra-duskvale');
  if (parsed === null) throw new Error('fixture path');
  return parsed;
}
const path = fixturePath();

const facts = { ownerKnown: true, own: true, entitled: true, hasScore: true };

function mount(extra: Partial<typeof facts> = {}): string {
  return render(HistoryMount, { props: { path, ...facts, ...extra } }).body.replace(/<!--[^>]*-->/g, '');
}

describe('HistoryMount', () => {
  it('holds the section s reserved place while its chunk loads, for the entitled owner', () => {
    const html = mount();
    expect(html).toContain('history-shell-reserve');
    expect(html).not.toContain('history-free-sentence');
  });

  it('says one sentence with a See plans link to the free owner', () => {
    const html = mount({ entitled: false });
    expect(html).toContain('History charts draw your score and parses over time.');
    expect(html).toMatch(/<a[^>]*href="\/premium#history"[^>]*>See plans<\/a>/);
    expect(html).toMatch(/over time\. <span class="whitespace-nowrap"><a[^>]*>See plans<\/a>\.<\/span>/);
    expect(html).not.toContain('history-shell-reserve');
  });

  it('draws nothing for a visitor', () => {
    expect(mount({ own: false, entitled: false }).trim()).toBe('');
    expect(mount({ own: false }).trim()).toBe('');
  });

  it('draws nothing for an owner with no score', () => {
    expect(mount({ hasScore: false }).trim()).toBe('');
    expect(mount({ hasScore: false, entitled: false }).trim()).toBe('');
  });

  it('holds a place, shown only under the pre-paint hint, until /v1/me answers', () => {
    expect(mount({ ownerKnown: false, own: false, entitled: false })).toContain('history-shell-reserve');
  });
});
