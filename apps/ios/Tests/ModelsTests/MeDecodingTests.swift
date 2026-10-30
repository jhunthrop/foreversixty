import Foundation
import Models
import Networking
import TestSupport
import Testing

/// The models must decode lane C's contract fixtures as they are. These tests
/// pin shapes, never catalogue counts or values (contracts are floors).
@Suite("Me decoding")
struct MeDecodingTests {
    private func decodeMe(_ json: Data) throws -> Me {
        let envelope = try WireFormat.decoder().decode(Envelope<Me>.self, from: json)
        return try #require(envelope.data)
    }

    @Test("decodes the /v1/me contract fixture")
    func decodesFixture() throws {
        let me = try decodeMe(FixtureLocation.data("me.json"))
        #expect(me.user.id > 0)
        #expect(me.characters.allSatisfy { !$0.key.isEmpty && !$0.name.isEmpty })
    }

    @Test("every /v1/me success fixture decodes", arguments: ["me.json", "me.premium.json", "me.empty.json", "me.patch.json"])
    func decodesEveryMeFixture(file: String) throws {
        let me = try decodeMe(FixtureLocation.data(file))
        #expect(me.user.id > 0)
    }

    @Test("the premium fixture carries its billing")
    func premiumBilling() throws {
        let me = try decodeMe(FixtureLocation.data("me.premium.json"))
        #expect(me.entitlements.billing != nil)
    }

    @Test("the empty fixture has no characters and no main")
    func emptyAccount() throws {
        let me = try decodeMe(FixtureLocation.data("me.empty.json"))
        #expect(me.characters.isEmpty)
        #expect(me.mainCharacter == nil)
    }

    @Test("main character prefers main_character_key, else the first")
    func mainCharacter() throws {
        let me = try decodeMe(FixtureLocation.data("me.json"))
        if let key = me.mainCharacterKey {
            #expect(me.mainCharacter?.key == key)
        } else {
            #expect(me.mainCharacter == me.characters.first)
        }
    }

    @Test("null slices and absent optionals decode as empty and nil")
    func minimalShape() throws {
        let json = """
        {"ok":true,"request_id":"r","error":null,"data":{
          "user":{"id":1,"battletag":null,"email":"a@b.test","role":"user","anonymize":true},
          "characters":null,"guilds":null,
          "entitlements":{"server_sims":true,"retention":false,"multi_compare":false,"history":false,
            "notifications":false,"officer_views":false,"roster_check":false,"log_coach":false,
            "supporter_mark":false,"billing":{"plan":"premium","source":"grant","status":"active",
            "current_period_end":null,"cancel_at_period_end":false,"price_cents":600},"personal_portal":false},
          "finished_sims":[]}}
        """
        let me = try decodeMe(Data(json.utf8))
        #expect(me.characters.isEmpty)
        #expect(me.guilds.isEmpty)
        #expect(me.mainCharacter == nil)
        #expect(me.bnetImportedAt == nil)
        #expect(me.entitlements.serverSims)
        #expect(me.entitlements.billing?.source == "grant")
    }

    @Test("a character card with only the required fields decodes")
    func sparseCharacter() throws {
        let json = #"{"key":"us/normal/x","region":"us","ruleset":"normal","name":"X","source":"export"}"#
        let card = try WireFormat.decoder().decode(CharacterCard.self, from: Data(json.utf8))
        #expect(card.level == nil)
        #expect(card.build == nil)
        #expect(card.guild == nil)
        #expect(card.id == "us/normal/x")
    }

    @Test("a full character card decodes build, talents and guild")
    func fullCharacter() throws {
        let json = """
        {"key":"us/normal/x","region":"us","ruleset":"normal","name":"X","class":"mage","level":60,
         "item_level":62,"avatar_url":"https://render.example.test/a.jpg","source":"bnet","spec":"frost",
         "build":{"source":"addon","captured_at":"2026-11-04T21:03:00.5Z","gear":{"head":1},
           "talents":{"trees":["1","2","3"],"points":[1,2,3]},"median_sync_gap_sec":null,"sync_error":null},
         "guild":{"id":2,"name":"G","rank":"Officer","verified":false,"attendance":{"present":1,"nights":2}}}
        """
        let card = try WireFormat.decoder().decode(CharacterCard.self, from: Data(json.utf8))
        #expect(card.itemLevel == 62)
        #expect(card.avatarUrl?.host == "render.example.test")
        #expect(card.build?.talents?.points == [1, 2, 3])
        #expect(card.build?.gear?["head"] == 1)
        #expect(card.guild?.attendance?.present == 1)
        #expect(card.guild?.attendance?.nights == 2)
    }
}

/// `GET /v1/characters/{key}`: lane C's character card fixtures.
@Suite("Character profile decoding")
struct CharacterProfileDecodingTests {
    private func decode(_ file: String) throws -> CharacterProfile {
        let envelope = try WireFormat.decoder().decode(Envelope<CharacterProfile>.self, from: FixtureLocation.data(file))
        return try #require(envelope.data)
    }

    @Test("a rated character decodes its score and gear panel")
    func rated() throws {
        let profile = try decode("character.card.json")
        #expect(!profile.character.key.isEmpty)
        #expect(profile.score != nil)
        #expect(profile.components != nil)
        #expect(profile.gear.allSatisfy { !$0.slot.isEmpty })
    }

    @Test("a character with no rated fight has every score field nil")
    func unrated() throws {
        let profile = try decode("character.card.empty.json")
        #expect(profile.score == nil)
        #expect(profile.tier == nil)
        #expect(profile.components == nil)
        #expect(profile.gearGapTotal == nil)
    }
}
