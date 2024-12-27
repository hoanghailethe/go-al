package cmd

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 type ListNode struct {
	     Val int
	     Next *ListNode
 }

 func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    h1 := list1
    h2 := list2
   //var head *ListNode
   head := &ListNode{
		Val: 0,
		Next: &ListNode{},
	}
    running := head.Next
    for running!=nil {
        fmt.Println(running)
        
        res, stop := PickYourBestGirlFriend(h1,h2)   
		if stop {
			break
		}
		running=res
          
        running = running.Next     
    }
    return head.Next
}

func PickYourBestGirlFriend(list1 *ListNode, list2 *ListNode) (*ListNode, bool) {
	if list1 == nil {
		return list2, true
	}
	if list2 == nil {
		return list1, true
	}
    res := &ListNode{}
	if(list1.Val<=list2.Val){
        res = list1
		if list1.Next!=nil {list1= list1.Next}
	} else {
        res = list2
        if list2.Next!=nil {list2= list2.Next}
    }
	return res, false
}

// FIX C
func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
    dummy := &ListNode{}
    running := dummy

    h1 := list1
    h2 := list2

    for h1 != nil || h2 != nil {
        res, stop, updatedH1, updatedH2 := PickYourBestGirlFriend2(h1, h2)
        if stop {
            running.Next = res
            break
        }
        running.Next = res
        running = running.Next
        h1, h2 = updatedH1, updatedH2
    }
    return dummy.Next
}

func PickYourBestGirlFriend2(list1 *ListNode, list2 *ListNode) (*ListNode, bool, *ListNode, *ListNode) {
    if list1 == nil {
        return list2, true, nil, nil
    }
    if list2 == nil {
        return list1, true, nil, nil
    }

    if list1.Val <= list2.Val {
        return list1, false, list1.Next, list2
    } else {
        return list2, false, list1, list2.Next
    }
}


// why h1 dont change after each loop hear:

// [1,2,4]
// [1,3,4]

// &{0 <nil>}
// &{1 0xc0000160d0}
// &{2 0xc0000160e0}
// end turn
// &{2 0xc0000160e0}
// &{1 0xc0000160d0}
// &{2 0xc0000160e0}
// end turn
// &{2 0xc0000160e0}
// &{1 0xc0000160d0}
// &{2 0xc0000160e0}
// end turn

// The issue lies in how you're updating the list1 and list2 pointers inside your PickYourBestGirlFriend function. Specifically, when you write:

// go
// Copy code
// if list1.Next != nil {
//     list1 = list1.Next
// }
// You are updating the local variable list1, which is a copy of the pointer passed to PickYourBestGirlFriend. This does not modify the h1 pointer in the mergeTwoLists function. In Go, function arguments are passed by value, meaning that the list1 and list2 variables in PickYourBestGirlFriend are new local variables that hold copies of the original pointers.

// Why h1 Doesn't Change
// h1 is passed to PickYourBestGirlFriend as list1.
// Inside PickYourBestGirlFriend, the line list1 = list1.Next modifies only the local copy of the pointer.
// This does not affect the original h1 pointer in the calling function.
// How to Fix It
// To ensure h1 (or list1) is updated correctly, you can either:

// Update the pointers directly in the calling function: Use the result of PickYourBestGirlFriend to update h1 and h2 explicitly in mergeTwoLists.

// Pass pointers to pointers: Modify the signature of PickYourBestGirlFriend to take **ListNode so that the changes to the pointer are reflected outside the function.