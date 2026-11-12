package gg.foreversixty.app.navigation

import android.content.Context
import androidx.compose.ui.test.SemanticsNodeInteraction
import androidx.compose.ui.test.assertIsDisplayed
import androidx.compose.ui.test.hasText
import androidx.compose.ui.test.junit4.v2.createComposeRule
import androidx.compose.ui.test.onNodeWithText
import androidx.compose.ui.test.performClick
import androidx.test.core.app.ApplicationProvider
import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.R
import gg.foreversixty.app.bis.BisModel
import gg.foreversixty.app.bis.BisTestData
import gg.foreversixty.app.bis.FakeBisSource
import gg.foreversixty.app.bis.MemoryPreferences
import gg.foreversixty.app.designsystem.ForeverTheme
import gg.foreversixty.app.flows.SELECTABLE
import gg.foreversixty.app.home.HomeDependencies
import gg.foreversixty.app.home.MemoryHomeCache
import gg.foreversixty.app.items.ItemCatalog
import gg.foreversixty.app.model.mainCharacter
import gg.foreversixty.app.network.AccountApi
import gg.foreversixty.app.network.ApiClient
import gg.foreversixty.app.network.ApiConfig
import gg.foreversixty.app.network.ApiMode
import gg.foreversixty.app.network.ApiResult
import gg.foreversixty.app.network.AppConfigRepository
import gg.foreversixty.app.network.CachedFile
import gg.foreversixty.app.network.ClasspathFixtureSource
import gg.foreversixty.app.network.Feature
import gg.foreversixty.app.network.FixtureEngine
import gg.foreversixty.app.network.FixtureSource
import gg.foreversixty.app.network.ForeverHttpClient
import gg.foreversixty.app.network.InMemoryAppConfigCache
import gg.foreversixty.app.network.NoToken
import gg.foreversixty.app.network.StaticDataClient
import gg.foreversixty.app.network.StaticDataStore
import gg.foreversixty.app.testing.PARITY_PIXEL_9
import kotlinx.coroutines.runBlocking
import org.junit.Assert.assertTrue
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.annotation.Config
import org.robolectric.annotation.GraphicsMode

/**
 * A fresh install has no cached config, so the shell first composes with every switch off (the
 * bundled defaults) and every tab a placeholder. The config lands a moment later: the tabs must
 * switch to their screens then, with no restart (the Play internal track's first-launch report).
 */
@RunWith(AndroidJUnit4::class)
@GraphicsMode(GraphicsMode.Mode.NATIVE)
@Config(qualifiers = PARITY_PIXEL_9)
class ShellSwitchesTest {
    @get:Rule val compose = createComposeRule()

    private val context: Context = ApplicationProvider.getApplicationContext()
    private val classpath = ClasspathFixtureSource(requireNotNull(javaClass.classLoader))

    /** The contract fixtures, with a config answer that turns Home and Gear on (the fixture's own turns only raids on). */
    private val source = FixtureSource { name -> if (name == CONFIG_FIXTURE) CONFIG_HOME_AND_GEAR_ON else classpath.read(name) }
    private val apiConfig = ApiConfig(baseUrl = "https://api.example.test", mode = ApiMode.Fixture)
    private val client = ApiClient(ForeverHttpClient.create(apiConfig, NoToken, FixtureEngine.create(source)))

    /** An empty cache: the fresh install. */
    private val appConfig = AppConfigRepository(client, InMemoryAppConfigCache())

    private val catalog = ItemCatalog(StaticDataStore(StaticDataClient(FixtureEngine.create(classpath))))
    private val home = HomeDependencies(AccountApi(client), client, catalog, MemoryHomeCache(), isOn = { appConfig.isEnabled(Feature.CharacterHome) })
    private val bis = BisModel(
        FakeBisSource(ApiResult.Success(CachedFile(BisTestData.bis(), BisTestData.NOW, isFromCache = false), ""), BisTestData.me().mainCharacter),
        BisTestData.catalog(),
        MemoryPreferences(),
        context.resources,
    )

    /** As AppGraph.shellFeatures: a feature is in the shell only while its switch is on. */
    private fun features() = ShellFeatures(home = home, bis = bis.takeIf { appConfig.isEnabled(Feature.Bis) })

    @Test
    fun `a fresh install's tabs switch to their screens when the first config lands`() {
        compose.setContent { ForeverTheme { ForeverAppShell(appConfig.config) { features() } } }
        placeholder(ForeverTab.Home).assertIsDisplayed()
        tab(ForeverTab.Gear).performClick()
        placeholder(ForeverTab.Gear).assertIsDisplayed()

        assertTrue(runBlocking { appConfig.refresh() } is ApiResult.Success)
        compose.waitForIdle()

        placeholder(ForeverTab.Gear).assertDoesNotExist()
        // Signed out, Gear opens on the classes (spec 2026-11-08 4.3).
        compose.onNodeWithText(text(R.string.gear_out_title)).assertIsDisplayed()
        tab(ForeverTab.Home).performClick()
        placeholder(ForeverTab.Home).assertDoesNotExist()
    }

    @Test
    fun `a switch the config leaves off keeps its placeholder`() {
        compose.setContent { ForeverTheme { ForeverAppShell(appConfig.config) { features() } } }
        assertTrue(runBlocking { appConfig.refresh() } is ApiResult.Success)
        compose.waitForIdle()

        tab(ForeverTab.Guild).performClick()
        placeholder(ForeverTab.Guild).assertIsDisplayed()
    }

    private fun tab(tab: ForeverTab): SemanticsNodeInteraction = compose.onNode(hasText(text(tab.label)) and SELECTABLE)

    private fun placeholder(tab: ForeverTab): SemanticsNodeInteraction =
        compose.onNodeWithText(context.getString(R.string.shell_arrivesNext, text(tab.label)))

    private fun text(id: Int): String = context.getString(id)

    private companion object {
        const val CONFIG_FIXTURE = "app.config.json"
        val CONFIG_HOME_AND_GEAR_ON = """
            {"ok":true,"data":{"web_checkout_regions":[],"min_app_version":{"ios":"1.0.0","android":"1.0.0"},
            "active_build":"${BisTestData.BUILD}","features":{"character_home":true,"bis":true}},"error":null,"request_id":"shell"}
        """.trimIndent()
    }
}
