// web/src/lib/guides/page-script.ts
// The one small script a spec guide runs, inline and after the page has loaded (the same
// "after paint" rule W.16 set: nothing here is needed for first paint, and the page works
// without it): it lights the contents rail's current section, `aria-current="location"` on the
// link whose heading is the last one above the upper part of the viewport, recomputed whenever
// a heading crosses it (IntersectionObserver); a heading a filter hides is skipped (the
// changelog's months), and the section an address points at stays marked while it is near the top
// (an arrival at `#cookies` marks Cookies, not the next head up the screen). The Best in slot band follows the answer's
// character instead (bis-lane.ts).
export const pageScript = `(function () {
  var run = function () {
    try {
      var links = {};
      document.querySelectorAll('[data-toc-link]').forEach(function (a) { links[a.getAttribute('data-toc-link')] = a; });
      var heads = Object.keys(links).map(function (id) { return document.getElementById(id); }).filter(Boolean);
      if (!heads.length || !('IntersectionObserver' in window)) return;
      var mark = function () {
        var current = heads.filter(function (h) { return h.offsetParent !== null; })[0] || heads[0];
        heads.forEach(function (h) { if (h.offsetParent !== null && h.getBoundingClientRect().top <= window.innerHeight * 0.35) current = h; });
        var arrived = location.hash ? document.getElementById(decodeURIComponent(location.hash.slice(1))) : null;
        if (arrived && links[arrived.id] && Math.abs(arrived.getBoundingClientRect().top) <= window.innerHeight * 0.35) current = arrived;
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
