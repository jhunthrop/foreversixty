/** The one "Install the addon" link the gear, talent and guide pages carry, to /setup. It is also the
 *  label of /setup's own addon step (addonCopy.setupInstallAddon). Dependency free: the BiS and planner
 *  islands import it without the rest of the addon copy. */
export const INSTALL_ADDON_LINK = {
  label: 'Install the addon',
  href: '/setup',
} as const;
