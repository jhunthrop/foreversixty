// companion/ui/install.js
// The install windows (install.html?mode=...), bound by the Go side
// (internal/install):
//   move       the Mac's first launch: fsInstallInfo says whether the system can
//              open the app at login; fsInstallMove installs it, starts the
//              installed copy and closes this window.
//   setup      the Windows setup download; uninstall, Apps & Features' Uninstall.
//              fsSetupInfo says whether a copy is installed already and whether it
//              starts with Windows; fsSetupRun starts the steps, which report back
//              through fsProgress, then fsDone or fsFailed, and the Go side closes
//              the window.
// Closing the window instead quits; the next launch asks again.
'use strict';

(() => {
  const params = new URLSearchParams(location.search);
  const mode = ['setup', 'uninstall'].includes(params.get('mode')) ? params.get('mode') : 'move';
  const C = COPY.install[mode];
  const $ = (id) => document.getElementById(id);
  const title = $('move-title');
  const lede = $('move-lede');
  const button = $('move');
  const toggle = $('move-login');
  const toggleRow = $('move-login-row');
  const error = $('move-error');
  const progress = $('move-progress');
  // The words in force: an upgrade swaps the setup's title, button and busy line.
  let words = { title: C.title, lede: C.lede, action: C.action, busy: C.busy, done: C.done };

  const show = () => {
    title.textContent = words.title;
    lede.textContent = words.lede;
    button.textContent = words.action;
  };
  $('move-login-label').textContent = C.toggle;
  show();
  if (mode === 'uninstall') button.classList.replace('primary', 'danger');
  // The keyboard starts on the one button, except in the uninstall: Enter must not remove the
  // app, so there it starts on the toggle. After the first frame, when the web view takes focus.
  requestAnimationFrame(() => (mode === 'uninstall' ? toggle : button).focus());

  const busy = () => {
    error.hidden = true;
    toggle.disabled = true;
    title.textContent = words.busy;
    lede.hidden = true;
  };
  const showFailed = () => {
    show();
    lede.hidden = false;
    error.textContent = C.failed;
    error.hidden = false;
    progress.hidden = true;
    button.hidden = false;
    button.disabled = false;
    toggle.disabled = false;
  };

  if (mode === 'move') {
    toggleRow.hidden = true;
    window.fsInstallInfo?.().then((info) => {
      toggleRow.hidden = !(info && info.loginItems);
    });
    button.addEventListener('click', async () => {
      busy();
      button.disabled = true;
      button.textContent = words.busy;
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
    if (mode !== 'setup' || !info || !info.upgrade) return;
    words = { title: C.titleUpgrade, lede: C.ledeUpgrade, action: C.actionUpgrade, busy: C.busyUpgrade, done: C.doneUpgrade };
    toggle.checked = !!info.startup;
    show();
  });
  // A step reports as it starts: the bar shows the steps already done, and is full only at fsDone.
  window.fsProgress = (step, of, key) => {
    $('move-bar').style.width = `${Math.round(((step - 1) / of) * 100)}%`;
    $('move-step').textContent = COPY.install.steps[key] || words.busy;
  };
  window.fsFailed = showFailed;
  // Done: the bar full, the title says so, and the window closes a moment later.
  window.fsDone = () => {
    $('move-bar').style.width = '100%';
    $('move-step').textContent = '';
    title.textContent = words.done;
    if (mode !== 'uninstall') return;
    lede.textContent = C.lede;
    lede.hidden = false;
    toggleRow.hidden = true;
    progress.hidden = true;
  };
  button.addEventListener('click', () => {
    busy();
    button.hidden = true;
    progress.hidden = false;
    window.fsProgress(1, 1, '');
    window.fsSetupRun(toggle.checked);
  });
})();
