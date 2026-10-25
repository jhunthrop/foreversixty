import { describe, expect, it } from 'vitest';
import { historyMount } from './history-mount';

const base = { ownerKnown: true, own: true, entitled: true, hasScore: true };

describe('historyMount', () => {
  it('mounts the charts for the entitled owner of a scored character', () => {
    expect(historyMount(base)).toBe('charts');
  });

  it('says one sentence to the free owner of a scored character', () => {
    expect(historyMount({ ...base, entitled: false })).toBe('sentence');
  });

  it('mounts nothing for an owner with no score, whatever the plan', () => {
    expect(historyMount({ ...base, hasScore: false })).toBe('none');
    expect(historyMount({ ...base, hasScore: false, entitled: false })).toBe('none');
  });

  it('mounts nothing for a visitor', () => {
    expect(historyMount({ ...base, own: false })).toBe('none');
    expect(historyMount({ ...base, own: false, entitled: false })).toBe('none');
  });

  it('holds the place of the section until /v1/me has answered', () => {
    expect(historyMount({ ...base, ownerKnown: false, own: false, entitled: false })).toBe('pending');
  });

  it('holds nothing while pending for a character with no score', () => {
    expect(historyMount({ ...base, ownerKnown: false, hasScore: false })).toBe('none');
  });
});
