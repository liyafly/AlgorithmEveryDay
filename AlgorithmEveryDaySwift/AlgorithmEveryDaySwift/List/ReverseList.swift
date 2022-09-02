//
// Created by 李亚非 on 2022/9/2.
//

import Foundation
// 反转列表 https://leetcode.cn/problems/reverse-linked-list/
extension Solution {
    func reverseList(_ head: ListNode?) -> ListNode? {
        var p: ListNode? = nil
        var c = head
        while c != nil {
            let t = c?.next
            c?.next = p
            p = c
            c = t
        }
        return  p
    }

    func reverseListR(_ head: ListNode?) -> ListNode? {
        if head == nil || head?.next == nil {
            return head
        }
        let p = reverseListR(head?.next)
        head?.next?.next = head
        head?.next = nil
        return p
    }


    func reverseListHeadInsert(_ head: ListNode?) -> ListNode? {
        let d = ListNode(-1, head)
        let p: ListNode? = d
        let c = p?.next
        var n: ListNode? = nil
        while c?.next != nil {
            n = c?.next
            c?.next = n?.next
            n?.next = p?.next
            p?.next = n
        }
        return d.next
    }
}
