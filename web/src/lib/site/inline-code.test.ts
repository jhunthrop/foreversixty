// web/src/lib/site/inline-code.test.ts
// Developer comments stay in source and never ship: the build minifies every inline script and
// drops our HTML comments (scripts/inline-code.mjs, run on dist/ by astro.config.mjs's
// `ship-minified` integration), and postbuild's check-no-dev-comments.mjs fails the build if a
// built page still carries one. This drives both halves over fixtures, and checks that every
// inline script in the layouts and components minifies.
import { mkdirSync, mkdtempSync, readdirSync, readFileSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import path from 'node:path';
import { afterEach, beforeEach, describe, expect, it } from 'vitest';
import {
  devCommentFailures,
  distDevCommentFailures,
  minifyCode,
  shipDist,
  shipHtml,
} from '../../../scripts/inline-code.mjs';

const SRC = path.resolve(import.meta.dirname, '../..');

const PAGE = [
  '<html><head>',
  '<script type="application/ld+json">{"@context":"https://schema.org","name":"Forever Sixty"}</script>',
  '<script is:inline data-me-storage-key="fs.q.me">',
  '  // Reads the key off its own tag before the first paint.',
  '  /* lib/account/layout.ts keeps the same numbers (spec 2026-10-01). */',
  '  const key = document.currentScript.dataset.meStorageKey;',
  '  window.seen = localStorage.getItem(key);',
  '</script>',
  '<script type="module" src="/_astro/page.js"></script>',
  '<script type="application/json" id="config">{"a":"/* kept */"}</script>',
  '</head><body>',
  '<!-- The island draws its own header; this reserve holds its height. -->',
  '<!--[--><p>Hunter</p><!--]--><!----><!--[!--><!--astro:end-->',
  '</body></html>',
].join('\n');

describe('shipHtml', () => {
  const shipped = shipHtml(PAGE);

  it('minifies an inline script and drops its comments', () => {
    expect(shipped).toContain(
      '<script is:inline data-me-storage-key="fs.q.me">const key=document.currentScript.dataset.meStorageKey;window.seen=localStorage.getItem(key);</script>',
    );
    expect(shipped).not.toContain('Reads the key');
    expect(shipped).not.toContain('lib/account');
  });

  it('keeps scripts with a src, JSON-LD and JSON blocks as written', () => {
    expect(shipped).toContain('<script type="module" src="/_astro/page.js"></script>');
    expect(shipped).toContain('{"@context":"https://schema.org","name":"Forever Sixty"}');
    expect(shipped).toContain('{"a":"/* kept */"}');
  });

  it('drops our HTML comments and keeps the hydration markers', () => {
    expect(shipped).not.toContain('The island draws');
    expect(shipped).toContain('<!--[--><p>Hunter</p><!--]--><!----><!--[!--><!--astro:end-->');
  });

  it('leaves a page with nothing to strip byte for byte', () => {
    const clean = '<p>Ready</p><!--[--><script>console.log(1);</script><!--]-->';
    expect(shipHtml(clean)).toBe(clean);
  });

  it('keeps the top-level names one script defines for another', () => {
    expect(minifyCode('function openMenu(button) { return button; }\nconst menuKey = "fs.menu";')).toBe(
      'function openMenu(e){return e}const menuKey=`fs.menu`;',
    );
  });

  it('refuses a script it cannot parse', () => {
    expect(() => shipHtml('<script>const = ;</script>')).toThrow(/does not minify/);
  });
});

describe('devCommentFailures', () => {
  it('flags every comment and internal path the raw page carries', () => {
    const failures = devCommentFailures(PAGE);
    expect(failures).toEqual([
      '// comment in an inline script: // Reads the key off its own tag before the first paint.',
      '/* */ comment inline: /* lib/account/layout.ts keeps the same numbers (spec 2026-1',
      'HTML comment <!--The island draws its own header; this reserve holds its heig-->',
      'internal path "lib/" in the page',
      'internal path "spec 20" in the page',
    ]);
  });

  it('passes the shipped page', () => {
    expect(devCommentFailures(shipHtml(PAGE))).toEqual([]);
  });

  it('flags a comment in an inline style', () => {
    expect(devCommentFailures('<style>/* reserve */ a{color:red}</style>')).toEqual([
      '/* */ comment inline: /* reserve */',
    ]);
  });
});

describe('over a dist directory', () => {
  let dist: string;

  beforeEach(() => {
    dist = mkdtempSync(path.join(tmpdir(), 'inline-code-'));
    mkdirSync(path.join(dist, 'guides'));
    writeFileSync(path.join(dist, 'index.html'), PAGE);
    writeFileSync(path.join(dist, 'guides', 'hunter.html'), PAGE);
  });

  afterEach(() => rmSync(dist, { recursive: true, force: true }));

  it('fails before shipping, rewrites every page and passes after', () => {
    expect(distDevCommentFailures(dist, dist)).toHaveLength(10);
    expect(shipDist(dist)).toBeGreaterThan(0);
    expect(distDevCommentFailures(dist, dist)).toEqual([]);
    expect(readFileSync(path.join(dist, 'guides', 'hunter.html'), 'utf8')).toBe(shipHtml(PAGE));
  });
});

/** Every .astro file under `dir`. */
function* astroFiles(dir: string): Generator<string> {
  for (const entry of readdirSync(dir, { withFileTypes: true })) {
    const file = path.join(dir, entry.name);
    if (entry.isDirectory()) yield* astroFiles(file);
    else if (entry.name.endsWith('.astro')) yield file;
  }
}

describe('the inline scripts in source', () => {
  // A literal `<script is:inline ...>body</script>` (no define:vars, whose body Astro wraps,
  // and no set:html, whose body is a string built in the frontmatter).
  const literal = /<script\s+is:inline(?![^>]*(?:define:vars|set:html|type=))[^>]*>([\s\S]*?)<\/script>/g;
  const scripts = [...astroFiles(SRC)].flatMap((file) =>
    [...readFileSync(file, 'utf8').matchAll(literal)].map(
      (match) => [path.relative(SRC, file), match[1]!] as const,
    ),
  );

  it('finds the layout scripts', () => {
    expect(scripts.filter(([file]) => file === path.join('layouts', 'Base.astro')).length).toBeGreaterThan(1);
  });

  it.each(scripts)('%s minifies without a comment left', (_file, body) => {
    const shipped = minifyCode(body);
    expect(shipped).not.toMatch(/\/\*|^\s*\/\//m);
  });
});
