// web/src/styles/hero-font.check.test.ts
// Guards src/styles/hero-font.css, the committed Cinzel 700 subset scripts/hero-font.py
// generates for the homepage hero (see that script's module docstring for the full why).
// If home-landing-copy.ts's `homeHeroCopy.headline` (or any future `*headline*`-named
// field scripts/hero-font.py also scans) ever needs a character the committed subset does
// not have, that glyph would silently fall back to the full Cinzel face for anyone who hits
// this page before Cinzel finishes downloading -- not a visual defect (same face), but it
// would quietly reopen the LCP gap this subset exists to close. This test fails loudly
// instead, with the fix (`python3 scripts/hero-font.py`) named right there in the script's
// own error message.
import { execFileSync } from 'node:child_process';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';

const webRoot = path.resolve(fileURLToPath(new URL('.', import.meta.url)), '..', '..');
const scriptPath = path.join(webRoot, 'scripts', 'hero-font.py');

describe('hero-font subset', () => {
  it('has every glyph the hero headline currently needs', () => {
    expect(() =>
      execFileSync('python3', [scriptPath, '--check'], { cwd: webRoot, stdio: 'pipe' }),
    ).not.toThrow();
  });
});
