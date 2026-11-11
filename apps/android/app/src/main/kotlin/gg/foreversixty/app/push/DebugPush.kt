package gg.foreversixty.app.push

import android.content.Context
import gg.foreversixty.app.BuildConfig
import gg.foreversixty.app.model.ApiJson
import gg.foreversixty.app.network.AssetFixtureSource
import kotlinx.serialization.json.JsonObject
import kotlinx.serialization.json.jsonObject
import kotlinx.serialization.json.jsonPrimitive

/**
 * Debug builds only: draws a push fixture exactly as an FCM message is drawn (PushMessage, then
 * PushNotifier, with the actions its kind carries), so the review walk can show any push on the
 * emulator: raid.posted with In, Late and Bench, raid.short with its Mention, guild.applicant with
 * Accept and Decline. Started by `adb shell am start ... --es fs.debug.push push.raid-short.fcm.json`.
 */
object DebugPush {
    const val EXTRA = "fs.debug.push"

    /** Only an FCM push fixture's file name: `push.<kind>.fcm.json`, nothing with a path in it. */
    private val FIXTURE = Regex("""push\.[a-z0-9-]+\.fcm\.json""")

    fun post(context: Context, fixture: String) {
        if (!BuildConfig.DEBUG) return
        require(FIXTURE.matches(fixture)) { "fs.debug.push takes a push.<kind>.fcm.json fixture name, not $fixture" }
        val raw = requireNotNull(AssetFixtureSource(context.assets).read(fixture)) { "no push fixture $fixture" }
        val parsed = requireNotNull(parse(ApiJson.parseToJsonElement(raw).jsonObject)) { "$fixture is not a push the app draws" }
        PushNotifier.show(context, parsed)
    }

    /** The FCM `message` of a fixture as the messaging service reads it; null when it is not a drawn push. */
    internal fun parse(fixture: JsonObject): PushMessage? {
        val message = fixture.getValue("message").jsonObject
        val data = message.getValue("data").jsonObject.mapValues { it.value.jsonPrimitive.content }
        return PushMessage.parse(data, notificationTitle = null, notificationBody = null, collapse = null)
    }
}
