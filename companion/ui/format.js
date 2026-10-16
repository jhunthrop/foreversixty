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

  // characterLabel is "Thrall · Hardcore · US": the ruleset in words.
  function characterLabel(c, rulesets) {
    const ruleset = rulesets[String(c.ruleset).toLowerCase()] || c.ruleset;
    return c.name + ' · ' + ruleset + ' · ' + String(c.region).toUpperCase();
  }

  const characterValue = (c) => [c.region, c.ruleset, c.name].join('|');

  // countdown is the whole minutes and seconds left, never negative.
  function countdown(expiresISO, now) {
    const left = Math.max(0, Math.floor((new Date(expiresISO).getTime() - now.getTime()) / 1000));
    return { m: Math.floor(left / 60), s: left % 60 };
  }

  return { midTruncate, when, lastPlayed, characterLabel, characterValue, countdown };
})();

if (typeof module !== 'undefined') module.exports = FORMAT;
