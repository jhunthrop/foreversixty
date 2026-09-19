import { describe, expect, it } from 'vitest';
import { defaultSimState, parseSimState, simIdFrom, simSearch, withSimState } from './url';

describe('parseSimState', () => {
  it('is the default for an empty query', () => {
    expect(parseSimState('')).toEqual(defaultSimState());
    expect(defaultSimState()).toEqual({ source: '', ref: '', mode: 'sim', fight: '' });
  });

  it('reads every parameter the page writes', () => {
    expect(parseSimState('?source=fight&ref=fixture2abcd%3A2&mode=compare&fight=fixture2abcd%3A3')).toEqual({
      source: 'fight',
      ref: 'fixture2abcd:2',
      mode: 'compare',
      fight: 'fixture2abcd:3',
    });
  });

  it('drops a source or mode it does not recognise rather than trusting the URL', () => {
    expect(parseSimState('?source=rumour&mode=hack').source).toBe('');
    expect(parseSimState('?source=rumour&mode=hack').mode).toBe('sim');
  });

  it('caps a ref at a length no real id reaches, since the query is attacker-controlled', () => {
    expect(parseSimState(`?ref=${'x'.repeat(5000)}`).ref).toHaveLength(0);
  });
});

describe('simSearch', () => {
  it('writes only what differs from the default, so a plain /sim stays clean', () => {
    expect(simSearch(defaultSimState())).toBe('');
    expect(simSearch({ ...defaultSimState(), source: 'build', ref: 'bld1' })).toBe('?source=build&ref=bld1');
  });

  it('round-trips every state it writes', () => {
    const state = {
      source: 'fight' as const,
      ref: 'fixture2abcd:2',
      mode: 'compare' as const,
      fight: 'fixture2abcd:3',
    };
    expect(parseSimState(simSearch(state))).toEqual(state);
  });
});

describe('withSimState', () => {
  it('returns a new state and never mutates the old one', () => {
    const base = defaultSimState();
    const next = withSimState(base, { mode: 'compare' });
    expect(next.mode).toBe('compare');
    expect(base.mode).toBe('sim');
  });
});

describe('simIdFrom', () => {
  it('reads the twelve base32 characters out of /sim/<id>', () => {
    expect(simIdFrom('/sim/simfixtureab')).toBe('simfixtureab');
    expect(simIdFrom('/sim/simfixtureab/')).toBe('simfixtureab');
  });

  it('is empty for /sim, /sim/specs and anything else', () => {
    expect(simIdFrom('/sim')).toBe('');
    expect(simIdFrom('/sim/specs')).toBe('');
    expect(simIdFrom('/sim/TOOSHORT')).toBe('');
    expect(simIdFrom('/reports/fixture2abcd')).toBe('');
  });
});
