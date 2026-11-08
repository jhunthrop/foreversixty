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
import gg.foreversixty.app.raids.push.ApplicantNotificationActions
import gg.foreversixty.app.raids.push.RaidNotificationActions

/**
 * Draws a push on its group's channel (spec 2026-11-08 F1, F4): the title and line, the buttons its
 * category names (the raid answers, an applicant's Decline), and a tap that opens its answer. One notification per collapse id, so a later
 * message for the same raid replaces the earlier one.
 */
object PushNotifier {
    fun show(context: Context, message: PushMessage) {
        PushGroup.createChannels(context)
        val tag = message.collapse ?: "${message.kind}-${message.title.hashCode()}"
        val id = message.group.ordinal + 1
        val builder = NotificationCompat.Builder(context, message.group.channelId)
            .setSmallIcon(R.drawable.ic_notification_seal)
            .setContentTitle(message.title)
            .setContentText(message.body)
            .setStyle(NotificationCompat.BigTextStyle().bigText(message.body))
            .setContentIntent(openIntent(context, message.landing(), tag))
            .setAutoCancel(true)
        message.data[PushMessage.KEY_RAID]?.let { builder.setGroup("raid-$it") }
        actions(context, message, tag, id).forEach(builder::addAction)
        notify(context, tag, id, builder.build())
    }

    /** The buttons the message's category names ([PushActionSet]). */
    private fun actions(context: Context, message: PushMessage, tag: String, id: Int): List<NotificationCompat.Action> =
        when (val set = PushActionSet.of(message)) {
            is PushActionSet.Raid -> RaidNotificationActions.actions(context, set.category, message.data, tag, id)
            is PushActionSet.Applicant -> ApplicantNotificationActions.actions(context, set.payload, tag, id)
            PushActionSet.None -> emptyList()
        }

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

    /** The tap: the app on [link] (an app link MainActivity routes), or the app as it was when null. */
    fun openIntent(context: Context, link: String?, tag: String?): PendingIntent {
        val intent = launch(context, link)
        return PendingIntent.getActivity(
            context, (tag ?: link ?: "").hashCode(), intent,
            PendingIntent.FLAG_UPDATE_CURRENT or PendingIntent.FLAG_IMMUTABLE,
        )
    }

    /** Opens the app on [url] from a receiver. */
    fun open(context: Context, url: String) {
        context.startActivity(launch(context, url).addFlags(Intent.FLAG_ACTIVITY_NEW_TASK))
    }

    private fun launch(context: Context, link: String?): Intent {
        val intent = link?.let { Intent(Intent.ACTION_VIEW, it.toUri(), context, MainActivity::class.java) }
            ?: Intent(context, MainActivity::class.java)
        return intent.addFlags(Intent.FLAG_ACTIVITY_CLEAR_TOP or Intent.FLAG_ACTIVITY_SINGLE_TOP)
    }
}
