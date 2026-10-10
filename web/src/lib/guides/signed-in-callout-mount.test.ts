// web/src/lib/guides/signed-in-callout-mount.test.ts
// @vitest-environment jsdom
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';
import type { Me, MeCharacter } from '../account/api';
import { CURRENT_CHARACTER_CHANGED, writeCurrent } from '../current-character';
import { pointerForCharacter } from '../account/main-character';
import { mountSignedInCallout } from './signed-in-callout-mount';

let session: Me | null = null;
vi.mock('../account/api', () => ({ fetchMeOnce: () => Promise.resolve(session) }));

function character(name: string, className: string, spec: string): MeCharacter {
  return {
    key: `us/normal/${name.toLowerCase()}`,
    region: 'us',
    ruleset: 'normal',
    name,
    class: className,
    spec,
  };
}

const arms = character('Armsy', 'Warrior', 'Arms');
const bow = character('Bow', 'Hunter', 'Beast Mastery');

function choose(target: MeCharacter): void {
  writeCurrent(pointerForCharacter(target));
  window.dispatchEvent(new Event(CURRENT_CHARACTER_CHANGED));
}

async function settle(): Promise<void> {
  await vi.advanceTimersByTimeAsync(5);
}

describe('mountSignedInCallout', () => {
  let root: HTMLElement;
  let stop: () => void = () => undefined;
  beforeEach(() => {
    vi.useFakeTimers();
    localStorage.clear();
    root = document.createElement('div');
    document.body.append(root);
    session = {
      user: { id: 1, battletag: null, email: null, role: 'user', anonymize: false },
      characters: [arms, bow],
      guilds: [],
    };
  });
  afterEach(() => {
    stop();
    root.remove();
    vi.useRealTimers();
  });

  it('draws the selected character and redraws for the next one', async () => {
    choose(arms);
    stop = mountSignedInCallout(root);
    await settle();
    expect(root.querySelector('a')?.getAttribute('href')).toBe('/guides/warrior/arms');

    choose(bow);
    await settle();
    expect(root.querySelectorAll('a')).toHaveLength(1);
    expect(root.querySelector('a')?.getAttribute('href')).toBe('/guides/hunter/beast-mastery');
  });

  it('hides the slot when signed out and shows it again once a character is selected', async () => {
    session = null;
    stop = mountSignedInCallout(root);
    await settle();
    expect(root.hidden).toBe(true);
    expect(root.childElementCount).toBe(0);

    session = {
      user: { id: 1, battletag: null, email: null, role: 'user', anonymize: false },
      characters: [bow],
      guilds: [],
    };
    choose(bow);
    await settle();
    expect(root.hidden).toBe(false);
    expect(root.querySelector('a')?.getAttribute('href')).toBe('/guides/hunter/beast-mastery');
  });
});
