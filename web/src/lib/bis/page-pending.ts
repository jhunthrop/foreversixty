// web/src/lib/bis/page-pending.ts
// The BiS page's head script: a deep link with a non-default state (a ?band= or other state
// param, the legacy #band-<faction>-<band> hash, or a remembered Sources choice) marks
// html[data-bis-pending], which hides the regions that change (geometry kept) until the island
// applies the state on mount, so the default render never visibly swaps. A failed hydration
// must never leave the page blank, so the mark clears itself after PENDING_TIMEOUT_MS.
import { STATE_PARAMS } from './page-state';
import { SOURCES_STORAGE_KEYS } from './source-store';

export const PENDING_ATTRIBUTE = 'data-bis-pending';
const PENDING_TIMEOUT_MS = 4000;

export const BIS_PENDING_SCRIPT = `(function(){try{var d=document.documentElement,s=new URLSearchParams(location.search),p=${JSON.stringify(
  STATE_PARAMS,
)}.some(function(k){return s.has(k)})||/^#band-/.test(location.hash);if(!p){var k=${JSON.stringify(
  SOURCES_STORAGE_KEYS,
)};for(var i=0;i<k.length;i++)if(localStorage.getItem(k[i])!==null)p=true}if(p){d.setAttribute('${PENDING_ATTRIBUTE}','');setTimeout(function(){d.removeAttribute('${PENDING_ATTRIBUTE}')},${PENDING_TIMEOUT_MS})}}catch(e){}})();`;
