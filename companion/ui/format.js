// companion/ui/format.js
// The window's pure helpers: shortening a path, saying a time, naming a
// character. They take their inputs as arguments (the clock too) so a
// node test can run them: `node --test companion/ui/format.test.js`.
'use strict';

const FORMAT = (() => {
  const DAYS = ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat'];
  const MONTHS = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'];

  // midTruncate keeps the start and the end of a path, which are the
  // drive and the flavour folder, and drops the middle.
  function midTruncate(path, max) {
    if (path.length <= max) return path;
    const keep = max - 1;
    const head = Math.ceil(keep * 0.4);
    const tail = keep - head;
    return path.slice(0, head) + '…' + path.slice(path.length - tail);
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
  // or that it is not yet. Never a hyphenated name: the log's form is the
  // companion's to log, the player reads "Bow Jackzon".
  function creditText(credit, words) {
    const c = credit || {};
    if (c.state === 'credited') return words.credited(spaced(c.name));
    return words.notYet;
  }

  const spaced = (name) => String(name || '').replace(/-/g, ' ');

  // currentReport is the Current report row's value: the lead ("In progress",
  // or the zone once known), the credit words, and the muted help line that
  // only a report not credited yet carries.
  function currentReport(report, status) {
    const credited = !!report.credit && report.credit.state === 'credited';
    return {
      lead: report.zone || status.inProgress,
      separator: status.credit.separator,
      credit: creditText(report.credit, status.credit),
      credited,
      help: credited ? '' : status.credit.notYetHelp,
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

  return { midTruncate, when, lastPlayed, ago, creditText, currentReport, advancedView, checker, checkedLine, characterLabel, characterValue, automaticLabel, countdown };
})();

if (typeof module !== 'undefined') module.exports = FORMAT;
