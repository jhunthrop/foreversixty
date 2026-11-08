// web/scripts/app-qr-png.mjs
// Writes the /app QR code as a PNG to design/assets/app-qr.png, for the addon (which cannot
// inline the page's SVG). The URL, error correction, quiet zone and colours are
// src/lib/app-beta/qr.json, the file the page's SVG (src/lib/app-beta/qr.ts) is drawn from.
//   node scripts/app-qr-png.mjs   (from web/; commit design/assets/app-qr.png)
import { mkdirSync, readFileSync } from 'node:fs';
import { dirname } from 'node:path';
import { fileURLToPath } from 'node:url';
import QRCode from 'qrcode';

const OUT = fileURLToPath(new URL('../../design/assets/app-qr.png', import.meta.url));
const config = JSON.parse(readFileSync(new URL('../src/lib/app-beta/qr.json', import.meta.url), 'utf8'));

mkdirSync(dirname(OUT), { recursive: true });
await QRCode.toFile(OUT, config.url, {
  errorCorrectionLevel: config.errorCorrectionLevel,
  margin: config.margin,
  color: config.color,
  scale: config.pngScale,
  type: 'png',
});
process.stdout.write(`${OUT}\n`);
