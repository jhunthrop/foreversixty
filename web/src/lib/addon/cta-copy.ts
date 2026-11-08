import {
  sharedCtaGetAddonAndCompanion,
  sharedCtaInstallAddonAndCompanion,
  sharedCtaInstallAddonAndCompanionInline,
} from '../copy/shared/cta.generated';

/** The one wording for "get the addon", so every call to action names the desktop companion too
 *  (it uploads logs and syncs the addon's data; both downloads are on /setup). */
export const ctaCopy = {
  getAddonAndCompanion: sharedCtaGetAddonAndCompanion,
  installAddonAndCompanion: sharedCtaInstallAddonAndCompanion,
  /** Mid-sentence form of installAddonAndCompanion. */
  installAddonAndCompanionInline: sharedCtaInstallAddonAndCompanionInline,
} as const;
