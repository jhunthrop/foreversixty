package gg.foreversixty.app.rankings.logic

import gg.foreversixty.app.model.AccountCharacter
import gg.foreversixty.app.model.CharacterHome
import gg.foreversixty.app.model.FeedPage
import gg.foreversixty.app.model.FeedRow
import gg.foreversixty.app.model.GuildHome
import gg.foreversixty.app.model.HomeBlocks
import gg.foreversixty.app.model.Me
import gg.foreversixty.app.model.RaidLine
import gg.foreversixty.app.network.ApiClient
import gg.foreversixty.app.network.ApiEndpoints
import gg.foreversixty.app.network.ApiResult

/**
 * The reads a character page makes beside its card (design/specs/2026-11-06-character.md C2 to C22):
 * the band's upgrades and the character's feed rows for anyone; for the owner the home's guild ladder
 * and spec week, the guild's next raid, and the account character (its level and exported talents).
 */
interface CharacterExtrasService {
    suspend fun home(region: String, ruleset: String, name: String): ApiResult<CharacterHome>
    suspend fun blocks(characterKey: String): ApiResult<HomeBlocks>
    suspend fun guild(id: Long): ApiResult<GuildHome>
    suspend fun activity(characterKey: String): ApiResult<FeedPage>
    suspend fun me(): ApiResult<Me>
}

class LiveCharacterExtras(private val client: ApiClient) : CharacterExtrasService {
    override suspend fun home(region: String, ruleset: String, name: String) = client.send(ApiEndpoints.getCharacterHome(region, ruleset, name))
    override suspend fun blocks(characterKey: String) = client.send(ApiEndpoints.getHome(characterKey))
    override suspend fun guild(id: Long) = client.send(ApiEndpoints.getGuildHome(id))
    override suspend fun activity(characterKey: String) = client.send(ApiEndpoints.getFeed(character = characterKey, limit = ACTIVITY_ROWS))
    override suspend fun me() = client.send(ApiEndpoints.getMe())

    companion object {
        /** C18: the activity rows the page shows. */
        const val ACTIVITY_ROWS = 5
    }
}

/** What the page shows beside the card; each piece is null (or empty) until it answers, and stays so on failure. */
data class CharacterExtras(
    val home: CharacterHome? = null,
    val blocks: HomeBlocks? = null,
    val nextRaid: RaidLine? = null,
    val activity: List<FeedRow> = emptyList(),
    /** The owner's account character: its level and exported talents decide `spend your points`. */
    val own: AccountCharacter? = null,
) {
    companion object {
        /** Reads the extras for [route]; the owner's pieces only when [owner]. A failed read leaves its piece out. */
        suspend fun load(service: CharacterExtrasService, route: CharacterRoute, characterKey: String, owner: Boolean): CharacterExtras {
            val home = service.home(route.region, route.ruleset, route.name).data()
            val activity = service.activity(characterKey).data()?.rows.orEmpty()
            if (!owner) return CharacterExtras(home = home, activity = activity)
            val blocks = service.blocks(characterKey).data()
            val raid = blocks?.guildSims?.guildId?.let { service.guild(it).data()?.nextRaid }
            val own = service.me().data()?.characters?.firstOrNull { it.key == characterKey }
            return CharacterExtras(home, blocks, raid, activity, own)
        }

        private fun <T> ApiResult<T>.data(): T? = (this as? ApiResult.Success)?.data
    }
}
