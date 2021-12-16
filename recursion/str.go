package main

import (
	"container/list"
	"fmt"
	"strings"
)

type EventMarketing struct {
	// sku
}

func Process1(str []byte, index int, ans1 *list.List, path string) {
	if index == len(str) || len(path) >= 2 {
		// fmt.Printf("index=%d,str=%s\n", index, path)
		if len(path) >= 2 {
			ans1.PushBack(path)
		}
		// fmt.Println(ans1)
		return
	}
	no := path
	Process1(str, index+1, ans1, no)
	yes := path + string(str[index])
	Process1(str, index+1, ans1, yes)
}

//列出所有排列情况
func Process2(str []byte, index int, ans *list.List) {
	if index == len(str) {
		ans.PushBack(string(str))
	} else {
		for i := index; i < len(str); i++ {
			swap(str, index, i)
			Process2(str, index+1, ans)
			swap(str, i, index)
		}
	}
}

func swap(str []byte, i, j int) {
	tmp := str[i]
	str[i] = str[j]
	str[j] = tmp
}
func main1() {
	// fmt.Println(int('0'), int('2'), int(rune('3')), int(rune('|')))
	str := []byte("abcdef1")
	// path := ""
	// l := list.New()
	// Process1(str, 0, l, path)
	// for i := l.Front(); i != nil; i = i.Next() {
	// 	fmt.Println(i.Value)
	// }

	l2 := list.New()
	Process2(str, 0, l2)
	fmt.Println("len=", l2.Len())
	for i := l2.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	// input := []rune("12+(27-3)/3+(3+2)*15")
	// e := len(input)
	// b := AfaToBtree(input, 0, e)
	// PrintInOrder(b, 0, "H", 15)
	// val := Cal(b)
	// fmt.Println(val)

	// input1 := []rune("(Count(a)+Count(b)>=6)&&(Amount(c)+Amount(d)>=16000)")
	// e1 := len(input1)
	// b1 := AfaToBtree1(input1, 0, e1)
	// PrintInOrder(b1, 0, "H", 8)
	// val1 := Cal(b1)
	// fmt.Println(val1)

}

func main() {
	activities := []Activities{
		{Id: "1", Name: "买2送1", Desc: "", Expression: "Count(A)+C ount(B)+Count(C)>=3"},
		{Id: "2", Name: "买3送1", Desc: "", Expression: "Count(B)+Count(C)>=2"},
		{Id: "3", Name: "纸尿裤买2送1", Desc: "", Expression: "Count(A)+Count(B)>=3"},
	}

	for i, _ := range activities {
		a := []rune(activities[i].Expression)
		lenA := len(a)
		activities[i].ExpressionTree = AfaToBtree1(a, 0, lenA)
		PrintInOrder(activities[i].ExpressionTree, 0, "H", 8)
	}

	cart := &Cart{
		SkuList: []CartSku{
			{Id: "1", Name: "a", Num: 100, Amount: 108, GoodID: "A", GroupId: "A", InActivitiesNum: 2},
			{Id: "2", Name: "b", Num: 200, Amount: 356, GoodID: "B", GroupId: "B", InActivitiesNum: 1},
			{Id: "3", Name: "c", Num: 300, Amount: 108, GoodID: "C", GroupId: "C", InActivitiesNum: 1},
			{Id: "4", Name: "d", Num: 100, Amount: 254, GoodID: "D", GroupId: "D", InActivitiesNum: 1},
			// {Id: "5", Name: "e", Num: 100, Amount: 254, GoodID: "E", GroupId: "E", InActivitiesNum: 1},
			// {Id: "6", Name: "f", Num: 100, Amount: 254, GoodID: "D", GroupId: "A", InActivitiesNum: 1},
			// {Id: "7", Name: "d", Num: 100, Amount: 254, GoodID: "D", GroupId: "A", InActivitiesNum: 1},
			// {Id: "8", Name: "d", Num: 100, Amount: 254, GoodID: "D", GroupId: "A", InActivitiesNum: 1},
			// {Id: "9", Name: "d", Num: 100, Amount: 254, GoodID: "D", GroupId: "A", InActivitiesNum: 1},
			// {Id: "10", Name: "d", Num: 100, Amount: 254, GoodID: "D", GroupId: "A", InActivitiesNum: 1},
			// {Id: "11", Name: "d", Num: 100, Amount: 254, GoodID: "D", GroupId: "A", InActivitiesNum: 1},
			// {Id: "12", Name: "d", Num: 100, Amount: 254, GoodID: "D", GroupId: "A", InActivitiesNum: 1},
			// {Id: "13", Name: "d", Num: 100, Amount: 254, GoodID: "D", GroupId: "A", InActivitiesNum: 1},
			// {Id: "14", Name: "d", Num: 100, Amount: 254, GoodID: "D", GroupId: "A", InActivitiesNum: 1},
		},
		// GroupInfo: map[string]struct {
		// 	Num    int
		// 	Amount int
		// 	Index  int
		// }{
		// 	"a": {Num: 2, Amount: 108, Index: 0},
		// 	"b": {Num: 1, Amount: 356, Index: 1},
		// 	"c": {Num: 1, Amount: 108, Index: 2},
		// 	"d": {Num: 1, Amount: 254, Index: 3},
		// },
	}
	list := list.New()
	inAcitiveSkus := PromotionSelectSku{}
	visited := make(map[string]struct{}, 0)
	GetJionAcitivieEnum2(cart, 0, 1, list, visited, inAcitiveSkus, activities[0])
	for i := list.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}

type PromotionSelectSku []CartSku

func (promotion PromotionSelectSku) Add(sku CartSku) PromotionSelectSku {
	ans := append(promotion, sku)

	return ans
}

func (promotion PromotionSelectSku) GetSelectedSkuId() string {
	buf := strings.Builder{}
	for _, item := range promotion {
		buf.WriteString(item.Id)
		buf.WriteRune('-')
	}
	return buf.String()
}

//返回当前组内Sku数量
func (c PromotionSelectSku) GroupCount(groupID string) int {
	count := 0
	for _, item := range c {
		if strings.ToLower(item.GroupId) == groupID {
			count++
		}
	}
	return count
}

//返回当前组内Sku 总金额
func (c PromotionSelectSku) GroupAmount(groupID string) float64 {
	amount := 0.0
	for _, item := range c {
		if strings.ToLower(item.GroupId) == groupID {
			amount += item.Amount
		}
	}
	return amount
}

func (c PromotionSelectSku) CalculateFunc(args ...string) CalValue {
	if len(args) < 2 {
		return CalValue{
			boolVal:     false,
			CurrentType: BoolType,
		}
	}
	v := 0.0
	switch args[0] {
	case "count":
		v = float64(c.GroupCount(args[1]))
	case "amount":
		v = c.GroupAmount(args[1])
	}
	return CalValue{
		CurrentType: ValType,
		Val:         float64(v),
	}
}

//当前购物车 满足当前活动的几种组合
func GetJionAcitivieEnum(cart *Cart, index, position int, ans1 *list.List, inAcitiveskus PromotionSelectSku, activies ...Activities) {
	if len(activies) > 0 {
		isAdd := true
		for _, item := range activies {
			if !item.Calculate(inAcitiveskus.CalculateFunc) {
				// ans1.PushBack(inAcitiveskus)
				isAdd = false
			}
		}
		if isAdd {
			ans1.PushBack(inAcitiveskus)
		}

	}

	if index >= len(cart.SkuList) {
		return
	}
	var yes, no PromotionSelectSku
	if cart.SkuList[index].Num > position {
		position++
		no = inAcitiveskus
		yes = inAcitiveskus.Add(cart.SkuList[index])
		// GetJionAcitivieEnum(cart, index, position, ans1, yes, activies...)
	} else {
		no = inAcitiveskus
		yes = inAcitiveskus.Add(cart.SkuList[index])
		index++
		position = 1

	}
	GetJionAcitivieEnum(cart, index, position, ans1, yes, activies...)
	GetJionAcitivieEnum(cart, index, position, ans1, no, activies...)
}

//当前购物车 满足当前活动的几种组合
func GetJionAcitivieEnum1(cart *Cart, index, position int, ans1 *list.List, inAcitiveskus PromotionSelectSku, activie Activities) {
	if activie.Calculate(inAcitiveskus.CalculateFunc) {
		ans1.PushBack(inAcitiveskus)
		return
	}

	if index >= len(cart.SkuList) {
		return
	}
	var yes, no PromotionSelectSku
	no = inAcitiveskus
	GetJionAcitivieEnum1(cart, index+1, position, ans1, no, activie)
	yes = inAcitiveskus.Add(cart.SkuList[index])
	GetJionAcitivieEnum1(cart, index+1, position, ans1, yes, activie)
}

//当前购物车 满足当前活动的几种组合
func GetJionAcitivieEnum2(cart *Cart, index, position int, ans1 *list.List, visited map[string]struct{}, inAcitiveskus PromotionSelectSku, activie Activities) {
	// 判断当前sku 是否参入活动
	for index1 := index; index1 < len(cart.SkuList); index1++ {
		if activie.SkuUseInCurrentAcitivity(&cart.SkuList[index]) {
			index = index1
			// fmt.Println(index, cart.SkuList[index].GroupId)
			break
		}
	}
	// if !activie.SkuUseInCurrentAcitivity(&cart.SkuList[index]) {
	// 	// GetJionAcitivieEnum2(cart, index+1, position, ans1, visited, inAcitiveskus, activie)
	// 	// index++
	// 	return
	// }
	if _, ok := visited[inAcitiveskus.GetSelectedSkuId()]; ok {

		return
	}

	if activie.Calculate(inAcitiveskus.CalculateFunc) {

		visited[inAcitiveskus.GetSelectedSkuId()] = struct{}{}
		ans1.PushBack(inAcitiveskus)
		return
	}

	if index >= len(cart.SkuList) {
		return
	}
	var yes, no PromotionSelectSku
	if cart.SkuList[index].Num > position && position <= 3 {
		position++
		no = inAcitiveskus
		yes = inAcitiveskus.Add(cart.SkuList[index])
	} else {
		no = inAcitiveskus
		yes = inAcitiveskus.Add(cart.SkuList[index])
		index++
		position = 1

	}
	GetJionAcitivieEnum2(cart, index, position, ans1, visited, no, activie)

	GetJionAcitivieEnum2(cart, index, position, ans1, visited, yes, activie)

}

//当前购物车 满足了已选活动，还可以满足那些活动
