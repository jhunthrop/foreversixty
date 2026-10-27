// The row's strings are read out of the experience spec, so a drift in either file fails here.
import { readFileSync } from 'node:fs';
import { describe, expect, it } from 'vitest';
import { specString } from '../../test-support/spec-table';
import { discordLinkCopy } from './discord-link-copy';

const SPEC = new URL('../../../../design/specs/2026-10-27-discord-bot.md', import.meta.url);
const HEADING = '### 6.1 On the site';
const cell = (label: string, index = 0, fill: Record<string, string> = {}): string =>
  specString(SPEC, HEADING, label, index, fill);

describe('discordLinkCopy against spec 6.1', () => {
  it.each([
    ['lead', discordLinkCopy.lead, cell('Account lead')],
    [
      'leadSuggested',
      discordLinkCopy.leadSuggested('Olympus XXVII'),
      cell('Account lead', 1, { Guild: 'Olympus XXVII' }),
    ],
    ['link', discordLinkCopy.link, cell('Account controls')],
    ['opening', discordLinkCopy.opening, cell('Account controls', 1)],
    ['linkedLead', discordLinkCopy.linkedLead, cell('Account controls', 2)],
    ['unlink', discordLinkCopy.unlink, cell('Account controls', 4)],
    ['unlinking', discordLinkCopy.unlinking, cell('Account controls', 5)],
    ['unlinkQuestion', discordLinkCopy.unlinkQuestion, cell('Account controls', 6)],
    ['cancel', discordLinkCopy.cancel, cell('Account controls', 8)],
    ['unlinked', discordLinkCopy.unlinked, cell('Account controls', 9)],
    ['unlinkFailed', discordLinkCopy.unlinkFailed, cell('Account controls', 10)],
    ['returnLinked', discordLinkCopy.returnLinked, cell('Account returns')],
    ['returnDenied', discordLinkCopy.returnDenied, cell('Account returns', 1)],
    ['returnFailed', discordLinkCopy.returnFailed, cell('Account returns', 2)],
    ['returnConflict', discordLinkCopy.returnConflict, cell('Account returns', 3)],
    ['conflictHelp', discordLinkCopy.conflictHelp, cell('Account returns', 4)],
    ['stuck', discordLinkCopy.stuck, cell('Account returns', 5)],
  ])('%s', (_name, actual, spec) => {
    expect(actual).toBe(spec);
  });
});

describe('the changelog entry', () => {
  it('is the spec string, verbatim', () => {
    const entry = readFileSync(
      new URL('../../content/changelog/2026-10-27-discord-bot.md', import.meta.url),
      'utf8',
    );
    expect(entry.trim().endsWith(cell('Changelog'))).toBe(true);
  });
});
