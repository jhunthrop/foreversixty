// web/src/lib/app-beta/qr.ts
// The QR code for foreversixty.gg/app, drawn at build time by the qrcode package (no QR service,
// nothing fetched). The page inlines the SVG; scripts/app-qr-png.mjs writes the same code from the
// same qr.json as a PNG to design/assets/app-qr.png for the addon. Medium error correction, the
// four-module quiet zone, near-black on the site's parchment so every phone camera reads it.
import QRCode from 'qrcode';
import config from './qr.json';

export const QR_OPTIONS = {
  errorCorrectionLevel: config.errorCorrectionLevel as 'M',
  margin: config.margin,
  color: config.color,
} as const;

/** The code as an SVG string, scaled by its viewBox. */
export function appQrSvg(url: string = config.url): Promise<string> {
  return QRCode.toString(url, { ...QR_OPTIONS, type: 'svg' });
}
