package gg.foreversixty.app.widgets

import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.raids.resources
import gg.foreversixty.app.widgets.logic.LivePush
import gg.foreversixty.app.widgets.logic.WidgetFigures
import org.junit.Assert.assertEquals
import org.junit.Assert.assertNull
import org.junit.Test
import org.junit.runner.RunWith
import java.time.Instant
import java.time.ZoneId

/** `live.raid.start` (fixture surface.live-raid-start.fcm) draws the raid-night notification on its own. */
@RunWith(AndroidJUnit4::class)
class LiveRaidPushTest {
    private val start = mapOf(
        "kind" to "live.raid.start", "raid_id" to "7", "guild_id" to "2", "guild_name" to "OLYMPUS XXVII",
        "zone_name" to "Onyxia's Lair", "character_name" to "Lightningbolltz", "starts_at" to "1796954400",
        "status" to "open", "my_state" to "in", "role" to "dps", "signed_up" to "38", "size" to "40",
        "zone_icon" to "inv_misc_head_dragon_black", "in" to "36", "late" to "2", "tanks" to "2", "healers" to "5",
    )
    private val figures = WidgetFigures(resources(), ZoneId.of("America/Chicago"))

    @Test fun startAndEndParse() {
        assertEquals(LivePush.Raid(LivePush.Event.Start, 2, 7), LivePush.parse(start))
        assertEquals(LivePush.Event.End, (LivePush.parse(start + ("kind" to "live.raid.end")) as LivePush.Raid).event)
    }

    @Test fun aStartMessageDrawsTheCountdown() {
        val content = figures.pushedLive(start)
        assertEquals("Onyxia's Lair", content?.title)
        assertEquals("OLYMPUS XXVII · You are In", content?.text)
        assertEquals(Instant.ofEpochSecond(1796954400L), content?.countdownTo)
        assertEquals("Starts 8:00 pm CST · 36 in, 2 late · 2 tanks, 5 heals", content?.line)
        assertEquals("inv_misc_head_dragon_black", content?.zoneIcon)
    }

    @Test fun nobodyOutGetsACountdown() {
        assertNull(figures.pushedLive(start + ("my_state" to "out")))
    }
}
