import { describe, expect, it } from 'vitest';
import {
  captionFields,
  exampleCaption,
  exampleHref,
  isPlaceholder,
  parseServerRunExample,
  serverRunExample,
  type CaptionTemplates,
} from './premium-examples';

const VALID = {
  simId: 'abc123',
  characterName: 'Kraggor',
  level: 60,
  specLabel: 'Fury Warrior',
  preset: true,
  placeholder: false,
  combinations: 4812,
  ranAt: '2026-10-27T18:00:00Z',
};
const T: CaptionTemplates = {
  preset: (f) => `P|${f.level}|${f.specLabel}|${f.combinations}|${f.ranAt}`,
  ownGear: (f) => `O|${f.characterName}|${f.specLabel}|${f.combinations}|${f.ranAt}`,
};

describe('premium examples', () => {
  it('parses a valid file', () => expect(parseServerRunExample({ serverRun: VALID })).toEqual(VALID));
  it.each(['simId', 'level', 'specLabel', 'combinations', 'ranAt', 'placeholder'])(
    'rejects a file missing %s',
    (k) => {
      const rest: Record<string, unknown> = { ...VALID };
      delete rest[k];
      expect(() => parseServerRunExample({ serverRun: rest })).toThrow(k);
    },
  );
  it('rejects a missing or non-boolean preset', () => {
    const rest: Record<string, unknown> = { ...VALID };
    delete rest.preset;
    expect(() => parseServerRunExample({ serverRun: rest })).toThrow('preset');
    expect(() => parseServerRunExample({ serverRun: { ...VALID, preset: 'true' } })).toThrow('preset');
    expect(() => parseServerRunExample({ serverRun: { ...VALID, preset: null } })).toThrow('preset');
  });
  it('rejects an empty characterName on an own-gear run', () => {
    expect(() =>
      parseServerRunExample({ serverRun: { ...VALID, preset: false, characterName: '' } }),
    ).toThrow('characterName');
  });
  it('rejects an out-of-range level and a bad instant', () => {
    expect(() => parseServerRunExample({ serverRun: { ...VALID, level: 61 } })).toThrow('level');
    expect(() => parseServerRunExample({ serverRun: { ...VALID, ranAt: 'soon' } })).toThrow('ranAt');
  });
  it('rejects a zero combination count, placeholder or not', () => {
    expect(() => parseServerRunExample({ serverRun: { ...VALID, combinations: 0 } })).toThrow('combinations');
    expect(() =>
      parseServerRunExample({ serverRun: { ...VALID, placeholder: true, combinations: 0 } }),
    ).toThrow('combinations');
  });
  it.each(['Oct 27 2026', '2026', '2026-10-27'])('rejects ranAt %s as not an RFC 3339 instant', (ranAt) => {
    expect(() => parseServerRunExample({ serverRun: { ...VALID, ranAt } })).toThrow('ranAt');
  });
  it.each([{}, null, 'x'])('rejects a file without serverRun with a dedicated message', (file) => {
    expect(() => parseServerRunExample(file)).toThrow('premium-examples.json: serverRun is missing');
  });
  it('picks the preset template when preset is true', () =>
    expect(exampleCaption(VALID, T)).toBe('P|60|Fury Warrior|4,812|27 Oct 2026'));
  it('picks the own-gear template when preset is false', () =>
    expect(exampleCaption({ ...VALID, preset: false }, T)).toBe('O|Kraggor|Fury Warrior|4,812|27 Oct 2026'));
  it('formats ranAt in UTC across the date line', () =>
    expect(captionFields({ ...VALID, ranAt: '2026-10-27T23:30:00-07:00' }).ranAt).toBe('28 Oct 2026'));
  it('flags the placeholder from the explicit flag only', () => {
    expect(isPlaceholder({ ...VALID, placeholder: true })).toBe(true);
    expect(isPlaceholder({ ...VALID, simId: 'placeholder' })).toBe(false);
    expect(isPlaceholder(VALID)).toBe(false);
  });
  it('separates thousands past a million', () =>
    expect(captionFields({ ...VALID, combinations: 1234567 }).combinations).toBe('1,234,567'));
  it('the shipped file parses', () => expect(() => serverRunExample).not.toThrow());
  it('links to the saved result', () => expect(exampleHref(VALID)).toBe('/sim/abc123'));
});
