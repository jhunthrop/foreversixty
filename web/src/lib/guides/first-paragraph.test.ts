// web/src/lib/guides/first-paragraph.test.ts
import { describe, expect, it } from 'vitest';
import { splitFirstParagraph } from './first-paragraph';

describe('splitFirstParagraph', () => {
  it('splits a single leading paragraph from the rest of the body', () => {
    const { first, rest } = splitFirstParagraph('Para one.\n\nPara two.\n\nPara three.');
    expect(first).toBe('Para one.');
    expect(rest).toBe('Para two.\n\nPara three.');
  });

  it('leaves rest empty when the whole body is one paragraph', () => {
    const { first, rest } = splitFirstParagraph('Only paragraph.');
    expect(first).toBe('Only paragraph.');
    expect(rest).toBe('');
  });

  it('trims leading/trailing whitespace off both halves', () => {
    const { first, rest } = splitFirstParagraph('\n\n  Para one.  \n\n  Para two.  \n\n');
    expect(first).toBe('Para one.');
    expect(rest).toBe('Para two.');
  });
});
