package gg.foreversixty.app.raids.push

import android.content.Context
import gg.foreversixty.app.BuildConfig
import gg.foreversixty.app.model.ApiJson
import gg.foreversixty.app.network.AssetFixtureSource
import gg.foreversixty.app.push.PushMessage
import gg.foreversixty.app.push.PushNotifier
import kotlinx.serialization.json.jsonObject
import kotlinx.serialization.json.jsonPrimitive

/**
 * Debug builds only: draws push.raid-posted.fcm exactly as an FCM message is drawn (PushMessage, then
 * PushNotifier, with In, Late, Bench and Out), so the action path can be driven on the emulator for the
 * review captures. Started by `adb shell am start ... --ez fs.debug.raidPush true`.
 */
object RaidDebugPush {
    const val EXTRA = "fs.debug.raidPush"

    fun post(context: Context) {
        if (!BuildConfig.DEBUG) return
        val raw = AssetFixtureSource(context.assets).read("push.raid-posted.fcm.json") ?: return
        val message = ApiJson.parseToJsonElement(raw).jsonObject.getValue("message").jsonObject
        val data = message.getValue("data").jsonObject.mapValues { it.value.jsonPrimitive.content }
        PushMessage.parse(data, notificationTitle = null, notificationBody = null, collapse = null)?.let { PushNotifier.show(context, it) }
    }
}
