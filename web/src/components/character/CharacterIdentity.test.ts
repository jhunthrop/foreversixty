// web/src/components/character/CharacterIdentity.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { createRawSnippet } from 'svelte';
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

  it('renders the name as a link with href, in the heading face only at lg', () => {
    const md = render(CharacterIdentity, {
      props: { character: CHAR, size: 'md', descriptor: 'none', href: '/character/x', nameTestid: 'n' },
    });
    expect(md.body).toContain('<a');
    expect(md.body).toContain('href="/character/x"');
    expect(md.body).not.toContain('heading-name');

    const lg = render(CharacterIdentity, {
      props: { character: CHAR, size: 'lg', descriptor: 'none', href: '/character/x' },
    });
    // heading-name, not heading-face: a character name keeps its own case (type pass 2026-10-29).
    expect(lg.body).toContain('heading-name');
  });

  it('uses a caller-chosen portrait test id prefix, defaulting to "character"', () => {
    const { body } = render(CharacterIdentity, {
      props: { character: CHAR, size: 'md', descriptor: 'none' },
    });
    expect(body).toContain('data-testid="character-avatar"');

    const custom = render(CharacterIdentity, {
      props: { character: CHAR, size: 'md', descriptor: 'none', testid: 'home-hero' },
    });
    expect(custom.body).toContain('data-testid="home-hero-avatar"');
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

  // Finding 3, 2026-09-24 landing pass: CharacterRow.svelte's own pill and guild line
  // render through this slot, in the text column, after the descriptor line.
  it('renders `below` content in the text column, after the descriptor line, when given', () => {
    const below = createRawSnippet(() => ({
      render: () => '<span data-testid="below-marker">extra</span>',
    }));
    const { body } = render(CharacterIdentity, {
      props: { character: CHAR, size: 'md', descriptor: 'full', descriptorTestid: 'd', below },
    });
    const descriptorIndex = body.indexOf('data-testid="d"');
    const belowIndex = body.indexOf('below-marker');
    expect(descriptorIndex).toBeGreaterThan(-1);
    expect(belowIndex).toBeGreaterThan(descriptorIndex);
  });

  // L8.4 fix round 4 item 8: the profile tops the crest against the name block, so the owner's
  // extra sync line does not move it; every other caller keeps it centred.
  it('tops the portrait against the name block with align="start", and centres it by default', () => {
    const block = (body: string): string[] =>
      (/^<div class="([^"]*)"/.exec(body.replace(/<!--[^>]*-->/g, ''))?.[1] ?? '').split(' ');
    const start = render(CharacterIdentity, {
      props: { character: CHAR, size: 'lg', descriptor: 'none', align: 'start' },
    }).body;
    expect(block(start)).toContain('items-start');
    expect(block(start)).not.toContain('items-center');
    const centred = render(CharacterIdentity, {
      props: { character: CHAR, size: 'lg', descriptor: 'none' },
    }).body;
    expect(block(centred)).toContain('items-center');
  });

  it('renders nothing extra when `below` is omitted (every other caller)', () => {
    const { body } = render(CharacterIdentity, {
      props: { character: CHAR, size: 'md', descriptor: 'none' },
    });
    expect(body).not.toContain('below-marker');
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

  it('sets the xl name at 34px, capped at two lines, in a heading box that reserves both', () => {
    const { body } = render(CharacterIdentity, {
      props: { character: CHAR, size: 'xl', descriptor: 'none', heading: true },
    });
    expect(body).toContain('text-[34px]');
    expect(body).toContain('line-clamp-2');
    expect(body).toContain('min-h-[calc(34px*1.1*2)]');
  });

  it('leaves every smaller name unclamped and unreserved', () => {
    const { body } = render(CharacterIdentity, {
      props: { character: CHAR, size: 'lg', descriptor: 'none', heading: true },
    });
    expect(body).not.toContain('line-clamp-2');
    expect(body).not.toContain('min-h-');
  });

  it('draws the Supporter chip after the name only when the character is one', () => {
    const base = { character: CHAR, size: 'lg', descriptor: 'none' } as const;
    const on = render(CharacterIdentity, { props: { ...base, supporter: true } }).body;
    const off = render(CharacterIdentity, { props: { ...base, supporter: false } }).body;
    const absent = render(CharacterIdentity, { props: base }).body;
    expect(on.split('data-testid="supporter-chip"')).toHaveLength(2);
    expect(on.indexOf('Elyra Duskvale')).toBeLessThan(on.indexOf('supporter-chip'));
    expect(off).not.toContain('supporter-chip');
    expect(absent).not.toContain('supporter-chip');
  });
});
