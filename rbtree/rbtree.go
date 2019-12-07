package main

import(
	"fmt"
	"github.com/pkg/errors"
	"log"
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
	rbtree:=&RBTree{
		root:&RBNode{color:BLACK},
	}
	rbtree.insertNode(rbtree.root,10)
	rbtree.insertNode(rbtree.root,9)
	rbtree.insertNode(rbtree.root,8)
	rbtree.insertNode(rbtree.root,6)
	rbtree.insertNode(rbtree.root,7)
	log.Println("------ ", rbtree.root.value)
    log.Println("----", rbtree.root.left.value, "---", rbtree.root.right.value)
    log.Println("--", rbtree.root.left.left.value, "-", rbtree.root.left.right.value)
}
const(
	LEFTROTATE bool = true
	RIGHTROTATE bool =false
)

type RBTree struct{
	root *RBNode
}

func (rbtree *RBTree)rotateLeft(node *RBNode){
	if tmproot,err:=node.rotate(LEFTROTATE);err!=nil{
		if tmproot!=nil{
			rbtree.root=tmproot
		}
	}else{
		log.Printf(err.Error())
	}
}

func (rbtree *RBTree)rotateRight(node *RBNode){
	if tmproot,err:=node.rotate(RIGHTROTATE);err!=nil{
		if tmproot!=nil{
			rbtree.root=tmproot
		}
	}else{
		log.Printf(err.Error())
	}
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
		if isleft{
			parent.left=rbnode.parent
		}else{
			parent.right=rbnode.parent
		}
		rbnode.parent.parent=parent
	}
	return root,nil
}

func (rbtree *RBTree)insertCheck(node *RBNode){
	if node.parent==nil{ //为根节点
		rbtree.root=node
		rbtree.root.color=BLACK
		return
	}
	if node.parent.color==RED{
		if node.getUncle()!=nil&&node.getUncle().color==RED{
			// 父节点及叔父节点都是红色，则转为黑色
			node.parent.color=BLACK
			node.getUncle().color=BLACK
			//祖父节点转为红色
			node.getGrandParent().color=RED
			//递归处理
			rbtree.insertCheck(node.getGrandParent())
		}else{ //父亲节点为红色，父亲兄弟节点为黑色
			isleft:=node == node.parent.left
			isparentleft:=node.parent==node.getGrandParent().left
			if isleft&&isparentleft{ //当前节点和父节点都为左节点 
				node.parent.color=BLACK
				node.getGrandParent().color=RED
				rbtree.rotateRight(node.getGrandParent())
			}else if !isleft&&isparentleft{ //right red|parent left red
				rbtree.rotateLeft(node.parent)//左转
				rbtree.rotateRight(node.parent)//右转
				node.color=BLACK
				// node.left.color=RED
				node.right.color=RED
			}else if isleft&&!isparentleft{ //left red|parent right red
				rbtree.rotateRight(node.parent)
				rbtree.rotateLeft(node.parent)
				node.color=BLACK
				node.right.color=RED
				node.left.color=RED
			}else if !isleft&&!isparentleft{//red right|parent red right
				node.parent.color=BLACK
				node.getGrandParent().color=RED
				rbtree.rotateLeft(node.getGrandParent())
			}
		}
	}
}

func (rbtree *RBTree)insertNode(pnode *RBNode,data int64){
	if pnode.value>=data{
		if pnode.left!=nil{
			rbtree.insertNode(pnode.left,data)
		}else{
			tmpnode:=&RBNode{}
			tmpnode.value=data
			tmpnode.parent=pnode
			pnode.left=tmpnode
			rbtree.insertCheck(tmpnode)
		}
	}else{
		if pnode.right!=nil{
			rbtree.insertNode(pnode.right,data)
		}else{
			tmpnode:=&RBNode{}
			tmpnode.value=data
			tmpnode.parent=pnode
			pnode.right=tmpnode
			rbtree.insertCheck(tmpnode)
		}
	}
}