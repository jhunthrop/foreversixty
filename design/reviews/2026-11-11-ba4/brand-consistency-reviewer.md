# brand-consistency-reviewer: companion after BA4 (2026-11-11)

Read: styles.css (no hex or rgba literals; px left are layout only), shell_gui.go (tray reads copygen.TrayOpen, ReportsOpen, TrayQuit), companion.json keys, live /reports ("Your reports"), captures logging-credited, reports, dmg-window, install-setup, tray-menu.

## Prior findings
1. Tray wording hardcoded: FIXED. Tray reads generated copygen; the job is "Open your reports" in tray, window and page title.
2. Formats twice: formats now render in the shared style ("Today 6:40 pm", "1 min ago", "Sun 18 Oct, 2:40 pm"); captures agree with the web.
3. Copy not keyed to shared: FIXED for the companion via companion.json plus generator (shared nouns are still not drawn from shared.json; NICE).
4. DMG light ground: ACCEPTED as ALLOWED platform idiom. Finder draws dark label ink over a background picture in both appearances, so a dark ground would be unreadable. The dark brand lives in the painted band and seal. Seal is the painted variant at that size, correct.
5. Windows setup: on-brand, matches page-setup.png; native title bar allowed.
NICE items (font literals, raw radii, veils, shadows): FIXED via tokens (--color-veil-*, --shadow-heading, --radius-round).

## DRIFT
BLOCKER: none.
SHOULD: none.
NICE
- The corner filigree on the lead panel is still the most ornate element in the product; the web's dense plates have none. Owner call, unchanged from the last audit.
- reports.png shows the job link "Open your reports" beside "Your reports" at the top; the web page only shows the title. Harmless but redundant on the same screen.
- Mac DMG Applications folder in the mock is a generic blue tile; fine, Finder draws the real one.
- Confirm the tray dot tones still match --tone-* by convention only (no test).

VERDICT: SHIP (CONSISTENT)
