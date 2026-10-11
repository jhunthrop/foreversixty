// web/src/lib/bis/landing-character.ts
// The signed-in half of the /bis landing, loaded only after the page has settled and only when a
// session is hinted: finds the character the header selector points at, reads its spec's BiS
// file to count the worn slots, fills the lead, and tells the page which spec and band are the
// character's. It runs again on every selector change, so the lead and the page follow in place;
// only the newest question may paint.
import { fetchMeOnce } from '../account/api';
import type { MeCharacter } from '../account/api';
import { selectedCharacter } from '../account/hero-character';
import { onCurrentCharacterChange, readCurrent } from '../current-character';
import { bandEntryFor, fetchBisFile } from './hover';
import { DEFAULT_LANDING_BAND, characterBand } from './landing-bands';
import type { LandingController } from './landing-client';
import { specKeyFor, type LandingConfig } from './landing-config';
import { buildLeadView, type LeadView } from './landing-lead';
import { hideLead, renderLead, showLeadShell } from './landing-lead-dom';
import type { BisBand, Faction } from './types';

const DEFAULT_LOOKUP_FACTION: Faction = 'alliance';

async function selectedMeCharacter(): Promise<MeCharacter | null> {
  const me = await fetchMeOnce().catch(() => null);
  return me === null ? null : selectedCharacter(readCurrent(), me);
}

/** The character's band in its spec's file; undefined when the spec is unknown or has no file. */
async function bandEntryOf(
  config: LandingConfig,
  character: MeCharacter,
  specKey: string | undefined,
): Promise<BisBand | undefined> {
  if (specKey === undefined) return undefined;
  const file = await fetchBisFile(config.build, specKey);
  if (file === null) return undefined;
  const band = characterBand(character.level) ?? DEFAULT_LANDING_BAND;
  return bandEntryFor(file, band, character.faction ?? DEFAULT_LOOKUP_FACTION);
}

export function followSelectedCharacter(config: LandingConfig, page: LandingController): void {
  const root = document.querySelector<HTMLElement>('[data-bis-lead]');
  if (root === null) return;
  let latest = 0;
  const answer = async (): Promise<void> => {
    const asked = ++latest;
    const character = await selectedMeCharacter();
    if (asked !== latest) return;
    const specKey = character === null ? undefined : specKeyFor(config, character.class, character.spec);
    const bandEntry = character === null ? undefined : await bandEntryOf(config, character, specKey);
    if (asked !== latest) return;
    const view: LeadView | null =
      character === null ? null : buildLeadView({ character, config, bandEntry, now: new Date() });
    if (character === null || view === null) {
      hideLead(root);
      page.setCharacter(null);
      return;
    }
    showLeadShell(root);
    renderLead(root, view);
    page.setCharacter({ level: character.level, faction: character.faction, specKey });
  };
  void answer();
  onCurrentCharacterChange(() => void answer());
}
