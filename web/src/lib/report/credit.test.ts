import { describe, expect, it } from 'vitest';
import { NOT_CREDITED, reportCredit } from './credit';

describe('reportCredit', () => {
  it('reads the API payload exactly: { key, name }', () => {
    const logging_character = JSON.parse('{"key":"us/pvp/bow-jackzon","name":"Bow Jackzon"}');
    expect(reportCredit({ logging_character })).toEqual({
      prefix: 'Credited to ',
      name: 'Bow Jackzon',
      href: '/character/us/pvp/bow-jackzon',
      credited: true,
    });
  });
  it('still accepts the older string shape', () => {
    expect(reportCredit({ logging_character: 'us/normal/bow-jackzon' })).toEqual({
      prefix: 'Credited to ',
      name: 'Bow Jackzon',
      href: '/character/us/normal/bow-jackzon',
      credited: true,
    });
  });
  it('shows the key last segment title-cased when the name is missing', () => {
    const credit = reportCredit({ logging_character: { key: 'us/pvp/bow-jackzon', name: '' } });
    expect(credit.name).toBe('Bow Jackzon');
    expect(credit.href).toBe('/character/us/pvp/bow-jackzon');
  });
  it('links nothing when only a name is known', () => {
    expect(reportCredit({ logging_character: { name: 'Bow Jackzon' } })).toMatchObject({
      name: 'Bow Jackzon',
      href: undefined,
      credited: true,
    });
  });
  it('says out loud that it is not credited, with the fix, when there is none', () => {
    for (const logging_character of [undefined, null, '', { key: '', name: '' }]) {
      expect(reportCredit({ logging_character })).toEqual({
        prefix: NOT_CREDITED,
        name: '',
        href: undefined,
        credited: false,
      });
    }
    expect(NOT_CREDITED).toBe('Not credited yet. Log out once in game; the companion credits it then.');
  });
});
