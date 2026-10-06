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

    it('renders the sign-in action as a secondary button, not a second primary, and the email-link alternative as text', () => {
      const { body } = render(SourceSwitcher, { props });
      expect(body).toContain('data-testid="sim-signin"');
      expect(body).toContain(landingCopy.signInWithBattlenet);
      // One primary per view (design system): ColdPasteHero.svelte's Run is the gold
      // PRIMARY_BUTTON whenever this hero renders, so this button takes the same
      // secondary look the non-hero branch's identical button already uses.
      const match = /<button[^>]*data-testid="sim-signin"[^>]*>/.exec(body);
      if (match === null) throw new Error('no sim-signin button rendered');
      expect(match[0]).not.toContain('bg-gold');
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

  // Persona review 2026-10-06 bis: SimView.svelte passes `showAddonCard={false}` whenever
  // ColdPasteHero.svelte already owns the one paste surface on /sim -- one primary paste
  // box per view, not two. The bulk tool pages never pass the prop, so it defaults true
  // and their four-card (or hero-plus-three) grid is unchanged (every test above).
  describe('showAddonCard: false', () => {
    it('drops the addon card from the hero row, keeping From a build and From a logged fight', () => {
      const { body } = render(SourceSwitcher, {
        props: { ...requiredProps, heroSignIn: true, showAddonCard: false },
      });
      expect(body).not.toContain('data-testid="sim-addon-input"');
      expect(body).toContain('data-testid="sim-build-input"');
      expect(body).toContain('data-testid="sim-fight-input"');
    });

    it('drops the addon card from the plain grid too', () => {
      const { body } = render(SourceSwitcher, {
        props: { ...requiredProps, showAddonCard: false },
      });
      expect(body).not.toContain('data-testid="sim-addon-input"');
      expect(body).toContain('data-testid="sim-build-input"');
      expect(body).toContain('data-testid="sim-fight-input"');
    });
  });
});
