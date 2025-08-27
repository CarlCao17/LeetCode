package lru_test

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

// 基准测试：Put操作
func BenchmarkLRU_Put(b *testing.B) {
	cache, _ := lru.New[int, string](1000)
	
	b.ResetTimer()
	b.ReportAllocs()
	
	for i := 0; i < b.N; i++ {
		cache.Put(i, strconv.Itoa(i))
	}
}

// 基准测试：Get操作（全命中）
func BenchmarkLRU_Get_Hit(b *testing.B) {
	cache, _ := lru.New[int, string](1000)
	
	// 预填充
	for i := 0; i < 1000; i++ {
		cache.Put(i, strconv.Itoa(i))
	}
	
	b.ResetTimer()
	b.ReportAllocs()
	
	for i := 0; i < b.N; i++ {
		cache.Get(i % 1000)
	}
}

// 基准测试：Get操作（全未命中）
func BenchmarkLRU_Get_Miss(b *testing.B) {
	cache, _ := lru.New[int, string](1000)
	
	b.ResetTimer()
	b.ReportAllocs()
	
	for i := 0; i < b.N; i++ {
		cache.Get(i + 2000) // 确保未命中
	}
}

// 基准测试：混合操作
func BenchmarkLRU_Mixed(b *testing.B) {
	cache, _ := lru.New[int, string](1000)
	
	// 预填充
	for i := 0; i < 500; i++ {
		cache.Put(i, strconv.Itoa(i))
	}
	
	b.ResetTimer()
	b.ReportAllocs()
	
	for i := 0; i < b.N; i++ {
		switch i % 4 {
		case 0, 1: // 50% Get操作
			cache.Get(rand.Intn(1000))
		case 2: // 25% Put操作
			cache.Put(rand.Intn(2000), strconv.Itoa(i))
		case 3: // 25% Contains操作
			cache.Contains(rand.Intn(1000))
		}
	}
}

// 基准测试：不同容量的性能
func BenchmarkLRU_DifferentCapacities(b *testing.B) {
	capacities := []int{10, 100, 1000, 10000}
	
	for _, capacity := range capacities {
		b.Run(fmt.Sprintf("Cap_%d", capacity), func(b *testing.B) {
			cache, _ := lru.New[int, string](capacity)
			
			b.ResetTimer()
			b.ReportAllocs()
			
			for i := 0; i < b.N; i++ {
				cache.Put(i, strconv.Itoa(i))
			}
		})
	}
}

// 基准测试：随机访问模式
func BenchmarkLRU_RandomAccess(b *testing.B) {
	cache, _ := lru.New[int, string](1000)
	
	// 预填充
	for i := 0; i < 1000; i++ {
		cache.Put(i, strconv.Itoa(i))
	}
	
	// 生成随机访问序列
	rand.Seed(time.Now().UnixNano())
	keys := make([]int, b.N)
	for i := range keys {
		keys[i] = rand.Intn(1000