import { readFileSync } from 'node:fs';
import { describe, expect, it } from 'vitest';
import { specCell, specQuoted } from '../../test-support/spec-table';
import { coachCopy, fill, type CoachCopyKey } from './copy';

const SPEC = new URL('../../../../design/specs/2026-10-28-log-coach.md', import.meta.url);
const HEADING = '### 8.2 Web copy';

function specKeys(): string[] {
  const text = readFileSync(SPEC, 'utf8');
  const start = text.indexOf(HEADING);
  const end = text.indexOf('### 8.3', start);
  return text
    .slice(start, end)
    .split('\n')
    .filter((line) => line.startsWith('| ') && !line.startsWith('| Key') && !line.startsWith('|---'))
    .map((line) => (line.split('|')[1] ?? '').trim());
}

describe('coachCopy against the experience spec', () => {
  it('holds exactly the keys of section 8.2', () => {
    expect(Object.keys(coachCopy).sort()).toEqual(specKeys().sort());
  });

  it.each(specKeys())('%s is the spec string word for word', (key) => {
    const value = coachCopy[key as CoachCopyKey];
    expect(value).toBe(specCell(SPEC, HEADING, key)[0]);
    expect(specQuoted(SPEC, value)).toBe(true);
  });

  it('fills a row line', () => {
    expect(fill(coachCopy.rowFree, { n: 3, title: 'Bloodthirst: 9 casts lost' })).toBe(
      '3 things to fix · Bloodthirst: 9 casts lost first · Read the full review with Premium',
    );
  });

  it('fill leaves an unknown slot visible', () => {
    expect(fill('{a} and {b}', { a: 'x' })).toBe('x and {b}');
  });

  it('no exclamation marks and never "heal more"', () => {
    expect(Object.values(coachCopy).join(' ')).not.toMatch(/!|heal more/i);
  });
});
