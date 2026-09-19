// web/src/lib/sim/history.test.ts
import { describe, expect, it } from 'vitest';
import { simCopy } from './copy';
import { KIND_FILTERS, headlineOf, kindOf, titleOf } from './history';
import type { SimListRow } from './types';

const row = (over: Partial<SimListRow> = {}): SimListRow => ({
  sim_id: 'aaaaaaaaaaaa',
  spec: 'warrior-fury',
  dps: 1204.5,
  engine_version: 'edc0c8e9a',
  created_at: '2026-09-19T10:00:00Z',
  title: '',
  ...over,
});

describe('kindOf', () => {
  it('is the row’s kind', () => {
    expect(kindOf(row({ kind: 'gear' }))).toBe('gear');
  });

  it('is run for a row saved before the kind column existed, and for one the API mislabels', () => {
    expect(kindOf(row())).toBe('run');
    expect(kindOf(row({ kind: 'nonsense' }))).toBe('run');
  });
});

describe('headlineOf', () => {
  it('is the API’s own sentence when it sends one', () => {
    expect(headlineOf(row({ headline: '+41 DPS from Vis’kag' }))).toBe('+41 DPS from Vis’kag');
  });

  it('falls back to the figure, so a row before the API composed headlines still says something', () => {
    expect(headlineOf(row())).toBe('1,205 DPS');
    expect(headlineOf(row({ headline: '' }))).toBe('1,205 DPS');
  });
});

describe('titleOf', () => {
  it('is the player’s title, and the spec when they never gave one', () => {
    expect(titleOf(row({ title: 'Pre-raid' }))).toBe('Pre-raid');
    expect(titleOf(row())).toBe('Fury');
  });
});

describe('KIND_FILTERS', () => {
  it('is “all” and the contract’s five kinds, each named', () => {
    expect(KIND_FILTERS).toEqual(['all', 'run', 'gear', 'talents', 'drops', 'weights']);
    for (const filter of KIND_FILTERS) {
      expect(simCopy.kindLabel[filter], filter).toBeTruthy();
    }
  });
});
