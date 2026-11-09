// The app mocks' fonts from the repo, not Google: the site's own subsets (web/src/assets/fonts), so a capture and a
// region file come out the same offline and in CI (design/parity.md "Native"). The faces go in as data: URLs, as the web
// parity helper serves them (web/tests/e2e/support/parity.ts): the @font-face rules arrive in a stylesheet answered for
// fonts.googleapis.com, and Chromium refuses a file:// font from an https stylesheet, so file URLs fell back to Georgia
// and Arial without a word (2026-11-08).
import { readFileSync } from "node:fs";

// The apps still set Barlow (round 5 moved only the site to Open Sans; the mobile lane follows), so the app mocks keep
// the Barlow subsets the site shipped until then, kept beside the mocks in data/fonts.
const SITE_FONTS = new URL("../../web/src/assets/fonts/", import.meta.url);
const APP_FONTS = new URL("./data/fonts/", import.meta.url);
const FACES = [
  ["Barlow", 400, "barlow-400-subset.woff2", APP_FONTS],
  ["Barlow", 600, "barlow-600-subset.woff2", APP_FONTS],
  ["Barlow", 700, "barlow-700-subset.woff2", APP_FONTS],
  ["Marcellus", 400, "marcellus-400-subset.woff2", SITE_FONTS],
];
const CSS = FACES.map(([family, weight, file, dir]) => {
  const data = readFileSync(new URL(file, dir)).toString("base64");
  return `@font-face{font-family:'${family}';font-weight:${weight};font-style:normal;font-display:block;src:url(data:font/woff2;base64,${data}) format('woff2')}`;
}).join("\n");

/** Answers the mocks' Google Fonts request with the repo's faces and drops every other network request. One handler:
 *  Playwright runs the route registered last first, so the catch-all abort registered after the fonts route used to
 *  swallow the stylesheet as well, and the mocks drew in Georgia and Arial. */
export async function routeFonts(context) {
  await context.route(/^https?:\/\//, (r) =>
    r.request().url().startsWith("https://fonts.googleapis.com/") ? r.fulfill({ contentType: "text/css", body: CSS }) : r.abort(),
  );
}

/** After `document.fonts.ready`: throws unless every face loaded, so a mock can never be measured or captured in a
 *  fallback face again. `file` names the mock in the error. */
export async function assertFonts(page, file) {
  const problems = await page.evaluate(async (faces) => {
    // Load every declared face, used on this page or not, so a broken one cannot hide behind an unused weight.
    await Promise.all([...document.fonts].map((f) => f.load().catch(() => null)));
    const declared = [...document.fonts].map((f) => ({ family: f.family.replace(/['"]/g, ""), weight: String(f.weight), status: f.status }));
    return faces.flatMap(([family, weight]) => {
      const face = declared.find((f) => f.family === family && f.weight === String(weight));
      if (!face) return [`${family} ${weight} not declared`];
      return face.status === "loaded" ? [] : [`${family} ${weight} ${face.status}`];
    });
  }, FACES.map(([family, weight]) => [family, weight]));
  if (problems.length) throw new Error(`${file}: the mock's fonts did not load (${problems.join("; ")})`);
}
