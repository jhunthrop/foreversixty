import { describe, expect, it } from 'vitest';
import { safeNextPath } from './safe-next';
import { setupReturnTarget } from './setup-next';

const CLAIM = '/guild/us/pvp/olympus-xxvii/claim?plan=guild&interval=yearly';

describe('safeNextPath', () => {
  it('answers the fallback, never null, for off-site and protocol-relative next', () => {
    expect(safeNextPath('https://evil.example/claim', '/')).toBe('/');
    expect(safeNextPath('//evil.example/claim', '/')).toBe('/');
    expect(safeNextPath(null, '/')).toBe('/');
  });

  it('passes a same-site claim path through', () => {
    expect(safeNextPath(CLAIM, '/')).toBe(CLAIM);
  });
});

describe('setupReturnTarget', () => {
  it('accepts a claim page with its query', () => {
    expect(setupReturnTarget(CLAIM)).toEqual({
      href: CLAIM,
      guild: { region: 'us', ruleset: 'pvp', slug: 'olympus-xxvii' },
    });
  });

  it.each([
    'https://evil.example/guild/us/pvp/olympus-xxvii/claim',
    '//evil.example/guild/us/pvp/olympus-xxvii/claim',
    '/account',
    '/guild/us/pvp/olympus-xxvii/settings',
    '',
    null,
  ])('has no target for %j', (raw) => {
    expect(setupReturnTarget(raw)).toBeNull();
  });
});
