// web/src/components/Character.test.ts
// The character page's first paint (character spec 2026-11-06 section 3, `loading`): a head and a
// Gear plate of skeleton slots at the final geometry, no "Loading." text in the page.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import Character from './Character.svelte';
import { parseCharacterPath } from '../lib/characters';

const path = parseCharacterPath('/character/us/hardcore/elyra-duskvale');

describe('Character loading state', () => {
  it('holds the head and the gear plate with skeletons', () => {
    const { body } = render(Character, { props: { path } });
    expect(body).toContain('data-testid="character-skeleton"');
    expect(body).toContain('data-region="head"');
    expect(body).toContain('data-region="slots"');
    expect(body).not.toContain('Loading.');
    expect(body).not.toContain('character-error');
  });
});
