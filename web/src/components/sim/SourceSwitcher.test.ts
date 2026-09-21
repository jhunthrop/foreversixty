// web/src/components/sim/SourceSwitcher.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { currentCharacterCopy } from '../../lib/current-character-copy';
import SourceSwitcher from './SourceSwitcher.svelte';

const requiredProps = {
  busy: false,
  message: null,
  signedIn: false,
  onaddon: () => {},
  onbuild: () => {},
  onfight: () => {},
  onsignin: () => {},
};

describe('SourceSwitcher', () => {
  it('links to /addon, with the shared "get the addon" copy', () => {
    const { body } = render(SourceSwitcher, { props: requiredProps });
    expect(body).toContain('href="/addon"');
    expect(body).toContain(currentCharacterCopy.getTheAddon);
  });

  // rowLink is the same 44px hit target CurrentCharacterChip.svelte's own nav links use
  // (CurrentCharacterChip.test.ts's "keeps the nav link colour" test), so a one-line note
  // still gives the anchor a real touch target on phone.
  it('gives the /addon link a 44px hit target', () => {
    const { body } = render(SourceSwitcher, { props: requiredProps });
    const match = /<a[^>]*href="\/addon"[^>]*>/.exec(body);
    if (match === null) throw new Error('no /addon anchor rendered');
    expect(match[0]).toContain('min-h-11');
  });
});
