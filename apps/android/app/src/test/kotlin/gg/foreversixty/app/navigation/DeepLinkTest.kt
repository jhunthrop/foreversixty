package gg.foreversixty.app.navigation

import kotlinx.serialization.json.Json
import kotlinx.serialization.json.jsonArray
import kotlinx.serialization.json.jsonObject
import kotlinx.serialization.json.jsonPrimitive
import org.junit.Assert.assertEquals
import org.junit.Assert.assertNull
import org.junit.Test
import java.io.File

/** apps/shared/links.json, read in place: every route's example opens its screen and round-trips. */
class DeepLinkTest {
    private val table = Json.parseToJsonElement(
        File(System.getProperty("user.dir"), "../../shared/links.json").readText(),
    ).jsonObject

    @Test
    fun `every route's example opens its screen and builds the same link back`() {
        val host = table.getValue("host").jsonPrimitive.content
        assertEquals(DeepLink.HOST, host)
        for (route in table.getValue("routes").jsonArray.map { it.jsonObject }) {
            val path = route.getValue("path").jsonPrimitive.content
            val example = route.getValue("example").jsonPrimitive.content
            val link = requireNotNull(DeepLink.parse("https://$host$example")) { "$path did not parse $example" }
            assertEquals(path, route.getValue("screen").jsonPrimitive.content, link.screen)
            assertEquals(path, link, DeepLink.parse(link.url()))
        }
    }

    @Test
    fun `each route reads its parameters`() {
        assertEquals(DeepLink.Raid(2, 41, "ask", "in"), DeepLink.parse("https://foreversixty.gg/app/raid/2/41?spec=ask&state=in"))
        assertEquals(
            DeepLink.Character("eu", "normal", "jäger"),
            DeepLink.parse("https://foreversixty.gg/app/character/eu/normal/j%C3%A4ger"),
        )
        assertEquals(DeepLink.RaidNight(2, 41), DeepLink.parse("https://foreversixty.gg/app/raid/2/41/night"))
        assertEquals(DeepLink.Guild(7), DeepLink.parse("https://foreversixty.gg/app/guild/7"))
        assertEquals(DeepLink.Bis("warrior-fury", "40"), DeepLink.parse("https://foreversixty.gg/app/bis/warrior-fury/40"))
        assertEquals(DeepLink.Sim("abc"), DeepLink.parse("https://foreversixty.gg/app/sim/abc"))
        assertEquals(DeepLink.Notifications, DeepLink.parse("https://foreversixty.gg/app/you/notifications"))
        assertEquals(DeepLink.Plan, DeepLink.parse("https://foreversixty.gg/app/you/plan"))
        assertEquals(DeepLink.SignIn("c", "s"), DeepLink.parse("https://foreversixty.gg/app/auth?code=c&state=s"))
    }

    @Test
    fun `anything else is not an app link`() {
        listOf(
            "https://foreversixty.gg/raid/2/41",
            "https://evil.example/app/raid/2/41",
            "http://foreversixty.gg/app/raid/2/41",
            "https://foreversixty.gg/app/raid/two/41",
            "https://foreversixty.gg/app/raid/2",
            "https://foreversixty.gg/app/you",
            "not a url at all",
        ).forEach { assertNull(it, DeepLink.parse(it)) }
    }

    @Test
    fun `an unknown query word is dropped, never trusted`() {
        assertEquals(DeepLink.Raid(2, 41), DeepLink.parse("https://foreversixty.gg/app/raid/2/41?state=maybe"))
    }
}
