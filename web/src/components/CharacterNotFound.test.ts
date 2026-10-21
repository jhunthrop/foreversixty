// web/src/components/CharacterNotFound.test.ts
// W.8: a name the site has never heard of says so, instead of drawing an empty character card.
// Server render only (the page's fetch and its 404 are the island's own $effect, which vitest's
// server entry never runs), so this pins the state's markup and Character.svelte's choice of it.
import { readFileSync } from 'node:fs';
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { scoreCopy } from '../lib/score/copy';
import CharacterNotFound from './CharacterNotFound.svelte';

const PATH = { region: 'us' as const, ruleset: 'pvp' as const, slug: 'nobody-here' };

describe('CharacterNotFound', () => {
  const { body } = render(CharacterNotFound, { props: { path: PATH, name: 'Nobody Here' } });

  it('says no character is called that on that ruleset and region', () => {
    expect(scoreCopy.notFoundLine('Nobody Here', 'PvP US')).toBe(
      'No character called Nobody Here on PvP US yet.',
    );
    expect(body).toContain('No character called Nobody Here on PvP US yet.');
  });

  it('is titled "Character not found" and draws no score card', () => {
    expect(scoreCopy.notFoundTitle).toBe('Character not found');
    expect(body).toContain('>Character not found</h1>');
    expect(body).not.toContain('forever-score');
    expect(body).not.toContain('No logs yet');
  });

  it('carries the lookup form, starting on the region and ruleset of the address', () => {
    expect(body).toContain('data-testid="character-not-found-form"');
    expect(body).toContain('name="name"');
    expect(body).toMatch(/<option value="us" selected/);
    expect(body).toMatch(/<option value="pvp" selected/);
  });

  it('offers "Your character? Get the addon"', () => {
    const next = /data-testid="character-not-found-next">([\s\S]*?)<\/p>/.exec(body)?.[1] ?? '';
    expect(next).toContain('Your character?');
    expect(next).toContain('href="/setup"');
    expect(next).toContain('>Get the addon</a>');
  });
});

describe('Character.svelte', () => {
  const source = readFileSync(new URL('./Character.svelte', import.meta.url), 'utf8');

  it('shows the not-found state to a visitor and keeps the empty card for the viewer’s own character', () => {
    expect(source).toMatch(/\{#if owned\}\s*<CharacterUnlogged[\s\S]*\{:else\}\s*<CharacterNotFound/);
  });
});
