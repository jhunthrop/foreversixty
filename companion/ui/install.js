// companion/ui/install.js
// The Mac's first-launch window (install.html). The Go side binds two
// functions on the window (internal/install/offer.go): fsInstallInfo says
// whether the system can open the app at login, and fsInstallMove installs
// it, starts the installed copy and closes this window.
'use strict';

(() => {
  const C = COPY.install;
  const $ = (id) => document.getElementById(id);
  const button = $('move');
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

  button.addEventListener('click', async () => {
    button.disabled = true;
    button.textContent = C.moving;
    error.hidden = true;
    try {
      const login = !$('move-login-row').hidden && $('move-login').checked;
      await window.fsInstallMove(login);
    } catch {
      error.textContent = C.failed;
      error.hidden = false;
      button.disabled = false;
      button.textContent = C.move;
    }
  });
})();
