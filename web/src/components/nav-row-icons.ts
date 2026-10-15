// The 20px gold stroke icon of each panel row (spec 2026-10-15-grouped-header.md section 3),
// as SVG path data on a 24px grid, keyed by the row's label in lib/nav.ts.
export const NAV_ROW_ICON_PATHS: Readonly<Record<string, string>> = {
  Planner: 'M12 3v6M12 9l-5 5M12 9l5 5M7 14v6M17 14v6',
  BiS: 'M12 3l7 3v5c0 5-3 8-7 10-4-2-7-5-7-10V6z',
  Simulator: 'M13 2L4 14h7l-1 8 9-12h-7z',
  'Get set up': 'M12 4v11M7 11l5 5 5-5M5 20h14',
  'Upload a log': 'M12 20V9M7 13l5-5 5 5M5 4h14',
  'Your reports': 'M7 3h8l4 4v14H7zM15 3v4h4M10 12h6M10 16h6',
  Rankings: 'M5 20v-8M12 20V5M19 20V9',
  'Tier list': 'M4 6h16M4 12h12M4 18h8',
  'Your guild': 'M6 3h12v12l-6 6-6-6z',
  Roster: 'M9 11a3 3 0 100-6 3 3 0 000 6zM3 20c0-3 3-5 6-5s6 2 6 5M17 11a3 3 0 000-6M17 15c2 0 4 2 4 5',
  'Loot council': 'M6 4h12l3 5-9 11L3 9zM3 9h18',
  'Guild logs': 'M4 5h12a3 3 0 013 3v11H7a3 3 0 01-3-3zM8 10h7M8 14h5',
};
