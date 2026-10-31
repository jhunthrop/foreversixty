package gg.foreversixty.app.raids.push

import android.Manifest
import android.app.NotificationChannel
import android.app.NotificationManager
import android.content.Context
import android.content.pm.PackageManager
import androidx.core.app.NotificationCompat
import androidx.core.app.NotificationManagerCompat
import androidx.core.content.ContextCompat
import gg.foreversixty.app.BuildConfig
import gg.foreversixty.app.R
import gg.foreversixty.app.model.ApiJson
import gg.foreversixty.app.network.AssetFixtureSource
import kotlinx.serialization.json.jsonObject
import kotlinx.serialization.json.jsonPrimitive

/**
 * Debug builds only: posts push.raid-posted.fcm with its In, Late and Out actions, as lane F7's
 * builder will, so the action path can be driven on the emulator for the review captures. Started
 * by `adb shell am start ... --ez fs.debug.raidPush true`.
 */
object RaidDebugPush {
    const val EXTRA = "fs.debug.raidPush"
    private const val TAG = "raid.posted-31"
    private const val ID = 31

    fun post(context: Context) {
        if (!BuildConfig.DEBUG) return
        val raw = AssetFixtureSource(context.assets).read("push.raid-posted.fcm.json") ?: return
        val message = ApiJson.parseToJsonElement(raw).jsonObject.getValue("message").jsonObject
        val data = message.getValue("data").jsonObject.mapValues { it.value.jsonPrimitive.content } + ("raid_id" to "31")
        val note = message.getValue("notification").jsonObject
        val manager = context.getSystemService(NotificationManager::class.java)
        manager.createNotificationChannel(
            NotificationChannel(RaidNotificationActions.CHANNEL_RAIDS, context.getString(R.string.notif_section_raids), NotificationManager.IMPORTANCE_HIGH),
        )
        if (ContextCompat.checkSelfPermission(context, Manifest.permission.POST_NOTIFICATIONS) != PackageManager.PERMISSION_GRANTED) return
        val builder = NotificationCompat.Builder(context, RaidNotificationActions.CHANNEL_RAIDS)
            .setSmallIcon(R.mipmap.ic_launcher)
            .setContentTitle(note.getValue("title").jsonPrimitive.content)
            .setContentText(note.getValue("body").jsonPrimitive.content)
        RaidNotificationActions.actions(context, data["category"] ?: RaidPushCategory.SIGNUP, data, TAG, ID).forEach(builder::addAction)
        NotificationManagerCompat.from(context).notify(TAG, ID, builder.build())
    }
}
