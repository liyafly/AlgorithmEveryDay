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
            var t = c?.next
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
        var p = reverseListR(head?.next)
        head?.next?.next = head
        head?.next = nil
        return p
    }

}
