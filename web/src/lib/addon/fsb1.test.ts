import { existsSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';
import vectors from '../../fixtures/addon/codec-vectors.json';
import { addonCopy } from './copy';
import { decodeFSB1, encodeFSB1, FSB1_PREFIX, MAX_CODE_LENGTH, type FSB1Build } from './fsb1';

describe('decodeFSB1', () => {
  it.each(vectors.fsb1)('reads the shared vector: $name', ({ code, build }) => {
    const result = decodeFSB1(code);
    expect(result.ok).toBe(true);
    if (!result.ok) return;
    expect(result.build).toEqual(build);
  });

  it.each(vectors.fsb1Invalid)('refuses with a reason: $name', ({ code }) => {
    const result = decodeFSB1(code);
    expect(result.ok).toBe(false);
    if (result.ok) return;
    expect(result.message.length).toBeGreaterThan(0);
  });

  it('names the prefix it found and the one it wanted', () => {
    const result = decodeFSB1('FS1:1:paladin:111:');
    expect(result.ok).toBe(false);
    if (result.ok) return;
    expect(result.message).toContain('FS1');
    expect(result.message).toContain('FSB1');
  });

  it('refuses a code past the bound before parsing it', () => {
    const result = decodeFSB1('x'.repeat(MAX_CODE_LENGTH + 1));
    expect(result).toEqual({ ok: false, message: addonCopy.tooLong });
  });

  it('names an unlabelled code from the copy file, not an inline literal, when there is no prefix to read', () => {
    const result = decodeFSB1('');
    expect(result.ok).toBe(false);
    if (result.ok) return;
    expect(result.message).toBe(addonCopy.wrongPrefix(addonCopy.unlabelledCode, FSB1_PREFIX));
  });

  it('reads a negative stat value: real items carry them and the penalty has to survive', () => {
    // Ring of Scorn is spirit -3 in data/builds/1.60.1.69893/items/*.json; 43 stat values
    // across the nine class files are negative. A decoder that refused them would refuse
    // the code the site's own "Copy addon code" button produces for a player wearing one.
    const result = decodeFSB1('FSB1:1.60.1.69893:warrior::finger1=3235:stamina=4;spirit=-3');
    expect(result.ok).toBe(true);
    if (!result.ok) return;
    expect(result.build.gear).toEqual([{ slot: 'finger1', itemId: 3235, stats: { stamina: 4, spirit: -3 } }]);
  });

  it.each(['-', '-0', '--3', '3-', '+3', '- 3'])('refuses the malformed stat value %j', (value) => {
    const result = decodeFSB1(`FSB1:1.60.1.69893:warrior::head=12640:spirit=${value}`);
    expect(result.ok).toBe(false);
    if (result.ok) return;
    expect(result.message).toBe(addonCopy.statPair(`spirit=${value}`));
  });

  it('refuses an empty data build field', () => {
    // `parts.length < 5` already refuses a code that is missing a field; a field that is
    // present but empty is the same defect with a colon in front of it. Neither the addon's
    // staleness check nor its tree lookup can do anything with one.
    const result = decodeFSB1('FSB1::warrior:111:');
    expect(result.ok).toBe(false);
    if (result.ok) return;
    expect(result.message).toBe(addonCopy.emptyField(addonCopy.dataBuildField));
  });

  it('refuses an empty class field', () => {
    const result = decodeFSB1('FSB1:1.60.1.69893::111:');
    expect(result.ok).toBe(false);
    if (result.ok) return;
    expect(result.message).toBe(addonCopy.emptyField(addonCopy.classField));
  });
});

describe('encodeFSB1', () => {
  // A vector's `code` is a decoder INPUT, hand-written when the fixture was authored, not
  // necessarily this format's canonical encoding of the build beside it: `encodeStats`
  // sorts a gear entry's stats by name and a hand-written code need not already agree.
  // The first shared vector spells its stats `stamina=17;spell_power=23`, which name order
  // reverses -- so "a code equals its own re-encoding" is simply not a property FSB1 has,
  // and asserting it here (as this file used to) was asserting against the authoritative
  // fixture rather than with it. The addon lane's codec_fsb1_spec.lua makes exactly this
  // distinction for exactly this reason; the properties below are the ones that do hold,
  // and they are the same three it checks.
  it.each(vectors.fsb1)('re-encodes the shared vector to a fixed point: $name', ({ code }) => {
    const decoded = decodeFSB1(code);
    expect(decoded.ok).toBe(true);
    if (!decoded.ok) return;

    // Nothing is lost going from a decoded build to a string and back.
    const encoded = encodeFSB1(decoded.build);
    const redecoded = decodeFSB1(encoded);
    expect(redecoded.ok).toBe(true);
    if (!redecoded.ok) return;
    expect(redecoded.build).toEqual(decoded.build);

    // And the string that came out is a fixed point under another pass, so the canonical
    // form is reached in one encode rather than converging over several.
    expect(encodeFSB1(redecoded.build)).toBe(encoded);
  });

  it('writes the first shared vector’s stats in name order, reversing the order the fixture spells them in', () => {
    // Pinned by value rather than only by the fixed point above: the fixture's own
    // `stamina=17;spell_power=23` is a hand-written input, and name order puts
    // `spell_power` first. Do not "fix" this assertion back toward the fixture's spelling
    // -- that is the mistake this whole change is undoing. The addon's Codec.lua asserts
    // the identical pair in the identical direction.
    const decoded = decodeFSB1('FSB1:1.60.1.69893:paladin:111112121:head=12640:stamina=17;spell_power=23');
    expect(decoded.ok).toBe(true);
    if (!decoded.ok) return;
    expect(encodeFSB1(decoded.build)).toBe(
      'FSB1:1.60.1.69893:paladin:111112121:head=12640:spell_power=23;stamina=17',
    );
  });

  it('sorts a slot’s stats so one build is one string', () => {
    const build: FSB1Build = {
      dataBuild: '1',
      classSlug: 'paladin',
      order: [],
      gear: [{ slot: 'head', itemId: 1, stats: { stamina: 2, agility: 1 } }],
    };
    expect(encodeFSB1(build)).toBe('FSB1:1:paladin::head=1:agility=1;stamina=2');
  });

  it('writes a negative stat with its sign, and never writes a negative zero', () => {
    const build: FSB1Build = {
      dataBuild: '1',
      classSlug: 'paladin',
      order: [],
      gear: [{ slot: 'head', itemId: 1, stats: { spirit: -3, stamina: -0.4 } }],
    };
    // Math.round(-0.4) is -0, and `${-0}` is "0" -- so the wire never carries a "-0" the
    // decoder would then have to refuse. `spirit` leads `stamina` because the stats are
    // written in name order ("spi" < "sta").
    expect(encodeFSB1(build)).toBe('FSB1:1:paladin::head=1:spirit=-3;stamina=0');
    expect(decodeFSB1(encodeFSB1(build)).ok).toBe(true);
  });

  it('writes a tier above 9 as a base-36 digit', () => {
    const build: FSB1Build = {
      dataBuild: '1',
      classSlug: 'paladin',
      order: [{ tab: 1, tier: 10, column: 2 }],
      gear: [],
    };
    expect(encodeFSB1(build)).toBe('FSB1:1:paladin:1a2:');
  });

  it('sorts every stat by byte order, pinned vocabulary or not, and never by locale collation', () => {
    // `stamina` is in contract 10.8's vocabulary and the other three are not, and that now
    // buys it nothing: names sort against each other as bytes, so `stamina` lands between
    // `aaa_unknown` and `zzz_unknown` rather than ahead of all three. Unknown names
    // therefore get the same stable, defined order as known ones instead of a separate
    // rule -- the property the parallel Lua codec has to match, with no shared table.
    //
    // The pair is deliberately case-mixed, and with name order that is now the primary
    // comparator rather than a tie-break: `'Bonus_unknown'.localeCompare('aaa_unknown')` is
    // 1 (collation folds case and sorts a before B) while `'Bonus_unknown' < 'aaa_unknown'`
    // is true (byte order puts every uppercase letter before every lowercase one). Lua's
    // `table.sort` on strings is byte order, so byte order is what this format uses, and
    // `localeCompare` would additionally vary by the browser's locale. A test with only
    // lowercase names could not tell the two apart.
    const build: FSB1Build = {
      dataBuild: '1',
      classSlug: 'paladin',
      order: [],
      gear: [
        { slot: 'head', itemId: 1, stats: { zzz_unknown: 9, stamina: 5, aaa_unknown: 3, Bonus_unknown: 7 } },
      ],
    };
    expect(encodeFSB1(build)).toBe(
      'FSB1:1:paladin::head=1:Bonus_unknown=7;aaa_unknown=3;stamina=5;zzz_unknown=9',
    );
  });
});

describe('codec-vectors.json byte identity', () => {
  // The addon lane's copy at addon/tests/fixtures/codec-vectors.json is authoritative; this
  // guards the web copy against drifting from it. Two copies rather than one shared path
  // because the web bundler and the addon's zip each need the file inside their own tree,
  // so this is the gate that a regeneration reached both. It used to be conditional on the
  // addon copy existing, because that lane had not landed; it has now, the file is tracked,
  // and the guard is unconditional -- a missing file is a real failure rather than a silent
  // skip. addon/tests/vectors_spec.lua asserts the same identity from the other side.
  const webFixturePath = fileURLToPath(new URL('../../fixtures/addon/codec-vectors.json', import.meta.url));
  const addonFixturePath = path.resolve(
    webFixturePath,
    '../../../../../addon/tests/fixtures/codec-vectors.json',
  );

  it('matches the addon lane’s copy byte-for-byte', () => {
    expect(existsSync(addonFixturePath)).toBe(true);
    expect(readFileSync(webFixturePath)).toEqual(readFileSync(addonFixturePath));
  });
});
