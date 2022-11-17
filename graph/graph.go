package main

import (
	"container/list"
	"fmt"
)

type Graph struct {
	Visited bool
	Name    string
	Friends []*Graph
}

func main() {
	var g1, g2, g3, g4, g5, g6, g7, g8 Graph

	g1 = Graph{
		Name:    "상인",
		Friends: []*Graph{&g2, &g3, &g4},
	}
	g2 = Graph{
		Name:    "은지",
		Friends: []*Graph{&g1},
	}
	g3 = Graph{
		Name:    "범기",
		Friends: []*Graph{&g1, &g5},
	}
	g4 = Graph{
		Name:    "혜진",
		Friends: []*Graph{&g1, &g6},
	}
	g5 = Graph{
		Name:    "영욱",
		Friends: []*Graph{&g2, &g7},
	}
	g6 = Graph{
		Name:    "윤희",
		Friends: []*Graph{&g4},
	}
	g7 = Graph{
		Name:    "상미",
		Friends: []*Graph{&g5, &g8},
	}
	g8 = Graph{
		Name:    "수현",
		Friends: []*Graph{&g7},
	}
	g1.Dfs()
}

func (g *Graph) Dfs() {
	// 방문할때 이름 출력
	queue := list.New()
	queue.PushBack(g)
	for queue.Len() > 0 {
		target := queue.Remove(queue.Front()).(*Graph)
		fmt.Println(target.Name)
		target.Visited = true
		for i := 0; i < len(target.Friends); i++ {
			if !target.Friends[i].Visited {
				queue.PushBack(target.Friends[i])
			}
		}
	}
}
