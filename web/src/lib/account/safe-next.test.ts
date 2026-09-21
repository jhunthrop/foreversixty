// web/src/lib/account/safe-next.test.ts
import { describe, expect, it } from 'vitest';
import { safeNextPath } from './safe-next';

describe('safeNextPath', () => {
  it('accepts a same-site path', () => {
    expect(safeNextPath('/premium/checkout?plan=premium&interval=monthly', '/logs')).toBe(
      '/premium/checkout?plan=premium&interval=monthly',
    );
  });

  it('falls back on null or missing', () => {
    expect(safeNextPath(null, '/logs')).toBe('/logs');
    expect(safeNextPath(undefined, '/logs')).toBe('/logs');
    expect(safeNextPath('', '/logs')).toBe('/logs');
  });

  it('rejects a protocol-relative path (host smuggled in)', () => {
    expect(safeNextPath('//evil.example/x', '/logs')).toBe('/logs');
    expect(safeNextPath('/\\evil.example/x', '/logs')).toBe('/logs');
  });

  it('rejects an absolute URL with a scheme', () => {
    expect(safeNextPath('https://evil.example', '/logs')).toBe('/logs');
    expect(safeNextPath('javascript:alert(1)', '/logs')).toBe('/logs');
  });

  it('rejects a path that does not start with a single slash', () => {
    expect(safeNextPath('logs', '/logs')).toBe('/logs');
  });
});
