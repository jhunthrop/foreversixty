// Both apps' feature keys and bundled defaults are generated from apps/shared/features.json
// (web/scripts/gen-mobile-features.mjs). After editing it run `node web/scripts/gen-mobile-features.mjs`
// and commit the result.
import { readFileSync } from 'node:fs';
import { describe, expect, it } from 'vitest';
import { OUTPUTS, buildFeatures, generate } from '../../../scripts/gen-mobile-features.mjs';

const HINT = 'run `node web/scripts/gen-mobile-features.mjs` and commit the result';

describe('mobile feature switches codegen', () => {
  const generated = generate() as Record<string, string>;

  it.each(Object.values(OUTPUTS) as string[])('%s matches the generator', (file) => {
    expect(readFileSync(file, 'utf8'), HINT).toBe(generated[file]);
  });

  it('every feature ships dark until its lane turns it on from the API', () => {
    const features = buildFeatures(
      readFileSync(new URL('../../../../apps/shared/features.json', import.meta.url), 'utf8'),
    ) as { key: string; default: boolean }[];
    expect(features.filter((f) => f.default).map((f) => f.key)).toEqual([]);
  });

  it('refuses a camelCase key and a missing default', () => {
    expect(() => buildFeatures('{"features":{"raidSignup":{"default":false,"about":"x"}}}')).toThrow(
      /snake_case/,
    );
    expect(() => buildFeatures('{"features":{"raids":{"about":"x"}}}')).toThrow(/default/);
  });
});
