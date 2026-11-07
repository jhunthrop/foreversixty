// web/tests/e2e/support/guild-parity/raids.ts
// The guild parity states for the Raids tab (list, plan form, log match) and the raid page (planner, assignments) (GUILD-ALL lane). Each opener routes the
// committed fixtures (design/mocks/guild_raids_fixtures.py) and waits until the page shows its data.
import type { GuildStateOpener } from '../guild-parity';

export const RAIDS_OPENERS: Readonly<Record<string, GuildStateOpener>> = {};
