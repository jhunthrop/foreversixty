// web/src/lib/not-found/suggest.test.ts
// @vitest-environment jsdom
import { describe, expect, it } from 'vitest';
import { editDistance, knownPages, suggestPage, SUGGEST_SCRIPT } from './suggest';

const PAGES = knownPages(
  [{ slug: 'hunter', name: 'Hunter' }],
  [{ classSlug: 'hunter', className: 'Hunter', specSlug: 'beast-mastery', specName: 'Beast Mastery' }],
);

describe('Did you mean', () => {
  it("knows the footer's pages and each guide and BiS spec page", () => {
    expect(PAGES).toEqual(
      expect.arrayContaining([
        ['/changelog', 'Changelog'],
        ['/privacy', 'Privacy'],
        ['/guides/hunter', 'Hunter guides'],
        ['/guides/hunter/beast-mastery', 'Beast Mastery Hunter guide'],
        ['/bis/hunter/beast-mastery', 'Beast Mastery Hunter best in slot'],
      ]),
    );
    expect(PAGES.some(([path]) => path.includes('#'))).toBe(false);
  });

  it('measures edits, stopping past the limit', () => {
    expect(editDistance('/guides/hunter/beastmastery', '/guides/hunter/beast-mastery')).toBe(1);
    expect(editDistance('kitten', 'sitting')).toBe(3);
    expect(editDistance('/a', '/abcdefgh')).toBe(4);
  });

  it('suggests the nearest page within 3 edits, case and a trailing slash aside', () => {
    expect(suggestPage('/guides/hunter/beastmastery', PAGES)).toEqual([
      '/guides/hunter/beast-mastery',
      'Beast Mastery Hunter guide',
    ]);
    expect(suggestPage('/CHANGELOG/', PAGES)?.[0]).toBe('/changelog');
    expect(suggestPage('/changelgo', PAGES)?.[0]).toBe('/changelog');
    expect(suggestPage('/bis/huntr', PAGES)).toBeNull();
    expect(suggestPage('/this-page-does-not-exist', PAGES)).toBeNull();
  });

  it('runs the same search inline: the failed address, the suggestion, the /character/ mark', () => {
    document.body.innerHTML = `<span id="nf-path"></span><script type="application/json" id="nf-pages">${JSON.stringify(PAGES)}</script>
      <div id="nf-suggest" hidden><p>Did you mean <a href="/"></a></p></div>`;
    window.history.replaceState(null, '', '/guides/hunter/beastmastery');
    new Function(SUGGEST_SCRIPT)();
    expect(document.getElementById('nf-path')?.textContent).toBe('/guides/hunter/beastmastery');
    expect(document.getElementById('nf-suggest')?.hidden).toBe(false);
    expect(document.querySelector('#nf-suggest a')?.getAttribute('href')).toBe(
      '/guides/hunter/beast-mastery',
    );
    expect(document.querySelector('#nf-suggest a')?.textContent).toBe('Beast Mastery Hunter guide ›');
    window.history.replaceState(null, '', '/character/us/normal');
    new Function(SUGGEST_SCRIPT)();
    expect(document.documentElement.hasAttribute('data-nf-character')).toBe(true);
  });
});
