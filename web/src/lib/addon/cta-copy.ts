import { sharedText } from '../copy/shared';

/** The one wording for "get the addon", so every call to action names the desktop companion too
 *  (it uploads logs and syncs the addon's data; both downloads are on /setup). */
export const ctaCopy = {
  getAddonAndCompanion: sharedText('shared.cta.getAddonAndCompanion'),
  installAddonAndCompanion: sharedText('shared.cta.installAddonAndCompanion'),
  /** Mid-sentence form of installAddonAndCompanion. */
  installAddonAndCompanionInline: sharedText('shared.cta.installAddonAndCompanionInline'),
} as const;
