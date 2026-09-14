import { describe, expect, it } from 'vitest';
import {
  OFFLINE_PROMPT,
  RATE_LIMITED,
  SUBSCRIBE_FAILED,
  SUBSCRIBED,
  subscribeMessageFor,
} from '../lib/subscribe';

describe('subscribeMessageFor', () => {
  it('accepts a 202 and says what happens next', () => {
    expect(subscribeMessageFor(202, null)).toEqual({ kind: 'done', message: SUBSCRIBED });
  });

  it('passes the API field message through on a 400', () => {
    expect(subscribeMessageFor(400, 'enter a valid email address')).toEqual({
      kind: 'error',
      message: 'enter a valid email address',
    });
  });

  it('uses fixed wording when rate limited', () => {
    expect(subscribeMessageFor(429, 'slow down')).toEqual({ kind: 'error', message: RATE_LIMITED });
  });

  it('offers the Discord fallback when the API cannot be reached', () => {
    expect(subscribeMessageFor(0, null)).toEqual({ kind: 'offline', message: OFFLINE_PROMPT });
    expect(subscribeMessageFor(503, null)).toEqual({ kind: 'offline', message: OFFLINE_PROMPT });
  });

  it('falls back to a generic error for anything else', () => {
    expect(subscribeMessageFor(418, null)).toEqual({ kind: 'error', message: SUBSCRIBE_FAILED });
  });
});
