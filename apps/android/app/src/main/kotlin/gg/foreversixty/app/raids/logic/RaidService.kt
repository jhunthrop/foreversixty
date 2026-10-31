package gg.foreversixty.app.raids.logic

import android.content.res.Resources
import gg.foreversixty.app.R
import gg.foreversixty.app.model.GuildRaids
import gg.foreversixty.app.model.Me
import gg.foreversixty.app.model.RaidNightSummary
import gg.foreversixty.app.model.RaidView
import gg.foreversixty.app.model.SignupAnswer
import gg.foreversixty.app.model.SignupRequest
import gg.foreversixty.app.network.ApiClient
import gg.foreversixty.app.network.ApiEndpoints
import gg.foreversixty.app.network.ApiError
import gg.foreversixty.app.network.ApiResult

/** The raid calls the screens make. Live goes through [ApiClient]; tests pass a fake. */
interface RaidService {
    suspend fun me(): ApiResult<Me>
    suspend fun raids(guild: Long): ApiResult<GuildRaids>
    suspend fun raid(guild: Long, raid: Long): ApiResult<RaidView>
    suspend fun signup(guild: Long, raid: Long, characterKey: String, request: SignupRequest): ApiResult<SignupAnswer>
    suspend fun lock(guild: Long, raid: Long): ApiResult<RaidView>
    suspend fun unlock(guild: Long, raid: Long): ApiResult<RaidView>
    suspend fun night(raid: Long): ApiResult<RaidNightSummary>
}

class LiveRaidService(private val client: ApiClient) : RaidService {
    override suspend fun me() = client.send(ApiEndpoints.getMe())
    override suspend fun raids(guild: Long) = client.send(ApiEndpoints.listGuildRaids(guild))
    override suspend fun raid(guild: Long, raid: Long) = client.send(ApiEndpoints.getGuildRaid(guild, raid))

    override suspend fun signup(guild: Long, raid: Long, characterKey: String, request: SignupRequest): ApiResult<SignupAnswer> {
        val key = CharacterKey.parse(characterKey)
            ?: return ApiResult.Failure(ApiError.Decoding(0, null, IllegalArgumentException("not a character key: $characterKey")))
        return client.send(ApiEndpoints.putRaidSignup(guild, raid, key.region, key.ruleset, key.name, request))
    }

    override suspend fun lock(guild: Long, raid: Long) = client.send(ApiEndpoints.lockGuildRaid(guild, raid))
    override suspend fun unlock(guild: Long, raid: Long) = client.send(ApiEndpoints.unlockGuildRaid(guild, raid))
    override suspend fun night(raid: Long) = client.send(ApiEndpoints.getRaidNight(raid))
}

/** A character key, `region/ruleset/name-slug`, split for the sign-up route. */
data class CharacterKey(val region: String, val ruleset: String, val name: String) {
    companion object {
        fun parse(key: String): CharacterKey? {
            val parts = key.split("/")
            return if (parts.size == 3 && parts.all { it.isNotEmpty() }) CharacterKey(parts[0], parts[1], parts[2]) else null
        }
    }
}

/** The sentence a refused or failed write shows (spec 5.4 raid states and section 13 `raid.*`). */
object SignupRefusal {
    fun sentence(error: ApiError, res: Resources): String = when (error) {
        is ApiError.Network -> res.getString(R.string.raid_offline)
        is ApiError.Server -> when (error.code) {
            "raid_locked" -> res.getString(R.string.raid_race_locked)
            "raid_cancelled" -> res.getString(R.string.raid_race_cancelled)
            "guild_plan_required" -> res.getString(R.string.raid_race_toolsEnded)
            "rate_limited" -> res.getString(R.string.raid_rateLimited)
            "forbidden" -> error.message
            else -> res.getString(R.string.common_failed)
        }
        else -> res.getString(R.string.common_failed)
    }

    fun isOffline(error: ApiError): Boolean = error is ApiError.Network
}
