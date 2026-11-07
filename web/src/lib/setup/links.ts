// web/src/lib/setup/links.ts
// Where /setup's actions go. CurseForge is addressed by project id (the ids in the two TOCs'
// X-Curse-Project-ID lines; the site redirects an id to the project's page, whatever its slug),
// Wago by the slug its pages answer to. The zip, the data pack and the companion go through the
// site's own /dl/ paths (lib/downloads/redirect.ts).
import { battlenetStartUrl } from '../account/api';
import { companionDownloadPath } from '../downloads/paths';
import type { StepActionId } from './steps';

export const CURSEFORGE_URL = 'https://www.curseforge.com/projects/1706689';
export const WAGO_URL = 'https://addons.wago.io/addons/foreversixty';

/** Sign-in and Refresh come back to /setup, where row 1 shows what changed (setup spec U6, G3). */
export const SETUP_SIGNED_IN = '/setup?signed_in=1';
export const SETUP_REFRESHED = '/setup?refreshed=1';

/** The href of a step's link action; null for a button (retry, refresh). */
export function stepHref(id: StepActionId, companionAsset: string): string | null {
  switch (id) {
    case 'signin':
      return battlenetStartUrl(SETUP_SIGNED_IN);
    case 'curseforge':
      return CURSEFORGE_URL;
    case 'download':
      return companionDownloadPath(companionAsset);
    case 'paste':
      return '#paste';
    default:
      return null;
  }
}
