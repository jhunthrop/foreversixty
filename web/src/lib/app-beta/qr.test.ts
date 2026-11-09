import { describe, expect, it } from 'vitest';
import { APP_PAGE_URL } from './copy';
import { QR_OPTIONS, appQrSvg } from './qr';
import config from './qr.json';

describe('the QR for /app', () => {
  it('is an SVG drawn here, dark modules on a light quiet zone', async () => {
    const svg = await appQrSvg();
    expect(svg.startsWith('<svg')).toBe(true);
    expect(svg).toContain(QR_OPTIONS.color.dark.slice(0, 7));
    expect(svg).toContain(QR_OPTIONS.color.light.slice(0, 7));
    expect(svg).not.toMatch(/https?:\/\/(?!www\.w3\.org)/);
  });

  it('is the same code every build', async () => {
    expect(await appQrSvg()).toBe(await appQrSvg());
  });

  it('encodes the permanent link, the one the page and the PNG script both read', async () => {
    expect(APP_PAGE_URL).toBe('https://foreversixty.gg/app');
    expect(config.url).toBe(APP_PAGE_URL);
    expect(await appQrSvg()).toBe(await appQrSvg(APP_PAGE_URL));
  });
});
