// web/src/lib/guides/page-script.ts
// The one small script a spec guide runs, inline and after the page has loaded (the same
// "after paint" rule W.16 set: nothing here is needed for first paint, and the page works
// without it). It does two things:
//   1. Opens Best in slot on the visitor's own level range. The current-character pointer in
//      localStorage names a character; when it is an addon or code build of this guide's class
//      (`ref` is an FS1 code), its spent talent points give its level (points + 9) and so its
//      range, and that radio is checked. Nothing stored, another class or another source: the
//      table stays on 60.
//   2. Lights the contents rail's current section: `aria-current="location"` on the link whose
//      heading is the last one above the upper part of the viewport, recomputed whenever a
//      heading crosses it (IntersectionObserver).
export function ownBand(ref: string, classSlug: string): number | undefined {
  const parts = ref.split(':');
  if (parts[0] !== 'FS1' || parts[2] !== classSlug) return undefined;
  let points = 0;
  for (const tree of (parts[4] ?? '').split('|')[0]!.split('/')) {
    for (const digit of tree) {
      const value = Number(digit);
      if (Number.isInteger(value)) points += value;
    }
  }
  const level = Math.min(60, points + 9);
  return level >= 60 ? 60 : Math.max(20, Math.floor(level / 10) * 10);
}

export const pageScript = `(function () {
  var run = function () {
    try {
      var guide = document.querySelector('[data-testid="guide-bis"]');
      var raw = localStorage.getItem('fs.currentCharacter');
      if (guide && raw) {
        var c = JSON.parse(raw);
        if (c && (c.source === 'addon' || c.source === 'code') && typeof c.ref === 'string') {
          var p = c.ref.split(':');
          if (p[0] === 'FS1' && p[2] === guide.getAttribute('data-guide-class')) {
            var pts = 0;
            (p[4] || '').split('|')[0].split('/').forEach(function (s) {
              for (var i = 0; i < s.length; i++) { var d = Number(s.charAt(i)); if (d >= 0 && d <= 9) pts += d; }
            });
            var level = Math.min(60, pts + 9);
            var band = level >= 60 ? 60 : Math.max(20, Math.floor(level / 10) * 10);
            var radio = document.getElementById('gbis-band-' + band);
            if (radio) radio.checked = true;
          }
        }
      }
    } catch (e) {}
    try {
      var links = {};
      document.querySelectorAll('[data-toc-link]').forEach(function (a) { links[a.getAttribute('data-toc-link')] = a; });
      var heads = Object.keys(links).map(function (id) { return document.getElementById(id); }).filter(Boolean);
      if (!heads.length || !('IntersectionObserver' in window)) return;
      var mark = function () {
        var current = heads[0];
        heads.forEach(function (h) { if (h.getBoundingClientRect().top <= window.innerHeight * 0.35) current = h; });
        Object.keys(links).forEach(function (id) {
          if (id === current.id) links[id].setAttribute('aria-current', 'location');
          else links[id].removeAttribute('aria-current');
        });
      };
      var io = new IntersectionObserver(mark, { rootMargin: '0px 0px -60% 0px' });
      heads.forEach(function (h) { io.observe(h); });
      window.addEventListener('hashchange', mark);
      mark();
    } catch (e) {}
  };
  var start = function () { setTimeout(run, 300); };
  if (document.readyState === 'complete') start(); else window.addEventListener('load', start);
})();`;
