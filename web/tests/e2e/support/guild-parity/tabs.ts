// web/tests/e2e/support/guild-parity/tabs.ts
// The guild parity states for the Roster, Progression and Readiness tabs (GUILD-ALL lane). Each opener routes the
// committed fixtures (design/mocks/guild_tabs_fixtures.py) and waits until the page shows its data.
import type { GuildStateOpener } from '../guild-parity';

export const TABS_OPENERS: Readonly<Record<string, GuildStateOpener>> = {};
