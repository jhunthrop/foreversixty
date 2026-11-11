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

/** GUILD_APPLICANT (GO N2): the payload's Accept and Decline on the notification, the receiver's POST, the answer in its place. */
@RunWith(AndroidJUnit4::class)
class ApplicantPushTest {
    private val data: Map<String, String> = RaidFixture.push("push.guild-applicant.fcm").getValue("message").jsonObject
        .getValue("data").jsonObject.mapValues { it.value.jsonPrimitive.content }
    private val payload = requireNotNull(ApplicantPushPayload.parse(data))
    private val decision = RaidFixture.load(GuildApplicationDecision.serializer(), "guild.application.accept")

    @Test fun `the payload names the application, Accept then Decline, each on its own route`() {
        assertEquals(2L, payload.guild)
        assertEquals(31L, payload.application)
        val (accept, decline) = payload.actions
        assertEquals(listOf(ApplicantPushCategory.ACTION_ACCEPT, ApplicantPushCategory.ACTION_DECLINE), payload.actions.map { it.id })
        assertEquals("/v1/guilds/2/applications/31/accept", requireNotNull(payload.call(accept)).path)
        assertEquals("/v1/guilds/2/applications/31/decline", requireNotNull(payload.call(decline)).path)
        assertNull(payload.call(decline.copy(path = "/v1/guilds/2/applications/32/decline")))
        assertNull(payload.call(accept.copy(path = decline.path)))
    }

    @Test fun `an accept or a decline that saves replaces the push with its answer`() = runTest {
        val pressed = mutableListOf<String>()
        val answerer = ApplicantPushAnswerer({ _, action -> pressed += action.id; ApiResult.Success(decision, "test") }, resources())
        assertEquals("Application accepted", answerer.answer(payload, ApplicantPushCategory.ACTION_ACCEPT).title)
        val declined = answerer.answer(payload, ApplicantPushCategory.ACTION_DECLINE)
        assertEquals("Application declined", declined.title)
        assertEquals(payload.title, declined.body)
        assertEquals(listOf(ApplicantPushCategory.ACTION_ACCEPT, ApplicantPushCategory.ACTION_DECLINE), pressed)
    }

    @Test fun `a refused press says what failed for whom and how to retry (round 7 row 2)`() = runTest {
        val refused = ApplicantPushAnswerer({ _, _ -> ApiResult.Failure(ApiError.Status(HTTP_CONFLICT)) }, resources())
        val declined = refused.answer(payload, ApplicantPushCategory.ACTION_DECLINE)
        assertEquals("Could not decline Mirelle", declined.title)
        assertEquals("Mirelle, Holy Priest 60. Open Forever Sixty to try again.", declined.body)
        assertEquals("Could not accept Mirelle", refused.answer(payload, ApplicantPushCategory.ACTION_ACCEPT).title)
    }

    @Test fun `an unknown press sends nothing`() = runTest {
        var calls = 0
        ApplicantPushAnswerer({ _, _ -> calls++; null }, resources()).answer(payload, "application_hold")
        assertEquals(0, calls)
    }

    @Test fun `a push without the applicant's words falls back to the plain failure`() = runTest {
        val bare = payload.copy(applicantName = null, applicant = null)
        val refused = ApplicantPushAnswerer({ _, _ -> ApiResult.Failure(ApiError.Status(HTTP_CONFLICT)) }, resources())
        assertEquals("That did not save. Try again.", refused.answer(bare, ApplicantPushCategory.ACTION_ACCEPT).title)
    }

    @Test fun `the notification draws Accept then Decline, posted through the receiver, and its tap opens the application`() {
        val app = ApplicationProvider.getApplicationContext<Application>()
        shadowOf(app).grantPermissions(Manifest.permission.POST_NOTIFICATIONS)
        PushNotifier.show(app, requireNotNull(PushMessage.parse(data, null, null, collapse = data["collapse"])))

        val posted = requireNotNull(shadowOf(app.getSystemService(NotificationManager::class.java)).allNotifications.singleOrNull())
        val tap: Intent = shadowOf(posted.contentIntent).savedIntent
        assertEquals(DeepLink.GuildApplicant(2, 31).url(), tap.dataString)
        assertEquals(listOf("Accept", "Decline"), posted.actions.map { it.title.toString() })
        val intent = shadowOf(posted.actions.last().actionIntent).savedIntent
        assertEquals(OfficerActionReceiver::class.java.name, intent.component?.className)
        assertEquals(ApplicantPushCategory.ACTION_DECLINE, intent.getStringExtra(OfficerActionReceiver.EXTRA_ACTION))
        assertEquals(ApplicantPushCategory.CATEGORY, intent.getStringExtra(OfficerActionReceiver.EXTRA_CATEGORY))
    }

    private companion object {
        const val HTTP_CONFLICT = 409
    }
}
