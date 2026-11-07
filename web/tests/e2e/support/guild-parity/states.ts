// web/tests/e2e/support/guild-parity/states.ts
// The guild parity states for the Overview states: before raids open, no raid logged, not in a guild, low-level character (GUILD-ALL lane). Each opener routes the
// committed fixtures (design/mocks/guild_states_fixtures.py) and waits until the page shows its data.
import type { GuildStateOpener } from '../guild-parity';

export const STATES_OPENERS: Readonly<Record<string, GuildStateOpener>> = {};
