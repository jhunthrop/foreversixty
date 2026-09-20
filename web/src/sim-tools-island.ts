// web/src/sim-tools-island.ts
// Entry point for dist/sim-tools-island.js, the bundle /sim/gear, /sim/talents, /sim/drops
// and /sim/weights load.
//
// It is a separate build from sim-island for one measured reason: sim-island is budgeted at
// 90 KB gzipped and carries /sim's own 1,600 ms mobile LCP, and four candidate grids, a
// source picker and two results views do not fit under that. Nothing here is loaded by /sim.
import { mount } from 'svelte';
import ToolsView from './components/sim/tools/ToolsView.svelte';
import { TOOLS, type SimTool } from './lib/sim/bulk-store.svelte';
import './styles/fonts.css';
import './styles/global.css';

const MOUNT_ID = 'sim-tools';

function isTool(value: string): value is SimTool {
  return (TOOLS as readonly string[]).includes(value);
}

/**
 * The mount's own stamp wins; the path is the fallback, so a shell that forgot the
 * attribute still renders the right page. Anything else lands on Top Gear rather than
 * throwing: the value reaches a component name lookup and is attacker-influenced through
 * the URL, so it is validated against a closed list and never used raw.
 */
export function toolFrom(element: HTMLElement, pathname: string): SimTool {
  const stamped = element.dataset.simTool ?? '';
  if (isTool(stamped)) return stamped;
  const last =
    pathname
      .replace(/\.html$/, '')
      .split('/')
      .pop() ?? '';
  return isTool(last) ? last : 'gear';
}

function boot(): void {
  const target = document.getElementById(MOUNT_ID);
  if (target === null) return;
  const tool = toolFrom(target, window.location.pathname);
  // The shell's skeleton and no-JS paragraph live inside the mount; Svelte 5 appends rather
  // than replaces, so they go before the mount or they stay under the island.
  target.replaceChildren();
  mount(ToolsView, { target, props: { tool } });
}

if (document.readyState === 'loading') {
  document.addEventListener('DOMContentLoaded', boot, { once: true });
} else {
  boot();
}
