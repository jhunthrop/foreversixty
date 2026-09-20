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
    expect(result).toEqual({ ok: false, message: 'That code is too long to read.' });
  });

  it('names an unlabelled code from the copy file, not an inline literal, when there is no prefix to read', () => {
    const result = decodeFSB1('');
    expect(result.ok).toBe(false);
    if (result.ok) return;
    expect(result.message).toBe(addonCopy.wrongPrefix(addonCopy.unlabelledCode, FSB1_PREFIX));
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

  it('writes a tier above 9 as a base-36 digit', () => {
    const build: FSB1Build = {
      dataBuild: '1',
      classSlug: 'paladin',
      order: [{ tab: 1, tier: 10, column: 2 }],
      gear: [],
    };
    expect(encodeFSB1(build)).toBe('FSB1:1:paladin:1a2:');
  });

  it('sorts a stat outside contract 10.8’s vocabulary after every known one, alphabetically among themselves', () => {
    // stamina is pinned (contract 10.8), so it always sorts first here; zzz_unknown and
    // aaa_unknown are both outside the pinned vocabulary, so this also proves two unknown
    // names get a stable, defined order between them rather than an arbitrary one -- the
    // exact property the parallel Lua codec has to match.
    const build: FSB1Build = {
      dataBuild: '1',
      classSlug: 'paladin',
      order: [],
      gear: [{ slot: 'head', itemId: 1, stats: { zzz_unknown: 9, stamina: 5, aaa_unknown: 3 } }],
    };
    expect(encodeFSB1(build)).toBe('FSB1:1:paladin::head=1:stamina=5;aaa_unknown=3;zzz_unknown=9');
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
