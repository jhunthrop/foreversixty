// web/src/lib/reports/layout.ts
// The height the two report feeds -- RecentReports.svelte ("Recent public reports") and
// MyReports.svelte ("Your reports") -- reserve while loading. One module rather than two
// inline literals so the pair cannot drift apart, which is guild/layout.ts's own reason.
//
// Both skeletons draw five rows, which is not the page size: MyReports pages
// REPORTS_PER_PAGE (100) rows at a time and the full RecentReports mount has no cap at
// all, so the row count is a legible stand-in and this reserve, not the rows, is what has
// to hold. A real ReportRow is min-h-11 with py-2 and a border, so ~61px on desktop and
// ~77px once its title wraps at 360px; 360px is five of those at the phone height, the
// screenful a reader sees before scrolling. Reserving to the taller of the two is
// deliberate -- over-reserving costs dead space, under-reserving costs the CLS budget,
// and /logs.html is one of the 13 URLs lighthouserc.json gates at CLS 0.05.
export const REPORTS_LOADING_MIN_H = 'min-h-[360px]';
