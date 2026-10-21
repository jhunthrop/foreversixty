// companion/ui/format.js
// The window's pure helpers: shortening a path, saying a time, naming a
// character. They take their inputs as arguments (the clock too) so a
// node test can run them: `node --test companion/ui/format.test.js`.
'use strict';

const FORMAT = (() => {
  const DAYS = ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat'];
  const BS = String.fromCharCode(92); // the Windows path separator
  const MONTHS = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'];

  // pathTruncate shortens a folder path at its separators only, never
  // inside a name. Past the budget it keeps the first segment (the drive)
  // and the last two: C:\…\World of Warcraft\_classic_era_; then the drive
  // and the last one: C:\…\_classic_era_. Only when that one name alone is
  // still too long is it cut, with the ellipsis at the end of the name.
  function pathTruncate(path, max) {
    if (path.length <= max) return path;
    const sep = path.includes(BS) ? BS : '/';
    const segs = path.split(sep);
    if (segs.length < 3) return cutName(path, max);
    const head = segs[0] + sep + '…' + sep;
    const two = head + segs.slice(-2).join(sep);
    if (segs.length > 3 && two.length <= max) return two;
    const one = head + segs[segs.length - 1];
    if (one.length <= max) return one;
    const room = max - head.length;
    return head + cutName(segs[segs.length - 1], Math.max(room, 2));
  }

  // cutName is a single name cut to max characters, ellipsis last.
  function cutName(name, max) {
    return name.length <= max ? name : name.slice(0, max - 1) + '…';
  }

  // pathBudget is how many characters of a folder path a row has room for:
  // 56 in the 800 px window, 38 at the 560 px minimum. The Status page's
  // rows have no control beside the path and take 6 more.
  function pathBudget(width, wide) {
    return (width >= 700 ? 56 : 38) + (wide ? 6 : 0);
  }

  const pad = (n) => String(n).padStart(2, '0');
  const sameDay = (a, b) => a.getFullYear() === b.getFullYear() && a.getMonth() === b.getMonth() && a.getDate() === b.getDate();

  // when is "Today 20:14", "Yesterday 20:14" or "Tue 14 Oct, 20:14".
  function when(iso, now, words) {
    const d = new Date(iso);
    if (Number.isNaN(d.getTime())) return '';
    const clock = pad(d.getHours()) + ':' + pad(d.getMinutes());
    if (sameDay(d, now)) return words.today + ' ' + clock;
    const yesterday = new Date(now.getFullYear(), now.getMonth(), now.getDate() - 1);
    if (sameDay(d, yesterday)) return words.yesterday + ' ' + clock;
    return DAYS[d.getDay()] + ' ' + d.getDate() + ' ' + MONTHS[d.getMonth()] + ', ' + clock;
  }

  // lastPlayed is "Last played 19:42" today, "Last played 15 Oct 19:42"
  // another day, and "Not played yet" when the time is not known.
  function lastPlayed(iso, now, words) {
    const d = iso ? new Date(iso) : null;
    if (!d || Number.isNaN(d.getTime())) return words.notPlayed;
    const clock = pad(d.getHours()) + ':' + pad(d.getMinutes());
    if (sameDay(d, now)) return words.lastPlayed + ' ' + clock;
    return words.lastPlayed + ' ' + d.getDate() + ' ' + MONTHS[d.getMonth()] + ' ' + clock;
  }

  // ago is how long since iso, in the Addon row's words: minutes under
  // an hour, the clock later the same day, then yesterday and days.
  // Unreadable time is an empty string.
  function ago(iso, now, words) {
    const d = new Date(iso);
    if (Number.isNaN(d.getTime())) return '';
    const mins = Math.floor((now.getTime() - d.getTime()) / 60000);
    if (mins < 1) return words.lessThanMinute;
    if (mins < 2) return words.oneMinute;
    if (mins < 60) return words.minutes(mins);
    if (sameDay(d, now)) return words.today(pad(d.getHours()) + ':' + pad(d.getMinutes()));
    const days = Math.round((new Date(now.getFullYear(), now.getMonth(), now.getDate()) -
      new Date(d.getFullYear(), d.getMonth(), d.getDate())) / 86400000);
    return days <= 1 ? words.yesterday : words.days(days);
  }

  // creditText is the suffix after a report's zone: who it is credited to,
  // that the log has named nobody yet (state none, a report still open: not
  // a fault), or that it is not credited yet. Never a hyphenated name: the
  // log's form is the companion's to log, the player reads "Bow Jackzon".
  // A report that is done and still named nobody is not credited, whatever
  // its state says: pass done=true for it.
  function creditText(credit, words, done) {
    const c = credit || {};
    if (c.state === 'credited') return words.credited(spaced(c.name));
    if (c.state === 'none' && !done) return words.noName;
    return words.notYet;
  }

  // reportCredit is how a Reports row ends: kind credited, notyet (the one
  // orange phrase, and only when the report is on the site: it is credited
  // by creating it with a character or by a patch), noname (muted), or
  // absent (the site has no such report yet, so there is nothing to
  // credit and no URL).
  function reportCredit(r, words) {
    if (!r.url) return { kind: 'absent', text: '' };
    const c = r.credit || {};
    const text = creditText(c, words, r.done);
    return { kind: c.state === 'credited' ? 'credited' : (text === words.noName ? 'noname' : 'notyet'), text };
  }

  // anyNotCredited is true when a row in the list reads "not credited yet".
  function anyNotCredited(reports, words) {
    return reports.some((r) => reportCredit(r, words).kind === 'notyet');
  }

  // saveStrip is the Settings strip's words and whether Save can be pressed.
  function saveStrip(dirty, saving, words) {
    return { text: dirty ? words.unsaved : words.noChanges, enabled: dirty && !saving };
  }

  const spaced = (name) => String(name || '').replace(/-/g, ' ');

  // currentReport is the Current report row's value: the lead ("In progress",
  // or the zone once known), the credit words, and the muted help line that
  // only a report not credited yet carries.
  function currentReport(report, status) {
    const state = (report.credit || {}).state;
    const credited = state === 'credited';
    const noName = state === 'none';
    return {
      lead: report.zone || status.inProgress,
      separator: status.credit.separator,
      credit: creditText(report.credit, status.credit),
      credited,
      noName,
      help: credited || noName ? '' : status.credit.notYetHelp,
    };
  }

  // advancedView is how a game folder's advanced-logging switch reads. Only
  // the watched install, whose raids are going to be uploaded, shouts when it
  // is off (a warning glyph, and the instruction unless the headline already
  // gave it); any other install is "off" in the muted token and nothing more.
  function advancedView(install, headlineKey, words) {
    if (!install.advanced_logging_known) return { kind: 'unknown', text: words.advancedUnknown, help: '' };
    if (install.advanced_logging) return { kind: 'on', text: words.advancedOn, help: '' };
    if (!install.watching) return { kind: 'muted', text: words.advancedOff, help: '' };
    return {
      kind: 'warn',
      text: words.advancedOffWatched,
      help: headlineKey === 'advanced_off' ? '' : (install.advanced_logging_help || ''),
    };
  }

  // The button asks for at most one check a half minute.
  const CHECK_GAP_MS = 30000;

  // checker is the Check for builds now button's brain: one press at a
  // time, one every half minute. post is the request; clock is Date.now.
  // press() answers 'ignored' when it did not ask the companion anything.
  function checker(post, clock) {
    let last = null;
    let running = false;
    const canPress = () => !running && (last === null || clock() - last >= CHECK_GAP_MS);
    return {
      get running() { return running; },
      canPress,
      async press() {
        if (!canPress()) return 'ignored';
        running = true;
        last = clock();
        try {
          return await post();
        } finally {
          running = false;
        }
      },
    };
  }

  // checkedLine is "Checked just now · 1 build for Bow Jackzon", or
  // "... · nothing new"; the time is the Addon row's words.
  function checkedLine(checked, now, words, agoWords) {
    const mins = Math.floor((now.getTime() - new Date(checked.at).getTime()) / 60000);
    const when = mins < 1 ? words.checkedJustNow : ago(checked.at, now, agoWords);
    const found = checked.new_builds > 0
      ? words.checkedBuilds(checked.new_builds, (checked.for || []).map(spaced))
      : words.checkedNothing;
    return words.checked(when, found);
  }

  // characterLabel is "Thrall · Hardcore · US": the ruleset in words.
  function characterLabel(c, rulesets) {
    const ruleset = rulesets[String(c.ruleset).toLowerCase()] || c.ruleset;
    return c.name + ' · ' + ruleset + ' · ' + String(c.region).toUpperCase();
  }

  const characterValue = (c) => [c.region, c.ruleset, c.name].join('|');

  // automaticLabel is the select's first option: the log decides, and
  // once it has named a character the option says who.
  function automaticLabel(detected, settings) {
    return detected ? settings.characterAutoNow(detected) : settings.characterAuto;
  }

  // countdown is the whole minutes and seconds left, never negative.
  function countdown(expiresISO, now) {
    const left = Math.max(0, Math.floor((new Date(expiresISO).getTime() - now.getTime()) / 1000));
    return { m: Math.floor(left / 60), s: left % 60 };
  }

  // The sign-in link works ten minutes; the expiry bar is the share left.
  const LINK_SECONDS = 600;
  const LATE_SECONDS = 60;

  // secondsLeft is the whole seconds until expiresISO, never negative.
  function secondsLeft(expiresISO, now) {
    const t = new Date(expiresISO).getTime();
    if (Number.isNaN(t)) return 0;
    return Math.max(0, Math.floor((t - now.getTime()) / 1000));
  }

  // expiryShare is how much of the link's life is left, from 1 to 0, for
  // the bar under the countdown.
  function expiryShare(expiresISO, now) {
    return Math.min(1, secondsLeft(expiresISO, now) / LINK_SECONDS);
  }

  // lastMinute is true once the countdown should turn amber.
  function lastMinute(expiresISO, now) {
    return secondsLeft(expiresISO, now) <= LATE_SECONDS;
  }

  return { pathTruncate, pathBudget, reportCredit, anyNotCredited, saveStrip, when, lastPlayed, ago, creditText, currentReport, advancedView, checker, checkedLine, characterLabel, characterValue, automaticLabel, countdown, secondsLeft, expiryShare, lastMinute };
})();

if (typeof module !== 'undefined') module.exports = FORMAT;
