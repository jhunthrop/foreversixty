// design/listing/check-store-text.mjs
// Checks the store listing text before the owner pastes it: each field's length against the
// store's limit, and the voice rules every player-facing line follows (no em or en dashes, no
// exclamation marks, none of the stock marketing words). Exits 1 on any failure.
// Run from anywhere: node design/listing/check-store-text.mjs
import { readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';

const here = path.dirname(fileURLToPath(import.meta.url));

/** Each store's fields (the "## " headings in its metadata.md) and their character limits. */
const STORES = {
  'app-store/metadata.md': {
    Name: 30,
    Subtitle: 30,
    'Promotional text': 170,
    Keywords: 100,
    Description: 4000,
    "What's new": 4000,
  },
  'play/metadata.md': { 'App name': 30, 'Short description': 80, 'Full description': 4000 },
};
/** The listing pages a player reads on the addon stores; checked for voice only. */
const VOICE_ONLY = ['curseforge-description.md', 'data-addon-description.md'];
const BANNED = [
  /[—–]/u,
  /!(?!\[)/u,
  /\b(seamless(ly)?|elevate|unlock|empower|journey|game[- ]changer|supercharge|effortless(ly)?|ultimate|revolution\w*)\b/iu,
];

/** The text under each "## " heading of a metadata file. */
function sections(markdown) {
  const fields = {};
  for (const block of markdown.split(/^## /mu).slice(1)) {
    const [heading, ...body] = block.split('\n');
    fields[heading.trim()] = body.join('\n').trim();
  }
  return fields;
}

/** A field's length as the store counts it: Play strips its HTML tags. */
const storeLength = (text) => text.replace(/<\/?b>/gu, '').length;

function voiceErrors(file, text) {
  return text
    .split('\n')
    .flatMap((line, index) =>
      BANNED.filter((rule) => rule.test(line)).map((rule) => `${file}:${index + 1}: ${rule} in "${line.trim()}"`),
    );
}

const errors = [];
for (const [file, limits] of Object.entries(STORES)) {
  const markdown = readFileSync(path.join(here, file), 'utf8');
  const fields = sections(markdown);
  for (const [field, limit] of Object.entries(limits)) {
    const text = fields[field];
    if (text === undefined) {
      errors.push(`${file}: no "## ${field}" section`);
      continue;
    }
    const length = storeLength(text);
    console.log(`${file} ${field.padEnd(18)} ${String(length).padStart(4)} of ${limit}`);
    if (length > limit) errors.push(`${file}: ${field} is ${length} characters, over ${limit}`);
  }
  errors.push(...voiceErrors(file, markdown));
}
for (const file of VOICE_ONLY) errors.push(...voiceErrors(file, readFileSync(path.join(here, file), 'utf8')));

if (errors.length > 0) {
  console.error(errors.join('\n'));
  process.exit(1);
}
console.log('store text ok');
