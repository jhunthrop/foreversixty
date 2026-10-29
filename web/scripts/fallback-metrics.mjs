// web/scripts/fallback-metrics.mjs
// Computes the size-adjusted fallback @font-face descriptors in src/styles/fonts.css: the
// `size-adjust`, `ascent-override`, `descent-override` and `line-gap-override` that make a
// system face occupy the same em-box and average advance as the web face it stands in for,
// so `font-display: swap` moves no line when the real face arrives (tenet 13).
//
// The formula is capsize's createFontStack (@capsizecss/core), applied to metrics read with
// @capsizecss/unpack: size-adjust is the ratio of the two faces' latin-weighted average
// advance widths, and each override is the web face's own vertical metric divided by its
// units-per-em scaled by that size-adjust.
//
// Usage, from web/ (the system faces are macOS paths; pass others as arguments):
//   node scripts/fallback-metrics.mjs <web-face.woff2> <system-face.ttf|ttc> [postscriptName]
// Reproduces `Barlow Fallback` against Helvetica Neue:
//   node scripts/fallback-metrics.mjs src/assets/fonts/barlow-400-subset.woff2 \
//     /System/Library/Fonts/HelveticaNeue.ttc HelveticaNeue
import { fromFile } from '@capsizecss/unpack/fs';

const PERCENT_DECIMALS = 2;

function percent(ratio) {
  return `${(ratio * 100).toFixed(PERCENT_DECIMALS)}%`;
}

/** capsize's createFontStack overrides for `fallback` standing in for `web`. */
export function fallbackDescriptors(web, fallback) {
  const sizeAdjust = web.xWidthAvg / web.unitsPerEm / (fallback.xWidthAvg / fallback.unitsPerEm);
  const scaledEm = web.unitsPerEm * sizeAdjust;
  return {
    sizeAdjust: percent(sizeAdjust),
    ascentOverride: percent(web.ascent / scaledEm),
    descentOverride: percent(Math.abs(web.descent) / scaledEm),
    lineGapOverride: percent(web.lineGap / scaledEm),
  };
}

async function main() {
  const [webPath, fallbackPath, postscriptName] = process.argv.slice(2);
  if (!webPath || !fallbackPath) {
    throw new Error('usage: node scripts/fallback-metrics.mjs <web-face> <system-face> [postscriptName]');
  }
  const web = await fromFile(webPath);
  const fallback = await fromFile(fallbackPath, postscriptName ? { postscriptName } : undefined);
  const d = fallbackDescriptors(web, fallback);
  console.log(`/* ${web.familyName} over ${fallback.fullName} */`);
  console.log(`  size-adjust: ${d.sizeAdjust};`);
  console.log(`  ascent-override: ${d.ascentOverride};`);
  console.log(`  descent-override: ${d.descentOverride};`);
  console.log(`  line-gap-override: ${d.lineGapOverride};`);
}

await main();
