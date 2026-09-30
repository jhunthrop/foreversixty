import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { Me } from '../../lib/account/api';
import HomeSwitchCharacterPanel from './HomeSwitchCharacterPanel.svelte';

const USER: Me['user'] = { id: 1, battletag: 'Fixture#1', email: null, role: 'user', anonymize: false };

const ME: Me = {
  user: USER,
  guilds: [],
  characters: [
    { key: 'us/normal/frostspine', region: 'us', ruleset: 'normal', name: 'Frostspine', class: 'mage' },
    { key: 'us/normal/zulmara', region: 'us', ruleset: 'normal', name: 'Zulmara', class: 'hunter' },
    { key: 'us/normal/grokmar', region: 'us', ruleset: 'normal', name: 'Grokmar', class: 'warrior' },
  ],
};

describe('HomeSwitchCharacterPanel', () => {
  it('lists the current character first, then every other in account order', () => {
    const { body } = render(HomeSwitchCharacterPanel, {
      props: { me: ME, currentKey: 'us/normal/zulmara', onswitch: () => {} },
    });
    const zulmaraIndex = body.indexOf('Zulmara');
    const frostspineIndex = body.indexOf('Frostspine');
    const grokmarIndex = body.indexOf('Grokmar');
    expect(zulmaraIndex).toBeGreaterThan(-1);
    expect(zulmaraIndex).toBeLessThan(frostspineIndex);
    expect(frostspineIndex).toBeLessThan(grokmarIndex);
  });

  it('marks the current character and offers Switch for every other', () => {
    const { body } = render(HomeSwitchCharacterPanel, {
      props: { me: ME, currentKey: 'us/normal/zulmara', onswitch: () => {} },
    });
    expect(body).toContain('data-testid="current-character-bar-switch-current"');
    expect(body).toContain('data-testid="current-character-bar-switch-row-us/normal/frostspine"');
  });

  it('renders the header label and the Add one link', () => {
    const { body } = render(HomeSwitchCharacterPanel, {
      props: { me: ME, currentKey: 'us/normal/zulmara', onswitch: () => {} },
    });
    expect(body).toContain('Switch character');
    expect(body).toContain('href="/account#add-character"');
    expect(body).toContain('Add one');
  });

  it('falls back to account order when no character matches currentKey', () => {
    const { body } = render(HomeSwitchCharacterPanel, {
      props: { me: ME, currentKey: 'us/normal/unknown', onswitch: () => {} },
    });
    expect(body.indexOf('Frostspine')).toBeLessThan(body.indexOf('Zulmara'));
  });

  it('draws every row portrait as the same 36px ringed ClassCrest, never an avatar (review round 1 item 3)', () => {
    const { body } = render(HomeSwitchCharacterPanel, {
      props: { me: ME, currentKey: 'us/normal/zulmara', onswitch: () => {} },
    });
    expect(body).toContain('src="/icons/hd/crests/hunter.webp"');
    expect(body).toContain('src="/icons/hd/crests/mage.webp"');
    expect(body).toContain('src="/icons/hd/crests/warrior.webp"');
    expect((body.match(/width="36"/g) ?? []).length).toBe(3);
  });

  it("gives every row crest ClassCrest's own rest/hover/focus-within ring recipe (ux-designer round 2)", () => {
    // svelte/server's render() does not inline a component's own scoped <style> text into
    // `body` (Svelte injects it separately at hydration time), so this checks the markup
    // the CSS selectors below hook into -- the row carries the `crest-row` hover/focus-
    // within scope and the crest carries `row-crest` plus its own `--c` custom property,
    // exactly like ClassCrest.astro's `.crest`/`--c` pair. The rendered ring/outline
    // itself is verified visually (recaptured hover/focus screenshots, this lane's report).
    const { body } = render(HomeSwitchCharacterPanel, {
      props: { me: ME, currentKey: 'us/normal/zulmara', onswitch: () => {} },
    });
    expect(body).toContain('class="crest-row');
    expect(body).toContain('class="row-crest');
    expect(body).toContain('style="--c: var(--color-class-hunter);"');
  });
});
