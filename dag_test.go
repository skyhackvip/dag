package main

import (
	"fmt"
	"github.com/eapache/queue"
	"testing"
	"time"
)

func TestDag(t *testing.T) {
	_, root := NewDAG()
	all := BFS(root)
	doTasks(all)
}

//广度遍历
func BFS(root *Vertex) []*Vertex {

	q := queue.New()
	q.Add(root)
	visited := make(map[string]*Vertex)
	all := make([]*Vertex, 0)

	for q.Length() > 0 {
		qSize := q.Length()
		for i := 0; i < qSize; i++ {
			//pop vertex
			currVert := q.Remove().(*Vertex)
			if _, ok := visited[currVert.Key]; ok {
				continue
			}
			visited[currVert.Key] = currVert
			all = append([]*Vertex{currVert}, all...)
			for _, val := range currVert.Children {
				if _, ok := visited[val.Key]; !ok {
					q.Add(val) //add child
				}
			}
		}
	}
	return all
}

//并发执行
func doTasks(vertexs []*Vertex) {
	for _, v := range vertexs {
		time.Sleep(5 * time.Second)
		fmt.Printf("do %v, result is %v \n", v.Key, v.Value)
	}
}
