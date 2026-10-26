// web/src/lib/guild/serial-save.ts
// One save at a time per sign-up row (raids spec 4): a press made while a save is in flight is
// never dropped. It marks the row as owing a save, and when the in-flight one settles a single
// follow-up sends whatever the row then shows (state, role and note merged), so the last answer
// the player gave is the one that lands.

export interface SerialSaveOptions<C, O> {
  /** The row's choice now, or `null` when there is nothing to send (an unsigned row's note). */
  read: () => C | null;
  send: (choice: C) => Promise<O>;
  /** Each answer; `more` is true when a follow-up is about to go, so the row keeps its edits. */
  settled: (outcome: O, choice: C, more: boolean) => void;
  onBusy?: (busy: boolean) => void;
}

export interface SerialSave {
  /** Saves now, or after the save in flight; resolves once the row owes nothing. */
  request: () => Promise<void>;
  busy: () => boolean;
}

export function createSerialSave<C, O>(options: SerialSaveOptions<C, O>): SerialSave {
  let inFlight: Promise<void> | null = null;
  let owed = false;

  async function drain(): Promise<void> {
    for (let choice = options.read(); choice !== null; choice = owed ? options.read() : null) {
      owed = false;
      const outcome = await options.send(choice);
      options.settled(outcome, choice, owed);
    }
  }

  function request(): Promise<void> {
    if (inFlight !== null) {
      owed = true;
      return inFlight;
    }
    options.onBusy?.(true);
    inFlight = drain().finally(() => {
      inFlight = null;
      owed = false;
      options.onBusy?.(false);
    });
    return inFlight;
  }

  return { request, busy: () => inFlight !== null };
}
