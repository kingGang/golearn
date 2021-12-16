package main

import "strings"

type CartSku struct {
	Id              string
	Name            string
	Num             int     //总数量
	Amount          float64 //总金额
	GoodID          string  //商品ID
	GroupId         string  //虚拟分类
	InActivitiesNum int     //参入活动的数量
}

type Cart struct {
	SkuList   []CartSku
	GroupInfo map[string]struct {
		Num    int //组内商品数量
		Amount int //组内商品总金额
		Index  int //sku信息下标
	}
}

type Activities struct {
	Id                  string
	Name                string
	Desc                string
	Expression          string
	ExpressionTree      *BtreeNode
	MutexActivitieIds   []string //互斥的活动
	MutuallyDissolvable []string //互溶的活动 //活动某人都是累加的
}

//返回当前组内Sku数量
func (c *Cart) GroupCount(groupID string) int {
	return c.GroupInfo[groupID].Num
}

//返回当前组内Sku 总金额
func (c *Cart) GroupAmount(groupID string) int {
	return c.GroupInfo[groupID].Amount
}

func (c *Cart) CalculateFunc(args ...string) CalValue {
	if len(args) < 2 {
		return CalValue{
			boolVal:     false,
			CurrentType: BoolType,
		}
	}
	v := 0
	switch args[0] {
	case "count":
		v = c.GroupCount(args[1])
	case "amount":
		v = c.GroupAmount(args[1])
	}
	return CalValue{
		CurrentType: ValType,
		Val:         float64(v),
	}
}

func (a *Activities) Calculate(processFun func(args ...string) CalValue) bool {
	v := OperationExpression(a.ExpressionTree, processFun)
	if v.CurrentType == BoolType {
		return v.boolVal
	}
	//如果不是bool类型的值，则返回false
	return false
}

// 当前SKU是否使用
func (a *Activities) SkuUseInCurrentAcitivity(sku *CartSku) bool {
	return strings.Contains(a.Expression, "("+sku.GroupId+")")
}

//Cal 计算btree运算结果
func OperationExpression(root *BtreeNode, processFunc func(args ...string) CalValue) CalValue {
	switch root.yunsuan {
	case "+", "-", "*", "/":
		if root.lchild.Data.CurrentType == ValType && root.rchild.Data.CurrentType == ValType {
			fVal := 0.0
			switch root.yunsuan {
			case "+":
				fVal = OperationExpression(root.lchild, processFunc).Val + OperationExpression(root.rchild, processFunc).Val
			case "-":
				fVal = OperationExpression(root.lchild, processFunc).Val - OperationExpression(root.rchild, processFunc).Val
			case "*":
				fVal = OperationExpression(root.lchild, processFunc).Val * OperationExpression(root.rchild, processFunc).Val
			case "/":
				fVal = OperationExpression(root.lchild, processFunc).Val / OperationExpression(root.rchild, processFunc).Val
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
				bVal = OperationExpression(root.lchild, processFunc).Val > OperationExpression(root.rchild, processFunc).Val
			case ">=":
				bVal = OperationExpression(root.lchild, processFunc).Val >= OperationExpression(root.rchild, processFunc).Val
			case "<":
				bVal = OperationExpression(root.lchild, processFunc).Val < OperationExpression(root.rchild, processFunc).Val
			case "<=":
				bVal = OperationExpression(root.lchild, processFunc).Val <= OperationExpression(root.rchild, processFunc).Val
			case "==", "=":
				bVal = OperationExpression(root.lchild, processFunc).Val == OperationExpression(root.rchild, processFunc).Val

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
				bVal = OperationExpression(root.lchild, processFunc).boolVal || OperationExpression(root.rchild, processFunc).boolVal
			case "&&":
				bVal = OperationExpression(root.lchild, processFunc).boolVal && OperationExpression(root.rchild, processFunc).boolVal

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
			arg = root.yunsuan[p+1 : lenOp-1]
			return processFunc(funcName, arg)
		}
	}

	return root.Data
}
