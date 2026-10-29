import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { EMPTY_NAMES } from '../../lib/coach/names';
import { coachTeaser } from '../../test-support/coach-fixtures';
import CoachTeaser from './CoachTeaser.svelte';

const names = { ...EMPTY_NAMES, iconUrl: () => '/data/b/icons/bt.webp' };

describe('CoachTeaser', () => {
  it('count, first title with its icon, and an upgrade link', () => {
    const { body } = render(CoachTeaser, { props: { answer: coachTeaser(), names } });
    expect(body).toContain('3 things to fix');
    expect(body).toContain('Bloodthirst: 9 casts lost');
    expect(body).not.toContain('First:');
    expect(body).toContain('/data/b/icons/bt.webp');
    expect(body).toContain('href="/premium#coach"');
  });
  it('one finding is singular', () => {
    const { body } = render(CoachTeaser, { props: { answer: coachTeaser({ finding_count: 1 }), names } });
    expect(body).toContain('1 thing to fix');
  });
  it('zero findings says nothing to fix and offers no button', () => {
    const { body } = render(CoachTeaser, {
      props: { answer: coachTeaser({ finding_count: 0, fix_first: null }), names },
    });
    expect(body).toContain('Nothing to fix on this pull that the log can show.');
    expect(body).not.toContain('/premium');
  });
});
