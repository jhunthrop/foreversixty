package gg.foreversixty.app.widgets.live

import android.Manifest
import android.app.Notification
import android.app.NotificationChannel
import android.app.NotificationManager
import android.app.PendingIntent
import android.content.Context
import android.content.Intent
import android.content.pm.PackageManager
import android.os.Build
import androidx.core.app.NotificationCompat
import androidx.core.app.NotificationManagerCompat
import androidx.core.content.ContextCompat
import androidx.core.net.toUri
import gg.foreversixty.app.MainActivity
import gg.foreversixty.app.R
import gg.foreversixty.app.navigation.DeepLink
import gg.foreversixty.app.raids.push.RaidNotificationActions
import gg.foreversixty.app.raids.push.RaidPushCategory
import gg.foreversixty.app.widgets.logic.RaidLiveContent
import gg.foreversixty.app.widgets.logic.SimLive
import gg.foreversixty.app.widgets.logic.SimLiveContent

/**
 * The ongoing notifications (spec §5.6): `Raid starting` with a chronometer counting down to the
 * invites and Late and Out, and `Server sims` with a determinate bar. Quiet hours never suppress
 * them (no sound, no buzz). The builders are pure functions of their content so tests capture them.
 */
class LiveNotifier(private val context: Context) {
    private val manager = NotificationManagerCompat.from(context)

    fun showRaid(content: RaidLiveContent) {
        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.TIRAMISU &&
            ContextCompat.checkSelfPermission(context, Manifest.permission.POST_NOTIFICATIONS) != PackageManager.PERMISSION_GRANTED
        ) {
            return
        }
        ensureChannels()
        manager.notify(RAID_TAG, RAID_ID, raidNotification(context, content))
    }

    fun cancelRaid() = manager.cancel(RAID_TAG, RAID_ID)

    fun showSim(sim: SimLive, content: SimLiveContent) {
        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.TIRAMISU &&
            ContextCompat.checkSelfPermission(context, Manifest.permission.POST_NOTIFICATIONS) != PackageManager.PERMISSION_GRANTED
        ) {
            return
        }
        ensureChannels()
        manager.notify(simTag(sim.simId), SIM_ID, simNotification(context, sim, content))
    }

    private fun ensureChannels() {
        val system = context.getSystemService(NotificationManager::class.java)
        // Low importance: an ongoing countdown neither sounds nor buzzes (Reduce Motion, quiet hours).
        system.createNotificationChannel(NotificationChannel(CHANNEL_RAID, context.getString(R.string.live_channel_raid), NotificationManager.IMPORTANCE_LOW))
        system.createNotificationChannel(NotificationChannel(CHANNEL_SIMS, context.getString(R.string.live_channel_sims), NotificationManager.IMPORTANCE_LOW))
    }

    companion object {
        const val CHANNEL_RAID = "live_raid"
        const val CHANNEL_SIMS = "live_sims"
        const val RAID_TAG = "live-raid"
        private const val RAID_ID = 1401
        private const val SIM_ID = 1402

        fun simTag(simId: String) = "live-sim-$simId"

        /** The raid countdown: Late and Out answer in place through the raids lane's receiver. */
        fun raidNotification(context: Context, content: RaidLiveContent): Notification {
            val data = mapOf("guild_id" to content.guildId.toString(), "raid_id" to content.raidId.toString())
            val builder = NotificationCompat.Builder(context, CHANNEL_RAID)
                .setSmallIcon(R.mipmap.ic_launcher)
                .setContentTitle(content.title)
                .setContentText(content.text)
                .setSubText(content.subText)
                .setStyle(NotificationCompat.BigTextStyle().bigText(content.bigText))
                .setWhen(content.countdownTo.toEpochMilli())
                .setShowWhen(true)
                .setUsesChronometer(true)
                .setChronometerCountDown(true)
                .setOngoing(true)
                .setOnlyAlertOnce(true)
                .setCategory(NotificationCompat.CATEGORY_EVENT)
                .setTicker(content.description)
                .setContentIntent(open(context, DeepLink.Raid(content.guildId, content.raidId).url()))
            RaidNotificationActions.actions(context, RaidPushCategory.LIVE, data, RAID_TAG, RAID_ID).forEach(builder::addAction)
            return builder.build()
        }

        fun simNotification(context: Context, sim: SimLive, content: SimLiveContent): Notification {
            val builder = NotificationCompat.Builder(context, CHANNEL_SIMS)
                .setSmallIcon(R.mipmap.ic_launcher)
                .setContentTitle(content.title)
                .setContentText(content.text)
                .setSubText(content.subText)
                .setShowWhen(false)
                .setOngoing(content.ongoing)
                .setAutoCancel(!content.ongoing)
                .setOnlyAlertOnce(true)
                .setCategory(NotificationCompat.CATEGORY_PROGRESS)
                .setContentIntent(open(context, DeepLink.Sim(sim.simId).url()))
            content.progress?.let { (done, max) -> builder.setProgress(max, done, content.indeterminate) }
            return builder.build()
        }

        private fun open(context: Context, url: String): PendingIntent {
            val intent = Intent(Intent.ACTION_VIEW, url.toUri(), context, MainActivity::class.java)
                .addFlags(Intent.FLAG_ACTIVITY_NEW_TASK or Intent.FLAG_ACTIVITY_CLEAR_TOP)
            return PendingIntent.getActivity(context, url.hashCode(), intent, PendingIntent.FLAG_UPDATE_CURRENT or PendingIntent.FLAG_IMMUTABLE)
        }
    }
}
