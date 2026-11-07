// The app mocks' fonts from the repo, not Google: the site's own subsets (web/src/assets/fonts), so a capture and a
// region file come out the same offline and in CI (design/parity.md "Native").
const FONTS = new URL("../../web/src/assets/fonts/", import.meta.url).href;
const FACES = [
  ["Barlow", 400, "barlow-400-subset.woff2"],
  ["Barlow", 600, "barlow-600-subset.woff2"],
  ["Barlow", 700, "barlow-700-subset.woff2"],
  ["Marcellus", 400, "marcellus-400-subset.woff2"],
];
const CSS = FACES.map(
  ([family, weight, file]) =>
    `@font-face{font-family:'${family}';font-weight:${weight};font-style:normal;font-display:block;src:url(${FONTS}${file}) format('woff2')}`,
).join("\n");

/** Answers the mocks' Google Fonts request with the repo's faces and drops every other network request. */
export async function routeFonts(context) {
  await context.route("https://fonts.googleapis.com/**", (r) => r.fulfill({ contentType: "text/css", body: CSS }));
  await context.route(/^https?:\/\//, (r) => r.abort());
}
