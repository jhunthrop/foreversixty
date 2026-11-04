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
  it('below 20 never asks for a spec: Talents teaches the leveling order and the sim card is quiet', () => {
    const { body } = render(HomeHeroCards, { props: { hero: { ...HERO, level: 5, spec: undefined } } });
    expect(body).not.toContain('Pick a spec');
    expect(body).toContain('data-testid="home-talents-leveling"');
    expect(body).toContain('Sim numbers start at level 20.');
  });

  it('is exactly two cards, Talents and Simulator: Best in slot is the header’s answer now', () => {
    const { body } = render(HomeHeroCards, { props: { hero: HERO } });
    expect(body).toContain('data-testid="home-hero-card-talents"');
    expect(body).toContain('data-testid="home-hero-card-sim"');
    expect(body).not.toContain('home-hero-card-bis');
    expect(body).not.toContain('href="#upgrades"');
  });

  it('points a character with no talent export at the guide build in the planner', () => {
    const { body } = render(HomeHeroCards, { props: { hero: HERO_WITH_SPEC } });
    expect(body).toContain('Your guide build for this level is in the planner.');
  });

  it('links the Talents card to the planner with ?spec= for a character with a learned spec', () => {
    const { body } = render(HomeHeroCards, { props: { hero: HERO_WITH_SPEC } });
    expect(body).toContain('href="/planner?spec=hunter-marksmanship"');
  });

  // A character with no level reads as below 20, whose sim card is quiet (the low-level
  // home); the sim card's link and skeleton belong to a character from 20 up.
  it('links the Simulator card to the armory sim href for this character', () => {
    const { body } = render(HomeHeroCards, { props: { hero: HERO_WITH_SPEC } });
    expect(body).toContain('href="/sim?source=armory&amp;ref=us%2Fnormal%2Fzulmara"');
  });

  it('shows the Simulator card loading skeleton on first render (SSR, before the sim fetch settles)', () => {
    const { body } = render(HomeHeroCards, { props: { hero: HERO_WITH_SPEC } });
    expect(body).toContain('data-testid="home-hero-card-sim-skeleton"');
  });
});
