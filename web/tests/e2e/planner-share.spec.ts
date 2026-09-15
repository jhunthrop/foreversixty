import { crc32, deflateSync } from 'node:zlib';
import { expect, test } from '@playwright/test';

import { ACTIVE_BUILD } from './support/active-build';

const SAVED = {
  ok: true,
  data: { id: 'k7x2qm4a', url: 'https://foreversixty.gg/b/k7x2qm4a' },
  error: null,
  request_id: 'req-1',
};

/**
 * Builds a minimal, valid solid-colour PNG of exactly `width` x `height`, with only Node's
 * built-in `zlib` (no image library dependency). Used to prove the preview card's box holds
 * its aspect ratio once a real 600x315 image decodes -- a 1x1 placeholder would not: the
 * attribute-mapped ratio only holds until the browser learns the image's real one.
 */
function chunk(type: string, data: Buffer): Buffer {
  const typeBuf = Buffer.from(type, 'ascii');
  const length = Buffer.alloc(4);
  length.writeUInt32BE(data.length);
  const crc = Buffer.alloc(4);
  crc.writeUInt32BE(crc32(Buffer.concat([typeBuf, data])) >>> 0);
  return Buffer.concat([length, typeBuf, data, crc]);
}

function solidPng(width: number, height: number): Buffer {
  const signature = Buffer.from([0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a]);
  const ihdr = Buffer.alloc(13);
  ihdr.writeUInt32BE(width, 0);
  ihdr.writeUInt32BE(height, 4);
  ihdr.set([8, 2, 0, 0, 0], 8); // 8-bit depth, RGB colour type, default compression/filter/interlace
  const rowBytes = 1 + width * 3;
  const raw = Buffer.alloc(height * rowBytes);
  for (let y = 0; y < height; y += 1) {
    const row = y * rowBytes;
    for (let x = 0; x < width; x += 1) {
      const pixel = row + 1 + x * 3;
      raw.set([200, 60, 60], pixel);
    }
  }
  return Buffer.concat([
    signature,
    chunk('IHDR', ihdr),
    chunk('IDAT', deflateSync(raw)),
    chunk('IEND', Buffer.alloc(0)),
  ]);
}

test('sharing a build posts the contract body and shows the link', async ({ page }) => {
  let body: unknown;
  await page.route('**/v1/builds', async (route) => {
    body = route.request().postDataJSON();
    await route.fulfill({ status: 201, contentType: 'application/json', body: JSON.stringify(SAVED) });
  });

  await page.goto('/planner');
  for (let i = 0; i < 3; i += 1) await page.getByTestId('talent-1001').click();
  await page.getByLabel('Title').fill('Arms leveling');
  await page.getByRole('button', { name: 'Share' }).click();

  await expect(page.getByTestId('share-link')).toHaveText('https://foreversixty.gg/b/k7x2qm4a');
  expect(body).toEqual({
    class_id: 1,
    race_id: 1,
    tree_version: ACTIVE_BUILD,
    point_order: [1001, 1001, 1001],
    gear: {},
    title: 'Arms leveling',
  });
});

test('the copy button reports that it copied', async ({ page, context, browserName }) => {
  test.skip(browserName !== 'chromium', 'clipboard permissions are Chromium-only here');
  await context.grantPermissions(['clipboard-read', 'clipboard-write']);
  await page.route('**/v1/builds', (route) =>
    route.fulfill({ status: 201, contentType: 'application/json', body: JSON.stringify(SAVED) }),
  );
  await page.goto('/planner');
  await page.getByTestId('talent-1001').click();
  await page.getByRole('button', { name: 'Share' }).click();
  await page.getByRole('button', { name: 'Copy link' }).click();
  await expect(page.getByRole('button', { name: 'Copied' })).toBeVisible();
});

test('a rate-limited save keeps the build and explains the wait', async ({ page }) => {
  await page.route('**/v1/builds', (route) =>
    route.fulfill({
      status: 429,
      contentType: 'application/json',
      body: JSON.stringify({ ok: false, data: null, error: { message: 'slow down' }, request_id: 'r' }),
    }),
  );
  await page.goto('/planner');
  await page.getByTestId('talent-1001').click();
  await page.getByRole('button', { name: 'Share' }).click();
  await expect(page.getByRole('alert')).toContainText(
    'Too many saves from this connection; try again in an hour.',
  );
  await expect(page.getByTestId('planner-spent')).toHaveText('1/51');
});

test('a rejected build shows the API field message and retries', async ({ page }) => {
  let attempts = 0;
  await page.route('**/v1/builds', async (route) => {
    attempts += 1;
    if (attempts === 1) {
      return route.fulfill({
        status: 400,
        contentType: 'application/json',
        body: JSON.stringify({
          ok: false,
          data: null,
          error: {
            message: 'build is not valid',
            fields: { 'point_order[0]': 'Talent 1001 is not in this class' },
          },
          request_id: 'r',
        }),
      });
    }
    return route.fulfill({ status: 201, contentType: 'application/json', body: JSON.stringify(SAVED) });
  });
  await page.goto('/planner');
  await page.getByTestId('talent-1001').click();
  await page.getByRole('button', { name: 'Share' }).click();
  await expect(page.getByRole('alert')).toContainText('Talent 1001 is not in this class');
  await page.getByRole('button', { name: 'Retry' }).click();
  await expect(page.getByTestId('share-link')).toBeVisible();
});

test('a save that resolves after the build changed is not shown as the current build', async ({ page }) => {
  let release: () => void = () => {};
  const held = new Promise<void>((resolve) => {
    release = resolve;
  });
  await page.route('**/v1/builds', async (route) => {
    await held;
    await route.fulfill({ status: 201, contentType: 'application/json', body: JSON.stringify(SAVED) });
  });

  await page.goto('/planner');
  await page.getByTestId('talent-1001').click();
  await page.getByRole('button', { name: 'Share' }).click();
  await expect(page.getByRole('button', { name: 'Saving' })).toBeVisible();

  // Edit the build while that save is still in flight.
  await page.getByTestId('talent-1001').click();

  release();
  await expect(page.getByRole('button', { name: 'Share' })).toBeVisible();
  await expect(page.getByTestId('share-link')).not.toBeVisible();
});

test('the preview card reserves its box before the image loads, and nothing above it moves', async ({
  page,
}) => {
  let releaseImage: () => void = () => {};
  const heldImage = new Promise<void>((resolve) => {
    releaseImage = resolve;
  });
  await page.route('**/v1/builds', (route) =>
    route.fulfill({ status: 201, contentType: 'application/json', body: JSON.stringify(SAVED) }),
  );
  // A real 600x315 image, not a 1x1 stand-in: the browser only honours the width/height
  // attributes' aspect ratio until it learns the image's real one, so a mismatched fixture
  // would report a stable box that a real card.png would not actually have.
  await page.route('**/b/k7x2qm4a/card.png', async (route) => {
    await heldImage;
    await route.fulfill({ status: 200, contentType: 'image/png', body: solidPng(600, 315) });
  });

  await page.goto('/planner');
  await page.getByTestId('talent-1001').click();
  await page.getByRole('button', { name: 'Share' }).click();
  await expect(page.getByTestId('share-link')).toBeVisible();

  // Measured after the Share click's own auto-scroll has already settled, so the only thing
  // that can move either box afterwards is the image itself deciding a different height.
  const card = page.getByAltText(/^Preview card for build/);
  await card.waitFor({ state: 'attached' });
  const titleBefore = await page.getByLabel('Title').boundingBox();
  const boxBefore = await card.boundingBox();

  releaseImage();
  await expect(card).toBeVisible();
  await page.waitForTimeout(100); // let the decoded image settle before re-measuring
  const boxAfter = await card.boundingBox();
  const titleAfter = await page.getByLabel('Title').boundingBox();

  expect(boxBefore).not.toBeNull();
  expect(boxAfter).not.toBeNull();
  expect(titleBefore).not.toBeNull();
  expect(titleAfter).not.toBeNull();
  expect(boxAfter!.height).toBeCloseTo(boxBefore!.height, 0);
  expect(titleAfter!.y).toBe(titleBefore!.y);
});
