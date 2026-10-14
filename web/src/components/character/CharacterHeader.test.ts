// web/src/components/character/CharacterHeader.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import CharacterHeader from './CharacterHeader.svelte';

const ANSWER = {
  strings: { failed: 'Could not load.' },
  load: () => Promise.reject(new Error('not used')),
};

/** The text a crawler would read: the markup with tags removed (the aria-hidden separator dots are decoration). */
function readableText(html: string): string {
  return html
    .replace(/<style[\s\S]*?<\/style>/g, '')
    .replace(/<[^>]*>/g, '')
    .replace(/&nbsp;|&middot;|\u00b7/g, '')
    .trim();
}

describe('CharacterHeader loading state (the server-rendered reserved box)', () => {
  const props = {
    action: { label: 'Open', href: '/bis' },
    answer: ANSWER,
    nameHeading: true,
    whenSignedOut: 'none' as const,
  };

  it('carries no readable text and never says "Placeholder"', () => {
    const { body } = render(CharacterHeader, { props });
    expect(body).toContain('data-testid="character-header-loading"');
    expect(body).not.toMatch(/placeholder/i);
    expect(readableText(body)).toBe('');
  });

  it('never renders the character name as a heading, even when the page asks for one', () => {
    const { body } = render(CharacterHeader, { props });
    expect(body).not.toContain('<h1');
  });
});
