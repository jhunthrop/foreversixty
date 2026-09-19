// web/src/lib/sim/notify.ts
// Design 5.4: "browser notification when a server run completes".
//
// The Notification API behind a two-method interface, so the decision logic is testable
// without a browser and so the store never touches `window` -- which is the rule
// store.svelte.ts's own header sets and which keeps it unit-testable.
//
// Permission is only ever asked for from the player's own click on the control. A page
// that asks on load is a page people permanently deny.

export interface Notifier {
  /** "default" | "granted" | "denied". */
  permission: string;
  request(): Promise<string>;
  show(title: string, body: string): void;
}

/** The real one, or null where the browser has no Notification API (Safari in a frame, and every SSR pass). */
export function browserNotifier(): Notifier | null {
  const ctor = (globalThis as { Notification?: typeof Notification }).Notification;
  if (ctor === undefined) return null;
  return {
    get permission() {
      return ctor.permission;
    },
    request: () => ctor.requestPermission(),
    show: (title, body) => {
      // A notification that throws must never take the results page down with it: a
      // browser can refuse to construct one even with permission (a private window, a
      // page that has lost its user gesture).
      try {
        new ctor(title, { body });
      } catch {
        /* no notification; the result is on screen regardless */
      }
    },
  };
}

/** True once notifications are usable. Asks at most once, and never a browser that said no. */
export async function enableNotifications(notifier: Notifier | null): Promise<boolean> {
  if (notifier === null) return false;
  if (notifier.permission === 'granted') return true;
  if (notifier.permission === 'denied') return false;
  return (await notifier.request()) === 'granted';
}

/** Shows one, if the player asked for them and the browser allows them. Returns whether it did. */
export function notifyFinished(
  notifier: Notifier | null,
  enabled: boolean,
  title: string,
  body: string,
): boolean {
  if (notifier === null || !enabled || notifier.permission !== 'granted') return false;
  notifier.show(title, body);
  return true;
}
