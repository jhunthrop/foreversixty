import { describe, expect, it } from 'vitest';
import { SimApiError } from './api';
import { rerunOutcome } from './server-runs-rerun';

describe('rerunOutcome', () => {
  it('is started when the call succeeds', async () => {
    await expect(rerunOutcome(() => Promise.resolve({ sim_id: 'new' }))).resolves.toBe('started');
  });

  it('is failed for a non-402 API error and for a network error', async () => {
    await expect(rerunOutcome(() => Promise.reject(new SimApiError('no', 500)))).resolves.toBe('failed');
    await expect(rerunOutcome(() => Promise.reject(new TypeError('offline')))).resolves.toBe('failed');
  });

  it('is lapsed for a 402', async () => {
    await expect(rerunOutcome(() => Promise.reject(new SimApiError('plan', 402)))).resolves.toBe('lapsed');
  });
});
