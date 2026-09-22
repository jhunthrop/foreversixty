// web/src/report-island.ts
// Entry point for dist/report-island.js, which every report shell loads:
//   <div id="report" data-report-mount data-report-id="…" data-report='…'></div>
//   <link rel="stylesheet" href="https://foreversixty.gg/report-island.css">
//   <script type="module" src="https://foreversixty.gg/report-island.js"></script>
//
// The two stylesheet imports are load-bearing for the same reason they are in
// src/planner-island.ts: report-island.css is the only stylesheet the shell links, so it
// has to carry the tokens, the base rules and the self-hosted faces or the page renders in
// Georgia and Arial.
//
// data-report is an optional bootstrap of GET /v1/reports/{id}'s payload. The prerendered
// fixture page inlines it so the page renders from static files alone -- which is what
// Lighthouse measures -- exactly as the API inlines data-build on /b/:id. Worker-served
// shells carry only the id and the island fetches the rest, so the shell stays small
// enough to be worth caching.
import { mount } from 'svelte';
import ReportView from './components/report/ReportView.svelte';
import { scheduleBoot } from './lib/islands/boot';
import type { ReportMeta } from './lib/report/types';
import './styles/fonts.css';
import './styles/global.css';

const MOUNT_ID = 'report';
const ID_IN_PATH = /^\/reports\/([a-z2-7]{12})\/?$/;

/** Exported for the unit test; the shell and the Worker both feed one of these two. */
export function reportIdFrom(element: HTMLElement, pathname: string): string {
  const fromData = element.dataset.reportId;
  if (fromData !== undefined && fromData !== '') return fromData;
  return ID_IN_PATH.exec(pathname)?.[1] ?? '';
}

function readInlineMeta(element: HTMLElement): ReportMeta | null {
  const raw = element.dataset.report;
  if (raw === undefined || raw === '') return null;
  try {
    return JSON.parse(raw) as ReportMeta;
  } catch (error) {
    console.error('report island: data-report is not valid JSON', error);
    return null;
  }
}

function boot(): void {
  const target = document.getElementById(MOUNT_ID);
  if (target === null) return;

  const reportId = reportIdFrom(target, window.location.pathname);
  const inlineMeta = readInlineMeta(target);

  // The shell's own markup is an empty state for people with no JavaScript and for the
  // moment before this runs. Svelte's mount appends, so it has to go first.
  target.replaceChildren();
  mount(ReportView, { target, props: { reportId, inlineMeta } });
}

scheduleBoot(boot);
