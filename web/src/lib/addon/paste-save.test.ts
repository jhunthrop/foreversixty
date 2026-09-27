// web/src/lib/addon/paste-save.test.ts
import { describe, expect, it } from 'vitest';
import { missingPasteSaveField } from './paste-save';

describe('missingPasteSaveField', () => {
  it('names the first empty field in form order', () => {
    expect(missingPasteSaveField({ name: '', region: '', ruleset: '' })).toBe('name');
    expect(missingPasteSaveField({ name: '  ', region: 'us', ruleset: 'pvp' })).toBe('name');
    expect(missingPasteSaveField({ name: 'Bow Jackzon', region: '', ruleset: 'pvp' })).toBe('region');
    expect(missingPasteSaveField({ name: 'Bow Jackzon', region: 'us', ruleset: '' })).toBe('ruleset');
  });

  it('is null once every field is filled', () => {
    expect(missingPasteSaveField({ name: 'Bow Jackzon', region: 'us', ruleset: 'pvp' })).toBeNull();
  });
});
