// web/src/lib/guild/raid-page-deps.ts
// The raid page controller's live wiring: the site's own reads and writes, the document's
// visibility and the browser's storage. Kept apart so the controller's tests never load them.
import { fetchMeFresh } from '../account/api';
import { API_BASE_URL } from '../planner/config';
import { fetchGuild } from '../rankings/api';
import {
  clearAttendanceOverride,
  clearMatch,
  fetchMatchCandidates,
  postReport,
  setAttendanceOverride,
  setMatch,
} from './attendance-api';
import { fetchGuildHome, fetchGuildRaids } from './api';
import { clearAssignment, postAssignments, setAssignment } from './assignments-api';
import { pingRaid, postRaidNow, resetCompTarget, setCompTarget } from './overview-api';
import { browserStorage } from './reserve-keys';
import type { RaidPageDeps, RaidRoute } from './raid-page-types';
import { fetchGuildRaid, refetchGuildRaid, setRaidStatus, updateRaid } from './raids-api';
import { sendSignup } from './signup-save';

export function liveRaidPageDeps(route: RaidRoute): RaidPageDeps {
  return {
    loadGuild: () => fetchGuild(route.guild),
    loadMe: () => fetchMeFresh(),
    loadHome: (guildId) => fetchGuildHome(guildId),
    fetchRaid: (guildId, raidId) => fetchGuildRaid(guildId, raidId),
    refetchRaid: (guildId, raidId) => refetchGuildRaid(guildId, raidId),
    saveSignup: sendSignup,
    setStatus: (guildId, raidId, action, body) => setRaidStatus(guildId, raidId, action, API_BASE_URL, body),
    updateRaid: (guildId, raidId, body) => updateRaid(guildId, raidId, body),
    loadMatchCandidates: (guildId, raidId) => fetchMatchCandidates(guildId, raidId),
    putMatch: (guildId, raidId, reportId) => setMatch(guildId, raidId, reportId),
    deleteMatch: (guildId, raidId) => clearMatch(guildId, raidId),
    putOverride: (guildId, raidId, key, status) => setAttendanceOverride(guildId, raidId, key, status),
    deleteOverride: (guildId, raidId, key) => clearAttendanceOverride(guildId, raidId, key),
    postReport: (guildId, raidId, repost) => postReport(guildId, raidId, repost),
    putCompTarget: (guildId, raidId, target) => setCompTarget(guildId, raidId, target),
    deleteCompTarget: (guildId, raidId) => resetCompTarget(guildId, raidId),
    pingRaid: (guildId, raidId) => pingRaid(guildId, raidId),
    postRaidNow: (guildId, raidId) => postRaidNow(guildId, raidId),
    putAssignment: (guildId, raidId, encounterId, slot, keys) =>
      setAssignment(guildId, raidId, encounterId, slot, keys),
    deleteAssignment: (guildId, raidId, encounterId, slot) =>
      clearAssignment(guildId, raidId, encounterId, slot),
    postAssignments: (guildId, raidId) => postAssignments(guildId, raidId),
    loadZones: (guildId) => fetchGuildRaids(guildId).then((page) => page.zones ?? []),
    isVisible: () => typeof document === 'undefined' || document.visibilityState === 'visible',
    now: () => new Date(),
    storage: browserStorage(),
  };
}
