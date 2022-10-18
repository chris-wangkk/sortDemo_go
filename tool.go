package main

import (
	"fmt"
	"math/rand"
)

type sortOper interface {
	Sort()
}

type sorter struct {
	Data []int
	//
	CntForComp int64
	CntForSwap int64
}

func (ins *sorter) Init() {
	for idx := 0; idx < 20; idx++ {
		ins.Data = append(ins.Data, rand.Intn(2000))
	}
}

func (ins *sorter) Swap(i, j int) {
	ins.Data[i], ins.Data[j] = ins.Data[j], ins.Data[i]
	ins.CntForSwap++
}

func (ins *sorter) Show() {
	for _, iData := range ins.Data {
		fmt.Printf("%d ", iData)
	}
	fmt.Printf("\n	替换次数=%v \n", ins.CntForSwap)
}

/*
	比较类排序->交换排序（具体执行方式是元素间两两比较）
	冒泡排序：
		比较相邻的元素。如果前面比后面大，就交换它们两个
		对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对，这样每次遍历在最后的元素应该会是最大的数
		针对所有的元素重复以上的步骤(每次都对少做最后一次，因为最后的元素已经排好序)
*/
type Bubble struct {
	sorter
}

func (ins *Bubble) Sort() {
	for i := 0; i < len(ins.Data); i++ {
		flag := false
		for j := 0; j < len(ins.Data)-i-1; j++ {
			ins.CntForComp++
			if ins.Data[j] > ins.Data[j+1] {
				ins.Swap(j, j+1)
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}

/*
	比较类排序->交换排序
	快速排序（重点）：
		分治法
		从数列中挑出一个元素作为基准(pivot),所有比基准值小的摆放在基准前面，所有比基准值大的摆在基准的后面
			此时该基准就处于数列的中间位置(partition)
		后续递归地在小于基准值元素的子数列和大于基准值元素的子数列再做类似处理
*/
type Quick struct {
	sorter
}

func (ins *Quick) partition(left, right int) int {
	var (
		pivot = ins.Data[left]
		i     = left + 1
		j     = right
	)
	for {
		for i <= j && ins.Data[i] <= pivot {
			ins.CntForComp++
			i++ //从右边扫描直至找到比pivot大的数值
		}
		for i <= j && ins.Data[j] >= pivot {
			ins.CntForComp++
			j-- //从左边扫描直至找到比pivot小的数值
		}
		if i >= j { //交汇了，说明正好扫描完，找到pivot对应的位置，退出
			break
		}
		ins.Swap(i, j)
	}
	ins.Data[left] = ins.Data[j]
	ins.Data[j] = pivot
	ins.CntForSwap++
	return j
}
func (ins *Quick) quickSort(left, right int) {
	if left < right {
		mid := ins.partition(left, right)
		ins.quickSort(left, mid-1)
		ins.quickSort(mid+1, right)
	}
}
func (ins *Quick) Sort() {
	ins.quickSort(0, len(ins.Data)-1)
}

/*
	比较类排序->插入排序
	简单插入排序（一般不会用）
		思路简单，通过构建有序序列，对于未排序数据，在已排序序列中从后向前扫描，找到相应位置并插入
		1.从第一个元素开始（该元素可以认为已经被排序）
		2.取出下一个元素，在已经排序的元素序列中从后向前扫描
		3.如果该元素大于新元素，将该元素移到下一位置
		4.重复步骤3，直到找到已排序的元素小于或者等于新元素的位置
		5.将新元素插入到该位置后
		重复上述步骤
*/
type SimplyInsert struct {
	sorter
}

func (ins *SimplyInsert) Sort() {
	for i := 0; i < len(ins.Data); i++ {
		tmp := ins.Data[i]
		for j := i - 1; j >= 0; j-- {
			ins.CntForComp++
			if ins.Data[j] > tmp {
				ins.Data[j+1] = ins.Data[j]
			} else {
				if j+1 != i {
					ins.Data[j+1] = tmp
				}
				break
			}
		}
	}
}

/*
	比较类排序->插入排序
	希尔排序（会有用）
		vs简单插入，会优先比较距离最远的元素(缩小增量排序)
		先让数组中任意间隔为 h 的元素有序，刚开始 h 的大小可以是 h = n / 2
		接着让 h = n / 4，让 h 一直缩小，当 h = 1 时，也就是此时数组中任意间隔为1的元素有序，此时的数组就是有序
		1.选择一个增量序列t1，t2，…，tk，其中ti>tj，tk=1；
		2.按增量序列个数k，对序列进行k 趟排序；
		3.每趟排序，根据对应的增量ti，将待排序列分割成若干长度为m 的子序列，分别对各子表进行直接插入排序
			这样排序下来能有个粗略的排序，后续排序触发的次数会少很多
		4.仅增量因子为1 时，整个序列作为一个表来处理（此时触发插入的次数就少很多，只是比较），表长度即为整个序列的长度
*/
type Shell struct {
	sorter
	interval int
}

func (ins *Shell) insert(h, loc int) {
	val := ins.Data[loc]
	k := loc - h
	for ; k >= 0 && val < ins.Data[k]; k -= h {
		ins.CntForSwap++
		ins.Data[k+h] = ins.Data[k]
	}
	ins.Data[k+h] = val
}
func (ins *Shell) Sort() {
	if 0 >= ins.interval {
		ins.interval = 3
	}
	for h := len(ins.Data) / ins.interval; h > 0; h /= ins.interval {
		for i := h; i < len(ins.Data); i++ {
			ins.insert(h, i)
		}
	}
}

/*
	比较类排序->选择排序
	简单选择排序（一般不会用）
		每次遍历找到最值，然后放到对应的位置
*/
type SimpleSelect struct {
	sorter
}

func (ins *SimpleSelect) Sort() {
	for i := 0; i < len(ins.Data); i++ {
		val := i
		for j := i + 1; j < len(ins.Data); j++ {
			if ins.Data[i] > ins.Data[j] {
				val = j
			}
			if i != val {
				ins.Swap(i, val)
			}
		}
	}
}

/*
	比较类排序->选择排序
	堆排序（重点）
		1.将初始待排序关键字序列(R1,R2….Rn)构建成大顶堆，此堆为初始的无序区
		2.将堆顶元素R[1]与最后一个元素R[n]交换，得到新的无序区(R1,R2,……Rn-1)和新的有序区(Rn),且满足R[1,2…n-1]<=R[n]
		3.由于交换后新的堆顶R[1]可能违反堆的性质，因此需要对当前无序区(R1,R2,……Rn-1)调整为新堆，然后再次将R[1]与无序区最后一个元素交换，得到新的无序区(R1,R2….Rn-2)和新的有序区(Rn-1,Rn)
		4.不断重复此过程直到有序区的元素个数为n-1
*/
type Heap struct {
	sorter
}

func (ins *Heap) heapity(i, len int) { //将[i,len]内的元素堆化
	var (
		lchild = i*2 + 1
		rchild = i*2 + 2
		max    = i
	)
	if lchild < len && ins.Data[lchild] > ins.Data[max] {
		max = lchild
	}
	if rchild < len && ins.Data[rchild] > ins.Data[max] {
		max = rchild
	}
	if i != max {
		ins.Swap(i, max)
		ins.heapity(max, len)
	}
}

func (ins *Heap) buildHeap() {
	for i := len(ins.Data)/2 - 1; i >= 0; i-- {
		ins.heapity(i, len(ins.Data))
	}
}

func (ins *Heap) Sort() {
	ins.buildHeap()
	for i := len(ins.Data) - 1; i > 0; i-- {
		ins.Swap(0, i)    //（将将堆顶元素与最后一个元素交换）这里是堆排序的一个缺陷：此时若尾部是一个小值，会反复触发交换操作
		ins.heapity(0, i) //重新建堆logn
	}
}

/*
	比较类排序->归并排序
*/
type Merge struct {
	sorter
	bRecur bool
}

func (ins *Merge) mergeProc(left, mid, right int) {
	var (
		len = right - left + 1
		i   = left
		j   = mid + 1
		tmp = make([]int, 0, len)
	)
	for i <= mid && j <= right {
		if ins.Data[i] <= ins.Data[j] {
			tmp = append(tmp, ins.Data[i])
			i++
		} else {
			tmp = append(tmp, ins.Data[j])
			j++
		}
	}
	for i <= mid {
		tmp = append(tmp, ins.Data[i])
		i++
	}
	for j <= right {
		tmp = append(tmp, ins.Data[j])
		j++
	}
	for idx, elem := range tmp {
		ins.Data[left+idx] = elem
	}
}

func (ins *Merge) mergeSortRecur(left, right int) {
	if left >= right {
		return
	}
	mid := (left + right) / 2
	ins.mergeSortRecur(left, mid)
	ins.mergeSortRecur(mid+1, right)
	ins.mergeProc(left, mid, right)
}

func (ins *Merge) mergeSortIteration() {
	for idx := 1; idx < len(ins.Data); idx *= 2 { //每次归并的间隔
		left := 0                      //最开始肯定从0起
		for left+idx < len(ins.Data) { //迭代式归并
			mid := left + idx - 1
			right := mid + idx
			if right >= len(ins.Data) {
				right = len(ins.Data) - 1 //是否还足够右边的
			}
			ins.mergeProc(left, mid, right)
			left = right + 1
		}
	}
}

func (ins *Merge) Sort() {
	if ins.bRecur {
		ins.mergeSortRecur(0, len(ins.Data)-1)
	} else {
		ins.mergeSortIteration()
	}
}

func main() {
	obj1 := Bubble{}
	obj1.Init()
	obj1.Sort()
	obj1.Show()
	//
	obj2 := Quick{}
	obj2.Init()
	obj2.Sort()
	obj2.Show()
	//
	obj3 := SimplyInsert{}
	obj3.Init()
	obj3.Sort()
	obj3.Show()
	//
	obj4 := Shell{}
	obj4.Init()
	obj4.Sort()
	obj4.Show()
	//
	obj5 := Heap{}
	obj5.Init()
	obj5.Sort()
	obj5.Show()
	//
	obj6 := Merge{bRecur: false}
	obj6.Init()
	obj6.Sort()
	obj6.Show()
}
