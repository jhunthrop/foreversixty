// web/src/components/character/ClassCrestRing.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import ClassCrestRing from './ClassCrestRing.svelte';

describe('ClassCrestRing', () => {
  it('draws the class crest for a known class', () => {
    const { body } = render(ClassCrestRing, { props: { characterClass: 'warrior', size: 26, testid: 't' } });
    expect(body).toContain('/icons/hd/crests/warrior.webp');
    expect(body).toContain('data-testid="t"');
    expect(body).not.toContain('data-testid="t-fallback"');
  });

  // Guild control-centre live-fix round, defect 4: a roster character with no known class
  // (readiness/roster rows read `class: ""` rather than omitting the field) must render
  // CharacterPortrait's own neutral ringed disc, never a broken `/icons/hd/crests/.webp`
  // request for an empty slug.
  it('renders the neutral ringed disc for an empty class, never a broken crest image', () => {
    const { body } = render(ClassCrestRing, { props: { characterClass: '', size: 26, testid: 't' } });
    expect(body).toContain('data-testid="t-fallback"');
    expect(body).toContain('rounded-full');
    expect(body).not.toContain('/icons/hd/crests/.webp');
    expect(body).not.toContain('data-testid="t"><img');
  });

  it('sizes the fallback disc the same pixel size the crest image would have taken', () => {
    const { body } = render(ClassCrestRing, { props: { characterClass: '', size: 40 } });
    expect(body).toContain('width: 40px');
    expect(body).toContain('height: 40px');
  });
});
