package _4_tree

type BST struct {
	*BinaryTree

	//比较函数，-1: v < nodeV 0: v = nodeV 1: v > nodeV
	compareFunc func(v, nodeV interface{}) int
}

//二叉查找树搜索
func (this *BST) Find(val interface{}) *Node {
	p := this.root
	for p != nil {
		compareResult := this.compareFunc(val, p.data)
		if compareResult == 0 {
			return p
		} else if compareResult > 0 {
			p = p.right
		} else {
			p = p.left
		}
	}

	return nil
}

func (this *BST) Delete(val interface{}) bool {
	p := this.root

	//暂存节点
	var temp *Node = nil

	deleteLeft := false

	for nil != p {
		compareResult := this.compareFunc(val, p.data)
		if compareResult == 0 {
			break
		} else if compareResult > 0 {
			temp = p
			p = p.right
			deleteLeft = false
		} else {
			temp = p
			p = p.left
			deleteLeft = true
		}
	}

	if nil == p { //要删除的节点不存在
		return false
	} else if nil == p.left && nil == p.right { //删除的是一个叶子节点
		if nil != temp {
			if deleteLeft {
				temp.left = nil
			} else {
				temp.right = nil
			}
		} else { //删除的是根节点
			this.root = nil
		}
	} else if nil != p.right { //删除的是包含右孩子的节点，可能也包含左孩子的节点，比较复杂
		//找到p节点的右子树最小节点
		temp2 := p  	//缓存当前节点
		q := p.right //向右走一步
		fromRight := true
		for nil != q.left { //向左走到底
			temp2 = q
			q = q.left
			fromRight = false
		}
		if fromRight {
			temp2.right = nil
		} else {
			temp2.left = nil
		}
		q.left = p.left
		q.right = p.right
		if nil == temp2 { //根节点被删除
			this.root = q
		} else {
			if deleteLeft {
				temp2.left = q
			} else {
				temp2.right = q
			}
		}
	} else { //删除的是只有左孩子的节点
		if nil != temp {
			if deleteLeft {
				temp.left = p.left
			} else {
				temp.right = p.left
			}
		} else { //删除根节点
			this.root = p.left
		}
	}

	return true
}
