// web/src/components/CurrentCharacterBar.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import CurrentCharacterBar from './CurrentCharacterBar.svelte';
import { currentCharacterCopy } from '../lib/current-character-copy';

// readCurrent() reads localStorage via `window`, which svelte/server's render() has no
// DOM for; CurrentCharacterBar's own $effect (where readCurrent runs) never fires during
// SSR (same reasoning Account.svelte's header comment gives for its own $effects), so the
// component always SSRs with current === null. That is exactly the case this test needs.

describe('CurrentCharacterBar', () => {
  it('shows the no-character line by default', () => {
    const { body } = render(CurrentCharacterBar, { props: {} });
    expect(body).toContain(currentCharacterCopy.noCharacterLine);
  });

  it('renders nothing when compact and there is no current character', () => {
    const { body } = render(CurrentCharacterBar, { props: { compact: true } });
    expect(body).not.toContain(currentCharacterCopy.noCharacterLine);
    expect(body).not.toContain('data-testid="current-character-chip"');
  });
});

describe('CurrentCharacterBar spine mode', () => {
  it('renders the signed-out, no-session line with no pointer (SSR: session/pointer both null)', () => {
    const { body } = render(CurrentCharacterBar, { props: { spine: true } });
    expect(body).toContain(currentCharacterCopy.barSignedOutLine);
    expect(body).toContain('/login');
    expect(body).toContain('/setup#paste');
  });

  it('reserves the same CHIP_HEIGHT classes as the plain chip', () => {
    const { body } = render(CurrentCharacterBar, { props: { spine: true } });
    expect(body).toContain('h-[88px]');
    expect(body).toContain('md:h-11');
  });
});

// L8.4 fix round 2 item 7: on a profile the character is the page's subject, so the visitor's
// sign-in band sits below the Forever Score card at every width, not above the identity. The
// profile orders its flex column: a signed-in chip -4, the header -3, the score card -2, this
// band -1, the rest in page order.
describe('CurrentCharacterBar on a profile', () => {
  it('orders the signed-out band after the score card at every width, not on phones alone', () => {
    const { body } = render(CurrentCharacterBar, { props: { spine: true, profile: true } });
    const bar = /<div class="([^"]*)"[^>]*data-testid="current-character-bar"/.exec(body)?.[1] ?? '';
    expect(bar.split(' ')).toContain('-order-1');
    expect(bar).not.toContain('max-md:-order-1');
  });
});
