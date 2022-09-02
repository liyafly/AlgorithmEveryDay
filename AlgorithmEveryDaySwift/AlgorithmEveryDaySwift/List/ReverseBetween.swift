//
// Created by 李亚非 on 2022/9/2.
//

import Foundation

extension Solution {
    func reverseBetween(_ head: ListNode?, _ left: Int, _ right: Int) -> ListNode? {
        let d = ListNode(-1, head)
        var p: ListNode? = d
        for _ in 0..<(left - 1) {
            p = p?.next
        }
        let c = p?.next
        var n: ListNode? = nil
        for _ in left..<right {
            n = c?.next
            c?.next = n?.next
            n?.next = p?.next
            p?.next = n
        }
        return d.next
    }
    }
}
