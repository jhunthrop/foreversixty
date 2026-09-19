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
  // entries, plus a version 2 code's optional bags, bank, named sets and loadouts), so it
  // gets its own, larger bound rather than sharing MAX_REF -- a real code comfortably fits
  // under it, and a query well past MAX_CODE_LENGTH still cannot reach the decoder.
  it('carries a real FS1 code whole, and caps a code far past what one reaches', () => {
    const code = `FS1:1.15.9.69722:warrior:orc:${'1'.repeat(40)}/0/0:${'head=12640,'.repeat(17).slice(0, -1)}`;
    expect(code.length).toBeGreaterThan(128);
    expect(code.length).toBeLessThan(16_384);
    expect(parseSimState(`?code=${encodeURIComponent(code)}`).code).toBe(code);
    expect(parseSimState(`?code=${'x'.repeat(20_000)}`).code).toHaveLength(0);
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

  // Fix round 1: `decodeRequestParam` used to accept any non-null, non-array object, so a
  // crafted `/sim?req=e30` ('e30' is base64url for '{}') passed it and then threw inside
  // `settingsFromRequest` on the first field it reached for -- an unhandled rejection for
  // anyone who followed the link. These pin the shape check that stops that at the door.
  it('refuses an empty object -- the crafted /sim?req=e30 the review demonstrated', () => {
    expect(decodeRequestParam('e30')).toBeNull();
  });

  it('refuses a partially-populated object that has some, but not all, required fields', () => {
    // character and encounter alone were enough to pass the old check (a non-null, non-array
    // object); every other required field -- engine_version, spec, source, iterations,
    // random_seed -- is still missing.
    const partial = { character: request.character, encounter: request.encounter } as unknown as SimRequest;
    const encoded = encodeRequestParam(partial);
    expect(encoded).not.toBeNull();
    expect(decodeRequestParam(encoded!)).toBeNull();
  });

  it('accepts a request whose encoded size lands exactly on the budget, and refuses one byte more', () => {
    // Grows a filler field one character at a time until the real, base64url-encoded
    // output lands exactly on MAX_REQUEST_PARAM -- the boundary an off-by-one in either
    // encodeRequestParam's `>` comparison or decodeRequestParam's own length gate would get
    // wrong, unlike the 20,000-buffs test above, which overshoots by a wide margin.
    let filler = '';
    let padded = { ...request, character: { ...request.character, name: filler } };
    let encoded = encodeRequestParam(padded);
    while (encoded !== null && encoded.length < MAX_REQUEST_PARAM) {
      filler += 'x';
      padded = { ...request, character: { ...request.character, name: filler } };
      encoded = encodeRequestParam(padded);
    }
    expect(encoded).not.toBeNull();
    expect(encoded!.length).toBe(MAX_REQUEST_PARAM);
    expect(decodeRequestParam(encoded!)).toEqual(padded);

    // The length gate itself, not the content: one character past the budget is refused
    // before decodeRequestParam even tries to parse anything.
    expect(decodeRequestParam(`${encoded!}x`)).toBeNull();

    // The real encoder agrees: one character more of input pushes the actual encoded
    // output past the budget too.
    filler += 'x';
    expect(encodeRequestParam({ ...request, character: { ...request.character, name: filler } })).toBeNull();
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
