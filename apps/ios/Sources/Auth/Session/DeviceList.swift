import Foundation
import Models

/// One row of You › Devices (spec F1.6), already ordered and dated; the view supplies the words.
public struct DeviceRow: Equatable, Sendable, Identifiable {
    public enum Meta: Equatable, Sendable {
        /// This device: `Signed in {date}`.
        case signedIn(Date)
        /// Another device: `Last used {age}`.
        case lastUsed(RelativeAge)
        /// A device the API sent no date for.
        case none
    }

    public let id: String
    /// The API's name; the view shows `This iPhone` instead when `isThisDevice`.
    public let name: String
    public let isThisDevice: Bool
    public let meta: Meta

    /// Every device but this one carries Revoke.
    public var canRevoke: Bool { !isThisDevice }
}

public enum DeviceList {
    /// This device first, then the others by last use, newest first.
    public static func rows(_ devices: [Device], thisDevice: String?, now: Date) -> [DeviceRow] {
        let this = devices.filter { $0.id == thisDevice }
        let others = devices.filter { $0.id != thisDevice }.sorted { lastUse($0) > lastUse($1) }
        return this.map { row($0, isThis: true, now: now) } + others.map { row($0, isThis: false, now: now) }
    }

    private static func lastUse(_ device: Device) -> Date {
        device.lastSeenAt ?? device.createdAt ?? .distantPast
    }

    private static func row(_ device: Device, isThis: Bool, now: Date) -> DeviceRow {
        let meta: DeviceRow.Meta = if isThis {
            device.createdAt.map(DeviceRow.Meta.signedIn) ?? .none
        } else {
            (device.lastSeenAt ?? device.createdAt).map { .lastUsed(RelativeAge(from: $0, to: now)) } ?? .none
        }
        return DeviceRow(id: device.id, name: device.name, isThisDevice: isThis, meta: meta)
    }
}

/// How long ago, in the one unit the meta lines use (`2 h ago`).
public enum RelativeAge: Equatable, Sendable {
    case now
    case minutes(Int)
    case hours(Int)
    case days(Int)

    private static let minute: TimeInterval = 60
    private static let hour: TimeInterval = 3600
    private static let day: TimeInterval = 86400

    public init(from date: Date, to now: Date) {
        let seconds = max(0, now.timeIntervalSince(date))
        switch seconds {
        case ..<Self.minute: self = .now
        case ..<Self.hour: self = .minutes(Int(seconds / Self.minute))
        case ..<Self.day: self = .hours(Int(seconds / Self.hour))
        default: self = .days(Int(seconds / Self.day))
        }
    }
}
