// web/src/lib/setup/links.ts
// Where /setup's actions go. CurseForge is addressed by project id (the ids in the two TOCs'
// X-Curse-Project-ID lines; the site redirects an id to the project's page, whatever its slug),
// Wago by the slug its pages answer to. The companion's download is its own component
// (components/setup/CompanionDownload.svelte, lib/setup/release.ts).
import { battlenetStartUrl } from '../account/api';
import type { StepActionId } from './steps';

export const CURSEFORGE_URL = 'https://www.curseforge.com/projects/1706689';
export const WAGO_URL = 'https://addons.wago.io/addons/foreversixty';

/** Sign-in and Refresh come back to /setup, where row 1 shows what changed (setup spec U6, G3). */
export const SETUP_SIGNED_IN = '/setup?signed_in=1';
export const SETUP_REFRESHED = '/setup?refreshed=1';

/** The href of a step's link action; null for a button (retry, refresh, copy link) and the download. */
export function stepHref(id: StepActionId): string | null {
  switch (id) {
    case 'signin':
      return battlenetStartUrl(SETUP_SIGNED_IN);
    case 'paste':
      return '#paste';
    default:
      return null;
  }
}
