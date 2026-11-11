package gg.foreversixty.app.push

import android.Manifest
import android.app.Notification
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
import gg.foreversixty.app.raids.push.OfficerNotificationActions
import gg.foreversixty.app.raids.push.RaidNotificationActions

/**
 * Draws a push on its group's channel (spec 2026-11-08 F1, F4): the title and the line in big text,
 * so the expanded view shows all of it (GO N3), the raid answers for RAID_SIGNUP, the payload's
 * Accept and Decline for GUILD_APPLICANT and its Mention for RAID_SHORT, and a tap that opens its
 * answer. One notification per collapse id, so a later message for the same raid replaces the earlier one.
 */
object PushNotifier {
    fun show(context: Context, message: PushMessage) {
        PushGroup.createChannels(context)
        val tag = message.collapse ?: "${message.kind}-${message.title.hashCode()}"
        val id = message.group.ordinal + 1
        val builder = builder(context, message.group.channelId)
            .setContentTitle(message.title)
            .setContentText(message.body)
            .setStyle(NotificationCompat.BigTextStyle().bigText(message.body))
            .setContentIntent(openIntent(context, message.landing(), tag))
            .setAutoCancel(true)
        message.data[PushMessage.KEY_RAID]?.let { builder.setGroup("raid-$it") }
        val actions = when (val category = message.category) {
            null -> emptyList()
            in OfficerNotificationActions.CATEGORIES -> OfficerNotificationActions.actions(context, category, message.data, tag, id)
            else -> RaidNotificationActions.actions(context, category, message.data, tag, id)
        }
        actions.forEach(builder::addAction)
        notify(context, tag, id, builder.build())
    }

    /** A push's builder on [channelId]: the flat seal as the small icon over the seal's turquoise. */
    fun builder(context: Context, channelId: String): NotificationCompat.Builder =
        NotificationCompat.Builder(context, channelId)
            .setSmallIcon(R.drawable.ic_notification_seal)
            .setColor(ContextCompat.getColor(context, R.color.notification_accent))

    /** Posts when the player allowed notifications (Android 13 and up ask; 9 to 12 never do). */
    fun notify(context: Context, tag: String?, id: Int, notification: Notification) {
        if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.TIRAMISU &&
            ContextCompat.checkSelfPermission(context, Manifest.permission.POST_NOTIFICATIONS) != PackageManager.PERMISSION_GRANTED
        ) {
            return
        }
        PushGroup.createChannels(context)
        NotificationManagerCompat.from(context).notify(tag, id, notification)
    }

    /**
     * The tap: the app on [link] (an app link MainActivity routes), or the app as it was when null. A
     * button that opens the app uses this too ([slot] apart from the tap), never a receiver that starts
     * an activity: Android 12 blocks notification trampolines.
     */
    fun openIntent(context: Context, link: String?, tag: String?, slot: Int = 0): PendingIntent {
        val intent = launch(context, link)
        return PendingIntent.getActivity(
            context, (tag ?: link ?: "").hashCode() + slot, intent,
            PendingIntent.FLAG_UPDATE_CURRENT or PendingIntent.FLAG_IMMUTABLE,
        )
    }

    private fun launch(context: Context, link: String?): Intent {
        val intent = link?.let { Intent(Intent.ACTION_VIEW, it.toUri(), context, MainActivity::class.java) }
            ?: Intent(context, MainActivity::class.java)
        return intent.addFlags(Intent.FLAG_ACTIVITY_CLEAR_TOP or Intent.FLAG_ACTIVITY_SINGLE_TOP)
    }
}
