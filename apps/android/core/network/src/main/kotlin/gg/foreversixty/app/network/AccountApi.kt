package gg.foreversixty.app.network

import gg.foreversixty.app.model.Me

/** The account endpoints. Feature lanes add one small class like this per API area. */
class AccountApi(private val client: ApiClient) {
    suspend fun me(): ApiResult<Me> = client.get(ME_PATH, Me.serializer())

    companion object {
        const val ME_PATH = "/v1/me"
    }
}
