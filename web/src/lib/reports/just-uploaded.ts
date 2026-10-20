// web/src/lib/reports/just-uploaded.ts
// The upload form marks the report id it is about to open; the report page reads the mark
// to treat a 404 for that id as "not visible yet" for a short while (load.ts,
// fetchReportMetaAfterUpload). sessionStorage, so the mark dies with the tab and a
// stranger's link to a missing report is never given the grace.
const PREFIX = 'fs.justUploaded.';

function store(): Storage | null {
  try {
    return globalThis.sessionStorage ?? null;
  } catch {
    return null;
  }
}

export function markJustUploaded(reportId: string, now: number = Date.now()): void {
  try {
    store()?.setItem(PREFIX + reportId, String(now));
  } catch {
    // Private mode: the page simply has no grace.
  }
}

/** When this id was uploaded from this tab, or null when it was not. */
export function justUploadedAt(reportId: string): number | null {
  try {
    const raw = store()?.getItem(PREFIX + reportId);
    if (raw === null || raw === undefined) return null;
    const at = Number(raw);
    return Number.isFinite(at) ? at : null;
  } catch {
    return null;
  }
}
