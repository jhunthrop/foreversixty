// web/src/components/sim/LandingCard.test.ts
// W.12 ruling 3: the simulator landing lists the current character first, a healing or tanking
// spec has no gold "Run sim" (a quiet "Not simulated yet" where the button would be), and the
// landing's character card has the same shape as the simulator page's card: crest, name,
// "Race Spec Class · level", the source badge.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { MeCharacter } from '../../lib/account/api';
import { landingCopy } from '../../lib/sim/landing-copy';
import LandingState from './LandingState.svelte';
import SimRunBlock from './SimRunBlock.svelte';

const build = { source: 'addon' as const, captured_at: '2026-10-20T00:00:00Z' };
const BOW: MeCharacter = {
  key: 'us/pvp/bow-jackzon',
  name: 'Bow Jackzon',
  class: 'hunter',
  race: 'High Order Skyborne',
  level: 60,
  region: 'us',
  ruleset: 'pvp',
  build: { ...build, talents: { trees: ['', '', ''], points: [31, 16, 4] } },
};
const PIPSQUEAK: MeCharacter = {
  key: 'us/pvp/pipsqueak',
  name: 'Pipsqueak',
  class: 'priest',
  race: 'Gnome',
  level: 60,
  region: 'us',
  ruleset: 'pvp',
  build: { ...build, talents: { trees: ['', '', ''], points: [0, 46, 5] } },
};

function order(body: string): string[] {
  return [...body.matchAll(/data-testid="sim-character-link-([^"]+)"/g)].map((match) => match[1]);
}

describe('the sim landing', () => {
  it('lists the current character first, then the rest by level and name', () => {
    const props = { busyKey: null, onpick: () => {}, onother: () => {} };
    const first = render(LandingState, {
      props: { ...props, characters: [BOW, PIPSQUEAK], currentKey: PIPSQUEAK.key },
    });
    expect(order(first.body)).toEqual([PIPSQUEAK.key, BOW.key]);
    const second = render(LandingState, {
      props: { ...props, characters: [PIPSQUEAK, BOW], currentKey: BOW.key },
    });
    expect(order(second.body)).toEqual([BOW.key, PIPSQUEAK.key]);
  });

  it('a healing spec has no Sim button, only "Not simulated yet"', () => {
    const { body } = render(LandingState, {
      props: { characters: [BOW, PIPSQUEAK], busyKey: null, onpick: () => {}, onother: () => {} },
    });
    expect(body).toContain(`data-testid="sim-pick-${BOW.key}"`);
    expect(body).not.toContain(`data-testid="sim-pick-${PIPSQUEAK.key}"`);
    expect(body).toContain(`data-testid="sim-not-simulated-${PIPSQUEAK.key}"`);
    expect(body).toContain(landingCopy.notSimulatedYet);
  });
});

describe('the landing card', () => {
  const props = { path: null, busy: false, onrun: () => {} };

  it('has the simulator card’s shape: name, "Race Spec Class · level", the source badge', () => {
    const { body } = render(SimRunBlock, { props: { ...props, character: PIPSQUEAK } });
    expect(body).toContain('Pipsqueak');
    expect(body).toContain('Gnome Holy Priest · 60');
    expect(body).toContain('data-testid="sim-run-block-pill"');
    expect(body).toContain('Addon');
  });

  it('names the spec by the talents, as the simulator card does', () => {
    const { body } = render(SimRunBlock, { props: { ...props, character: BOW } });
    expect(body).toContain('High Order Skyborne Beast Mastery Hunter · 60');
  });

  it('keeps the gold "Run sim" for a damage spec and drops it for a healer', () => {
    expect(render(SimRunBlock, { props: { ...props, character: BOW } }).body).toContain(
      'data-testid="sim-run-block-action"',
    );
    const healer = render(SimRunBlock, { props: { ...props, character: PIPSQUEAK } }).body;
    expect(healer).not.toContain('data-testid="sim-run-block-action"');
    expect(healer).toContain('Not simulated yet');
  });
});
