package gg.foreversixty.app.widgets.live

import android.Manifest
import android.app.Notification
import android.app.NotificationChannel
import android.app.NotificationManager
import android.app.PendingIntent
import android.content.Context
import android.content.Intent
import android.content.pm.PackageManager
import android.graphics.Bitmap
import android.graphics.Color
import android.view.View
import android.os.Build
import android.os.SystemClock
import android.widget.RemoteViews
import java.time.Duration
import java.time.Instant
import androidx.core.app.NotificationCompat
import androidx.core.app.NotificationManagerCompat
import androidx.core.content.ContextCompat
import androidx.core.net.toUri
import gg.foreversixty.app.MainActivity
import gg.foreversixty.app.R
import androidx.compose.ui.graphics.toArgb
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.navigation.DeepLink
import gg.foreversixty.app.push.PushNotifier
import gg.foreversixty.app.raids.push.RaidNotificationActions
import gg.foreversixty.app.raids.push.RaidPushCategory
import gg.foreversixty.app.raids.logic.SignupState
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

    fun showRaid(content: RaidLiveContent, zoneIcon: Bitmap? = null) {
        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.TIRAMISU &&
            ContextCompat.checkSelfPermission(context, Manifest.permission.POST_NOTIFICATIONS) != PackageManager.PERMISSION_GRANTED
        ) {
            return
        }
        ensureChannels()
        manager.notify(RAID_TAG, RAID_ID, raidNotification(context, content, zoneIcon))
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
        private const val CANCELLED_FOR_MILLIS = 5L * 60 * 1000

        fun simTag(simId: String) = "live-sim-$simId"

        /**
         * Raid night (spec F2, mock m-android-lock), one body (layout notification_raid_live): the zone with
         * its icon, `Olympus XXVII · You are In` and the countdown big beside them, red inside 10 minutes;
         * expanded, the time and who is coming and In, Late, Bench and Out answering in place through the
         * raids lane's receiver. Cancelled: the zone and `Cancelled`, gone after 5 minutes.
         */
        fun raidNotification(
            context: Context,
            content: RaidLiveContent,
            zoneIcon: Bitmap? = null,
            now: Instant = Instant.ofEpochMilli(System.currentTimeMillis()),
        ): Notification {
            val builder = PushNotifier.builder(context, CHANNEL_RAID)
                .setContentTitle(content.title)
                .setContentText(content.text)
                .setShowWhen(false)
                .setOnlyAlertOnce(true)
                .setCategory(NotificationCompat.CATEGORY_EVENT)
                .setTicker(content.description)
                .setContentIntent(open(context, DeepLink.Raid(content.guildId, content.raidId).url()))
            val target = content.countdownTo
            if (content.cancelled || target == null) {
                // Nothing to expand: the system's own card, the zone's icon at its end.
                return builder.setLargeIcon(zoneIcon).setOngoing(false).setAutoCancel(true).setTimeoutAfter(CANCELLED_FOR_MILLIS).build()
            }
            val left = Duration.between(now, target).toMillis()
            val data = mapOf("guild_id" to content.guildId.toString(), "raid_id" to content.raidId.toString())
            // Android draws three action buttons at most; the four answers are pills in the expanded body.
            val answers = RaidNotificationActions.actions(context, RaidPushCategory.LIVE, data, RAID_TAG, RAID_ID)
            return builder
                .setStyle(NotificationCompat.DecoratedCustomViewStyle())
                .setOngoing(true)
                .setCustomContentView(liveBody(context, content, zoneIcon, left, answers = emptyList()))
                .setCustomBigContentView(liveBody(context, content, zoneIcon, left, answers))
                .build()
        }

        /** The raid-night body: the top row always; the line and the four answers (SignupState's order) when [answers] are given. */
        private fun liveBody(context: Context, content: RaidLiveContent, zoneIcon: Bitmap?, countdownMillis: Long, answers: List<NotificationCompat.Action>): RemoteViews =
            RemoteViews(context.packageName, R.layout.notification_raid_live).apply {
                if (zoneIcon != null) setImageViewBitmap(R.id.live_icon, zoneIcon) else setViewVisibility(R.id.live_icon, View.GONE)
                setTextViewText(R.id.live_title, content.title)
                setTextViewText(R.id.live_text, content.text)
                setChronometer(R.id.live_countdown, SystemClock.elapsedRealtime() + countdownMillis, null, true)
                setChronometerCountDown(R.id.live_countdown, true)
                if (countdownMillis < URGENT_MILLIS) setTextColor(R.id.live_countdown, FsColor.wipe.toArgb())
                if (answers.isEmpty() || content.line == null) setViewVisibility(R.id.live_line, View.GONE) else setTextViewText(R.id.live_line, content.line)
                if (answers.isEmpty()) setViewVisibility(R.id.live_answers, View.GONE)
                SignupState.entries.zip(ANSWER_VIEWS).zip(answers) { (state, view), action ->
                    setTextViewText(view, action.title)
                    setOnClickPendingIntent(view, action.actionIntent)
                    if (state == content.answer) {
                        setInt(view, "setBackgroundResource", R.drawable.bg_live_answer_on)
                        setTextColor(view, Color.WHITE)
                    }
                }
            }

        /** Inside 10 minutes the countdown turns red (the pull is close). */
        private const val URGENT_MILLIS = 10L * 60 * 1000

        /** The answer pills, in SignupState's order (In, Late, Bench, Out). */
        private val ANSWER_VIEWS = listOf(R.id.live_in, R.id.live_late, R.id.live_bench, R.id.live_out)

        fun simNotification(context: Context, sim: SimLive, content: SimLiveContent): Notification {
            // The seal's turquoise tints the icon, like every notification; the bar below matches it.
            val builder = PushNotifier.builder(context, CHANNEL_SIMS)
                .setContentTitle(content.title)
                .setContentText(content.text)
                .setSubText(content.subText)
                .setShowWhen(false)
                .setOngoing(content.ongoing)
                .setAutoCancel(!content.ongoing)
                .setOnlyAlertOnce(true)
                .setCategory(NotificationCompat.CATEGORY_PROGRESS)
                .setContentIntent(open(context, DeepLink.Sim(sim.simId).url()))
            content.progress?.let { (done, max) ->
                val body = RemoteViews(context.packageName, R.layout.notification_sim_progress).apply {
                    setTextViewText(R.id.sim_title, content.title)
                    setTextViewText(R.id.sim_text, content.text)
                    setProgressBar(R.id.sim_progress, max, done, content.indeterminate)
                }
                builder.setStyle(NotificationCompat.DecoratedCustomViewStyle()).setCustomContentView(body).setCustomBigContentView(body)
            }
            return builder.build()
        }

        private fun open(context: Context, url: String): PendingIntent {
            val intent = Intent(Intent.ACTION_VIEW, url.toUri(), context, MainActivity::class.java)
                .addFlags(Intent.FLAG_ACTIVITY_NEW_TASK or Intent.FLAG_ACTIVITY_CLEAR_TOP)
            return PendingIntent.getActivity(context, url.hashCode(), intent, PendingIntent.FLAG_UPDATE_CURRENT or PendingIntent.FLAG_IMMUTABLE)
        }
    }
}
