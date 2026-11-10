package gg.foreversixty.app.raids.push

import android.Manifest
import android.app.Application
import android.app.NotificationManager
import android.content.Intent
import androidx.test.core.app.ApplicationProvider
import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.model.GuildApplicationDecision
import gg.foreversixty.app.navigation.DeepLink
import gg.foreversixty.app.network.ApiError
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.push.PushMessage
import gg.foreversixty.app.push.PushNotifier
import gg.foreversixty.app.raids.RaidFixture
import gg.foreversixty.app.raids.resources
import kotlinx.coroutines.test.runTest
import kotlinx.serialization.json.jsonObject
import kotlinx.serialization.json.jsonPrimitive
import org.junit.Assert.assertEquals
import org.junit.Assert.assertNull
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.Shadows.shadowOf

/** GUILD_APPLICANT: the payload's Decline on the notification, the receiver's POST, the answer in its place. */
@RunWith(AndroidJUnit4::class)
class ApplicantPushTest {
    private val data: Map<String, String> = RaidFixture.push("push.guild-applicant.fcm").getValue("message").jsonObject
        .getValue("data").jsonObject.mapValues { it.value.jsonPrimitive.content }
    private val payload = requireNotNull(ApplicantPushPayload.parse(data))
    private val decision = RaidFixture.load(GuildApplicationDecision.serializer(), "guild.application.accept")

    @Test fun `the payload names the application and its Decline`() {
        assertEquals(2L, payload.guild)
        assertEquals(31L, payload.application)
        val decline = payload.actions.single()
        assertEquals(ApplicantPushCategory.ACTION_DECLINE, decline.id)
        assertEquals("/v1/guilds/2/applications/31/decline", requireNotNull(payload.declineCall(decline)).path)
        assertNull(payload.declineCall(decline.copy(path = "/v1/guilds/2/applications/32/decline")))
    }

    @Test fun `a decline that saves replaces the push with the answer`() = runTest {
        var calls = 0
        val answerer = ApplicantPushAnswerer({ _, _ -> calls++; ApiResult.Success(decision, "test") }, resources())
        val answer = answerer.answer(payload, ApplicantPushCategory.ACTION_DECLINE)
        assertEquals(1, calls)
        assertEquals("Application declined", answer.title)
        assertEquals(payload.title, answer.body)
    }

    @Test fun `a refused or unknown press answers that it did not save`() = runTest {
        val refused = ApplicantPushAnswerer({ _, _ -> ApiResult.Failure(ApiError.Status(HTTP_CONFLICT)) }, resources())
        val failed = refused.answer(payload, ApplicantPushCategory.ACTION_DECLINE)
        assertEquals("Open the application to answer it.", failed.body)
        var calls = 0
        val unknown = ApplicantPushAnswerer({ _, _ -> calls++; null }, resources()).answer(payload, "application_accept")
        assertEquals(0, calls)
        assertEquals(failed, unknown)
    }

    @Test fun `the notification draws the payload's Decline, posted through the receiver, and its tap opens the application`() {
        val app = ApplicationProvider.getApplicationContext<Application>()
        shadowOf(app).grantPermissions(Manifest.permission.POST_NOTIFICATIONS)
        PushNotifier.show(app, requireNotNull(PushMessage.parse(data, null, null, collapse = data["collapse"])))

        val posted = requireNotNull(shadowOf(app.getSystemService(NotificationManager::class.java)).allNotifications.singleOrNull())
        val tap: Intent = shadowOf(posted.contentIntent).savedIntent
        assertEquals(DeepLink.GuildApplicant(2, 31).url(), tap.dataString)
        val button = posted.actions.single()
        assertEquals("Decline", button.title.toString())
        val intent = shadowOf(button.actionIntent).savedIntent
        assertEquals(ApplicantActionReceiver::class.java.name, intent.component?.className)
        assertEquals(ApplicantPushCategory.ACTION_DECLINE, intent.getStringExtra(ApplicantActionReceiver.EXTRA_ACTION))    }

    private companion object {
        const val HTTP_CONFLICT = 409
    }
}
