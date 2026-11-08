package gg.foreversixty.app.raids.push

import android.app.PendingIntent
import android.content.BroadcastReceiver
import android.content.Context
import android.content.Intent
import androidx.core.app.NotificationCompat
import gg.foreversixty.app.ForeverApplication
import gg.foreversixty.app.R
import gg.foreversixty.app.push.PushGroup
import gg.foreversixty.app.push.PushNotifier
import gg.foreversixty.app.raids.flows.LiveRaidFlowService
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.SupervisorJob
import kotlinx.coroutines.launch
import kotlinx.coroutines.withTimeoutOrNull

/**
 * Decline on a guild.applicant notification, answered in the background like the raid answers
 * ([RaidActionReceiver]): `goAsync`, the POST inside the receiver's window, then the answer replaces
 * the notification by its tag. A refusal the shade cannot word opens the applicant sheet.
 */
class ApplicantActionReceiver : BroadcastReceiver() {
    override fun onReceive(context: Context, intent: Intent) {
        if (intent.getStringExtra(EXTRA_ACTION) != ApplicantPushPayload.ACTION_DECLINE) return
        val guild = intent.getLongExtra(EXTRA_GUILD, 0L).takeIf { it > 0 } ?: return
        val application = intent.getLongExtra(EXTRA_APPLICATION, 0L).takeIf { it > 0 } ?: return
        val payload = ApplicantPushPayload(guild, application)
        val tag = intent.getStringExtra(EXTRA_TAG)
        val id = intent.getIntExtra(EXTRA_ID, 0)
        val pending = goAsync()
        val app = context.applicationContext as ForeverApplication
        val decliner = ApplicantPushDecliner(LiveRaidFlowService(app.graph.apiClient), context.resources)
        CoroutineScope(SupervisorJob() + Dispatchers.IO).launch {
            try {
                when (val outcome = withTimeoutOrNull(BUDGET_MILLIS) { decliner.decline(payload) } ?: decliner.failed()) {
                    is ApplicantDeclineOutcome.Answered -> post(context, outcome, payload, tag, id)
                    is ApplicantDeclineOutcome.Open -> PushNotifier.open(context, outcome.link.url())
                }
            } finally {
                pending.finish()
            }
        }
    }

    /** The answer replaces the applicant: the same tag and id; a failure keeps Decline. */
    private fun post(context: Context, answer: ApplicantDeclineOutcome.Answered, payload: ApplicantPushPayload, tag: String?, id: Int) {
        val builder = NotificationCompat.Builder(context, PushGroup.Officer.channelId)
            .setSmallIcon(R.drawable.ic_notification_seal)
            .setContentTitle(answer.title)
            .setContentText(answer.body)
            .setStyle(NotificationCompat.BigTextStyle().bigText(answer.body))
            .setContentIntent(PushNotifier.openIntent(context, payload.deepLink.url(), tag))
            .setAutoCancel(true)
        if (answer.retry) {
            ApplicantNotificationActions.actions(context, payload, tag, id).forEach(builder::addAction)
        }
        PushNotifier.notify(context, tag, id, builder.build())
    }

    companion object {
        const val EXTRA_ACTION = "gg.foreversixty.applicant.ACTION"
        const val EXTRA_GUILD = "gg.foreversixty.applicant.GUILD"
        const val EXTRA_APPLICATION = "gg.foreversixty.applicant.APPLICATION"
        const val EXTRA_TAG = "gg.foreversixty.applicant.TAG"
        const val EXTRA_ID = "gg.foreversixty.applicant.ID"

        /** Inside the receiver's ~10 s window, with room to post the answer. */
        private const val BUDGET_MILLIS = 8_000L
    }
}

/** The applicant notification's one button, Decline (Review is the notification itself). */
object ApplicantNotificationActions {
    fun actions(context: Context, payload: ApplicantPushPayload, tag: String?, id: Int): List<NotificationCompat.Action> {
        val intent = Intent(context, ApplicantActionReceiver::class.java)
            .setAction(ApplicantPushPayload.ACTION_DECLINE)
            .putExtra(ApplicantActionReceiver.EXTRA_ACTION, ApplicantPushPayload.ACTION_DECLINE)
            .putExtra(ApplicantActionReceiver.EXTRA_GUILD, payload.guild)
            .putExtra(ApplicantActionReceiver.EXTRA_APPLICATION, payload.application)
            .putExtra(ApplicantActionReceiver.EXTRA_TAG, tag)
            .putExtra(ApplicantActionReceiver.EXTRA_ID, id)
        val pending = PendingIntent.getBroadcast(
            context, (tag ?: payload.toString()).hashCode(), intent,
            PendingIntent.FLAG_UPDATE_CURRENT or PendingIntent.FLAG_IMMUTABLE,
        )
        return listOf(NotificationCompat.Action.Builder(0, context.getString(R.string.push_action_decline), pending).build())
    }
}
