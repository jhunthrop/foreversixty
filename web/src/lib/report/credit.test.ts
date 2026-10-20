import { describe, expect, it } from 'vitest';
import { NOT_CREDITED_TITLE, reportCredit } from './credit';

describe('reportCredit', () => {
  it('names the logging character', () => {
    expect(reportCredit({ logging_character: 'us/normal/bow-jackzon' })).toEqual({
      text: 'Credited to Bow Jackzon',
      title: undefined,
      credited: true,
    });
    expect(reportCredit({ logging_character: { name: 'Bow Jackzon' } }).text).toBe('Credited to Bow Jackzon');
  });
  it('says so, and how to fix it, when there is none', () => {
    for (const logging_character of [undefined, null, '', { name: '' }]) {
      expect(reportCredit({ logging_character })).toEqual({
        text: 'Not credited to a character',
        title: NOT_CREDITED_TITLE,
        credited: false,
      });
    }
  });
});
