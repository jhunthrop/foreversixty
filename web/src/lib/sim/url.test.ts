import { describe, expect, it } from 'vitest';
import envelope from '../../fixtures/sim/envelope-v2.json';
import type { SimRequest } from './types';
import {
  decodeRequestParam,
  defaultSimState,
  encodeRequestParam,
  MAX_REQUEST_PARAM,
  parseSimState,
  simIdFrom,
  simSearch,
  withSimState,
} from './url';

describe('parseSimState', () => {
  it('is the default for an empty query', () => {
    expect(parseSimState('')).toEqual(defaultSimState());
    expect(defaultSimState()).toEqual({
      source: '',
      ref: '',
      code: '',
      req: '',
      mode: 'sim',
      fight: '',
    });
  });

  it('reads every parameter the page writes', () => {
    expect(
      parseSimState(
        '?source=fight&ref=fixture2abcd%3A2&code=FS1%3A1.15.9.69722%3Awarrior%3Aorc%3A3%2F0%2F0%3A&mode=compare&fight=fixture2abcd%3A3',
      ),
    ).toEqual({
      source: 'fight',
      ref: 'fixture2abcd:2',
      code: 'FS1:1.15.9.69722:warrior:orc:3/0/0:',
      req: '',
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

  // An FS1 code is far longer than a ref (three tree strings, up to seventeen gear
  // entries), so it gets its own, larger bound rather than sharing MAX_REF -- a real code
  // comfortably fits under it, and a 3000-character query still cannot reach the decoder.
  it('carries a real FS1 code whole, and caps a code far past what one reaches', () => {
    const code = `FS1:1.15.9.69722:warrior:orc:${'1'.repeat(40)}/0/0:${'head=12640,'.repeat(17).slice(0, -1)}`;
    expect(code.length).toBeGreaterThan(128);
    expect(code.length).toBeLessThan(2048);
    expect(parseSimState(`?code=${encodeURIComponent(code)}`).code).toBe(code);
    expect(parseSimState(`?code=${'x'.repeat(3000)}`).code).toHaveLength(0);
  });
});

describe('simSearch', () => {
  it('writes only what differs from the default, so a plain /sim stays clean', () => {
    expect(simSearch(defaultSimState())).toBe('');
    expect(simSearch({ ...defaultSimState(), source: 'build', ref: 'bld1' })).toBe('?source=build&ref=bld1');
  });

  it('writes a code the same way it writes a ref', () => {
    expect(simSearch({ ...defaultSimState(), code: 'FS1:1.15.9.69722:warrior:orc:0/0/0:' })).toBe(
      '?code=FS1%3A1.15.9.69722%3Awarrior%3Aorc%3A0%2F0%2F0%3A',
    );
  });

  it('round-trips every state it writes', () => {
    const state = {
      source: 'fight' as const,
      ref: 'fixture2abcd:2',
      code: 'FS1:1.15.9.69722:warrior:orc:0/0/0:',
      req: '',
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

describe('a request in the URL', () => {
  const request = (envelope as unknown as { request: SimRequest }).request;

  it('round-trips a request through the query string', () => {
    const encoded = encodeRequestParam(request);
    expect(encoded).not.toBeNull();
    expect(decodeRequestParam(encoded!)).toEqual(request);
  });

  it('is URL-safe: no +, / or = to be mangled by a chat client', () => {
    expect(encodeRequestParam(request)!).toMatch(/^[A-Za-z0-9_-]+$/);
  });

  it('refuses a request past the budget rather than writing a link that will be cut', () => {
    const huge = {
      ...request,
      character: { ...request.character, buffs: Array.from({ length: 20_000 }, (_, i) => `b${i}`) },
    };
    expect(encodeRequestParam(huge)).toBeNull();
  });

  it('answers null for a value that is not a request, however it is malformed', () => {
    expect(decodeRequestParam('not-base64!!')).toBeNull();
    expect(decodeRequestParam(btoa('[1,2,3]').replaceAll('=', ''))).toBeNull();
    expect(decodeRequestParam('')).toBeNull();
  });

  it('parses and writes the req parameter beside the others', () => {
    const encoded = encodeRequestParam(request)!;
    const state = parseSimState(`?req=${encoded}`);
    expect(state.req).toBe(encoded);
    expect(simSearch({ ...defaultSimState(), req: encoded })).toBe(`?req=${encoded}`);
  });

  it('drops a req parameter past the budget instead of handing a truncated one to the decoder', () => {
    expect(parseSimState(`?req=${'a'.repeat(MAX_REQUEST_PARAM + 1)}`).req).toBe('');
  });
});
