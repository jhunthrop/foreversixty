// web/tests/e2e/support/guild-parity/directory.ts
// The guild parity states for the Guild landing (/guild), join, claim, and the home page Top guilds (GUILD-ALL lane). Each opener routes the
// committed fixtures (design/mocks/guild_directory_fixtures.py) and waits until the page shows its data.
import type { GuildStateOpener } from '../guild-parity';

export const DIRECTORY_OPENERS: Readonly<Record<string, GuildStateOpener>> = {};
