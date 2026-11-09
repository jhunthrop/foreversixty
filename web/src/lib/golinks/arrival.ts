// web/src/lib/golinks/arrival.ts
// The one part of the tracked links that has to run before anything else: a /go/ landing carries
// ?golink=<slug>, which is kept (localStorage, as the newest click) and taken out of the address bar
// before any island reads or rewrites the URL. Base.astro inlines it at the end of every page.
//
// It is a few hundred bytes, written from config.ts's constants here, because every page's HTML
// carries it: Lighthouse's simulated load adds a whole round trip to the BiS pages' largest paint
// once their HTML grows past the next congestion-window step, and 1.4 KB more did exactly that.
// Everything that can wait (the arrival send and the visit beacon) is lib/golinks/later.ts,
// imported after paint by the header's loader. Nothing here sets a cookie.
import { ARRIVAL_NEXT_KEY, LINK_PARAM, SLUG_PATTERN } from './config';

const param = JSON.stringify(LINK_PARAM);

export const ARRIVAL_CAPTURE_SCRIPT =
  '(function(){try{var u=new URL(location.href),s=u.searchParams.get(' +
  param +
  ');if(s===null)return;if(new RegExp(' +
  JSON.stringify(SLUG_PATTERN) +
  ').test(s))localStorage.setItem(' +
  JSON.stringify(ARRIVAL_NEXT_KEY) +
  ',JSON.stringify({slug:s,at:Date.now()}));u.searchParams.delete(' +
  param +
  ');history.replaceState(history.state,"",u.pathname+u.search+u.hash)}catch(e){}})()';
