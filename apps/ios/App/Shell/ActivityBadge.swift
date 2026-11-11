import Foundation

/// The Activity tab's number (flows spec 3): the unread count from `GET /v1/me/activity`, read at
/// shell start and on every foreground so it shows on every tab before Activity opens (round 5, i1;
/// Android's `refreshUnread`). Signed out, or a failed read, shows none.
enum ActivityBadgeCount {
    static func load(signedIn: Bool, service: any ActivityService) async -> Int {
        guard signedIn, let page = try? await service.page(group: nil, before: nil) else { return 0 }
        return page.unread
    }
}
