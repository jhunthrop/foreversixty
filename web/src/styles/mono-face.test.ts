// The monospace face is for literal code only (type pass 2026-10-29, polish spec 4.4): export
// strings, slash commands, file names, build strings, pairing codes, ids and every field a player
// copies or pastes. Every UI number, date, count, stat and label is the body face with tabular
// numerals (the `numeric` utility in global.css). This walks every source file that can style
// text and fails on a use of --font-mono or the `font-mono` class outside the list below, so a
// new stat cannot quietly bring the code face back. A new entry names what the code is.
import { readdirSync, readFileSync } from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';

const SRC = path.resolve(path.dirname(fileURLToPath(import.meta.url)), '..');
const STYLED_EXTENSIONS = new Set(['.astro', '.svelte', '.ts', '.css']);

/** Every place the code face may appear, with what it sets. */
const ALLOWED = new Map<string, number>([
  ['styles/nav.css', 2], // the header export paste box and the search box's "/" key hint
  ['components/Character.svelte', 1], // <code> a character address
  ['components/Guild.svelte', 1], // <code> a guild address
  ['components/CharacterSearch.astro', 1], // the "/" key hint
  ['components/CopyForGame.svelte', 1], // the copy-for-the-game string field
  ['components/Connect.svelte', 2], // the pairing code the companion asks for
  ['components/Account.svelte', 2], // the pairing code the companion asks for
  ['components/Upload.svelte', 7], // file and folder names, /combatlog, the chosen file name
  ['components/AddonPasteBox.svelte', 1], // the addon export paste box
  ['components/ui/CopyLinkField.svelte', 1], // a link to copy
  ['components/planner/SharePanel.svelte', 3], // the copy-as-text string, a field id, the share link
  ['components/planner/ImportBox.svelte', 1], // the talent code paste box
  ['components/report/QueriesView.svelte', 2], // the SQL box and its result cells
  ['components/report/MechanicsMode.svelte', 2], // <code> the forever-logs command
  ['components/sim/ColdPasteHero.svelte', 1], // the export paste box
  ['components/sim/SavedSim.svelte', 1], // the engine build string
  ['components/sim/DetailsCard.svelte', 1], // the engine build string
  ['components/sim/RunControl.svelte', 1], // the engine's own error sentence
  ['components/sim/tools/BulkRunBar.svelte', 1], // the engine's own error sentence
  ['components/sim/SavedWeights.svelte', 1], // the weights export string
  ['components/sim/tools/StatWeights.svelte', 1], // the weights export string
  ['components/sim/RequestDrawer.svelte', 3], // the request JSON, a field id, the buff ids
  ['components/guild/RaidReport.svelte', 1], // the copy fallback box
  ['components/guild/RaidOfficerRow.svelte', 1], // the copy fallback box
  ['components/guild/GuildDiscordChannel.svelte', 1], // the channel id
  ['pages/setup.astro', 1], // the slash commands to type
]);

/** `var(--font-mono)`, or the `font-mono` class. */
const MONO_FACE_USE = /var\(--font-mono\)|(?<![\w-])font-mono(?![\w-])/g;

function styledFiles(dir: string): string[] {
  return readdirSync(dir, { withFileTypes: true }).flatMap((entry) => {
    const full = path.join(dir, entry.name);
    if (entry.isDirectory()) return entry.name === 'data' ? [] : styledFiles(full);
    if (entry.name.endsWith('.test.ts')) return [];
    return STYLED_EXTENSIONS.has(path.extname(entry.name)) ? [full] : [];
  });
}

/** Counts of code-face uses per file, keyed relative to src/; comments are not uses. */
function monoFaceUses(): Map<string, number> {
  const uses = new Map<string, number>();
  for (const file of styledFiles(SRC)) {
    const text = readFileSync(file, 'utf8')
      .replace(/\/\*[\s\S]*?\*\//g, '')
      .replace(/^\s*\/\/.*$/gm, '')
      .replace(/<!--[\s\S]*?-->/g, '');
    const count = text.match(MONO_FACE_USE)?.length ?? 0;
    if (count > 0) uses.set(path.relative(SRC, file), count);
  }
  return uses;
}

describe('the code face', () => {
  it('appears only on literal code', () => {
    expect(Object.fromEntries(monoFaceUses())).toEqual(Object.fromEntries(ALLOWED));
  });
});
