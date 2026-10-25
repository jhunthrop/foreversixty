import { describe, expect, it, vi } from 'vitest';
import { parseCharacterPath } from '../characters';
import { characterHistoryFixture } from '../../fixtures/character-history';
import { createHistoryController } from './history-controller.svelte';

const path = parseCharacterPath('/character/us/hardcore/elyra-duskvale');
if (path === null) throw new Error('fixture path');

const settle = (): Promise<void> => new Promise((resolve) => setTimeout(resolve, 0));

describe('createHistoryController', () => {
  it('starts loading and holds the history once it is read', async () => {
    const history = characterHistoryFixture();
    const controller = createHistoryController(path, () => Promise.resolve(history));
    expect(controller.status).toBe('loading');
    controller.load();
    await settle();
    expect(controller.status).toBe('ready');
    expect(controller.history).toBe(history);
  });

  it('reports an error and re-reads in place on the next load', async () => {
    const fetcher = vi
      .fn()
      .mockRejectedValueOnce(new Error('nope'))
      .mockResolvedValueOnce(characterHistoryFixture());
    const controller = createHistoryController(path, fetcher);
    controller.load();
    await settle();
    expect(controller.status).toBe('error');
    expect(controller.history).toBeNull();
    controller.load();
    expect(controller.status).toBe('loading');
    await settle();
    expect(controller.status).toBe('ready');
  });

  it('drops an answer that a newer load has superseded', async () => {
    let resolveFirst: (value: ReturnType<typeof characterHistoryFixture>) => void = () => undefined;
    const second = characterHistoryFixture({ as_of: '2027-01-01' });
    const fetcher = vi
      .fn()
      .mockReturnValueOnce(new Promise((resolve) => (resolveFirst = resolve)))
      .mockResolvedValueOnce(second);
    const controller = createHistoryController(path, fetcher);
    controller.load();
    controller.load();
    await settle();
    resolveFirst(characterHistoryFixture({ as_of: '2020-01-01' }));
    await settle();
    expect(controller.history).toBe(second);
  });
});
