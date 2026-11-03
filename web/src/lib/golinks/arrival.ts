// web/src/lib/golinks/arrival.ts
// The base layout's arrival script and the data it reads. The script itself is plain browser
// JavaScript in arrival.inline.js, inlined into every page (Base.astro) rather than shipped as
// a module: it has to take ?golink= out of the address bar before any island reads the URL, and
// an inline script costs no request ahead of the largest paint. Its constants come from
// config.ts through these data attributes, so the two can never disagree.
import ARRIVAL_SOURCE from './arrival.inline.js?raw';
import {
  ARRIVAL_KEY,
  ARRIVAL_PATH,
  ARRIVAL_TTL_MS,
  BEACON_HOST,
  BEACON_SKIP_PREFIXES,
  BEACON_SRC,
  BEACON_TOKEN,
  LINK_PARAM,
  SLUG_PATTERN,
} from './config';

export const ARRIVAL_SCRIPT: string = ARRIVAL_SOURCE;

/**
 * The beat after `load` before the beacon and the arrival go out: the same 800 ms
 * lib/after-paint.ts gives the page's own deferred work, so neither competes with first paint.
 */
export const ARRIVAL_BEAT_MS = 800;

/** The script tag's data attributes, by attribute name. */
export function arrivalScriptData(apiBase: string): Record<string, string> {
  return {
    'data-param': LINK_PARAM,
    'data-pattern': SLUG_PATTERN,
    'data-key': ARRIVAL_KEY,
    'data-ttl': String(ARRIVAL_TTL_MS),
    'data-api': apiBase,
    'data-path': ARRIVAL_PATH,
    'data-beat': String(ARRIVAL_BEAT_MS),
    'data-beacon-src': BEACON_SRC,
    'data-beacon-token': BEACON_TOKEN,
    'data-beacon-host': BEACON_HOST,
    'data-beacon-skip': BEACON_SKIP_PREFIXES.join(','),
  };
}
