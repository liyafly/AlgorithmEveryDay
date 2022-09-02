//
// Created by 李亚非 on 2022/9/2.
//

import Foundation

open class ListNode {
    public var val: Int
    public var next: ListNode?
    public init() {
        val = 0
        next = nil
    }
    public init(_ val: Int) {
        self.val = val
        next = nil
    }
    public init(_ val: Int, _ next: ListNode?) {
        self.val = val
        self.next = next
    }
}