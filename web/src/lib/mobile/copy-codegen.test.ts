// Both apps' strings are generated from apps/shared/copy/en.json (web/scripts/gen-mobile-copy.mjs).
// This fails when a checked-in catalog differs from a fresh generation: after editing en.json run
// `node web/scripts/gen-mobile-copy.mjs` and commit the result.
import { readFileSync } from 'node:fs';
import { describe, expect, it } from 'vitest';
// @ts-expect-error plain .mjs script, no type declarations
import {
  OUTPUTS,
  generate,
  parseEntry,
  renderAndroid,
  renderXcstrings,
} from '../../../scripts/gen-mobile-copy.mjs';

const HINT = 'run `node web/scripts/gen-mobile-copy.mjs` and commit the result';

describe('mobile copy codegen', () => {
  const generated = generate() as Record<string, string>;

  it.each(Object.values(OUTPUTS) as string[])('%s matches the generator', (file) => {
    expect(readFileSync(file, 'utf8'), HINT).toBe(generated[file]);
  });

  it('numbers placeholders by first appearance, ints as integers', () => {
    const entry = parseEntry('raid.count', {
      text: '{n} of {size} signed up',
      args: { n: 'int', size: 'int' },
    });
    expect(renderAndroid([entry], '// x')).toContain(
      '<string name="raid_count">%1$d of %2$d signed up</string>',
    );
    expect(renderXcstrings([entry])).toContain('"value": "%1$lld of %2$lld signed up"');
  });

  it('writes plurals for both platforms', () => {
    const entry = parseEntry('talents.unspent', {
      count: 'n',
      one: '1 point unspent.',
      other: '{n} points unspent.',
    });
    expect(renderAndroid([entry], '// x')).toContain('<item quantity="other">%1$d points unspent.</item>');
    const ios = JSON.parse(renderXcstrings([entry])).strings['talents.unspent'].localizations.en;
    expect(ios.variations.plural.other.stringUnit.value).toBe('%1$lld points unspent.');
  });

  it('keeps edge spaces and escapes apostrophes for aapt', () => {
    const android = renderAndroid([parseEntry('a.b', " · it's {x}")], '// x');
    expect(android).toContain(`<string name="a_b">" · it\\'s %1$s"</string>`);
  });

  it('refuses mismatched placeholders, stray plural args and banned words', () => {
    expect(() => parseEntry('a.b', { ios: 'Hi {name}', android: 'Hi' })).toThrow(/same placeholders/);
    expect(() => parseEntry('a.b', { count: 'n', one: '1 {x}', other: '{n} {x}' })).toThrow(/names only/);
    expect(() => parseEntry('a.b', 'Go Premium')).toThrow(/banned/);
    expect(() => parseEntry('a.b', 'Done!')).toThrow(/banned/);
    expect(() => parseEntry('A.b', 'x')).toThrow(/lowerCamel/);
  });
});
