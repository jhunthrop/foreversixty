import { SETUP_NAV_ITEM } from '../nav';
import { addonCopy } from './copy';

/** The one "Install the addon" link the gear, talent and guide pages carry: /setup's addon step
 *  wording, to /setup. */
export const INSTALL_ADDON_LINK = {
  label: addonCopy.setupInstallAddon,
  href: SETUP_NAV_ITEM.href,
} as const;
