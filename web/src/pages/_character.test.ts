// web/src/pages/_character.test.ts
// L8.4 fix round 3 item 9: the character page with JavaScript off. The island fetches the
// character in the browser, so the server HTML is only the loading skeleton; without
// JavaScript that skeleton would shimmer forever. A <noscript> hides it and says, in one line
// from the copy module, why nothing is there and where to go instead (the Planner works
// without this page). Both shells carry it: the Worker's (character.astro) and the fixture
// build's (character/[...path].astro).
import type { AstroComponentFactory } from 'astro/runtime/server/index.js';
import { describe, expect, it } from 'vitest';
import { characterNoScriptCopy } from '../lib/character-profile';
import { createPageContainer } from '../test-support/page-container';
import CharacterShell from './character.astro';
import CharacterFixture from './character/[...path].astro';

const SENTENCE =
  'Forever Sixty needs JavaScript to show this character. Turn it on and reload, or open the Planner.';

describe('the character page without JavaScript', () => {
  it('words the line once, in the copy module', () => {
    expect(`${characterNoScriptCopy.lead} ${characterNoScriptCopy.link}${characterNoScriptCopy.tail}`).toBe(
      SENTENCE,
    );
    expect(characterNoScriptCopy.href).toBe('/planner');
  });

  type Container = Awaited<ReturnType<typeof createPageContainer>>;
  const SHELLS: [string, (container: Container) => Promise<string>][] = [
    ['the Worker shell', (container) => container.renderToString(CharacterShell)],
    [
      'the fixture page',
      (container) =>
        // The route's props type comes from its getStaticPaths (as in _sim.test.ts); the
        // container renders it from params alone.
        container.renderToString(CharacterFixture as unknown as AstroComponentFactory, {
          params: { path: 'us/hardcore/elyra-duskvale' },
        }),
    ],
  ];
  for (const [name, renderShell] of SHELLS) {
    it(`${name}: hides the skeleton and shows the line with the Planner link`, async () => {
      const html = await renderShell(await createPageContainer());
      const noscript = /<noscript>([\s\S]*?)<\/noscript>/.exec(html)?.[1] ?? '';
      expect(noscript).toMatch(
        /<style>\s*\[data-testid=["']character-skeleton["']\]\s*\{\s*display:\s*none;?\s*\}\s*<\/style>/,
      );
      expect(noscript).toContain('data-testid="character-noscript"');
      const words = noscript
        .replace(/<style>[\s\S]*?<\/style>/, '')
        .replace(/<[^>]+>/g, '')
        .replace(/\s+/g, ' ')
        .trim();
      expect(words).toBe(SENTENCE);
      expect(noscript).toMatch(/<a [^>]*href="\/planner"[^>]*>Planner<\/a>/);
      // The skeleton it hides is in the same server HTML.
      expect(html).toContain('data-testid="character-skeleton"');
    });
  }
});
