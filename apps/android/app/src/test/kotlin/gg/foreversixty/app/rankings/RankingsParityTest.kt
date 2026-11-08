package gg.foreversixty.app.rankings

import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.verticalScroll
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.test.junit4.v2.createComposeRule
import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.flows.FlowBackBar
import gg.foreversixty.app.model.CharacterRank
import gg.foreversixty.app.model.RankedEncounter
import gg.foreversixty.app.model.RankingGuild
import gg.foreversixty.app.model.RankingPlayer
import gg.foreversixty.app.model.RankingRow
import gg.foreversixty.app.model.RankingsPage
import gg.foreversixty.app.navigation.ForeverTab
import gg.foreversixty.app.navigation.ForeverTabScaffold
import gg.foreversixty.app.navigation.TabBadges
import gg.foreversixty.app.rankings.logic.Board
import gg.foreversixty.app.rankings.logic.BoardPhase
import gg.foreversixty.app.rankings.logic.Cohort
import gg.foreversixty.app.rankings.logic.Measure
import gg.foreversixty.app.rankings.logic.RankingsFilters
import gg.foreversixty.app.rankings.logic.RankingsViewer
import gg.foreversixty.app.rankings.logic.ViewerCharacter
import gg.foreversixty.app.rankings.ui.RankingsContent
import gg.foreversixty.app.rankings.ui.RankingsScreenState
import gg.foreversixty.app.testing.PARITY_PIXEL_9
import gg.foreversixty.app.testing.ParityCheck
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.annotation.Config
import org.robolectric.annotation.GraphicsMode
import java.time.Duration
import java.time.Instant
import java.time.ZoneId

/** Rankings pushed under Home (design/mocks/web web-rankings, web-rankings-you, web-rankings-empty). */
@RunWith(AndroidJUnit4::class)
@GraphicsMode(GraphicsMode.Mode.NATIVE)
@Config(qualifiers = PARITY_PIXEL_9)
class RankingsParityTest {
    @get:Rule val compose = createComposeRule()

    @Test fun rankings() = pushed("rankings") {
        RankingsContent(RankingsScreenState(BoardPhase.Choose(RankingsParityFixtures.encounters), null, emptyMap(), RankingsViewer.SignedOut), NOW, zone = ZONE)
    }

    @Test fun rankingsEmpty() = pushed("rankings-empty") {
        RankingsContent(RankingsScreenState(BoardPhase.Choose(emptyList()), null, emptyMap(), RankingsViewer.SignedOut), NOW, zone = ZONE)
    }

    @Test fun rankingsYou() = pushed("rankings-you") { RankingsContent(RankingsParityFixtures.board(), NOW, zone = ZONE) }

    @Test fun rankingsYouAt2() = pushed("rankings-you", fontScale = 2f) { RankingsContent(RankingsParityFixtures.board(), NOW, zone = ZONE) }

    private fun pushed(screen: String, fontScale: Float = 1f, content: @Composable () -> Unit) =
        ParityCheck.state(compose, screen, fontScale) {
            ForeverTabScaffold(selected = ForeverTab.Home, onSelect = {}, badges = TabBadges(raid = 1, activity = 3)) {
                Column {
                    FlowBackBar("Home", {})
                    Box(Modifier.fillMaxSize().verticalScroll(rememberScrollState())) { content() }
                }
            }
        }

    private companion object {
        val NOW: Instant = Instant.parse("2026-12-17T20:00:00Z")
        val ZONE: ZoneId = ZoneId.of("America/Chicago")
    }
}

/** The mock's board: Obnoxious Yell, 14th of 212 Fury Warriors on Onyxia, Normal ruleset. */
object RankingsParityFixtures {
    val encounters = listOf(
        RankedEncounter(1, "onyxia", "Onyxia", "Onyxia's Lair", 40, 38),
        RankedEncounter(2, "warden-kelthas", "Warden Kelthas", null, null, 4),
    )

    private val names = listOf(
        "Thorgrim Ashveil", "Valkyra", "Brannoch", "Kaelith Rowe", "Rhoswen", "Dunmar", "Tessiel Vane", "Ironvow",
        "Gaelwyn", "Hrafnild", "Morvain Holt", "Sellar", "Orsolya", "Obnoxious Yell", "Bladewick",
    )

    private val olympus = RankingGuild("Olympus XXVII", "us", "normal")
    private val now: Instant = Instant.parse("2026-12-17T20:00:00Z")
    private const val TOP = 888.0
    private const val STEP = 12.0
    private const val FIRST_LENGTH_MS = 161_000
    private const val SECOND_MS = 1_000
    private const val YOURS = 14

    private fun row(index: Int): RankingRow {
        val rank = index + 1
        val name = names[index]
        return RankingRow(
            rank = rank,
            player = RankingPlayer("us/normal/${name.lowercase().replace(' ', '-')}", name, "Warrior", "Fury"),
            guild = if (rank % 2 == 0) olympus else null,
            value = TOP - STEP * index,
            size = 40,
            foughtAt = now.minus(Duration.ofDays(rank.toLong())).toString(),
            durationMs = FIRST_LENGTH_MS + SECOND_MS * index,
            reportId = "r$rank",
            fightIndex = 3,
            state = "ranked",
            executionScore = 0.9,
        )
    }

    private val rows = List(names.size) { row(it) }

    val page = RankingsPage(
        rows = rows,
        total = 212,
        page = 1,
        perPage = 100,
        updatedAt = "2026-12-13T18:00:00Z",
        deduped = true,
        characterRank = CharacterRank(YOURS, rows[YOURS - 1]),
    )

    private val obnoxious = ViewerCharacter("us/normal/obnoxious-yell", "Obnoxious Yell", "us", "normal", "warrior", "Fury")

    val viewer = RankingsViewer(signedIn = true, current = obnoxious, accountKeys = setOf(obnoxious.key), history = true)

    val filters = RankingsFilters("onyxia", Measure.Damage, "normal", "", "", "", todayOnly = false, cohort = Cohort("warrior", "Fury"))

    fun board() = RankingsScreenState(
        BoardPhase.Loaded(Board.of(page, null)),
        filters,
        encounters.associate { it.slug to it.name },
        viewer,
        reportsOn = true,
    )
}
