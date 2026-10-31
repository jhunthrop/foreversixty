// web/src/lib/sim/follow-current-character.test.ts
// @vitest-environment jsdom
import { afterEach, describe, expect, it, vi, type Mock } from 'vitest';
import {
  CURRENT_CHARACTER_CHANGED,
  announcePageLoadedCharacter,
  announceSessionHeal,
  clearCurrent,
  writeCurrent,
  type CurrentCharacter,
} from '../current-character';
import type { CharacterLoaders } from './character-bootstrap';
import { followCurrentCharacter, simHoldsWork } from './follow-current-character';

function pointer(overrides: Partial<CurrentCharacter>): CurrentCharacter {
  return {
    source: 'armory',
    ref: 'us/normal/bow-jackzon',
    label: 'Bow Jackzon · Beast Mastery Hunter',
    classSlug: 'hunter',
    savedAt: '2026-10-09T00:00:00.000Z',
    ...overrides,
  };
}

type FakeLoaders = { [K in keyof CharacterLoaders]: Mock<CharacterLoaders[K]> };

function fakeLoaders(): FakeLoaders {
  return {
    loadCode: vi.fn<CharacterLoaders['loadCode']>().mockResolvedValue(undefined),
    loadAddon: vi.fn<CharacterLoaders['loadAddon']>().mockResolvedValue(undefined),
    loadBuild: vi.fn<CharacterLoaders['loadBuild']>().mockResolvedValue(undefined),
    loadFight: vi.fn<CharacterLoaders['loadFight']>().mockResolvedValue(undefined),
    loadStored: vi.fn<CharacterLoaders['loadStored']>().mockResolvedValue(undefined),
    setMessage: vi.fn<CharacterLoaders['setMessage']>(),
  };
}

function choose(value: CurrentCharacter): void {
  writeCurrent(value);
  window.dispatchEvent(new Event(CURRENT_CHARACTER_CHANGED));
}

const flush = async (): Promise<void> => {
  await new Promise((resolve) => setTimeout(resolve, 0));
};

describe('followCurrentCharacter', () => {
  let stop: () => void = () => undefined;
  afterEach(() => {
    stop();
    clearCurrent();
  });

  it('loads a newly chosen armory character in place, through the stored-character loader', async () => {
    const loaders = fakeLoaders();
    stop = followCurrentCharacter({ loaders, characterLoaded: () => true });
    choose(pointer({}));
    await flush();
    expect(loaders.loadStored).toHaveBeenCalledWith({ region: 'us', ruleset: 'normal', slug: 'bow-jackzon' });
  });

  it('loads a chosen pasted export through the addon loader', async () => {
    const loaders = fakeLoaders();
    stop = followCurrentCharacter({ loaders, characterLoaded: () => true });
    choose(pointer({ source: 'addon', ref: 'FS1:code', label: 'Simfury · Warrior', classSlug: 'warrior' }));
    await flush();
    expect(loaders.loadAddon).toHaveBeenCalledWith('FS1:code');
  });

  it('ignores a write the page made by loading that character itself', async () => {
    const loaders = fakeLoaders();
    stop = followCurrentCharacter({ loaders, characterLoaded: () => true });
    writeCurrent(pointer({}));
    announcePageLoadedCharacter();
    await flush();
    expect(loaders.loadStored).not.toHaveBeenCalled();
  });

  it('follows the session heal, except while the page holds work', async () => {
    const loaders = fakeLoaders();
    let holding = false;
    stop = followCurrentCharacter({ loaders, characterLoaded: () => true, holdsWork: () => holding });
    writeCurrent(pointer({}));
    announceSessionHeal();
    await flush();
    expect(loaders.loadStored).toHaveBeenCalledTimes(1);
    holding = true;
    announceSessionHeal();
    await flush();
    expect(loaders.loadStored).toHaveBeenCalledTimes(1);
    choose(pointer({}));
    await flush();
    expect(loaders.loadStored).toHaveBeenCalledTimes(2);
  });

  it('never follows on a pinned page', async () => {
    const loaders = fakeLoaders();
    stop = followCurrentCharacter({ loaders, characterLoaded: () => true, pinned: () => true });
    choose(pointer({}));
    await flush();
    expect(loaders.loadStored).not.toHaveBeenCalled();
  });

  it('runs back-to-back choices in order and reports busy until the last settles', async () => {
    const loaders = fakeLoaders();
    const order: string[] = [];
    loaders.loadStored.mockImplementation(async (path) => {
      await flush();
      order.push(path.slug);
    });
    const busy = vi.fn();
    stop = followCurrentCharacter({ loaders, characterLoaded: () => true, onBusy: busy });
    choose(pointer({ ref: 'us/normal/first' }));
    choose(pointer({ ref: 'us/normal/second' }));
    await flush();
    await flush();
    await flush();
    expect(order).toEqual(['first', 'second']);
    expect(busy.mock.calls.at(-1)).toEqual([false]);
  });
});

describe('simHoldsWork', () => {
  const idle = { phase: 'idle', serverRunning: false };

  it('holds while a run or its engine load is under way, or a server run is', () => {
    expect(simHoldsWork({ ...idle, phase: 'loading-engine' }, '')).toBe(true);
    expect(simHoldsWork({ ...idle, phase: 'running' }, '')).toBe(true);
    expect(simHoldsWork({ ...idle, phase: 'counting' }, '')).toBe(true);
    expect(simHoldsWork({ ...idle, serverRunning: true }, '')).toBe(true);
  });

  it('holds a ?code= build, and lets an idle or finished page move', () => {
    expect(simHoldsWork(idle, 'FS1:x')).toBe(true);
    expect(simHoldsWork(idle, '')).toBe(false);
    expect(simHoldsWork({ ...idle, phase: 'done' }, '')).toBe(false);
  });
});
