package main

import (
	"fmt"
	"github.com/eapache/queue"
	"sync"
	"testing"
	"time"
)

func TestDagNew(t *testing.T) {
	_, root := NewDAG()
	all := BFSNew(root)

	for _, layer := range all {
		fmt.Println("------------------")
		doTasksNew(layer)
	}
}

//广度遍历 返回双层结构
func BFSNew(root *Vertex) [][]*Vertex {

	q := queue.New()
	q.Add(root)
	visited := make(map[string]*Vertex)
	all := make([][]*Vertex, 0)

	for q.Length() > 0 {
		qSize := q.Length()
		tmp := make([]*Vertex, 0)
		for i := 0; i < qSize; i++ {
			//pop vertex
			currVert := q.Remove().(*Vertex)
			if _, ok := visited[currVert.Key]; ok {
				continue
			}
			visited[currVert.Key] = currVert
			tmp = append(tmp, currVert)

			//fmt.Println(level, currVert.Key, currVert.Value)

			for _, val := range currVert.Children {
				if _, ok := visited[val.Key]; !ok {
					q.Add(val) //add child
				}
			}
		}
		all = append([][]*Vertex{tmp}, all...)
	}
	return all
}

//并发执行
func doTasksNew(vertexs []*Vertex) {
	var wg sync.WaitGroup
	for _, v := range vertexs {
		wg.Add(1)
		go func(v *Vertex) {
			defer wg.Done()
			time.Sleep(5 * time.Second)
			fmt.Printf("do %v, result is %v \n", v.Key, v.Value)
		}(v) //notice
	}
	wg.Wait()
}
