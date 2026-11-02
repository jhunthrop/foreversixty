// web/src/components/character/HomeUpgradesPanel.test.ts
// Loading, "Your upgrades" is exactly the height its slot reserved: a skeleton drawn at its own
// height is taller than the settled smallest states, and it pushed the home page down for the
// frame before a cached answer replaced it (header-finished-sim's layout shift, 2026-11-02).
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { MeCharacter } from '../../lib/account/me-types';
import { HOME_UPGRADES_RESERVE_HEIGHT } from '../../lib/home/upgrades-layout';
import HomeUpgradesPanel from './HomeUpgradesPanel.svelte';

const HERO = { key: 'us/normal/tester', name: 'Tester', class: 'Warrior', spec: 'Fury' } as MeCharacter;

describe('HomeUpgradesPanel', () => {
  it('holds the reserved height while loading, the card clipping its skeleton', () => {
    const { body } = render(HomeUpgradesPanel, { props: { hero: HERO } });
    expect(body).toContain('data-testid="home-upgrades-skeleton"');
    expect(body).toContain(`height: ${HOME_UPGRADES_RESERVE_HEIGHT}`);
    expect(HOME_UPGRADES_RESERVE_HEIGHT).toBe('var(--home-upgrades-reserve, 147px)');
    expect(body).toMatch(/<section[^>]*class="[^"]*\bh-full\b[^"]*"[^>]*data-testid="home-upgrades-body"/);
  });
});
