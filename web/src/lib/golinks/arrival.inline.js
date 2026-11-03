// web/src/lib/golinks/arrival.inline.js
// Inlined at the end of every page by Base.astro (see arrival.ts); its settings are the data
// attributes arrival.ts's arrivalScriptData writes. Three jobs, none of which sets a cookie
// or reads anything about the visitor:
//  1. A /go/ landing carries ?golink=<slug>: keep the slug (localStorage, an hour at most) and
//     take the parameter out of the address bar, so a shared or bookmarked URL never carries it.
//  2. On a signed-in page, send a kept slug to the API once, which records it as the account's
//     first link only when the account is new. Kept while signed out, for the sign-in.
//  3. Load the cookieless Cloudflare Web Analytics beacon, on the production host only.
// 2 and 3 wait for `load` and a beat after it, so neither competes with the first paint.
(function () {
  var el = document.currentScript;
  if (!el) return;
  var d = el.dataset;

  try {
    var url = new URL(location.href);
    var ref = url.searchParams.get(d.param);
    if (ref !== null) {
      if (new RegExp(d.pattern).test(ref)) {
        localStorage.setItem(d.key, JSON.stringify({ slug: ref, at: Date.now() }));
      }
      url.searchParams.delete(d.param);
      history.replaceState(history.state, '', url.pathname + url.search + url.hash);
    }
  } catch {
    /* storage or history unavailable: the visit simply goes unattributed */
  }

  function sendArrival() {
    try {
      var raw = localStorage.getItem(d.key);
      if (raw === null) return;
      var kept = JSON.parse(raw);
      var age = kept && typeof kept.at === 'number' ? Date.now() - kept.at : -1;
      if (!kept || typeof kept.slug !== 'string' || age < 0 || age >= Number(d.ttl)) {
        localStorage.removeItem(d.key);
        return;
      }
      var csrf = /(?:^|;\s*)fs_csrf=([^;]*)/.exec(document.cookie);
      if (csrf === null || csrf[1] === '') return;
      fetch(d.api + d.path, {
        method: 'POST',
        credentials: 'include',
        headers: { 'content-type': 'application/json', 'x-csrf-token': decodeURIComponent(csrf[1]) },
        body: JSON.stringify({ slug: kept.slug }),
      })
        .then(function (response) {
          // Sent (recorded or not), or refused for good. A 401 or 403 means the session is
          // gone, so the slug waits for the sign-in.
          if (response.ok || response.status === 400) localStorage.removeItem(d.key);
        })
        .catch(function () {
          /* offline: tried again on the next page */
        });
    } catch {
      /* storage unavailable or malformed: nothing to send */
    }
  }

  function loadBeacon() {
    if (location.hostname !== d.beaconHost) return;
    var skip = d.beaconSkip ? d.beaconSkip.split(',') : [];
    for (var i = 0; i < skip.length; i++) {
      if (location.pathname.indexOf(skip[i]) === 0) return;
    }
    var beacon = document.createElement('script');
    beacon.type = 'module';
    beacon.src = d.beaconSrc;
    beacon.setAttribute('data-cf-beacon', JSON.stringify({ token: d.beaconToken, spa: false }));
    document.head.appendChild(beacon);
  }

  function later() {
    loadBeacon();
    sendArrival();
  }

  function afterLoad() {
    setTimeout(function () {
      if ('requestIdleCallback' in window) window.requestIdleCallback(later);
      else later();
    }, Number(d.beat));
  }

  if (document.readyState === 'complete') afterLoad();
  else window.addEventListener('load', afterLoad, { once: true });
})();
