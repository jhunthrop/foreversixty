package gg.foreversixty.app.raids.flows

import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.model.GuildRaids
import gg.foreversixty.app.model.Me
import gg.foreversixty.app.model.mainCharacter
import gg.foreversixty.app.model.ReadyCheck
import gg.foreversixty.app.raids.RaidFixture
import gg.foreversixty.app.raids.RaidTime
import gg.foreversixty.app.raids.clock
import gg.foreversixty.app.raids.logic.GuildHead
import gg.foreversixty.app.raids.logic.RaidsTabContent
import gg.foreversixty.app.raids.logic.RaidsTabState
import gg.foreversixty.app.raids.resources
import org.junit.Assert.assertTrue
import org.junit.Test
import org.junit.runner.RunWith

/** Before you pull is a 60's check: the Guild tab shows it at 60 and leaves it out below (spec 2026-11-10-guild-tab.md 2.1). */
@RunWith(AndroidJUnit4::class)
class GuildTabReadyTest {
    private val check = RaidFixture.load(ReadyCheck.serializer(), "character.ready")

    private fun guild(level: Int): RaidTabView.Guild {
        val me = RaidFixture.load(Me.serializer(), "me")
        val raids = RaidFixture.load(GuildRaids.serializer(), "guild.raids")
        val character = requireNotNull(me.mainCharacter).copy(level = level)
        val content = RaidsTabContent(
            head = GuildHead.of(requireNotNull(character.guild), me.guilds.firstOrNull(), resources()),
            character = character,
            planned = raids.planned.orEmpty(),
            nights = raids.rows,
            waiting = raids.waitingForApproval,
            offlineSince = null,
        )
        val view = RaidTabPresenter(resources(), clock(), RaidTime.now)
            .view(RaidsTabState.Loaded(content), null, null, true, RaidTabExtras(ready = check))
        return view as RaidTabView.Guild
    }

    @Test fun `at 60 Before you pull lists the check`() {
        assertTrue(guild(60).ready.isNotEmpty())
    }

    @Test fun `below 60 Before you pull is left out, and so are the 60s' lines under Who is coming`() {
        val g = guild(36)
        assertTrue(g.ready.isEmpty())
        assertTrue(g.foot.isEmpty())
    }
}
