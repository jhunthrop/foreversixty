// @vitest-environment jsdom
// web/src/lib/sim/phase.test.ts
import { afterEach, beforeEach, describe, expect, it } from 'vitest';
import { createSimApi, TEST_API } from '../../test-support/sim-api';
import {
  BUILT_IN_PHASES,
  PHASE_LATER,
  fetchPhases,
  hasOpened,
  openDateLabel,
  phaseAt,
  phaseLabel,
  phaseStart,
} from './phase';

const api = createSimApi();
const phases = BUILT_IN_PHASES;

beforeEach(() => api.install());
afterEach(() => api.reset());

describe('phaseAt', () => {
  it('names the phase a moment falls in, the way api/internal/phase does', () => {
    expect(phaseAt(phases, new Date('2026-09-01T00:00:00Z'))).toBe('pre-beta');
    expect(phaseAt(phases, new Date('2026-09-17T00:00:00Z'))).toBe('beta');
    expect(phaseAt(phases, new Date('2026-11-04T22:59:00Z'))).toBe('beta');
    expect(phaseAt(phases, new Date('2026-11-04T23:00:00Z'))).toBe('launch');
    expect(phaseAt(phases, new Date('2026-12-09T00:00:00Z'))).toBe('raids-1');
    expect(phaseAt(phases, new Date('2027-03-01T00:00:00Z'))).toBe('raids-1');
  });
});

describe('hasOpened', () => {
  it('treats a source with no phase as open from launch of the data', () => {
    expect(hasOpened(phases, undefined, new Date('2026-09-01T00:00:00Z'))).toBe(true);
  });

  it('gates a later phase and opens it on the instant', () => {
    expect(hasOpened(phases, 'raids-1', new Date('2026-12-08T23:59:00Z'))).toBe(false);
    expect(hasOpened(phases, 'raids-1', new Date('2026-12-09T00:00:00Z'))).toBe(true);
  });

  it('never opens the literal "later" (contract 10.4: an unknown date)', () => {
    expect(PHASE_LATER).toBe('later');
    expect(hasOpened(phases, PHASE_LATER, new Date('2099-01-01T00:00:00Z'))).toBe(false);
  });

  it('treats a phase nobody has heard of as not open, rather than as open', () => {
    expect(hasOpened(phases, 'season-of-mastery', new Date('2030-01-01T00:00:00Z'))).toBe(false);
  });
});

describe('labels', () => {
  it('names each phase from the display table, not from the data file', () => {
    expect(phaseLabel('raids-1')).toBe('First raids');
    expect(phaseLabel(PHASE_LATER)).toBe('Later');
    expect(phaseLabel('nope')).toBe('nope');
  });

  it('dates the phases that have a date and nothing else', () => {
    expect(openDateLabel(phases, 'raids-1')).toBe('9 December 2026');
    expect(openDateLabel(phases, 'pre-beta')).toBe('');
    expect(openDateLabel(phases, PHASE_LATER)).toBe('');
    expect(phaseStart(phases, 'pre-beta')).toBeNull();
    expect(BUILT_IN_PHASES).toHaveLength(4);
  });

  it('treats the Go zero time api/internal/phase marshals as "no start", not a real date', () => {
    // task-2-report.md: the checked-in phases.json carries pre-beta's start as the literal
    // wire value of a Go zero time.Time ("0001-01-01T00:00:00Z"), not "". Any later reader
    // of phaseStart/openDateLabel must treat that string the same way it treats "".
    const preBeta = phases.find((phase) => phase.name === 'pre-beta');
    expect(preBeta?.start).toBe('0001-01-01T00:00:00Z');
  });
});

describe('fetchPhases', () => {
  it('prefers GET /v1/phases at runtime', async () => {
    api.route({
      method: 'GET',
      pattern: /\/v1\/phases$/,
      respond: () =>
        new Response(
          JSON.stringify({
            ok: true,
            request_id: 'r',
            error: null,
            data: { phases: [{ name: 'raids-2', start: '2027-03-01T00:00:00Z' }] },
          }),
          { status: 200, headers: { 'content-type': 'application/json' } },
        ),
    });
    await expect(fetchPhases(TEST_API)).resolves.toEqual([
      { name: 'raids-2', start: '2027-03-01T00:00:00Z' },
    ]);
  });

  it('falls back to the build-time table rather than leaving the gate unknown', async () => {
    api.route({
      method: 'GET',
      pattern: /\/v1\/phases$/,
      respond: () => new Response(null, { status: 500 }),
    });
    await expect(fetchPhases(TEST_API)).resolves.toEqual(BUILT_IN_PHASES);
  });
});
