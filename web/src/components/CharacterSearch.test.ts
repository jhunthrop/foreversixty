// web/src/components/CharacterSearch.test.ts
// Task L8.2 ruling 7 and fix round 1 item 2: the home page's character search, in the design
// system's Search form. The routing itself is lib/character-search.ts's (its own test) and
// the submit is driven end to end by tests/e2e/home-character-search.spec.ts.
import type { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { createPageContainer } from '../test-support/page-container';
import { beforeAll, describe, expect, it } from 'vitest';
import { homeCharacterSearchCopy as copy } from '../lib/home-landing-copy';
import { BUTTON_PRIMARY } from '../lib/ui/button';
import CharacterSearch from './CharacterSearch.astro';

let html: string;

beforeAll(async () => {
  const container: AstroContainer = await createPageContainer();
  html = await container.renderToString(CharacterSearch);
});

function classOf(pattern: RegExp): string[] {
  return (pattern.exec(html)?.[1] ?? '').split(/\s+/);
}

describe('CharacterSearch', () => {
  it('labels all three fields in the quiet field-label recipe, Character name, Region and Ruleset', () => {
    for (const [id, label] of [
      ['home-character-search-name', copy.nameLabel],
      ['home-character-search-region', copy.regionLabel],
      ['home-character-search-ruleset', copy.rulesetLabel],
    ]) {
      expect(html).toMatch(new RegExp(`<label class="quiet-label"[^>]*for="${id}"[^>]*>${label}</label>`));
    }
  });

  it('says how to use it, with one muted line beside the heading', () => {
    expect(copy.hint).toBe('Type the name, then pick the region and ruleset.');
    expect(html).toContain(copy.hint);
    expect(copy.aside).toBe('See any character’s Forever Score.');
    expect(html).toMatch(/class="text-muted text-\[14px\]" data-testid="home-character-search-aside"/);
  });

  it('is the design system’s Search: 52px, warm border, gold icon, a "/" hint', () => {
    const input = classOf(/<input[^>]*class="([^"]*)"/);
    expect(input).toEqual(expect.arrayContaining(['h-[52px]', 'border-line-warm-strong']));
    expect(html).toMatch(/<svg class="text-gold[^"]*"/);
    expect(html).toMatch(/<kbd[^>]*>\/<\/kbd>/);
    expect(html).toContain('aria-keyshortcuts="/"');
  });

  it('keeps the placeholder at body size so it does not read as a typed value', () => {
    expect(classOf(/<input[^>]*class="([^"]*)"/)).toEqual(
      expect.arrayContaining(['placeholder:text-[15px]', 'placeholder:text-muted']),
    );
  });

  it('makes Look up the filled gold primary button, full width on a phone', () => {
    const button = classOf(/<button[^>]*class="([^"]*)"/);
    expect(button).toEqual(expect.arrayContaining(BUTTON_PRIMARY.split(' ')));
    expect(button).toEqual(expect.arrayContaining(['col-span-2', 'w-full', 'md:col-span-1', 'md:w-auto']));
    expect(/<button[^>]*>([\s\S]*?)<\/button>/.exec(html)?.[1]?.trim()).toBe(copy.submit);
  });

  it('stacks on a phone: the name full width, region and ruleset side by side on the 52px Select', () => {
    expect(classOf(/<form class="([^"]*)"/)).toEqual(expect.arrayContaining(['grid', 'grid-cols-2']));
    for (const name of ['region', 'ruleset']) {
      expect(classOf(new RegExp(`<select[^>]*name="${name}"[^>]*class="([^"]*)"`))).toEqual(
        expect.arrayContaining(['h-[52px]', 'appearance-none']),
      );
    }
  });

  it('starts on US and Normal', () => {
    expect(html).toMatch(/<option value="us"[^>]* selected/);
    expect(html).toMatch(/<option value="normal"[^>]* selected/);
  });
});

describe('CharacterSearch (fix round 2 item 7)', () => {
  it('reads "Character name" as its placeholder', () => {
    expect(copy.namePlaceholder).toBe('Character name');
    expect(html).toContain('placeholder="Character name"');
  });

  it('is the #lookup target a "Look up another character" link lands on, and the box "/" focuses', () => {
    expect(html).toMatch(/<section id="lookup"/);
    expect(html).toContain('data-character-search-primary');
  });
});
