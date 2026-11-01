package gg.foreversixty.app.system

import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.raids.resources
import gg.foreversixty.app.system.logic.ShareMessage
import gg.foreversixty.app.system.logic.ShareText
import gg.foreversixty.app.system.logic.TierRole
import java.io.File
import java.time.Instant
import java.time.ZoneId
import org.junit.Assert.assertEquals
import org.junit.Assert.assertNull
import org.junit.Assert.assertTrue
import org.junit.Test
import org.junit.runner.RunWith

/** The share texts (spec §5.7 Share) and the two shortcuts.xml copies. */
@RunWith(AndroidJUnit4::class)
class ShareAndShortcutsXmlTest {
    private val res = resources()

    @Test fun `a character shares its page and its one line`() {
        val url = "https://foreversixty.gg/character/us/normal/lightningbolltz"
        assertEquals(
            ShareMessage("Lightningbolltz, Elemental Shaman, level 60", url),
            ShareText.character(res, url, "Lightningbolltz", "Elemental", "Shaman", 60),
        )
        assertEquals(ShareMessage("Lightningbolltz", url), ShareText.character(res, url, "Lightningbolltz", null, "Shaman", 60))
    }

    @Test fun `a guide shares its class and spec page`() {
        assertEquals(
            ShareMessage("Elemental Shaman guide", "https://foreversixty.gg/guides/shaman/elemental"),
            ShareText.guide(res, "shaman", "elemental", "Elemental", "Shaman"),
        )
    }

    @Test fun `a tier list shares its role's page`() {
        assertEquals(ShareMessage("DPS tier list", "https://foreversixty.gg/tiers"), ShareText.tiers(res, TierRole.Dps))
        assertEquals(ShareMessage("Tank tier list", "https://foreversixty.gg/tiers/tank"), ShareText.tiers(res, TierRole.Tank))
        assertEquals(ShareMessage("Healer tier list", "https://foreversixty.gg/tiers/healer"), ShareText.tiers(res, TierRole.Healer))
    }

    @Test fun `public and unlisted reports share, guild and private ones do not`() {
        val at = Instant.parse("2026-11-06T03:00:00Z")
        val zone = ZoneId.of("America/Chicago")
        assertEquals(
            ShareMessage("Onyxia's Lair · 5 Nov", "https://foreversixty.gg/reports/abc123"),
            ShareText.report(res, "abc123", "Onyxia's Lair", at, "public", zone),
        )
        assertTrue(ShareText.report(res, "abc123", "Onyxia's Lair", at, "unlisted", zone) != null)
        assertNull(ShareText.report(res, "abc123", "Onyxia's Lair", at, "guild", zone))
        assertNull(ShareText.report(res, "abc123", "Onyxia's Lair", at, "private", zone))
    }

    @Test fun `the share sheet sends the sentence then the link`() {
        assertEquals("DPS tier list\nhttps://foreversixty.gg/tiers", ShareText.tiers(res, TierRole.Dps).body)
    }

    @Test fun `the debug shortcuts are the release ones with the debug package`() {
        val release = File("src/main/res/xml/shortcuts.xml").readText()
        val debug = File("src/debug/res/xml/shortcuts.xml").readText()
        fun strip(xml: String) = xml.replace(Regex("<!--.*?-->", RegexOption.DOT_MATCHES_ALL), "")
        assertEquals(strip(release).replace("\"gg.foreversixty.app\"", "\"gg.foreversixty.app.debug\""), strip(debug))
        listOf("next_raid", "signup_in", "upgrades", "search").forEach { assertTrue(release.contains("android:shortcutId=\"$it\"")) }
        assertTrue(release.contains("actions.intent.OPEN_APP_FEATURE"))
    }
}
