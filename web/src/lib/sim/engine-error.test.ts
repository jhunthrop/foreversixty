import { describe, expect, it } from 'vitest';
import { SimApiError } from './api';
import { simCopy } from './copy';
import { humaniseEngineError, humaniseServerFailure } from './engine-error';
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

// Fix round 1: store.svelte.ts's runOnServer() is the fourth lane, beside a browser run's
// runAndSettle and a bulk run's/count's runBulkAndSettle/recount, and the last one to route
// through humaniseEngineError -- it read SimApiError.message straight onto `message`.
describe('humaniseServerFailure', () => {
  it('humanises a SimApiError whose message names an unsupported spec', () => {
    const error = new SimApiError(
      'combine: part 0 failed: request: unsupported spec: "druid-restoration"',
      0,
    );
    expect(humaniseServerFailure(error, simCopy.failed)).toBe(
      'The engine does not simulate Restoration Druid yet.',
    );
  });

  it('leaves a SimApiError with any other message exactly as it arrived', () => {
    expect(humaniseServerFailure(new SimApiError(simCopy.premiumRequired, 402), simCopy.failed)).toBe(
      simCopy.premiumRequired,
    );
    expect(humaniseServerFailure(new SimApiError(simCopy.notFound, 404), simCopy.failed)).toBe(
      simCopy.notFound,
    );
  });

  it('falls back for anything that is not a SimApiError, the same as runOnServer’s own three catches', () => {
    expect(humaniseServerFailure(new Error('network down'), simCopy.failed)).toBe(simCopy.failed);
    expect(humaniseServerFailure('not even an Error', simCopy.failed)).toBe(simCopy.failed);
    expect(humaniseServerFailure(undefined, simCopy.failed)).toBe(simCopy.failed);
  });
});
