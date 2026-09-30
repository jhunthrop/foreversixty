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

describe('HomeHeroCards', () => {
  it('renders the Best in slot and Talents cards as the honest not-available sentence, never a fabricated figure', () => {
    const { body } = render(HomeHeroCards, { props: { hero: HERO } });
    expect(body).toContain('Not available yet: the addon does not send worn gear.');
    expect(body).toContain('Not available yet: the addon does not send talents.');
    expect(body).toContain('href="#upgrades"');
    expect(body).toContain('href="/planner"');
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
