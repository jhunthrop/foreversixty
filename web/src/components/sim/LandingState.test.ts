// web/src/components/sim/LandingState.test.ts
// Fix round 1, Important #1: `busyKey` is what gates every pick, including the background
// stored-pointer restore's own sentinel (character-bootstrap.ts's `RESTORE_BUSY_KEY`) --
// `svelte/server`'s render cannot click the row's own link (`follow`'s busyKey check is
// component wiring around a DOM MouseEvent, not pure logic with a seam of its own), but it
// can pin the static contract every busy state has to keep: the pick button itself is
// `disabled` whenever `busyKey` is anything but null, sentinel or a real character key alike.
import { render } from 'svelte/server';
import { describe, expect, it } from 'vitest';
import type { MeCharacter } from '../../lib/account/api';
import LandingState from './LandingState.svelte';

const CHARACTERS: MeCharacter[] = [
  { key: 'us/normal/simfury', name: 'Simfury', class: 'warrior', region: 'us', ruleset: 'normal' },
];

function pickButtonTag(body: string): string {
  const match = /<button[^>]*data-testid="sim-pick-us\/normal\/simfury"[^>]*>/.exec(body);
  if (match === null) throw new Error('pick button not found');
  return match[0];
}

describe('LandingState', () => {
  it('leaves the pick button enabled when nothing is busy', () => {
    const { body } = render(LandingState, {
      props: { characters: CHARACTERS, busyKey: null, onpick: () => {}, onother: () => {} },
    });
    // Svelte's own SSR renders a `false` boolean attribute as absent, not as `disabled="false"`
    // -- checked as the exact `disabled=""` attribute, not a bare substring match, since the
    // button's own class carries the unrelated `disabled:opacity-50` Tailwind variant.
    expect(pickButtonTag(body)).not.toContain('disabled=""');
  });

  it('disables the pick button while another row is busy', () => {
    const { body } = render(LandingState, {
      props: { characters: CHARACTERS, busyKey: 'us/normal/otherchar', onpick: () => {}, onother: () => {} },
    });
    expect(pickButtonTag(body)).toContain('disabled=""');
  });

  // The exact sentinel `restoreFromPointer` (SimView.svelte) holds `landingBusyKey` at
  // during a background restore -- disables the same way a real character key does.
  it('disables the pick button while the background restore sentinel is set', () => {
    const { body } = render(LandingState, {
      props: {
        characters: CHARACTERS,
        busyKey: 'restoring-current-character',
        onpick: () => {},
        onother: () => {},
      },
    });
    expect(pickButtonTag(body)).toContain('disabled=""');
  });
});
