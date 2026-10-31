package gg.foreversixty.app.raids.push

import android.Manifest
import android.app.PendingIntent
import android.content.BroadcastReceiver
import android.content.Context
import android.content.Intent
import android.content.pm.PackageManager
import android.os.Bundle
import androidx.core.app.NotificationCompat
import androidx.core.app.NotificationManagerCompat
import androidx.core.content.ContextCompat
import androidx.core.net.toUri
import gg.foreversixty.app.ForeverApplication
import gg.foreversixty.app.MainActivity
import gg.foreversixty.app.R
import gg.foreversixty.app.raids.logic.LiveRaidService
import gg.foreversixty.app.widgets.WidgetRefreshWork
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.SupervisorJob
import kotlinx.coroutines.launch
import kotlinx.coroutines.withTimeoutOrNull

/**
 * In, Late and Out answered from the notification shade (spec 6.1 Background actions): `goAsync`,
 * the PUT inside the receiver's window, then the answer replaces the notification by its tag. A
 * foreground action, or a press the server cannot place, opens the raid page instead. Lane F7 builds
 * the notification from the FCM message and attaches [RaidNotificationActions.actions].
 */
class RaidActionReceiver : BroadcastReceiver() {
    override fun onReceive(context: Context, intent: Intent) {
        val action = intent.getStringExtra(EXTRA_ACTION) ?: return
        val data = intent.getBundleExtra(EXTRA_DATA)?.toStringMap() ?: return
        val payload = RaidPushPayload.parse(data) ?: return
        val tag = intent.getStringExtra(EXTRA_TAG)
        val id = intent.getIntExtra(EXTRA_ID, 0)
        val state = RaidPushCategory.stateFor(action)
        if (state == null) {
            open(context, payload.deepLink(askingSpec = action == RaidPushCategory.ACTION_ASK_SPEC).url())
            return
        }
        val pending = goAsync()
        val app = context.applicationContext as ForeverApplication
        val answerer = RaidPushAnswerer(LiveRaidService(app.graph.apiClient), context.resources)
        CoroutineScope(SupervisorJob() + Dispatchers.IO).launch {
            try {
                val outcome = withTimeoutOrNull(BUDGET_MILLIS) { answerer.answer(payload, state) }
                    ?: RaidPushOutcome.Answered(payload.answer(RaidPushPayload.FAILED, context.resources))
                when (outcome) {
                    is RaidPushOutcome.Answered -> {
                        post(context, outcome.answer, tag, id)
                        // Lane F14: the widgets, tile and countdown redraw from the server's answer.
                        WidgetRefreshWork.now(context)
                    }
                    is RaidPushOutcome.Open -> open(context, outcome.link.url())
                }
            } finally {
                pending.finish()
            }
        }
    }

    private fun open(context: Context, url: String) {
        val intent = Intent(Intent.ACTION_VIEW, url.toUri(), context, MainActivity::class.java)
            .addFlags(Intent.FLAG_ACTIVITY_NEW_TASK or Intent.FLAG_ACTIVITY_CLEAR_TOP)
        context.startActivity(intent)
    }

    /** The answer replaces the question: the same tag and id, no actions. */
    private fun post(context: Context, answer: RaidPushPayload.Answer, tag: String?, id: Int) {
        if (ContextCompat.checkSelfPermission(context, Manifest.permission.POST_NOTIFICATIONS) != PackageManager.PERMISSION_GRANTED &&
            android.os.Build.VERSION.SDK_INT >= android.os.Build.VERSION_CODES.TIRAMISU
        ) {
            return
        }
        val notification = NotificationCompat.Builder(context, RaidNotificationActions.CHANNEL_RAIDS)
            .setSmallIcon(R.mipmap.ic_launcher)
            .setContentTitle(answer.title)
            .setContentText(answer.body)
            .setStyle(NotificationCompat.BigTextStyle().bigText(answer.body))
            .setAutoCancel(true)
            .build()
        NotificationManagerCompat.from(context).notify(tag, id, notification)
    }

    companion object {
        const val EXTRA_ACTION = "gg.foreversixty.raids.ACTION"
        const val EXTRA_DATA = "gg.foreversixty.raids.DATA"
        const val EXTRA_TAG = "gg.foreversixty.raids.TAG"
        const val EXTRA_ID = "gg.foreversixty.raids.ID"

        /** Inside the receiver's ~10 s window, with room to post the answer. */
        private const val BUDGET_MILLIS = 8_000L

        private fun Bundle.toStringMap(): Map<String, String> =
            keySet().mapNotNull { key -> getString(key)?.let { key to it } }.toMap()
    }
}

/** The In, Late and Out buttons for a raid notification, for lane F7's notification builder. */
object RaidNotificationActions {
    /** The `Raids` channel (spec 6.1); lane F7 creates the channels. */
    const val CHANNEL_RAIDS = "raids"

    fun actions(context: Context, category: String, data: Map<String, String>, tag: String?, id: Int): List<NotificationCompat.Action> {
        val ids = when (category) {
            RaidPushCategory.SIGNUP -> listOf(
                RaidPushCategory.ACTION_IN to R.string.push_action_in,
                RaidPushCategory.ACTION_LATE to R.string.push_action_late,
                RaidPushCategory.ACTION_OUT to R.string.push_action_out,
            )
            RaidPushCategory.ASK_SPEC -> listOf(
                RaidPushCategory.ACTION_ASK_SPEC to R.string.push_action_askSpec,
                RaidPushCategory.ACTION_OUT to R.string.push_action_out,
            )
            RaidPushCategory.CHOOSE -> listOf(RaidPushCategory.ACTION_CHOOSE to R.string.push_action_choose)
            RaidPushCategory.LIVE -> listOf(
                RaidPushCategory.ACTION_LATE to R.string.push_action_late,
                RaidPushCategory.ACTION_OUT to R.string.push_action_out,
            )
            else -> emptyList()
        }
        val bundle = Bundle().apply { data.forEach { (key, value) -> putString(key, value) } }
        return ids.mapIndexed { index, (action, label) ->
            val intent = Intent(context, RaidActionReceiver::class.java)
                .setAction(action)
                .putExtra(RaidActionReceiver.EXTRA_ACTION, action)
                .putExtra(RaidActionReceiver.EXTRA_DATA, bundle)
                .putExtra(RaidActionReceiver.EXTRA_TAG, tag)
                .putExtra(RaidActionReceiver.EXTRA_ID, id)
            val pending = PendingIntent.getBroadcast(
                context, (tag ?: "").hashCode() * ACTION_SLOTS + index, intent,
                PendingIntent.FLAG_UPDATE_CURRENT or PendingIntent.FLAG_IMMUTABLE,
            )
            NotificationCompat.Action.Builder(0, context.getString(label), pending).build()
        }
    }

    private const val ACTION_SLOTS = 4
}
