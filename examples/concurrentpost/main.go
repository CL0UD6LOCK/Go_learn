package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"
)

func postStudent(name string, wg *sync.WaitGroup) {
	defer wg.Done()
	body := fmt.Sprintf(`{"name":%q,"age":20}`, name)
	fmt.Println(name, "开始发送请求")

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Post(
		"http://localhost:8080/students",
		"application/json",
		strings.NewReader(body),
	)

	fmt.Println(name, "请求已结束")
	if err != nil {
		fmt.Println(name, "请求失败", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		fmt.Println(name, "HTTP状态码:", resp.StatusCode)
		return
	}
	var result struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		fmt.Println("解析响应失败:", err)
		return
	}
	fmt.Printf("%s -> HTTP %d, ID=%d\n", name, resp.StatusCode, result.ID)
}

func getStudents(wg *sync.WaitGroup) {
	defer wg.Done()

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(
		"http://localhost:8080/students",
	)
	if err != nil {
		fmt.Println("GET 请求失败:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("GET 状态码:", resp.StatusCode)
}

func patchStudent(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	url := fmt.Sprintf(
		"http://localhost:8080/students/%d", id,
	)

	req, err := http.NewRequest(
		http.MethodPatch,
		url,
		strings.NewReader(`{"age":21}`),
	)
	if err != nil {
		fmt.Println("创建 PATCH 请求失败:", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("PATCH 请求失败:", err)
		return
	}

	defer resp.Body.Close()

	fmt.Println("PATCH 状态码:", resp.StatusCode)
}

func deleteStudent(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	url := fmt.Sprintf(
		"http://localhost:8080/students/%d", id,
	)

	req, err := http.NewRequest(
		http.MethodDelete,
		url,
		nil,
	)
	if err != nil {
		fmt.Println("创建DELETE请求失败:", err)
		return
	}
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("DELETE 请求失败:", err)
		return
	}

	defer resp.Body.Close()
	fmt.Println("DELETE 状态码:", resp.StatusCode)

}

func getStudentByID(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	url := fmt.Sprintf(
		"http://localhost:8080/students/%d", id,
	)

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)

	if err != nil {
		fmt.Println("GET 请求失败:", err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("GET 状态码:", resp.StatusCode)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	go deleteStudent(37, &wg)
	go getStudentByID(37, &wg)

	wg.Wait()

	fmt.Println("DELETE + GET 测试完成")
}
