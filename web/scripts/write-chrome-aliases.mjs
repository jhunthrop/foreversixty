// web/scripts/write-chrome-aliases.mjs
// Run from `postbuild`: writes dist/_chrome/<name>.js, the stable header scripts the API's
// /b/{id} page loads (scripts/chrome-aliases.mjs says why).
import { fileURLToPath } from 'node:url';
import { CHROME_DIR, writeChromeAliases } from './chrome-aliases.mjs';

const names = writeChromeAliases(fileURLToPath(new URL('../dist/', import.meta.url)));
console.log(`chrome-aliases: /${CHROME_DIR}/ carries ${names.join(', ')}.`);
