// web/src/lib/app-links/handoff.ts
// The script /app/auth runs from its <head>, before first paint. The API's app sign-in
// (api/internal/auth/app_auth.go) sends the browser here with a one-time code; normally the
// app catches the link first, and this page is what shows when it did not (the app is not
// installed, the link is not verified yet, or Chrome needs a tap to leave a Custom Tab).
//
// The code and state are taken out of the address bar at once with history.replaceState, so
// they are never in history, a screenshot of the URL, or a link someone is asked to paste.
// From then on they live only in this closure: the button is a <button>, not a link, and
// carries nothing in the DOM.
//
// The page inlines this function's source (`(${appHandoff})(window, package)`), so it must
// stay self-contained: no imports and nothing from outside its own body.
export function appHandoff(
  win: Window,
  androidPackage: string,
  navigate: (href: string) => void = (href) => win.location.assign(href),
): void {
  const params = new URLSearchParams(win.location.search);
  const code = params.get('code');
  const state = params.get('state');
  if (win.location.search !== '') win.history.replaceState(null, '', win.location.pathname);

  win.document.addEventListener('DOMContentLoaded', () => {
    const button = win.document.querySelector<HTMLButtonElement>('[data-app-handoff]');
    const missing = win.document.querySelector<HTMLElement>('[data-app-handoff-missing]');
    if (button === null || missing === null) return;
    if (code === null || code === '' || state === null || state === '') {
      button.hidden = true;
      missing.hidden = false;
      return;
    }
    button.hidden = false;
    const query = `?${new URLSearchParams({ code, state }).toString()}`;
    button.addEventListener('click', () => {
      const { host, origin, pathname } = win.location;
      // Chrome on Android opens a verified App Link from the same site only through an
      // explicit intent naming the package; everywhere else the https link is the app link.
      const href = /Android/i.test(win.navigator.userAgent)
        ? `intent://${host}${pathname}${query}#Intent;scheme=https;package=${androidPackage};end`
        : `${origin}${pathname}${query}`;
      navigate(href);
    });
  });
}
