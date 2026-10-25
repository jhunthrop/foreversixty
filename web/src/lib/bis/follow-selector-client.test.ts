// web/src/lib/bis/follow-selector-client.test.ts
// @vitest-environment jsdom
import { afterEach, describe, expect, it, vi } from 'vitest';
import {
  CURRENT_CHARACTER_CHANGED,
  announcePageLoadedCharacter,
  announceSessionHeal,
} from '../current-character-events';
import { mountBisFollow } from './follow-selector-client';

const follow = vi.fn<(root: HTMLElement) => Promise<void>>(() => Promise.resolve());
// The run module is the heavy half; a mock factory is only evaluated when something imports it,
// so `loaded` records whether the page's own listener pulled it in.
let loaded = false;
vi.mock('./follow-selector-run', () => {
  loaded = true;
  return { follow: (root: HTMLElement) => follow(root) };
});

describe('mountBisFollow', () => {
  let stop: (() => void) | undefined;
  afterEach(() => {
    stop?.();
    follow.mockClear();
  });

  it('loads nothing until the pointer changes, so the page pays no fetch for it at load', () => {
    stop = mountBisFollow(document.createElement('div'));
    expect(loaded).toBe(false);
  });

  it('ignores the change the page itself announced after loading a character', async () => {
    stop = mountBisFollow(document.createElement('div'));
    announcePageLoadedCharacter();
    await Promise.resolve();
    expect(follow).not.toHaveBeenCalled();
  });

  it('follows a pointer change by loading the run module and calling it with the page root', async () => {
    const root = document.createElement('div');
    stop = mountBisFollow(root);
    window.dispatchEvent(new Event(CURRENT_CHARACTER_CHANGED));
    await vi.waitFor(() => expect(follow).toHaveBeenCalledWith(root));
    expect(loaded).toBe(true);
  });

  it("follows the session's pointer heal: a fresh session lands on the main", async () => {
    const root = document.createElement('div');
    stop = mountBisFollow(root);
    announceSessionHeal();
    await vi.waitFor(() => expect(follow).toHaveBeenCalledWith(root));
  });
});
