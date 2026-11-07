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
  it('links to /setup with the shared "get the addon" copy, at a 44px hit target on phone', () => {
    const { body } = render(SourceSwitcher, { props: requiredProps });
    const match = /<a[^>]*href="\/setup"[^>]*>/.exec(body);
    expect(match?.[0]).toContain('min-h-11');
    expect(body).toContain(currentCharacterCopy.getTheAddon);
  });

  it('puts the addon export on the one Plate and the other ways in as open rows', () => {
    const { body } = render(SourceSwitcher, { props: requiredProps });
    expect(body.match(/data-region="plate"/g)).toHaveLength(1);
    expect(body.indexOf('sim-addon-input')).toBeLessThan(body.indexOf('sim-build-input'));
    expect(body.indexOf('sim-fight-input')).toBeLessThan(body.indexOf('sim-signin'));
    expect(body).toContain(simCopy.sourceOtherWays);
  });

  it('signed out offers Battle.net; signed in, the way back to the characters', () => {
    expect(render(SourceSwitcher, { props: requiredProps }).body).toContain(landingCopy.signInWithBattlenet);
    const signedIn = render(SourceSwitcher, { props: { ...requiredProps, signedIn: true } }).body;
    expect(signedIn).toContain('sim-back-to-characters');
    expect(signedIn).not.toContain('sim-signin');
    const none = render(SourceSwitcher, {
      props: { ...requiredProps, signedIn: true, hasCharacters: false },
    }).body;
    expect(none).toContain(simCopy.charactersOnTheirWay);
  });

  it('says why a source did not load', () => {
    const { body } = render(SourceSwitcher, { props: { ...requiredProps, message: 'No such build.' } });
    expect(body).toContain('sim-source-message');
    expect(body).toContain('No such build.');
  });
});
