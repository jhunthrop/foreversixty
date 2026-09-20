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
  it.each(vectors.fsb1)('round-trips the shared vector: $name', ({ code, build }) => {
    expect(encodeFSB1(build as FSB1Build)).toBe(code);
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
    // decoder would then have to refuse.
    expect(encodeFSB1(build)).toBe('FSB1:1:paladin::head=1:stamina=0;spirit=-3');
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

  it('sorts a stat outside contract 10.8’s vocabulary after every known one, by byte order among themselves', () => {
    // stamina is pinned (contract 10.8), so it always sorts first here; the other three are
    // outside the pinned vocabulary, so this also proves unknown names get a stable, defined
    // order between them rather than an arbitrary one -- the exact property the parallel Lua
    // codec has to match.
    //
    // The pair is deliberately case-mixed: `'Bonus_unknown'.localeCompare('aaa_unknown')` is
    // 1 (collation folds case and sorts a before B) while `'Bonus_unknown' < 'aaa_unknown'`
    // is true (byte order puts every uppercase letter before every lowercase one). Lua's
    // `table.sort` on strings is byte order, so byte order is what this format uses; a test
    // with only lowercase names could not tell the two apart.
    const build: FSB1Build = {
      dataBuild: '1',
      classSlug: 'paladin',
      order: [],
      gear: [
        { slot: 'head', itemId: 1, stats: { zzz_unknown: 9, stamina: 5, aaa_unknown: 3, Bonus_unknown: 7 } },
      ],
    };
    expect(encodeFSB1(build)).toBe(
      'FSB1:1:paladin::head=1:stamina=5;Bonus_unknown=7;aaa_unknown=3;zzz_unknown=9',
    );
  });
});

describe('codec-vectors.json byte identity', () => {
  // The addon lane's copy at addon/tests/fixtures/codec-vectors.json is authoritative;
  // this guards the web copy against drifting from it. That lane has not landed yet (see
  // task-15's controller notes), so until its copy exists on disk this check is a
  // documented no-op rather than a failure -- it starts asserting the moment the file
  // shows up.
  const webFixturePath = fileURLToPath(new URL('../../fixtures/addon/codec-vectors.json', import.meta.url));
  const addonFixturePath = path.resolve(
    webFixturePath,
    '../../../../../addon/tests/fixtures/codec-vectors.json',
  );
  const addonCopyExists = existsSync(addonFixturePath);

  it.runIf(addonCopyExists)('matches the addon lane’s copy byte-for-byte', () => {
    expect(readFileSync(webFixturePath)).toEqual(readFileSync(addonFixturePath));
  });

  it.skipIf(addonCopyExists)('is skipped: the addon lane has not published its copy yet', () => {
    expect(addonCopyExists).toBe(false);
  });
});
