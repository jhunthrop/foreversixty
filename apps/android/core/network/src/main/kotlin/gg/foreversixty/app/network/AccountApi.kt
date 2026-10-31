package gg.foreversixty.app.network

import gg.foreversixty.app.model.Me

/** The account calls the shell makes. Every route is an `ApiEndpoints` member; this names the ones used here. */
class AccountApi(private val client: ApiClient) {
    suspend fun me(): ApiResult<Me> = client.send(ApiEndpoints.getMe())
}
