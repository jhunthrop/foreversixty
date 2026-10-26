import { describe, expect, it } from 'vitest';
import { specString } from '../../test-support/spec-table';
import { discordText, yourNightText } from './night-copy';

const SPEC = new URL('../../../../design/specs/2026-10-27-guild-attendance.md', import.meta.url);
const sentence = (index: number, fill: Record<string, string>): string =>
  specString(SPEC, '### 6.1 On the page', 'Your night sentences', index, fill);

describe('yourNightText, verbatim from spec 6.1', () => {
  it('writes every sentence of the spec row', () => {
    expect(yourNightText.present('In', 'DPS', 11, 11)).toBe(
      sentence(0, { State: 'In', Role: 'DPS', n: '11', t: '11' }),
    );
    expect(yourNightText.late('Late', 'Healer', 22, 9, 11)).toBe(
      sentence(1, { State: 'Late', Role: 'Healer', m: '22', n: '9', t: '11' }),
    );
    expect(yourNightText.cameAs('In', 'Pyrewisp', 'Grimtotem', 8, 11)).toBe(
      sentence(2, { Name: 'Pyrewisp', Alt: 'Grimtotem', n: '8', t: '11' }),
    );
    expect(yourNightText.notInLog('In', 'DPS')).toBe(sentence(3, { Role: 'DPS' }));
    expect(yourNightText.unsigned(4, 11)).toBe(sentence(4, { n: '4', t: '11' }));
    expect(yourNightText.joined('Bench', 6, 11)).toBe(sentence(5, { n: '6', t: '11' }));
  });

  it('drops the role clause when there is no role', () => {
    expect(yourNightText.present('In', null, 1, 2)).toBe(
      'Signed In. Present from the first pull, 1 of 2 pulls.',
    );
  });

  it('labels the overflow line', () => {
    expect(discordText.more(15)).toBe('+15 more');
  });
});
