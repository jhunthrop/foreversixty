// web/scripts/inline-code.mjs
// Astro emits `<script is:inline>` bodies and our own `<!-- -->` template comments verbatim,
// so every developer comment in them shipped on every page. The comments stay in source; the
// build ships the scripts minified and the comments stripped. Two
// halves of one rule, as pure functions of HTML text so src/lib/site/inline-code.test.ts can
// drive them (a plain .mjs: astro.config.mjs and the postbuild check import it under node):
//   - shipHtml rewrites one built page (the `ship-minified` integration runs it on dist/).
//   - devCommentFailures lists what a built page still leaks (check-no-dev-comments.mjs).
import { readdirSync, readFileSync, writeFileSync } from 'node:fs';
import { join, relative } from 'node:path';
import { minifySync } from 'vite';

/**
 * One pass over the page: a script block, a style block or an HTML comment, in document
 * order, so a `<!--` inside a script body is never read as a comment.
 */
const BLOCK = /<script\b([^>]*)>([\s\S]*?)<\/script>|<style\b([^>]*)>([\s\S]*?)<\/style>|<!--([\s\S]*?)-->/gi;

/** `type` values that mean JavaScript; anything else (JSON-LD, JSON config) ships as written. */
const JS_TYPES = new Set(['', 'text/javascript', 'application/javascript', 'module']);

/** Words that only ever appear in a source note, never in what a player reads. */
export const INTERNAL_MARKERS = ['lib/', '.test.ts', 'spec 20'];

const typeOf = (attributes) => /\btype\s*=\s*["']?([^"'\s>]+)/i.exec(attributes)?.[1]?.toLowerCase() ?? '';

/** An inline script esbuild may minify: no `src`, a JavaScript type and a non-empty body. */
const isInlineJs = (attributes, body) =>
  !/\bsrc\s*=/i.test(attributes) && JS_TYPES.has(typeOf(attributes)) && body.trim() !== '';

/**
 * A comment we wrote. Svelte's hydration markers (`<!--[-->`, `<!--]-->`, `<!---->`) and
 * Astro's (`<!--astro:end-->`) never hold whitespace; every prose comment does.
 */
const isDevComment = (body) => /\s/.test(body);

/**
 * A script's `code` minified by Vite's own minifier (Oxc, what the bundles ship through),
 * every comment dropped. A classic script is read as a sloppy-mode script, a `module` one as a
 * module. Top-level names are kept, so a global one script defines for another survives, and
 * `document.currentScript` still names the running script. Throws on a parse error rather than
 * shipping a script it could not read.
 */
export function minifyCode(code, module = false) {
  const result = minifySync('inline.js', code, {
    module,
    compress: true,
    mangle: { toplevel: false },
    codegen: { legalComments: 'none' },
  });
  if (result.errors.length > 0) {
    throw new Error(`inline script does not minify: ${result.errors.map((e) => e.message).join('; ')}`);
  }
  return result.code.trim();
}

/** minifyCode once per distinct body: the same layout scripts repeat on every page. */
const cached = (cache, code, module) => {
  const key = `${module}\0${code}`;
  if (!cache.has(key)) cache.set(key, minifyCode(code, module));
  return cache.get(key);
};

/**
 * The page with its inline scripts minified and our HTML comments removed. Styles ship as
 * written: Vite minifies every processed one and the inline ones carry no comments, which
 * devCommentFailures holds them to.
 */
export function shipHtml(html, cache = new Map()) {
  return html.replace(BLOCK, (whole, scriptAttrs, scriptBody, styleAttrs, _styleBody, comment) => {
    if (scriptAttrs !== undefined && isInlineJs(scriptAttrs, scriptBody)) {
      const module = typeOf(scriptAttrs) === 'module';
      return `<script${scriptAttrs}>${cached(cache, scriptBody, module)}</script>`;
    }
    if (comment !== undefined && isDevComment(comment)) return '';
    return whole;
  });
}

/** Every developer comment a built page still carries, and any internal path anywhere in it; empty when clean. */
export function devCommentFailures(html) {
  const failures = [];
  for (const [, scriptAttrs, scriptBody, , styleBody, comment] of html.matchAll(BLOCK)) {
    if (comment !== undefined) {
      if (isDevComment(comment)) failures.push(`HTML comment <!--${comment.trim().slice(0, 60)}-->`);
      continue;
    }
    const isScript = scriptAttrs !== undefined;
    // JSON-LD and JSON config blocks are data, not code, and ship as written.
    if (isScript && !isInlineJs(scriptAttrs, scriptBody)) continue;
    const code = isScript ? scriptBody : styleBody;
    const line = isScript ? /^\s*\/\/.*$/m.exec(code) : null;
    if (line) failures.push(`// comment in an inline script: ${line[0].trim().slice(0, 60)}`);
    const block = /\/\*[\s\S]*?\*\//.exec(code);
    if (block) failures.push(`/* */ comment inline: ${block[0].slice(0, 60)}`);
  }
  for (const marker of INTERNAL_MARKERS) {
    if (html.includes(marker)) failures.push(`internal path "${marker}" in the page`);
  }
  return failures;
}

function* htmlFiles(dir) {
  for (const entry of readdirSync(dir, { withFileTypes: true })) {
    const path = join(dir, entry.name);
    if (entry.isDirectory()) yield* htmlFiles(path);
    else if (entry.isFile() && entry.name.endsWith('.html')) yield path;
  }
}

/** Rewrites every built page under `dist` with shipHtml; returns the bytes saved. */
export function shipDist(dist) {
  const cache = new Map();
  let saved = 0;
  for (const file of htmlFiles(dist)) {
    const html = readFileSync(file, 'utf8');
    const shipped = shipHtml(html, cache);
    if (shipped === html) continue;
    saved += Buffer.byteLength(html) - Buffer.byteLength(shipped);
    writeFileSync(file, shipped);
  }
  return saved;
}

/** devCommentFailures for every built page under `dist`, each prefixed with its path from `base`. */
export function distDevCommentFailures(dist, base) {
  const failures = [];
  for (const file of htmlFiles(dist)) {
    const name = relative(base, file);
    for (const failure of devCommentFailures(readFileSync(file, 'utf8')))
      failures.push(`${name}: ${failure}`);
  }
  return failures;
}
