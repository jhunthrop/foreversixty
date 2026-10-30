import Auth
import Foundation
import Models

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

    static func signInFailure(_ failure: SignInFailure) -> String {
        switch failure {
        case .battleNet: Copy.signinErrorBnet
        case .expired: Copy.signinErrorExpired
        case .offline: Copy.signinErrorOffline
        }
    }

    static func age(_ age: RelativeAge) -> String {
        switch age {
        case .now: Copy.accountAgeNow
        case let .minutes(n): Copy.accountAgeMinutes(n: n)
        case let .hours(n): Copy.accountAgeHours(n: n)
        case let .days(n): Copy.accountAgeDays(n: n)
        }
    }

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
