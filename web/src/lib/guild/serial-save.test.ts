// web/src/lib/guild/serial-save.test.ts
import { describe, expect, it } from 'vitest';
import { createSerialSave } from './serial-save';

interface Choice {
  state: string;
  note: string;
}

function deferred(): { promise: Promise<string>; resolve: (value: string) => void } {
  let resolve: (value: string) => void = () => {};
  const promise = new Promise<string>((done) => (resolve = done));
  return { promise, resolve };
}

describe('createSerialSave', () => {
  it('sends at once when idle and reports the outcome with nothing queued', async () => {
    const sent: Choice[] = [];
    const settled: [string, boolean][] = [];
    const save = createSerialSave<Choice, string>({
      read: () => ({ state: 'in', note: '' }),
      send: (choice) => {
        sent.push(choice);
        return Promise.resolve('ok');
      },
      settled: (outcome, _choice, more) => settled.push([outcome, more]),
    });
    await save.request();
    expect(sent).toEqual([{ state: 'in', note: '' }]);
    expect(settled).toEqual([['ok', false]]);
    expect(save.busy()).toBe(false);
  });

  it('never drops a press made while a save is in flight: one follow-up carries every change', async () => {
    let current: Choice = { state: 'in', note: 'late by 10' };
    const first = deferred();
    const sent: Choice[] = [];
    const more: boolean[] = [];
    const save = createSerialSave<Choice, string>({
      read: () => current,
      send: (choice) => {
        sent.push(choice);
        return sent.length === 1 ? first.promise : Promise.resolve('ok');
      },
      settled: (_outcome, _choice, queued) => more.push(queued),
    });
    const done = save.request();
    expect(save.busy()).toBe(true);
    current = { state: 'late', note: 'late by 10' };
    void save.request();
    current = { state: 'bench', note: 'late by 10' };
    void save.request();
    first.resolve('ok');
    await done;
    expect(sent).toEqual([
      { state: 'in', note: 'late by 10' },
      { state: 'bench', note: 'late by 10' },
    ]);
    expect(more).toEqual([true, false]);
    expect(save.busy()).toBe(false);
  });

  it('sends nothing when there is nothing to save', async () => {
    const sent: unknown[] = [];
    const save = createSerialSave<Choice, string>({
      read: () => null,
      send: (choice) => {
        sent.push(choice);
        return Promise.resolve('ok');
      },
      settled: () => {},
    });
    await save.request();
    expect(sent).toEqual([]);
  });
});
