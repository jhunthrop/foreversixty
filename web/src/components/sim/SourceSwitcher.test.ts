// web/src/components/sim/SourceSwitcher.test.ts
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { currentCharacterCopy } from '../../lib/current-character-copy';
import { landingCopy } from '../../lib/sim/landing-copy';
import { simCopy } from '../../lib/sim/copy';
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
    expect(body).toContain('href="/setup"');
    expect(body).toContain(currentCharacterCopy.getTheAddon);
  });

  // rowLink is the same 44px hit target CurrentCharacterChip.svelte's own nav links use
  // (CurrentCharacterChip.test.ts's "keeps the nav link colour" test), so a one-line note
  // still gives the anchor a real touch target on phone.
  it('gives the /addon link a 44px hit target', () => {
    const { body } = render(SourceSwitcher, { props: requiredProps });
    const match = /<a[^>]*href="\/setup"[^>]*>/.exec(body);
    if (match === null) throw new Error('no /addon anchor rendered');
    expect(match[0]).toContain('min-h-11');
  });

  // 2026-09-26 layout pass, Finding 2: `heroSignIn` promotes the sign-in card to a
  // full-width hero above the three paste cards, and carries the one copy of the DPS-only
  // restriction as its own caption -- the switcher's separate duplicate line is gone.
  describe('heroSignIn, signed out', () => {
    const props = { ...requiredProps, heroSignIn: true };

    it('renders the intro line and the static example ahead of the hero card', () => {
      const { body } = render(SourceSwitcher, { props });
      expect(body).toContain('data-testid="sim-intro-line"');
      expect(body).toContain(landingCopy.introLine);
      expect(body).toContain('data-testid="sim-example-result"');
      const introIndex = body.indexOf('data-testid="sim-intro-line"');
      const accountIndex = body.indexOf('data-testid="sim-account-card"');
      expect(introIndex).toBeGreaterThan(-1);
      expect(accountIndex).toBeGreaterThan(introIndex);
    });

    it('renders the restriction once, as the hero card’s own caption, never the old duplicate', () => {
      const { body } = render(SourceSwitcher, { props });
      expect(body).toContain('data-testid="sim-scope-note"');
      expect(body).toContain(landingCopy.scopeCaveat);
      expect(body).not.toContain('data-testid="sim-sources-scope-note"');
    });

    it('renders the one PRIMARY_BUTTON sign-in action and the email-link alternative as text', () => {
      const { body } = render(SourceSwitcher, { props });
      expect(body).toContain('data-testid="sim-signin"');
      expect(body).toContain(landingCopy.signInWithBattlenet);
      expect(body).toContain('bg-gold');
      expect(body).toContain(landingCopy.emailLinkInstead);
      expect(body).toContain('href="/login"');
    });

    it('puts the hero card ahead of the three paste cards', () => {
      const { body } = render(SourceSwitcher, { props });
      const accountIndex = body.indexOf('data-testid="sim-account-card"');
      const addonIndex = body.indexOf('data-testid="sim-addon-input"');
      expect(accountIndex).toBeGreaterThan(-1);
      expect(addonIndex).toBeGreaterThan(accountIndex);
    });
  });

  it('ignores heroSignIn once signed in -- the plain grid and its own scope line stay', () => {
    const { body } = render(SourceSwitcher, {
      props: { ...requiredProps, signedIn: true, heroSignIn: true, hasCharacters: true },
    });
    expect(body).not.toContain('data-testid="sim-intro-line"');
    expect(body).not.toContain('data-testid="sim-example-result"');
    expect(body).toContain('data-testid="sim-sources-scope-note"');
    expect(body).toContain(simCopy.scopeNote);
    expect(body).toContain('data-testid="sim-back-to-characters"');
  });
});
