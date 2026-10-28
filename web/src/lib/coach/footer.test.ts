import { describe, expect, it } from 'vitest';
import { footerLine } from './footer';

const AT = '2026-10-27T12:00:00Z';

describe('footerLine', () => {
  it('names the bases, each without its own full stop, joined by "; "', () => {
    expect(footerLine(AT, ['Your log.', 'Top 25 Fury warriors on Onyxia (40).'])).toBe(
      'Reviewed 27 Oct · Based on: Your log; Top 25 Fury warriors on Onyxia (40).',
    );
  });

  it('with no basis reads "Reviewed {date}." and never an empty "Based on: ."', () => {
    expect(footerLine(AT, [])).toBe('Reviewed 27 Oct.');
    expect(footerLine(AT, ['', '  '])).toBe('Reviewed 27 Oct.');
  });
});
