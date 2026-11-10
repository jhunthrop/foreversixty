// companion/ui/install.js
// The install windows (install.html?mode=...), bound by the Go side
// (internal/install):
//   move       the Mac's first launch: fsInstallInfo says whether the system can
//              open the app at login; fsInstallMove installs it, starts the
//              installed copy and closes this window.
//   setup      the Windows setup download; uninstall, Apps & Features' Uninstall.
//              fsSetupInfo says whether a copy is installed already; fsSetupRun
//              starts the steps, which report back through fsProgress, then
//              fsDone (uninstall) or fsFailed, and the Go side closes the window.
// Closing the window instead quits; the next launch asks again.
'use strict';

(() => {
  const params = new URLSearchParams(location.search);
  const mode = ['setup', 'uninstall'].includes(params.get('mode')) ? params.get('mode') : 'move';
  const C = COPY.install[mode];
  const $ = (id) => document.getElementById(id);
  const button = $('move');
  const toggle = $('move-login');
  const toggleRow = $('move-login-row');
  const error = $('move-error');
  const progress = $('move-progress');

  $('move-title').textContent = C.title;
  $('move-lede').textContent = C.lede;
  $('move-login-label').textContent = C.toggle;
  button.textContent = C.action;
  if (mode === 'uninstall') button.classList.replace('primary', 'danger');

  const showFailed = () => {
    error.textContent = C.failed;
    error.hidden = false;
    progress.hidden = true;
    button.hidden = false;
    button.disabled = false;
    toggle.disabled = false;
    button.textContent = C.action;
  };

  if (mode === 'move') {
    toggleRow.hidden = true;
    window.fsInstallInfo?.().then((info) => {
      toggleRow.hidden = !(info && info.loginItems);
    });
    button.addEventListener('click', async () => {
      button.disabled = true;
      toggle.disabled = true;
      button.textContent = C.busy;
      error.hidden = true;
      try {
        await window.fsInstallMove(!toggleRow.hidden && toggle.checked);
      } catch {
        showFailed();
      }
    });
    return;
  }

  toggleRow.hidden = false;
  window.fsSetupInfo?.().then((info) => {
    if (mode === 'setup' && info && info.upgrade) $('move-lede').textContent = C.ledeUpgrade;
  });
  window.fsProgress = (step, of, key) => {
    $('move-bar').style.width = `${Math.round((step / of) * 100)}%`;
    $('move-step').textContent = COPY.install.steps[key] || C.busy;
  };
  window.fsFailed = showFailed;
  window.fsDone = () => {
    $('move-title').textContent = C.done;
    $('move-lede').hidden = true;
    toggleRow.hidden = true;
    progress.hidden = true;
  };
  button.addEventListener('click', () => {
    error.hidden = true;
    toggle.disabled = true;
    button.hidden = true;
    progress.hidden = false;
    window.fsProgress(0, 1, '');
    window.fsSetupRun(toggle.checked);
  });
})();
