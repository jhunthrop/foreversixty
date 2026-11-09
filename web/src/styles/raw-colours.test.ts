// web/src/styles/raw-colours.test.ts
// Tokens are the only source of colour (owner 2026-11-09: "we should use shared components and
// resources as a first class citizen"). design/tokens/tokens.json generates tokens.css; the system
// layer (styles/*.css and the components/ui/ primitives) may compose them, and every page and
// feature component reads a token by name: a Tailwind colour utility (`text-muted`, `bg-raised`,
// `bg-white/[.06]`), `var(--color-*)`, or a primitive. This fails on a hex, rgb(), rgba(), hsl() or
// hsla() literal anywhere else under src/, so a page cannot invent a colour.
import { readdirSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';

const SRC = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '..');
const STYLED = new Set(['.astro', '.svelte', '.css']);
/** The system layer: the token output, the global sheets and the primitives. */
const SYSTEM = ['styles/', 'components/ui/'];
const RAW_COLOUR =
  /(?<![\w&-])#(?:[0-9a-f]{8}|[0-9a-f]{6}|[0-9a-f]{3,4})(?![\w-])|(?<![a-z])(?:rgba?|hsla?)\(/gi;

function files(dir: string): string[] {
  return readdirSync(dir, { withFileTypes: true }).flatMap((entry) => {
    const full = path.join(dir, entry.name);
    if (entry.isDirectory()) return entry.name === 'data' || entry.name === 'fixtures' ? [] : files(full);
    return STYLED.has(path.extname(entry.name)) && !entry.name.includes('.test.') ? [full] : [];
  });
}

function relative(file: string): string {
  return path.relative(SRC, file).split(path.sep).join('/');
}

/** Comments name colours to explain them; only code counts. */
function code(text: string): string {
  return text
    .replace(/\/\*[\s\S]*?\*\//g, '')
    .replace(/<!--[\s\S]*?-->/g, '')
    .replace(/^\s*\/\/.*$/gm, '');
}

describe('raw colours', () => {
  it('appear only in the token output, the global sheets and the ui/ primitives', () => {
    const offenders = files(SRC)
      .filter((file) => !SYSTEM.some((prefix) => relative(file).startsWith(prefix)))
      .flatMap((file) => {
        const found = code(readFileSync(file, 'utf8')).match(RAW_COLOUR) ?? [];
        return found.map((literal) => `${relative(file)}: ${literal}`);
      });
    expect(offenders, 'use a token (design/tokens/tokens.json) or a ui/ primitive').toEqual([]);
  });
});
