// @vitest-environment jsdom
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';
import * as layoutReserve from '../layout-reserve';
import {
  clearMeGuildCount,
  clearPremiumHint,
  ME_GUILD_COUNT_KEY,
  PREMIUM_SOURCE_KEY,
  readMeGuildCount,
  readPremiumHint,
  writeMeGuildCount,
  writePremiumHint,
} from './premium-hint';

beforeEach(() => {
  localStorage.clear();
  delete document.documentElement.dataset.premium;
});
afterEach(() => vi.restoreAllMocks());

describe('writePremiumHint', () => {
  it('stores own for own_then_seat', () => {
    writePremiumHint('own_then_seat', document);
    expect(localStorage.getItem(PREMIUM_SOURCE_KEY)).toBe('own');
    expect(document.documentElement.dataset.premium).toBe('own');
  });

  it.each(['seat', 'grant'] as const)('stores %s as is', (source) => {
    writePremiumHint(source, document);
    expect(localStorage.getItem(PREMIUM_SOURCE_KEY)).toBe(source);
    expect(readPremiumHint()).toBe(source);
  });

  it.each(['fallback', 'none'] as const)('%s removes the key and the attribute', (source) => {
    localStorage.setItem(PREMIUM_SOURCE_KEY, 'grant');
    document.documentElement.dataset.premium = 'grant';
    writePremiumHint(source, document);
    expect(localStorage.getItem(PREMIUM_SOURCE_KEY)).toBeNull();
    expect(document.documentElement.dataset.premium).toBeUndefined();
  });
});

describe('readPremiumHint', () => {
  it('removes a value that is not one of the three words', () => {
    localStorage.setItem(PREMIUM_SOURCE_KEY, 'premium');
    expect(readPremiumHint()).toBeNull();
    expect(localStorage.getItem(PREMIUM_SOURCE_KEY)).toBeNull();
  });

  it('answers null when nothing is stored', () => expect(readPremiumHint()).toBeNull());
});

describe('clearPremiumHint', () => {
  it('drops the key and the attribute', () => {
    writePremiumHint('own', document);
    clearPremiumHint(document);
    expect(localStorage.getItem(PREMIUM_SOURCE_KEY)).toBeNull();
    expect(document.documentElement.dataset.premium).toBeUndefined();
  });
});

describe('the guild count', () => {
  it('round-trips a stored count', () => {
    writeMeGuildCount(3);
    expect(readMeGuildCount()).toBe(3);
  });

  it('reads a stored zero as one (spec §5.7: at least one skeleton card) and clearMeGuildCount forgets it', () => {
    writeMeGuildCount(0);
    expect(readMeGuildCount()).toBe(1);
    clearMeGuildCount();
    expect(readMeGuildCount()).toBe(2);
  });

  it('clamps to 1..4, else 2', () => {
    localStorage.setItem(ME_GUILD_COUNT_KEY, '9');
    expect(readMeGuildCount()).toBe(4);
    localStorage.setItem(ME_GUILD_COUNT_KEY, '-3');
    expect(readMeGuildCount()).toBe(1);
    localStorage.setItem(ME_GUILD_COUNT_KEY, 'x');
    expect(readMeGuildCount()).toBe(2);
    localStorage.removeItem(ME_GUILD_COUNT_KEY);
    expect(readMeGuildCount()).toBe(2);
  });
});

describe('a throwing storage', () => {
  beforeEach(() => {
    vi.spyOn(layoutReserve, 'storage').mockImplementation(() => {
      throw new Error('blocked');
    });
  });

  it('reads answer null and 2, writes are no-ops', () => {
    expect(readPremiumHint()).toBeNull();
    expect(readMeGuildCount()).toBe(2);
    expect(() => {
      writePremiumHint('own', document);
      writePremiumHint('none', document);
      clearPremiumHint(document);
      writeMeGuildCount(3);
    }).not.toThrow();
  });
});
