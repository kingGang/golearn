package main

import(
	"fmt"
	"github.com/pkg/errors"
)
const(
	RED bool = true
	BLACK bool = false
)

type RBNode struct{
	value int64
	color bool
	left,right,parent *RBNode
}

//返回当前节点的爷爷节点
func (rb *RBNode)getGrandParent()*RBNode{
	if rb.parent!=nil{
		return rb.parent.parent
	}
	return nil
}

// 返回当前节点的兄弟节点
func (rb *RBNode)getSibling()*RBNode{
	if rb.parent!=nil{
		if rb.parent.left==rb{
			return rb.parent.right
		}else{
			return rb.parent.left
		}
	}
	return nil
}

func (rb *RBNode)getUncle()*RBNode{
	grandparent:=rb.getGrandParent()
	if grandparent==nil{
		return nil
	}
	if grandparent.left==rb.parent{
		return grandparent.right
	}else{
		return grandparent.left
	}
	return nil
}

func main(){
	fmt.Println("start ...")
}
const(
	LEFTROTATE bool = true
	RIGHTROTATE bool =false
)
type RBTree sturct{
	root *RBNode
}
//rotate ture左旋|false右旋
func (rbnode *RBNode)rotate(isRotateLeft bool)(*RBNode,error){
	var root *RBNode
	if rbnode==nil{
		return root,nil
	}
	if !isRotateLeft&&rbnode.left==nil{
		return root,errors.New("右旋左节点不能为空")
	}else if isRotateLeft&&rbnode.right==nil{
		return root,errors.New("左旋右节点不能为空")
	}

	parent:=rbnode.parent
	var isleft bool
	//判断当前节点是否为左节点
	if parent!=nil{
		isleft=parent.left==rbnode
	}
	if isRotateLeft{//左旋
		grandson:=rbnode.right.left
		rbnode.right.left=rbnode
		rbnode.parent=rbnode.right
		rbnode.right=grandson
	}else{//右旋
		grandson:=rbnode.left.right
		rbnode.left.right=rbnode
		rbnode.parent=rbnode.left
		rbnode.left=grandson
	}
	if parent==nil{
		rbnode.parent.parent=nil
		root=rbnode.parent
	}else{
		if left{
			parent.left=rbnode.parent
		}else{
			parent.right=rbnode.parent
		}
		rbnode.parent.parent=parent
	}
	return root,nil
}