package gg.foreversixty.app.firstrun

import androidx.compose.ui.test.junit4.v2.createComposeRule
import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.model.AccountCharacter
import gg.foreversixty.app.model.CharacterGuild
import gg.foreversixty.app.testing.PARITY_PIXEL_9
import gg.foreversixty.app.testing.ParityCheck
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.annotation.Config
import org.robolectric.annotation.GraphicsMode

/** The first run (design/mocks/ws_app_flows.py F8): Pick your main and the guild member's ask. */
@RunWith(AndroidJUnit4::class)
@GraphicsMode(GraphicsMode.Mode.NATIVE)
@Config(qualifiers = PARITY_PIXEL_9)
class FirstRunParityTest {
    @get:Rule val compose = createComposeRule()

    @Test fun pickMain() = pickMain(1f)

    @Test fun pickMainAt2() = pickMain(2f)

    @Test fun permission() = permission("first-permission", "Olympus XXVII")

    @Test fun permissionAt2() = permission("first-permission", "Olympus XXVII", 2f)

    @Test fun permissionLong() = permission("first-permission-long", "Ranchers of Thunderbluff")

    private fun pickMain(fontScale: Float) = ParityCheck.state(compose, "first-pick-main", fontScale = fontScale) {
        PickMainScreen(FirstRunFixtures.characters, FirstRunFixtures.characters.first().key, saving = false, failed = false, onSelect = {}, onContinue = {})
    }

    private fun permission(screen: String, guild: String, fontScale: Float = 1f) = ParityCheck.state(compose, screen, fontScale = fontScale) {
        FirstPermissionScreen(guild, onTurnOn = {}, onNotNow = {})
    }
}

/** The mock's account: Thandrel and five alts, highest level first. */
object FirstRunFixtures {
    private val olympus = CharacterGuild(id = 2, name = "Olympus XXVII", rank = "Member", verified = true)

    private fun character(name: String, className: String, spec: String, level: Int, guild: CharacterGuild? = null) = AccountCharacter(
        key = "us/normal/${name.lowercase()}",
        region = "us",
        ruleset = "normal",
        name = name,
        source = "bnet",
        className = className,
        level = level,
        spec = spec,
        guild = guild,
    )

    val characters = listOf(
        character("Thandrel", "warrior", "fury", 60, olympus),
        character("Lightvow", "paladin", "holy", 60, olympus),
        character("Kaelor", "mage", "frost", 47),
        character("Felsnap", "hunter", "beast mastery", 34),
        character("Brewmist", "shaman", "enhancement", 22),
        character("Ashmantle", "warlock", "affliction", 18),
    )
}
