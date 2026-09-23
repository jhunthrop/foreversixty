// web/src/test-support/reset-query-cache.ts
// vitest setup: the client data cache (lib/data/query.ts) is module state, and a test that
// fetched something would otherwise hand that answer to the next test in the same file.
import { beforeEach } from 'vitest';
import { resetQueryCache } from '../lib/data/query';

beforeEach(() => {
  resetQueryCache();
});
