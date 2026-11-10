package gg.foreversixty.app.raids.push

import android.app.PendingIntent
import android.content.BroadcastReceiver
import android.content.Context
import android.content.Intent
import android.os.Bundle
import androidx.core.app.NotificationCompat
import gg.foreversixty.app.ForeverApplication
import gg.foreversixty.app.push.PushActions
import gg.foreversixty.app.push.PushGroup
import gg.foreversixty.app.push.PushNotifier
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.SupervisorJob
import kotlinx.coroutines.launch
import kotlinx.coroutines.withTimeoutOrNull

/**
 * Decline answered from the shade (GUILD_APPLICANT): `goAsync`, the POST inside the receiver's window,
 * then the answer replaces the notification by its tag; a tap on the answer opens the application.
 */
class ApplicantActionReceiver : BroadcastReceiver() {
    override fun onReceive(context: Context, intent: Intent) {
        val action = intent.getStringExtra(EXTRA_ACTION) ?: return
        val raw = intent.getBundleExtra(EXTRA_DATA)?.let { bundle -> bundle.keySet().mapNotNull { key -> bundle.getString(key)?.let { key to it } }.toMap() } ?: return
        val payload = ApplicantPushPayload.parse(raw) ?: return
        val tag = intent.getStringExtra(EXTRA_TAG)
        val id = intent.getIntExtra(EXTRA_ID, 0)
        val pending = goAsync()
        val app = context.applicationContext as ForeverApplication
        val answerer = ApplicantPushAnswerer(LiveApplicantDecliner(app.graph.apiClient), context.resources)
        CoroutineScope(SupervisorJob() + Dispatchers.IO).launch {
            try {
                val answer = withTimeoutOrNull(BUDGET_MILLIS) { answerer.answer(payload, action) } ?: answerer.failed()
                val notification = PushNotifier.builder(context, PushGroup.Officer.channelId)
                    .setContentTitle(answer.title)
                    .setContentText(answer.body)
                    .setStyle(NotificationCompat.BigTextStyle().bigText(answer.body))
                    .setContentIntent(PushNotifier.openIntent(context, payload.link, tag))
                    .setAutoCancel(true)
                    .build()
                PushNotifier.notify(context, tag, id, notification)
            } finally {
                pending.finish()
            }
        }
    }

    companion object {
        const val EXTRA_ACTION = "gg.foreversixty.applicant.ACTION"
        const val EXTRA_DATA = "gg.foreversixty.applicant.DATA"
        const val EXTRA_TAG = "gg.foreversixty.applicant.TAG"
        const val EXTRA_ID = "gg.foreversixty.applicant.ID"

        /** Inside the receiver's ~10 s window, with room to post the answer. */
        private const val BUDGET_MILLIS = 8_000L
    }
}

/** The applicant notification's buttons, from the payload's `actions` (titles as the server words them). */
object ApplicantNotificationActions {
    fun actions(context: Context, data: Map<String, String>, tag: String?, id: Int): List<NotificationCompat.Action> {
        val bundle = Bundle().apply { data.forEach { (key, value) -> putString(key, value) } }
        return ApplicantPushCategory.answerable(PushActions.of(data)).mapIndexed { index, action ->
            val intent = Intent(context, ApplicantActionReceiver::class.java)
                .setAction(action.id)
                .putExtra(ApplicantActionReceiver.EXTRA_ACTION, action.id)
                .putExtra(ApplicantActionReceiver.EXTRA_DATA, bundle)
                .putExtra(ApplicantActionReceiver.EXTRA_TAG, tag)
                .putExtra(ApplicantActionReceiver.EXTRA_ID, id)
            val pending = PendingIntent.getBroadcast(
                context, (tag ?: "").hashCode() + index, intent,
                PendingIntent.FLAG_UPDATE_CURRENT or PendingIntent.FLAG_IMMUTABLE,
            )
            NotificationCompat.Action.Builder(0, action.title, pending).build()
        }
    }
}
