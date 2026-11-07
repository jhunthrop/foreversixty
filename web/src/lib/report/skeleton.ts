// web/src/lib/report/skeleton.ts
// The report's loading frame is components/report/ReportSkeleton.svelte, drawn by the page shell
// before the island boots and by the island until the first summary arrives. Its heading carries
// this testid: the Worker writes the report's title into it (src/worker.ts), so the title paints
// on the first render instead of after the island hydrates.
export const REPORT_SHELL_TITLE_TESTID = 'report-shell-title';
