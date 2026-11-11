package gg.foreversixty.app.raids.push

import android.Manifest
import android.app.Application
import android.app.Notification
import android.app.NotificationManager
import androidx.test.core.app.ApplicationProvider
import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.model.RaidPingResult
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

/** RAID_SHORT (GO N1): one `Mention healers` button, the raid's Discord ping from the shade, the answer in its place. */
@RunWith(AndroidJUnit4::class)
class RaidShortPushTest {
    private val data: Map<String, String> = RaidFixture.push("push.raid-short.fcm").getValue("message").jsonObject
        .getValue("data").jsonObject.mapValues { it.value.jsonPrimitive.content }
    private val payload = requireNotNull(RaidShortPushPayload.parse(data))
    private val mention = payload.actions.single()

    @Test fun `the payload names the raid and the ping for its short role`() {
        assertEquals(2L, payload.guild)
        assertEquals(7L, payload.raid)
        assertEquals(RaidShortPushCategory.ACTION_MENTION, mention.id)
        assertEquals("healer", payload.role(mention))
        assertEquals("/v1/guilds/2/raids/7/ping", requireNotNull(payload.call(mention)).path)
        assertNull(payload.call(mention.copy(path = "/v1/guilds/2/raids/8/ping")))
    }

    @Test fun `a ping that posts says who it mentioned over the raid`() = runTest {
        val result = RaidPingResult(posted = true, mentioned = 3, noDiscord = 0)
        val answer = RaidShortPushAnswerer({ _, _ -> ApiResult.Success(result, "test") }, resources()).answer(payload, mention.id)
        assertEquals("Mentioned 3 healers in Discord.", answer.title)
        assertEquals(payload.title, answer.body)
    }

    @Test fun `a refused press says who it could not mention and how to retry, an unknown one sends nothing`() = runTest {
        val refused = RaidShortPushAnswerer({ _, _ -> ApiResult.Failure(ApiError.Status(HTTP_TOO_MANY)) }, resources())
        val failed = refused.answer(payload, mention.id)
        // Round 7 row 2: the role as the button names it, then the raid and the retry.
        assertEquals("Could not mention the healers", failed.title)
        assertEquals("Onyxia's Lair. Open the raid to try again.", failed.body)
        // Without the server's words for who, the plain failure title.
        assertEquals("That did not save. Try again.", refused.failed(payload.copy(mentionWho = null)).title)
        var calls = 0
        RaidShortPushAnswerer({ _, _ -> calls++; null }, resources()).answer(payload, "raid_mention_all")
        assertEquals(0, calls)
    }

    @Test fun `the notification carries Mention healers through the receiver, in big text`() {
        val app = ApplicationProvider.getApplicationContext<Application>()
        shadowOf(app).grantPermissions(Manifest.permission.POST_NOTIFICATIONS)
        PushNotifier.show(app, requireNotNull(PushMessage.parse(data, null, null, collapse = data["collapse"])))

        val posted = requireNotNull(shadowOf(app.getSystemService(NotificationManager::class.java)).allNotifications.singleOrNull())
        assertEquals(listOf("Mention healers"), posted.actions.map { it.title.toString() })
        val intent = shadowOf(posted.actions.single().actionIntent).savedIntent
        assertEquals(OfficerActionReceiver::class.java.name, intent.component?.className)
        assertEquals(RaidShortPushCategory.CATEGORY, intent.getStringExtra(OfficerActionReceiver.EXTRA_CATEGORY))
        // GO N3: the expanded view shows the whole line, the short role first.
        assertEquals(data["body"], posted.extras.getCharSequence(Notification.EXTRA_BIG_TEXT)?.toString())
    }

    private companion object {
        const val HTTP_TOO_MANY = 429
    }
}
