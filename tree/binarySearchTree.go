package main

import (
	"fmt"
)

type SearchTree struct {
	//节点元素
	K int

	Used bool

	Left  *SearchTree

	Right *SearchTree
}

type HashSet struct {
	Set map[int]bool
}

func NewHashSet() *HashSet {
	return &HashSet{make(map[int]bool)}
}

func (set *HashSet) Add(i int) bool {
	_, found := set.Set[i]
	set.Set[i] = true
	return !found //False if it existed already
}

func (set *HashSet) Get(i int) bool {
	_, found := set.Set[i]
	return found //true if it existed already
}

func (set *HashSet) Remove(i int) {
	delete(set.Set, i)
}

func (t *SearchTree) Add (k int) {
	if t.Used == false {
		t.K = k
		t.Used = true
		return
	}

	if k == t.K {
		return
	}

	if k < t.K {
		if t.Left == nil {
			t.Left = new(SearchTree)
		}
		t.Left.Add(k)
	} else {
		if t.Right == nil {
			t.Right = new(SearchTree)
		}
		t.Right.Add(k)
	}
	return
}

func (t *SearchTree) Remove(k int, parent *SearchTree) {
	if t.Used == false {
		return
	}

	if k < t.K {
		if t.Left != nil {
			t.Left.Remove(k, t)			
		}
	} else {
		if t.Right != nil {
			t.Right.Remove(k, t)			
		}
	}

	if t.K == k {
		if t.Left != nil && t.Right != nil {
			//根据二叉树的中序遍历， 需要找到"右子树"的最小节点
			fmt.Println(t.K)
			rightMinK := t.FindMin(parent)

			t.Remove(rightMinK, parent)

			t.K = rightMinK
		} else {
			if t.Left != nil {
				parent.Left = t.Left
			} else {
				parent.Right = t.Right
			}
		}
	}
}

func (t *SearchTree) FindMin(parent *SearchTree) int {
	if t.Used == false {
		return parent.K
	}

	if t.Right == nil {
		if (t.Left == nil) {
			return t.K
		} else {
			return t.Left.FindMin(t)			
		}
	} else {
		return t.Right.FindMin(t)
	}
}

func (t *SearchTree) SearchRange(min, max int, hashSet *HashSet)  {
	if t.Used == false {
		return
	}

	//遍历左子树
	if min < t.K {
		if t.Left != nil {
			t.Left.SearchRange(min, max, hashSet)
		}
	}

	if min <= t.K && max >= t.K {
		hashSet.Add(t.K)
	}

	//遍历右子树
	if min > t.K || max > t.K {
		if t.Right != nil {		
			t.Right.SearchRange(min, max, hashSet)
		}
	}
}

func main () {
	intArr := []int{11,23,4,4,5,64,32,33,0,1,2,3,4,5,6,7}
	t := new(SearchTree)
	h := NewHashSet()

	for _, v := range intArr {
		t.Add(v)
	}

	t.Remove(4, t)

	t.SearchRange(1, 11, h)

	fmt.Println(h)
}