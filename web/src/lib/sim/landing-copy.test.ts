// web/src/lib/sim/landing-copy.test.ts
import { describe, expect, it } from 'vitest';
import { landingCopy } from './landing-copy';

describe('landingCopy has no example', () => {
  it.each(['exampleCaption', 'exampleHint', 'exampleSpec', 'exampleDps'])('no %s key', (key) => {
    expect(landingCopy).not.toHaveProperty(key);
  });
});
