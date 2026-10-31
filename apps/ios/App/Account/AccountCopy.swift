import Auth
import Foundation
import Models
import Networking

/// The words for F1's states, all from the copy catalog (`Copy`). Dates follow the spec's
/// meta lines: `30 Oct` within the year, `27 Oct 2027` otherwise.
enum AccountCopy {
    static func notice(_ notice: SignedOutNotice) -> String {
        switch notice {
        case let .signInFailed(failure): signInFailure(failure)
        case .sessionEnded: Copy.signinSignedOut
        case let .accountDeleted(completeBy): Copy.accountDeleteDone(date: longDate(completeBy))
        }
    }

    /// Red is for failures only (spec §5.0): a completed deletion reads neutral.
    static func isFailure(_ notice: SignedOutNotice) -> Bool {
        switch notice {
        case .signInFailed, .sessionEnded: true
        case .accountDeleted: false
        }
    }

    /// `US · PvP`: Forever has no realms, so a character is its region, ruleset and name.
    static func place(_ character: AccountCharacter) -> String {
        Copy.accountCharacterWhere(region: character.region.uppercased(), ruleset: ruleset(character.ruleset))
    }

    /// The site's ruleset labels (web/src/lib/characters.ts); an unknown one is echoed.
    static func ruleset(_ id: String) -> String {
        switch id {
        case "normal": Copy.accountRulesetNormal
        case "pvp": Copy.accountRulesetPvp
        case "rp": Copy.accountRulesetRp
        case "hardcore": Copy.accountRulesetHardcore
        default: id
        }
    }

    static func consequence(_ consequence: DeletionConsequence) -> String {
        switch consequence {
        case .premiumPlan: Copy.accountDeleteLosePremium
        case .founderPrice: Copy.accountDeleteLoseFounder
        case let .guildPlan(guild): Copy.accountDeleteLoseGuildPlan(guild: guild)
        case let .seat(guild): Copy.accountDeleteLoseSeat(guild: guild)
        }
    }

    /// `Version 0.1.0 (1) · data 1.60.1.70658`, without the data part before the config answers.
    static func version(_ identity: ClientIdentity, data: String?) -> String {
        guard let data else { return Copy.accountVersionNoData(version: identity.version, build: identity.build) }
        return Copy.youVersion(version: identity.version, build: identity.build, data: data)
    }

    static func signInFailure(_ failure: SignInFailure) -> String {
        switch failure {
        case .battleNet: Copy.signinErrorBnet
        case .expired: Copy.signinErrorExpired
        case .offline: Copy.signinErrorOffline
        }
    }

    static func age(_ age: RelativeAge) -> String { age.phrase }

    static func deviceName(_ row: DeviceRow) -> String {
        row.isThisDevice ? Copy.devicesThis : row.name
    }

    static func deviceMeta(_ row: DeviceRow) -> String? {
        switch row.meta {
        case let .signedIn(date): Copy.devicesSignedIn(date: shortDate(date))
        case let .lastUsed(age): Copy.devicesLastUsed(age: Self.age(age))
        case .none: nil
        }
    }

    /// `Orc Fury Warrior · Level 58` (switcher rows).
    static func switcherDetail(_ character: AccountCharacter) -> String {
        let kind = [character.race, character.spec, character.className]
            .compactMap { $0.map(words) }
            .joined(separator: " ")
        guard let level = character.level else { return kind }
        return Copy.accountSwitcherDetail(kind: kind, n: level)
    }

    /// `Fury · 58` (the selector's short line); the class when no spec is saved.
    static func selectorDetail(_ character: AccountCharacter) -> String? {
        guard let level = character.level else { return nil }
        let spec = (character.spec ?? character.className).map(words) ?? ""
        return Copy.accountSelectorDetail(spec: spec, n: level)
    }

    /// The VoiceOver line: `Character, Lightningbolltz, Elemental Shaman, level 60`.
    static func selectorLabel(_ character: AccountCharacter) -> String {
        Copy.selectorA11y(
            name: character.name,
            spec: character.spec.map(words) ?? "",
            className: character.className.map(words) ?? "",
            n: character.level ?? 0
        )
    }

    /// The API's lower-case words (`night elf`, `warrior`) as display words.
    static func words(_ raw: String) -> String {
        raw.split(separator: " ").map { $0.prefix(1).uppercased() + $0.dropFirst() }.joined(separator: " ")
    }

    static func shortDate(_ date: Date) -> String {
        date.formatted(.dateTime.day().month(.abbreviated).locale(locale))
    }

    static func longDate(_ date: Date) -> String {
        date.formatted(.dateTime.day().month(.abbreviated).year().locale(locale))
    }

    /// The catalog is English (en.json); dates match it.
    private static let locale = Locale(identifier: "en_GB")
}
