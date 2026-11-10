// web/src/lib/site/chrome-aliases.test.ts
// The API's /b/{id} header loads its scripts by stable path (scripts/chrome-aliases.mjs): each
// build writes /_chrome/<name>.js re-exporting its own hashed /_astro/ file, so a site deploy
// never strands the API's copy of the chrome on hashes that no longer exist.
import { mkdirSync, mkdtempSync, readFileSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import path from 'node:path';
import { afterEach, beforeEach, describe, expect, it } from 'vitest';
import {
  aliasModule,
  hashedScripts,
  stableChromePath,
  writeChromeAliases,
} from '../../../scripts/chrome-aliases.mjs';

const PAGE = [
  '<astro-island component-url="/_astro/CharacterSelector.B1w9MGHd.js" component-export="default"',
  ' renderer-url="/_astro/client.svelte.CUVksGWg.js"></astro-island>',
  '<script type="module" src="/_astro/Header.astro_astro_type_script_index_0_lang.Cfbe9qLK.js"></script>',
  '<link rel="stylesheet" href="/_astro/about.Dx1_aB-9.css">',
].join('');

describe('hashedScripts', () => {
  it('names every hashed script once, by its name without the hash, and skips stylesheets', () => {
    expect(hashedScripts(PAGE + PAGE)).toEqual([
      { name: 'CharacterSelector', url: '/_astro/CharacterSelector.B1w9MGHd.js' },
      { name: 'client.svelte', url: '/_astro/client.svelte.CUVksGWg.js' },
      {
        name: 'Header.astro_astro_type_script_index_0_lang',
        url: '/_astro/Header.astro_astro_type_script_index_0_lang.Cfbe9qLK.js',
      },
    ]);
  });

  it('refuses two builds of one name: the stable path would be ambiguous', () => {
    expect(() => hashedScripts('/_astro/a.AAAAAAAA.js /_astro/a.BBBBBBBB.js')).toThrow(/built twice/);
  });
});

describe('aliasModule', () => {
  it('re-exports the default too when the hashed module has one (an island or a renderer)', () => {
    expect(aliasModule('/_astro/x.AAAAAAAA.js', 'const a=1;export{a as default};')).toBe(
      "export * from '/_astro/x.AAAAAAAA.js';\nexport { default } from '/_astro/x.AAAAAAAA.js';\n",
    );
  });

  it('only imports a script with no default (the header script)', () => {
    expect(aliasModule('/_astro/h.AAAAAAAA.js', 'document.body;export{b as c};')).toBe(
      "export * from '/_astro/h.AAAAAAAA.js';\n",
    );
  });
});

describe('writeChromeAliases', () => {
  let dist = '';
  beforeEach(() => {
    dist = mkdtempSync(path.join(tmpdir(), 'chrome-aliases-'));
    mkdirSync(path.join(dist, '_astro'));
    writeFileSync(path.join(dist, 'about.html'), PAGE);
    writeFileSync(path.join(dist, '_astro/CharacterSelector.B1w9MGHd.js'), 'export{s as default};');
    writeFileSync(path.join(dist, '_astro/client.svelte.CUVksGWg.js'), 'export default 1;');
    writeFileSync(path.join(dist, '_astro/Header.astro_astro_type_script_index_0_lang.Cfbe9qLK.js'), 'x();');
  });
  afterEach(() => rmSync(dist, { recursive: true, force: true }));

  it('writes one stable module per header script, where stableChromePath says', () => {
    expect(writeChromeAliases(dist)).toHaveLength(3);
    const selector = readFileSync(path.join(dist, stableChromePath('CharacterSelector')), 'utf8');
    expect(selector).toContain("export { default } from '/_astro/CharacterSelector.B1w9MGHd.js';");
    const header = readFileSync(
      path.join(dist, stableChromePath('Header.astro_astro_type_script_index_0_lang')),
      'utf8',
    );
    expect(header).toBe("export * from '/_astro/Header.astro_astro_type_script_index_0_lang.Cfbe9qLK.js';\n");
  });

  it('fails the build when the page names no script: the API header would load nothing', () => {
    writeFileSync(path.join(dist, 'about.html'), '<p>no scripts</p>');
    expect(() => writeChromeAliases(dist)).toThrow(/names no/);
  });
});
