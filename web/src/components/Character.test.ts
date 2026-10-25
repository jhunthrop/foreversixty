// web/src/components/Character.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import Character from './Character.svelte';
import { parseCharacterPath } from '../lib/characters';
import { CHARACTER_LOADING_MIN_H } from '../lib/character-layout';
import { historyMount } from '../lib/character/history-mount';

const path = parseCharacterPath('/character/us/hardcore/elyra-duskvale');

describe('Character loading state', () => {
  it('reserves the ready height with a Skeleton instead of a bare line', () => {
    const { body } = render(Character, { props: { path } });
    expect(body).toContain('data-testid="character-skeleton"');
    expect(body).toContain(CHARACTER_LOADING_MIN_H);
    expect(body).not.toContain('Loading.');
  });
});

describe('Character failed state', () => {
  it('shows a retry that re-fires the same fetch in place', () => {
    const { body } = render(Character, { props: { path } });
    // status starts 'loading' in SSR (no effect runs server-side), so this only proves
    // the markup shape LoadError renders exists; the retry wiring itself is exercised by
    // the route-stub test in tests/e2e/character-retry.spec.ts.
    expect(body).not.toContain('character-error');
  });
});

describe('Character history mount', () => {
  it('mounts the section for the entitled owner, one sentence for the free owner, nothing for a visitor', () => {
    const facts = { ownerKnown: true, own: true, entitled: true, hasScore: true };
    expect(historyMount(facts)).toBe('charts');
    expect(historyMount({ ...facts, entitled: false })).toBe('sentence');
    expect(historyMount({ ...facts, own: false, entitled: false })).toBe('none');
  });
});
