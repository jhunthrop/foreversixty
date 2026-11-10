// web/src/components/CharacterSearch.test.ts
// "Look up a character" in the flat system (pages spec 2026-11-08 §4.4; /characters and the 404):
// a Section with FormRow's look. The routing itself is lib/character-search.ts's (its own test)
// and the submit is driven end to end by tests/e2e/home-character-search.spec.ts.
import type { experimental_AstroContainer as AstroContainer } from 'astro/container';
import { createPageContainer } from '../test-support/page-container';
import { beforeAll, describe, expect, it } from 'vitest';
import { homeCharacterSearchCopy as copy } from '../lib/home-landing-copy';
import { BUTTON_PRIMARY } from '../lib/ui/button';
import { FIELD_CONTROL, FIELD_LABEL } from '../lib/ui/field';
import CharacterSearch from './CharacterSearch.astro';

let html: string;

beforeAll(async () => {
  const container: AstroContainer = await createPageContainer();
  html = (await container.renderToString(CharacterSearch)).replace(/<!--[^>]*-->/g, '');
});

function classOf(pattern: RegExp): string[] {
  return (pattern.exec(html)?.[1] ?? '').split(/\s+/);
}

describe('CharacterSearch', () => {
  it('is a Section titled Look up a character, the #lookup target, with no card', () => {
    expect(html).toMatch(/<section id="lookup"[^>]*data-region="section-look-up-a-character"/);
    expect(html).toMatch(/<h2[^>]*>Look up a character<\/h2>/);
    expect(html).not.toMatch(/card-surface|rounded-plate|bg-raised/);
    expect(html).toContain('data-character-search-primary');
  });

  it('labels all three fields in the 12 px caps field label, Character name, Region and Ruleset', () => {
    for (const [id, label] of [
      ['home-character-search-name', copy.nameLabel],
      ['home-character-search-region', copy.regionLabel],
      ['home-character-search-ruleset', copy.rulesetLabel],
    ]) {
      expect(html).toMatch(new RegExp(`<label class="${FIELD_LABEL}" for="${id}"[^>]*>${label}</label>`));
    }
  });

  it('says how to use it, with the muted aside and hint', () => {
    expect(copy.hint).toBe('Type the name, then pick the region and ruleset.');
    expect(html).toContain(copy.hint);
    expect(copy.aside).toBe('See any character’s Forever Score.');
    expect(html).toMatch(/class="text-small text-muted[^"]*" data-testid="home-character-search-aside"/);
  });

  it('draws the one field: the shared field classes, a combobox for the typeahead, "/" focuses it', () => {
    expect(classOf(/<input[^>]*class="([^"]*)"/)).toEqual(expect.arrayContaining(FIELD_CONTROL.split(' ')));
    expect(html).toContain('aria-keyshortcuts="/"');
    expect(html).toContain('role="combobox"');
    expect(html).toContain('placeholder="Character name"');
  });

  it('makes Look up the filled gold primary button, the form named for screen readers', () => {
    const button = classOf(/<button[^>]*class="([^"]*)"/);
    expect(button).toEqual(expect.arrayContaining(BUTTON_PRIMARY.split(' ')));
    expect(/<button[^>]*>([\s\S]*?)<\/button>/.exec(html)?.[1]?.trim()).toBe(copy.submit);
    expect(html).toMatch(/<form[^>]*aria-label="Look up a character"/);
  });

  it('starts on US and Normal, on the compact Select', () => {
    expect(html).toMatch(/<option value="us"[^>]* selected/);
    expect(html).toMatch(/<option value="normal"[^>]* selected/);
    expect(html.match(/data-size="compact"/g)).toHaveLength(2);
  });

  it('keeps a status line for its two validation sentences', () => {
    expect(html).toMatch(/<p id="home-character-search-error"[^>]*role="status"/);
  });
});
