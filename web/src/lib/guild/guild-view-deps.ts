// web/src/lib/guild/guild-view-deps.ts
// The live reads and writes of the views under a guild page (recruit, applications, a member
// profile). Kept apart from the page states so their tests never load the network modules.
import { fetchMeFresh } from '../account/api';
import { fetchGuild } from '../rankings/api';
import type { ApplicationsPageDeps } from './applications-page.svelte';
import type { ProfilePageDeps } from './profile-page.svelte';
import type { RecruitPageDeps } from './recruit-page.svelte';
import {
  addMemberNote,
  applyToGuild,
  decideApplication,
  deleteMemberNote,
  fetchMemberProfile,
  fetchMyApplications,
  fetchOfficerApplications,
  fetchRecruitCard,
  patchRecruit,
  voteApplication,
  withdrawApplication,
} from './recruit-api';

const viewDeps = {
  loadGuild: fetchGuild,
  loadMe: () => fetchMeFresh(),
};

export function liveRecruitDeps(): RecruitPageDeps {
  return {
    ...viewDeps,
    loadCard: (guildId) => fetchRecruitCard(guildId),
    loadMine: (guildId) => fetchMyApplications(guildId),
    apply: (guildId, input) => applyToGuild(guildId, input),
    withdraw: (guildId, appId) => withdrawApplication(guildId, appId),
  };
}

export function liveApplicationsDeps(): ApplicationsPageDeps {
  return {
    ...viewDeps,
    loadCard: (guildId) => fetchRecruitCard(guildId),
    loadList: (guildId) => fetchOfficerApplications(guildId),
    vote: (guildId, appId, vote) => voteApplication(guildId, appId, vote),
    decide: (guildId, appId, decision) => decideApplication(guildId, appId, decision),
    saveCard: (guildId, patch) => patchRecruit(guildId, patch),
  };
}

export function liveProfileDeps(): ProfilePageDeps {
  return {
    ...viewDeps,
    loadProfile: (guildId, account) => fetchMemberProfile(guildId, account),
    addNote: (guildId, account, body) => addMemberNote(guildId, account, body),
    deleteNote: (guildId, account, noteId) => deleteMemberNote(guildId, account, noteId),
  };
}
