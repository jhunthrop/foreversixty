// web/src/lib/guild/raid-page-deps.ts
// The raid page controller's live wiring: the site's own reads and writes, the document's
// visibility and the browser's storage. Kept apart so the controller's tests never load them.
import { fetchMeFresh } from '../account/api';
import { API_BASE_URL } from '../planner/config';
import { fetchGuild } from '../rankings/api';
import { fetchGuildHome, fetchGuildRaids } from './api';
import { browserStorage } from './planned-reserve';
import type { RaidPageDeps, RaidRoute } from './raid-page-controller.svelte';
import { fetchGuildRaid, putSignup, refetchGuildRaid, setRaidStatus, updateRaid } from './raids-api';

export function liveRaidPageDeps(route: RaidRoute): RaidPageDeps {
  return {
    loadGuild: () => fetchGuild(route.guild),
    loadMe: () => fetchMeFresh(),
    loadHome: (guildId) => fetchGuildHome(guildId),
    fetchRaid: (guildId, raidId) => fetchGuildRaid(guildId, raidId),
    refetchRaid: (guildId, raidId) => refetchGuildRaid(guildId, raidId),
    putSignup: (guildId, raidId, key, body) => putSignup(guildId, raidId, key, body),
    setStatus: (guildId, raidId, action, body) => setRaidStatus(guildId, raidId, action, API_BASE_URL, body),
    updateRaid: (guildId, raidId, body) => updateRaid(guildId, raidId, body),
    loadZones: (guildId) => fetchGuildRaids(guildId).then((page) => page.zones ?? []),
    isVisible: () => typeof document === 'undefined' || document.visibilityState === 'visible',
    now: () => new Date(),
    storage: browserStorage(),
  };
}
