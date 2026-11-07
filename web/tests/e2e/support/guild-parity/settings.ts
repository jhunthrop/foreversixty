// web/tests/e2e/support/guild-parity/settings.ts
// The guild parity states for the Settings tab and the guild settings page (billing, Discord, crest, invite link) (GUILD-ALL lane). Each opener routes the
// committed fixtures (design/mocks/guild_settings_fixtures.py) and waits until the page shows its data.
import type { GuildStateOpener } from '../guild-parity';

export const SETTINGS_OPENERS: Readonly<Record<string, GuildStateOpener>> = {};
