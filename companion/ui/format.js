// companion/ui/format.js
// The window's pure helpers: shortening a path, saying a time, naming a
// character. They take their inputs as arguments (the clock, the zone and the
// locale too) so a node test can run them: `node --test companion/ui/format.test.js`.
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

  // The house dates and clock, the site's own (web/src/lib/dates.ts), held to the shared examples
  // in apps/shared/formats.json by format.test.js (owner rulings 2026-11-10: "Sep", the clock in
  // the viewer's locale). A zone of undefined is the window's own; a locale of undefined is the
  // window's language.
  const SITE_LOCALE = 'en-US';
  const WEEKDAYS = ['Sunday', 'Monday', 'Tuesday', 'Wednesday', 'Thursday', 'Friday', 'Saturday'];
  const FULL_MONTHS = ['January', 'February', 'March', 'April', 'May', 'June', 'July', 'August', 'September', 'October', 'November', 'December'];
  const MINUTE_MS = 60000;
  const HOUR_MS = 60 * MINUTE_MS;
  const DAY_MS = 24 * HOUR_MS;
  const AGO_CEILING_DAYS = 30;
  const two = (n) => String(n).padStart(2, '0');

  const viewerLocale = () => (typeof navigator !== 'undefined' && navigator.language) || SITE_LOCALE;

  const twelveHour = new Map();
  // usesTwelveHourClock: en-US reads 8 pm, en-GB and de-DE read 20:00.
  function usesTwelveHourClock(locale) {
    if (!twelveHour.has(locale)) {
      twelveHour.set(locale, new Intl.DateTimeFormat(locale, { hour: 'numeric' }).resolvedOptions().hour12 === true);
    }
    return twelveHour.get(locale);
  }

  // zoned is the instant on the zone's calendar and 24-hour clock.
  function zoned(d, zone) {
    const parts = {};
    const f = new Intl.DateTimeFormat(SITE_LOCALE, {
      timeZone: zone, hourCycle: 'h23', year: 'numeric', month: 'numeric', day: 'numeric',
      weekday: 'short', hour: 'numeric', minute: '2-digit',
    });
    for (const p of f.formatToParts(d)) parts[p.type] = p.value;
    return {
      year: Number(parts.year), month: Number(parts.month) - 1, day: Number(parts.day),
      weekday: DAYS.indexOf(parts.weekday), hour: Number(parts.hour) % 24, minute: Number(parts.minute),
    };
  }

  const valid = (d) => d instanceof Date && !Number.isNaN(d.getTime());

  // clock is the time of day in the viewer's hour cycle: "8 pm", "8:30 pm" (the minutes on the
  // hour too with alwaysMinutes), or "20:00".
  function clock(d, zone, locale, alwaysMinutes) {
    if (!valid(d)) return '';
    const { hour, minute } = zoned(d, zone);
    if (!usesTwelveHourClock(locale || viewerLocale())) return two(hour) + ':' + two(minute);
    const twelve = hour % 12 === 0 ? 12 : hour % 12;
    return twelve + (alwaysMinutes || minute !== 0 ? ':' + two(minute) : '') + (hour < 12 ? ' am' : ' pm');
  }

  // dayMonth is "9 Dec"; weekdayDayMonth "Wed 9 Dec", "Wed 6 Jan 2027" outside now's year;
  // longDay "Wednesday 9 December".
  function dayMonth(d, zone) {
    if (!valid(d)) return '';
    const at = zoned(d, zone);
    return at.day + ' ' + MONTHS[at.month];
  }
  function weekdayDayMonth(d, zone, now) {
    if (!valid(d)) return '';
    const at = zoned(d, zone);
    const base = DAYS[at.weekday] + ' ' + at.day + ' ' + MONTHS[at.month];
    return at.year === zoned(now, zone).year ? base : base + ' ' + at.year;
  }
  function longDay(d, zone) {
    if (!valid(d)) return '';
    const at = zoned(d, zone);
    return WEEKDAYS[at.weekday] + ' ' + at.day + ' ' + FULL_MONTHS[at.month];
  }

  // stamp is "19 Oct, 9:40 pm": a day and its clock.
  function stamp(d, zone, locale) {
    if (!valid(d)) return '';
    return dayMonth(d, zone) + ', ' + clock(d, zone, locale);
  }

  // compactAgo is the site's freshness age: "just now", "5 min ago", "2 h ago", "3 d ago", then
  // the date past 30 days ("1 Oct", UTC as the site writes it). A time ahead of the clock is
  // "just now".
  function compactAgo(d, now) {
    if (!valid(d)) return '';
    const elapsed = Math.max(0, now.getTime() - d.getTime());
    if (elapsed < MINUTE_MS) return 'just now';
    if (elapsed < HOUR_MS) return Math.floor(elapsed / MINUTE_MS) + ' min ago';
    if (elapsed < DAY_MS) return Math.floor(elapsed / HOUR_MS) + ' h ago';
    const days = Math.floor(elapsed / DAY_MS);
    return days < AGO_CEILING_DAYS ? days + ' d ago' : dayMonth(d, 'UTC');
  }

  const sameDay = (a, b, zone) => {
    const x = zoned(a, zone);
    const y = zoned(b, zone);
    return x.year === y.year && x.month === y.month && x.day === y.day;
  };

  // when is "Today 8:14 pm", "Yesterday 8:14 pm" or "Tue 14 Oct, 8:14 pm".
  function when(iso, now, words, locale, zone) {
    const d = new Date(iso);
    if (!valid(d)) return '';
    const time = clock(d, zone, locale);
    if (sameDay(d, now, zone)) return words.today + ' ' + time;
    if (sameDay(d, new Date(now.getTime() - DAY_MS), zone)) return words.yesterday + ' ' + time;
    return weekdayDayMonth(d, zone, now) + ', ' + time;
  }

  // lastPlayed is "Last played 7:42 pm" today, "Last played 15 Oct, 7:42 pm" another day, and
  // "Not played yet" when the time is not known.
  function lastPlayed(iso, now, words, locale, zone) {
    const d = iso ? new Date(iso) : null;
    if (!valid(d)) return words.notPlayed;
    if (sameDay(d, now, zone)) return words.lastPlayed + ' ' + clock(d, zone, locale);
    return words.lastPlayed + ' ' + stamp(d, zone, locale);
  }

  // ago is how long since iso, the site's compact age; an unreadable time is an empty string.
  function ago(iso, now) {
    return compactAgo(new Date(iso), now);
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

  // usualCredit is the name most of the listed reports are credited to: the
  // rows leave it out and name only a different character, or a missing one.
  function usualCredit(reports) {
    const counts = new Map();
    for (const r of reports || []) {
      const c = r && r.credit;
      if (c && c.state === 'credited' && c.name) counts.set(c.name, (counts.get(c.name) || 0) + 1);
    }
    let best = '';
    for (const [name, n] of counts) if (!best || n > counts.get(best)) best = name;
    return best;
  }

  // addonSummary folds version lines that all say "up to date" into one,
  // "Addon 2.0.16 · Data 2026.11.09 · up to date"; any other set stays as it is.
  function addonSummary(lines, suffix) {
    const tail = ' · ' + suffix;
    if (lines.length < 2 || !lines.every((l) => !l.muted && l.text.endsWith(tail))) return lines;
    return [{ text: lines.map((l) => l.text.slice(0, -tail.length)).join(' · ') + tail, muted: false }];
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
  // "... · nothing new"; the time is the site's compact age.
  function checkedLine(checked, now, words) {
    const when = ago(checked.at, now);
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

  // fightRows says what the Fights uploaded and Queued rows hold, so the two
  // never count the same fights: uploaded is the fights the site has (every fight
  // the report closed, less those still waiting), queued is the fights waiting
  // while some are already up. While none are up yet the one row says they are
  // on their way. With no report open there are no rows, the Current report
  // row has said "None yet".
  function fightRows(p, w) {
    if (!p.report_key && !p.report_id) return { show: false };
    const total = p.fights || 0;
    const waiting = Math.min(p.queued_fights || 0, total);
    const up = total - waiting;
    if (waiting > 0 && up === 0) return { show: true, text: w.uploadingFights(waiting), help: '', queued: 0 };
    if (total === 0) {
      return p.report_id && p.logging
        ? { show: true, text: w.noFightsYet, help: w.noFightsHelp, queued: 0 }
        : { show: true, text: w.noReport, help: '', queued: 0 };
    }
    return { show: true, text: String(up), help: '', queued: waiting };
  }

  // addonRows reads the snapshot's addon_update: the Addon row's version lines (each worded by
  // the Go side) and the Install the addon into <folder> buttons. A snapshot from before it
  // existed has neither. A line with no text is dropped; it is never rebuilt from a number.
  function addonRows(au) {
    const lines = ((au && au.lines) || []).filter((l) => l && l.text).map((l) => ({ text: l.text, muted: l.tone === 'muted' }));
    const offers = ((au && au.offers) || []).filter((o) => o && o.path && o.label).map((o) => ({ path: o.path, label: o.label }));
    return { lines, offers };
  }

  return { addonRows, addonSummary, usualCredit, fightRows, pathTruncate, pathBudget, reportCredit, saveStrip, when, lastPlayed, ago, clock, dayMonth, weekdayDayMonth, longDay, stamp, compactAgo, SITE_LOCALE, creditText, currentReport, advancedView, checker, checkedLine, characterLabel, characterValue, automaticLabel, countdown, secondsLeft, expiryShare, lastMinute };
})();

if (typeof module !== 'undefined') module.exports = FORMAT;
