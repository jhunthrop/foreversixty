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
