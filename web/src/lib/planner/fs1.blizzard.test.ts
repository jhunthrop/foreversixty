// web/src/lib/planner/fs1.blizzard.test.ts
// Spec 2026-09-22 §3.5: the API lane's bnetbuild encoder writes
// api/internal/bnetbuild/testdata/era-kiloz.fs1, a checked-in FS1 string encoded from the
// Era Kiloz fixture profile (api/internal/bnetapi/testdata/era-kiloz/). This decodes it
// with the exact same decoder every addon paste and sim-input read goes through, proving
// the two lanes' grammar agrees without either lane importing the other's code. Skips
// cleanly (not a failure) when the API lane has not landed that file on this branch yet.
import { readFileSync, existsSync } from 'node:fs';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';
import { decodeFS1 } from './fs1';

const FIXTURE_PATH = fileURLToPath(
  new URL('../../../../api/internal/bnetbuild/testdata/era-kiloz.fs1', import.meta.url),
);

describe('a Battle.net-built FS1 string (bnetbuild)', () => {
  it.skipIf(!existsSync(FIXTURE_PATH))("decodes with this site's own FS1 decoder", () => {
    const code = readFileSync(FIXTURE_PATH, 'utf-8').trim();
    const result = decodeFS1(code);
    expect(result.ok).toBe(true);
    if (result.ok) {
      expect(result.build.classSlug).toBe('warrior');
    }
  });

  if (!existsSync(FIXTURE_PATH)) {
    it('is skipped: api/internal/bnetbuild/testdata/era-kiloz.fs1 is absent on this branch', () => {
      expect(existsSync(FIXTURE_PATH)).toBe(false);
    });
  }
});
