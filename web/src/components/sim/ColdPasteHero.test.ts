// web/src/components/sim/ColdPasteHero.test.ts
// `render()` from 'svelte/server' never runs client interaction, so the decode -> Run
// enables itself path lives in tests/e2e/sim-cold-paste.spec.ts, a real browser; this pins
// the one thing an SSR render can prove -- the box's static shape, and that Run starts
// disabled with an empty box and no error shown for it (cold-paste.test.ts already covers
// `validateColdPaste` itself, including the decode and bad-paste cases this box renders).
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import { landingCopy } from '../../lib/sim/landing-copy';
import ColdPasteHero from './ColdPasteHero.svelte';

describe('ColdPasteHero', () => {
  it('renders the title, textarea and Run disabled while the box is empty, with no error shown', () => {
    const { body } = render(ColdPasteHero, { props: { busy: false, onrun: () => {} } });
    expect(body).toContain(landingCopy.pasteHeroTitle);
    expect(body).toContain('data-testid="sim-cold-paste-input"');
    const match = /<button[^>]*data-testid="sim-cold-paste-run"[^>]*>/.exec(body);
    if (match === null) throw new Error('run button not found');
    expect(match[0]).toContain('disabled=""');
    expect(body).not.toContain('data-testid="sim-cold-paste-error"');
  });

  it('says why Run is off while the box is empty, and carries the damage-specs line above the box', () => {
    const { body } = render(ColdPasteHero, { props: { busy: false, onrun: () => {} } });
    expect(body).toContain(landingCopy.pasteHeroRunReason);
    expect(body.indexOf(landingCopy.scopeCaveat)).toBeLessThan(
      body.indexOf('data-testid="sim-cold-paste-input"'),
    );
  });

  it('says the run is local and points the in-game hint at /setup', () => {
    const { body } = render(ColdPasteHero, { props: { busy: false, onrun: () => {} } });
    expect(body).toContain(landingCopy.pasteHeroRunsLocally);
    expect(body).toContain('href="/setup"');
    expect(body).toContain(landingCopy.pasteHeroHintLink);
  });

  it('puts the reason right after Run and the get-set-up link last', () => {
    const { body } = render(ColdPasteHero, { props: { busy: false, onrun: () => {} } });
    const run = body.indexOf('data-testid="sim-cold-paste-run"');
    const reason = body.indexOf(landingCopy.pasteHeroRunReason);
    expect(run).toBeLessThan(reason);
    expect(reason).toBeLessThan(body.indexOf(landingCopy.pasteHeroRunsLocally));
    expect(body.indexOf('data-testid="sim-cold-paste-setup"')).toBeGreaterThan(
      body.indexOf(landingCopy.pasteHeroRunsLocally),
    );
  });

  it('disables the textarea and shows the busy look while a load/run is in flight', () => {
    const { body } = render(ColdPasteHero, { props: { busy: true, onrun: () => {} } });
    const input = /<textarea[^>]*data-testid="sim-cold-paste-input"[^>]*>/.exec(body);
    if (input === null) throw new Error('textarea not found');
    expect(input[0]).toContain('disabled=""');
    const button = /<button[^>]*data-testid="sim-cold-paste-run"[^>]*>/.exec(body);
    if (button === null) throw new Error('run button not found');
    expect(button[0]).toContain('aria-busy="true"');
  });
});
