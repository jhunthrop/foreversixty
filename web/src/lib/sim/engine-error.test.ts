import { describe, expect, it } from 'vitest';
import { simCopy } from './copy';
import { humaniseEngineError } from './engine-error';
import { specLabel } from './spec-label';

describe('humaniseEngineError', () => {
  it('translates the unsupported-spec shape, naming the spec by its display label', () => {
    const raw = 'combine: part 0 failed: request: shard 0: unsupported spec: "druid-restoration"';
    expect(humaniseEngineError(raw)).toBe(simCopy.engineUnsupportedSpec(specLabel('druid-restoration')));
    expect(humaniseEngineError(raw)).toBe('The engine does not simulate Restoration Druid yet.');
  });

  it('matches the shape at the top level too, not only wrapped inside combine’s own prefix', () => {
    expect(humaniseEngineError('unsupported spec: "warrior-protection"')).toBe(
      simCopy.engineUnsupportedSpec(specLabel('warrior-protection')),
    );
  });

  it('leaves every other engine message exactly as it arrived', () => {
    expect(humaniseEngineError('boom')).toBe('boom');
    expect(humaniseEngineError('request: unknown buff: "battle-shout"')).toBe(
      'request: unknown buff: "battle-shout"',
    );
    expect(humaniseEngineError('bulk: the build has no such item: 12345')).toBe(
      'bulk: the build has no such item: 12345',
    );
    expect(humaniseEngineError('')).toBe('');
  });
});
