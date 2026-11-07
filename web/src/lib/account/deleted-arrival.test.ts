// @vitest-environment jsdom
// web/src/lib/account/deleted-arrival.test.ts
import { afterEach, describe, expect, it } from 'vitest';
import { DELETED_ARRIVAL_HREF, takeDeletedArrival } from './deleted-arrival';

afterEach(() => window.history.replaceState(null, '', '/'));

describe('takeDeletedArrival', () => {
  it('reads the flag once and takes it out of the address bar', () => {
    window.history.replaceState(null, '', DELETED_ARRIVAL_HREF);
    expect(takeDeletedArrival()).toBe(true);
    expect(window.location.pathname + window.location.search).toBe('/');
    expect(takeDeletedArrival()).toBe(false);
  });

  it('keeps the other parameters and the hash', () => {
    window.history.replaceState(null, '', '/?golink=x&deleted=1#feed');
    expect(takeDeletedArrival()).toBe(true);
    expect(window.location.search).toBe('?golink=x');
    expect(window.location.hash).toBe('#feed');
  });

  it('an ordinary visit is not an arrival and the URL is left alone', () => {
    window.history.replaceState(null, '', '/?band=30');
    expect(takeDeletedArrival()).toBe(false);
    expect(window.location.search).toBe('?band=30');
  });

  it('a forged value is dropped without a notice', () => {
    window.history.replaceState(null, '', '/?deleted=yes');
    expect(takeDeletedArrival()).toBe(false);
    expect(window.location.search).toBe('');
  });
});
