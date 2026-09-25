// web/src/components/character/CharacterIdentity.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { MeCharacter } from '../../lib/account/api';
import CharacterIdentity from './CharacterIdentity.svelte';

const CHAR: MeCharacter = {
  key: 'us/hardcore/elyra-duskvale',
  region: 'us',
  ruleset: 'hardcore',
  name: 'Elyra Duskvale',
  class: 'priest',
  race: 'Night Elf',
  realm: 'Whitemane',
  level: 60,
};

describe('CharacterIdentity', () => {
  it('renders the full descriptor', () => {
    const { body } = render(CharacterIdentity, {
      props: { character: CHAR, size: 'md', descriptor: 'full' },
    });
    expect(body).toContain('Night Elf Priest · Level 60 · Whitemane (Hardcore US)');
  });

  it('renders the realm descriptor from ruleset and region alone', () => {
    const { body } = render(CharacterIdentity, {
      props: { character: CHAR, size: 'md', descriptor: 'realm' },
    });
    expect(body).toContain('Hardcore · US');
    expect(body).not.toContain('Night Elf');
  });

  it('renders no descriptor line for "none"', () => {
    const { body } = render(CharacterIdentity, {
      props: { character: CHAR, size: 'md', descriptor: 'none', descriptorTestid: 'd' },
    });
    expect(body).not.toContain('data-testid="d"');
  });

  it('renders the name as plain text with no href', () => {
    const { body } = render(CharacterIdentity, {
      props: { character: CHAR, size: 'md', descriptor: 'none', nameTestid: 'n' },
    });
    expect(body).toContain('<span');
    expect(body).toContain('data-testid="n"');
    expect(body).not.toContain('<a');
  });

  it('renders the name as a link with href, in the display face only at lg', () => {
    const md = render(CharacterIdentity, {
      props: { character: CHAR, size: 'md', descriptor: 'none', href: '/character/x', nameTestid: 'n' },
    });
    expect(md.body).toContain('<a');
    expect(md.body).toContain('href="/character/x"');
    expect(md.body).not.toContain('font-display');

    const lg = render(CharacterIdentity, {
      props: { character: CHAR, size: 'lg', descriptor: 'none', href: '/character/x' },
    });
    expect(lg.body).toContain('font-display');
  });

  it('uses a caller-chosen portrait test id prefix, defaulting to "character"', () => {
    const { body } = render(CharacterIdentity, {
      props: { character: CHAR, size: 'md', descriptor: 'none' },
    });
    expect(body).toContain('data-testid="character-avatar-fallback"');

    const custom = render(CharacterIdentity, {
      props: { character: CHAR, size: 'md', descriptor: 'none', testid: 'home-hero' },
    });
    expect(custom.body).toContain('data-testid="home-hero-avatar-fallback"');
  });

  it('wraps the name link in an <h1> when heading is set (Ruling 5)', () => {
    const { body } = render(CharacterIdentity, {
      props: { character: CHAR, size: 'lg', descriptor: 'none', href: '/character/x', heading: true },
    });
    expect(body).toContain('<h1');
    const h1Index = body.indexOf('<h1');
    const aIndex = body.indexOf('<a');
    const closeH1Index = body.indexOf('</h1>');
    expect(aIndex).toBeGreaterThan(h1Index);
    expect(aIndex).toBeLessThan(closeH1Index);
  });

  it('wraps the plain name span in an <h1> when heading is set and href is omitted', () => {
    const { body } = render(CharacterIdentity, {
      props: { character: CHAR, size: 'lg', descriptor: 'none', heading: true, nameTestid: 'n' },
    });
    expect(body).toContain('<h1');
    const h1Index = body.indexOf('<h1');
    const nameSpanIndex = body.indexOf('data-testid="n"');
    const closeH1Index = body.indexOf('</h1>');
    expect(nameSpanIndex).toBeGreaterThan(h1Index);
    expect(nameSpanIndex).toBeLessThan(closeH1Index);
    expect(body).not.toContain('<a');
  });

  it('renders no <h1> when heading is omitted or false (every other caller)', () => {
    const omitted = render(CharacterIdentity, {
      props: { character: CHAR, size: 'md', descriptor: 'none', href: '/character/x' },
    });
    expect(omitted.body).not.toContain('<h1');

    const explicit = render(CharacterIdentity, {
      props: { character: CHAR, size: 'md', descriptor: 'none', heading: false },
    });
    expect(explicit.body).not.toContain('<h1');
  });
});
