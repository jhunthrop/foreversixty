// web/src/components/CopyForGame.test.ts
// Task L8.1: one "Copy for the game" button on every surface that holds a build or a sim
// result. This project's vitest config hands components Svelte's server entry (see
// sim/HelpNote.test.ts), so there is no client-side click to simulate: the static render
// proves the button's label, test id and disabled state, and the click's whole behaviour --
// clipboard, the two-second "Copied" and the failure fallback -- lives in `copy-for-game.ts`'s
// controller, which the component calls and which runs here against a stubbed clipboard
// and fake timers.
import { render } from 'svelte/server';
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';
import { addonCopy } from '../lib/addon/copy';
import { COPIED_VISIBLE_MS as COPIED_FOR_GAME_MS } from '../lib/ui/copy-feedback';
import { copyForGameLabel, createCopyForGame, type CopyForGameState } from '../lib/addon/copy-for-game';
import CopyForGame from './CopyForGame.svelte';

const CODE = 'FSB1:1.60.1.69893:paladin:111:head=12640:stamina=17';

function buttonOf(body: string): string {
  const match = /<button[^>]*>[\s\S]*?<\/button>/.exec(body);
  if (match === null) throw new Error('button not found');
  return match[0];
}

describe('CopyForGame (static render)', () => {
  it('reads "Copy for the game" and carries the caller’s test id', () => {
    const { body } = render(CopyForGame, { props: { code: CODE, testid: 'copy-addon-code' } });
    const button = buttonOf(body);
    expect(button).toContain(addonCopy.copyForGame);
    expect(addonCopy.copyForGame).toBe('Copy for the game');
    expect(button).toContain('data-testid="copy-addon-code"');
    expect(button).not.toMatch(/\sdisabled(?=[\s>=])/);
  });

  it('disables the button when there is no code to copy', () => {
    const { body } = render(CopyForGame, { props: { code: '' } });
    expect(buttonOf(body)).toMatch(/\sdisabled(?=[\s>=])/);
  });

  it('shows no fallback field until a copy has failed', () => {
    const { body } = render(CopyForGame, { props: { code: CODE } });
    expect(body).not.toContain(addonCopy.copyFailed);
    expect(body).not.toContain('<input');
  });
});

describe('CopyForGame (the click)', () => {
  let states: CopyForGameState[];
  let writeText: ReturnType<typeof vi.fn>;

  beforeEach(() => {
    vi.useFakeTimers();
    states = [];
    writeText = vi.fn().mockResolvedValue(undefined);
    vi.stubGlobal('navigator', { clipboard: { writeText } });
  });

  afterEach(() => {
    vi.useRealTimers();
    vi.unstubAllGlobals();
  });

  it('copies the code and reads "Copied, paste it in game with /fs" for 2 s', async () => {
    const copier = createCopyForGame((state) => states.push(state));
    await copier.copy(CODE);
    expect(writeText).toHaveBeenCalledWith(CODE);
    expect(states).toEqual(['copied']);
    expect(copyForGameLabel('copied')).toBe(addonCopy.copiedForGame);
    expect(addonCopy.copiedForGame).toBe('Copied, paste it in game with /fs');

    vi.advanceTimersByTime(COPIED_FOR_GAME_MS - 1);
    expect(states).toEqual(['copied']);
    vi.advanceTimersByTime(1);
    expect(states).toEqual(['copied', 'idle']);
    expect(copyForGameLabel('idle')).toBe(addonCopy.copyForGame);
    expect(COPIED_FOR_GAME_MS).toBe(2000);
  });

  it('a second copy inside the 2 s restarts the window rather than reverting early', async () => {
    const copier = createCopyForGame((state) => states.push(state));
    await copier.copy(CODE);
    vi.advanceTimersByTime(1500);
    await copier.copy(CODE);
    vi.advanceTimersByTime(1500);
    expect(states).toEqual(['copied', 'copied']);
    vi.advanceTimersByTime(500);
    expect(states).toEqual(['copied', 'copied', 'idle']);
  });

  it('says the copy failed and offers the code to select when the clipboard refuses', async () => {
    writeText.mockRejectedValue(new Error('denied'));
    const copier = createCopyForGame((state) => states.push(state));
    await copier.copy(CODE);
    expect(states).toEqual(['failed']);
    expect(addonCopy.copyFailed).toBe('Could not copy; select the code below instead');
    // The button itself goes back to its own label; the message and the field say the rest.
    expect(copyForGameLabel('failed')).toBe(addonCopy.copyForGame);
    vi.advanceTimersByTime(COPIED_FOR_GAME_MS * 2);
    expect(states).toEqual(['failed']);
  });

  it('fails the same way when the browser has no clipboard at all', async () => {
    vi.stubGlobal('navigator', {});
    const copier = createCopyForGame((state) => states.push(state));
    await copier.copy(CODE);
    expect(states).toEqual(['failed']);
  });

  it('copies nothing for an empty code', async () => {
    const copier = createCopyForGame((state) => states.push(state));
    await copier.copy('');
    expect(writeText).not.toHaveBeenCalled();
    expect(states).toEqual([]);
  });

  it('stops its timer on dispose, so an unmounted button is never written to', async () => {
    const copier = createCopyForGame((state) => states.push(state));
    await copier.copy(CODE);
    copier.dispose();
    vi.advanceTimersByTime(COPIED_FOR_GAME_MS);
    expect(states).toEqual(['copied']);
  });
});

describe('the addon hint', () => {
  it('names /fs, not the retired /fs follow', () => {
    expect(addonCopy.addonCodeHint).toBe(
      'Paste this in game with /fs; the Forever Sixty line and the talent window follow it from then on.',
    );
    expect(addonCopy.addonCodeHint).not.toContain('/fs follow');
  });
});
