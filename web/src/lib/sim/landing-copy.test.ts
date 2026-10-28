// web/src/lib/sim/landing-copy.test.ts
import { describe, expect, it } from 'vitest';
import { specKeyCell } from '../../test-support/spec-table';
import { landingCopy } from './landing-copy';

const SPEC = new URL('../../../../design/specs/2026-10-27-launch-pricing.md', import.meta.url);

describe('landingCopy example', () => {
  it('label is the B5 label', () =>
    expect(landingCopy.exampleCaption).toBe(
      specKeyCell(SPEC, '### B5. Copy, verbatim', 'landingCopy.exampleCaption')[0],
    ));

  it('has no invented figure keys', () => {
    expect(landingCopy).not.toHaveProperty('exampleSpec');
    expect(landingCopy).not.toHaveProperty('exampleDps');
  });
});
