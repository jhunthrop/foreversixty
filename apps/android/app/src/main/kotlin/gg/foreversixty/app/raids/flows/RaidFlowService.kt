package gg.foreversixty.app.raids.flows

import gg.foreversixty.app.model.GuildApplicationDecision
import gg.foreversixty.app.model.GuildOfficerApplications
import gg.foreversixty.app.model.ReadyCheck
import gg.foreversixty.app.model.RecruitingGuilds
import gg.foreversixty.app.network.ApiClient
import gg.foreversixty.app.network.ApiEndpoints
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.raids.logic.CharacterKey

/** The Raid tab's calls beyond the raids themselves (spec 2026-11-08 4.2, F2, F3), behind an interface for tests. */
interface RaidFlowService {
    suspend fun applications(guild: Long): ApiResult<GuildOfficerApplications>
    suspend fun accept(guild: Long, application: Long): ApiResult<GuildApplicationDecision>
    suspend fun decline(guild: Long, application: Long): ApiResult<GuildApplicationDecision>
    suspend fun recruiting(spec: String, faction: String, region: String?, ruleset: String?): ApiResult<RecruitingGuilds>

    /** The character's Ready check for [raid]; null for a key that does not parse. */
    suspend fun ready(characterKey: String, raid: Long): ApiResult<ReadyCheck>?
}

class LiveRaidFlowService(private val client: ApiClient) : RaidFlowService {
    override suspend fun applications(guild: Long) = client.send(ApiEndpoints.listGuildApplications(guild))
    override suspend fun accept(guild: Long, application: Long) = client.send(ApiEndpoints.acceptGuildApplication(guild, application))
    override suspend fun decline(guild: Long, application: Long) = client.send(ApiEndpoints.declineGuildApplication(guild, application))
    override suspend fun recruiting(spec: String, faction: String, region: String?, ruleset: String?) =
        client.send(ApiEndpoints.listRecruitingGuilds(spec, faction, region, ruleset))

    override suspend fun ready(characterKey: String, raid: Long): ApiResult<ReadyCheck>? {
        val key = CharacterKey.parse(characterKey) ?: return null
        return client.send(ApiEndpoints.getCharacterReady(key.region, key.ruleset, key.name, raid))
    }
}
