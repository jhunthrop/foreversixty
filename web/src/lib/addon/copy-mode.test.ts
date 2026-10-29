// web/src/lib/addon/copy-mode.test.ts
// L8.5 ruling 5: the addon's copy mode has no "Copy my character" button any more (L3.4 fix 2):
// `/fs copy` opens the dialog on the Copy my character tab with the box selected, and Ctrl+C
// is the action. `/fs export` is gone too (the addon answers it with its command list). Every
// place the site tells a player how to get the export says so.
import { readdirSync, readFileSync, statSync } from 'node:fs';
import path from 'node:path';
import { describe, expect, it } from 'vitest';
import { scoreCopy } from '../score/copy';
import { landingCopy } from '../sim/landing-copy';
import { addonCopy } from './copy';

const WEB_ROOT = path.resolve(import.meta.dirname, '..', '..', '..');

/** Every source file under `dir` (tests left out), recursively. */
function sources(dir: string): string[] {
  return readdirSync(dir).flatMap((name) => {
    const full = path.join(dir, name);
    if (statSync(full).isDirectory()) return sources(full);
    return /\.(ts|svelte|astro|json|md|mdx|html)$/.test(name) && !/\.test\.ts$/.test(name) ? [full] : [];
  });
}

describe('the site’s instruction for the addon’s copy mode (L8.5 ruling 5)', () => {
  it('says the page updates by itself, and /fs copy only without the companion, on the profile’s In game section', () => {
    expect(`${scoreCopy.inGamePlannerLead} ${scoreCopy.inGamePasteLink}.`).toBe(
      'This page updates by itself each time you log out with the addon and the companion running. Without the companion: /fs copy in game, then paste it on the setup page.',
    );
  });

  it('says the same on /setup’s "Out of the game" card and the sim’s paste box', () => {
    expect(addonCopy.flowOutBody).toBe(
      'In game, type /fs copy, press Ctrl+C, then paste it into the planner’s Import from addon box.',
    );
    // The lead once ran on into a bare sign-in link ("…here, or"); the polish lane's Task 8
    // removed that link, so the sentence closes.
    expect(landingCopy.pasteHeroHintLead).toBe('In game, type /fs copy, press Ctrl+C, then paste it here.');
  });

  it('leaves no "Copy my character" button or /fs export command anywhere in the site’s copy', () => {
    const swept = [
      ...sources(path.join(WEB_ROOT, 'src')),
      ...sources(path.join(WEB_ROOT, '..', 'docs/contracts/fixtures')),
    ];
    // Fix round 1 item 6: the changelog and guide pages (.md, .mdx) and any .html are swept too.
    expect(swept.some((file) => file.endsWith('.md'))).toBe(true);
    const offenders = swept
      .filter((file) => /press Copy my character|\/fs export/.test(readFileSync(file, 'utf8')))
      .map((file) => path.relative(WEB_ROOT, file));
    expect(offenders).toEqual([]);
  });
});
