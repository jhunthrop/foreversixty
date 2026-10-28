# Task 27 report
Files: web/tests/captures/subscribe-and-join/{capture.spec.ts,playwright.capture.config.ts}. Lint/prettier clean; npm run check has no error in these files (one unrelated error in CheckoutGuildStep.test.ts:267).
Rows 1-17 covered (1/1b at 360/390/1024/1440/1920/2000). Output: design/reviews/2026-10-27-subscribe-and-join/.
Commands (from web/): npm run build; cp -R dist /tmp/sj-dist; npx astro preview --outDir /tmp/sj-dist --port 4483 & (kill by PID);
 CAPTURE_PORT=4483 npx playwright test -c tests/captures/subscribe-and-join/playwright.capture.config.ts
Row 17 baseline first: build main, serve, run with CAPTURE_VARIANT=before; then the default run diffs (fails on any differing pixel).
Concerns: no subscribe-and-join mock renders in design/mocks/renders (boards are written only if subscribe-and-join-<row>[-<state>].png exists); testids taken from committed components, unverified by a run.
