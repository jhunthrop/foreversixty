// design/mocks/2026-11-02-class-landing/render.mjs
// Renders mock.html (spec design/specs/2026-11-02-class-landing.md §16) full page at the four
// review widths, then writes a same-scale side-by-side of each width's mock beside a built
// capture when one exists in design/reviews/2026-11-02-class-landing/<prefix>-hunter-<width>.png.
//
//   node design/mocks/2026-11-02-class-landing/render.mjs [prefix]
//     -> design/mocks/2026-11-02-class-landing/mock-<width>.png
//     -> design/reviews/2026-11-02-class-landing/side-by-side-<prefix>-<width>.png
import { createRequire } from "node:module";
import { existsSync, readFileSync } from "node:fs";
import { fileURLToPath } from "node:url";

const require = createRequire(
  new URL("../../../web/package.json", import.meta.url),
);
const { chromium } = require("@playwright/test");

const HERE = fileURLToPath(new URL("./", import.meta.url));
const REVIEWS = fileURLToPath(
  new URL("../../reviews/2026-11-02-class-landing/", import.meta.url),
);
const PREFIX = process.argv[2] ?? "after";
const WIDTHS = [390, 1024, 1440, 2000];
const GAP = 24;

const browser = await chromium.launch();
try {
  for (const width of WIDTHS) {
    const page = await browser.newPage({ viewport: { width, height: 900 } });
    await page.goto(`file://${HERE}mock.html`, { waitUntil: "networkidle" });
    const mock = `${HERE}mock-${width}.png`;
    await page.screenshot({ path: mock, fullPage: true });
    const built = `${REVIEWS}${PREFIX}-hunter-${width}.png`;
    if (existsSync(built)) {
      const uri = (file) =>
        `data:image/png;base64,${readFileSync(file).toString("base64")}`;
      await page.setViewportSize({ width: width * 2 + GAP, height: 900 });
      await page.setContent(
        `<body style="margin:0;background:#333;display:flex;gap:${GAP}px;align-items:flex-start">` +
          `<img src="${uri(mock)}" width="${width}"><img src="${uri(built)}" width="${width}"></body>`,
      );
      await page.screenshot({
        path: `${REVIEWS}side-by-side-${PREFIX}-${width}.png`,
        fullPage: true,
      });
    }
    await page.close();
    console.log(mock);
  }
} finally {
  await browser.close();
}
