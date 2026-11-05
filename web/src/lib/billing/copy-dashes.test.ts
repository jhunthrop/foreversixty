// Guild round 3, item 5: player-facing guild and billing copy never carries an em or en dash
// inside a sentence ("soon — nothing was charged"). A lone dash is the design system's empty-cell
// mark and stays.
import { readFileSync, readdirSync } from 'node:fs';
import { join } from 'node:path';
import { describe, expect, it } from 'vitest';

const LIB = join(__dirname, '..');
const COPY_DIRS = ['billing', 'guild'];
const DASH = /[–—]/;
/** A quoted literal on one line: '...', "..." or `...`. */
const LITERAL = /'(?:[^'\\\n]|\\.)*'|"(?:[^"\\\n]|\\.)*"|`(?:[^`\\]|\\.)*`/g;
const EMPTY_CELL_MARK = /^['"`][–—]['"`]$/;

function sourceFiles(dir: string): string[] {
  return readdirSync(join(LIB, dir))
    .filter((name) => name.endsWith('.ts') && !name.endsWith('.test.ts'))
    .map((name) => join(LIB, dir, name));
}

/** Code with its comments blanked, so a dash in a doc comment is not read as copy. */
function withoutComments(source: string): string {
  return source.replace(/\/\*[\s\S]*?\*\//g, '').replace(/(^|[^:'"`])\/\/.*$/gm, '$1');
}

function dashedLiterals(file: string): string[] {
  const code = withoutComments(readFileSync(file, 'utf8'));
  return (code.match(LITERAL) ?? []).filter((lit) => DASH.test(lit) && !EMPTY_CELL_MARK.test(lit));
}

describe('guild and billing copy', () => {
  it('writes no em or en dash inside a sentence', () => {
    const found = COPY_DIRS.flatMap(sourceFiles).flatMap((file) =>
      dashedLiterals(file).map((lit) => `${file.slice(LIB.length + 1)}: ${lit}`),
    );
    expect(found).toEqual([]);
  });
});
