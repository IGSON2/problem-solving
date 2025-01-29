# [ 베스트 앨범 ](https://school.programmers.co.kr/learn/courses/30/lessons/42579?language=go)
0번부터 n번까지 트랙에 대한 장르명과 플레이 수가 배열로 주어질 때  
각 장르별 최대 2곡의 노래를 총 장르의 플레이수 -> 장르 내 플레이수로 정렬하여 추려내는 문제

플레이 수 기준 우선순위 큐를 활용하여 해결하려 했지만  
데이터를 POP하여 제거해 나가는 게 오히려 로직을 복잡하게 만들 것 같아  
문제 취지에 맞는 해싱(Map)만 사용하여 해결하였음


## 우선순위 큐를 사용하여 문제를 해결한 예시
```go
package main

import (
    "sort"
)

type Genre struct {
    id    int
    name  string
    count int
}

func NewGenre(id int, name string, count int) Genre {
    return Genre{
        id:    id,
        name:  name,
        count: count,
    }
}

type GenreHeap []Genre

func (h GenreHeap) Len() int                  { return len(h) }
func (h GenreHeap) Less(i, j int) bool        { return h[i].count > h[j].count }
func (h GenreHeap) Swap(i, j int)             { h[i], h[j] = h[j], h[i] }
func (h *GenreHeap) Push(element interface{}) { *h = append(*h, element.(Genre)) }
func (h GenreHeap) Pop(i int) Genre           { return h[i] }

type Ranking struct {
    hash map[string]int
    rank []string
}

func NewRanking(hash map[string]int) Ranking {
    ranking := Ranking{}
    ranking.hash = hash
    for key := range hash {
        ranking.rank = append(ranking.rank, key)
    }
    return ranking
}
func (r Ranking) Len() int           { return len(r.hash) }
func (r Ranking) Swap(i, j int)      { r.rank[i], r.rank[j] = r.rank[j], r.rank[i] }
func (r Ranking) Less(i, j int) bool { return r.hash[r.rank[i]] > r.hash[r.rank[j]] }

func solution(genres []string, plays []int) []int {
    var result []int
    hash := make(map[string]*GenreHeap)
    rank_hash := make(map[string]int)

    for i, genre := range genres {
        rank_hash[genre] += plays[i]
        if hash[genre] != nil {
            h := hash[genre]
            h.Push(NewGenre(i, genre, plays[i]))
            hash[genre] = h
        } else {
            h := &GenreHeap{}
            h.Push(NewGenre(i, genre, plays[i]))
            hash[genre] = h
        }
    }

    ranking := NewRanking(rank_hash)
    sort.Sort(ranking)
    for genre := range hash {
        sort.Sort(hash[genre])
    }

    for _, genre := range ranking.rank {
        h := hash[genre]
        if h.Len() == 1 {
            result = append(result, h.Pop(0).id)
        } else {
            result = append(result, h.Pop(0).id)
            result = append(result, h.Pop(1).id)
        }
    }
    return result
}
```