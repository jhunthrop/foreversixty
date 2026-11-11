package gg.foreversixty.app.menu

import androidx.compose.foundation.layout.width
import androidx.compose.ui.Modifier
import androidx.compose.ui.graphics.asImageBitmap
import androidx.compose.ui.test.junit4.v2.createComposeRule
import androidx.compose.ui.unit.dp
import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.model.AccountCharacter
import gg.foreversixty.app.model.ApiJson
import gg.foreversixty.app.model.CharacterGuild
import gg.foreversixty.app.model.Me
import gg.foreversixty.app.push.NotificationsLine
import gg.foreversixty.app.model.NotificationSettings
import gg.foreversixty.app.raids.RaidFixture
import gg.foreversixty.app.raids.resources
import gg.foreversixty.app.testing.PARITY_PIXEL_9_TALL
import gg.foreversixty.app.testing.ParityCheck
import kotlinx.serialization.json.jsonObject
import org.junit.Assert.assertEquals
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.annotation.Config
import org.robolectric.annotation.GraphicsMode
import java.time.Instant

/** The profile menu's mocked states (design/mocks/ws_app_flows.py `menu`) against their regions. */
@RunWith(AndroidJUnit4::class)
@GraphicsMode(GraphicsMode.Mode.NATIVE)
@Config(qualifiers = PARITY_PIXEL_9_TALL)
class MenuParityTest {
    @get:Rule val compose = createComposeRule()

    @Test fun menu() = state("menu", MenuFixtures.all)

    @Test fun menuAt1_5() = state("menu", MenuFixtures.all, fontScale = 1.5f)

    @Test fun menuAt2() = state("menu", MenuFixtures.all, fontScale = 2f)

    @Test fun menuOne() = state("menu-one", listOf(MenuFixtures.thandrel))

    @Test fun menuLow() = state("menu-low", MenuFixtures.all, lowOpen = true)

    @Test fun menuOut() = ParityCheck.state(compose, "menu-out") { Menu(null) }

    @Test fun menuOutAt2() = ParityCheck.state(compose, "menu-out", fontScale = 2f) { Menu(null) }

    @Test fun theModelFoldsLowLevelsAndSortsByLevel() {
        val model = MenuFixtures.model(MenuFixtures.all)
        assertEquals(listOf("Lightvow", "Kaelor", "Felsnap", "Brewmist", "Ashmantle"), model.switchTo.map { it.name })
        assertEquals(listOf("Pockets", "Bankalt"), model.lowLevel.map { it.name })
        assertEquals(true, model.isMain)
    }

    private fun state(screen: String, characters: List<AccountCharacter>, fontScale: Float = 1f, lowOpen: Boolean = false) =
        ParityCheck.state(compose, screen, fontScale) { Menu(MenuFixtures.model(characters), lowOpen) }

    @androidx.compose.runtime.Composable
    private fun Menu(model: MenuModel?, lowOpen: Boolean = false) {
        androidx.compose.runtime.CompositionLocalProvider(LocalPortraitPainter provides MenuFixtures.repoPortraits) {
            ProfileMenu(model, MenuFixtures.lines, MenuActions(), Modifier.width(DRAWER_WIDTH), lowLevelOpen = lowOpen)
        }
    }

    private companion object {
        /** 85% of 412. */
        val DRAWER_WIDTH = 350.dp
    }
}

/** The mock's account: Thandrel (the main) and the alts of `ws_app_flows.py` CHARACTERS and LOW, avatars and factions as there. */
object MenuFixtures {
    /** The parity clock (design/parity.md "Native"). */
    val NOW: Instant = Instant.parse("2026-11-06T18:00:00Z")

    private val guild = CharacterGuild(id = 2, name = "Olympus XXVII", rank = "Raider", verified = true)

    private fun character(
        name: String,
        className: String,
        spec: String?,
        level: Int,
        guilded: Boolean = false,
        avatar: Boolean = true,
        faction: String = "alliance",
        itemLevel: Int? = null,
    ) = AccountCharacter(
        key = "us/normal/${name.lowercase()}",
        region = "us",
        ruleset = "normal",
        name = name,
        source = "bnet",
        className = className,
        level = level,
        faction = faction,
        spec = spec,
        avatarUrl = if (avatar) "https://render.worldofwarcraft.com/us/character/forever/${name.lowercase()}-avatar.jpg" else null,
        guild = if (guilded) guild else null,
        itemLevel = itemLevel,
    )

    val thandrel = character("Thandrel", "warrior", "fury", 60, guilded = true)
    val all = listOf(
        thandrel,
        character("Lightvow", "paladin", "holy", 60, guilded = true, itemLevel = 61),
        character("Kaelor", "mage", "frost", 47, faction = "horde", itemLevel = 44),
        character("Felsnap", "hunter", "beast mastery", 34, itemLevel = 31),
        character("Brewmist", "shaman", "enhancement", 22, faction = "horde", itemLevel = 19),
        character("Ashmantle", "warlock", "affliction", 18, avatar = false, itemLevel = 15),
        character("Pockets", "rogue", null, 6, avatar = false, itemLevel = 5),
        character("Bankalt", "mage", null, 1, avatar = false),
    )

    /** The mock's portraits (design/mocks/data/avatars/<name>.jpg) for the fixtures' `<name>-avatar.jpg` URLs. */
    val repoPortraits = PortraitPainterSource { url ->
        androidx.compose.runtime.remember(url) {
            val name = url.substringAfterLast('/').removeSuffix("-avatar.jpg")
            java.io.File(gg.foreversixty.app.home.HomeFixtures.repoRoot, "design/mocks/data/avatars/$name.jpg").takeIf { it.exists() }
                ?.let { android.graphics.BitmapFactory.decodeFile(it.path) }
                ?.let { androidx.compose.ui.graphics.painter.BitmapPainter(it.asImageBitmap()) }
        }
    }

    val lines = MenuLines(
        // The drawer's own builder over the account fixture (round 7 row 4), every channel on, so the
        // golden and the app cannot disagree.
        notifications = NotificationsLine.account(RaidFixture.load(NotificationSettings.serializer(), "app.notifications")).let { account ->
            NotificationsLine.line(resources(), account.keys.toList(), account, account.keys.toList())
        },
        plan = "Premium · Founder price",
        devices = 3,
        newestDevice = "Pixel 9",
        offlineData = "26.6 MB",
        version = "Version 0.1.0 (7) · data 1.60.1.70790",
        year = 2026,
    )

    private fun me(characters: List<AccountCharacter>): Me {
        val raw = requireNotNull(javaClass.classLoader?.getResource("fixtures/app/me.json")) { "no me.json" }.readText()
        val data = requireNotNull(ApiJson.parseToJsonElement(raw).jsonObject["data"])
        return ApiJson.decodeFromJsonElement(Me.serializer(), data)
            .copy(characters = characters, mainCharacterKey = thandrel.key, bnetImportedAt = "2026-11-06T16:00:00Z")
    }

    fun model(characters: List<AccountCharacter>): MenuModel = MenuModel.of(me(characters), thandrel, NOW)
}
