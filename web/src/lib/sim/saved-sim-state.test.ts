// @vitest-environment jsdom
// web/src/lib/sim/saved-sim-state.test.ts
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';
import { createSimApi, fixtureResult, FIXTURE_SIM_ID } from '../../test-support/sim-api';
import { createSavedSimState } from './saved-sim-state.svelte';

const api = createSimApi();

beforeEach(() => api.install());
afterEach(() => api.reset());

describe('createSavedSimState', () => {
  it('starts with the initial result and no error', () => {
    const state = createSavedSimState(fixtureResult);
    expect(state.result).toStrictEqual(fixtureResult);
    expect(state.error).toBeNull();
  });

  it('load() fills in the result on success', async () => {
    const state = createSavedSimState(null);
    state.load(FIXTURE_SIM_ID);
    await vi.waitFor(() => expect(state.result).not.toBeNull());
    expect(state.result?.sim_id).toBe(FIXTURE_SIM_ID);
    expect(state.error).toBeNull();
  });

  it('load() sets the error on failure and clears a stale one on retry', async () => {
    const state = createSavedSimState(null);
    state.load('zzzzzzzzzzzz');
    await vi.waitFor(() => expect(state.error).not.toBeNull());
    const firstError = state.error;
    expect(firstError).not.toBeNull();
    expect(state.result).toBeNull();

    // A retry must not leave the previous message showing while the new request is in
    // flight -- load() clears `error` synchronously, before the fetch resolves either way.
    state.load(FIXTURE_SIM_ID);
    expect(state.error).toBeNull();
    await vi.waitFor(() => expect(state.result).not.toBeNull());
    expect(state.result?.sim_id).toBe(FIXTURE_SIM_ID);
  });
});
