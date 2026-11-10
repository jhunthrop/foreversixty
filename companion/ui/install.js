// companion/ui/install.js
// The Mac's first-launch window (install.html). The Go side binds two
// functions on the window (internal/install/offer.go): fsInstallInfo says
// whether the system can open the app at login, and fsInstallMove installs
// it, starts the installed copy and closes this window. Closing the window
// instead quits the companion; the next launch asks again.
'use strict';

(() => {
  const C = COPY.install;
  const $ = (id) => document.getElementById(id);
  const button = $('move');
  const login = $('move-login');
  const error = $('move-error');

  $('move-title').textContent = C.title;
  $('move-lede').textContent = C.lede;
  $('move-login-label').textContent = C.openAtLogin;
  button.textContent = C.move;

  if (typeof window.fsInstallInfo === 'function') {
    window.fsInstallInfo().then((info) => {
      $('move-login-row').hidden = !(info && info.loginItems);
    });
  }

  const busy = (on) => {
    button.disabled = on;
    login.disabled = on;
    button.textContent = on ? C.moving : C.move;
  };

  button.addEventListener('click', async () => {
    busy(true);
    error.hidden = true;
    try {
      await window.fsInstallMove(!$('move-login-row').hidden && login.checked);
    } catch {
      error.textContent = C.failed;
      error.hidden = false;
      busy(false);
    }
  });
})();
