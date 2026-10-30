package gg.foreversixty.app.network

import gg.foreversixty.app.model.CharacterProfile

/** The public character routes. */
class CharacterApi(private val client: ApiClient) {
    /** GET /v1/characters/{region}/{ruleset}/{name}; [key] is `region/ruleset/slug`. */
    suspend fun profile(key: String): ApiResult<CharacterProfile> =
        client.get("$CHARACTERS_PATH/$key", CharacterProfile.serializer())

    companion object {
        const val CHARACTERS_PATH = "/v1/characters"
    }
}
