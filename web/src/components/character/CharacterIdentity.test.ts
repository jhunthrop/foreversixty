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
});
