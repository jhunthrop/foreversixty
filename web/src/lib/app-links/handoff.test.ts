// /app/auth's handoff: the code and state leave the address bar at once and live only in
// the script's memory; the button carries them back to the app.
import { JSDOM } from 'jsdom';
import { describe, expect, it, vi } from 'vitest';
import { appHandoff } from './handoff';

const PAGE = `<!doctype html><html><body>
  <button data-app-handoff hidden>Open Forever Sixty</button>
  <p data-app-handoff-missing hidden>Start again</p>
</body></html>`;

function open(url: string, userAgent = 'Mozilla/5.0 (iPhone; CPU iPhone OS 26_0 like Mac OS X)') {
  const dom = new JSDOM(PAGE, { url });
  const win = dom.window as unknown as Window;
  Object.defineProperty(win.navigator, 'userAgent', { value: userAgent });
  const navigate = vi.fn<(href: string) => void>();
  appHandoff(win, 'gg.foreversixty.app', navigate);
  win.document.dispatchEvent(new (win as Window & typeof globalThis).Event('DOMContentLoaded'));
  const doc = win.document;
  return {
    win,
    navigate,
    button: doc.querySelector<HTMLButtonElement>('[data-app-handoff]')!,
    missing: doc.querySelector<HTMLElement>('[data-app-handoff-missing]')!,
  };
}

describe('appHandoff', () => {
  it('strips the code and state from the address bar before anything else runs', () => {
    const { win } = open('https://foreversixty.gg/app/auth?code=c0de&state=st4te');
    expect(win.location.href).toBe('https://foreversixty.gg/app/auth');
    expect(win.document.documentElement.outerHTML).not.toContain('c0de');
  });

  it('re-opens the app link with the code from memory on a tap', () => {
    const { button, navigate } = open('https://foreversixty.gg/app/auth?code=c0de&state=st4te');
    expect(button.hidden).toBe(false);
    button.click();
    expect(navigate).toHaveBeenCalledWith('https://foreversixty.gg/app/auth?code=c0de&state=st4te');
  });

  it('opens the Android app by package from Chrome, which needs the explicit intent', () => {
    const { button, navigate } = open(
      'https://foreversixty.gg/app/auth?code=c0de&state=st4te',
      'Mozilla/5.0 (Linux; Android 15; Pixel 9) Chrome/130.0',
    );
    button.click();
    expect(navigate).toHaveBeenCalledWith(
      'intent://foreversixty.gg/app/auth?code=c0de&state=st4te#Intent;scheme=https;package=gg.foreversixty.app;end',
    );
  });

  it('offers no button when there is nothing to hand over', () => {
    const { button, missing, navigate } = open('https://foreversixty.gg/app/auth');
    expect(button.hidden).toBe(true);
    expect(missing.hidden).toBe(false);
    button.click();
    expect(navigate).not.toHaveBeenCalled();
  });
});
