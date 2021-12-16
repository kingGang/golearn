package main

import (
	"fmt"
	"strconv"
	"strings"
)

type BtreeNode struct {
	Data    CalValue
	yunsuan string
	lchild  *BtreeNode
	rchild  *BtreeNode
}

type ValueMod uint8

const (
	ValType  ValueMod = 1 //值类型
	BoolType ValueMod = 2 //bool类型
)

//
type CalValue struct {
	Val         float64
	boolVal     bool
	CurrentType ValueMod //当前值使用那个类型；1：float64,2:bool类型

}

//Cal 计算btree运算结果
func Cal(root *BtreeNode, processFunc func(args ...string) CalValue) CalValue {
	switch root.yunsuan {
	case "+", "-", "*", "/":
		if root.lchild.Data.CurrentType == ValType && root.rchild.Data.CurrentType == ValType {
			fVal := 0.0
			switch root.yunsuan {
			case "+":
				fVal = Cal(root.lchild, processFunc).Val + Cal(root.rchild, processFunc).Val
			case "-":
				fVal = Cal(root.lchild, processFunc).Val - Cal(root.rchild, processFunc).Val
			case "*":
				fVal = Cal(root.lchild, processFunc).Val * Cal(root.rchild, processFunc).Val
			case "/":
				fVal = Cal(root.lchild, processFunc).Val / Cal(root.rchild, processFunc).Val
			}
			root.Data = CalValue{
				CurrentType: ValType,
				Val:         fVal,
			}
			return root.Data
		}
	case ">", ">=", "<", "<=":
		if root.lchild.Data.CurrentType == ValType && root.rchild.Data.CurrentType == ValType {
			bVal := false
			switch root.yunsuan {
			case ">":
				bVal = Cal(root.lchild, processFunc).Val > Cal(root.rchild, processFunc).Val
			case ">=":
				bVal = Cal(root.lchild, processFunc).Val >= Cal(root.rchild, processFunc).Val
			case "<":
				bVal = Cal(root.lchild, processFunc).Val < Cal(root.rchild, processFunc).Val
			case "<=":
				bVal = Cal(root.lchild, processFunc).Val <= Cal(root.rchild, processFunc).Val
			case "==", "=":
				bVal = Cal(root.lchild, processFunc).Val == Cal(root.rchild, processFunc).Val

			}
			root.Data = CalValue{
				CurrentType: BoolType,
				boolVal:     bVal,
			}
			return root.Data
		}
	case "||", "&&":
		if root.lchild.Data.CurrentType == BoolType && root.rchild.Data.CurrentType == BoolType {
			bVal := false
			switch root.yunsuan {
			case "||":
				bVal = Cal(root.lchild, processFunc).boolVal || Cal(root.rchild, processFunc).boolVal
			case "&&":
				bVal = Cal(root.lchild, processFunc).boolVal && Cal(root.rchild, processFunc).boolVal

			}
			root.Data = CalValue{
				CurrentType: BoolType,
				boolVal:     bVal,
			}
			return root.Data
		}
	default:
		op := []rune(root.yunsuan)
		lenOp := len(op)
		if lenOp > 0 && op[0] >= '0' && op[0] <= '9' {
			return CalValue{
				CurrentType: ValType,
				Val:         root.Data.Val,
			}
		} else {
			var funcName, arg string
			p := strings.IndexRune(root.yunsuan, '(')
			funcName = root.yunsuan[0:p]
			arg = root.yunsuan[p : lenOp-1]
			return processFunc(funcName, arg)
		}
	}

	return root.Data
}

func CalFunc(head *BtreeNode) CalValue {
	op := []rune(head.yunsuan)
	var reply CalValue
	if len(op) > 0 && op[0] >= '0' && op[0] <= '9' {
		reply = CalValue{
			CurrentType: ValType,
			Val:         head.Data.Val,
		}
	} else {
		switch head.yunsuan {
		case "func1":
			reply = CalValue{
				CurrentType: ValType,
				Val:         head.Data.Val,
			}
		default:
			reply = CalValue{
				CurrentType: ValType,
				Val:         8000,
			}
		}
	}
	return reply

}

// func CountVirtualTypeGoods() CalCalValue {

// }

func AfaToBtree1(input []rune, s, e int) *BtreeNode {
	var local_r, flag, m_m_p, a_s_p, x, compare_p, logical_p, xishu int
	for i := s; i < e; i++ {
		if string(input[i]) == "+" || string(input[i]) == "-" ||
			string(input[i]) == "*" || string(input[i]) == "/" ||
			string(input[i]) == ">" || string(input[i]) == "<" || string(input[i]) == "=" || string(input[i]) == "|" ||
			string(input[i]) == "&" {
			x++
		}
	}
	//为0，则为函数表达式或者纯数字
	if x == 0 {
		var bn *BtreeNode
		if int(input[s]) >= int('0') && int(input[s]) <= int('9') { //如果第一个字母为数字，则认为该字符串为数字
			y := int(input[e-1]) - int('0')
			xishu = 1
			for i := e - 2; i >= s; i-- {
				xishu = 10 * xishu
				y = y + xishu*(int(input[i])-int('0'))
			}
			bn = &BtreeNode{
				Data: CalValue{
					Val:         float64(y),
					CurrentType: ValType,
				},
				yunsuan: strconv.Itoa(y),
				lchild:  nil,
				rchild:  nil,
			}
		} else { //否则字符串为字母
			y := ""
			for i := s; i < e; i++ {
				y = y + string(input[i])
			}
			bn = &BtreeNode{
				Data: CalValue{
					Val:         0,
					CurrentType: ValType,
				},
				yunsuan: strings.ToLower(y),
				lchild:  nil,
				rchild:  nil,
			}
		}

		return bn
	}
	for i := s; i < e; i++ {
		if input[i] == rune('(') {
			flag++
		} else if input[i] == rune(')') {
			flag--
		}
		if flag == 0 { //在括号外
			if input[i] == rune('+') || input[i] == rune('-') {
				//将括号外的最右侧的+或-号的位置记录下来
				a_s_p = i
			} else if input[i] == rune('*') || input[i] == rune('/') {
				//将括号外的最右侧的*或/号的位置记录下来
				m_m_p = i
			} else if input[i] == rune('|') || input[i] == rune('&') {
				//将括号外的最右侧的&&或||号的位置记录下来
				logical_p = i
			} else if input[i] == rune('>') || input[i] == rune('<') || input[i] == rune('=') {
				//将括号外的最右侧的>、<、>=、<=号的位置记录下来
				compare_p = i
			}
		}
	}
	if m_m_p == 0 && a_s_p == 0 && logical_p == 0 && compare_p == 0 { //去掉最外层的括号
		return AfaToBtree1(input, s+1, e-1)
	} else {
		valMod := ValType
		followCharOp := false
		//注意优先级 乘除》加减》比较》&&||
		if logical_p > 0 {
			local_r = logical_p
			followCharOp = true
			valMod = BoolType //该节点返回bool类型
		} else if compare_p > 0 {
			local_r = compare_p
			valMod = BoolType //该节点返回bool类型
			followCharOp = followCharIsOp(input[local_r-1])
		} else if a_s_p > 0 { //加减
			local_r = a_s_p
		} else if m_m_p > 0 { //乘除
			local_r = m_m_p
		}
		yunsuan := ""
		rightBegin := 0
		leftEnd := 0
		if followCharOp {
			yunsuan = string(input[local_r-1]) + string(input[local_r])
			rightBegin = local_r + 1
			leftEnd = local_r - 1
		} else {
			yunsuan = string(input[local_r])
			rightBegin = local_r + 1
			leftEnd = local_r
		}
		bn := &BtreeNode{
			Data: CalValue{
				Val:         0.0,
				CurrentType: valMod,
			},
			yunsuan: yunsuan,
			lchild:  AfaToBtree1(input, s, leftEnd),
			rchild:  AfaToBtree1(input, rightBegin, e),
		}
		return bn
	}
}

func followCharIsOp(s rune) bool {

	if s == rune('|') || s == rune('&') || s == rune('=') || s == rune('>') || s == rune('<') {
		return true
	}
	return false
}

//AfaToBtree 计算算术表达式
func AfaToBtree(input []rune, s, e int) *BtreeNode {
	var local_r, flag, m_m_p, a_s_p, x, xishu int
	for i := s; i < e; i++ {
		if string(input[i]) == "+" || string(input[i]) == "-" || string(input[i]) == "*" || string(input[i]) == "/" {
			x++
		}
	}
	//判断是否全为数字，如果是x取零
	if x == 0 {
		y := int(input[e-1]) - int('0')
		xishu = 1
		for i := e - 2; i >= s; i-- {
			xishu = 10 * xishu
			y = y + xishu*(int(input[i])-int('0'))
		}
		bn := &BtreeNode{
			Data: CalValue{
				Val:         float64(y),
				CurrentType: ValType,
			},
			yunsuan: strconv.Itoa(y),
			lchild:  nil,
			rchild:  nil,
		}
		return bn
	}
	for i := s; i < e; i++ {
		if input[i] == rune('(') {
			flag++
		} else if input[i] == rune(')') {
			flag--
		}
		if flag == 0 { //在括号外
			if input[i] == rune('*') || input[i] == rune('/') {
				//将括号外的最右侧的*或/号的位置记录下来
				m_m_p = i
			} else if input[i] == rune('+') || input[i] == rune('-') {
				//;//将括号外的最右侧的+或-号的位置记录下来
				a_s_p = i
			}
		}
	}
	if m_m_p == 0 && a_s_p == 0 { //去掉最外层的括号
		return AfaToBtree(input, s+1, e-1)
	} else {
		if a_s_p > 0 {
			local_r = a_s_p
		} else if m_m_p > 0 {
			local_r = m_m_p
		}
		bn := &BtreeNode{
			Data: CalValue{
				Val:         0.0,
				CurrentType: ValType,
			},
			yunsuan: string(input[local_r]),
			lchild:  AfaToBtree(input, s, local_r),
			rchild:  AfaToBtree(input, local_r+1, e),
		}
		return bn
	}
}

//PrintInOrder 逆时针90°输出树结构
// head 数的根节点
//height 当前节点高度
// to 当前节点分隔符
// length 每个节点打印宽度
func PrintInOrder(head *BtreeNode, height int, to string, length int) {
	if head == nil {
		return
	}
	PrintInOrder(head.rchild, height+1, "v", length)
	val := to + head.yunsuan + to
	lenM := len(val)
	lenL := (length - lenM) / 2
	lenR := length - lenM - lenL
	val = getSpace(lenL) + val + getSpace(lenR)
	fmt.Println(getSpace(height*length) + val)
	PrintInOrder(head.lchild, height+1, "^", length)

}

func getSpace(num int) string {
	sapce := " "
	var buf strings.Builder
	for i := 0; i < num; i++ {
		buf.WriteString(sapce)
	}
	return buf.String()
}
