package gg.foreversixty.app.gear.flows

import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.runtime.Composable
import androidx.compose.runtime.CompositionLocalProvider
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.test.junit4.v2.createComposeRule
import androidx.compose.ui.unit.dp
import androidx.test.core.app.ApplicationProvider
import androidx.test.ext.junit.runners.AndroidJUnit4
import gg.foreversixty.app.bis.BisItemRef
import gg.foreversixty.app.bis.BisList
import gg.foreversixty.app.bis.BisPick
import gg.foreversixty.app.bis.BisRow
import gg.foreversixty.app.bis.BisRunnerUp
import gg.foreversixty.app.bis.BisUnit
import gg.foreversixty.app.designsystem.FsColor
import gg.foreversixty.app.designsystem.parityContent
import gg.foreversixty.app.flows.FlowLead
import gg.foreversixty.app.gear.ItemSource
import gg.foreversixty.app.gear.RunPick
import gg.foreversixty.app.items.ItemLook
import gg.foreversixty.app.menu.LocalOpenMenu
import gg.foreversixty.app.menu.ProfileBar
import gg.foreversixty.app.model.ItemHolder
import gg.foreversixty.app.model.ItemNeed
import gg.foreversixty.app.model.ItemNeeds
import gg.foreversixty.app.navigation.ForeverTab
import gg.foreversixty.app.navigation.ForeverTabScaffold
import gg.foreversixty.app.navigation.TabBadges
import gg.foreversixty.app.testing.PARITY_PIXEL_9
import gg.foreversixty.app.testing.ParityCheck
import org.junit.Assert.assertEquals
import org.junit.Rule
import org.junit.Test
import org.junit.runner.RunWith
import org.robolectric.annotation.Config
import org.robolectric.annotation.GraphicsMode

/** Gear (design/mocks/ws_app_flows.py `gear`, `item_sheet`) inside the tab shell. */
@RunWith(AndroidJUnit4::class)
@GraphicsMode(GraphicsMode.Mode.NATIVE)
@Config(qualifiers = PARITY_PIXEL_9)
class GearParityTest {
    @get:Rule val compose = createComposeRule()

    @Test fun gear() = ParityCheck.state(compose, "gear") { Tab { GearFlowScreen(GearFixtures.content, BUILD, GearActions()) } }

    @Test fun gearAt2() = ParityCheck.state(compose, "gear", fontScale = 2f) { Tab { GearFlowScreen(GearFixtures.content, BUILD, GearActions()) } }

    @Test fun gearOut() = ParityCheck.state(compose, "gear-out") { Tab { GearSignedOut(GearFixtures.classes) {} } }

    @Test fun itemSheet() = ParityCheck.state(compose, "item-sheet") {
        Column(Modifier.fillMaxSize().parityContent()) {
            Spacer(Modifier.fillMaxWidth().height(DIM))
            Column(Modifier.fillMaxWidth().background(FsColor.raised, RoundedCornerShape(topStart = 16.dp, topEnd = 16.dp))) {
                Box(Modifier.fillMaxWidth().padding(top = 9.dp, bottom = 9.dp), Alignment.Center) {
                    Box(Modifier.size(36.dp, 5.dp).background(FsColor.line, RoundedCornerShape(3.dp)))
                }
                GearSheetBody(GearFixtures.sheet, BUILD)
            }
        }
    }

    @Test fun placesGroupInstancesThenProfessionsThenReputations() {
        val composer = GearComposer(ApplicationProvider.getApplicationContext<android.content.Context>().resources)
        val rows = composer.where(GearFixtures.picks, GearFixtures.list)
        assertEquals(listOf("Blackrock Spire", "Blacksmithing", "Stormpike Guard"), rows.map { it.title })
        assertEquals("Blackcrow · Shadow Hunter Vosh'gajin, 8%", rows[0].line)
        assertEquals("Lionheart Helm, Greenhammer, Ironwood Blade · ask a blacksmith", rows[1].line)
        assertEquals("Don Julio's Band, needs revered", rows[2].line)
    }

    @Test fun theSummaryCountsWornSlots() {
        val composer = GearComposer(ApplicationProvider.getApplicationContext<android.content.Context>().resources)
        val content = composer.content(GearFixtures.list, "alliance", emptyList(), emptyList())
        assertEquals("Fury Warrior · 60 raid · Alliance", content.listTitle)
        assertEquals("1 of 3 slots on the list", content.onList)
    }

    @Test fun theSheetNamesWhoNeedsItAndWhoHasIt() {
        val res = ApplicationProvider.getApplicationContext<android.content.Context>().resources
        val row = GearSlotRow("head", ItemLook(2, "Lionheart Helm", 4), "Head", "+14.5", worn = false, row = GearFixtures.list.rows.first())
        val sheet = GearSheetComposer(res, BUILD).content(row, "DPS", null, emptyList(), GearFixtures.needs)
        assertEquals(listOf("Protection Warrior · needs it, the pick on their list", "Enhancement Shaman · has it"), sheet.guild?.map { it.line })
        assertEquals(null, sheet.guildNote)
    }

    @Composable
    private fun Tab(content: @Composable () -> Unit) {
        ForeverTabScaffold(selected = ForeverTab.Gear, onSelect = {}, badges = TabBadges(raid = 1, activity = 3)) {
            Column {
                CompositionLocalProvider(LocalOpenMenu provides {}) { ProfileBar(null) }
                content()
            }
        }
    }

    private companion object {
        const val BUILD = "1.60.1.70790"
        val DIM = 120.dp
    }
}

/** The mock's Thandrel: the Fury Warrior raid list, his upgrades and where they come from. */
object GearFixtures {
    private data class Doll(val slot: String, val label: String, val name: String, val icon: String, val quality: Int, val source: String, val gain: String?)

    private val doll = listOf(
        Doll("head", "Head", "Lionheart Helm", "inv_helmet_36", 4, "Blacksmithing", "+14.5"),
        Doll("neck", "Neck", "Skibi's Pendant", "inv_jewelry_amulet_07", 3, "World drop", null),
        Doll("shoulder", "Shoulder", "Theramore Pauldrons", "inv_shoulder_06", 4, "Theramore Expeditionary Force", null),
        Doll("back", "Back", "Howler's Furs", "inv_misc_cape_10", 3, "Pix Xizzix, vendor", "+2.4"),
        Doll("chest", "Chest", "Timbermaw Tunic", "inv_chest_fur", 4, "Leatherworking", "+5.3"),
        Doll("wrist", "Wrist", "Forest Stalker's Bracers", "inv_bracer_12", 4, "Silverwing Sentinels", "+0.1"),
        Doll("hands", "Hands", "Theramore Gauntlets", "inv_gauntlets_29", 4, "Theramore Expeditionary Force", "+1.0"),
        Doll("waist", "Waist", "Valiant's Waistguard", "inv_belt_33", 4, "Pix Xizzix, vendor", null),
        Doll("legs", "Legs", "Sentinel's Plate Legguards", "inv_pants_plate_20", 4, "Illiyana Moonblaze, vendor", "+4.7"),
        Doll("feet", "Feet", "Scalegut Treaders", "inv_boots_chain_11", 3, "Leatherworking", "+0.9"),
        Doll("finger_1", "Finger", "Don Julio's Band", "inv_jewelry_ring_35", 4, "Stormpike Guard", "+6.8"),
        Doll("finger_2", "Finger", "Cutthroat's Signet", "inv_jewelry_ring_01", 3, "Pix Xizzix, vendor", "+0.8"),
        Doll("trinket_1", "Trinket", "Hand of Justice", "inv_jewelry_talisman_01", 3, "Blackrock Depths: Emperor Dagran Thaurissan", null),
        Doll("trinket_2", "Trinket", "Molten Heart of the Mountain", "inv_misc_gem_ruby_01", 3, "Enchanting", "+3.2"),
        Doll("main_hand", "Main hand", "Ironwood Blade", "inv_sword_07", 4, "Blacksmithing", "+5.9"),
        Doll("off_hand", "Off hand", "Greenhammer", "inv_mace_1h_raidhyjalc60_d_01", 4, "Blacksmithing", "+9.7"),
        Doll("ranged", "Ranged", "Blackcrow", "inv_weapon_crossbow_04", 3, "Blackrock Spire: Shadow Hunter Vosh'gajin, 8%", "+1.9"),
    )

    private fun look(i: Int, d: Doll) = ItemLook(1000 + i, d.name, d.quality, d.icon)

    val content = GearContent(
        listTitle = "Fury Warrior · 60 raid · Alliance",
        onList = "4 of 17 slots on the list",
        where = listOf(
            GearWhereRow("Blackrock Spire", "Blackcrow · Shadow Hunter Vosh'gajin, 8%", ItemLook(1, "Blackcrow", 3, "inv_weapon_crossbow_04")),
            GearWhereRow("Blacksmithing", "Lionheart Helm, Greenhammer, Ironwood Blade · ask a blacksmith", ItemLook(2, "Lionheart Helm", 4, "inv_helmet_36")),
            GearWhereRow("Stormpike Guard", "Don Julio's Band, needs revered", ItemLook(3, "Don Julio's Band", 4, "inv_jewelry_ring_35")),
        ),
        slots = doll.mapIndexed { i, d ->
            GearSlotRow(d.slot, look(i, d), "${d.label} · ${d.source}", d.gain, worn = d.gain == null, row = BisRow(d.slot, d.label, null, null))
        },
        unitWord = "DPS",
    )

    val classes = listOf("warrior", "paladin", "hunter", "rogue", "priest", "shaman", "mage", "warlock", "druid")
        .map { it to it.replaceFirstChar(Char::titlecase) }

    private const val BUILD = "1.60.1.70790"
    private val deathbringer = ItemLook(17068, "Deathbringer", 4, "inv_axe_09")

    val sheet = GearSheetContent(
        look = deathbringer,
        slot = "Main hand",
        gain = "+12.6",
        unitWord = "DPS",
        tooltip = listOf(
            "Binds when picked up" to false,
            "One-Hand · Axe" to false,
            "114 - 213 Damage · Speed 2.90" to false,
            "(56.4 damage per second)" to false,
            "Chance on hit: Sends a shadowy bolt at the enemy causing 125 Shadow damage." to true,
            "Requires Level 60" to false,
        ),
        where = listOf(GearSheetLine("Onyxia, Onyxia's Lair", "8% drop chance · raid", FlowLead.Item(deathbringer, BUILD))),
        runnerUp = GearSheetLine(
            "Vis'kag the Bloodletter", "Onyxia · 8% · −1.8 DPS",
            FlowLead.Item(ItemLook(17075, "Vis'kag the Bloodletter", 4, "inv_sword_18"), BUILD), titleColor = null,
        ),
        guild = listOf(
            GearSheetLine("Grunvald", "Protection Warrior · needs it, the pick on their list", FlowLead.Crest("warrior")),
            GearSheetLine("Torvak", "Enhancement Shaman · has it", FlowLead.Crest("shaman")),
        ),
        guildNote = null,
    )

    private fun pick(id: Int, name: String, worn: Boolean = false) =
        BisPick(BisItemRef(id, name, null), "", "", isUnmodelled = false, isWorn = worn, runnersUp = emptyList<BisRunnerUp>(), source = "")

    val list = BisList(
        setDpsLine = null,
        rows = listOf(
            BisRow("head", "Head", pick(2, "Lionheart Helm"), null),
            BisRow("ranged", "Ranged", pick(1, "Blackcrow"), null),
            BisRow("finger_1", "Finger", pick(3, "Don Julio's Band", worn = true), null),
        ),
        unit = BisUnit.Dps,
        specLabel = "Fury Warrior",
        rangeLabel = "60 raid",
    )

    private val smith = ItemSource("crafted", "Blacksmithing", null, profession = "blacksmithing")
    val picks = listOf(
        RunPick("head", 2, "Lionheart Helm", 14.5, smith),
        RunPick("off_hand", 5, "Greenhammer", 9.7, smith),
        RunPick("finger_1", 3, "Don Julio's Band", 6.8, ItemSource("rep", "Stormpike Guard", null, standing = "Revered")),
        RunPick("main_hand", 4, "Ironwood Blade", 5.9, smith),
        RunPick("ranged", 1, "Blackcrow", 1.9, ItemSource("dungeon", "Blackrock Spire", 8.0, boss = "Shadow Hunter Vosh'gajin")),
        RunPick("legs", 6, "Sentinel's Plate Legguards", 4.7, ItemSource("vendor", "Illiyana Moonblaze", null)),
    )

    val needs = ItemNeeds(
        17068,
        listOf(ItemNeed("us/pvp/grunvald", "Grunvald", "Warrior", "Protection", 60, 60, "main_hand", true)),
        listOf(ItemHolder("us/pvp/torvak", "Torvak", "Shaman", "Enhancement")),
        0,
    )
}
