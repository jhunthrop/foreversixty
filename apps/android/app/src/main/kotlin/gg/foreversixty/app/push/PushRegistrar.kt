package gg.foreversixty.app.push

import android.util.Log
import gg.foreversixty.app.model.LiveStartTokenRequest
import gg.foreversixty.app.model.PushTokenRequest
import gg.foreversixty.app.network.ApiClient
import gg.foreversixty.app.network.ApiEndpoints
import gg.foreversixty.app.network.ApiResult

/**
 * Registers this install's FCM token with `POST /v1/me/push-tokens` while signed in, and as the raid
 * countdown's start token (`POST /v1/me/live-activities/start-token`), so `live.raid.start` can draw
 * the ongoing raid-night notification before the app has run that day. The token is tied to the
 * device token on the server, so revoking the device at sign-out drops it there (migration 0062);
 * nothing is unregistered by hand.
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
            is ApiResult.Success -> {
                registerStartToken(value)
                true
            }
            is ApiResult.Failure -> {
                Log.w(TAG, "push token not registered: ${result.error}")
                false
            }
        }
    }

    private suspend fun registerStartToken(token: String) {
        val request = LiveStartTokenRequest(platform = PLATFORM, kind = RAID, token = token)
        (api.send(ApiEndpoints.storeLiveStartToken(request)) as? ApiResult.Failure)?.let {
            Log.w(TAG, "raid start token not stored: ${it.error}")
        }
    }

    private companion object {
        const val PLATFORM = "android"
        const val RAID = "raid"
        const val TAG = "Push"
    }
}
