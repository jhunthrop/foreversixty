package gg.foreversixty.app.push

import android.util.Log
import gg.foreversixty.app.model.PushTokenRequest
import gg.foreversixty.app.network.ApiClient
import gg.foreversixty.app.network.ApiEndpoints
import gg.foreversixty.app.network.ApiResult

/**
 * Registers this install's FCM token with `POST /v1/me/push-tokens` while signed in. The token is
 * tied to the device token on the server, so revoking the device at sign-out drops it there
 * (migration 0062); nothing is unregistered by hand.
 */
class PushRegistrar(
    private val api: ApiClient,
    private val appVersion: String,
    private val fcmToken: suspend () -> String?,
    private val signedIn: suspend () -> Boolean,
    private val enabled: () -> Boolean,
) {
    /** Registers [token] (a fresh one from FCM), else the current one. False when nothing was sent. */
    suspend fun register(token: String? = null): Boolean {
        if (!enabled() || !signedIn()) return false
        val value = token ?: fcmToken() ?: return false
        val request = PushTokenRequest(platform = PLATFORM, token = value, appVersion = appVersion)
        return when (val result = api.send(ApiEndpoints.registerPushToken(request))) {
            is ApiResult.Success -> true
            is ApiResult.Failure -> {
                Log.w(TAG, "push token not registered: ${result.error}")
                false
            }
        }
    }

    private companion object {
        const val PLATFORM = "android"
        const val TAG = "Push"
    }
}
