import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { MeCharacter } from '../../lib/account/api';
import HomeHeroCards from './HomeHeroCards.svelte';

const HERO: MeCharacter = {
  key: 'us/normal/zulmara',
  region: 'us',
  ruleset: 'normal',
  name: 'Zulmara',
  class: 'hunter',
};

const HERO_WITH_SPEC: MeCharacter = {
  ...HERO,
  spec: 'Marksmanship',
  level: 24,
  faction: 'horde',
};

describe('HomeHeroCards', () => {
  it('shows "Pick a spec" on Best in slot and Talents for a character with no spec yet, never a fabricated figure', () => {
    const { body } = render(HomeHeroCards, { props: { hero: HERO } });
    expect(body).toContain('Pick a spec');
    expect(body).toContain('href="#upgrades"');
    expect(body).toContain('href="/planner"');
  });

  it('shows the honest not-yet-synced sentence for a character with a spec but no addon export yet', () => {
    const { body } = render(HomeHeroCards, { props: { hero: HERO_WITH_SPEC } });
    expect(body).toContain(
      'Not available yet: no gear export for this character. Open the addon once to send it.',
    );
    expect(body).toContain(
      'Not available yet: no talent export for this character. Open the addon once to send it.',
    );
  });

  it('links the Simulator card to the armory sim href for this character', () => {
    const { body } = render(HomeHeroCards, { props: { hero: HERO } });
    expect(body).toContain('href="/sim?source=armory&amp;ref=us%2Fnormal%2Fzulmara"');
  });

  it('shows the Simulator card loading skeleton on first render (SSR, before the sim fetch settles)', () => {
    const { body } = render(HomeHeroCards, { props: { hero: HERO } });
    expect(body).toContain('data-testid="home-hero-card-sim-skeleton"');
  });
});
