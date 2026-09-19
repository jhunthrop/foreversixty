import { describe, expect, it } from 'vitest';
import { simCopy } from './copy';
import { executionHref, executionLabel, executionTitle } from './execution';

describe('executionLabel', () => {
  it('is a whole percentage', () => {
    expect(executionLabel(0.92)).toBe('92%');
    expect(executionLabel(1)).toBe('100%');
    expect(executionLabel(0.014)).toBe('1%');
  });

  it('is an em dash when the fight has no score', () => {
    expect(executionLabel(null)).toBe('—');
  });

  it('shows a score above the sim without clipping it, since the API already clamped', () => {
    expect(executionLabel(1.14)).toBe('114%');
  });
});

describe('executionTitle', () => {
  it('explains an absent score as a normal state, not a failure', () => {
    expect(executionTitle(null)).toBe(simCopy.executionUnscored);
    expect(simCopy.executionUnscored).toBe(
      'Not scored yet: this spec is not validated, or the fight predates scoring.',
    );
  });

  it('says what the number means when there is one', () => {
    expect(executionTitle(0.92)).toBe('92% of what this gear can do, simulated');
  });
});

describe('executionHref', () => {
  it('opens compare mode on that exact fight', () => {
    expect(executionHref('fixture2abcd', 2)).toBe('/sim?source=fight&ref=fixture2abcd%3A2&mode=compare');
  });
});
