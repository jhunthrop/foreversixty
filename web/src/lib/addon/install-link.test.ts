import { describe, expect, it } from 'vitest';
import { INSTALL_ADDON_LINK } from './install-link';

describe('INSTALL_ADDON_LINK', () => {
  it('says Install the addon and goes to /setup', () => {
    expect(INSTALL_ADDON_LINK).toEqual({ label: 'Install the addon', href: '/setup' });
  });
});
